
<script setup>
import { ref } from 'vue';
import DadosGerais from '@/components/Evaluation/DadosGerais.vue';
import { useDadosGeraisStore } from '@/store/evaluation/dadosGerais';
import { useEnderecoStore } from "@/store/evaluation/endereco";
import { useContatoStore } from '@/store/evaluation/contato';
import { useCuidadoresStore } from '@/store/evaluation/cuidadores';
import { useSegurancaStore } from '@/store/evaluation/seguranca';
import { useHistoricoStore } from '@/store/evaluation/historico';
import { useExameFisicoStore } from '@/store/evaluation/exameFisico';
import Cardiorrespiratório from '@/components/Evaluation/Cardiorrespiratório.vue';

const dadosGeraisStore = useDadosGeraisStore();
const enderecoStore = useEnderecoStore();
const contatoStore = useContatoStore();
const cuidadorStore = useCuidadoresStore();
const segurancaStore = useSegurancaStore();
const historicoStore = useHistoricoStore();
const exameFisicoStore = useExameFisicoStore();



const forma_contato = ref([
    { name: 'Residencial', code: 'residencial' },
    { name: 'Celular', code: 'celular' },
    { name: 'Whatsapp', code: 'whatsapp' },
    { name: 'Celular e Whatsapp', code: 'celular_whatsapp' }
]);

const turnos_cuidador = ref([
    { name: 'Diurno', code: 'diurno' },
    { name: 'Noturno', code: 'noturno' },
    { name: '24h', code: '24h' }
]);

const tipo_alergia = ref([
    { name: 'Medicamento', code: 'medicamento' },
    { name: 'Alimento', code: 'alimento' }
]);

const tipo_precaucao = ref([
    { name: 'Padrão', code: 'padrao' },
    { name: 'Contato', code: 'contato' },
    { name: 'Gotícula', code: 'goticula' },
    { name: 'Aerossóis', code: 'aerossois' },
    { name: 'Reverso', code: 'reverso' }
]);

const estado_geral = ref([
    { name: 'Bom', code: 'bom' },
    { name: 'Regular', code: 'regular' },
    { name: 'Grave', code: 'grave' },
    { name: 'Gravíssimo', code: 'gravissimo' },
    { name: 'Decaído', code: 'decaido' }
]);

const locomocao = ref([
    { name: 'Deambula sem apoio', code: 'deambula_sem_apoio' },
    { name: 'Deambula com apoio', code: 'deambula_com_apoio' },
    { name: 'Restrito ao leito', code: 'restrito_leito' },
    { name: 'Cadeirante', code: 'cadeirante' }
]);

const piccline = ref([
    { name: 'Monolumen', code: 'monolumen' },
    { name: 'Duplolumen', code: 'duplolumen' },
    { name: 'Triplolumen', code: 'triplolumen' }
]);

</script>

