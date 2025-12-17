            <script setup>
            import { ref, computed } from 'vue';
            import { useScoreStore } from '@/store/evaluation/score';
            import { useStepAccordion } from "@/composable/useStepAccordion";
            import { useAvaliationForm } from "@/store/evaluation/form";
            import { AvaliationService } from '@/service/AvaliationService';

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
                                    <InputSwitch 
                                        id="domicilio_risco" 
                                        v-model="scoreStore.score.domicilio_risco" 
                                    />
                                </div>
                                <div class="flex flex-col gap-2 w-full md:w-1/2">
                                    <label for="impedimento_deslocamento" class="dark:text-gray-200">Impedimento de deslocamento?</label>
                                    <InputSwitch 
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
                            <div class="flex flex-col md:flex-row gap-4">Para indicação de Planejamento de Atenção Domiciliar (P.A.D.) Recomendação atual: RECOMENDAÇÃO DE ACORDO COM GRUPO 2    </div>

                            <!-- KATZ -->
                            <div class="flex justify-between items-center border-b pb-2 mt-4 dark:border-gray-700">
                                <h5 class="font-medium text-gray-700 m-0 dark:text-gray-300">Escore de Katz</h5>
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
                            <div class="p-4 rounded-lg mt-4 border" :class="katzColorClass">
                                <h3 class="text-lg font-bold mb-1">Resultado Katz: TOTALKATZ / 6</h3>
                                <p class="font-medium m-0">CLASSIFICAO KATZ</p>
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
                                    <label for="exercicios_ventilatorios" class="dark:text-gray-200">Exercícios Ventilatórios</label>
                                    <Select 
                                        id="exercicios_ventilatorios" 
                                        v-model="scoreStore.score.exercicios_ventilatorios" 
                                        :options="props.formFields.exercicios_ventilatorios"
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

                            <div class="p-4 rounded-lg mt-4 border" :class="neadColorClass">
                                <h3 class="text-lg font-bold mb-1">Resultado Final: TOTAL NEAD Pontos</h3>
                                <p class="font-medium m-0">CLASSIFICAO NEAD FINAL</p>
                            </div>

                        </div>
                        <Button class="mt-3" v-on:click="handleSave">
                            <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
                            Próximo
                        </Button>
                    </AccordionTab>
                </Accordion>
            </template>