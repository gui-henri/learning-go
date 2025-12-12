<script setup>
import { ref, computed, watch } from 'vue';
import { useDadosGeraisStore } from '@/store/evaluation/dadosGerais';
import { useAvaliationForm } from "@/store/evaluation/form";
import { useStepAccordion } from "@/composable/useStepAccordion";
import { AvaliationService } from '@/service/AvaliationService';

const avaliationFormStore = useAvaliationForm();

const dadosGeraisStore = useDadosGeraisStore();
const emit = defineEmits(['next-step']);

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

const isFilled = computed(() => {
    return !!dadosGeraisStore.dadosGerais.nome_paciente && 
           dadosGeraisStore.dadosGerais.nome_paciente.length > 3;
});

const calculateAge = (dateOfBirth) => {
    if (!dateOfBirth) return null;

    const dob = new Date(dateOfBirth);
    if (isNaN(dob.getTime())) return null;

    const today = new Date();
    let age = today.getFullYear() - dob.getFullYear();
    const m = today.getMonth() - dob.getMonth();

    if (m < 0 || (m === 0 && today.getDate() < dob.getDate())) {
        age--;
    }

    return age;
};

watch(() => dadosGeraisStore.dadosGerais.data_nascimento, (newDate) => {
    const age = calculateAge(newDate);

    if (age !== null && age >= 0) {
        dadosGeraisStore.dadosGerais.idade = age;
    } else {

        dadosGeraisStore.dadosGerais.idade = null;
    }
}, { immediate: true });

const { internalIndex, nextStep } = useStepAccordion(props, emit);


