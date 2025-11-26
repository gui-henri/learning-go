<script setup>
import { ref } from 'vue';
import { useCuidadoresStore } from '@/store/evaluation/cuidadores';

const cuidadorStore = useCuidadoresStore();


const turnos_cuidador = ref([
    { name: 'Diurno', code: 'diurno' },
    { name: 'Noturno', code: 'noturno' },
    { name: '24h', code: '24h' }
]);

</script>

<template>
    <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <div class="flex flex-col gap-4 w-full">
            
            <h4 class="font-semibold text-xl">Cuidadores</h4>
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/2">
                    <label for="medico_solicitante">Médico solicitante</label>
                    <InputText 
                        id="medico_solicitante" 
                        v-model="cuidadorStore.cuidadores.medico_solicitante" 
                        placeholder="Nome do médico" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/8">
                    <label for="contato_medico">Contato do médico</label>
                    <InputText 
                        id="contato_medico" 
                        v-model="cuidadorStore.cuidadores.contato_medico" 
                        placeholder="(XX) 9 9999-9999" 
                        class="w-full"
                    />
                </div>
            </div>
             <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="possui_cuidador">Possui cuidador?</label>
                    <InputSwitch 
                        id="possui_cuidador" 
                        v-model="cuidadorStore.cuidadores.possui_cuidador" 
                        class="w-full"
                    ></InputSwitch>
                </div>

            <template v-if="cuidadorStore.cuidadores.possui_cuidador">
                
                <hr class="border-gray-200 my-2"/>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_cuidador">Nome do cuidador</label>
                        <InputText 
                            id="nome_cuidador" 
                            v-model="cuidadorStore.cuidadores.nome_cuidador" 
                            placeholder="Nome completo" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="contato_cuidador">Contato do cuidador</label>
                        <InputText 
                            id="contato_cuidador" 
                            v-model="cuidadorStore.cuidadores.contato_cuidador" 
                            placeholder="(XX) 9 9999-9999" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="turno_cuidador">Turno</label>
                        <Select 
                            id="turno_cuidador" 
                            v-model="cuidadorStore.cuidadores.turno_cuidador" 
                            :options="turnos_cuidador" 
                            optionLabel="name" 
                            placeholder="Selecione o turno" 
                            class="w-full"
                        ></Select>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4 items-center">
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="precisa_treinamento">Precisa de treinamento?</label>
                        <InputSwitch 
                            id="precisa_treinamento" 
                            v-model="cuidadorStore.cuidadores.precisa_treinamento" 
                            class="w-full"
                        ></InputSwitch>
                    </div>
                    
                    <div v-if="cuidadorStore.cuidadores.precisa_treinamento" class="flex flex-col gap-2 w-full flex-1 transition-all duration-300">
                        <label for="obs_treinamento">Observações do treinamento</label>
                        <InputText 
                            id="obs_treinamento" 
                            v-model="cuidadorStore.cuidadores.obs_treinamento" 
                            placeholder="Descreva a necessidade..." 
                            class="w-full"
                        />
                    </div>
                </div>

            </template>
            </div>
    </div>
</template>