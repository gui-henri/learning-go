<script setup lang="ts">
    const config = useRuntimeConfig();
    const descricao = ref('');

    async function sendTask() {
        try {
            const response = await fetch(config.apiBase + "/tasks", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ descricao: descricao.value, prazo: "amanhã" })
            })

            if (!response.ok) {
                throw new Error("Resposta diferente de ok")
            }

        } catch (error) {
            console.log("Erro ao realizar requisição", error)
        }
    }
</script>

<template>
    <div class="items-center justify-center gap-2">
        <Input v-model="descricao" class="p-2 rounded-full border-2 border-red-500" type="text" name="description" placeholder="Digite a tarefa" />
        <Button @click="sendTask" class="bg-red-600 hover:bg-red-500 p-5 rounded-3xl">Enviar</Button>
    </div>
</template>

<style lang="css" scoped>
    
    div {
        display: flex;
        width: 100%;
        max-width: 520px;
    }
    input {
        width: 100%;
    }

</style>