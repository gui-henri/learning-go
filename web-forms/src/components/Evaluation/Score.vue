<script setup>
import { ref, computed } from 'vue';
import { useScoreStore } from '@/store/evaluation/score';

const scoreStore = useScoreStore();
const activeIndex = ref(null);

const freq12hOpts = ref([
    { name: 'Não utiliza', code: 'nao_utiliza' },
    { name: 'Até 12 horas por dia', code: 'ate_12h' },
    { name: 'Mais de 12 horas por dia', code: 'mais_12h' }
]);

const freq5xOpts = ref([
    { name: 'Não utiliza', code: 'nao_utiliza' },
    { name: 'Até 5x por dia', code: 'ate_5x' },
    { name: 'Mais que 5x por dia', code: 'mais_5x' }
]);

const freq4xOpts = ref([
    { name: 'Não utiliza', code: 'nao_utiliza' },
    { name: 'Até 4x por dia', code: 'ate_4x' },
    { name: 'Mais que 4x por dia', code: 'mais_4x' }
]);


const katzOpts = ref([
    { name: 'Independente', code: 'independente' },
    { name: 'Com ajuda', code: 'com_ajuda' },
    { name: 'Dependente', code: 'dependente' }
]);


const estadoNutricionalOpts = ref([
    { name: 'Eutrófico', code: 'eutrofico' },
    { name: 'Sobrepeso/Emagrecido', code: 'sobrepeso_emagrecido' },
    { name: 'Obeso/Desnutrido', code: 'obeso_desnutrido' }
]);

const viaAlimentacaoOpts = ref([
    { name: 'Sem auxílio', code: 'sem_auxilio' },
    { name: 'Assistida', code: 'assistida' },
    { name: 'Gastrostomia/Jejunostomia', code: 'gastro_jejuno' },
    { name: 'SNG/SNE', code: 'sng_sne' }
]);

const internacoesOpts = ref([
    { name: '0-1', code: '0-1' },
    { name: '2-3', code: '2-3' },
    { name: '4+', code: '4+' }
]);

const aspiracaoViasAereasOpts = ref([
    { name: 'Ausente', code: 'ausente' },
    { name: 'Até 5x por dia', code: 'ate_5x' },
    { name: 'Mais que 5x por dia', code: 'mais_5x' }
]);

const lesoesOpts = ref([
    { name: 'Nenhuma ou lesão única (curativo simples)', code: 'nenhuma_simples' },
    { name: 'Múltiplas (simples) ou Única (complexo)', code: 'multiplas_simples_unica_complexa' },
    { name: 'Múltiplas lesões (complexas)', code: 'multiplas_complexas' }
]);

const viaMedicacoesOpts = ref([
    { name: 'Via enteral', code: 'enteral' },
    { name: 'Intramuscular ou subcutânea', code: 'im_sc' },
    { name: 'Intravenosa até 4x/dia ou hipodermóclise', code: 'iv_hipo' }
]);

const frequenciaRespOpts = ref([
    { name: 'Ausente', code: 'ausente' },
    { name: 'Intermitente', code: 'intermitente' },
    { name: 'Contínuo', code: 'continuo' }
]);

const nivelConscienciaOpts = ref([
    { name: 'Alerta', code: 'alerta' },
    { name: 'Confuso/Desorientado', code: 'confuso' },
    { name: 'Comatoso', code: 'comatoso' }
]);

const isFilled = computed(() => {
    return !!scoreStore.score.diagnostico_primario && 
           scoreStore.score.diagnostico_primario.length > 3;
});

