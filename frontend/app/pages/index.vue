<script setup>
    const { data, pending, error } = await useFetch('http://backend:8090/task/all-active')
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