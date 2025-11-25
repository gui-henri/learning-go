<script setup>
import DadosGerais from '@/components/Evaluation/DadosGerais.vue';
import Endereço from '@/components/Evaluation/Endereço.vue';
import Cardiorrespiratório from '@/components/Evaluation/Cardiorrespiratório.vue';
import Nutricional from '@/components/Evaluation/Nutricional.vue';
import Eliminações from '@/components/Evaluation/Eliminações.vue';
import CondiçõesdaPele from '@/components/Evaluation/CondiçõesdaPele.vue';
import Observações from '@/components/Evaluation/Observações.vue';
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