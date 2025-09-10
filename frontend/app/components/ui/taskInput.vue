<script setup lang="ts">
    import type { DateValue } from "@internationalized/date"
    import {
        DateFormatter,
        getLocalTimeZone,
        today
    } from "@internationalized/date"
    import { CalendarIcon, SendIcon } from "lucide-vue-next"
    import { Popover, PopoverTrigger, PopoverContent } from "@/components/ui/popover"
    import { Calendar } from "@/components/ui/calendar"

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

    const df = new DateFormatter("pt-BR", {})

    const dateValue = ref<DateValue | null>(null)
    watch(dateValue, (val: any) => console.log(val))
    const todayDate = today(getLocalTimeZone())

    async function sendTask() {
        try {
            const response = await fetch(config.public.apiBase + "/tasks", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ 
                    descricao: descricao.value, 
                    prazo: dateValue.value ? 
                        dateValue.value.toDate(getLocalTimeZone()).toISOString() 
                        : "sem prazo" 
                })
            })

            if (!response.ok) {
                console.log(response.status)
                console.log(response.url)
                throw new Error(String(response.status))
            }

            const responseJson: InsertResponse = await response.json()

            emit("taskSended", responseJson.tarefa);
            descricao.value = ""
            dateValue.value = null;

        } catch (error) {
            console.log("Erro ao realizar requisição", error)
        }
    }
</script>

<template>
    <div class="items-center justify-center gap-2">
        <div class="p-0.5 rounded-full border-2 border-red-500">
            <Input v-model="descricao" class="p-3 border-0 focus-visible:border-0 focus-visible:ring-ring/0 focus-visible:ring-[3px]" type="text" name="description" placeholder="Digite a tarefa" />
            <Popover>
                <PopoverTrigger as-child>
                    <Button variant="outline"
                        class="w-34 bg-amber-200 hover:bg-red-400 transition justify-start text-left font-normal border-0 rounded-full"
                    >
                        <CalendarIcon />
                        {{ dateValue ? df.format(dateValue?.toDate(getLocalTimeZone())) : "Definir prazo" }}
                    </Button>
                </PopoverTrigger>
                <PopoverContent class="w-auto p-0">
                    <Calendar locale="pt-BR" :minValue="todayDate" class="flex flex-col" v-model="dateValue" initial-focus />
                </PopoverContent>
            </Popover>
        </div>
        <Button @click="sendTask" class="bg-red-600 hover:bg-red-500 p-5 rounded-2xl cursor-pointer"><SendIcon /></Button>
    </div>
</template>

<style lang="css" scoped>
    
    div {
        display: flex;
        width: 100%;
        max-width: 640px;
    }
    input {
        width: 100%;
    }

    :deep([data-selected]) {
        background-color: #dc2626; /* bg-red-600 */
        color: white;
        font-weight: 700;
    }

    :deep([data-selected]:hover) {
        background-color: #b91c1c; /* bg-red-700 */
    }

</style>