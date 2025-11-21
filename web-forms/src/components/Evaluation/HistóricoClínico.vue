<script setup>
import { ref } from 'vue';
import { useHistoricoStore } from '@/store/evaluation/historico';   
const historicoStore = useHistoricoStore();
</script>

<template>
    <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <div class="flex flex-col gap-4 w-full">

            <h4 class="font-semibold text-xl">Histórico e Planejamento</h4>
                
                <div class="flex flex-col gap-2 w-full">
                    <label for="historia_doenca_atual">História da doença atual</label>
                    <Textarea 
                        id="historia_doenca_atual" 
                        v-model="historicoStore.historico.historia_doenca_atual" 
                        rows="4" 
                        placeholder="Descreva a história clínica mais atual..." 
                        class="w-full" 
                        autoResize
                    />
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="anamnese_enfermagem">Anamnese de enfermagem</label>
                    <Textarea 
                        id="anamnese_enfermagem" 
                        v-model="historicoStore.historico.anamnese_enfermagem" 
                        rows="4" 
                        placeholder="Dados sobre o histórico de saúde..." 
                        class="w-full" 
                        autoResize
                    />
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="plano_terapeutico">Plano terapêutico</label>
                    <Textarea 
                        id="plano_terapeutico" 
                        v-model="historicoStore.historico.plano_terapeutico" 
                        rows="4" 
                        placeholder="Planejamento do cuidado..." 
                        class="w-full" 
                        autoResize
                    /> 
                </div>
            </div>
    </div>
</template>