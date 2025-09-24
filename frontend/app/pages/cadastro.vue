<script setup>
import { ref } from "vue";

// Recurso FHIR Patient
const paciente = ref({
  resourceType: "Patient",
  identifier: [{ system: "http://hl7.org.br/fhir/r4/sid/CPF", value: "" }],
  name: [{ use: "official", family: "", given: [] }],
  gender: "",
  birthDate: "",
  telecom: [{ system: "phone", value: "" }, { system: "email", value: "" }],
});

// Função de cadastro com envio para backend
const cadastrar = async () => {
  try {
    console.log("[INFO] Patient being send: ", paciente.value)
    const res = await $fetch('http://localhost:8090/Patient', {
      method: 'POST',
      body: JSON.stringify({
        paciente: paciente.value
      }),
      headers: {
        'Content-Type': 'application/json',
      },
    });

    alert('Paciente cadastrado com sucesso!');
    console.log('Resposta do backend:', res);
  } catch (err) {
    alert('Erro ao cadastrar paciente');
    console.error(err);
  }
};

// Função para formatar telefone enquanto digita
const formatarTelefone = (e) => {
  let valor = e.target.value.replace(/\D/g, ""); // remove tudo que não é número
  if (valor.length > 2) {
    valor = `(${valor.slice(0, 2)}) ${valor.slice(2)}`;
  }
  if (valor.length > 9) {
    valor = `${valor.slice(0, 9)}-${valor.slice(9, 13)}`;
  }
  paciente.value.telecom[0].value = valor;
};
</script>

<template>
  <div class="min-h-screen flex items-center justify-center p-6">
    <!-- Caixa com borda vermelha interna -->
    <div class="bg-white shadow-2xl rounded-2xl w-full max-w-lg p-6 border-4 border-red-600">
      <h1 class="text-3xl font-bold text-red-600 text-center mb-6">
        Cadastro de Paciente (FHIR)
      </h1>

      <form @submit.prevent="cadastrar" class="space-y-4">
<!-- Nome -->
<div>
  <label class="block text-gray-700 font-semibold">Nome</label>
  <input
    v-model="paciente.name[0].given[0]"
    type="text"
    required
    pattern="[A-Za-zÀ-ÿ\s]+"
    title="Apenas letras são permitidas"
    @input="paciente.name[0].given[0] = paciente.name[0].given[0].replace(/[^A-Za-zÀ-ÿ\s]/g, '')"
    class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-red-500"
  />
</div>

<!-- Sobrenome -->
<div>
  <label class="block text-gray-700 font-semibold">Sobrenome</label>
  <input
    v-model="paciente.name[0].family"
    type="text"
    required
    pattern="[A-Za-zÀ-ÿ\s]+"
    title="Apenas letras são permitidas"
    @input="paciente.name[0].family = paciente.name[0].family.replace(/[^A-Za-zÀ-ÿ\s]/g, '')"
    class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-red-500"
  />
</div>


        <!-- CPF -->
        <div>
          <label class="block text-gray-700 font-semibold">CPF</label>
          <input
            v-model="paciente.identifier[0].value"
            type="text"   
            title="Apenas números."
            required
            class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-red-500"
          />
        </div>

        <!-- Data de Nascimento -->
        <div>
          <label class="block text-gray-700 font-semibold">Data de Nascimento</label>
          <input
            v-model="paciente.birthDate"
            type="date"
            title="Utilize o padrão dd/mm/aaaa."
            required
            class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-red-500"
          />
        </div>

        <!-- Gênero -->
        <div>
          <label class="block text-gray-700 font-semibold">Gênero</label>
          <select
            v-model="paciente.gender"
            required
            class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-red-500"
          >
            <option disabled value="">Selecione</option>
            <option value="male">Masculino</option>
            <option value="female">Feminino</option>
            <option value="other">Outro</option>
          </select>
        </div>

        <!-- Telefone -->
        <div>
          <label class="block text-gray-700 font-semibold">Telefone</label>
          <input
            v-model="paciente.telecom[0].value"
            type="tel"
            title="Utilize apenas números."
            class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-red-500"
            @input="formatarTelefone"
            maxlength="16"
          />
        </div>

        <!-- Email -->
        <div>
          <label class="block text-gray-700 font-semibold">Email</label>
          <input
            v-model="paciente.telecom[1].value"
            type="email"
            class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-red-500"
          />
        </div>

        <!-- Botão -->
        <button
          type="submit"
          class="w-full bg-red-600 hover:bg-red-700 text-white font-bold py-3 rounded-xl transition-all shadow-lg"
        >
          Cadastrar (FHIR JSON)
        </button>
      </form>
    </div>
  </div>
</template>
