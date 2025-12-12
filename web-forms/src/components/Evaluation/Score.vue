    <script setup>
    import { ref, computed } from 'vue';
    import { useScoreStore } from '@/store/evaluation/score';

    const scoreStore = useScoreStore();
    const activeIndex = ref(null);

    const getPoints = (value, options) => {
    if (!value) return 0;

    let opt = options.find(o => o.code === value);

    if (!opt) {
        opt = options.find(o => o.name === value);
    }

    return opt ? opt.points : 0;
};

const props = defineProps({
  formFields: {
    type: Object,
    required: false,
    default: null
  }
})


    const freq12hOpts = ref([
        { name: 'Não utiliza', code: 'nao_utiliza', points: 0 },
        { name: 'Até 12 horas por dia', code: 'ate_12h', points: 0 },
        { name: 'Mais de 12 horas por dia', code: 'mais_12h', points: 0 }
    ]);

    const freq5xOpts = ref([
        { name: 'Não utiliza', code: 'nao_utiliza', points: 0 },
        { name: 'Até 5x por dia', code: 'ate_5x', points: 0 },
        { name: 'Mais que 5x por dia', code: 'mais_5x', points: 0 }
    ]);

    const freq4xOpts = ref([
        { name: 'Não utiliza', code: 'nao_utiliza', points: 0 },
        { name: 'Até 4x por dia', code: 'ate_4x', points: 0 },
        { name: 'Mais que 4x por dia', code: 'mais_4x', points: 0 }
    ]);

    const katzOpts = ref([
    { name: 'Independente (Sem ajuda)', code: 'independente', points: 1 }, // Mudado de 'false' para string
    { name: 'Com ajuda / Dependente', code: 'dependente', points: 0 } // Mudado de 'true' para string
]);
    const estadoNutricionalOpts = ref([
        { name: 'Eutrófico', code: 'eutrofico', points: 0 },
        { name: 'Sobrepeso/Emagrecido', code: 'sobrepeso_emagrecido', points: 1 },
        { name: 'Obeso/Desnutrido', code: 'obeso_desnutrido', points: 2 }
    ]);

    const viaAlimentacaoOpts = ref([
        { name: 'Sem auxílio', code: 'sem_auxilio', points: 0 },
        { name: 'Assistida', code: 'assistida', points: 1 },
        { name: 'Gastrostomia/Jejunostomia', code: 'gastro_jejuno', points: 2 },
        { name: 'SNG/SNE', code: 'sng_sne', points: 3 }
    ]);

    const internacoesOpts = ref([
        { name: '0-1', code: '0-1', points: 0 },
        { name: '2-3', code: '2-3', points: 1 },
        { name: '4+', code: '4+', points: 2 }
    ]);

    const aspiracaoViasAereasOpts = ref([
        { name: 'Ausente', code: 'ausente', points: 0 },
        { name: 'Até 5x por dia', code: 'ate_5x', points: 1 },
        { name: 'Mais que 5x por dia', code: 'mais_5x', points: 2 }
    ]);

    const lesoesOpts = ref([
        { name: 'Nenhuma ou Única Simples', code: 'nenhuma_simples', points: 0 },
        { name: 'Múltiplas (simples) ou Única (complexo)', code: 'multiplas_simples_unica_complexa', points: 1 },
        { name: 'Múltiplas lesões (complexas)', code: 'multiplas_complexas', points: 2 }
    ]);

    const viaMedicacoesOpts = ref([
        { name: 'Via enteral', code: 'enteral', points: 0 },
        { name: 'Intramuscular ou subcutânea', code: 'im_sc', points: 1 },
        { name: 'IV até 4x/dia ou Hipodermóclise', code: 'iv_hipo', points: 2 }
    ]);

    const frequenciaRespOpts = ref([
        { name: 'Ausente', code: 'ausente', points: 0 },
        { name: 'Intermitente', code: 'intermitente', points: 1 },
        { name: 'Contínuo', code: 'continuo', points: 2 }
    ]);

    const nivelConscienciaOpts = ref([
        { name: 'Alerta', code: 'alerta', points: 0 },
        { name: 'Confuso/Desorientado', code: 'confuso', points: 1 },
        { name: 'Comatoso', code: 'comatoso', points: 2 }
    ]);

    // ... (outras computed properties e imports)

