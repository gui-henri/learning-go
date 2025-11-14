<script setup>
import { ref } from 'vue';

const genero = ref([
    { name: 'Feminino', code: 'feminino' },
    { name: 'Masculino', code: 'masculino' }
    
]);

const complexidade = ref([
    { name: 'Baixa', code: 'baixa' },
    { name: 'Média', code: 'media' },
    { name: 'Alta', code: 'alta' }
    
]);

const selectedComplexidade = ref(null);
const selectedGender = ref(null);
</script>

<template>
    <Fluid >

        <div class="card shadow-2xl rounded-2xl w-full max-w-6xl sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                <h4 class="font-semibold text-xl">Dados Gerais</h4>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-wrap gap-2 w-full">
                        <label for="lastname2">Nome completo</label>
                        <InputText id="lastname2" type="text" placeholder="Glauberthy glauberthy"/>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4 flex-wrap">
                    <div class="flex flex-wrap gap-2 lg:w-40">
                        <label for="state">Gênero</label>
                        <Select id="state" v-model="selectedGender" :options="genero" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>
                    <div class="flex gap-4">
                        <div class="flex flex-wrap gap-2 w-24">
                            <label for="zip">Peso</label>
                            <InputGroup>
                                <InputText id="weight" type="text" maxlength="3" placeholder="64"/>
                                <InputGroupAddon>kg</InputGroupAddon>
                            </InputGroup>
                        </div>

                        <div class="flex flex-wrap gap-2 w-24">
                            <label for="zip">Altura</label>
                            <InputGroup>
                                <InputText id="height" type="text" maxlength="3" placeholder="181"/>
                                <InputGroupAddon>cm</InputGroupAddon>
                            </InputGroup>
                        </div>
                        <div class="flex flex-wrap gap-2 w-28">
                            <label for="zip">Idade</label>
                            <InputGroup>
                                <InputText id="age" type="text" maxlength="3" placeholder="32"/>
                                <InputGroupAddon>anos</InputGroupAddon>
                            </InputGroup>
                        </div>
                    </div>
                     <div class="flex flex-col md:flex-row gap-4 flex-wrap">
                       <label for="zip">Previsão de alta?</label>
                            <InputSwitch v-model="checked" />
                        

                    </div>
                    <div class="flex flex-wrap gap-2 lg:w-40">
                        <label for="state">Complexidade</label>
                        <Select id="state" v-model="selectedComplexidade" :options="complexidade" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>

                </div>
            </div>
        </div>
    </Fluid>

    <Fluid>


    </Fluid>
</template>
