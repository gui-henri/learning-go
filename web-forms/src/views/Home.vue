<script setup>
import { ref } from 'vue';
import InputText from 'primevue/inputtext';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
// Certifique-se de importar os componentes da DataTable
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Card from 'primevue/card';
import Button from 'primevue/button';

const searchQuery = ref(''); // Adicionado para não dar erro no v-model

const pacientes = ref([
    { id: 1, nome: 'João Silva', genero: 'Masculino', idade: 45 },
    { id: 2, nome: 'Maria Oliveira', genero: 'Feminino', idade: 32 },
    { id: 3, nome: 'Pedro Santos', genero: 'Masculino', idade: 28 },
]);

const emdesenvolvimento = () => alert('Em desenvolvimento');
</script>

<template>
    <div class="flex justify-center w-full p-2">
        <InputGroup :style="{ width: '700px' }">
            <InputGroupAddon>
                <i class="pi pi-search"></i>
            </InputGroupAddon>
            <InputText v-model="searchQuery" placeholder="Digite" />
        </InputGroup>
    </div>
        <div class="card border-t-8 border-red-600 shadow-md rounded-2xl overflow-hidden">
            <DataTable :value="pacientes" tableStyle="min-width: 50rem">
                <Column field="nome" header="Nome"></Column>
                <Column field="genero" header="Gênero"></Column>
                <Column field="idade" header="Idade"></Column>
                <Column header="Ações">
                    <template #body>
                        <Button icon="pi pi-search" text rounded @click="emdesenvolvimento" />
                    </template>
                </Column>
            </DataTable>
        </div>

</template>