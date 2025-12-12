<script setup>
import { onMounted, computed, ref, watch } from 'vue';

import { useDadosGeraisStore } from '@/store/evaluation/dadosGerais';
import { useEnderecoStore } from '@/store/evaluation/endereco';
import { useContatoStore } from '@/store/evaluation/contato';
import { useCuidadoresStore } from '@/store/evaluation/cuidadores';
import { useSegurancaStore } from '@/store/evaluation/seguranca';
import { useHistoricoStore } from '@/store/evaluation/historico';
import { useRespiratorioStore } from '@/store/evaluation/cardiorespiratorio';
import { useNutricionalStore } from '@/store/evaluation/nutricional';
import { useEliminacoesStore } from '@/store/evaluation/eliminacoes';
import { useCondicoesPeleStore } from '@/store/evaluation/condicoesPele';
import { useScoreStore } from '@/store/evaluation/score';
import { useObservacoesStore } from '@/store/evaluation/observacao';
import { useExameFisicoStore } from '@/store/evaluation/exameFisico';
import { AvaliationService } from '@/service/AvaliationService';
import { useAvaliationForm } from '@/store/evaluation/form';

const dadosGeraisStore = useDadosGeraisStore();
const enderecoStore = useEnderecoStore();
const contatoStore = useContatoStore();
const cuidadoresStore = useCuidadoresStore();
const segurancaPacienteStore = useSegurancaStore();
const historicoClinicoStore = useHistoricoStore();
const exameFisicoStore = useExameFisicoStore();
const cardioRespiratorioStore = useRespiratorioStore();
const nutricionalStore = useNutricionalStore();
const eliminacoes = useEliminacoesStore();
const condicoesPele = useCondicoesPeleStore();
const score = useScoreStore();
const obs = useObservacoesStore();
const avaliationFormStore = useAvaliationForm();

const activeStep = ref(0);

const formData = computed(() => avaliationFormStore.avaliationForm);
onMounted(async () => {
    await AvaliationService.getFormData();
});

async function submitFormData() {
  const payload = {
      dadosGerais: dadosGeraisStore.dadosGerais,
      endereco: enderecoStore.endereco,
      contato: contatoStore.contato,
      cuidadores: cuidadoresStore.cuidadores,
      seguranca: segurancaPacienteStore.seguranca,
      historicoClinico: historicoClinicoStore.historico,
      exameFisico: exameFisicoStore.exameFisico,
      cardioRespiratorio: cardioRespiratorioStore.respiratorio,
      nutricional: nutricionalStore.nutricional,
      eliminacoes: eliminacoes.eliminacoes,
      condicoesPele: condicoesPele.condicoesPele,
      score: score.score,
      observacoes: obs.observacoes
  };
  
  await AvaliationService.submitFormData(payload)
}

function setStep(index) {
    if (activeStep.value === index) {
        activeStep.value = -1;
        
        setTimeout(() => {
            activeStep.value = index;
        }, 50);
    } else {
        activeStep.value = index;
    }
}

</script>

<template>
    <Fluid>
        <div v-if="formData">
            <DadosGerais :is-active="activeStep === 0" :formFields="formData.dados_gerais" @next-step="setStep(1)"/>
            <Endereço :is-active="activeStep === 1" @next-step="setStep(2)"/>
            <Contato :is-active="activeStep === 2" :formFields="formData.contato" @next-step="setStep(3)"/>
            <Cuidadores :is-active="activeStep === 3" :formFields="formData.cuidadores" @next-step="setStep(4)"/>
            <SegurançadoPaciente :is-active="activeStep === 4" :formFields="formData.seguranca" @next-step="setStep(5)"/>
            <HistóricoClínico :is-active="activeStep === 5" @next-step="setStep(6)"/>
            <ExameFísico :is-active="activeStep === 6" :formFields="formData.dados_clinicos" @next-step="setStep(7)"/>
            <Cardiorrespiratório :is-active="activeStep === 7" :formFields="formData.cardio_respiratorio" @next-step="setStep(8)"/>
            <Nutricional :is-active="activeStep === 8" :formFields="formData.nutricional" @next-step="setStep(9)"/>
            <Eliminações :is-active="activeStep === 9" :formFields="formData.eliminacoes" @next-step="setStep(10)"/>
            <CondiçõesdaPele :is-active="activeStep === 10" :formFields="formData.condicoes_pele" @next-step="setStep(11)"/>
            <Score :is-active="activeStep === 11" :formFields="formData.avaliacao" @next-step="setStep(12)"/>
            <Observações :is-active="activeStep === 12" @next-step="setStep(13)"/>
            <Button @click="submitFormData">Enviar</Button>
        </div>
        <div v-else>
            <p>Carregando formulário...</p>
        </div>
    </Fluid>
</template>