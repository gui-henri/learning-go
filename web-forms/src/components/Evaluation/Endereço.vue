<script setup>
import { useEnderecoStore } from '@/store/evaluation/endereco';
import { useStepAccordion } from "@/composable/useStepAccordion";
import { computed, ref, watch } from 'vue';
import { AvaliationService } from '@/service/AvaliationService';
import { useAvaliationForm } from "@/store/evaluation/form";

const avaliationFormStore = useAvaliationForm();
const enderecoStore = useEnderecoStore();
const isLoadingCep = ref(false);

const emit = defineEmits(['next-step']);

const props = defineProps({
  isActive: {
    type: Boolean,
    default: false
  }
})

const isFilled = computed(() => {
    return !!enderecoStore.endereco.cep && 
           enderecoStore.endereco.cep.length > 3;
});

const buscarCep = async (cep) => {
    isLoadingCep.value = true;
    try {
        const cepLimpo = cep.replace(/\D/g, '');
        
        const response = await fetch(`https://viacep.com.br/ws/${cepLimpo}/json/`);
        const data = await response.json();

        if (!data.erro) {
            enderecoStore.endereco.cidade = data.localidade; 
            enderecoStore.endereco.logradouro = data.logradouro; 
            enderecoStore.endereco.bairro = data.bairro;
            enderecoStore.endereco.estado = data.estado; 
            
            if(data.complemento) {
                enderecoStore.endereco.complemento = data.complemento;
            }
            document.getElementById('numero-logradoro')?.focus();
        } else {
            alert('CEP não encontrado. Verifique os números.');
        }
    } catch (error) {
        console.error("Erro ao buscar CEP:", error);
    } finally {
        isLoadingCep.value = false;
    }
};

watch(() => enderecoStore.endereco.cep, (novoCep) => {
    const apenasNumeros = novoCep?.replace(/\D/g, '') || '';
    if (apenasNumeros.length === 8) {
        buscarCep(apenasNumeros);
    }
});

const { internalIndex, nextStep } = useStepAccordion(props, emit);

const handleSave = async () => {
    const payload = {
        endereco: enderecoStore.endereco,
    };
    try {
        if (avaliationFormStore.avaliationId === null) {
            await AvaliationService.submitFormData(payload)
        } else {
            await AvaliationService.appendToAvaliation(
                avaliationFormStore.avaliationId, 
                payload
            )
        }

        internalIndex.value = null;
        setTimeout(() => {
            const self = document.getElementById("endereco");
            if (self) {
                self.scrollIntoView({ 
                    behavior: 'instant', 
                    block: 'start',
                });
            }
        }, 0);
        nextStep()
    } catch (error) {
        console.error(error);
    }
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

<Accordion v-model:activeIndex="internalIndex" id="endereco" class=" scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
    <AccordionTab>
    <template #header>
        <div class="flex items-center gap-3 w-full">
            <i class="pi text-xl" 
                :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Endereço <i class="pi pi-map text-3xl"></i></h4>
                <span class="text-xs text-gray-500 font-normal -mt-4">
                    {{ isFilled ? '' : 'Toque para preencher' }}
                </span>
            </div>
        </div>
    </template>
    <div class="flex flex-col gap-4 w-full">
        <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="cep">CEP</label>
                <div class="p-input-icon-right w-full">
                    <i v-show="isLoadingCep" class="pi pi-spin pi-spinner" />
                    <InputMask 
                        id="cep" 
                        v-model="enderecoStore.endereco.cep" 
                        placeholder="55050-550"
                        mask="99999-999" 
                        class="w-full"
                        :disabled="isLoadingCep"
                    />
                </div>
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="estado">Estado</label>
                <AutoComplete
                    id="estado" 
                    v-model="enderecoStore.endereco.estado" 
                    :suggestions="estadoFiltrados" 
                    @complete="searchEstado" 
                    dropdown
                    placeholder="Pernambuco" 
                    class="w-full"
                    maxlength="20"
                    v-keyfilter="/[a-zA-Z\sáàãâéèêíìîóòõôúùûçÇÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/"
                    />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="cidade">Cidade</label>
                <InputText 
                    id="cidade" 
                    v-model="enderecoStore.endereco.cidade" 
                    placeholder="Recife" 
                    class="w-full"
                    maxlength="50"
                    v-keyfilter="/[a-zA-Z\sáàãâéèêíìîóòõôúùûçÇÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/"
                    
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="bairro">Bairro</label>
                <InputText 
                    id="bairro" 
                    v-model="enderecoStore.endereco.bairro" 
                    placeholder="Boa vista" 
                    class="w-full"
                    maxlength="50"
                    v-keyfilter="/[a-zA-Z\sáàãâéèêíìîóòõôúùûçÇÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/" 
                />
            </div>
        </div>

        <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="logradouro">Logradouro</label>
                <InputText 
                    id="logradouro" 
                    v-model="enderecoStore.endereco.logradouro" 
                    placeholder="Agamenon magalhães" 
                    class="w-full"
                    maxlength="50"
                    v-keyfilter="/[a-zA-Z\sáàãâéèêíìîóòõôúùûçÇÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="numero-logradoro">Número</label>
                <InputText 
                    id="numero-logradoro" 
                    v-model="enderecoStore.endereco.numero" 
                    placeholder="23-A" 
                    class="w-full"
                    maxlength="10"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="complemento">Complemento</label>
                <InputText 
                    id="complemento" 
                    v-model="enderecoStore.endereco.complemento" 
                    placeholder="Prédio azul" 
                    class="w-full"
                    maxlength="50"
                    v-keyfilter="/[a-zA-Z\sáàãâéèêíìîóòõôúùûçÇÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/4">
                <label for="ponto_referencia">Ponto de referência</label>
                <InputText 
                    id="ponto_referencia" 
                    v-model="enderecoStore.endereco.ponto_referencia" 
                    placeholder="perto do posto de saúde" 
                    class="w-full"
                    maxlength="50"
                    v-keyfilter="/[a-zA-Z\sáàãâéèêíìîóòõôúùûçÇÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/"
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