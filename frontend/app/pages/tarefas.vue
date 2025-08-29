<template>
  <h1 style="text-align: center;">Tarefas Concluídas</h1>

</template>

  <script setup>

  
    const config = useRuntimeConfig();

    const { data, pending, error } = await useAsyncData('tasks', () =>
        $fetch('/task/all', {
            baseURL: config.apiBase ?? "http://localhost:8090"
        }
    )
)
</script>