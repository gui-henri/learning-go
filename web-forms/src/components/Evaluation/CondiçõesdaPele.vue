<script setup>
import { ref, computed } from 'vue';
import { useCondicoesPeleStore } from '@/store/evaluation/condicoesPele';

const emit = defineEmits(['next-step']);

const condicoesPeleStore = useCondicoesPeleStore();

const condicaoCutaneaOpts = ref([
    { name: 'Íntegra', code: 'integra' },
    { name: 'Com Lesão', code: 'lesao' }
]);

const activeIndex = ref(null);

const isFilled = computed(() => {
    return !!condicoesPeleStore.condicoesPele.condicao_cutanea_mucosa;
});

const addCurativo = () => {
    condicoesPeleStore.adicionarCurativo();
};

const removeCurativo = (index) => {
    condicoesPeleStore.removerCurativo(index);
};

const handleSave = () => {
    activeIndex.value = null;
    
    setTimeout(() => {
        const self = document.getElementById("condicoesPele");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    emit('next-step');
};
</script>

<template>
    <Accordion v-model:activeIndex="activeIndex" id="condicoesPele" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Condições da Pele</h4>
                <span class="text-xs text-gray-500 font-normal -mt-4">
                    {{ isFilled ? '' : 'Toque para preencher' }}
                </span>
            </div>
        </div>
    </template>
        <div class="flex flex-col gap-4 w-full">
        
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="condicao_cutanea_mucosa">Como é a Cutânea-mucosa</label>
                    <Select 
                        id="condicao_cutanea_mucosa" 
                        v-model="condicoesPeleStore.condicoesPele.condicao_cutanea_mucosa" 
                        :options="condicaoCutaneaOpts" 
                        optionLabel="name" 
                        optionValue="name"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
            </div>
                <div v-for="(curativo, index) in condicoesPeleStore.condicoesPele.curativos" :key="index" class="p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10 flex flex-col gap-4 relative">
                    
                    <div class="flex justify-between items-center">
                        <h6 class="font-bold text-gray-600 m-0">Lesão {{ index + 1 }}</h6>
                        <Button v-if="condicoesPeleStore.condicoesPele.curativos.length > 0" 
                                icon="pi pi-trash" 
                                class="p-button-rounded p-button-danger p-button-text" 
                                @click="removeCurativo(index)" 
                                v-tooltip="'Remover curativo'" />
                    </div>

                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="realiza_curativo">Realiza curativo?</label>
                    <InputSwitch 
                        id="realiza_curativo" 
                        v-model="condicoesPeleStore.condicoesPele.realiza_curativo" 
                    />
                </div>

                    <div class="flex flex-col md:flex-row gap-4">
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label :for="'local_curativo_' + index">Local do curativo</label>
                            <InputText 
                                :id="'local_curativo_' + index" 
                                v-model="curativo.local_curativo" 
                                placeholder="Local anatômico" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label :for="'tipo_cobertura_' + index">Tipo de cobertura</label>
                            <InputText 
                                :id="'tipo_cobertura_' + index" 
                                v-model="curativo.tipo_cobertura" 
                                placeholder="Tipo de cobertura" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label :for="'tamanho_curativo_' + index">Tamanho do curativo</label>
                            <InputText 
                                :id="'tamanho_curativo_' + index" 
                                v-model="curativo.tamanho_curativo" 
                                placeholder="Dimensões" 
                                class="w-full"
                            />
                        </div>
                    </div>

                    <div class="flex flex-col md:flex-row gap-4">
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label :for="'qtd_curativo_' + index">Quantidade</label>
                            <InputText 
                                :id="'qtd_curativo_' + index" 
                                v-model="curativo.qtd_curativo" 
                                placeholder="Qtd necessária" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label :for="'frequencia_troca_' + index">Frequência de troca</label>
                            <InputText 
                                :id="'frequencia_troca_' + index" 
                                v-model="curativo.frequencia_troca" 
                                placeholder="Ex: 1x ao dia" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label :for="'materiais_curativo_' + index">Materiais descartáveis</label>
                            <InputText 
                                :id="'materiais_curativo_' + index" 
                                v-model="curativo.materiais_curativo" 
                                placeholder="Lista de materiais" 
                                class="w-full"
                            />
                        </div>
                    </div>
                </div>

                <div class="flex justify-end w-full">
                    <Button label="Adicionar curativo" icon="pi pi-plus" class="p-button-outlined p-button-secondary" @click="addCurativo" />
                </div>

        </div>
        <Button class="mt-3" v-on:click="handleSave">
            <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
            Próximo
        </Button>
    </AccordionTab>
</Accordion>
</template>