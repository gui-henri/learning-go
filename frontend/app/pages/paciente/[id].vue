<script setup>
import { useRoute } from 'vue-router';
import { ArrowLeft } from "lucide-vue-next";
import { Button } from '~/components/ui/button';
import { Popover } from '~/components/ui/popover';
import PopoverTrigger from '~/components/ui/popover/PopoverTrigger.vue';
import PopoverContent from '~/components/ui/popover/PopoverContent.vue';

const route = useRoute();
const { id } = route.params;

const config = useRuntimeConfig();

const { data, pending, error } = await useAsyncData(`patient/${id}`, () =>
  $fetch(`/Patient/${id}`, {
    baseURL: config.apiBase ?? "http://localhost:8090"
  })
);

const { data: tarefasData, pending: tarefasPending, error: tarefasError } = await useAsyncData('tasks/all', () =>
  $fetch('/tasks/all', {
    baseURL: config.apiBase ?? "http://localhost:8090"
  })
)

const tarefas = computed(() => {
  if (!tarefasData.value || !tarefasData.value.tarefas) return []
  return tarefasData.value.tarefas.filter(t => t.paciente.id === id)
})

async function deadPatient(id, patient) {
  try {
    // 1. Espera a requisição para o backend finalizar
    await $fetch(`/Patient/${id}`, {
      method: 'PUT',
      baseURL: config.apiBase ?? "http://localhost:8090",
      body: {
        ...patient,
        active: false,
        deceasedBoolean: true
      }
    });

    await navigateTo('/pacientes');

  } catch (error) {

    console.error("Erro ao atualizar o paciente:", error);

  }
}
</script>

<template>
  <div>
    <div v-if="pending" class="min-h-screen flex items-center justify-center">
      <p class="text-xl text-gray-500 animate-pulse">Carregando informações do paciente...</p>
    </div>

    <div v-else-if="error" class="min-h-screen flex items-center justify-center p-4">
       <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg text-center" role="alert">
        <strong>Ocorreu um erro ao buscar o paciente:</strong>
        <p>{{ error.message || 'Verifique se o backend está rodando e o ID é válido.' }}</p>
        <NuxtLink to="/pacientes" class="text-blue-600 hover:underline mt-4 inline-block">
          &larr; Voltar para a lista
        </NuxtLink>
      </div>
    </div>

    <div v-else-if="data" class="flex flex-col md:flex-row items-start justify-center p-4 sm:p-6 gap-x-8">
      
      <div class="bg-white shadow-2xl rounded-2xl w-full md:w-1/2 p-6 sm:p-8 border-t-8 border-red-600 mb-8 md:mb-0">
        
        <header class="relative w-full text-center mb-8">
          <NuxtLink 
            to="/pacientes" 
            class="absolute left-0 top-1/2 -translate-y-1/2 bg-red-600 hover:bg-red-700 text-white p-3 rounded-full transition-colors"
            title="Voltar"
          >
            <ArrowLeft class="w-5 h-5" />
          </NuxtLink>
          <h1 class="text-3xl font-bold text-gray-800 inline-block">
            Ficha do Paciente
          </h1>
        </header>

        <main class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-8 pt-4 border-t border-gray-200">
          
          <div>
            <h3 class="text-sm font-semibold text-gray-600">Nome Completo</h3>
            <p class="text-lg text-gray-900 mt-1">{{ data.name[0].given.join(' ') }}</p>
          </div>

          <div>
            <h3 class="text-sm font-semibold text-gray-600">CPF</h3>
            <p class="text-lg text-gray-900 mt-1">{{ data.identifier[0].value }}</p>
          </div>

          <div>
            <h3 class="text-sm font-semibold text-gray-600">Data de Nascimento</h3>
            <p class="text-lg text-gray-900 mt-1">
              {{ new Date(data.birthDate).toLocaleDateString('pt-BR', { timeZone: 'UTC' }) }}
            </p>
          </div>
          
          <div>
            <h3 class="text-sm font-semibold text-gray-600">Gênero</h3>
            <p class="text-lg text-gray-900 mt-1 capitalize">{{ data.gender }}</p>
          </div>

          <div>
            <h3 class="text-sm font-semibold text-gray-600">Telefone</h3>
            <p class="text-lg text-gray-900 mt-1">{{ data.telecom.find(t => t.system === 'phone')?.value || 'Não informado' }}</p>
          </div>
          
          <div>
            <h3 class="text-sm font-semibold text-gray-600">Email</h3>
            <p class="text-lg text-gray-900 mt-1">{{ data.telecom.find(t => t.system === 'email')?.value || 'Não informado' }}</p>
          </div>

          <Popover>
            <PopoverTrigger>
              <Button class="bg-gray-600 hover:bg-gray-700">Registrar falecimento</Button>
            </PopoverTrigger>
            <PopoverContent class="flex flex-col gap-2">
              <b>Você tem certeza?</b>
              <Button @click="() => deadPatient(id, data)" class="bg-red-600 hover:bg-red-400">Sim</Button>
            </PopoverContent>
          </Popover>

        </main>
      </div>

      <div class="bg-white shadow-2xl rounded-2xl w-full md:w-1/3 p-6 sm:p-8">
        <h2 class="text-2xl font-bold text-gray-800 mb-4 border-b pb-2">Tarefas Associadas</h2>
        <div v-if="tarefasPending">
          <p class="text-gray-500">Carregando tarefas...</p>
        </div>
        <div v-else-if="tarefasError">
          <p class="text-red-500">Erro ao buscar as tarefas.</p>
        </div>
        <ul v-else-if="tarefas && tarefas.length > 0" class="space-y-4 max-h-96 overflow-y-auto">
          <li v-for="tarefa in tarefas" :key="tarefa.id" class="p-3 rounded-lg border" :class="tarefa.concluida ? 'bg-green-50 border-green-200' : 'bg-orange-50 border-orange-200'">
            <p class="font-semibold">{{ tarefa.descricao }}</p>
            <p class="text-sm text-gray-600">
              Prazo: {{ tarefa.prazo ? new Date(tarefa.prazo).toLocaleDateString('pt-BR', {timeZone: 'UTC'}) : 'Não definido' }}
            </p>
            <span class="text-xs font-medium px-2 py-0.5 rounded-full" :class="tarefa.concluida ? 'bg-green-100 text-green-800' : 'bg-orange-100 text-orange-800'">
              {{ tarefa.concluida ? 'Concluída' : 'Pendente' }}
            </span>
          </li>
        </ul>
        <div v-else class="text-center p-6 bg-gray-50 rounded-lg">
          <p class="text-gray-600">Nenhuma tarefa encontrada para este paciente.</p>
        </div>
      </div>

    </div>
  </div>
</template>