const handleSave = async () => {

    const payload = {
        dadosGerais: dadosGeraisStore.dadosGerais,
    };
    if (avaliationFormStore.avaliationId === null) {
        await AvaliationService.submitFormData(payload)
    } else {
        await AvaliationService.appendToAvaliation(
            avaliationFormStore.avaliationId, 
            dadosGeraisStore.dadosGerais
        )
    }

    internalIndex.value = null;
    setTimeout(() => {
        const self = document.getElementById("dados-gerais");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    nextStep()
};
</script>
<template>

<Accordion v-model:activeIndex="internalIndex" id="dados-gerais" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
    <AccordionTab>
    <template #header>
        <div class="flex items-center gap-3 w-full">
            <i class="pi text-xl" 
                :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Dados Gerais <i class="pi pi-id-card text-3xl"></i></h4>
                <span class="text-xs text-gray-500 font-normal -mt-4">
                    {{ isFilled ? '' : 'Toque para preencher' }}
                </span>
            </div>
        </div>
    </template>
    <div class="flex flex-col gap-4 w-full">
        <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col gap-2 w-full md:w-4/10 ">
                <label for="nome_paciente">Nome completo</label>
                <InputText 
                    id="nome_paciente" 
                    v-model="dadosGeraisStore.dadosGerais.nome_paciente" 
                    type="text" 
                    placeholder="Glauberthy júnior"
                    maxlength="100"
                    v-keyfilter="/[a-zA-Z\sáàãâéèêíìîóòõôúùûçÇÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/10 lg:w-2/12">
                <label for="genero">Gênero</label>
                <Select 
                    id="genero" 
                    v-model="dadosGeraisStore.dadosGerais.genero" 
                    :options="props.formFields.genero" 
                    optionLabel="label"
                    optionValue="label"
                    placeholder="Selecione" 
                    class="w-full"
                ></Select>
            </div>
            <div class="flex flex-col gap-2 w-full md:w-2/10 lg:w-3/14">
                <label for="nascimento">Data de nascimento</label>

                <InputMask
                    id="nascimento"
                    v-model="dadosGeraisStore.dadosGerais.data_nascimento"
                    placeholder="dd/mm/aaaa"
                    class="w-full"
                    showIcon="true"
                    mask="99/99/9999"

                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-3/10 lg:w-2/12">
                <label for="cpf">CPF</label>
                <InputMask 
                    id="cpf" 
                    v-model="dadosGeraisStore.dadosGerais.cpf" 
                    placeholder="123.456.789-00"
                    class="w-full"
                    mask="999.999.999-99"
                />
            </div>
        </div>

        <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col gap-2 w-full md:w-2/12">
                <label for="rg">RG</label>
                <InputMask 
                    id="rg" 
                    v-model="dadosGeraisStore.dadosGerais.rg" 
                    placeholder="12.123-456" 
                    class="w-full"
                    mask="99.999-999"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-2/12">
                <label for="peso">Peso</label>
                <InputGroup>
                    <InputText 
                        id="peso" 
                        v-model.number="dadosGeraisStore.dadosGerais.peso" 
                        type="text"
                        v-keyfilter.int
                        maxlength="3" 
                        placeholder="64"
                    />
                    <InputGroupAddon>kg</InputGroupAddon>
                </InputGroup>
            </div>
            <div class="flex flex-col gap-2 w-full md:w-2/12">
                <label for="altura">Altura</label>
                <InputGroup>
                    <InputText 
                        id="altura" 
                        v-model.number="dadosGeraisStore.dadosGerais.altura" 
                        type="text"
                        v-keyfilter.int
                        maxlength="3" 
                        placeholder="181"
                    />
                    <InputGroupAddon>cm</InputGroupAddon>
                </InputGroup>
            </div>
            <div class="flex flex-col gap-2 w-full md:w-2/12">
                <label for="age">Idade</label>
                <InputGroup>
                    <InputText 
                        id="idade" 
                        v-model.number="dadosGeraisStore.dadosGerais.idade" 
                        type="text"
                        v-keyfilter.int
                        maxlength="3"
                        placeholder="32"
                    />
                    <InputGroupAddon>anos</InputGroupAddon>
                </InputGroup>
            </div>
            <div class="flex flex-col gap-2 w-full md:w-2/12">
                <label for="complexidade">Complexidade</label>
                <Select 
                    id="complexidade" 
                    v-model="dadosGeraisStore.dadosGerais.complexidade" 
                    :options="props.formFields.complexidade" 
                    optionLabel="label"
                    optionValue="label"
                    placeholder="Selecione" 
                    class="w-full"
                ></Select>
            </div>
            <div class="flex flex-col gap-2 py-2 w-full md:w-2/12">
                <label for="previsao_alta">Previsão de alta?</label>
                <InputSwitch 
                    id="previsao_alta" 
                    v-model="dadosGeraisStore.dadosGerais.previsao_alta" 
                />
            </div>
        </div>

        <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col gap-2 w-full md:w-1/2">
                <label for="nome_pai">Nome do pai</label>
                <InputText 
                    id="nome_pai" 
                    v-model="dadosGeraisStore.dadosGerais.nome_pai" 
                    type="text" 
                    placeholder="Glauberthy glauberthy"
                    v-keyfilter="/[a-zA-Z\sáàãâéèêíìîóòõôúùûçÇÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/"
            
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-1/2">
                <label for="nome_mae">Nome da mãe</label>
                <InputText 
                    id="nome_mae" 
                    v-model="dadosGeraisStore.dadosGerais.nome_mae" 
                    type="text" 
                    placeholder="Júlia silva"
                    v-keyfilter="/[a-zA-Z\sáàãâéèêíìîóòõôúùûçÇÁÀÃÂÉÈÊÍÌÎÓÒÕÔÚÙÛ]/"
                />
            </div>
        </div>

        <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col gap-2 w-full md:w-2/8">
                <label for="convenio">Convênio</label>
                <InputText 
                    id="convenio" 
                    v-model="dadosGeraisStore.dadosGerais.convenio" 
                    type="text" 
                    placeholder="Convenio médico xyz"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-2/8">
                <label for="hospital">Hospital de internação</label>
                <InputText 
                    id="hospital" 
                    v-model="dadosGeraisStore.dadosGerais.hospital" 
                    type="text" 
                    placeholder="Português"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-2/8">
                <label for="data_admissao">Data de admissão</label>
                <InputMask
                    id="data_admissao" 
                    v-model="dadosGeraisStore.dadosGerais.data_admissao" 
                    placeholder="dd/mm/aaaa" 
                    class="w-full"
                    showIcon="true"
                    mask="99/99/9999"
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-2/8">
                <label for="apartamento-hospital">Apartamento</label>
                <InputText 
                    id="apartamento_hospital" 
                    v-model="dadosGeraisStore.dadosGerais.apartamento" 
                    type="text" 
                    placeholder="Bloco A"
                />
            </div>
          
        </div>
        <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col gap-2 md:w-2/8">
                <label for="carteirinha">Nº Carteirinha</label>
                <InputText 
                    id="carteirinha" 
                    v-model="dadosGeraisStore.dadosGerais.carteirinha" 
                    type="text" 
                    placeholder="198310293"
                    v-keyfilter.int
                />
            </div>
            <div class="flex flex-col gap-2 w-full md:w-2/8">
                <label for="vencimento_carteirinha">Validade da carteirinha</label>
                <InputMask
                    id="vencimento_carteirinha" 
                    v-model="dadosGeraisStore.dadosGerais.vencimento_carteirinha" 
                    placeholder="dd/mm/aaaa" 
                    class="w-full"
                    showIcon="true"
                    mask="99/99/9999"
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