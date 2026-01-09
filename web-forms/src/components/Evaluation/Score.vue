            <script setup>
            import { ref, computed } from 'vue';
            import { useScoreStore } from '@/store/evaluation/score';
            import { useStepAccordion } from "@/composable/useStepAccordion";
            import { useAvaliationForm } from "@/store/evaluation/form";
            import { AvaliationService } from '@/service/AvaliationService';
            import { ToggleSwitch } from 'primevue';

            const avaliationFormStore = useAvaliationForm();
            const scoreStore = useScoreStore();
            
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
            });
            
            const isFilled = computed(() => {
            const s = scoreStore.score
            return typeof s.diagnostico_primario === 'string' && s.diagnostico_primario.trim().length > 3
            })

            const { internalIndex, nextStep } = useStepAccordion(props, emit);

            const handleSave = async () => {
                const payload = {
                    score: scoreStore.score,
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
                        const self = document.getElementById("score");
                        if (self) {
                            self.scrollIntoView({ 
                                behavior: 'instant', 
                                block: 'start',
                            });
                        }
                    }, 0);
                    nextStep()
                } catch (error) {
                    console.error(error);
                }
            };

            // 1. Função que busca o valor de 'points' dentro do JSON que veio do back
const getPoints = (fieldKey, selectedValue) => {
    if (!props.formFields[fieldKey] || !selectedValue) return 0;
    const option = props.formFields[fieldKey].find(opt => opt.value === selectedValue);
    return option ? (option.points || 0) : 0;
};

// 2. Cálculo do Katz (Soma 1 ponto se for 'independente')
const totalKatzRaw = computed(() => {
    const s = scoreStore.score;
    const fields = ['banho', 'vestir', 'higiene_pessoal', 'transferencia', 'continencia', 'alimentacao'];
    return fields.reduce((acc, field) => acc + (s[field] === 'independente' ? 1 : 0), 0);
});

// 3. Recomendação do Grupo 2 (Baseada nos IDs/Values)
const recomendacaoGrupo2 = computed(() => {
    const s = scoreStore.score;
    const isHigh = (v) => v === 'mais_12h' || v === 'mais_5x' || v === 'mais_4x';
    const isMed = (v) => v === 'ate_12h' || v === 'ate_5x' || v === 'ate_4x';

    if (isHigh(s.alimentacao_parenteral) || isHigh(s.aspiracao_traqueo) || isHigh(s.ventilacao_mecanica) || s.medicacao_parenteral === 'mais_4x') {
        return 'Internação Domiciliar 24h';
    }
    if (isMed(s.alimentacao_parenteral) || isMed(s.aspiracao_traqueo) || isMed(s.ventilacao_mecanica)) {
        return 'Internação Domiciliar 12h';
    }
    return 'Atendimento Domiciliar / Outros Programas';
});

// 4. Cálculo do NEAD (Soma pontos do Grupo 3 + Bônus do Katz)
const totalNead = computed(() => {
    const s = scoreStore.score;
    let total = 0;
    
    // Soma pontos do Grupo 3 usando os valores do formFields
    total += getPoints('estado_nutricional', s.estado_nutricional);
    total += getPoints('via_alimentacao', s.via_alimentacao_medicacao);
    total += getPoints('internacoes', s.internacoes_ultimo_ano);
    total += getPoints('aspiracao_vias_aereas', s.aspiracao_vias_aereas);
    total += getPoints('lesoes', s.lesoes);
    total += getPoints('via_medicacoes', s.via_medicacoes);
    total += getPoints('frequencia_respiratoria', s.uso_oxigenoterapia);
    total += getPoints('nivel_consciencia', s.nivel_consciencia);
    total += getPoints('exercicio_ventilatorio', s.exercicio_ventilatorio);

    // Bônus Katz (Lógica: Menos de 2 no Katz = +2 no NEAD | 3 ou 4 no Katz = +1 no NEAD)
    const ptsK = totalKatzRaw.value;
    if (ptsK < 2) total += 2;
    else if (ptsK === 3 || ptsK === 4) total += 1;

    return total;
});

