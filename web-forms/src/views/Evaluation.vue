
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
<<<<<<< HEAD
import SegurançadoPaciente from '@/components/Evaluation/SegurançadoPaciente.vue';
=======
import Cardiorrespiratório from '@/components/Evaluation/Cardiorrespiratório.vue';
>>>>>>> e75a1840d723a6f29d973465d621bdd16beb8e6e

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

        <DadosGerais/>
        <Endereço/>
        <Contato/>
        <Cuidadores/>
        <SegurançadoPaciente/>
        <HistóricoClínico/>
       
        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Exame Físico</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="estado_geral">Estado geral</label>
                        <Select 
                            id="estado_geral" 
                            v-model="exameFisicoStore.exameFisico.estado_geral" 
                            :options="estado_geral" 
                            optionLabel="name" 
                            placeholder="Selecione" 
                            class="w-full"
                        ></Select>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="avaliacao_locomotora">Avaliação locomotora</label>
                        <Select 
                            id="avaliacao_locomotora" 
                            v-model="exameFisicoStore.exameFisico.avaliacao_locomotora" 
                            :options="locomocao" 
                            optionLabel="name" 
                            placeholder="Selecione" 
                            class="w-full"
                        ></Select>
                    </div>
                </div>

                <div class="flex flex-col gap-4 border-t pt-4">
                    <div class="flex flex-col md:flex-row gap-4 items-center">
                        <div class="flex flex-col gap-2 w-full md:w-1/6">
                            <label for="possui_dreno">Drenos?</label>
                            <InputSwitch 
                                id="possui_dreno" 
                                v-model="exameFisicoStore.exameFisico.possui_dreno" 
                            />
                        </div>
                    </div>

                    <div v-if="exameFisicoStore.exameFisico.possui_dreno" class="flex flex-col md:flex-row gap-4 p-4 rounded-lg">
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label for="local_dreno">Local do dreno</label>
                            <InputText 
                                id="local_dreno" 
                                v-model="exameFisicoStore.exameFisico.local_dreno" 
                                placeholder="Ex: Braço direito" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/6">
                            <label for="data_implantacao_dreno">Data implantação</label>
                            <Calendar 
                                id="data_implantacao_dreno" 
                                v-model="exameFisicoStore.exameFisico.data_implantacao_dreno" 
                                dateFormat="dd/mm/yy" 
                                placeholder="dd/mm/aaaa" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full flex-1">
                            <label for="curativos_dreno">Curativos do dreno</label>
                            <InputText 
                                id="curativos_dreno" 
                                v-model="exameFisicoStore.exameFisico.curativos_dreno" 
                                placeholder="Descrição dos curativos" 
                                class="w-full"
                            />
                        </div>
                    </div>
                </div>
                <div class="flex flex-col gap-4 border-t pt-4">
                
                        <div class="flex flex-col gap-2 w-full md:w-1/6">
                            <label for="possui_acessos">Acessos?</label>
                            <InputSwitch 
                                id="possui_acesso" 
                                v-model="exameFisicoStore.exameFisico.possui_acesso" 
                            />
                        </div>
                        </div>
                        <div v-if="exameFisicoStore.exameFisico.possui_dreno" class="flex flex-col md:flex-row gap-4 p-4 rounded-lg">
                            <div class="flex flex-col gap-2 w-full md:w-1/2">
                                <label for="estado_geral">Estado geral</label>
                                <Select 
                                id="estado_geral" 
                                v-model="exameFisicoStore.exameFisico.estado_geral" 
                                :options="estado_geral" 
                                optionLabel="name" 
                                placeholder="Selecione" 
                                class="w-full"
                            ></Select>
                            </div>    
                </div>
            </div>
        </div>
        <Cardiorrespiratório />
    </Fluid>
</template>
```