// Funções auxiliares (opcional, mas melhora a clareza)
const getGroup2Priority = (value, opts) => {
    // 1. Tenta encontrar a opção pelo 'code' (que é o que está salvo no store)
    const opt = opts.find(o => o.code === value);
    if (!opt) return 0; // Se não selecionado ou 'nao_utiliza' (0)

    // 2. Mapeamento da prioridade de complexidade:
    // Mais de 12h / Mais de 5x / Mais de 4x = 3 (24h)
    if (opt.name.includes('Mais de 12 horas') || opt.name.includes('Mais que 5x') || opt.name.includes('Mais que 4x')) {
        return 3; // 24 HORAS
    }
    // Até 12h / Até 5x / Até 4x = 2 (12h)
    if (opt.name.includes('Até 12 horas') || opt.name.includes('Até 5x') || opt.name.includes('Até 4x')) {
        return 2; // 12 HORAS (ou AD/Outros Programas para a Medicação)
    }

    // Não utiliza = 1 (AD / Outros Programas)
    return 1;
};


const recomendacaoGrupo2 = computed(() => {
    const s = scoreStore.score;
    let maxPriority = 0;


    maxPriority = Math.max(maxPriority, getGroup2Priority(s.alimentacao_parenteral, freq12hOpts.value));


    maxPriority = Math.max(maxPriority, getGroup2Priority(s.aspiracao_traqueo, freq5xOpts.value));

  
    maxPriority = Math.max(maxPriority, getGroup2Priority(s.ventilacao_mecanica, freq12hOpts.value));


    if (s.medicacao_parenteral === 'mais_4x') {
        maxPriority = Math.max(maxPriority, 3); // 24 HORAS
    } else if (s.medicacao_parenteral === 'ate_4x') {
        maxPriority = Math.max(maxPriority, 1);
    }
    

    if (maxPriority === 3) {
        return 'Internação Domiciliar 24h';
    } else if (maxPriority === 2) {
        return 'Internação Domiciliar 12h';
    } else if (maxPriority === 1) {
        return 'Atendimento Domiciliar / Outros Programas';
    } else {
        return 'Nenhuma indicação imediata no Grupo 2';
    }
});


    const totalKatzRaw = computed(() => {
        let total = 0;
        const s = scoreStore.score;
        total += getPoints(s.banho, katzOpts.value);
        total += getPoints(s.vestir, katzOpts.value);
        total += getPoints(s.higiene_pessoal, katzOpts.value);
        total += getPoints(s.transferencia, katzOpts.value);
        total += getPoints(s.continencia, katzOpts.value);
        total += getPoints(s.alimentacao, katzOpts.value);
        return total;
    });

    const classificacaoKatz = computed(() => {
        const pts = totalKatzRaw.value;
        if (pts >= 5) return 'Independente';
        if (pts >= 3) return 'Dependência Parcial';
        return 'Dependente Total';
    });


    const katzColorClass = computed(() => {
        const pts = totalKatzRaw.value;
        if (pts >= 5) { 
            return 'bg-green-50 border-green-200 text-green-900 dark:bg-green-900/30 dark:border-green-800 dark:text-green-300';
        }
        if (pts >= 3) { 
            return 'bg-orange-50 border-orange-200 text-orange-900 dark:bg-orange-900/30 dark:border-orange-800 dark:text-orange-300';
        }
        
        return 'bg-red-50 border-red-200 text-red-900 dark:bg-red-900/30 dark:border-red-800 dark:text-red-300';
    });

    const pontosNeadKatz = computed(() => {
        const pts = totalKatzRaw.value;
        if (pts >= 5) return 0;
        if (pts >= 3) return 1;
        return 2;
    });

    const totalNead = computed(() => {
    let total = 0;
    const s = scoreStore.score;
    
    // LINHAS REMOVIDAS (GRUPO 2):
    // total += getPoints(s.alimentacao_parenteral, freq12hOpts.value);
    // total += getPoints(s.ventilacao_mecanica, freq12hOpts.value);
    // total += getPoints(s.aspiracao_traqueo, freq5xOpts.value);
    // total += getPoints(s.medicacao_parenteral, freq4xOpts.value);

    // GRUPO 3 (CRITÉRIOS DE APOIO) e KATZ são mantidos:
    total += getPoints(s.estado_nutricional, estadoNutricionalOpts.value);
    total += getPoints(s.via_alimentacao_medicacao, viaAlimentacaoOpts.value);
    total += getPoints(s.internacoes_ultimo_ano, internacoesOpts.value);
    total += getPoints(s.aspiracao_vias_aereas, aspiracaoViasAereasOpts.value);
    total += getPoints(s.lesoes, lesoesOpts.value);
    total += getPoints(s.via_medicacoes, viaMedicacoesOpts.value);
    total += getPoints(s.exercicios_ventilatorios, frequenciaRespOpts.value);
    total += getPoints(s.uso_oxigenoterapia, frequenciaRespOpts.value);
    total += getPoints(s.nivel_consciencia, nivelConscienciaOpts.value);
    total += pontosNeadKatz.value;
    return total;
});    

    const classificacaoFinal = computed(() => {
        const pts = totalNead.value;
        if (pts <= 5) return 'Até 5 pts: Procedimentos Pontuais / Outros Programas';
        if (pts <= 11) return '6 a 11 pts: AD Multiprofissional';
        if (pts <= 17) return '12 a 17 pts: Internação Domiciliar 12h';
        return '18+ pts: Internação Domiciliar 24h';
    });


    const neadColorClass = computed(() => {
        if (totalNead.value >= 12) {
           
            return 'bg-red-50 border-red-200 text-red-900 dark:bg-red-900/30 dark:border-red-800 dark:text-red-300';
        }
            return 'bg-green-50 border-green-200 text-green-900 dark:bg-green-900/30 dark:border-green-800 dark:text-green-300';
    });

    const isFilled = computed(() => {
        return !!scoreStore.score.diagnostico_primario && 
            scoreStore.score.diagnostico_primario.length > 3;
    });

    const handleSave = () => {
        activeIndex.value = null;
        setTimeout(() => {
            const self = document.getElementById("score");
            if (self) self.scrollIntoView({ behavior: 'instant', block: 'start' });
        }, 0);
        emit('next-step');
    };


    </script>

    <template>
        <Accordion v-model:activeIndex="activeIndex" id="score" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600 dark:bg-surface-900">
            <AccordionTab>
                <template #header>
                    <div class="flex items-center gap-3 w-full">
                        <i class="pi text-xl" 
                            :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
                        </i>
                        <div class="flex flex-col text-left">
                            <h4 class="font-semibold text-xl p-0 m-0 dark:text-gray-100">Score NEAD e Katz</h4>
                            <span class="text-xs text-gray-500 font-normal -mt-4 dark:text-gray-400">
                                {{ isFilled ? `Pontuação: ${totalNead}` : 'Toque para preencher' }}
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
                                optionValue="label"
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
                                optionValue="label"
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
                                optionValue="label"
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
                                optionValue="label"
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                    </div>
                    <div class="flex flex-col md:flex-row gap-4">Para indicação de Planejamento de Atenção Domiciliar (P.A.D.) Recomendação atual: {{ recomendacaoGrupo2 }}    </div>

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
                                optionValue="label"
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
                                optionValue="label"
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
                                optionValue="label"
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
                                optionValue="label"
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
                                optionValue="label"
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
                                optionValue="label"
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                    </div>

                    <!-- CARD RESULTADO KATZ -->
                    <div class="p-4 rounded-lg mt-4 border" :class="katzColorClass">
                        <h3 class="text-lg font-bold mb-1">Resultado Katz: {{ totalKatzRaw }} / 6</h3>
                        <p class="font-medium m-0">{{ classificacaoKatz }} </p>
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
                                optionValue="label"
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="via_alimentacao_medicacao" class="dark:text-gray-200">Aliment / Medic Enteral</label>
                            <Select 
                                id="via_alimentacao_medicacao" 
                                v-model="scoreStore.score.via_alimentacao_medicacao" 
                                :options="props.formFields.via_aliementacao"
                                optionLabel="label" 
                                optionValue="label"
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="internacoes_ultimo_ano" class="dark:text-gray-200">Internações no último ano</label>
                            <Select 
                                id="internacoes_ultimo_ano" 
                                v-model="props.formFields.internacoes_ultimo_ano" 
                                :options="props.formFields.internacoes_ultimo_ano"
                                optionLabel="label" 
                                optionValue="label"
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
                                optionValue="label"
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="lesoes" class="dark:text-gray-200">Lesões</label>
                            <Select 
                                id="lesoes" 
                                :options="props.formFields.lesoes"
                                optionLabel="label" 
                                optionValue="label"
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
                                optionValue="label"
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
                                optionValue="label"
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="uso_oxigenoterapia" class="dark:text-gray-200">Uso de Oxigenoterapia</label>
                            <Select 
                                id="uso_oxigenoterapia" 
                                v-model="scoreStore.score.uso_oxigenoterapia" 
                                :options="props.formFields.avaliacao.frequencia_respiratoria" 
                                optionLabel="label" 
                                optionValue="label"
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
                                optionValue="label"
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                    </div>

                    <div class="p-4 rounded-lg mt-4 border" :class="neadColorClass">
                        <h3 class="text-lg font-bold mb-1">Resultado Final: {{ totalNead }} Pontos</h3>
                        <p class="font-medium m-0">{{ classificacaoFinal }}</p>
                    </div>

                </div>
                <Button class="mt-3" v-on:click="handleSave">
                    <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
                    Próximo
                </Button>
            </AccordionTab>
        </Accordion>
    </template>