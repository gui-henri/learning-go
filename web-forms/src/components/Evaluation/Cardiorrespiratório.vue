<script setup>
import { ref, computed } from 'vue';
import { useRespiratorioStore } from '@/store/evaluation/cardiorespiratorio';

const respiratorioStore = useRespiratorioStore();

const padraoRespiratorioOpts = ref([
    { name: 'Eupnéico', code: 'eupneico' },
    { name: 'Dispnéico', code: 'dispneico' },
    { name: 'Bradipnéico', code: 'bradipneico' },
    { name: 'Taquipnéico', code: 'taquipneico' }
]);

const viaAereaOpts = ref([
    { name: 'Fisiológica', code: 'fisiologica' },
    { name: 'Orotraqueal', code: 'orotraqueal' },
    { name: 'Traqueóstomo', code: 'traqueostomo' }
]);

const suporteVentilatorioOpts = ref([
    { name: 'Espontânea', code: 'espontanea' },
    { name: 'Ventilação Não Invasiva (VNI)', code: 'vni' },
    { name: 'Assistência Ventilatória Mecânica', code: 'mecanica' }
]);

const modoVniOpts = ref([
    { name: 'CPAP Simples', code: 'cpap_simples' },
    { name: 'CPAP Automático', code: 'cpap_automatico' },
    { name: 'BiPAP', code: 'bipap' }
]);

const frequenciaVniOpts = ref([
    { name: 'Contínuo', code: 'continuo' },
    { name: 'Intermitente', code: 'intermitente' }
]);

const mascaraVniOpts = ref([
    { name: 'Nasal', code: 'nasal' },
    { name: 'Facial', code: 'facial' },
    { name: 'Oronasal', code: 'oronasal' }
]);

const tamanhoMascaraOpts = ref([
    { name: 'P', code: 'P' },
    { name: 'M', code: 'M' },
    { name: 'G', code: 'G' }
]);

const fonteOxigenioOpts = ref([
    { name: 'Cilindro', code: 'cilindro' },
    { name: 'Concentrador', code: 'concentrador' }
]);

const frequenciaOxigenioOpts = ref([
    { name: 'Intermitente', code: 'intermitente' },
    { name: 'Contínua', code: 'continua' }
]);

const modalidadeOxigenioOpts = ref([
    { name: 'Cateter Nasal', code: 'cateter_nasal' },
    { name: 'Venturi', code: 'venturi' }
]);

const emit = defineEmits(['next-step']);
const activeIndex = ref(null);

const isFilled = computed(() => {
    return !!respiratorioStore.respiratorio.padrao_respiratorio && 
           respiratorioStore.respiratorio.padrao_respiratorio.length > 3;
});

