<script setup lang="ts">

    interface InsertResponse {
        tarefa: {
            id: number,
            descricao: string,
            prazo: string,
            concluida: boolean,
            criada_em: string
        }
    }

    const config = useRuntimeConfig();
    const descricao = ref('');
    const emit = defineEmits(["taskSended"])

    async function sendTask() {
        try {
            const response = await fetch(config.public.apiBase + "/tasks", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ descricao: descricao.value, prazo: "amanhã" })
            })

            if (!response.ok) {
                console.log(response.status)
                console.log(response.url)
                throw new Error(String(response.status))
            }

            const responseJson: InsertResponse = await response.json()

            emit("taskSended", responseJson.tarefa);

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