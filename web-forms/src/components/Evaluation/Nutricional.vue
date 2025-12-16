<script setup>
import { ref, computed, watch } from 'vue'; 
import { useNutricionalStore } from '@/store/evaluation/nutricional';
import { InputMask } from 'primevue';
import { useStepAccordion } from "@/composable/useStepAccordion";
import { useAvaliationForm } from "@/store/evaluation/form";
import { AvaliationService } from '@/service/AvaliationService';

const avaliationFormStore = useAvaliationForm();
const nutricionalStore = useNutricionalStore();

const emit = defineEmits(['next-step']);

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

const isFilled = computed(() => {
    return !!nutricionalStore.nutricional.alimentacao_oral;
});

watch(() => nutricionalStore.nutricional.via_enteral, (novoValor) => {
    if (novoValor === 'Gastrostomia' || novoValor === 'Nasoenteral') {
        nutricionalStore.nutricional.sonda = true;
    } else {
        nutricionalStore.nutricional.sonda = false;
    }
});

const { internalIndex, nextStep } = useStepAccordion(props, emit);

const handleSave = async () => {
    const payload = {
        nutricional: nutricionalStore.nutricional,
    };
    if (avaliationFormStore.avaliationId === null) {
        await AvaliationService.submitFormData(payload)
    } else {
        await AvaliationService.appendToAvaliation(
            avaliationFormStore.avaliationId, 
            payload
        )
    }

    internalIndex.value = null;
    setTimeout(() => {
        const self = document.getElementById("nutricional");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    nextStep()
};

const marcasBombaOpts = ref([
    'Kangaroo', 
    'Halyard', 
    'Kimberly-Clark', 
    'Wilson Cook', 
    'Folley'
]);

const marcasFiltradas = ref([]);

const searchMarcaBomba = (event) => {
    if (!event.query.trim().length) {
        marcasFiltradas.value = [...marcasBombaOpts.value];
    } else {
        marcasFiltradas.value = marcasBombaOpts.value.filter((marca) => {
            return marca.toLowerCase().startsWith(event.query.toLowerCase());
        });
    }
};

watch(() => nutricionalStore.nutricional.alimentacao_enteral, (novoValor) => {
    if (novoValor === true) {
        nutricionalStore.nutricional.alimentacao_parenteral = false;
    }
});

watch(() => nutricionalStore.nutricional.alimentacao_parenteral, (novoValor) => {
    if (novoValor === true) {
        nutricionalStore.nutricional.alimentacao_enteral = false;
    }
});



</script>

<template>
    <Accordion v-model:activeIndex="internalIndex" id="nutricional" class="scroll-mt-24 mt-8 mb-8 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Nutricional <i class="pi pi-apple text-3xl"></i></h4>
                <span class="text-xs text-gray-500 font-normal -mt-4">
                    {{ isFilled ? '' : 'Toque para preencher' }}
                </span>
            </div>
        </div>
    </template>
        
        <div class="flex flex-col gap-4 w-full">

            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="alimentacao_oral">Alimentação Oral?</label>
                    <InputSwitch 
                        id="alimentacao_oral" 
                        v-model="nutricionalStore.nutricional.alimentacao_oral" 
                    />
                </div>
                
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="alimentacao_enteral">Alimentação Enteral?</label>
                    <InputSwitch 
                        id="alimentacao_enteral" 
                        v-model="nutricionalStore.nutricional.alimentacao_enteral" 
                    />
                </div>

                <div v-if="nutricionalStore.nutricional.alimentacao_enteral" class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="via_enteral">Via de alimentação enteral <span class="text-red-500">*</span></label>
                    <Select 
                        id="via_enteral" 
                        v-model="nutricionalStore.nutricional.via_enteral" 
                        :options="props.formFields.via_enteral" 
                        optionLabel="label" 
                        optionValue="label"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/5">
                    <label for="alimentacao_parenteral">Alimentação Parenteral?</label>
                    <InputSwitch    
                        id="alimentacao_parenteral" 
                        v-model="nutricionalStore.nutricional.alimentacao_parenteral" 
                    />
                </div>

                <div v-if="nutricionalStore.nutricional.alimentacao_parenteral" class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="via_parenteral">Via de alimentação parenteral <span class="text-red-500">*</span></label>
                    <Select 
                        id="via_parenteral" 
                        v-model="nutricionalStore.nutricional.via_parenteral" 
                        :options="props.formFields.via_parenteral" 
                        optionLabel="label" 
                        optionValue="label"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
            </div>

            <div class="flex flex-col md:flex-row gap-4 border-t pt-4">
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="sonda">Usa Sonda?</label>
                    <InputSwitch 
                        id="sonda" 
                        v-model="nutricionalStore.nutricional.sonda" 
                    />
                </div>

                <template v-if="nutricionalStore.nutricional.sonda">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="data_ultima_troca">Data da última troca</label>
                        <InputMask
                            id="data_ultima_troca" 
                            v-model="nutricionalStore.nutricional.data_ultima_troca" 
                            mask="99/99/9999"
                            placeholder="dd/mm/aaaa" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="adaptador_sonda">Adaptador da sonda</label>
                        <Select 
                            id="adaptador_sonda" 
                            v-model="nutricionalStore.nutricional.adaptador_sonda" 
                            :options="props.formFields.adaptador_sonda" 
                            optionLabel="label" 
                            optionValue="label"
                            placeholder="Selecione" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="botton">Usa Botton?</label>
                        <InputSwitch 
                            id="botton" 
                            v-model="nutricionalStore.nutricional.botton" 
                        />
                    </div>
                </template>
            </div>

            <div class="flex flex-col gap-2 w-full border-t pt-4">
                <label for="restricao_alimentar">Restrição alimentar <span class="text-red-500">*</span></label>
                <InputText 
                    id="restricao_alimentar" 
                    v-model="nutricionalStore.nutricional.restricao_alimentar" 
                    placeholder="Descrição da restrição alimentar" 
                    class="w-full"
                />
            </div>
            <div class="flex flex-col gap-4 border-t pt-4">
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-2/10">
                        <label for="tipo_dieta">Tipo de dieta <span class="text-red-500">*</span></label>
                        
                        <Select 
                            id="tipo_dieta" 
                            v-model="nutricionalStore.nutricional.tipo_dieta" 
                            :options="props.formFields.tipo_dieta" 
                            optionLabel="label" 
                            optionValue="label"
                            placeholder="Selecione" 
                            class="w-full"
                        />
                    </div>

                    <template v-if="nutricionalStore.nutricional.tipo_dieta === 'Industrializada' || nutricionalStore.nutricional.tipo_dieta === 'Mista'">
                        <div class="flex flex-col gap-2 w-full md:w-2/10">
                            <label for="gavando">Gavando (ml/h) <span class="text-red-500">*</span></label>
                            <InputText 
                                id="gavando" 
                                v-model="nutricionalStore.nutricional.gavando" 
                                type="number"
                                placeholder="ml/h" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-2/10">
                            <label for="volume_diario">Volume diário <span class="text-red-500">*</span></label>
                            <InputText 
                                id="volume_diario" 
                                v-model="nutricionalStore.nutricional.volume_diario" 
                                type="number"
                                placeholder="Total em ml" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-3/10">
                            <label for="forma_administracao">Forma de administração <span class="text-red-500">*</span></label>
                            <Select 
                                id="forma_administracao" 
                                v-model="nutricionalStore.nutricional.forma_administracao" 
                                :options="props.formFields.forma_administracao" 
                                optionLabel="label" 
                                optionValue="label"
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                        <div v-if="nutricionalStore.nutricional.forma_administracao === 'Bomba de Infusão'" class="flex flex-col gap-2 w-full md:w-1/3 fadein animation-duration-300">
                            <label for="marca_bomba">Marca da Bomba <span class="text-red-500">*</span></label>  
                            <AutoComplete 
                                id="marca_bomba" 
                                v-model="nutricionalStore.nutricional.marca_bomba" 
                                :suggestions="marcasFiltradas" 
                                @complete="searchMarcaBomba" 
                                placeholder="Ex: Kangaroo, Halyard..." 
                                class="w-full"
                                :dropdown="true" 
                            />
                        </div>
                    </template>
                </div>
            </div>

            <div class="flex flex-col gap-4 border-t pt-4">
                <div class="flex flex-col md:flex-row gap-4 items-center">
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="suplemento">Suplemento?</label>
                        <InputSwitch 
                            id="suplemento" 
                            v-model="nutricionalStore.nutricional.suplemento" 
                        />
                    </div>
                    
                    <div v-if="nutricionalStore.nutricional.suplemento" class="flex flex-col gap-2 w-full flex-1">
                        <label for="qual_suplemento">Qual suplemento? <span class="text-red-500">*</span></label>
                        <InputText 
                            id="qual_suplemento" 
                            v-model="nutricionalStore.nutricional.qual_suplemento" 
                            placeholder="Nome do suplemento" 
                            class="w-full"
                        />
                    </div>
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