<script setup lang="ts">
    import { ref, watch, computed } from "vue"
    import type { DateValue } from "@internationalized/date"
    import {
        DateFormatter,
        getLocalTimeZone,
        today
    } from "@internationalized/date"
    import { CalendarIcon, SendIcon, SearchIcon } from "lucide-vue-next"
    import { Popover, PopoverTrigger, PopoverContent } from "@/components/ui/popover"
    import { Calendar } from "@/components/ui/calendar"
    import { Input } from "@/components/ui/input"
    import { Button } from "@/components/ui/button"

    // Interface para a resposta do backend (mantida por enquanto, mas não usada nesta versão)
    interface InsertResponse {
        tarefa: {
            id: number,
            descricao: string,
            prazo: string,
            concluida: boolean,
            criada_em: string,
            paciente?: string
        }
    }

    // Variáveis reativas para os campos do formulário
    const descricao = ref('');
    const dateValue = ref<DateValue | null>(null)
    const todayDate = today(getLocalTimeZone())

    // Lógica para o popover de paciente
    const buscaPaciente = ref('');
    const pacienteSelecionado = ref<string | null>(null);

    // Lista de pacientes de exemplo (substitua por uma chamada à API no futuro)
    const pacientes = [
      "Samuel Vitor",
      "Guilherme Henrique"
    ];

    // Propriedade computada para filtrar pacientes com base na busca
    const pacientesFiltrados = computed(() => {
      if (!buscaPaciente.value) {
        return pacientes;
      }
      return pacientes.filter(p =>
        p.toLowerCase().includes(buscaPaciente.value.toLowerCase())
      );
    });

    // Função para selecionar um paciente da lista
    const selecionarPaciente = (paciente: string) => {
      pacienteSelecionado.value = paciente;
    };

    // Função de exemplo para simular o envio da tarefa
    function sendTask() {
        console.log("Tarefa enviada (apenas front-end):", {
            descricao: descricao.value,
            prazo: dateValue.value ? dateValue.value.toDate(getLocalTimeZone()).toISOString() : "sem prazo",
            paciente: pacienteSelecionado.value
        });
        // Resetar os campos após a "ação de envio"
        descricao.value = "";
        dateValue.value = null;
        pacienteSelecionado.value = null;
    }

</script>

<template>
    <div class="items-center justify-center gap-2">
        <div class="p-0.5 rounded-full border-2 border-red-500 flex items-center">
            <Input v-model="descricao" class="p-3 border-0 focus-visible:border-0 focus-visible:ring-ring/0 focus-visible:ring-[3px]" type="text" name="description" placeholder="Digite a tarefa" />

            <Popover>
                <PopoverTrigger as-child>
                    <Button variant="outline"
                        class="w-34 bg-amber-200 hover:bg-red-400 transition justify-start text-left font-normal border-0 rounded-l-full"
                    >
                        <CalendarIcon />
                        {{ dateValue ? new DateFormatter("pt-BR", {}).format(dateValue.toDate(getLocalTimeZone())) : "Definir prazo" }}
                    </Button>
                </PopoverTrigger>
                <PopoverContent class="w-auto p-0">
                    <Calendar locale="pt-BR" :minValue="today(getLocalTimeZone())" class="flex flex-col" v-model="dateValue" initial-focus />
                </PopoverContent>
            </Popover>

            <Popover>
              <PopoverTrigger as-child>
                <Button
                  variant="outline"
                  class="w-36 bg-blue-300 hover:bg-red-400 transition justify-start text-left font-normal border-0 rounded-r-full contain-content "
                >
                  <SearchIcon class=" h-4 w-4" />
                  {{ pacienteSelecionado ? pacienteSelecionado.split(" ")[0] : "Definir paciente" }}
                </Button>
              </PopoverTrigger>

              <PopoverContent class="w-72 p-4">
                <div class="flex items-center border rounded-full px-3 py-2 bg-white shadow-sm">
                  <SearchIcon class="h-4 w-4 text-gray-500 mr-2" />
                  <input
                    v-model="buscaPaciente"
                    type="text"
                    placeholder="Digite o nome do paciente..."
                    class="flex-1 outline-none text-sm bg-transparent"
                  />
                </div>

                <!-- Resultados da pesquisa -->
                <ul v-if="pacientesFiltrados.length" class="mt-2 max-h-48 overflow-y-auto">
                  <li
                    v-for="(paciente, i) in pacientesFiltrados"
                    :key="i"
                    class="px-3 py-2 hover:bg-gray-100 rounded cursor-pointer"
                    @click="selecionarPaciente(paciente)"
                  >
                    {{ paciente }}
                  </li>
                </ul>
                <p v-else class="text-sm text-gray-400 mt-2">Nenhum paciente encontrado.</p>
              </PopoverContent>
            </Popover>
        </div>
        <Button @click="sendTask" class="bg-red-600 hover:bg-red-500 p-5 rounded-2xl cursor-pointer"><SendIcon /></Button>
    </div>
</template>

<style lang="css" scoped>
    div {
        display: flex;
        width: 100%;
        max-width: 640px;
    }
    input {
        width: 100%;
    }

    :deep([data-selected]) {
        background-color: #dc2626; /* bg-red-600 */
        color: white;
        font-weight: 700;
    }

    :deep([data-selected]:hover) {
        background-color: #b91c1c; /* bg-red-700 */
    }
</style>