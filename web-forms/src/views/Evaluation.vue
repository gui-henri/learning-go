<script setup>
import { ref } from 'vue';
import DadosGerais from '@/components/Evaluation/DadosGerais.vue';
import Endereço from '@/components/Evaluation/Endereço.vue';
import { useDadosGeraisStore } from '@/store/evaluation/dadosGerais';
import { useEnderecoStore } from "@/store/evaluation/endereco";
import { useContatoStore } from '@/store/evaluation/contato';
import { useCuidadoresStore } from '@/store/evaluation/cuidadores';
import { useSegurancaStore } from '@/store/evaluation/seguranca';
import { historicoModel, useHistoricoStore } from '@/store/evaluation/historico';
import { useExameFisicoStore } from '@/store/evaluation/exameFisico';
import Cardiorrespiratório from '@/components/Evaluation/Cardiorrespiratório.vue';
import Nutricional from '@/components/Evaluation/Nutricional.vue';
import Eliminações from '@/components/Evaluation/Eliminações.vue';
import CondiçõesdaPele from '@/components/Evaluation/CondiçõesdaPele.vue';
import Observações from '@/components/Evaluation/Observações.vue';

const dadosGeraisStore = useDadosGeraisStore();
const enderecoStore = useEnderecoStore();
const contatoStore = useContatoStore();
const cuidadorStore = useCuidadoresStore();
const segurancaStore = useSegurancaStore();
const historicoStore = useHistoricoStore();
const exameFisicoStore = useExameFisicoStore();

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
           

    </Fluid>
</template>