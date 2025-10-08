<script setup>
import { computed } from 'vue' // Adicionado
import { Ellipsis, Eye, Trash2, BookCheck, Plus, CrossIcon, MailWarningIcon } from "lucide-vue-next"
import Popover from '~/components/ui/popover/Popover.vue';
import PopoverTrigger from '~/components/ui/popover/PopoverTrigger.vue';
import PopoverContent from '~/components/ui/popover/PopoverContent.vue';
const config = useRuntimeConfig();

const { data, pending, error, refresh } = await useAsyncData('Patient/all', () =>
  $fetch('/Patient', {
    baseURL: config.apiBase ?? "http://localhost:8090"
  })
)

// Lógica de contagem adicionada
const totalPacientes = computed(() => {
  if (!data.value?.data?.entry) {
    return 0;
  }
  const activePatients = data.value.data.entry.filter(p => p.resource.active);
  return activePatients.length;
});

async function deactivatePatient(id, patient) {
  await $fetch(`/Patient/${id}`, {
    method: 'PUT',
    baseURL: config.apiBase ?? "http://localhost:8090",
    body: {
      ...patient,
      active: false
    }
  })

  refresh()
}

function viewPatientDetails(id) {
  // navigateTo is a Nuxt composable for programmatic routing
  navigateTo(`/paciente/${id}`);
}
</script>
<template>
  <div class="w-full flex justify-center mt-4">
  </div>
  <div class="w-full flex flex-col items-center">
    <div class="mt-2 flex flex-col items-end gap-6">
      <div class="w-full flex justify-between">
        <h1 class="font-bold text-3xl">{{ totalPacientes }} Paciente(s) em nosso cuidado.</h1>
        <NuxtLink to="/dashboard" >
          <Button class="bg-red-600 hover:bg-red-700"><BookCheck></BookCheck> Entender Meus Pacientes</Button>
        </NuxtLink>
        <NuxtLink to="/cadastro" >
          <Button class="bg-red-600 hover:bg-red-700"><Plus />Novo Paciente</Button>
        </NuxtLink>
       
      </div>

      
      <div v-if="pending" class="text-center text-gray-500">
          <p>Carregando pacientes...</p>
        </div>
      <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg text-center" role="alert">
        <strong>Ocorreu um erro:</strong>
        <p>{{ error.message }}</p>
      </div>
      <table v-else class="w-[1024px] pl-4 pr-4 pt-3 pb-3 border-2 table-auto border-t-red-400 border-l-red-400 border-e-red-600 border-b-red-600 rounded-2xl border-separate">
        <thead>
          <tr>
            <th class="text-start border-b-2 border-red-300"><Button class="bg-white text-black hover:bg-red-300">Nome</Button></th>
            <th class="text-start border-b-2 border-red-300"><Button class="bg-white text-black hover:bg-red-300">CPF</Button></th>
            <th class="text-start border-b-2 border-red-300"><Button class="bg-white text-black hover:bg-red-300">Data de Nascimento</Button></th>
            <th class="text-start border-b-2 border-red-300"><Button class="bg-white text-black hover:bg-red-300">Gênero</Button></th>
          </tr>
        </thead>
        <tbody>
          <tr class="border-2 border-red-600" v-for="(patient, index) in data.data.entry" :key="index">
            <td class="p-2">{{ (patient.resource.name.find(n => n.use === "official" || n.use === "usual")?.given || ["Sem nome"]).join(" ") + patient.resource.name.find(n => n.use === "official" || n.use === "usual")?.family }}</td>
            <td class="p-2">{{ (patient.resource.identifier.find(t => t.system ==="http://hl7.org.br/fhir/r4/sid/CPF")?.value) || "Sem CPF" }}</td>
            <td class="p-2">{{ (patient.resource.birthDate) || "Sem Data" }}</td>
            <td class="p-2">{{ (patient.resource.gender === "male" ? "Masculino" : patient.resource.gender === "female" ? "Feminino" : "Outro") }}</td>
            <td class="flex gap-1 justify-center items-center">
                <Popover>
                  <PopoverTrigger>
                    <Button class="bg-red-300 text-red-800 hover:bg-red-400"><Ellipsis /></Button>
                  </PopoverTrigger>
                  <PopoverContent class="flex flex-col gap-2 w-36 border-2 border-red-300">
                    <Button @click="() => viewPatientDetails(patient.resource.id)" class="w-28 flex justify-start bg-white text-gray-700 hover:bg-red-300"><Eye />Detalhes</Button>
                    <Button @click="() => deactivatePatient(patient.resource.id, patient.resource)" class="w-28 flex justify-start bg-white text-red-700 hover:bg-red-300"><Trash2 />Deletar</Button>
                  </PopoverContent>
                </Popover>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
  
</template>
