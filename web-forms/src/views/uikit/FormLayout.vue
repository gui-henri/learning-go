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

        <div class="bg-white shadow-2xl rounded-2xl w-full max-w-6xl p-6 sm:p-8 border-t-8 border-red-600">
            <div class="card flex flex-col gap-4 w-full">
                <div class="font-semibold text-xl">Dados Gerais
                </div>


                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-wrap gap-2 w-full">
                        <label for="lastname2">Nome completo</label>
                        <InputText id="lastname2" type="text" placeholder="Glauberthy glauberthy"/>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-wrap gap-2 w-full">
                        <label for="state">Gênero</label>
                        <Select id="state" v-model="selectedGender" :options="genero" optionLabel="name" placeholder="Selecione Gênero" class="w-full"></Select>
                    </div>
                    <div class="flex flex-wrap gap-2 w-72">
                        <label for="zip">Peso</label>
                        <InputGroup>
                            <InputText id="weight" type="text" maxlength="3"/>
                            <InputGroupAddon>kg</InputGroupAddon>
                        </InputGroup>
                    </div>

                     <div class="flex flex-wrap gap-2 w-72">
                        <label for="zip">Altura</label>
                        <InputGroup>
                            <InputText id="height" type="text" maxlength="3"/>
                            <InputGroupAddon>cm</InputGroupAddon>
                        </InputGroup>
                    </div>

                    <div class="flex flex-wrap gap-2 w-full">
                        <label for="state">Complexidade</label>
                        <Select id="state" v-model="selectedComplexidade" :options="complexidade" optionLabel="name" placeholder="Selecione Complexidade" class="w-full"></Select>
                    </div>

                </div>
            </div>
        </div>
    </Fluid>

    <Fluid>


    </Fluid>
</template>
