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

const activeStep = ref(0);
const isLoading = ref(false);

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

async function handleNextStep() {
    isLoading.value = true;
    
    try {
        const payload = {
            dadosGerais: dadosGeraisStore.dadosGerais,
        };

        // await AvaliationService.submitFormData(payload); // UNCOMMENT TO ENABLE SAVE
        console.log("Saving to database..."); 

        activeStep.value = activeStep.value + 1;        
        window.scrollTo({ top: 0, behavior: 'smooth' });

    } catch (error) {
        console.error("Error saving:", error);
        alert("Erro ao salvar dados.");
    } finally {
        isLoading.value = false;
    }
}

</script>

<template>
    <Fluid>
        <div v-if="formData">
            <DadosGerais :formFields="formData.dados_gerais"/>
            <Endereço/>
            <Contato :formFields="formData.contato"/>
            <Cuidadores :formFields="formData.cuidadores"/>
            <SegurançadoPaciente :formFields="formData.seguranca"/>
            <HistóricoClínico/>
            <ExameFísico :formFields="formData.dados_clinicos"/>
            <Cardiorrespiratório :formFields="formData.cardio_respiratorio"/>
            <Nutricional  :formFields="formData.nutricional"/>
            <Eliminações  :formFields="formData.eliminacoes"/>
            <CondiçõesdaPele  :formFields="formData.condicoes_pele"/> <p>em condicao o unico select que tem não esta no bruno,
                 imagino que por causa do jeito que o front é

            </p>
            <Score :formFields="formData.avaliacao" />
            <Observações/>
            <Button @click="submitFormData">Enviar</Button>
        </div>
        <div v-else>
            <p>Carregando formulário...</p>
        </div>
    </Fluid>
</template>