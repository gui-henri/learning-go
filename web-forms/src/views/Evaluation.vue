<script setup>
import { onMounted, computed } from 'vue';

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

const formData = computed(() => avaliationFormStore.avaliationForm);
onMounted(async () => {
    await AvaliationService.getFormData(avaliationFormStore);
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

</script>

<template>
    <Fluid>
        <div v-if="formData">
            <DadosGerais :formFields="formData.dados_gerais"/>
            <Endereço/>
            <Contato :formFields="formData.contato"/>
            <Cuidadores/>
            <SegurançadoPaciente/>
            <HistóricoClínico/>
            <ExameFísico/>
            <Cardiorrespiratório />
            <Nutricional/>
            <Eliminações/>
            <CondiçõesdaPele/>
            <Score/>
            <Observações/>
            <Button @click="submitFormData">Enviar</Button>
        </div>
        <div v-else>
            <p>Carregando formulário...</p>
        </div>
    </Fluid>
</template>