// 5. Classes de cores dinâmicas
const neadColorClass = computed(() => {
    return totalNead.value >= 12 
        ? 'bg-red-50 border-red-200 text-red-900 dark:bg-red-900/20 dark:text-red-300' 
        : 'bg-green-50 border-green-200 text-green-900 dark:bg-green-900/20 dark:text-green-300';
});

            </script>

            <template>
                <Accordion v-model:activeIndex="internalIndex" id="score" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600 dark:bg-surface-900">
                    <AccordionTab>
                        <template #header>
                            <div class="flex items-center gap-3 w-full">
                                <i class="pi text-xl" 
                                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
                                </i>
                                <div class="flex flex-col text-left">
                                    <h4 class="font-semibold text-xl p-0 m-0 dark:text-gray-100">Score NEAD e Katz <i class="pi pi-folder-plus text-3xl"></i></h4>
                                    <span class="text-xs text-gray-500 font-normal -mt-4 dark:text-gray-400">
                                        {{ isFilled ? '' : 'Toque para preencher' }}
                                    </span>
                                </div>
                            </div>
                        </template>
                        
                        <div class="flex flex-col gap-6 w-full">

                            <!-- GRUPO 1: DIAGNÓSTICO -->
                            <h5 class="font-medium text-gray-700 border-b pb-2 dark:text-gray-300 dark:border-gray-700">Diagnóstico e Elegibilidade (Grupo 1)</h5>
                            
                            <div class="flex flex-col md:flex-row gap-4">
                                <div class="flex flex-col gap-2 w-full md:w-1/2">
                                    <label for="diagnostico_primario" class="dark:text-gray-200">Diagnóstico Primário</label>
                                    <InputText 
                                        id="diagnostico_primario" 
                                        v-model="scoreStore.score.diagnostico_primario" 
                                        placeholder="CID ou Descrição" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/2">
                                    <label for="diagnostico_secundario" class="dark:text-gray-200">Diagnóstico Secundário</label>
                                    <InputText 
                                        id="diagnostico_secundario" 
                                        v-model="scoreStore.score.diagnostico_secundario" 
                                        placeholder="CID ou Descrição" 
                                        class="w-full"
                                    />
                                </div>
                            </div>
                            
                            <div class="flex flex-col md:flex-row gap-4 border-t pt-4 dark:border-gray-700">
                                <div class="flex flex-col gap-2 w-full md:w-1/2">
                                    <label for="domicilio_risco" class="dark:text-gray-200">O domicílio é de risco?</label>
                                    <ToggleSwitch 
                                        id="domicilio_risco" 
                                        v-model="scoreStore.score.domicilio_risco" 
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/2">
                                    <label for="impedimento_deslocamento" class="dark:text-gray-200">Impedimento de deslocamento?</label>
                                    <ToggleSwitch 
                                        id="impedimento_deslocamento" 
                                        v-model="scoreStore.score.impedimento_deslocamento" 
                                    />
                                </div>
                            </div>

                            <!-- GRUPO 2: INDICAÇÃO IMEDIATA -->
                            <h5 class="font-medium text-gray-700 border-b pb-2 mt-2 dark:text-gray-300 dark:border-gray-700">Indicação Imediata (Grupo 2)</h5>
                            
                            <div class="flex flex-col md:flex-row gap-4">
                                <div class="flex flex-col gap-2 w-full md:w-1/2">
                                    <label for="alimentacao_parenteral" class="dark:text-gray-200">Alimentação Parenteral</label>
                                    <Select 
                                        id="alimentacao_parenteral" 
                                        v-model="scoreStore.score.alimentacao_parenteral" 
                                        :options="props.formFields.freq_12h" 
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/2">
                                    <label for="aspiracao_traqueo" class="dark:text-gray-200">Aspiração Traqueostomia</label>
                                    <Select 
                                        id="aspiracao_traqueo" 
                                        v-model="scoreStore.score.aspiracao_traqueo" 
                                        :options="props.formFields.freq_5x" 
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                            </div>

                            <div class="flex flex-col md:flex-row gap-4 border-t pt-4 dark:border-gray-700">
                                <div class="flex flex-col gap-2 w-full md:w-1/2">
                                    <label for="ventilacao_mecanica" class="dark:text-gray-200">Ventilação Mecânica</label>
                                    <Select 
                                        id="ventilacao_mecanica" 
                                        v-model="scoreStore.score.ventilacao_mecanica" 
                                        :options="props.formFields.freq_12h"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/2">
                                    <label for="medicacao_parenteral" class="dark:text-gray-200">Medicação Parenteral / Hipodermóclise</label>
                                    <Select 
                                        id="medicacao_parenteral" 
                                        v-model="scoreStore.score.medicacao_parenteral" 
                                        :options="props.formFields.freq_4x"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                            </div>
                            <!-- KATZ -->
                            <div class="flex justify-between items-center border-b pb-2 mt-4 dark:border-gray-700">
                                <h5 class="font-medium text-gray-700 m-0 dark:text-gray-300">Escore de Katz</h5>
                            </div>
                            <div class="flex flex-col gap-2 w-full justify-center   md:w-1/3">
                                    <label for="atividades" class="dark:text-gray-200">Atividades</label>
                                    <Select 
                                        id="atividades" 
                                        v-model="scoreStore.score.atividades" 
                                        :options="props.formFields.katz" 
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>

                            <div class="flex flex-col md:flex-row gap-4">
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="banho" class="dark:text-gray-200">Banhar-se</label>
                                    <Select 
                                        id="banho" 
                                        v-model="scoreStore.score.banho" 
                                        :options="props.formFields.katz" 
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="vestir" class="dark:text-gray-200">Vestir-se</label>
                                    <Select 
                                        id="vestir" 
                                        v-model="scoreStore.score.vestir" 
                                        :options="props.formFields.katz"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="higiene_pessoal" class="dark:text-gray-200">Ir ao banheiro</label>
                                    <Select 
                                        id="higiene_pessoal" 
                                        v-model="scoreStore.score.higiene_pessoal" 
                                        :options="props.formFields.katz"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                            </div>

                            <div class="flex flex-col md:flex-row gap-4 border-t pt-4 dark:border-gray-700">
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="transferencia" class="dark:text-gray-200">Transferência</label>
                                    <Select 
                                        id="transferencia" 
                                        v-model="scoreStore.score.transferencia" 
                                        :options="props.formFields.katz"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="continencia" class="dark:text-gray-200">Continência</label>
                                    <Select 
                                        id="continencia" 
                                        v-model="scoreStore.score.continencia" 
                                        :options="props.formFields.katz"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="alimentacao" class="dark:text-gray-200">Alimentação</label>
                                    <Select 
                                        id="alimentacao" 
                                        v-model="scoreStore.score.alimentacao" 
                                        :options="props.formFields.katz"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                            </div>

                            <!-- CARD RESULTADO KATZ -->
                            <div class="p-4 rounded-lg mt-4 border transition-all" 
     :class="totalKatzRaw >= 5 ? 'bg-green-50' : (totalKatzRaw >= 3 ? 'bg-orange-50' : 'bg-red-50')">
    <h3 class="text-lg font-bold mb-1">Resultado Katz: {{ totalKatzRaw }} / 6</h3>
    <p class="font-medium m-0">
        {{ totalKatzRaw >= 5 ? 'Independente' : (totalKatzRaw >= 3 ? 'Dependência Parcial' : 'Dependente Total') }}
    </p>
