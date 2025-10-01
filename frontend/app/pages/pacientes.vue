<script setup>
const config = useRuntimeConfig();

const { data, pending, error, refresh } = await useAsyncData('Patient/all', () =>
  $fetch('/Patient', {
    baseURL: config.apiBase ?? "http://localhost:8090"
  })
)
</script>

<template>
  <div class="w-full flex justify-center mt-4">
    <NuxtLink 
      to="/cadastro" 
      class="bg-red-600 hover:bg-red-700 text-white font-bold py-3 px-15 rounded-xl transition-all shadow-lg">
      <b>Cadastrar Paciente</b>
    </NuxtLink>
  </div>
  <div v-if="pending" class="text-center text-gray-500">
        <p>Carregando pacientes...</p>
      </div>
  <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg text-center" role="alert">
    <strong>Ocorreu um erro:</strong>
    <p>{{ error.message }}</p>
  </div>
  
  <table v-else class="w-[896px]">
    <thead>
      <tr>
        <th class="text-start">Nome</th>
        <th class="text-start">Email</th>
        <th class="text-start">CPF</th>
        <th class="text-start">Data de Nascimento</th>
        <th class="text-start">Gênero</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(patient, index) in data.data.entry" :key="index">
        <td>{{ (patient.resource.name.find(n => n.use === "official" || n.use === "usual")?.given || ["Sem nome"]).join(" ") + patient.resource.name.find(n => n.use === "official" || n.use === "usual")?.family }}</td>
        <td>{{ (patient.resource.telecom.find(t => t.system ==="email"))?.value || "Sem Email" }}</td>
        <td>{{ (patient.resource.identifier.find(t => t.system ==="http://hl7.org.br/fhir/r4/sid/CPF")?.value) || "Sem CPF" }}</td>
        <td>{{ (patient.resource.birthDate) || "Sem Data" }}</td>
        <td>{{ (patient.resource.gender === "male" ? "Masculino" : patient.resource.gender === "female" ? "Feminino" : "Outro") }}</td>
      </tr>
    </tbody>
  </table>

</template>
