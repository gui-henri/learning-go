<script setup>
import { ref } from 'vue';

const genero = ref([
    { name: 'Feminino', code: 'feminino' },
    { name: 'Masculino', code: 'masculino' }
    
]);

const dropdownItem = ref(null);
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
                        <InputText id="lastname2" type="text" placeholder="GLAUBERTHY"/>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-wrap gap-2 w-full">
                        <label for="state">Gênero</label>
                        <Select id="state" v-model="genero" :options="genero" optionLabel="name" placeholder="Selecione Gênero" class="w-full"></Select>
                    </div>
                    <div class="flex flex-wrap gap-2 w-full">
                        <label for="zip">Peso</label>
                        <InputText id="zip" type="text" />
                    </div>
                     <div class="flex flex-wrap gap-2 w-full">
                        <label for="zip">Altura</label>
                        <InputText id="zip" type="text" />
                    </div>
                     <div class="flex flex-wrap gap-2 w-full">
                        <label for="zip">Zip</label>
                        <InputText id="zip" type="text" />
                    </div>

                </div>
            </div>
        </div>
    </Fluid>

    <Fluid>


    </Fluid>
</template>
