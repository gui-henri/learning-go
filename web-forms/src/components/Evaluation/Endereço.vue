<script setup>
import { ref, computed } from 'vue';
import { useEnderecoStore } from '@/store/evaluation/endereco';

const emit = defineEmits(['next-step']);
const enderecoStore = useEnderecoStore();
const activeIndex = ref(null);

const isFilled = computed(() => {
    return !!enderecoStore.endereco.cep && 
           enderecoStore.endereco.cep.length > 3;
});

const handleSave = () => {
    activeIndex.value = null;
    
    setTimeout(() => {
        const self = document.getElementById("endereco");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    emit('next-step');
};

const estadosOpts = ref([
    "Acre",
"Alagoas",
"Amapá",
"Amazonas",
"Bahia",
"Ceará",
"Distrito Federal",
"Espírito Santo",
"Goiás",
"Maranhão",
"Mato Grosso",
"Mato Grosso do Sul",
"Minas Gerais",
"Pará",
"Paraíba",
"Paraná",
"Pernambuco",
"Piauí",
"Rio de Janeiro",
"Rio Grande do Norte",
"Rio Grande do Sul",
"Rondônia",
"Roraima",
"Santa Catarina",
"São Paulo",
"Sergipe",
"Tocantins"
]);

const estadoFiltrados = ref([]);


const searchEstado = (event) => {
    if (!event.query.trim().length) {

        estadoFiltrados.value = [...estadosOpts.value];
    } else {
        estadoFiltrados.value = estadosOpts.value.filter((estado) => {
            return estado.toLowerCase().startsWith(event.query.toLowerCase());
        });
    }
};

</script>

<template>

<Accordion v-model:activeIndex="activeIndex" id="endereco" class=" scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
    <AccordionTab>
    <template #header>
        <div class="flex items-center gap-3 w-full">
            <i class="pi text-xl" 
                :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Endereço</h4>
                <span class="text-xs text-gray-500 font-normal -mt-4">
                    {{ isFilled ? '' : 'Toque para preencher' }}
                </span>
            </div>
        </div>
    </template>
    <div class="flex flex-col gap-4 w-full">
        <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col gap-2 w-full md:w-1/6">
                <label for="cep">CEP</label>
                <InputText 
                    id="cep" 
                    v-model="enderecoStore.endereco.cep" 
                    placeholder="55050-550" 
                    class="w-full"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/6">
                <label for="estado">Estado</label>
                <AutoComplete
                    id="estado" 
                    v-model="enderecoStore.endereco.estado" 
                    :suggestions="estadoFiltrados" 
                    @complete="searchEstado" 
                    dropdown
                    placeholder="Pernambuco" 
                    class="w-full"
                    />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/6">
                <label for="cidade">Cidade</label>
                <InputText 
                    id="cidade" 
                    v-model="enderecoStore.endereco.cidade" 
                    placeholder="Recife" 
                    class="w-full"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/6">
                <label for="bairro">Bairro</label>
                <InputText 
                    id="bairro" 
                    v-model="enderecoStore.endereco.bairro" 
                    placeholder="Boa vista" 
                    class="w-full"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/6">
                <label for="logradoro">Logradoro</label>
                <InputText 
                    id="logradoro" 
                    v-model="enderecoStore.endereco.logradoro" 
                    placeholder="Agamenon magalhães" 
                    class="w-full"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/10">
                <label for="numero-logradoro">Número</label>
                <InputText 
                    id="numero-logradoro" 
                    v-model="enderecoStore.endereco.numero" 
                    placeholder="23-A" 
                    class="w-full"
                />
            </div>
        </div>

        <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="complemento">Complemento</label>
                <InputText 
                    id="complemento" 
                    v-model="enderecoStore.endereco.complemento" 
                    placeholder="Prédio azul" 
                    class="w-full"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="ponto_referencia">Ponto de referência</label>
                <InputText 
                    id="ponto_referencia" 
                    v-model="enderecoStore.endereco.ponto_referencia" 
                    placeholder="perto do posto de saúde" 
                    class="w-full"
                />
            </div>
        </div>
    </div>
    <Button class="mt-3" v-on:click="handleSave">
        <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
        Próximo
    </Button>
    </AccordionTab>
</Accordion>
</template>