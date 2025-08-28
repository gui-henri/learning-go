<template>
    <div class="card-body">
        <div>
            <h3>{{ descricao }}</h3>
            <h4>{{ criadaEm }}</h4>
            <p>{{ concluida ? "Concluída" : "Em andamento" }}</p>
        </div>
        <div>
            <button>Concluir</button>
        </div>
    </div>
</template>

<style lang="css" scoped>
    .card-body {
        display: flex;
        align-items: center;
        width: 300px;
        justify-content: space-between;
    }

    h3, h4, p {
        padding: 0;
        margin: 0;
    }
</style>

<script setup>
defineProps({
  descricao: String,
  criadaEm: String,
  concluida: Boolean
})
</script>