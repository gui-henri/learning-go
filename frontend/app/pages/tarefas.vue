<script setup>
    import Card from "@/components/ui/card.vue";

    const config = useRuntimeConfig();

    const { data, pending, error, refresh } = await useAsyncData('tasks/all', () =>
        $fetch('/tasks/all', {
            baseURL: config.apiBase ?? "http://localhost:8090"
            }
        )
    )

    async function removeTask() {
        await refresh()
    }

</script>

<template>
    <main class="gap-2">
        <h3 class="text-3xl mb-5 font-bold">Tarefas</h3>
        <div v-if="pending">Loading...</div>
        <div v-else-if="error">Error: {{ error.message }}</div>    
        <div v-else id="task-container">
            <Card v-for="task in data.tarefas.slice().reverse()" :key="task.id" :id="task.id" :descricao="task.descricao" :prazo="task.prazo" :criadaEm="task['criada_em']" :concluida="task.concluida" @taskFinished="removeTask" />
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