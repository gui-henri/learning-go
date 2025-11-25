<script setup>
import { useObservacoesStore } from '@/store/evaluation/observacao';

const observacoesStore = useObservacoesStore();
</script>

<template>
    <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600 mb-8 mt-8">
        <div class="flex flex-col gap-4 w-full">
            
            <h4 class="font-semibold text-xl">Observações</h4>
            
            <div class="flex flex-col gap-2 w-full">
                <label for="observacoes_gerais">Anotações adicionais</label>
                <Textarea 
                    id="observacoes_gerais" 
                    v-model="observacoesStore.observacoes.texto" 
                    rows="10" 
                    placeholder="Descreva aqui qualquer observação adicional relevante sobre o paciente, condições do domicílio ou detalhes que não se encaixaram nos campos anteriores..." 
                    class="w-full" 
                    autoResize
                />
            </div>

        </div>
    </div>
</template>