<script setup>
    const config = useRuntimeConfig();

    const { data, pending, error } = await useAsyncData('tasks', () =>
        $fetch('/task/all-active', {
            baseURL: config.apiBase ?? "http://localhost:8090"
        }
    )
)

</script>

<template>
    INPUT DE NOVAS TAREFAS
    ----------------
    <div v-if="pending">Loading...</div>
    <div v-else-if="error">Error: {{ error.message }}</div>
    <div>
        <UiCard v-for="task in data" :key="task.id" :descricao="task.descricao" :criadaEm="task['criada-em']" :concluida="task.concluida" />
    </div>
</template>