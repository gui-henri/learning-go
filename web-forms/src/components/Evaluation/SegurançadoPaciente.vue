<script setup>
import { ref, computed } from 'vue';
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

const emit = defineEmits(['next-step']);
const activeIndex = ref(null);

const addAlergia = () => {
    segurancaStore.adicionarAlergia();
};

const removeAlergia = (index) => {
    segurancaStore.removerAlergia(index);
};

const isFilled = computed(() => {
    const temAlergia = segurancaStore.seguranca.alergias.length > 0 && 
                       !!segurancaStore.seguranca.alergias[0].tipo_alergia;
    const temCuidado = !!segurancaStore.seguranca.cuidados_paliativos;
    
    return temAlergia || temCuidado;
});

const handleSave = () => {
    activeIndex.value = null;
    
    setTimeout(() => {
        const self = document.getElementById("segurança");
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
<Accordion v-model:activeIndex="activeIndex" id="segurança" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
    <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
                </i>
                <div class="flex flex-col text-left">
                    <h4 class="font-semibold text-xl p-0 m-0">Segurança do Paciente</h4>
                    <span class="text-xs text-gray-500 font-normal -mt-4">
                        {{ isFilled ? '' : 'Toque para preencher' }}
                    </span>
                </div>
            </div>
        </template>
    <div class="flex flex-col gap-4 w-full">        
        <div v-for="(item, index) in segurancaStore.seguranca.alergias" :key="index" class="p-4 border border-gray-200 rounded-lg bg-gray-50 dark:bg-gray-800 flex flex-col gap-4">
            
            <div class="flex justify-between items-center">
                <h5 class="font-bold text-gray-600">Alergia {{ index + 1 }}</h5>
                <Button v-if="segurancaStore.seguranca.alergias.length > 0" 
                        icon="pi pi-trash" 
                        class="p-button-rounded p-button-danger p-button-text" 
                        @click="removeAlergia(index)" />
            </div>

            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/4">
                    <label :for="'tipo_alergia_' + index">Tipo de alergia</label>
                    <Select 
                        :id="'tipo_alergia_' + index" 
                        v-model="item.tipo_alergia" 
                        :options="tipo_alergia" 
                        optionLabel="name" 
                        placeholder="Selecione" 
                        class="w-full"
                    ></Select>
                </div>

                <div class="flex flex-col gap-2 w-full flex-1">
                    <label :for="'quais_alergias_' + index">Descrição (Quais?)</label>
                    <InputText 
                        :id="'quais_alergias_' + index" 
                        v-model="item.quais_alergias" 
                        placeholder="Descreva as alergias..." 
                        class="w-full"
                    />
                </div>
            </div>

            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/2">
                    <label :for="'precaucao_' + index">Precaução</label>
                    <Select 
                        :id="'precaucao_' + index" 
                        v-model="item.precaucao" 
                        :options="tipo_precaucao" 
                        optionLabel="name" 
                        placeholder="Selecione" 
                        class="w-full"
                    ></Select>
                </div>
            </div>
            <div class="flex flex-col gap-2 w-full">
                <label for="cuidados_paliativos">Cuidados Paliativos</label>
                <InputText 
                    id="cuidados_paliativos" 
                    v-model="item.cuidados_paliativos" 
                    placeholder="Descreva os cuidados gerais com o paciente..." 
                    class="w-full"
                />
            </div>
        </div>

        <div class="flex justify-end w-full">
            <Button label="Adicionar alergia" icon="pi pi-plus" class="p-button-outlined p-button-secondary" @click="addAlergia" />
        </div>
        <Divider />
    </div>
    <Button class="mt-3" v-on:click="handleSave">
        <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
        Próximo
    </Button>
    </AccordionTab>
</Accordion>        
</template>