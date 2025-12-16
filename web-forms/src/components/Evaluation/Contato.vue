<script setup>
import { computed } from 'vue';
import { useContatoStore } from '@/store/evaluation/contato';
import { useStepAccordion } from "@/composable/useStepAccordion";
import { InputMask } from 'primevue';
import { useAvaliationForm } from "@/store/evaluation/form";
import { AvaliationService } from '@/service/AvaliationService';

const avaliationFormStore = useAvaliationForm();
const contatoStore = useContatoStore();

const props = defineProps({
  formFields: {
    type: Object,
    required: false,
    default: null
  },
  isActive: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['next-step']);

const isFilled = computed(() => {
    return !!contatoStore.contato.telefone_paciente && 
           contatoStore.contato.telefone_paciente.length > 3;
});

const { internalIndex, nextStep } = useStepAccordion(props, emit);

const handleSave = async () => {

    const payload = {
        contato: contatoStore.contato,
    };
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
        const self = document.getElementById("contato");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    nextStep('next-step');
};

const addResponsavel = () => {
    contatoStore.adicionarResponsavel();
};

const removeResponsavel = (index) => {
    contatoStore.removerResponsavel(index);
};

</script>

<template>
    <Accordion v-model:activeIndex="internalIndex" id="contato" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
                </i>
                <div class="flex flex-col text-left">
                    <h4 class="font-semibold text-xl p-0 m-0">Contato <i class="pi pi-phone text-3xl"></i>   </h4>
                    <span class="text-xs text-gray-500 font-normal -mt-4">
                        {{ isFilled ? '' : 'Toque para preencher' }}
                    </span>
                </div>
            </div>
        </template>
        <div class="flex flex-col gap-4 w-full">
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="telefone_residencial">Telefone residencial</label>
                    <InputMask
                        id="telefone_residencial" 
                        v-model="contatoStore.contato.telefone_residencial" 
                        w
                        class="w-full"
                        mask="(99) 9 9999-9999"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="telefone_paciente">Telefone paciente</label>
                    <InputMask
                        id="telefone_paciente" 
                        v-model="contatoStore.contato.telefone_paciente" 
                        placeholder="(81) 9 9999-9999" 
                        class="w-full"
                        mask="(99) 9 9999-9999"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/3">
                    <label for="email_paciente">E-mail paciente</label>
                    <InputText
                        id="email_paciente" 
                        v-model="contatoStore.contato.email_paciente" 
                        placeholder="exemplo@dominio.com" 
                        class="w-full"
                        type="email"
                    />
                </div>
            </div>
            <div v-for="(responsavel, index) in contatoStore.contato.responsaveis" :key="index" class="flex flex-col gap-4 p-4 border border-gray-200 rounded-lg relative bg-gray-50 dark:bg-gray-800">
                
                <div class="flex justify-between items-center">
                    <h5 class="font-bold text-gray-600">Responsável {{ index + 1 }}</h5>
                    <Button v-if="contatoStore.contato.responsaveis.length > 1" 
                            icon="pi pi-trash" 
                            class="p-button-rounded p-button-danger p-button-text" 
                            @click="removeResponsavel(index)" 
                            v-tooltip="'Remover este responsável'" />
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label :for="'nome_responsavel_' + index">Nome do responsável</label>
                        <InputText 
                            :id="'nome_responsavel_' + index" 
                            v-model="responsavel.nome" 
                            placeholder="João albério" 
                            class="w-full"
                        />
                    </div>
                    
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label :for="'parentesco_responsavel_' + index">Parentesco</label>
                        <InputText 
                            :id="'parentesco_responsavel_' + index" 
                            v-model="responsavel.parentesco" 
                            placeholder="Pai, mãe..." 
                            class="w-full"
                        />
                    </div>
                
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label :for="'email_responsavel_' + index">E-mail do responsável</label>
                        <InputText 
                            :id="'email_responsavel_' + index" 
                            v-model="responsavel.email" 
                            placeholder="resp@email.com" 
                            class="w-full"
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/3">
                        <label :for="'telefone_responsavel_' + index">Telefone do responsável</label>
                        <InputMask
                            :id="'telefone_responsavel_' + index" 
                            v-model="responsavel.telefone" 
                            placeholder="(81) 9 9999-9999" 
                            class="w-full"
                            mask="(99) 9 9999-9999"
                        />
                    </div>
                    <div class="flex flex-col gap-2 md:w-1/3">
                        <label :for="'forma_contato_' + index">Forma de contato</label>
                        <Select 
                            :id="'forma_contato_' + index" 
                            v-model="responsavel.forma_contato" 
                            :options="formFields.responsavel" 
                            optionLabel="label"
                            optionValue="label"
                            placeholder="Selecione" 
                            class="w-full"
                        ></Select>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/3 items-start justify-center">
                        <label :for="'numero_prioritario_' + index" class="mb-2">Contato prioritário?</label>
                        <InputSwitch 
                            :id="'numero_prioritario_' + index" 
                            v-model="responsavel.numero_prioritario" 
                        />
                    </div>
                </div>
            </div>

            <div class="flex justify-end w-full">
                <Button label="Adicionar responsável" icon="pi pi-plus" class="p-button-outlined p-button-secondary" @click="addResponsavel" />
            </div>
        </div>
    <Button class="mt-3" v-on:click="handleSave">
        <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
        Próximo
    </Button>
    </AccordionTab>
</Accordion>
</template>