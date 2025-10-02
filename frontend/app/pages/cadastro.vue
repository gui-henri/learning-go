<script setup>
import { ArrowLeft } from "lucide-vue-next";
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

// Função para formatar telefone (com suporte ao 9º dígito)
const formatarTelefone = (e) => {
  let numeros = e.target.value.replace(/\D/g, "").slice(0, 11);

  if (numeros.length > 10) {
    // Máscara para celular com 9º dígito: (XX) 9XXXX-XXXX
    numeros = numeros.replace(/^(\d{2})(\d{5})(\d{4})/, "($1) $2-$3");
  } else if (numeros.length > 6) {
    // Máscara para telefone fixo: (XX) XXXX-XXXX
    numeros = numeros.replace(/^(\d{2})(\d{4})(\d{0,4})/, "($1) $2-$3");
  } else if (numeros.length > 2) {
    // Máscara para DDD e início do número: (XX) XXXX
    numeros = numeros.replace(/^(\d{2})(\d*)/, "($1) $2");
  } else {
    if (numeros.length > 0) {
      numeros = numeros.replace(/^(\d*)/, "($1");
    }
  }
  paciente.value.telecom[0].value = numeros;
};

// Função para formatar o CPF
const formatarCPF = (e) => {
  let valor = e.target.value.replace(/\D/g, "");
  valor = valor.slice(0, 11);
  valor = valor.replace(/(\d{3})(\d)/, "$1.$2");
  valor = valor.replace(/(\d{3})\.(\d{3})(\d)/, "$1.$2.$3");
  valor = valor.replace(/(\d{3})\.(\d{3})\.(\d{3})(\d{1,2})/, "$1.$2.$3-$4");
  paciente.value.identifier[0].value = valor;
};
</script>

<template>
  <div>
    <div class="min-h-screen flex items-start justify-center p-4 sm:p-6 mt-8">
      <div class="bg-white shadow-2xl rounded-2xl w-full max-w-6xl p-6 sm:p-8 border-t-8 border-red-600">
        
        <div class="relative w-full text-center mb-8">

          <NuxtLink 
            to="/pacientes" 
            class="absolute left-0 top-1/2 -translate-y-1/2 bg-red-600 hover:bg-red-700 p-3 rounded-full transition-colors duration-200"
            aria-label="Voltar para a lista de pacientes"
            title="Voltar para a lista de pacientes"
          >
            <ArrowLeft />
          </NuxtLink>
  
          <h1 class="text-3xl font-bold text-gray-800 inline-block">
            Ficha de Cadastro do Paciente
          </h1>
        </div>
        

        <form @submit.prevent="cadastrar" class="grid grid-cols-1 md:grid-cols-3 gap-x-8 gap-y-6">
          
          <div>
            <label class="block text-gray-700 font-semibold mb-1">Nome Completo</label>
            <input
              v-model="paciente.name[0].given[0]"
              placeholder="Ex: João Almeida do Carmo"
              type="text"
              required
              pattern="[A-Za-zÀ-ÿ\s]+"
              title="Apenas letras são permitidas"
              @input="paciente.name[0].given[0] = paciente.name[0].given[0].replace(/[^A-Za-zÀ-ÿ\s]/g, '')"
              class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 transition"
            />
          </div>

          <div>
            <label class="block text-gray-700 font-semibold mb-1">CPF</label>
            <input
              v-model="paciente.identifier[0].value"
              type="text"
              required
              maxlength="14"
              placeholder="000.000.000-00"
              title="Digite apenas números. O formato será aplicado automaticamente."
              @input="formatarCPF"
              class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 transition"
            />
          </div>

          <div>
            <label class="block text-gray-700 font-semibold mb-1">Data de Nascimento</label>
            <input
              v-model="paciente.birthDate"
              type="date"
              title="Utilize o padrão dd/mm/aaaa."
              required
              class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 transition"
            />
          </div>
          
          <div>
            <label class="block text-gray-700 font-semibold mb-1">Gênero</label>
            <select
              v-model="paciente.gender"
              title="Escolha seu genêro."
              required
              class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 transition bg-white"
            >
              <option disabled value="">Selecione uma opção</option>
              <option value="male">Masculino</option>
              <option value="female">Feminino</option>
              <option value="other">Desconhecido</option>
              <option value="other">Outro</option>
            </select>
          </div>

          <div>
            <label class="block text-gray-700 font-semibold mb-1">Telefone</label>
            <input
              v-model="paciente.telecom[0].value"
              type="tel"
              placeholder="(XX) 9XXXX-XXXX"
              title="Utilize apenas números."
              class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 transition"
              @input="formatarTelefone"
              maxlength="15"
            />
          </div>
          
          <div>
            <label class="block text-gray-700 font-semibold mb-1">Email</label>
            <input
              placeholder="Ex: seuemail@dominio.com"
              v-model="paciente.telecom[1].value"
              type="email"
              title="Utilize um email válido"
              class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 transition"
            />
          </div>
        
          <div class="md:col-span-3 flex justify-center mt-4">
            <button
              type="submit"
              class="w-full md:w-auto bg-red-600 hover:bg-red-700 text-white font-bold py-3 px-8 rounded-xl transition-all shadow-lg">
              Cadastrar Paciente
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

