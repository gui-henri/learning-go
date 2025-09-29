<script setup>
import { ref } from "vue";
const pacientes = ref([
  { name: [{ text: "João" }], identifier: [{ value: "123.456.789-00" }] },
  { name: [{ text: "Maria" }], identifier: [{ value: "987.654.321-00" }] }
]);

const open = ref(false);

function editarPaciente(paciente) {
  console.log("Editar paciente:", paciente);
}
</script>

<template>
  <div class="w-full flex justify-center mt-4">
    <NuxtLink 
      to="/cadastro" 
      class="bg-red-600 hover:bg-red-700 text-white font-bold py-3 px-15 rounded-xl transition-all shadow-lg">
      <b>Cadastrar Paciente</b>
    </NuxtLink>
  </div>
  <h1><b>PACIENTES APARECEM AQUI EM BAIXO</b> </h1>

  <div class="cards-container">
    <div v-for="paciente in pacientes" :key="paciente.identifier[0].value" class="card">
      <h3>{{ paciente.name[0].text }}</h3>
      <p>CPF: {{ paciente.identifier[0].value }}</p>
      <button @click="editarPaciente(paciente)"><a>Editar</a></button>
    </div>
  </div>
  
    <div class="relative inline-block">
    <!-- Ícone de filtro -->
    <button @click="open = !open" class="p-2 rounded hover:bg-gray-100">
      ⚙️ <!-- você pode trocar por ícone SVG -->
    </button>

    <!-- Popover -->
    <div
      v-if="open"
      class="absolute right-0 mt-2 w-48 bg-white border rounded shadow-lg z-10"
    >
      <ul class="p-2">
        <li class="hover:bg-gray-100 p-2 cursor-pointer">Nome</li>
        <li class="hover:bg-gray-100 p-2 cursor-pointer">CPF</li>
        <li class="hover:bg-gray-100 p-2 cursor-pointer">Data de Nascimento</li>
      </ul>
    </div>
  </div>

</template>

<style scoped>
.cards-container {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.card {
  border: 1px solid #ccc;
  border-radius: 12px;
  padding: 15px;
  width: 200px;
  box-shadow: 2px 2px 6px rgba(0,0,0,0.1);
}
</style>