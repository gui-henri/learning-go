<script setup>
import { ref } from 'vue';
import { useEliminacoesStore } from '@/store/evaluation/eliminacoes';

const eliminacoesStore = useEliminacoesStore();

const funcaoIntestinalOpts = ref([
    { name: 'Normal', code: 'normal' },
    { name: 'Constipado', code: 'constipado' },
    { name: 'Com diarreia', code: 'com_diarreia' }
]);

const tipoEstomiaOpts = ref([
    { name: 'Ileostomia', code: 'ileostomia' },
    { name: 'Colostomia', code: 'colostomia' }
]);

const diureseOpts = ref([
    { name: 'Espontânea', code: 'espontanea' },
    { name: 'Induzida', code: 'induzida' }
]);

// Usado para SVD, Cistostomia e Preservativo
const volumeDiureseOpts = ref([
    { name: 'Normal', code: 'normal' },
    { name: 'Anúria', code: 'anuria' },
    { name: 'Oligúria', code: 'oliguria' },
    { name: 'Poliúria', code: 'poliuria' }
]);

const viasSondaOpts = ref([
    { name: '2 Vias', code: '2' },
    { name: '3 Vias', code: '3' }
]);

</script>

<template>
    <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600 mb-8 mt-8">
        <div class="flex flex-col gap-4 w-full">
            
            <h4 class="font-semibold text-xl">Eliminações</h4>

            <h5 class="font-medium text-gray-700 border-b pb-2">Função Intestinal</h5>
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="funcao_intestinal">Função Intestinal</label>
                    <Select 
                        id="funcao_intestinal" 
                        v-model="eliminacoesStore.eliminacoes.funcao_intestinal" 
                        :options="funcaoIntestinalOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="via_evacuacao_fisiologica">Via Fisiológica?</label>
                    <InputSwitch 
                        id="via_evacuacao_fisiologica" 
                        v-model="eliminacoesStore.eliminacoes.via_evacuacao_fisiologica" 
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="via_evacuacao_estomia">Usa Estomia?</label>
                    <InputSwitch 
                        id="via_evacuacao_estomia" 
                        v-model="eliminacoesStore.eliminacoes.via_evacuacao_estomia" 
                    />
                </div>
            </div>

            <div v-if="eliminacoesStore.eliminacoes.via_evacuacao_estomia" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 mt-2 dark:bg-gray-400/10">
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="tipo_estomia">Tipo de Estomia</label>
                        <Select 
                            id="tipo_estomia" 
                            v-model="eliminacoesStore.eliminacoes.tipo_estomia" 
                            :options="tipoEstomiaOpts" 
                            optionLabel="name" 
                            optionValue="code"
                            placeholder="Selecione" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="placa_estomia">Placa de Estomia</label>
                        <InputText 
                            id="placa_estomia" 
                            v-model="eliminacoesStore.eliminacoes.placa_estomia" 
                            placeholder="Modelo/Marca" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="bolsa_estomia">Bolsa de Estomia</label>
                        <InputText 
                            id="bolsa_estomia" 
                            v-model="eliminacoesStore.eliminacoes.bolsa_estomia" 
                            placeholder="Modelo/Marca" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>

            <h5 class="font-medium text-gray-700 border-b pb-2 mt-4">Diurese e Dispositivos Urinários</h5>
            
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/4">
                    <label for="diurese">Diurese</label>
                    <Select 
                        id="diurese" 
                        v-model="eliminacoesStore.eliminacoes.diurese" 
                        :options="diureseOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="usa_fralda">Usa Fralda?</label>
                    <InputSwitch 
                        id="usa_fralda" 
                        v-model="eliminacoesStore.eliminacoes.usa_fralda" 
                    />
                </div>
                <div v-if="eliminacoesStore.eliminacoes.usa_fralda" class="flex flex-col gap-2 w-full md:w-1/4">
                    <label for="trocas_fralda_dia">Nº trocas por dia</label>
                    <InputText 
                        id="trocas_fralda_dia" 
                        v-model="eliminacoesStore.eliminacoes.trocas_fralda_dia" 
                        placeholder="Ex: 4" 
                        class="w-full"
                    />
                </div>
                
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="sva">SVA (Alívio)?</label>
                    <InputSwitch 
                        id="sva" 
                        v-model="eliminacoesStore.eliminacoes.sva" 
                    />
                </div>
            </div>

            <div class="border-t my-2 pt-2">
                <div class="flex flex-col md:flex-row gap-4 items-center">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="cateterismo_intermitente">Cateterismo Intermitente?</label>
                        <InputSwitch 
                            id="cateterismo_intermitente" 
                            v-model="eliminacoesStore.eliminacoes.cateterismo_intermitente" 
                        />
                    </div>
                    <div v-if="eliminacoesStore.eliminacoes.cateterismo_intermitente" class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="vezes_cateterismo_dia">Nº de vezes ao dia</label>
                        <InputText 
                            id="vezes_cateterismo_dia" 
                            v-model="eliminacoesStore.eliminacoes.vezes_cateterismo_dia" 
                            placeholder="Quantidade" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>

            <div class="border-t my-2 pt-2">
                <div class="flex flex-col gap-2 w-full mb-2">
                    <label for="svd">Sonda Vesical de Demora (SVD)?</label>
                    <InputSwitch id="svd" v-model="eliminacoesStore.eliminacoes.svd" />
                </div>

                <div v-if="eliminacoesStore.eliminacoes.svd" class="grid grid-cols-1 md:grid-cols-4 gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                    <div class="flex flex-col gap-2">
                        <label for="num_svd">Nº da Sonda (Fr)</label>
                        <InputText id="num_svd" v-model="eliminacoesStore.eliminacoes.num_svd" placeholder="Ex: 16" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="vias_svd">Nº de Vias</label>
                        <Select id="vias_svd" v-model="eliminacoesStore.eliminacoes.vias_svd" :options="viasSondaOpts" optionLabel="name" optionValue="code" placeholder="Selecione" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="volume_diurese_svd">Volume Diurese</label>
                        <Select id="volume_diurese_svd" v-model="eliminacoesStore.eliminacoes.volume_diurese_svd" :options="volumeDiureseOpts" optionLabel="name" optionValue="code" placeholder="Selecione" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="data_troca_svd">Última Troca</label>
                        <Calendar id="data_troca_svd" v-model="eliminacoesStore.eliminacoes.data_troca_svd" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full" />
                    </div>
                    <div class="flex items-center gap-2 md:col-span-4">
                        <label for="irrigacao_svd">Realiza Irrigação?</label>
                        <InputSwitch id="irrigacao_svd" v-model="eliminacoesStore.eliminacoes.irrigacao_svd" />
                    </div>
                </div>
            </div>

            <div class="border-t my-2 pt-2">
                <div class="flex flex-col gap-2 w-full mb-2">
                    <label for="cistostomia">Cistostomia?</label>
                    <InputSwitch id="cistostomia" v-model="eliminacoesStore.eliminacoes.cistostomia" />
                </div>

                <div v-if="eliminacoesStore.eliminacoes.cistostomia" class="grid grid-cols-1 md:grid-cols-4 gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                    <div class="flex flex-col gap-2">
                        <label for="num_cistostomia">Nº da Sonda</label>
                        <InputText id="num_cistostomia" v-model="eliminacoesStore.eliminacoes.num_cistostomia" placeholder="Calibre" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="vias_cistostomia">Nº de Vias</label>
                        <Select id="vias_cistostomia" v-model="eliminacoesStore.eliminacoes.vias_cistostomia" :options="viasSondaOpts" optionLabel="name" optionValue="code" placeholder="Selecione" class="w-full" />
                    </div>
                    <div class="flex flex-col gap-2">
                        <label for="volume_diurese_cistostomia">Volume Diurese</label>
                        <Select id="volume_diurese_cistostomia" v-model="eliminacoesStore.eliminacoes.volume_diurese_cistostomia" :options="volumeDiureseOpts" optionLabel="name" optionValue="code" placeholder="Selecione" class="w-full" />
                    </div>
                    
                    <div class="md:col-span-4 grid grid-cols-1 md:grid-cols-3 gap-4 border-t pt-2 mt-2">
                        <div class="flex flex-col gap-2">
                            <label for="irrigacao_cistostomia">Realiza Irrigação?</label>
                            <InputSwitch id="irrigacao_cistostomia" v-model="eliminacoesStore.eliminacoes.irrigacao_cistostomia" />
                        </div>
                        <div v-if="eliminacoesStore.eliminacoes.irrigacao_cistostomia" class="flex flex-col gap-2">
                            <label for="qtd_irrigacao_cistostomia">Quantidade Irrigação</label>
                            <InputText id="qtd_irrigacao_cistostomia" v-model="eliminacoesStore.eliminacoes.qtd_irrigacao_cistostomia" placeholder="ml" class="w-full" />
                        </div>
                    </div>

                     <div class="md:col-span-4 grid grid-cols-1 md:grid-cols-3 gap-4 border-t pt-2">
                        <div class="flex flex-col gap-2">
                            <label for="medicacao_irrigacao_cistostomia">Medicação na irrigação?</label>
                            <InputSwitch id="medicacao_irrigacao_cistostomia" v-model="eliminacoesStore.eliminacoes.medicacao_irrigacao_cistostomia" />
                        </div>
                        <div v-if="eliminacoesStore.eliminacoes.medicacao_irrigacao_cistostomia" class="flex flex-col gap-2 md:col-span-2">
                            <label for="nome_medicacao_cistostomia">Qual medicação?</label>
                            <InputText id="nome_medicacao_cistostomia" v-model="eliminacoesStore.eliminacoes.nome_medicacao_cistostomia" placeholder="Nome do medicamento" class="w-full" />
                        </div>
                    </div>
                </div>
            </div>

            <div class="border-t my-2 pt-2">
                <div class="flex flex-col md:flex-row gap-4 items-center">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="preservativo">Preservativo Urinário?</label>
                        <InputSwitch id="preservativo" v-model="eliminacoesStore.eliminacoes.preservativo" />
                    </div>
                    <div v-if="eliminacoesStore.eliminacoes.preservativo" class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="volume_diurese_preservativo">Volume Diurese</label>
                        <Select 
                            id="volume_diurese_preservativo" 
                            v-model="eliminacoesStore.eliminacoes.volume_diurese_preservativo" 
                            :options="volumeDiureseOpts" 
                            optionLabel="name" 
                            optionValue="code"
                            placeholder="Selecione" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>

        </div>
    </div>
</template> 