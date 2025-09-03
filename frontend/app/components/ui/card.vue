<template>
    <div class="card-body">
        <div>
            <h3>{{ descricao }}</h3>
            <h4>{{ `Criada em: ${new Date(criadaEm).toLocaleDateString()}` }}</h4>
            <p>{{ concluida ? "Concluída" : "Em andamento" }}</p>
        </div>
        <div>
            <button>Concluir</button>
        </div>
    </div>
</template>

<style lang="css" scoped>
    .card-body {
        width: 100%;
        max-width: 520px;
        display: flex;
        align-items: center;
        justify-content: space-between;

        border: 2px solid black;
        box-shadow: 2px 2px 10px gray;
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