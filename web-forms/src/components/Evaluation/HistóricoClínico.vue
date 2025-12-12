<script setup>
import { ref, computed } from 'vue';
import { useHistoricoStore } from '@/store/evaluation/historico';
import { useStepAccordion } from "@/composable/useStepAccordion";

const historicoStore = useHistoricoStore();

const props = defineProps({
  isActive: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['next-step']);

const isFilled = computed(() => {
    return !!historicoStore.historico.historia_doenca_atual &&
            historicoStore.historico.historia_doenca_atual.length > 3;
});

const { internalIndex, nextStep } = useStepAccordion(props, emit);

const handleSave = () => {
    internalIndex.value = null;
    
    setTimeout(() => {
        const self = document.getElementById("historico");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    nextStep('next-step');
};

</script>

<template>
    <Accordion v-model:activeIndex="internalIndex" id="historico" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
                </i>
                <div class="flex flex-col text-left">
                    <h4 class="font-semibold text-xl p-0 m-0">Histórico e Planejamento</h4>
                    <span class="text-xs text-gray-500 font-normal -mt-4">
                        {{ isFilled ? '' : 'Toque para preencher' }}
                    </span>
                </div>
            </div>
        </template>
        <div class="flex flex-col gap-4 w-full">
                <div class="flex flex-col gap-2 w-full">
                    <label for="historia_doenca_atual">História da doença atual</label>
                    <Textarea 
                        id="historia_doenca_atual" 
                        v-model="historicoStore.historico.historia_doenca_atual" 
                        rows="4" 
                        placeholder="Descreva a história clínica mais atual..." 
                        class="w-full" 
                        autoResize
                    />
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="anamnese_enfermagem">Anamnese de enfermagem</label>
                    <Textarea 
                        id="anamnese_enfermagem" 
                        v-model="historicoStore.historico.anamnese_enfermagem" 
                        rows="4" 
                        placeholder="Dados sobre o histórico de saúde..." 
                        class="w-full" 
                        autoResize
                    />
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="plano_terapeutico">Plano terapêutico</label>
                    <Textarea 
                        id="plano_terapeutico" 
                        v-model="historicoStore.historico.plano_terapeutico" 
                        rows="4" 
                        placeholder="Planejamento do cuidado..." 
                        class="w-full" 
                        autoResize
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