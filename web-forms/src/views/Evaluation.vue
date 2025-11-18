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
        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                <h4 class="font-semibold text-xl">Dados Gerais</h4>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full">
                        <label for="nome_paciente">Nome completo</label>
                        <InputText id="nome_paciente" v-model="store.evaluation.nome_paciente" type="text" placeholder="Glauberthy júnior"/>
                    </div>
                </div>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2  md:w-1/4">
                        <label for="genero">Gênero</label>
                        <Select id="genero" v-model="store.evaluation.genero" :options="genero" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="birthdate">Data de nascimento</label>
                        <Calendar id="birthdate" v-model="store.evaluation.dat_nascimento" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full"/>
                    </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="birthdate">CPF</label>
                        <inputText id="cpf" v-model="store.evaluation.cpf" placeholder="123.456.789-00"class="w-full"/>
                    </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="birthdate">RG</label>
                        <inputText id="rg" v-model="store.evaluation.rg" placeholder="12.123-456" class="w-full"/>
                    </div>
                </div>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="weight">Peso</label>
                        <InputGroup>
                            <InputText id="weight" v-model="store.evaluation.peso" type="text" maxlength="3" placeholder="64"/>
                            <InputGroupAddon>kg</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="height">Altura</label>
                        <InputGroup>
                            <InputText id="height" v-model="store.evaluation.altura" type="text" maxlength="3" placeholder="181"/>
                            <InputGroupAddon>cm</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="age">Idade</label>
                        <InputGroup>
                            <InputText id="age" v-model="store.evaluation.idade" type="text" maxlength="3" placeholder="32"/>
                            <InputGroupAddon>anos</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="complexidade">Complexidade</label>
                        <Select id="complexidade" v-model="store.evaluation.complexidade" :options="complexidade" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>
                    <div class="flex flex-col gap-2 py-2 w-full m"> 
                        <label for="previsao_alta">Previsão de alta?</label>
                        <InputSwitch id="previsao_alta" v-model="store.evaluation.previsao_alta" />
                    </div>
                </div>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full">
                        <label for="nome_pai">Nome do pai</label>
                        <InputText id="nome_pai" v-model="store.evaluation.nome_pai" type="text" placeholder="Glauberthy glauberthy"/>
                    </div>
                </div>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full">
                        <label for="nome_mae">Nome da mãe</label>
                        <InputText id="nome_mae" v-model="store.evaluation.nome_mae" type="text" placeholder="Júlia silva"/>
                    </div>
                </div>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_mae">Convenio médico</label>
                        <InputText id="convenio" v-model="store.evaluation.convenio" type="text" placeholder="Convenio médico xyz"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/2" >
                        <label for="nome_mae">Hospital de internação</label>
                        <InputText id="hospital" v-model="store.evaluation.hospital" type="text" placeholder="Portugês"/>
                    </div>
                </div>
                  <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="birthdate">Data de admissão</label>
                        <Calendar id="admissao" v-model="store.evaluation.dat_admissao" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4" >
                        <label for="nome_mae">Apartamento hospitalar</label>
                        <InputText id="apartamento" v-model="store.evaluation.apartamento" type="text" placeholder="Portugês"/>
                    </div>
                    <div class="flex flex-col gap-2 md:w-1/8" >
                        <label for="nome_mae">Nº Carteirinha</label>
                        <InputText id="carteirinha" v-model="store.evaluation.carteirinha" type="text" placeholder="198310293"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="birthdate">Data de validade da Carteirinha</label>
                        <Calendar id="vencimento-carteirinha" v-model="store.evaluation.vencimento_carteirinha" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full"/>
                    </div>
                </div>
            </div>
        </div>
    </Fluid>
</template>