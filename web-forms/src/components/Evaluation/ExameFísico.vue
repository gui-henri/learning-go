<script setup>
import { ref, computed} from 'vue';
import { useExameFisicoStore } from '@/store/evaluation/exameFisico';

const estado_geral = ref([
    { name: 'Bom', code: 'bom' },
    { name: 'Regular', code: 'regular' },
    { name: 'Grave', code: 'grave' },
    { name: 'Gravíssimo', code: 'gravissimo' },
    { name: 'Decaído', code: 'decaido' }
]);

const locomocao = ref([
    { name: 'Deambula sem apoio', code: 'deambula_sem_apoio' },
    { name: 'Deambula com apoio', code: 'deambula_com_apoio' },
    { name: 'Restrito ao leito', code: 'restrito_leito' },
    { name: 'Cadeirante', code: 'cadeirante' }
]);

const tipos_piccline = ref([
    { name: 'Monolumen', code: 'monolumen' },
    { name: 'Duplolumen', code: 'duplolumen' },
    { name: 'Triplolumen', code: 'triplolumen' }
]);

const emit = defineEmits(['next-step']);

const exameFisicoStore = useExameFisicoStore();

const activeIndex = ref(null);

const isFilled = computed(() => {
    return !!exameFisicoStore.exameFisico.estado_geral
});

