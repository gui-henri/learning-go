<script setup>
import { useObservacoesStore } from '@/store/evaluation/observacao';
import { computed } from 'vue';
import { useStepAccordion } from "@/composable/useStepAccordion";

const props = defineProps({
  isActive: {
    type: Boolean,
    default: false
  }
})

const observacoesStore = useObservacoesStore();
const emit = defineEmits(['next-step']);

const { internalIndex, nextStep } = useStepAccordion(props, emit);

    const handleSave = () => {
        internalIndex.value = null;
        setTimeout(() => {
            const self = document.getElementById("observacoes");
            if (self) {
                self.scrollIntoView({ 
                    behavior: 'instant', 
                    block: 'start',
                });
            }
        }, 0);
        nextStep()
    };

const isFilled = computed(() => {
    return !!observacoesStore.observacoes.texto && 
           observacoesStore.observacoes.texto.length > 3;
});

</script>
<template>
    <Accordion v-model:activeIndex="internalIndex" id="observacoes" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Observações <i class="pi pi-link text-3xl text-black"></i></h4>
                <span class="text-xs text-gray-500 font-normal -mt-4">
                    {{ isFilled ? '' : 'Toque para preencher' }}
                </span>
            </div>
        </div>
    </template>

        <div class="flex flex-col gap-4 w-full">
            <div class="flex flex-col gap-2 w-full">
                <label for="observacoes_gerais">Anotações adicionais</label>
                <Textarea 
                    id="observacoes_gerais" 
                    v-model="observacoesStore.observacoes.texto" 
                    rows="10" 
                    placeholder="Descreva aqui qualquer observação adicional relevante sobre o paciente, condições do domicílio ou detalhes que não se encaixaram nos campos anteriores..." 
                    class="w-full" 
                    autoResize
                    maxlength="300"
                />
            </div>

        </div>

<Button class="mt-3" v-on:click="handleSave">
        <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
        Próximo
    </Button>
    </AccordionTab>
</Accordion>

</template>