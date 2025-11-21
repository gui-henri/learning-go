<script setup>
import { ref } from 'vue';
import { useSegurancaStore } from '@/store/evaluation/seguranca';


const segurancaStore = useSegurancaStore();

const tipo_alergia = ref([
    { name: 'Medicamentosa', value: 'medicamentosa' },
    { name: 'Alimentar', value: 'alimentar' },
    { name: 'Respiratória', value: 'respiratoria' },
    { name: 'Outros', value: 'outros' }
]);

const tipo_precaucao = ref([
    { name: 'Padrão', value: 'padrao' },
    { name: 'Contato', value: 'contato' },
    { name: 'Gotículas', value: 'goticulas' },
    { name: 'Aerossóis', value: 'aerossois' }
]);
</script>

<template>
    
        
        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Segurança do paciente</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="alergia_relatada">Alergia relatada?</label>
                        <InputSwitch 
                            id="alergia_relatada" 
                            v-model="segurancaStore.seguranca.alergia_relatada" 
                            class="w-full"
                        ></InputSwitch>
                    </div>
                    
                    <template v-if="segurancaStore.seguranca.alergia_relatada">
                        <div class="flex flex-col gap-2 w-full md:w-1/6">
                            <label for="tipo_alergia">Tipo de alergia</label>
                            <Select 
                                id="tipo_alergia" 
                                v-model="segurancaStore.seguranca.tipo_alergia" 
                                :options="tipo_alergia" 
                                optionLabel="name" 
                                placeholder="Selecione" 
                                class="w-full"
                            ></Select>
                        </div>

                        <div class="flex flex-col gap-2 w-full flex-1">
                            <label for="quais_alergias">Quais alergias?</label>
                            <InputText 
                                id="quais_alergias" 
                                v-model="segurancaStore.seguranca.quais_alergias" 
                                placeholder="Descreva as alergias..." 
                                class="w-full"
                            />
                        </div>
                    </template>
                </div>

                <div v-if="segurancaStore.seguranca.alergia_relatada" class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="precaucao">Precaução</label>
                        <Select 
                            id="precaucao" 
                            v-model="segurancaStore.seguranca.precaucao" 
                            :options="tipo_precaucao" 
                            optionLabel="name" 
                            placeholder="Selecione" 
                            class="w-full"
                        ></Select>
                    </div>

                    <div class="flex flex-col gap-2 w-full flex-1">
                        <label for="cuidados_paliativos">Cuidados Paliativos</label>
                        <InputText 
                            id="cuidados_paliativos" 
                            v-model="segurancaStore.seguranca.cuidados_paliativos" 
                            placeholder="Descreva os cuidados..." 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>
        </div>


        
</template>