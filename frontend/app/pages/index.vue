<script setup>

    import TaskInput from "@/components/ui/taskInput.vue";

    const config = useRuntimeConfig();

    const { data, pending, error } = await useAsyncData('tasks', () =>
        $fetch('/tasks/all-unfinished', {
            baseURL: config.apiBase ?? "http://localhost:8090"
            }
        )
    )

</script>

<template>
        <div v-if="pending">Loading...</div>
        <div v-else-if="error">Error: {{ error.message }}</div>
        <main v-else>
            <TaskInput />
            <div id="task-container">
                <h3>Tarefas</h3>
                <UiCard v-for="task in data.tarefas" :key="task.id" :descricao="task.descricao" :criadaEm="task['criada_em']" :concluida="task.concluida" />
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