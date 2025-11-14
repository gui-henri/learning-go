<script setup>
import { ref } from 'vue';
import { useEvaluationStore } from '@/store/evaluation'

const store = useEvaluationStore();

const genero = ref([
    { name: 'Feminino', code: 'feminino' },
    { name: 'Masculino', code: 'masculino' }
    
]);

const complexidade = ref([
    { name: 'Baixa', code: 'baixa' },
    { name: 'Média', code: 'media' },
    { name: 'Alta', code: 'alta' }
    
]);

</script>

<template>
    <Fluid >

        <div class="card shadow-2xl rounded-2xl w-full sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                <h4 class="font-semibold text-xl">Dados Gerais</h4>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-wrap gap-2 w-full">
                        <label for="lastname2">Nome completo</label>
                        <InputText id="lastname2" v-model="store.evaluation.nome_paciente" type="text" placeholder="Glauberthy júnior"/>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4 flex-wrap">
                    <div class="flex flex-wrap gap-2 lg:w-40">
                        <label for="state">Gênero</label>
                        <Select id="state" v-model="store.evaluation.genero" :options="genero" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>
                    <div class="flex gap-4">
                        <div class="flex flex-wrap gap-2 w-24">
                            <label for="zip">Peso</label>
                            <InputGroup>
                                <InputText id="weight" v-model="store.evaluation.peso" type="text" maxlength="3" placeholder="64"/>
                                <InputGroupAddon>kg</InputGroupAddon>
                            </InputGroup>
                        </div>

                        <div class="flex flex-wrap gap-2 w-24">
                            <label for="zip">Altura</label>
                            <InputGroup>
                                <InputText id="height" v-model="store.evaluation.altura" type="text" maxlength="3" placeholder="181"/>
                                <InputGroupAddon>cm</InputGroupAddon>
                            </InputGroup>
                        </div>
                        <div class="flex flex-wrap gap-2 w-28">
                            <label for="zip">Idade</label>
                            <InputGroup>
                                <InputText id="age" v-model="store.evaluation.idade" type="text" maxlength="3" placeholder="32"/>
                                <InputGroupAddon>anos</InputGroupAddon>
                            </InputGroup>
                        </div>
                        <div class="flex flex-col gap-2"> <label for="zip">Previsão de alta?</label>
                         <InputSwitch v-model="checked" />
                    </div>
                     <div class="flex flex-col md:flex-row gap-4 flex-wrap">
                       <label for="zip">Previsão de alta?</label>
                        <InputSwitch v-model="store.previsao_alta" />
                    </div>
                    <div class="flex flex-wrap gap-2 lg:w-40">
                        <label for="state">Complexidade</label>
                        <Select id="state" v-model="store.evaluation.complexidade" :options="complexidade" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>
                </div>
              
                <div class="flex flex-col md:flex-row gap-4 flex-wrap" >
                      <div class="flex flex-wrap gap-2 lg:w-40">
                            <label for="zip">Data de nascimento</label>
                            <Calendar id="birthdate" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full"/>
                      </div>
                    
                </div>

                <div class="flex flex-wrap gap-2 w-full">
                        <label for="lastname2">Nome do pai</label>
                        <InputText id="lastname2" type="text" placeholder="Glauberthy glauberthy"/>
                    </div>
                <div class="flex flex-wrap gap-2 w-full">
                        <label for="lastname2">Nome da mãe</label>
                        <InputText id="lastname2" type="text" placeholder="Júlia silva"/>
                </div>
                <div class="flex flex-wrap gap-2 w-full">
                    <label for="lastname2">Endereço</label>
                    <InputText id="lastname2" type="text" placeholder="Rua das flores, 123 - Centro"/>

                </div>

            </div>
        </div>
    </Fluid>

</template>
