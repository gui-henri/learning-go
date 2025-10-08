<script setup>

import { ArrowLeft } from "lucide-vue-next";


const route = useRoute();
const pacienteId = route.params.id;


const { data: paciente, pending, error } = useAsyncData(
  `paciente-${pacienteId}`,
  () => $fetch(`http://localhost:8090/Patient/${pacienteId}`)
);


useHead({
  title: computed(() => {
    if (pending.value) return 'Carregando Paciente...';
    if (error.value || !paciente.value) return 'Paciente não encontrado';
    

    const nameArray = paciente.value?.name?.[0]?.given ?? [];
    const name = nameArray.join(' ');
    return `Detalhes de ${name || 'Paciente'}`;
  }),
});
</script>

<template>
  <div>
    <div v-if="pending" class="min-h-screen flex items-center justify-center">
      <p class="text-xl text-gray-500 animate-pulse">Carregando informações do paciente...</p>
    </div>

    <div v-else-if="error" class="min-h-screen flex items-center justify-center p-4">
       <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg text-center" role="alert">
        <strong>Ocorreu um erro ao buscar o paciente:</strong>
        <p>{{ error.message || 'Verifique se o backend está rodando e o ID é válido.' }}</p>
        <NuxtLink to="/pacientes" class="text-blue-600 hover:underline mt-4 inline-block">
          &larr; Voltar para a lista
        </NuxtLink>
      </div>
    </div>

    <div v-else-if="paciente" class="min-h-screen flex items-start justify-start p-4 sm:p-6">
      <div class="bg-white shadow-2xl rounded-2xl w-full md:w-1/2 p-6 sm:p-8 border-t-8 border-red-600">
        
        <header class="relative w-full text-center mb-8">
          <NuxtLink 
            to="/" 
            class="absolute left-0 top-1/2 -translate-y-1/2 bg-red-600 hover:bg-red-700 text-white p-3 rounded-full transition-colors"
            title="Voltar"
          >
            <ArrowLeft class="w-5 h-5" />
          </NuxtLink>
          <h1 class="text-3xl font-bold text-gray-800 inline-block">
            Ficha do Paciente
          </h1>
        </header>

        <main class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-8 pt-4 border-t border-gray-200">
          
          <div>
            <h3 class="text-sm font-semibold text-gray-600">Nome Completo</h3>
            <p class="text-lg text-gray-900 mt-1">{{ paciente.name[0].given.join(' ') }}</p>
          </div>

          <div>
            <h3 class="text-sm font-semibold text-gray-600">CPF</h3>
            <p class="text-lg text-gray-900 mt-1">{{ paciente.identifier[0].value }}</p>
          </div>

          <div>
            <h3 class="text-sm font-semibold text-gray-600">Data de Nascimento</h3>
            <p class="text-lg text-gray-900 mt-1">
              {{ new Date(paciente.birthDate).toLocaleDateString('pt-BR', { timeZone: 'UTC' }) }}
            </p>
          </div>
          
          <div>
            <h3 class="text-sm font-semibold text-gray-600">Gênero</h3>
            <p class="text-lg text-gray-900 mt-1 capitalize">{{ paciente.gender }}</p>
          </div>

          <div>
            <h3 class="text-sm font-semibold text-gray-600">Telefone</h3>
            <p class="text-lg text-gray-900 mt-1">{{ paciente.telecom.find(t => t.system === 'phone')?.value || 'Não informado' }}</p>
          </div>
          
          <div>
            <h3 class="text-sm font-semibold text-gray-600">Email</h3>
            <p class="text-lg text-gray-900 mt-1">{{ paciente.telecom.find(t => t.system === 'email')?.value || 'Não informado' }}</p>
          </div>

        </main>
      </div>
    </div>
  </div>
</template>