const handleSave = () => {
    activeIndex.value = null;
    
    setTimeout(() => {
        const self = document.getElementById("score");
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
    <Accordion v-model:activeIndex="activeIndex" id="score" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Score NEAD e Katz</h4>
                <span class="text-xs text-gray-500 font-normal -mt-4">
                    {{ isFilled ? '' : 'Toque para preencher' }}
                </span>
            </div>
        </div>
    </template>
        <div class="flex flex-col gap-6 w-full">

            <h5 class="font-medium text-gray-700 border-b pb-2">Diagnóstico e Elegibilidade (Grupo 1)</h5>
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/2">
                    <label for="diagnostico_primario">Diagnóstico Primário</label>
                    <InputText 
                        id="diagnostico_primario" 
                        v-model="scoreStore.score.diagnostico_primario" 
                        placeholder="CID ou Descrição" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/2">
                    <label for="diagnostico_secundario">Diagnóstico Secundário</label>
                    <InputText 
                        id="diagnostico_secundario" 
                        v-model="scoreStore.score.diagnostico_secundario" 
                        placeholder="CID ou Descrição" 
                        class="w-full"
                    />
                </div>
            </div>
            
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="domicilio_risco">O domicílio é de risco?</label>
                    <InputSwitch 
                        id="domicilio_risco" 
                        v-model="scoreStore.score.domicilio_risco" 
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="impedimento_deslocamento">Impedimento de deslocamento?</label>
                    <InputSwitch 
                        id="impedimento_deslocamento" 
                        v-model="scoreStore.score.impedimento_deslocamento" 
                    />
                </div>
            </div>

            <h5 class="font-medium text-gray-700 border-b pb-2 mt-2">Indicação Imediata (Grupo 2)</h5>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div class="flex flex-col gap-2">
                    <label for="alimentacao_parenteral">Alimentação Parenteral</label>
                    <Select 
                        id="alimentacao_parenteral" 
                        v-model="scoreStore.score.alimentacao_parenteral" 
                        :options="freq12hOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="aspiracao_traqueo">Aspiração Traqueostomia</label>
                    <Select 
                        id="aspiracao_traqueo" 
                        v-model="scoreStore.score.aspiracao_traqueo" 
                        :options="freq5xOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="ventilacao_mecanica">Ventilação Mecânica (Inv. ou não)</label>
                    <Select 
                        id="ventilacao_mecanica" 
                        v-model="scoreStore.score.ventilacao_mecanica" 
                        :options="freq12hOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="medicacao_parenteral">Medicação Parenteral / Hipodermóclise</label>
                    <Select 
                        id="medicacao_parenteral" 
                        v-model="scoreStore.score.medicacao_parenteral" 
                        :options="freq4xOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
            </div>

            <h5 class="font-medium text-gray-700 border-b pb-2 mt-2">Escore de Katz</h5>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div class="flex flex-col gap-2">
                    <label for="banho">Banhar-se</label>
                    <Select 
                        id="banho" 
                        v-model="scoreStore.score.banho" 
                        :options="katzOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="vestir">Vestir-se</label>
                    <Select 
                        id="vestir" 
                        v-model="scoreStore.score.vestir" 
                        :options="katzOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="higiene_pessoal">Ir ao banheiro</label>
                    <Select 
                        id="higiene_pessoal" 
                        v-model="scoreStore.score.higiene_pessoal" 
                        :options="katzOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="transferencia">Transferência</label>
                    <Select 
                        id="transferencia" 
                        v-model="scoreStore.score.transferencia" 
                        :options="katzOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="continencia">Continência</label>
                    <Select 
                        id="continencia" 
                        v-model="scoreStore.score.continencia" 
                        :options="katzOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="alimentacao">Alimentação</label>
                    <Select 
                        id="alimentacao" 
                        v-model="scoreStore.score.alimentacao" 
                        :options="katzOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
            </div>

            <h5 class="font-medium text-gray-700 border-b pb-2 mt-2">Critérios de Apoio (Grupo 3)</h5>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div class="flex flex-col gap-2">
                    <label for="estado_nutricional">Estado Nutricional</label>
                    <Select 
                        id="estado_nutricional" 
                        v-model="scoreStore.score.estado_nutricional" 
                        :options="estadoNutricionalOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="via_alimentacao_medicacao">Alimentação/Medicação Enteral</label>
                    <Select 
                        id="via_alimentacao_medicacao" 
                        v-model="scoreStore.score.via_alimentacao_medicacao" 
                        :options="viaAlimentacaoOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="internacoes_ultimo_ano">Internações no último ano</label>
                    <Select 
                        id="internacoes_ultimo_ano" 
                        v-model="scoreStore.score.internacoes_ultimo_ano" 
                        :options="internacoesOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="aspiracao_vias_aereas">Aspiração Vias Aéreas Sup.</label>
                    <Select 
                        id="aspiracao_vias_aereas" 
                        v-model="scoreStore.score.aspiracao_vias_aereas" 
                        :options="aspiracaoViasAereasOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="lesoes">Lesões</label>
                    <Select 
                        id="lesoes" 
                        v-model="scoreStore.score.lesoes" 
                        :options="lesoesOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="via_medicacoes">Medicações (Via)</label>
                    <Select 
                        id="via_medicacoes" 
                        v-model="scoreStore.score.via_medicacoes" 
                        :options="viaMedicacoesOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="exercicios_ventilatorios">Exercícios Ventilatórios</label>
                    <Select 
                        id="exercicios_ventilatorios" 
                        v-model="scoreStore.score.exercicios_ventilatorios" 
                        :options="frequenciaRespOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="uso_oxigenoterapia">Uso de Oxigenoterapia</label>
                    <Select 
                        id="uso_oxigenoterapia" 
                        v-model="scoreStore.score.uso_oxigenoterapia" 
                        :options="frequenciaRespOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2">
                    <label for="nivel_consciencia">Nível de Consciência</label>
                    <Select 
                        id="nivel_consciencia" 
                        v-model="scoreStore.score.nivel_consciencia" 
                        :options="nivelConscienciaOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
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