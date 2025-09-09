<script setup>
    import TaskInput from "@/components/ui/taskInput.vue";
    import Card from "@/components/ui/card.vue";

    const config = useRuntimeConfig();

    const { data, pending, error } = await useAsyncData('tasks', () =>
        $fetch('/tasks/all-unfinished', {
            baseURL: config.apiBase ?? "http://localhost:8090"
            }
        )
    )

    function removeTask(taskId) {
        data.value = { ...data.value, tarefas: data.value.tarefas.filter(task => task.id !== taskId) }
    }

    function addTask(tarefa) {
        data.value.tarefas =  [...data.value.tarefas, tarefa]
    }

</script>

<template>
    <main class="gap-2">
        <h3 class="text-3xl mb-5 font-bold">Tarefas</h3>
        <TaskInput @taskSended="addTask" />
        <div v-if="pending">Loading...</div>
        <div v-else-if="error">Error: {{ error.message }}</div>    
        <div v-else id="task-container">
            <Card v-for="task in data.tarefas" :key="task.id" :id="task.id" :descricao="task.descricao" :criadaEm="task['criada_em']" :concluida="task.concluida" @taskFinished="removeTask" />
        </div>
    </main>
</template>

<style scoped>
    main {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        width: 100%;
    }

    #task-container {
        width: 100%;
        max-width: 520px;
        gap: 16px;
        display: flex;
        flex-direction: column;
    }
</style>