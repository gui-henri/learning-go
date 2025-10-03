<script setup lang="ts">
import { computed } from 'vue'
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from '@/components/ui/carousel'

// --- LÓGICA UNIFICADA PARA PACIENTES E TAREFAS ---

const config = useRuntimeConfig();

// 1. Busca dos dados de PACIENTES
const { data: patientData, pending: patientsPending, error: patientsError } = await useAsyncData('patientCount', () =>
  $fetch('/Patient', {
    baseURL: config.apiBase ?? "http://localhost:8090"
  })
)

// 2. Busca dos dados de TAREFAS
const { data: taskData, pending: tasksPending, error: tasksError } = await useAsyncData('tasks/all', () =>
  $fetch('/tasks/all', {
    baseURL: config.apiBase ?? "http://localhost:8090"
  })
);

// 3. Propriedades computadas para PACIENTES
const totalPacientes = computed(() => {
  return patientData.value?.data?.entry?.length ?? 0;
})

// 4. Propriedades computadas para TAREFAS
const totalTasks = computed(() => {
  return taskData.value?.tarefas?.length ?? 0;
});

const completedTasksCount = computed(() => {
  if (!taskData.value?.tarefas) return 0;
  return taskData.value.tarefas.filter(t => t.concluida).length;
});

const completionPercentage = computed(() => {
  if (totalTasks.value === 0) return 0;
  const percentage = (completedTasksCount.value / totalTasks.value) * 100;
  return Math.round(percentage);
});
</script>

<template>
  <div class="flex items-start justify-center h-screen">
    <Carousel class="w-150 h-150">
      <CarouselContent>
        
        <CarouselItem>    
          <div class="flex flex-col items-center">
            <img 
              src="https://cdn.bitrix24.com.br/b22035513/landing/750/750d952c46e82fa4ef81b1103248c565/Atendimento_Domiciliar_-_HomeCare_1x.png" 
              class="w-full h-auto rounded-lg"
            > 
            <p class="mt-2 text-center text-red-700 text-sm">
              <b>Atendimento 24H para nossos pacientes</b>
            </p>
            <p class="mt-2 text-justify text-white-700 text-sm">
              O Home Care engloba serviços de saúde oferecidos em domicílio, que exigem um planejamento especial junto ao paciente e às equipes 
              médicas assistentes, em função das necessidades específicas de cada caso. 
              Os critérios da assistência seguem as normas orientadas pela Associação Brasileira de Home Care, além da avaliação específica do paciente. 
            </p>
            <div class="mt-4 p-4 w-full bg-red-50 border-t-4 border-red-500 rounded-b text-red-900 shadow-md" role="alert">
              <div class="flex items-center">
                <div class="py-1">
                  
                </div>
                <div>
                  <p v-if="patientsPending" class="font-bold">Carregando dados dos pacientes...</p>
                  <p v-else-if="patientsError" class="font-bold">Erro ao buscar pacientes</p>
                  <p v-else class="font-bold">
                    Atualmente, contamos com {{ totalPacientes }} pacientes cadastrados!
                  </p>
                  <p class="text-sm">Cuidando de cada um com a máxima atenção.</p>
                </div>
              </div>
            </div>
          </div>
        </CarouselItem>

        <CarouselItem>
          <div class="flex flex-col items-center">
            <img 
              src="https://telemedicinamorsch.com.br/wp-content/uploads/2017/02/iStock-592647828.jpg" 
              class="w-150 h-80 rounded-lg"
            >
            <p class="mt-2 text-center text-blue-700 text-sm">
              <b>Tecnologia em medicina</b>
            </p>
            <p class="mt-2 text-justify text-white-700 text-sm">
              O Home Care pode ser dividido em dois formatos: Internação Domiciliar e Assistência Domiciliar.
              Ambos possuem benefícios ao paciente: Conforto, Humanização e presença constante dos familiares, Atendimento com equipe multidisciplinar e diminuição no risco de infecção hospitalar 
            </p>

            <div class="mt-6 p-4 w-full bg-blue-50 border-t-4 border-blue-500 rounded-b text-blue-900 shadow-md text-left">
              <div>
                <p v-if="tasksPending" class="font-bold text-gray-500">Carregando estatísticas de tarefas...</p>
                <p v-else-if="tasksError" class="font-bold text-red-500">Não foi possível carregar os dados.</p>
                <div v-else>
                  <p class="font-bold">Acompanhamento de Tarefas do Sistema:</p>
                  <p class="text-sm">
                    <b>{{ completionPercentage }}%</b> de progresso 
                    ({{ completedTasksCount }} de {{ totalTasks }} tarefas concluídas).
                  </p>
                </div>
              </div>
            </div>
          </div>
        </CarouselItem>

        <CarouselItem>
          <div class="flex flex-col items-center">
            <img 
              src="https://cdn.prod.website-files.com/647e410aa35e6ce70fed96cd/67449063bc02321b87ce595a_bom-atendimento-medico.jpeg" 
              class="w-150 h-80 rounded-lg"
            >
            <p class="mt-2 text-center text-green-700 text-sm">
              <b>Atendimento humanizado com qualidade</b>
            </p>
            <p class="mt-2 text-center text-white-700 text-sm">
              Serviço de atendimento do paciente para aplicações de medicamentos, terapias entre outras necessidades específicas por período pré-determinado com objetivo de evolução do paciente.
              Esse serviço atende às necessidades pontuais do paciente.</p>
          </div>
        </CarouselItem>

        <CarouselItem>    
          <div class="flex flex-col items-center">
            <img
              src="https://cdn.bitrix24.com.br/b22035513/landing/2a4/2a44b7c2b4e7e248ec783a2eae611962/Logo_Interne_180_x_60_px_1x.png"
              class=""
              title="Serviço de assistência médica domiciliar no Recife, Pernambuco"
            >
            <p class="mt-2 text-sm text-red-700 text-justify"><b>INTERNE SOLUÇÕES EM SAÚDE, A ESCOLHA CERTA PARA HOME CARE</b></p>
            <div class="mt-2 text-sm text-white-700 whitespace-pre-line text-justify">
              <ul class="space-y-3 justify-middle">
                <li><b>O Home Care engloba serviços de saúde oferecidos em domicílio.</b></li>
                <li>Técnicos qualificados.</li>
                <li>Atendimento humanizado.</li>
                <li>Colaboradores especializado.</li>
                <li>Sistemas internos ágeis e seguros.</li>
                <li>Espaço amplo para atendimento.</li>
                <li>Equipamentos eficientes.</li>
                <li>Focada no bem-estar de seus pacientes.</li>
                <li><b>Saiba mais em : </b><a href="https://interne.com.br" target="_blank" class="text-blue-500 underline">
                  Nosso site parceiro.
                </a></li>
              </ul>
            </div>
          </div>
        </CarouselItem> 

      </CarouselContent> 
      <CarouselPrevious />
      <CarouselNext />
    </Carousel>
  </div>
</template>