const handleSave = () => {
    activeIndex.value = null;
    
    setTimeout(() => {
        const self = document.getElementById("exame-fisico");
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
    <Accordion v-model:activeIndex="activeIndex" id="exame-fisico" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
            </i>
            <div class="flex flex-col text-left">
                <h4 class="font-semibold text-xl p-0 m-0">Exame Físico</h4>
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
                        :options="estado_geral" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/4">
                    <label for="avaliacao_locomotora">Avaliação locomotora <span class="text-red-500">*</span></label>
                    <Select 
                        id="avaliacao_locomotora" 
                        v-model="exameFisicoStore.exameFisico.avaliacao_locomotora" 
                        :options="locomocao" 
                        optionLabel="name" 
                        optionValue="code"
                        placeholder="Selecione" 
                        class="w-full"
                    />
                </div>
            </div>

            <div class="border-t my-2"></div>

            <div class="flex flex-col gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="possui_dreno">Drenos?</label>
                    <InputSwitch 
                        id="possui_dreno" 
                        v-model="exameFisicoStore.exameFisico.possui_dreno" 
                    />
                </div>

                <div v-if="exameFisicoStore.exameFisico.possui_dreno" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                    <div class="flex flex-col md:flex-row gap-4">
                        <div class="flex flex-col gap-2 w-full md:w-1/3">
                            <label for="local_dreno">Local do dreno</label>
                            <InputText 
                                id="local_dreno" 
                                v-model="exameFisicoStore.exameFisico.local_dreno" 
                                placeholder="Local anatômico" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label for="data_implantacao_dreno">Data implantação</label>
                            <Calendar 
                                id="data_implantacao_dreno" 
                                v-model="exameFisicoStore.exameFisico.data_implantacao_dreno" 
                                dateFormat="dd/mm/yy" 
                                placeholder="dd/mm/aaaa" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full flex-1">
                            <label for="curativos_dreno">Curativos do dreno</label>
                            <InputText 
                                id="curativos_dreno" 
                                v-model="exameFisicoStore.exameFisico.curativos_dreno" 
                                placeholder="Informações sobre curativos" 
                                class="w-full"
                            />
                        </div>
                    </div>
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
                        
                        <div v-if="exameFisicoStore.exameFisico.piccline" class="grid grid-cols-1 md:grid-cols-4 gap-4 pl-4 border-l-4 border-blue-200">
                            <div class="flex flex-col gap-2">
                                <label for="tipo_piccline">Tipo</label>
                                <Select id="tipo_piccline" v-model="exameFisicoStore.exameFisico.tipo_piccline" :options="tipos_piccline" optionLabel="name" optionValue="code" placeholder="Selecione" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_imp_piccline">Data Implantação</label>
                                <Calendar id="data_imp_piccline" v-model="exameFisicoStore.exameFisico.data_imp_piccline" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_troca_piccline">Última Troca</label>
                                <Calendar id="data_troca_piccline" v-model="exameFisicoStore.exameFisico.data_troca_piccline" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full" />
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
                        
                        <div v-if="exameFisicoStore.exameFisico.acesso_central" class="grid grid-cols-1 md:grid-cols-4 gap-4 pl-4 border-l-4 border-blue-200">
                            <div class="flex flex-col gap-2">
                                <label for="local_central">Local</label>
                                <InputText id="local_central" v-model="exameFisicoStore.exameFisico.local_central" placeholder="Local anatômico" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_imp_central">Data Implantação</label>
                                <Calendar id="data_imp_central" v-model="exameFisicoStore.exameFisico.data_imp_central" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_troca_central">Última Troca</label>
                                <Calendar id="data_troca_central" v-model="exameFisicoStore.exameFisico.data_troca_central" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full" />
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
                        
                        <div v-if="exameFisicoStore.exameFisico.acesso_periferico" class="grid grid-cols-1 md:grid-cols-3 gap-4 pl-4 border-l-4 border-blue-200">
                            <div class="flex flex-col gap-2">
                                <label for="local_periferico">Local</label>
                                <InputText id="local_periferico" v-model="exameFisicoStore.exameFisico.local_periferico" placeholder="Local anatômico" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_troca_periferico">Data Implantação</label>
                                <Calendar id="data_troca_periferico" v-model="exameFisicoStore.exameFisico.data_troca_periferico" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full" />
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
                            <label for="hipodermoclise">Hipodermóclise?</label>
                            <InputSwitch id="hipodermoclise" v-model="exameFisicoStore.exameFisico.hipodermoclise" />
                        </div>
                        
                        <div v-if="exameFisicoStore.exameFisico.hipodermoclise" class="grid grid-cols-1 md:grid-cols-2 gap-4 pl-4 border-l-4 border-blue-200">
                            <div class="flex flex-col gap-2">
                                <label for="data_troca_hipo">Última Troca</label>
                                <Calendar id="data_troca_hipo" v-model="exameFisicoStore.exameFisico.data_troca_hipo" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full" />
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
                        
                        <div v-if="exameFisicoStore.exameFisico.outros_acessos" class="grid grid-cols-1 md:grid-cols-3 gap-4 pl-4 border-l-4 border-blue-200">
                             <div class="flex flex-col gap-2">
                                <label for="nome_outros_acessos">Nome/Tipo</label>
                                <InputText id="nome_outros_acessos" v-model="exameFisicoStore.exameFisico.nome_outros_acessos" placeholder="Descreva o acesso" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_outro_acesso">Data Implantação</label>
                                <Calendar id="data_outro_acesso" v-model="exameFisicoStore.exameFisico.data_outro_acesso" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                            <div class="flex flex-col gap-2">
                                <label for="data_acesso_ultimo">Último acesso</label>
                                <Calendar id="data_acesso_ultimo" v-model="exameFisicoStore.exameFisico.data_acesso_ultimo" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full" />
                            </div>
                        </div>
                    </div>

                </div>
            </div>

            <div class="border-t my-2"></div>

            <div class="flex flex-col gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/4">
                    <label for="uso_antimicrobiano">Uso de antimicrobiano?</label>
                    <InputSwitch 
                        id="uso_antimicrobiano" 
                        v-model="exameFisicoStore.exameFisico.uso_antimicrobiano" 
                    />
                </div>

                <div v-if="exameFisicoStore.exameFisico.uso_antimicrobiano" class="flex flex-col gap-4 p-4 bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-400/10">
                    <div class="flex flex-col md:flex-row gap-4">
                        <div class="flex flex-col gap-2 w-full md:w-1/2">
                            <label for="nome_antimicrobiano">Nome do Antimicrobiano</label>
                            <InputText 
                                id="nome_antimicrobiano" 
                                v-model="exameFisicoStore.exameFisico.nome_antimicrobiano" 
                                placeholder="Nome da medicação" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label for="data_inicio_atb">Data Início</label>
                            <Calendar 
                                id="data_inicio_atb" 
                                v-model="exameFisicoStore.exameFisico.data_inicio_atb" 
                                dateFormat="dd/mm/yy" 
                                placeholder="dd/mm/aaaa" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label for="data_termino_atb">Data Término</label>
                            <Calendar 
                                id="data_termino_atb" 
                                v-model="exameFisicoStore.exameFisico.data_termino_atb" 
                                dateFormat="dd/mm/yy" 
                                placeholder="dd/mm/aaaa" 
                                class="w-full"
                            />
                        </div>
                    </div>
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