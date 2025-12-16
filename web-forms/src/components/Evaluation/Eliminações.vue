<script setup>
import { computed, watch } from 'vue';
import { useEliminacoesStore } from '@/store/evaluation/eliminacoes';
import { InputMask } from 'primevue';
import { useStepAccordion } from "@/composable/useStepAccordion";
import { useAvaliationForm } from "@/store/evaluation/form";
import { AvaliationService } from '@/service/AvaliationService';

const avaliationFormStore = useAvaliationForm();
const eliminacoesStore = useEliminacoesStore();

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
    return !!eliminacoesStore.eliminacoes.funcao_intestinal;
});

const { internalIndex, nextStep } = useStepAccordion(props, emit);

const handleSave = async () => {
    const payload = {
        eliminacoes: eliminacoesStore.eliminacoes,
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
        const self = document.getElementById("eliminacoes");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    nextStep()
};

watch(() => eliminacoesStore.eliminacoes.sva, (novoValor) => {
    if (novoValor === true) {
        eliminacoesStore.eliminacoes.svd = false;
    }
});

watch(() => eliminacoesStore.eliminacoes.svd, (novoValor) => {
    if (novoValor === true) {
        eliminacoesStore.eliminacoes.sva = false;
    }
});

</script>

<template>
    <Accordion v-model:activeIndex="internalIndex" id="eliminacoes" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Eliminações <i class="pi pi-list text-3xl "></i></h4>
                <span class="text-xs text-gray-500 font-normal -mt-4">
                    {{ isFilled ? '' : 'Toque para preencher' }}
                </span>
            </div>
        </div>
    </template>
        <div class="flex flex-col gap-4 w-full">
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/5">
                    <label for="funcao_intestinal">Função Intestinal</label>
                    <Select 
                        id="funcao_intestinal" 
                        v-model="eliminacoesStore.eliminacoes.funcao_intestinal" 
                        :options="props.formFields.funcao_intestinal" 
                        optionLabel="label" 
                        optionValue="label"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/8">
                    <label for="via_evacuacao_fisiologica">Via Fisiológica?</label>
                    <InputSwitch 
                        id="via_evacuacao_fisiologica" 
                        v-model="eliminacoesStore.eliminacoes.via_evacuacao_fisiologica" 
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="via_evacuacao_estomia">Usa Estomia?</label>
                    <InputSwitch 
                        id="via_evacuacao_estomia" 
                        v-model="eliminacoesStore.eliminacoes.via_evacuacao_estomia" 
                    />
                </div>
            </div>

            <div v-if="eliminacoesStore.eliminacoes.via_evacuacao_estomia" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 mt-2 dark:bg-gray-400/10">
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="tipo_estomia">Tipo de Estomia</label>
                        <Select 
                            id="tipo_estomia" 
                            v-model="eliminacoesStore.eliminacoes.tipo_estomia" 
                            :options="props.formFields.tipo_estomia"
                            optionLabel="label" 
                            optionValue="label"
                            placeholder="Selecione" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="placa_estomia">Placa de Estomia</label>
                        <InputText 
                            id="placa_estomia" 
                            v-model="eliminacoesStore.eliminacoes.placa_estomia" 
                            placeholder="Modelo/Marca" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="bolsa_estomia">Bolsa de Estomia</label>
                        <InputText 
                            id="bolsa_estomia" 
                            v-model="eliminacoesStore.eliminacoes.bolsa_estomia" 
                            placeholder="Modelo/Marca" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>

            <h5 class="font-medium text-gray-700 border-b pb-2 mt-4">Diurese e Dispositivos Urinários</h5>
            
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/5">
                    <label for="diurese">Diurese</label>
                    <Select 
                        id="diurese" 
                        v-model="eliminacoesStore.eliminacoes.diurese" 
                        :options="props.formFields.volume_diurese" 
                        optionLabel="label" 
                        optionValue="label"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>

                <div v-if="eliminacoesStore.eliminacoes.diurese === 'induzida'" class="flex flex-col gap-2 w-full md:w-1/4">
                    <label for="num_snd_diurese">Nº da sonda (Fr)</label>
                    <InputText 
                        id="num_snd_diurese" 
                        v-model="eliminacoesStore.eliminacoes.num_snd_diurese" 
                        placeholder="Ex: 16" 
                        class="w-full"
                    />
                </div>
                
                
                <div class="flex flex-col gap-2 w-full md:w-1/8">
                    <label for="usa_fralda">Usa Fralda?</label>
                    <InputSwitch 
                        id="usa_fralda" 
                        v-model="eliminacoesStore.eliminacoes.usa_fralda" 
                    />
                </div>
                <div v-if="eliminacoesStore.eliminacoes.usa_fralda" class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="trocas_fralda_dia">Nº trocas por dia</label>
                    <InputText 
                        id="trocas_fralda_dia" 
                        v-model="eliminacoesStore.eliminacoes.trocas_fralda_dia" 
                        placeholder="Ex: 4" 
                        class="w-full"
                    />
                </div>
                
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="sva">SVA (Alívio)?</label>
                    <InputSwitch 
                        id="sva" 
                        v-model="eliminacoesStore.eliminacoes.sva" 
                    />
                </div>
            </div>

            <div class="border-t my-2 pt-2">
                <div class="flex flex-col md:flex-row gap-4 items-center">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="cateterismo_intermitente">Cateterismo Intermitente?</label>
                        <InputSwitch 
                            id="cateterismo_intermitente" 
                            v-model="eliminacoesStore.eliminacoes.cateterismo_intermitente" 
                        />
                    </div>
                    <div v-if="eliminacoesStore.eliminacoes.cateterismo_intermitente" class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="vezes_cateterismo_dia">Nº de vezes ao dia</label>
                        <InputText 
                            id="vezes_cateterismo_dia" 
                            v-model="eliminacoesStore.eliminacoes.vezes_cateterismo_dia" 
                            placeholder="Quantidade" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>

            <div class="border-t my-2 pt-2">
                <div class="flex flex-col gap-2 w-full mb-2">
                    <label for="svd">Sonda Vesical de Demora (SVD)?</label>
                    <InputSwitch id="svd" v-model="eliminacoesStore.eliminacoes.svd" />
                </div>

                <div v-if="eliminacoesStore.eliminacoes.svd" class="grid grid-cols-1 md:grid-cols-4 gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                    <div class="flex flex-col gap-2">
                        <label for="num_svd">Nº da Sonda (Fr)</label>
                        <InputText id="num_svd" v-model="eliminacoesStore.eliminacoes.num_svd" placeholder="Ex: 16" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="vias_svd">Nº de Vias</label>
                        <Select id="vias_svd" v-model="eliminacoesStore.eliminacoes.vias_svd" :options="props.formFields.vias_sonda" optionLabel="label" optionValue="label" placeholder="Selecione" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="volume_diurese_svd">Volume Diurese</label>
                        <Select id="volume_diurese_svd" v-model="eliminacoesStore.eliminacoes.volume_diurese_svd" :options="props.formFields.volume_diurese" optionLabel="label" optionValue="label" placeholder="Selecione" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="data_troca_svd">Última Troca</label>
                        <InputMask id="data_troca_svd" v-model="eliminacoesStore.eliminacoes.data_troca_svd" mask="99/99/9999" placeholder="dd/mm/aaaa" class="w-full" />
                    </div>
                    <div class="flex items-center gap-2 md:col-span-4">
                        <label for="irrigacao_svd">Realiza Irrigação?</label>
                        <InputSwitch id="irrigacao_svd" v-model="eliminacoesStore.eliminacoes.irrigacao_svd" />
                    </div>
                </div>
            </div>

            <div class="border-t my-2 pt-2">
                <div class="flex flex-col gap-2 w-full mb-2">
                    <label for="cistostomia">Cistostomia?</label>
                    <InputSwitch id="cistostomia" v-model="eliminacoesStore.eliminacoes.cistostomia" />
                </div>

                <div v-if="eliminacoesStore.eliminacoes.cistostomia" class="grid grid-cols-1 md:grid-cols-4 gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                    <div class="flex flex-col gap-2">
                        <label for="num_cistostomia">Nº da Sonda</label>
                        <InputText id="num_cistostomia" v-model="eliminacoesStore.eliminacoes.num_cistostomia" placeholder="Calibre" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="vias_cistostomia">Nº de Vias</label>
                        <Select id="vias_cistostomia" v-model="eliminacoesStore.eliminacoes.vias_cistostomia" :options="props.formFields.vias_sonda" optionLabel="label" optionValue="label" placeholder="Selecione" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="volume_diurese_cistostomia">Volume Diurese</label>
                        <Select id="volume_diurese_cistostomia" v-model="eliminacoesStore.eliminacoes.volume_diurese_cistostomia" :options="props.formFields.volume_diurese" optionLabel="label" optionValue="label" placeholder="Selecione" class="w-full" />
                    </div>
                    
                    <div class="md:col-span-4 grid grid-cols-1 md:grid-cols-3 gap-4 border-t pt-2 mt-2">
                        <div class="flex flex-col gap-2">
                            <label for="irrigacao_cistostomia">Realiza Irrigação?</label>
                            <InputSwitch id="irrigacao_cistostomia" v-model="eliminacoesStore.eliminacoes.irrigacao_cistostomia" />
                        </div>
                        <div v-if="eliminacoesStore.eliminacoes.irrigacao_cistostomia" class="flex flex-col gap-2">
                            <label for="qtd_irrigacao_cistostomia">Quantidade Irrigação</label>
                            <InputText id="qtd_irrigacao_cistostomia" v-model="eliminacoesStore.eliminacoes.qtd_irrigacao_cistostomia" placeholder="ml" class="w-full" />
                        </div>
                    </div>

                     <div class="md:col-span-4 grid grid-cols-1 md:grid-cols-3 gap-4 border-t pt-2">
                        <div class="flex flex-col gap-2">
                            <label for="medicacao_irrigacao_cistostomia">Medicação na irrigação?</label>
                            <InputSwitch id="medicacao_irrigacao_cistostomia" v-model="eliminacoesStore.eliminacoes.medicacao_irrigacao_cistostomia" />
                        </div>
                        <div v-if="eliminacoesStore.eliminacoes.medicacao_irrigacao_cistostomia" class="flex flex-col gap-2 md:col-span-2">
                            <label for="nome_medicacao_cistostomia">Qual medicação?</label>
                            <InputText id="nome_medicacao_cistostomia" v-model="eliminacoesStore.eliminacoes.nome_medicacao_cistostomia" placeholder="Nome do medicamento" class="w-full" />
                        </div>
                    </div>
                </div>
            </div>

            <div class="border-t my-2 pt-2">
                <div class="flex flex-col md:flex-row gap-4 items-center">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="preservativo">Preservativo Urinário?</label>
                        <InputSwitch id="preservativo" v-model="eliminacoesStore.eliminacoes.preservativo" />
                    </div>
                    <div v-if="eliminacoesStore.eliminacoes.preservativo" class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="volume_diurese_preservativo">Volume Diurese</label>
                        <Select 
                            id="volume_diurese_preservativo" 
                            v-model="eliminacoesStore.eliminacoes.volume_diurese_preservativo" 
                            :options="props.formFields.volume_diurese" 
                            optionLabel="label" 
                            optionValue="label"
                            placeholder="Selecione" 
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