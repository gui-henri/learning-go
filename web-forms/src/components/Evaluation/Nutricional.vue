<script setup>
import { ref } from 'vue';
import { useNutricionalStore } from '@/store/evaluation/nutricional';

const nutricionalStore = useNutricionalStore();

const viaEnteralOpts = ref([
    { name: 'Nasogástrica', code: 'nasogastrica' },
    { name: 'Orogástrica', code: 'orogastrica' },
    { name: 'Gastrostomia', code: 'gastrostomia' },
    { name: 'Jejunostomia', code: 'jejunostomia' },
    { name: 'Oroenteral', code: 'oroenteral' },
    { name: 'Nasoenteral', code: 'nasoenteral' }
]);

const viaParenteralOpts = ref([
    { name: 'Central', code: 'central' },
    { name: 'Periférica', code: 'periferica' }
]);

const adaptadorSondaOpts = ref([
    { name: 'Torneirinha 3 vias', code: 'torneirinha_3_vias' },
    { name: 'Polifix 2 vias', code: 'polifix_2_vias' },
    { name: 'Outros', code: 'outros' }
]);

const tipoDietaOpts = ref([
    { name: 'Artesanal', code: 'artesanal' },
    { name: 'Industrializada', code: 'industrializada' },
    { name: 'Mista', code: 'mista' }
]);

const formaAdministracaoOpts = ref([
    { name: 'Gravitacional', code: 'gravitacional' },
    { name: 'Bomba de infusão', code: 'bomba_infusao' },
    { name: 'Bomba de seringa', code: 'bomba_seringa' }
]);

</script>

<template>

   <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600 mb-8 mt-8">
        <div class="flex flex-col gap-4 w-full">
            
            <h4 class="font-semibold text-xl">Nutricional</h4>

            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="alimentacao_oral">Alimentação Oral?</label>
                    <InputSwitch 
                        id="alimentacao_oral" 
                        v-model="nutricionalStore.nutricional.alimentacao_oral" 
                    />
                </div>
                
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="alimentacao_enteral">Alimentação Enteral?</label>
                    <InputSwitch 
                        id="alimentacao_enteral" 
                        v-model="nutricionalStore.nutricional.alimentacao_enteral" 
                    />
                </div>

                <div v-if="nutricionalStore.nutricional.alimentacao_enteral" class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="via_enteral">Via de alimentação enteral <span class="text-red-500">*</span></label>
                    <Select 
                        id="via_enteral" 
                        v-model="nutricionalStore.nutricional.via_enteral" 
                        :options="viaEnteralOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
            </div>

            <div class="flex flex-col md:flex-row gap-4 border-t pt-4">
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="alimentacao_parenteral">Alimentação Parenteral?</label>
                    <InputSwitch 
                        id="alimentacao_parenteral" 
                        v-model="nutricionalStore.nutricional.alimentacao_parenteral" 
                    />
                </div>

                <div v-if="nutricionalStore.nutricional.alimentacao_parenteral" class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="via_parenteral">Via de alimentação parenteral <span class="text-red-500">*</span></label>
                    <Select 
                        id="via_parenteral" 
                        v-model="nutricionalStore.nutricional.via_parenteral" 
                        :options="viaParenteralOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
            </div>

            <div class="flex flex-col md:flex-row gap-4 border-t pt-4">
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="sonda">Usa Sonda?</label>
                    <InputSwitch 
                        id="sonda" 
                        v-model="nutricionalStore.nutricional.sonda" 
                    />
                </div>

                <template v-if="nutricionalStore.nutricional.sonda">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="data_ultima_troca">Data da última troca</label>
                        <Calendar 
                            id="data_ultima_troca" 
                            v-model="nutricionalStore.nutricional.data_ultima_troca" 
                            dateFormat="dd/mm/yy" 
                            placeholder="dd/mm/aaaa" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="adaptador_sonda">Adaptador da sonda</label>
                        <Select 
                            id="adaptador_sonda" 
                            v-model="nutricionalStore.nutricional.adaptador_sonda" 
                            :options="adaptadorSondaOpts" 
                            optionLabel="name" 
                            optionValue="code"
                            placeholder="Selecione" 
                            class="w-full"
                        />
                    </div>
                </template>

                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="botton">Usa Botton?</label>
                    <InputSwitch 
                        id="botton" 
                        v-model="nutricionalStore.nutricional.botton" 
                    />
                </div>
            </div>

            <div class="flex flex-col gap-2 w-full border-t pt-4">
                <label for="restricao_alimentar">Restrição alimentar <span class="text-red-500">*</span></label>
                <InputText 
                    id="restricao_alimentar" 
                    v-model="nutricionalStore.nutricional.restricao_alimentar" 
                    placeholder="Descrição da restrição alimentar" 
                    class="w-full"
                />
            </div>

            <div class="flex flex-col gap-4 border-t pt-4">
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="tipo_dieta">Tipo de dieta <span class="text-red-500">*</span></label>
                        <Select 
                            id="tipo_dieta" 
                            v-model="nutricionalStore.nutricional.tipo_dieta" 
                            :options="tipoDietaOpts" 
                            optionLabel="name" 
                            optionValue="code"
                            placeholder="Selecione" 
                            class="w-full"
                        />
                    </div>

                    <template v-if="nutricionalStore.nutricional.tipo_dieta === 'industrializada' || nutricionalStore.nutricional.tipo_dieta === 'mista'">
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label for="gavando">Gavando (ml/h) <span class="text-red-500">*</span></label>
                            <InputText 
                                id="gavando" 
                                v-model="nutricionalStore.nutricional.gavando" 
                                type="number"
                                placeholder="ml/h" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label for="volume_diario">Volume diário <span class="text-red-500">*</span></label>
                            <InputText 
                                id="volume_diario" 
                                v-model="nutricionalStore.nutricional.volume_diario" 
                                type="number"
                                placeholder="Total em ml" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="forma_administracao">Forma de administração <span class="text-red-500">*</span></label>
                            <Select 
                                id="forma_administracao" 
                                v-model="nutricionalStore.nutricional.forma_administracao" 
                                :options="formaAdministracaoOpts" 
                                optionLabel="name" 
                                optionValue="code"
                                placeholder="Selecione" 
                                class="w-full"
                            />
                        </div>
                    </template>
                </div>
            </div>

            <div class="flex flex-col gap-4 border-t pt-4">
                <div class="flex flex-col md:flex-row gap-4 items-center">
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="suplemento">Suplemento?</label>
                        <InputSwitch 
                            id="suplemento" 
                            v-model="nutricionalStore.nutricional.suplemento" 
                        />
                    </div>
                    
                    <div v-if="nutricionalStore.nutricional.suplemento" class="flex flex-col gap-2 w-full flex-1">
                        <label for="qual_suplemento">Qual suplemento? <span class="text-red-500">*</span></label>
                        <InputText 
                            id="qual_suplemento" 
                            v-model="nutricionalStore.nutricional.qual_suplemento" 
                            placeholder="Nome do suplemento" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>

        </div>
    </div>
</template>