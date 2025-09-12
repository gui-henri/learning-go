<script setup lang="ts">
    import { SquareCheckBig, LoaderIcon, PartyPopperIcon, Trash2 } from "lucide-vue-next"
    import { Button } from "@/components/ui/button"

    const props = defineProps<{
        id: number
        descricao: string
        criadaEm: string
        concluida: boolean
        prazo: string
    }>()

    const emit = defineEmits(["taskFinished"])

    const formattedDate = (() => {
        try {
            return new Intl.DateTimeFormat("pt-BR", {
                day: "2-digit",
                month: "2-digit",
                year: "numeric"
            }).format(new Date(props.criadaEm))
        } catch (err) {
            console.error("Invalid criadaEm date:", props.criadaEm, err)
            return "Data inválida"
        }
    })()


    const formatedPrazo = (() => {
        try {
            return new Intl.DateTimeFormat("pt-BR", {
                day: "2-digit",
                month: "2-digit",
                year: "numeric"
            }).format(new Date(props.prazo))
        } catch (err) {
            console.error("Invalid criadaEm date:", props.criadaEm, err)
            return "Indefinido"
        }
    })()


    async function finishTask() {
        try {
            const res = await fetch("http://localhost:8090/tasks/finish", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ id: props.id })
            });

            if (!res.ok) throw new Error("Failed to mark task as complete");

            emit("taskFinished", props.id);
        } catch (err) {
            console.error("Error finishing task:", err);
        }
    }

    async function deleteTask() {
        try {
            const res = await fetch("http://localhost:8090/tasks", {
                method: "DELETE",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ id: props.id })
            });

            if (!res.ok) throw new Error("Failed to delete task");

            emit("taskFinished", props.id);
        } catch (error) {
            console.error("Error finishing task:", error);
        }
    }


</script>

<template>
    <div class="card-body pl-3 pr-3 pt-2 pb-2 rounded-2xl border-2 border-b-red-600 border-l-red-400 border-r-red-600 border-t-red-400">
        <div>
            <h3 class="font-bold text-xl">{{ descricao }}</h3>
            <p>{{ formatedPrazo === "Indefinido" ? "Sem prazo" : formatedPrazo}}</p>
            <h4>Criada em: {{ formattedDate }}</h4>
        </div>
        <div class="gap-2 flex">
            <Button v-if="concluida === false" class="bg-neutral-700"><LoaderIcon/> Em andamento</Button>
            <Button v-else class="bg-green-500"> <PartyPopperIcon /> Concluída</Button>
            <Button v-if="concluida === false" @click="finishTask" size="icon" class="bg-green-600 hover:bg-green-500">
                <SquareCheckBig class="w-4 h-4 text-green-200" />
            </Button>
            <Button @click="deleteTask" size="icon" class="bg-red-600 hover:bg-red-500">
                <Trash2 class="w-4 h-4 text-white" />
            </Button>
        </div>
    </div>
</template>

<style lang="css" scoped>
    .card-body {
        width: 100%;
        max-width: 640px;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    h3, h4, p {
        padding: 0;
        margin: 0;
    }
</style>