<template>
    <Fluid>

        <DadosGerais/>

        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Endereço</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="cep">CEP</label>
                        <InputText 
                            id="cep" 
                            v-model="enderecoStore.endereco.cep" 
                            placeholder="55050-550" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="estado">Estado</label>
                        <InputText 
                            id="estado" 
                            v-model="enderecoStore.endereco.estado" 
                            placeholder="Pernambuco" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="cidade">Cidade</label>
                        <InputText 
                            id="cidade" 
                            v-model="enderecoStore.endereco.cidade" 
                            placeholder="Recife" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="bairro">Bairro</label>
                        <InputText 
                            id="bairro" 
                            v-model="enderecoStore.endereco.bairro" 
                            placeholder="Boa vista" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="logradoro">Logradoro</label>
                        <InputText 
                            id="logradoro" 
                            v-model="enderecoStore.endereco.logradoro" 
                            placeholder="Agamenon magalhães" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="numero-logradoro">Número</label>
                        <InputText 
                            id="numero-logradoro" 
                            v-model="enderecoStore.endereco.numero" 
                            placeholder="23-A" 
                            class="w-full"
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="complemento">Complemento</label>
                        <InputText 
                            id="complemento" 
                            v-model="enderecoStore.endereco.complemento" 
                            placeholder="Prédio azul" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="ponto_referencia">Ponto de referência</label>
                        <InputText 
                            id="ponto_referencia" 
                            v-model="enderecoStore.endereco.ponto_referencia" 
                            placeholder="perto do posto de saúde" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>
        </div>
        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Contato</h4>
                <p class="text-red-600">TODO: adicionar mais de um responsável</p>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="telefone_residencial">Telefone residencial</label>
                        <InputText 
                            id="telefone_residencial" 
                            v-model="contatoStore.contato.telefone_residencial" 
                            placeholder="(81) 9 9999-9999" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="telefone_paciente">Telefone paciente</label>
                        <InputText 
                            id="telefone_paciente" 
                            v-model="contatoStore.contato.telefone_paciente" 
                            placeholder="(81) 9 9999-9999" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="email_paciente">E-mail paciente</label>
                        <InputText 
                            id="email_paciente" 
                            v-model="contatoStore.contato.email_paciente" 
                            placeholder="exemplo@dominio.com" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_responsavel">Nome do responsável</label>
                        <InputText 
                            id="nome_responsavel" 
                            v-model="contatoStore.contato.nome_responsavel" 
                            placeholder="João albério" 
                            class="w-full"
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="parentesco_responsavel">Parentesco do responsável</label>
                        <InputText 
                            id="parentesco_responsavel" 
                            v-model="contatoStore.contato.parentesco_responsavel" 
                            placeholder="Pai, mãe, irmão etc" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="email_responsavel">E-mail do responsável</label>
                        <InputText 
                            id="email_responsavel" 
                            v-model="contatoStore.contato.email_responsavel" 
                            placeholder="responsavel@gmail.com" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="telefone_responsavel">Telefone do responsável</label>
                        <InputText 
                            id="telefone_responsavel" 
                            v-model="contatoStore.contato.telefone_responsavel" 
                            placeholder="(81) 9 9999-9999" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 md:w-1/6">
                        <label for="forma_contato">Forma de contato</label>
                        <Select 
                            id="forma_contato" 
                            v-model="contatoStore.contato.forma_contato" 
                            :options="forma_contato" 
                            optionLabel="name" 
                            placeholder="Selecione forma de contato" 
                            class="w-full"
                        ></Select>
                    </div>
                    <div class="flex flex-col gap-2 py-2 w-">
                        <label for="previsao_alta">Número prioritário</label>
                        <InputSwitch 
                            id="previsao_alta" 
                            v-model="contatoStore.contato.numero_prioritario" 
                        />
                    </div>
                </div>
            </div>
        </div>
                <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Cuidadores</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="medico_solicitante">Médico solicitante</label>
                        <InputText 
                            id="medico_solicitante" 
                            v-model="cuidadorStore.cuidadores.medico_solicitante" 
                            placeholder="Nome do médico" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="contato_medico">Contato do médico</label>
                        <InputText 
                            id="contato_medico" 
                            v-model="cuidadorStore.cuidadores.contato_medico" 
                            placeholder="(XX) 9 9999-9999" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="possui_cuidador">Possui cuidador?</label>
                        <InputSwitch 
                            id="possui_cuidador" 
                            v-model="cuidadorStore.cuidadores.possui_cuidador" 
                            class="w-full"
                        ></InputSwitch>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_cuidador">Nome do cuidador</label>
                        <InputText 
                            id="nome_cuidador" 
                            v-model="cuidadorStore.cuidadores.nome_cuidador" 
                            placeholder="Nome completo" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="contato_cuidador">Contato do cuidador</label>
                        <InputText 
                            id="contato_cuidador" 
                            v-model="cuidadorStore.cuidadores.contato_cuidador" 
                            placeholder="(XX) 9 9999-9999" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="turno_cuidador">Turno</label>
                        <Select 
                            id="turno_cuidador" 
                            v-model="cuidadorStore.cuidadores.turno_cuidador" 
                            :options="turnos_cuidador" 
                            optionLabel="name" 
                            placeholder="Selecione o turno" 
                            class="w-full"
                        ></Select>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="precisa_treinamento">Precisa de treinamento?</label>
                        <InputSwitch 
                            id="precisa_treinamento" 
                            v-model="cuidadorStore.cuidadores.precisa_treinamento" 
                            class="w-full"
                        ></InputSwitch>
                    </div>
                    <div class="flex flex-col gap-2 w-full flex-1 md:w-9/10">
                        <label for="obs_treinamento">Observações do treinamento</label>
                        <InputText 
                            id="obs_treinamento" 
                            v-model="cuidadorStore.cuidadores.obs_treinamento" 
                            placeholder="Descreva a necessidade..." 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>
        </div>
                <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Segurança do paciente</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="alergia_relatada">Alergia relatada?</label>
                        <InputSwitch 
                            id="alergia_relatada" 
                            v-model="segurancaStore.seguranca.alergia_relatada" 
                            class="w-full"
                        ></InputSwitch>
                    </div>
                    
                    <template v-if="segurancaStore.seguranca.alergia_relatada">
                        <div class="flex flex-col gap-2 w-full md:w-1/6">
                            <label for="tipo_alergia">Tipo de alergia</label>
                            <Select 
                                id="tipo_alergia" 
                                v-model="segurancaStore.seguranca.tipo_alergia" 
                                :options="tipo_alergia" 
                                optionLabel="name" 
                                placeholder="Selecione" 
                                class="w-full"
                            ></Select>
                        </div>

                        <div class="flex flex-col gap-2 w-full flex-1">
                            <label for="quais_alergias">Quais alergias?</label>
                            <InputText 
                                id="quais_alergias" 
                                v-model="segurancaStore.seguranca.quais_alergias" 
                                placeholder="Descreva as alergias..." 
                                class="w-full"
                            />
                        </div>
                    </template>
                </div>

                <div v-if="segurancaStore.seguranca.alergia_relatada" class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="precaucao">Precaução</label>
                        <Select 
                            id="precaucao" 
                            v-model="segurancaStore.seguranca.precaucao" 
                            :options="tipo_precaucao" 
                            optionLabel="name" 
                            placeholder="Selecione" 
                            class="w-full"
                        ></Select>
                    </div>

                    <div class="flex flex-col gap-2 w-full flex-1">
                        <label for="cuidados_paliativos">Cuidados Paliativos</label>
                        <InputText 
                            id="cuidados_paliativos" 
                            v-model="segurancaStore.seguranca.cuidados_paliativos" 
                            placeholder="Descreva os cuidados..." 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>
            
        </div>
        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Histórico e Planejamento</h4>
                
                <div class="flex flex-col gap-2 w-full">
                    <label for="historia_doenca_atual">História da doença atual</label>
                    <Textarea 
                        id="historia_doenca_atual" 
                        v-model="historicoStore.historico.historia_doenca_atual" 
                        rows="4" 
                        placeholder="Descreva a história clínica mais atual..." 
                        class="w-full" 
                        autoResize
                    />
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="anamnese_enfermagem">Anamnese de enfermagem</label>
                    <Textarea 
                        id="anamnese_enfermagem" 
                        v-model="historicoStore.historico.anamnese_enfermagem" 
                        rows="4" 
                        placeholder="Dados sobre o histórico de saúde..." 
                        class="w-full" 
                        autoResize
                    />
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="plano_terapeutico">Plano terapêutico</label>
                    <Textarea 
                        id="plano_terapeutico" 
                        v-model="historicoStore.historico.plano_terapeutico" 
                        rows="4" 
                        placeholder="Planejamento do cuidado..." 
                        class="w-full" 
                        autoResize
                    /> 
                </div>
            </div>
        </div>
        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Exame Físico</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="estado_geral">Estado geral</label>
                        <Select 
                            id="estado_geral" 
                            v-model="exameFisicoStore.exameFisico.estado_geral" 
                            :options="estado_geral" 
                            optionLabel="name" 
                            placeholder="Selecione" 
                            class="w-full"
                        ></Select>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="avaliacao_locomotora">Avaliação locomotora</label>
                        <Select 
                            id="avaliacao_locomotora" 
                            v-model="exameFisicoStore.exameFisico.avaliacao_locomotora" 
                            :options="locomocao" 
                            optionLabel="name" 
                            placeholder="Selecione" 
                            class="w-full"
                        ></Select>
                    </div>
                </div>

                <div class="flex flex-col gap-4 border-t pt-4">
                    <div class="flex flex-col md:flex-row gap-4 items-center">
                        <div class="flex flex-col gap-2 w-full md:w-1/6">
                            <label for="possui_dreno">Drenos?</label>
                            <InputSwitch 
                                id="possui_dreno" 
                                v-model="exameFisicoStore.exameFisico.possui_dreno" 
                            />
                        </div>
                    </div>

                    <div v-if="exameFisicoStore.exameFisico.possui_dreno" class="flex flex-col md:flex-row gap-4 p-4 rounded-lg">
                        <div class="flex flex-col gap-2 w-full md:w-1/4">
                            <label for="local_dreno">Local do dreno</label>
                            <InputText 
                                id="local_dreno" 
                                v-model="exameFisicoStore.exameFisico.local_dreno" 
                                placeholder="Ex: Braço direito" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full md:w-1/6">
                            <label for="data_implantacao_dreno">Data implantação</label>
                            <Calendar 
                                id="data_implantacao_dreno" 
                                v-model="exameFisicoStore.exameFisico.data_implantacao_dreno" 
                                dateFormat="dd/mm/yy" 
                                placeholder="dd/mm/aaaa" 
                                class="w-full"
                            />
                        </div>
                        <div class="flex flex-col gap-2 w-full flex-1">
                            <label for="curativos_dreno">Curativos do dreno</label>
                            <InputText 
                                id="curativos_dreno" 
                                v-model="exameFisicoStore.exameFisico.curativos_dreno" 
                                placeholder="Descrição dos curativos" 
                                class="w-full"
                            />
                        </div>
                    </div>
                </div>
                <div class="flex flex-col gap-4 border-t pt-4">
                
                        <div class="flex flex-col gap-2 w-full md:w-1/6">
                            <label for="possui_acessos">Acessos?</label>
                            <InputSwitch 
                                id="possui_acesso" 
                                v-model="exameFisicoStore.exameFisico.possui_acesso" 
                            />
                        </div>
                        </div>
                        <div v-if="exameFisicoStore.exameFisico.possui_dreno" class="flex flex-col md:flex-row gap-4 p-4 rounded-lg">
                            <div class="flex flex-col gap-2 w-full md:w-1/2">
                                <label for="estado_geral">Estado geral</label>
                                <Select 
                                id="estado_geral" 
                                v-model="exameFisicoStore.exameFisico.estado_geral" 
                                :options="estado_geral" 
                                optionLabel="name" 
                                placeholder="Selecione" 
                                class="w-full"
                            ></Select>
                            </div>    
                </div>
            </div>
        </div>
        <Cardiorrespiratório />
    </Fluid>
</template>
```