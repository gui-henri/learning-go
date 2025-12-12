<script setup>
import { ref, computed } from 'vue';
import { useCuidadoresStore } from '@/store/evaluation/cuidadores';
import { InputMask } from 'primevue';
import { useStepAccordion } from "@/composable/useStepAccordion";

const cuidadorStore = useCuidadoresStore();

const props = defineProps({
  formFields: {
    type: Object,
    required: false,
    default: null
  },
  isActive: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['next-step']);

const isFilled = computed(() => {
    return !!cuidadorStore.cuidadores.medico_solicitante && 
           cuidadorStore.cuidadores.medico_solicitante.length > 3;
});

const { internalIndex, nextStep } = useStepAccordion(props, emit);

const handleSave = () => {
    internalIndex.value = null;
    
    setTimeout(() => {
        const self = document.getElementById("cuidadores");
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
    <Accordion v-model:activeIndex="internalIndex" id="cuidadores" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
                </i>
                <div class="flex flex-col text-left">
                    <h4 class="font-semibold text-xl p-0 m-0">Cuidadores</h4>
                    <span class="text-xs text-gray-500 font-normal -mt-4">
                        {{ isFilled ? '' : 'Toque para preencher' }}
                    </span>
                </div>
            </div>
        </template>
        <div class="flex flex-col gap-2 w-full md:w-1/6">
            <label for="possui_cuidador">Possui cuidador?</label>
            <InputSwitch 
                id="possui_cuidador" 
                v-model="cuidadorStore.cuidadores.possui_cuidador" 
                class="w-full"
            ></InputSwitch>
        </div>
        <template v-if="cuidadorStore.cuidadores.possui_cuidador">   
            <hr class="border-gray-200 my-2"/>
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/2">
                    <label for="nome_cuidador">Nome do cuidador</label>
                    <InputText 
                        id="nome_cuidador" 
                        v-model="cuidadorStore.cuidadores.nome_cuidador" 
                        placeholder="Nome completo" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/4">
                    <label for="contato_cuidador">Contato do cuidador</label>
                    <InputMask
                        id="contato_cuidador" 
                        v-model="cuidadorStore.cuidadores.contato_cuidador" 
                        placeholder="(81) 9 9999-9999" 
                        class="w-full"
                        mask="(99) 9 9999-9999"     
                        
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/4">
                    <label for="turno_cuidador">Turno</label>
                    <Select 
                        id="turno_cuidador" 
                        v-model="cuidadorStore.cuidadores.turno_cuidador" 
                        :options="formFields.turno_cuidador" 
                        optionLabel="label"
                        optionValue="label"
                        placeholder="Selecione o turno" 
                        class="w-full"
                    ></Select>
                </div>
            </div>
            <div class="flex flex-col md:flex-row gap-4 items-center mt-6 min-h-16">
                <div class="flex flex-col gap-2 w-full md:w-1/5">
                    <label for="precisa_treinamento">Precisa de treinamento?</label>
                    <InputSwitch 
                        id="precisa_treinamento" 
                        v-model="cuidadorStore.cuidadores.precisa_treinamento" 
                        class="w-full"
                    ></InputSwitch>
                </div>
                <div v-if="cuidadorStore.cuidadores.precisa_treinamento" class="flex flex-col gap-2 w-full md:w-4/6 flex-1 transition-all duration-300">
                    <label for="obs_treinamento">Observações do treinamento</label>
                    <InputText 
                        id="obs_treinamento" 
                        v-model="cuidadorStore.cuidadores.obs_treinamento" 
                        placeholder="Descreva a necessidade..." 
                        class="w-full"
                    />
                </div>
            </div>
        </template>
        <hr class="border-gray-200 my-2"/>
        <div class="flex flex-col gap-4 w-full">
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/2">
                    <label for="medico_solicitante">Médico solicitante</label>
                    <InputText 
                        id="medico_solicitante" 
                        v-model="cuidadorStore.cuidadores.medico_solicitante" 
                        placeholder="Nome do médico" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/2">
                    <label for="contato_medico">Contato do médico</label>
                    <InputMask
                        id="contato_medico" 
                        v-model="cuidadorStore.cuidadores.contato_medico" 
                        placeholder="(XX) 9 9999-9999" 
                        class="w-full"
                        mask="(99) 9 9999-9999"
                    />
                </div>
            </div>
        </div>
        <Button class="mt-3" v-on:click="handleSave">
            <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
            Próximo
        </Button>
    </AccordionTab>
</Accordion>
</template>