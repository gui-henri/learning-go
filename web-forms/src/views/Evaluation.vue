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

const forma_contato = ref([
    { name: 'Residencial', code: 'residencial' },
    { name: 'Celular', code: 'celular' },
    { name: 'Whatsapp', code: 'whatsapp' },
    { name: 'Celular e Whatsapp', code: 'celular_whatsapp' }
]);

const turnos_cuidador = ref([
    { name: 'Diurno', code: 'diurno' },
    { name: 'Noturno', code: 'noturno' },
    { name: '24h', code: '24h' }
]);

const tipo_alergia = ref([
    { name: 'Medicamento', code: 'medicamento' },
    { name: 'Alimento', code: 'alimento' }
]);

const tipo_precaucao = ref([
    { name: 'Padrão', code: 'padrao' },
    { name: 'Contato', code: 'contato' },
    { name: 'Gotícula', code: 'goticula' },
    { name: 'Aerossóis', code: 'aerossois' },
    { name: 'Reverso', code: 'reverso' }
]);

const estado_geral = ref([
    { name: 'Bom', code: 'bom' },
    { name: 'Regular', code: 'regular' },
    { name: 'Grave', code: 'grave' },
    { name: 'Gravíssimo', code: 'gravissimo' },
    { name: 'Decaído', code: 'decaido' }
]);

const locomocao = ref([
    { name: 'Deambula sem apoio', code: 'deambula_sem_apoio' },
    { name: 'Deambula com apoio', code: 'deambula_com_apoio' },
    { name: 'Restrito ao leito', code: 'restrito_leito' },
    { name: 'Cadeirante', code: 'cadeirante' }
]);

const piccline = ref([
    { name: 'Monolumen', code: 'monolumen' },
    { name: 'Duplolumen', code: 'duplolumen' },
    { name: 'Triplolumen', code: 'triplolumen' }
]);

</script>

<template>
    <Fluid>
        <div class="flex flex-col gap-1 w-full">

            <DadosGerais/>
            <Endereço/>
            <Contato/>
            <Cuidadores/>
            <SegurançadoPaciente/>
            <HistóricoClínico/>
            <ExameFísico/>
            <Cardiorrespiratório />
            <Nutricional/>
           

        </div> 
    </Fluid>
</template>