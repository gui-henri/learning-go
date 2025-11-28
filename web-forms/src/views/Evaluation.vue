<script setup>
import DadosGerais from '@/components/Evaluation/DadosGerais.vue';
import Endereço from '@/components/Evaluation/Endereço.vue';
import Cardiorrespiratório from '@/components/Evaluation/Cardiorrespiratório.vue';
import Nutricional from '@/components/Evaluation/Nutricional.vue';
import Eliminações from '@/components/Evaluation/Eliminações.vue';
import CondiçõesdaPele from '@/components/Evaluation/CondiçõesdaPele.vue';
import Observações from '@/components/Evaluation/Observações.vue';
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

const dadosGeraisStore = useDadosGeraisStore();
const enderecoStore = useEnderecoStore();
const contatoStore = useContatoStore();
const cuidadoresStore = useCuidadoresStore();
const segurancaPacienteStore = useSegurancaStore();
const historicoClinicoStore = useHistoricoStore();
const exameFisicoStore = useHistoricoStore();
const cardioRespiratorioStore = useRespiratorioStore();
const nutricionalStore = useNutricionalStore();
const eliminacoes = useEliminacoesStore();
const condicoesPele = useCondicoesPeleStore();
const score = useScoreStore();
const obs = useObservacoesStore();

async function submitFormData() {
  try {
    const payload = {
      dadosGerais: dadosGeraisStore.$state,
      endereco: enderecoStore.$state,
      contato: contatoStore.$state,
      cuidadores: cuidadoresStore.$state,
      seguranca: segurancaPacienteStore.$state,
      historicoClinico: historicoClinicoStore.$state,
      exameFisico: exameFisicoStore.$state,
      cardioRespiratorio: cardioRespiratorioStore.$state,
      nutricional: nutricionalStore.$state,
      eliminacoes: eliminacoes.$state,
      condicoesPele: condicoesPele.$state,
      score: score.$state,
      observacoes: obs.$state
    };

    const response = await fetch('http://localhost:8090/Avaliation', { 
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload)
    });

    if (!response.ok) {
      throw new Error(`Error: ${response.statusText}`);
    }

    const result = await response.json();
    console.log('Success:', result);
    alert('Formulário salvo com sucesso!');

  } catch (error) {
    console.error('Failed to save form:', error);
    alert('Erro ao salvar o formulário.');
  }
}


</script>

<template>
    <Fluid>
        <DadosGerais/>
        <Endereço/>
        <Contato/>
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
        <Button v-on:click="submitFormData">Enviar</Button>
    </Fluid>
</template>