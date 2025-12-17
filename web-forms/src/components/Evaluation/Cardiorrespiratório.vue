    <script setup>
    import { ref, computed } from 'vue';
    import { useRespiratorioStore } from '@/store/evaluation/cardiorespiratorio';
    import { useStepAccordion } from "@/composable/useStepAccordion";
    import { useAvaliationForm } from "@/store/evaluation/form";
    import { AvaliationService } from '@/service/AvaliationService';
import { ToggleSwitch } from 'primevue';
    
    const avaliationFormStore = useAvaliationForm();
    const respiratorioStore = useRespiratorioStore();

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

    const isFilled = computed(() => {
        return !!respiratorioStore.respiratorio.padrao_respiratorio && 
            respiratorioStore.respiratorio.padrao_respiratorio.length > 3;
    });

    const { internalIndex, nextStep } = useStepAccordion(props, emit);

    const handleSave = async () => {
        const payload = {
            cardioRespiratorio: respiratorioStore.respiratorio,
        };
        try {
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
                const self = document.getElementById("cardiorrespiratorio");
                if (self) {
                    self.scrollIntoView({ 
                        behavior: 'instant', 
                        block: 'start',
                    });
                }
            }, 0);
            nextStep('next-step');
        } catch (error) {
            console.error(error);
        }
    };


    </script>

    <template>
        <Fluid>
            <Accordion v-model:activeIndex="internalIndex" id="cardiorrespiratorio" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
                <AccordionTab>
                <template #header>
                    <div class="flex items-center gap-3 w-full">
                        
                        <i class="pi text-xl" 
                            :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
                        </i>
                        <div class="flex flex-col w-full text-left">
                            <h4 class="font-semibold text-xl p-0 m-0">Cardiorrespiratório <i class="pi pi-cloud text-2xl hidden sm:inline"></i></h4>

                            <span class="text-xs text-gray-500 font-normal -mt-4">
                                {{ isFilled ? '' : 'Toque para preencher' }}
                            </span>
                        </div>
                        
                    </div>
                    <div class="w-full text-right"> 
                <button 
                    type="button" 
                    class="p-2 mr-6 rounded-lg bg-red-600 hover:bg-red-700 transition-colors" 
                    @click="respiratorioStore.resetRespiratorio"> 
                    <i class="pi pi-trash text-white"></i>
                </button>   
            </div>
                </template>
    

                <div class="flex flex-col gap-4 w-full">
                    <div class="flex flex-col md:flex-row gap-4">
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label for="padrao_respiratorio">Padrão respiratório <span class="text-red-500">*</span></label>
                            <Select 
                                id="padrao_respiratorio" 
                                v-model="respiratorioStore.respiratorio.padrao_respiratorio" 
                                :options="props.formFields.padrao_respiratorio" 
                                optionLabel="label"
                                optionValue="label" 
                                placeholder="Selecione" 
                                class="w-full"

                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label for="via_aerea">Via aérea <span class="text-red-500">*</span></label>
                            <Select 
                                id="via_aerea" 
                                v-model="respiratorioStore.respiratorio.via_aerea" 
                                :options="props.formFields.via_aerea" 
                                optionLabel="label"
                                optionValue="label" 
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                    </div>

                    <div v-show="respiratorioStore.respiratorio.via_aerea === 'Orotraqueal'" class="flex flex-col md:flex-row gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
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
                            <InputMask
                                id="data_orotraqueal" 
                                v-model="respiratorioStore.respiratorio.data_orotraqueal" 
                                mask="99/99/9999"
                                placeholder="dd/mm/aaaa" 
                                class="w-full"
                                showIcon
                            />
                        </div>
                    </div>

                    <div class="border-t my-2"></div>

                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="suporte_ventilatorio">Suporte ventilatório <span class="text-red-500">*</span></label>
                        <Select 
                            id="suporte_ventilatorio" 
                            v-model="respiratorioStore.respiratorio.suporte_ventilatorio" 
                            :options="props.formFields.suporte_ventilatorio" 
                            optionLabel="label"
                            optionValue="label" 
                            placeholder="Selecione o tipo de suporte" 
                            class="w-full"
                        />
                    </div>
                    <div v-show="respiratorioStore.respiratorio.suporte_ventilatorio === 'Ventilação Não Invasiva (VNI)'" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 mt-2 dark:bg-gray-400/10">
                    <h5 class="font-medium text-gray-700">Detalhes da VNI</h5>
                        
                        <div class="flex flex-col md:flex-row gap-4">
                            <div class="flex flex-col gap-2 w-full md:w-1/3">
                                <label for="modo_vni">Modo</label>
                                <Select 
                                    id="modo_vni" 
                                    v-model="respiratorioStore.respiratorio.modo_vni" 
                                    :options="props.formFields.modo_vni" 
                                    optionLabel="label"
                                    optionValue="label" 
                                    placeholder="Selecione" 
                                    class="w-full"
                                />
                            </div>
                            <div class="flex flex-col gap-2 w-full md:w-1/3">
                                <label for="frequencia_vni">Frequência</label>
                                <Select 
                                    id="frequencia_vni" 
                                    v-model="respiratorioStore.respiratorio.frequencia_vni" 
                                    :options="props.formFields.frequencia_vni" 
                                    optionLabel="label"
                                    optionValue="label" 
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
                                    :options="props.formFields.mascara_vni" 
                                    optionLabel="label"
                                    optionValue="label" 
                                    placeholder="Selecione" 
                                    class="w-full"
                                />
                            </div>
                            <div class="flex flex-col gap-2 w-full md:w-1/3">
                                <label for="tamanho_mascara_vni">Tamanho da Máscara</label>
                                <Select 
                                    id="tamanho_mascara_vni" 
                                    v-model="respiratorioStore.respiratorio.tamanho_mascara_vni" 
                                    :options="props.formFields.tamanho_mascara" 
                                    optionLabel="label"
                                    optionValue="label" 
                                    placeholder="Selecione" 
                                    class="w-full"
                                />
                            </div>
                        </div>
                    </div>
                    <div v-show="respiratorioStore.respiratorio.suporte_ventilatorio === 'Assistência Ventilatória Mecânica'" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 mt-2 dark:bg-gray-400/10">
                        <h5 class="font-medium text-gray-700">Detalhes da AVM</h5>
                        
                        <div class="flex flex-col md:flex-row gap-4">
                            <div class="flex flex-col gap-2 w-full md:w-1/3">
                                <label for="modo_avm">Modo</label>
                                <Select 
                                    id="modo_avm" 
                                    v-model="respiratorioStore.respiratorio.modo_avm" 
                                    :options="props.formFields.modo_avm" 
                                    optionLabel="label"
                                    optionValue="label" 
                                    placeholder="Selecione" 
                                    class="w-full"
                                />
                            </div>
                            <div class="flex flex-col gap-2 w-full md:w-1/3">
                                <label for="frequencia_avm">Frequência</label>
                                <Select 
                                    id="frequencia_avm" 
                                    v-model="respiratorioStore.respiratorio.frequencia_avm" 
                                    :options="frequenciaVniOpts" 
                                    optionLabel="name"
                                    optionValue="name" 
                                    placeholder="Selecione" 
                                    class="w-full"
                                />
                            </div>
                            <div class="flex flex-col gap-2 w-full md:w-1/3">
                                <label for="fio2_avm">FIO2 (%)</label>
                                <InputGroup>
                                    <InputText 
                                        id="fio2_avm" 
                                        v-model="respiratorioStore.respiratorio.fio2_avm" 
                                        placeholder="Ex: 40" 
                                        type="number"
                                    />
                                    <InputGroupAddon>%</InputGroupAddon>
                                </InputGroup>
                            </div>
                        </div>

                        <div class="flex flex-col md:flex-row gap-4">
                            <div class="flex flex-col gap-2 w-full md:w-1/3">
                                <label for="mascara_avm">Tipo de Máscara</label>
                                <Select 
                                    id="mascara_avm" 
                                    v-model="respiratorioStore.respiratorio.mascara_avm" 
                                    :options="mascaraVniOpts" 
                                    optionLabel="name"
                                    optionValue="name" 
                                    placeholder="Selecione" 
                                    class="w-full"
                                />
                            </div>
                            <div class="flex flex-col gap-2 w-full md:w-1/3">
                                <label for="tamanho_mascara_avm">Tamanho da Máscara</label>
                                <Select 
                                    id="tamanho_mascara_avm" 
                                    v-model="respiratorioStore.respiratorio.tamanho_mascara_avm" 
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
                            <ToggleSwitch 
                                id="usa_oxigenioterapia"
                                v-model="respiratorioStore.respiratorio.usa_oxigenioterapia" 
                            />
                        </div>

                        <div v-show="respiratorioStore.respiratorio.usa_oxigenioterapia" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                            <div class="flex flex-col md:flex-row gap-4">
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="fonte_oxigenio">Fonte</label>
                                    <Select 
                                        id="fonte_oxigenio" 
                                        v-model="respiratorioStore.respiratorio.fonte_oxigenio" 
                                        :options="fonteOxigenioOpts" 
                                        optionLabel="name"
                                        optionValue="name" 
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="modalidade_oxigenio">Modalidade</label>
                                    <Select 
                                        id="modalidade_oxigenio" 
                                        v-model="respiratorioStore.respiratorio.modalidade_oxigenio" 
                                        :options="modalidadeOxigenioOpts" 
                                        optionLabel="name"
                                        optionValue="name" 
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
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
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="frequencia_oxigenio">Frequência</label>
                                    <Select 
                                        id="frequencia_oxigenio" 
                                        v-model="respiratorioStore.respiratorio.frequencia_oxigenio" 
                                        :options="frequenciaOxigenioOpts" 
                                        optionLabel="name"
                                        optionValue="name" 
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="vezes_dia_oxigenio">Vezes ao dia</label>
                                    <InputText 
                                        id="vezes_dia_oxigenio" 
                                        v-model="respiratorioStore.respiratorio.vezes_dia_oxigenio" 
                                        placeholder="Se intermitente" 
                                        type="number" 
                                        class="w-full"
                                        :disabled="respiratorioStore.respiratorio.frequencia_oxigenio !== 'Intermitente'"
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