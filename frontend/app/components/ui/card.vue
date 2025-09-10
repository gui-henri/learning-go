<script setup lang="ts">
    import { SquareCheckBig } from "lucide-vue-next"
    import { Button } from "@/components/ui/button"

    const props = defineProps<{
        id: number
        descricao: string
        criadaEm: string
        concluida: boolean
    }>()

    const emit = defineEmits(["taskFinished"])

    const formattedDate = new Intl.DateTimeFormat("pt-BR", {
        day: "2-digit",
        month: "2-digit",
        year: "numeric"
    }).format(new Date(props.criadaEm))


    async function finishTask() {
        try {
            const res = await fetch("http://localhost:8090/tasks/finish", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ id: props.id })
            });

            if (!res.ok) throw new Error("Failed to mark task as complete");

            console.log("emmiting for", props.id)
            emit("taskFinished", props.id);
        } catch (err) {
            console.error("Error finishing task:", err);
        }
    }


</script>

<template>
    <div class="card-body pl-3 pr-3 pt-2 pb-2 rounded-2xl border-2 border-b-red-600 border-l-red-400 border-r-red-600 border-t-red-400">
        <div>
            <h3>{{ descricao }}</h3>
            <h4>Criada em {{ formattedDate }}</h4>
            <p>{{ concluida ? "Concluída" : "Em andamento" }}</p>
        </div>
        <div>
            <Button @click="finishTask" size="icon" class="bg-red-600 hover:bg-red-500">
                <SquareCheckBig class="w-4 h-4 text-red-200" />
            </Button>
           <Button @click="deleteTask" size="icon" class="bg-gray-600 hover:bg-gray-500">
            <Trash2 class="w-4 h-4 text-white" />
            </Button>
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
    }

    h3, h4, p {
        padding: 0;
        margin: 0;
    }
</style>