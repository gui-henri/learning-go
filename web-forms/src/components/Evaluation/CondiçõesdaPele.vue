<script setup>
import { ref } from 'vue';
import { useCondicoesPeleStore } from '@/store/evaluation/condicoesPele';

const condicoesPeleStore = useCondicoesPeleStore();

const condicaoCutaneaOpts = ref([
    { name: 'Íntegra', code: 'integra' },
    { name: 'Com Lesão', code: 'lesao' }
]);

</script>

<template>
    <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600 mb-8 mt-8">
        <div class="flex flex-col gap-4 w-full">
            
            <h4 class="font-semibold text-xl">Condições da Pele</h4>

            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="condicao_cutanea_mucosa">Como é a Cutânea-mucosa</label>
                    <Select 
                        id="condicao_cutanea_mucosa" 
                        v-model="condicoesPeleStore.condicoesPele.condicao_cutanea_mucosa" 
                        :options="condicaoCutaneaOpts" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="realiza_curativo">Realiza curativo?</label>
                    <InputSwitch 
                        id="realiza_curativo" 
                        v-model="condicoesPeleStore.condicoesPele.realiza_curativo" 
                    />
                </div>
            </div>

            <div v-if="condicoesPeleStore.condicoesPele.realiza_curativo" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 mt-2 dark:bg-gray-400/10">
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="local_curativo">Local do curativo</label>
                        <InputText 
                            id="local_curativo" 
                            v-model="condicoesPeleStore.condicoesPele.local_curativo" 
                            placeholder="Local anatômico" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="tipo_cobertura">Tipo de cobertura</label>
                        <InputText 
                            id="tipo_cobertura" 
                            v-model="condicoesPeleStore.condicoesPele.tipo_cobertura" 
                            placeholder="Tipo de cobertura" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="tamanho_curativo">Tamanho do curativo</label>
                        <InputText 
                            id="tamanho_curativo" 
                            v-model="condicoesPeleStore.condicoesPele.tamanho_curativo" 
                            placeholder="Dimensões" 
                            class="w-full"
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="qtd_curativo">Quantidade</label>
                        <InputText 
                            id="qtd_curativo" 
                            v-model="condicoesPeleStore.condicoesPele.qtd_curativo" 
                            placeholder="Qtd necessária" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="frequencia_troca">Frequência de troca</label>
                        <InputText 
                            id="frequencia_troca" 
                            v-model="condicoesPeleStore.condicoesPele.frequencia_troca" 
                            placeholder="Ex: 1x ao dia" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label for="materiais_curativo">Materiais descartáveis</label>
                        <InputText 
                            id="materiais_curativo" 
                            v-model="condicoesPeleStore.condicoesPele.materiais_curativo" 
                            placeholder="Lista de materiais" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>

        </div>
    </div>
</template>