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
    {{ data }}
    <thead>
      <tr>
        <th class="text-start">Invoice</th>
        <th class="text-start">Payment Status</th>
        <th class="text-start">Payment Method</th>
        <th class="text-start">Amount</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="invoice in invoices" :key="invoice.invoice">
        <td>{{ invoice.invoice }}</td>
        <td>{{ invoice.paymentStatus }}</td>
        <td>{{ invoice.paymentMethod }}</td>
        <td>{{ invoice.totalAmount }}</td>
      </tr>
    </tbody>
  </table>

</template>
