<script setup>
import { ref, computed} from 'vue';
import { useExameFisicoStore } from '@/store/evaluation/exameFisico';
import { InputMask } from 'primevue';
import { useStepAccordion } from "@/composable/useStepAccordion";

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

const exameFisicoStore = useExameFisicoStore();

const isFilled = computed(() => {
    return !!exameFisicoStore.exameFisico.estado_geral
});

const { internalIndex, nextStep } = useStepAccordion(props, emit);

const handleSave = () => {
    internalIndex.value = null;
    
    setTimeout(() => {
        const self = document.getElementById("exame-fisico");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    nextStep('next-step');
};

const addAntimicrobiano = () => {
    exameFisicoStore.adicionarAntimicrobiano();
};

const removeAntimicrobiano = (index) => {
    exameFisicoStore.removerAntimicrobiano(index);
};

</script>
<template>
    <Accordion v-model:activeIndex="internalIndex" id="exame-fisico" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Exame Físico <i class="pi pi-user text-3xl text-black"></i></h4>
                <span class="text-xs text-gray-500 font-normal -mt-4">
                    {{ isFilled ? '' : 'Toque para preencher' }}
                </span>
            </div>
        </div>
    </template>
        <div class="flex flex-col gap-4 w-full">
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/4">
                    <label for="estado_geral">Estado geral <span class="text-red-500">*</span></label>
                    <Select 
                        id="estado_geral" 
                        v-model="exameFisicoStore.exameFisico.estado_geral" 
                        :options="formFields.estado_geral" 
                        optionLabel="label" 
                        optionValue="label"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-2/4">
                    <label for="avaliacao_locomotora">Avaliação locomotora <span class="text-red-500">*</span></label>
                    <Select 
                        id="avaliacao_locomotora" 
                        v-model="exameFisicoStore.exameFisico.avaliacao_locomotora" 
                        :options="formFields.locomocao" 
                        optionLabel="label" 
                        optionValue="label"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
            </div>

            <div class="border-t my-2"></div>
            <div class="flex flex-col gap-4 w-full">
                <div v-for="(dreno, index) in exameFisicoStore.exameFisico.drenos" :key="index" class="p-4 border border-gray-200 rounded-lg bg-gray-50 dark:bg-gray-800 flex flex-col gap-4">
                    
                    <div class="flex justify-between items-center">
                        <h5 class="font-bold text-gray-600">Dreno {{ index + 1 }}</h5>
                        <Button icon="pi pi-trash" 
                                class="p-button-rounded p-button-danger p-button-text" 
                                @click="exameFisicoStore.removerDreno(index)" />
                    </div>

                    <div class="flex flex-col md:flex-row gap-4">
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label :for="'local_dreno_' + index">Local do dreno</label>
                            <InputText 
                                :id="'local_dreno_' + index" 
                                v-model="dreno.local_dreno" 
                                placeholder="Local anatômico" 
                                class="w-full"
                            />
                        </div>

                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label :for="'data_implantacao_dreno_' + index">Data implantação</label>
                            <InputMask
                                :id="'data_implantacao_dreno_' + index" 
                                v-model="dreno.data_implantacao_dreno" 
                                placeholder="dd/mm/aaaa" 
                                class="w-full"
                                mask="99/99/9999"
                            />
                        </div>

                        <div class="flex flex-col gap-2 w-full flex-1">
                            <label :for="'curativos_dreno_' + index">Curativos do dreno</label>
                            <InputText 
                                :id="'curativos_dreno_' + index" 
                                v-model="dreno.curativos_dreno" 
                                placeholder="Informações sobre curativos" 
                                class="w-full"
                            />
                        </div>
                    </div>
                </div>

                <div v-if="exameFisicoStore.exameFisico.drenos.length === 0" class="text-gray-500 italic text-sm p-2">
                    Nenhum dreno registrado.
                </div>

                <div class="flex justify-end w-full">
                    <Button label="Adicionar dreno" icon="pi pi-plus" class="p-button-outlined p-button-secondary" @click="exameFisicoStore.adicionarDreno" />
                </div>
            </div>

            <div class="border-t my-2"></div>

            <div class="flex flex-col gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="possui_acesso">Possui acessos?</label>
                    <InputSwitch 
                        id="possui_acesso" 
                        v-model="exameFisicoStore.exameFisico.possui_acesso" 
                    />
                </div>

                <div v-if="exameFisicoStore.exameFisico.possui_acesso" class="flex flex-col gap-6 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                    
                    <div class="flex items-center gap-4">
                        <label for="cateter_total">Cateter total implantado?</label>
                        <InputSwitch id="cateter_total" v-model="exameFisicoStore.exameFisico.cateter_total" />
                    </div>

                    <hr class="border-gray-300"/>

                    <div class="flex flex-col gap-4">
                        <div class="flex items-center gap-4">
                            <label for="piccline">PiccLine?</label>
                            <InputSwitch id="piccline" v-model="exameFisicoStore.exameFisico.piccline" />
                        </div>
                        
                        <div v-if="exameFisicoStore.exameFisico.piccline" class="grid grid-cols-1 md:grid-cols-4 gap-4 pl-4 border-l-4 border-red-600   ">
                            <div class="flex flex-col gap-2">
                                <label for="tipo_piccline">Tipo</label>
                                <Select id="tipo_piccline" v-model="exameFisicoStore.exameFisico.tipo_piccline" :options="formFields.tipo_piccline" optionLabel="label" optionValue="label" placeholder="Selecione" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_imp_piccline">Data Implantação</label>
                                <InputMask id="data_imp_piccline" v-model="exameFisicoStore.exameFisico.data_imp_piccline" mask="99/99/9999" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_troca_piccline">Última Troca</label>
                                <InputMask id="data_troca_piccline" v-model="exameFisicoStore.exameFisico.data_troca_piccline" mask="99/99/9999" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="curativo_piccline">Curativo</label>
                                <InputText id="curativo_piccline" v-model="exameFisicoStore.exameFisico.curativo_piccline" placeholder="Nome do curativo" class="w-full" />
                            </div>
                        </div>
                    </div>

                    <hr class="border-gray-300"/>

                    <div class="flex flex-col gap-4">
                        <div class="flex items-center gap-4">
                            <label for="acesso_central">Acesso Venoso Central?</label>
                            <InputSwitch id="acesso_central" v-model="exameFisicoStore.exameFisico.acesso_central" />
                        </div>

                        <div v-if="exameFisicoStore.exameFisico.acesso_central" class="grid grid-cols-1 md:grid-cols-4 gap-4 pl-4 border-l-4 border-red-600">
                            <div class="flex flex-col gap-2">
                                <label for="local_central">Local</label>
                                <InputText id="local_central" v-model="exameFisicoStore.exameFisico.local_central" placeholder="Local anatômico" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_imp_central">Data Implantação</label>
                                <InputMask id="data_imp_central" v-model="exameFisicoStore.exameFisico.data_imp_central" mask="99/99/9999" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_troca_central">Última Troca</label>
                                <InputMask id="data_troca_central" v-model="exameFisicoStore.exameFisico.data_troca_central" mask="99/99/9999" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="curativo_central">Curativo</label>
                                <InputText id="curativo_central" v-model="exameFisicoStore.exameFisico.curativo_central" placeholder="Nome do curativo" class="w-full" />
                            </div>
                        </div>
                    </div>

                    <hr class="border-gray-300"/>

                    <div class="flex flex-col gap-4">
                        <div class="flex items-center gap-4">
                            <label for="acesso_periferico">Acesso Venoso Periférico?</label>
                            <InputSwitch id="acesso_periferico" v-model="exameFisicoStore.exameFisico.acesso_periferico" />
                        </div>
                        
                        <div v-if="exameFisicoStore.exameFisico.acesso_periferico" class="grid grid-cols-1 md:grid-cols-3 gap-4 pl-4 border-l-4 border-red-600">
                            <div class="flex flex-col gap-2">
                                <label for="local_periferico">Local</label>
                                <InputText id="local_periferico" v-model="exameFisicoStore.exameFisico.local_periferico" placeholder="Local anatômico" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_troca_periferico">Data Última Troca</label>
                                <InputMask id="data_troca_periferico" v-model="exameFisicoStore.exameFisico.data_troca_periferico" mask="99/99/9999" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="curativo_periferico">Curativo</label>
                                <InputText id="curativo_periferico" v-model="exameFisicoStore.exameFisico.curativo_periferico" placeholder="Nome do curativo" class="w-full" />
                            </div>
                        </div>
                    </div>

                    <hr class="border-gray-300"/>

                    <div class="flex flex-col gap-4">
                        <div class="flex items-center gap-4">
                            <label for="hipodermoclise">hipodermoclise?</label>
                            <InputSwitch id="hipodermoclise" v-model="exameFisicoStore.exameFisico.hipodermoclise" />
                        </div>

                        <div v-if="exameFisicoStore.exameFisico.hipodermoclise" class="grid grid-cols-1 md:grid-cols-2 gap-4 pl-4 border-l-4 border-red-600">
                            <div class="flex flex-col gap-2">
                                <label for="data_troca_hipo">Última Troca</label>
                                <InputMask id="data_troca_hipo" v-model="exameFisicoStore.exameFisico.data_troca_hipo" mask="99/99/9999" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="curativo_hipo">Curativo</label>
                                <InputText id="curativo_hipo" v-model="exameFisicoStore.exameFisico.curativo_hipo" placeholder="Nome do curativo" class="w-full" />
                            </div>
                        </div>
                    </div>

                    <hr class="border-gray-300"/>

                    <div class="flex flex-col gap-4">
                        <div class="flex items-center gap-4">
                            <label for="outros_acessos">Outros acessos?</label>
                            <InputSwitch id="outros_acessos" v-model="exameFisicoStore.exameFisico.outros_acessos" />
                        </div>

                        <div v-if="exameFisicoStore.exameFisico.outros_acessos" class="grid grid-cols-1 md:grid-cols-3 gap-4 pl-4 border-l-4 border-red-600">
                             <div class="flex flex-col gap-2">
                                <label for="nome_outros_acessos">Nome/Tipo</label>
                                <InputText id="nome_outros_acessos" v-model="exameFisicoStore.exameFisico.nome_outros_acessos" placeholder="Descreva o acesso" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_outro_acesso">Data Implantação</label>
                                <InputMask id="data_outro_acesso" v-model="exameFisicoStore.exameFisico.data_outro_acesso" mask="99/99/9999" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_acesso_ultimo">Último acesso</label>
                                <InputMask id="data_acesso_ultimo" v-model="exameFisicoStore.exameFisico.data_acesso_ultimo" mask="99/99/9999" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                        </div>
                    </div>

                </div>
            </div>

            <div class="border-t my-2"></div>
            <div class="flex flex-col gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <span for="possui_dreno">Antimicrobianos</span>
                </div>
            </div>
            <div class="flex flex-col gap-4">
                <div v-if="exameFisicoStore.exameFisico.antimicrobianos.length === 0" class="text-gray-500 italic text-sm p-2">
                    Nenhum antimicrobiano registrado.
                </div>
                <div v-for="(item, index) in exameFisicoStore.exameFisico.antimicrobianos" :key="index" class="p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10 flex flex-col gap-4 relative">
                    
                    <div class="flex justify-between items-center">
                        <h6 class="font-bold text-gray-600 m-0">Antimicrobiano {{ index + 1 }}</h6>
                        <Button v-if="exameFisicoStore.exameFisico.antimicrobianos.length > 0" 
                                icon="pi pi-trash" 
                                class="p-button-rounded p-button-danger p-button-text w-8 h-8" 
                                @click="removeAntimicrobiano(index)" 
                                v-tooltip="'Remover item'" />
                    </div>

                    <div class="flex flex-col md:flex-row gap-4">
                        <div class="flex flex-col gap-2 w-full md:w-1/2">
                            <label :for="'nome_antimicrobiano_' + index">Nome do Antimicrobiano</label>
                            <InputText 
                                :id="'nome_antimicrobiano_' + index" 
                                v-model="item.nome_antimicrobiano" 
                                placeholder="Nome da medicação" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label :for="'data_inicio_' + index">Data Início</label>
                            <InputMask 
                                :id="'data_inicio_' + index" 
                                v-model="item.data_inicio" 
                                mask="99/99/9999"
                                placeholder="dd/mm/aaaa" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label :for="'data_termino_' + index">Data Término</label>
                            <InputMask
                                :id="'data_termino_' + index"
                                v-model="item.data_termino"
                                placeholder="dd/mm/aaaa"
                                class="w-full"
                                mask="99/99/9999"
                            />
                        </div>
                    </div>
                </div>

                <div class="flex justify-end w-full">
                    <Button label="Adicionar antimicrobiano" icon="pi pi-plus" class="p-button-outlined p-button-secondary" @click="addAntimicrobiano" />
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