</div>

                            <!-- GRUPO 3: CRITÉRIOS DE APOIO -->
                            <div class="flex justify-between items-center border-b pb-2 mt-4 dark:border-gray-700">
                                <h5 class="font-medium text-gray-700 m-0 dark:text-gray-300">Critérios de Apoio (Grupo 3)</h5>
                                </div>

                            <div class="flex flex-col md:flex-row gap-4">
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="estado_nutricional" class="dark:text-gray-200">Estado Nutricional</label>
                                    <Select 
                                        id="estado_nutricional" 
                                        v-model="scoreStore.score.estado_nutricional" 
                                        :options="props.formFields.estado_nutricional"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="via_alimentacao_medicacao" class="dark:text-gray-200">Aliment / Medic Enteral</label>
                                    <Select 
                                        id="via_alimentacao_medicacao" 
                                        v-model="scoreStore.score.via_alimentacao_medicacao" 
                                        :options="props.formFields.via_alimentacao"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="internacoes_ultimo_ano" class="dark:text-gray-200">Internações no último ano</label>
                                    <Select 
                                        id="internacoes_ultimo_ano" 
                                        v-model="scoreStore.score.internacoes_ultimo_ano" 
                                        :options="props.formFields.internacoes"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                            </div>

                            <div class="flex flex-col md:flex-row gap-4 border-t pt-4 dark:border-gray-700">
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="aspiracao_vias_aereas" class="dark:text-gray-200">Aspiração Vias Aéreas Sup.</label>
                                    <Select 
                                        id="aspiracao_vias_aereas" 
                                        v-model="scoreStore.score.aspiracao_vias_aereas" 
                                        :options="props.formFields.aspiracao_vias_aereas"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="lesoes" class="dark:text-gray-200">Lesões</label>
                                    <Select 
                                        id="lesoes"
                                        v-model="scoreStore.score.lesoes"
                                        :options="props.formFields.lesoes"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="via_medicacoes" class="dark:text-gray-200">Medicações (Via)</label>
                                    <Select 
                                        id="via_medicacoes" 
                                        v-model="scoreStore.score.via_medicacoes" 
                                        :options="props.formFields.via_medicacoes"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                            </div>

                            <div class="flex flex-col md:flex-row gap-4 border-t pt-4 dark:border-gray-700">
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="exercicio_ventilatorio" class="dark:text-gray-200">Exercícios Ventilatórios</label>
                                    <Select 
                                        id="exercicio_ventilatorio" 
                                        v-model="scoreStore.score.exercicio_ventilatorio" 
                                        :options="props.formFields.exercicio_ventilatorio"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="uso_oxigenoterapia" class="dark:text-gray-200">Uso de Oxigenoterapia</label>
                                    <Select 
                                        id="uso_oxigenoterapia" 
                                        v-model="scoreStore.score.uso_oxigenoterapia" 
                                        :options="props.formFields.frequencia_respiratoria" 
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/3">
                                    <label for="nivel_consciencia" class="dark:text-gray-200">Nível de Consciência</label>
                                    <Select 
                                        id="nivel_consciencia" 
                                        v-model="scoreStore.score.nivel_consciencia" 
                                        :options="props.formFields.nivel_consciencia"
                                        optionLabel="label" 
                                        optionValue="value"
                                        placeholder="Selecione" 
                                        class="w-full"
                                    />
                                </div>
                            </div>

                            <div class="p-4 rounded-lg mt-4 border transition-all" :class="neadColorClass">
    <h3 class="text-lg font-bold mb-1">Resultado Final: {{ totalNead }} Pontos</h3>
    <p class="font-medium m-0">
        {{ totalNead <= 5 ? 'Outros Programas' : (totalNead <= 11 ? 'AD Multiprofissional' : (totalNead <= 17 ? 'Internação Domiciliar 12h' : 'Internação Domiciliar 24h')) }}
    </p>
</div>

                        </div>
                        <Button class="mt-3" v-on:click="handleSave">
                            <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
                            Próximo
                        </Button>
                    </AccordionTab>
                </Accordion>
            </template>