const handleSave = () => {
    activeIndex.value = null;
    setTimeout(() => {
        const self = document.getElementById("cardiorrespiratorio");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    emit('next-step');
};
</script>

<template>
    <Fluid>
        <Accordion v-model:activeIndex="activeIndex" id="cardiorrespiratorio" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <AccordionTab>
            <template #header>
                <div class="flex items-center gap-3 w-full">
                    <i class="pi text-xl" 
                        :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
                    </i>
                    <div class="flex flex-col text-left">
                        <h4 class="font-semibold text-xl p-0 m-0">Cardiorrespiratório</h4>
                        <span class="text-xs text-gray-500 font-normal -mt-4">
                            {{ isFilled ? '' : 'Toque para preencher' }}
                        </span>
                    </div>
                </div>
            </template>
            <div class="flex flex-col gap-4 w-full">
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="padrao_respiratorio">Padrão respiratório <span class="text-red-500">*</span></label>
                        <Select 
                            id="padrao_respiratorio" 
                            v-model="respiratorioStore.respiratorio.padrao_respiratorio" 
                            :options="padraoRespiratorioOpts" 
                            optionLabel="name"
                            optionValue="code" 
                            placeholder="Selecione" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="via_aerea">Via aérea <span class="text-red-500">*</span></label>
                        <Select 
                            id="via_aerea" 
                            v-model="respiratorioStore.respiratorio.via_aerea" 
                            :options="viaAereaOpts" 
                            optionLabel="name"
                            optionValue="code" 
                            placeholder="Selecione" 
                            class="w-full"
                        />
                    </div>
                </div>

                <div v-if="respiratorioStore.respiratorio.via_aerea === 'orotraqueal'" class="flex flex-col md:flex-row gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="numero_tubo">Número do tubo</label>
                        <InputText 
                            id="numero_tubo" 
                            v-model="respiratorioStore.respiratorio.numero_tubo" 
                            placeholder="Ex: 7.5" 
                            type="number"
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="data_orotraqueal">Data orotraqueal</label>
                        <Calendar 
                            id="data_orotraqueal" 
                            v-model="respiratorioStore.respiratorio.data_orotraqueal" 
                            dateFormat="dd/mm/yy" 
                            placeholder="dd/mm/aaaa" 
                            class="w-full"
                        />
                    </div>
                </div>

                <div class="border-t my-2"></div>

                <div class="flex flex-col gap-2 w-full md:w-1/2">
                    <label for="suporte_ventilatorio">Suporte ventilatório <span class="text-red-500">*</span></label>
                    <Select 
                        id="suporte_ventilatorio" 
                        v-model="respiratorioStore.respiratorio.suporte_ventilatorio" 
                        :options="suporteVentilatorioOpts" 
                        optionLabel="name"
                        optionValue="code" 
                        placeholder="Selecione o tipo de suporte" 
                        class="w-full"
                    />
                </div>

                <div v-if="respiratorioStore.respiratorio.suporte_ventilatorio === 'vni'" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 mt-2 dark:bg-gray-400/10">
                    <h5 class="font-medium text-gray-700">Detalhes da VNI</h5>
                    
                    <div class="flex flex-col md:flex-row gap-4">
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="modo_vni">Modo</label>
                            <Select 
                                id="modo_vni" 
                                v-model="respiratorioStore.respiratorio.modo_vni" 
                                :options="modoVniOpts" 
                                optionLabel="name"
                                optionValue="code" 
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="frequencia_vni">Frequência</label>
                            <Select 
                                id="frequencia_vni" 
                                v-model="respiratorioStore.respiratorio.frequencia_vni" 
                                :options="frequenciaVniOpts" 
                                optionLabel="name"
                                optionValue="code" 
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="fio2_vni">FIO2 (%)</label>
                            <InputGroup>
                                <InputText 
                                    id="fio2_vni" 
                                    v-model="respiratorioStore.respiratorio.fio2_vni" 
                                    placeholder="Ex: 40" 
                                    type="number"
                                />
                                <InputGroupAddon>%</InputGroupAddon>
                            </InputGroup>
                        </div>
                    </div>

                    <div class="flex flex-col md:flex-row gap-4">
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="mascara_vni">Tipo de Máscara</label>
                            <Select 
                                id="mascara_vni" 
                                v-model="respiratorioStore.respiratorio.mascara_vni" 
                                :options="mascaraVniOpts" 
                                optionLabel="name"
                                optionValue="code" 
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="tamanho_mascara_vni">Tamanho da Máscara</label>
                            <Select 
                                id="tamanho_mascara_vni" 
                                v-model="respiratorioStore.respiratorio.tamanho_mascara_vni" 
                                :options="tamanhoMascaraOpts" 
                                optionLabel="name"
                                optionValue="code" 
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                    </div>
                </div>

                <div class="border-t my-2"></div>

                <div class="flex flex-col gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="usa_oxigenioterapia">Oxigenioterapia?</label>
                        <InputSwitch 
                            id="usa_oxigenioterapia" 
                            v-model="respiratorioStore.respiratorio.usa_oxigenioterapia" 
                        />
                    </div>

                    <div v-if="respiratorioStore.respiratorio.usa_oxigenioterapia" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                        <div class="flex flex-col md:flex-row gap-4">
                            <div class="flex flex-col gap-2 w-full md:w-1/4">
                                <label for="fonte_oxigenio">Fonte</label>
                                <Select 
                                    id="fonte_oxigenio" 
                                    v-model="respiratorioStore.respiratorio.fonte_oxigenio" 
                                    :options="fonteOxigenioOpts" 
                                    optionLabel="name"
                                    optionValue="code" 
                                    placeholder="Selecione" 
                                    class="w-full"
                                />
                            </div>
                            <div class="flex flex-col gap-2 w-full md:w-1/4">
                                <label for="modalidade_oxigenio">Modalidade</label>
                                <Select 
                                    id="modalidade_oxigenio" 
                                    v-model="respiratorioStore.respiratorio.modalidade_oxigenio" 
                                    :options="modalidadeOxigenioOpts" 
                                    optionLabel="name"
                                    optionValue="code" 
                                    placeholder="Selecione" 
                                    class="w-full"
                                />
                            </div>
                            <div class="flex flex-col gap-2 w-full md:w-1/4">
                                <label for="fluxo_oxigenio">Fluxo de O2</label>
                                <InputGroup>
                                    <InputText 
                                        id="fluxo_oxigenio" 
                                        v-model="respiratorioStore.respiratorio.fluxo_oxigenio" 
                                        placeholder="Ex: 2" 
                                        type="number"
                                    />
                                    <InputGroupAddon>L/min</InputGroupAddon>
                                </InputGroup>
                            </div>
                        </div>

                        <div class="flex flex-col md:flex-row gap-4">
                            <div class="flex flex-col gap-2 w-full md:w-1/4">
                                <label for="frequencia_oxigenio">Frequência</label>
                                <Select 
                                    id="frequencia_oxigenio" 
                                    v-model="respiratorioStore.respiratorio.frequencia_oxigenio" 
                                    :options="frequenciaOxigenioOpts" 
                                    optionLabel="name"
                                    optionValue="code" 
                                    placeholder="Selecione" 
                                    class="w-full"
                                />
                            </div>
                            <div class="flex flex-col gap-2 w-full md:w-1/4">
                                <label for="vezes_dia_oxigenio">Vezes ao dia</label>
                                <InputText 
                                    id="vezes_dia_oxigenio" 
                                    v-model="respiratorioStore.respiratorio.vezes_dia_oxigenio" 
                                    placeholder="Se intermitente" 
                                    type="number" 
                                    class="w-full"
                                    :disabled="respiratorioStore.respiratorio.frequencia_oxigenio !== 'intermitente'"
                                />
                            </div>
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
    </Fluid>
</template>