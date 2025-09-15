<script setup>
import { ref } from 'vue'

const family = ref('')
const given1 = ref('')
const given2 = ref('')
const birthDate = ref('')

async function submitForm() {

    await $fetch('/patients', {
        method: 'POST',
        baseURL: 'http://localhost:8090',
        body: {
            family: family.value,
            given1: given1.value,
            given2: given2.value,
            birthDate: birthDate.value
        }
    })

    family.value = ''
    given1.value = ''
    given2.value = ''
    birthDate.value = ''
}
</script>

<template>
  <div>
    <h2 class=" justify-center ">Cadastro de Paciente</h2>
    <input v-model="family" placeholder="Sobrenome" />
    <input v-model="given1" placeholder="Nome 1" />
    <input v-model="given2" placeholder="Nome 2" />
    <input v-model="birthDate" type="date" placeholder="Data de Nascimento" />
    <button @click="submitForm">Enviar</button>
  </div>
</template>
