<template>
  <main class="w-full bg-gray-50 min-h-screen py-10 px-4">
    <div class="mx-auto w-full max-w-xl">
      <h3 class="text-3xl font-bold text-gray-800 text-center mb-5">
        Requisite sua tarefa abaixo
      </h3>
      <div class="mb-12">
        <TaskInput @taskSended="handleTaskChange" />
      </div>

      <div v-if="pending" class="text-center text-gray-500">
        <p>Carregando tarefas...</p>
      </div>
      <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg text-center" role="alert">
        <strong>Ocorreu um erro:</strong>
        <p>{{ error.message }}</p>
      </div>

      <div v-else class="space-y-12">
        <section>
          <div class="flex items-center gap-x-3 mb-4">
            <h4 class="text-red-700 text-2xl font-semibold">Tarefas Pendentes</h4>
            <span class="bg-red-100 text-red-800 text-xs font-medium px-2.5 py-1 rounded-full">{{ pendingTasks.length }}</span>
          </div>
          <div v-if="pendingTasks.length > 0" class="space-y-4">
            <Card
              v-for="task in pendingTasks"
              :key="task.id"
              :id="task.id"
              :descricao="task.descricao"
              :prazo="task.prazo"
              :criadaEm="task['criada_em']"
              :concluida="task.concluida"
              @taskFinished="handleTaskChange"
            />
          </div>
          <div v-else class="bg-white text-center p-8 rounded-lg border border-dashed border-gray-300">
            <p class="font-medium text-gray-600"><b>Nenhuma tarefa pendente. Continue assim !</b> 👍</p>
          </div>
        </section>

        <section>
          <div class="flex items-center gap-x-3 mb-4">
            <h4 class="text-green-600 text-2xl font-semibold">Tarefas Concluídas</h4>
            <span class="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-1 rounded-full">{{ completedTasks.length }}</span>
          </div>
          <div v-if="completedTasks.length > 0" class="space-y-4">
            <Card
              v-for="task in completedTasks"
              :key="task.id"
              :id="task.id"
              :descricao="task.descricao"
              :prazo="task.prazo"
              :criadaEm="task['criada_em']"
              :concluida="task.concluida"
              @taskFinished="handleTaskChange"
            />
          </div>
          <div v-else class="bg-white text-center p-8 rounded-lg border border-dashed border-gray-300">
            <p class="font-medium text-gray-600"><b>Aguardando conclusão de tarefas.</b></p>
          </div>
        </section>
      </div>
    </div>
  </main>
</template>

<script setup>
import Card from "@/components/ui/card.vue";
import TaskInput from "@/components/ui/taskInput.vue"; // Importação adicionada
import { computed } from 'vue';

const config = useRuntimeConfig();

// Busca TODAS as tarefas de uma única vez
const { data, pending, error, refresh } = await useAsyncData('tasks/all', () =>
  $fetch('/tasks/all', { // Endpoint unificado para buscar todas as tarefas
    baseURL: config.apiBase ?? "http://localhost:8090"
  })
)

async function handleTaskChange() {
  await refresh();
}

// Propriedades computadas para separar as tarefas em pendentes e concluídas
const pendingTasks = computed(() => {
  const tasks = data.value?.tarefas || [];
  // Exibe as tarefas pendentes mais recentes primeiro
  return tasks.filter(t => !t.concluida).slice().reverse();
});

const completedTasks = computed(() => {
  const tasks = data.value?.tarefas || [];
  // Exibe as tarefas concluídas mais recentes primeiro
  return tasks.filter(t => t.concluida).slice().reverse();
});
</script>