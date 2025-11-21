
<script setup>
import { ref } from 'vue';
import { useDadosGeraisStore } from '@/store/evaluation/dadosGerais';
import { useEnderecoStore } from "@/store/evaluation/endereco";

const dadosGeraisStore = useDadosGeraisStore();
const enderecoStore = useEnderecoStore();

const genero = ref([
    { name: 'Feminino', code: 'feminino' },
    { name: 'Masculino', code: 'masculino' }
]);

const complexidade = ref([
    { name: 'Baixa', code: 'baixa' },
    { name: 'Média', code: 'media' },
    { name: 'Alta', code: 'alta' }
]);

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
        
        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Dados Gerais</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_paciente">Nome completo</label>
                        <InputText 
                            id="nome_paciente" 
                            v-model="dadosGeraisStore.dadosGerais.nome_paciente" 
                            type="text" 
                            placeholder="Glauberthy júnior"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="genero">Gênero</label>
                        <Select 
                            id="genero" 
                            v-model="dadosGeraisStore.dadosGerais.genero" 
                            :options="genero" 
                            optionLabel="name" 
                            placeholder="Selecione" 
                            class="w-full"
                        ></Select>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="nascimento">Data de nascimento</label>
                        <Calendar 
                            id="nascimento" 
                            v-model="dadosGeraisStore.dadosGerais.data_nascimento" 
                            dateFormat="dd/mm/yy" 
                            placeholder="dd/mm/aaaa" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/9">
                        <label for="cpf">CPF</label>
                        <InputText 
                            id="cpf" 
                            v-model="dadosGeraisStore.dadosGerais.cpf" 
                            placeholder="123.456.789-00" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="rg">RG</label>
                        <InputText 
                            id="rg" 
                            v-model="dadosGeraisStore.dadosGerais.rg" 
                            placeholder="12.123-456" 
                            class="w-full"
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/7">
                        <label for="peso">Peso</label>
                        <InputGroup>
                            <InputText 
                                id="peso" 
                                v-model="dadosGeraisStore.dadosGerais.peso" 
                                type="text" 
                                maxlength="3" 
                                placeholder="64"
                            />
                            <InputGroupAddon>kg</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/7">
                        <label for="altura">Altura</label>
                        <InputGroup>
                            <InputText 
                                id="altura" 
                                v-model="dadosGeraisStore.dadosGerais.altura" 
                                type="text" 
                                maxlength="3" 
                                placeholder="181"
                            />
                            <InputGroupAddon>cm</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/7">
                        <label for="age">Idade</label>
                        <InputGroup>
                            <InputText 
                                id="idade" 
                                v-model="dadosGeraisStore.dadosGerais.idade" 
                                type="text" 
                                maxlength="3" 
                                placeholder="32"
                            />
                            <InputGroupAddon>anos</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="complexidade">Complexidade</label>
                        <Select 
                            id="complexidade" 
                            v-model="dadosGeraisStore.dadosGerais.complexidade" 
                            :options="complexidade" 
                            optionLabel="name" 
                            placeholder="Selecione" 
                            class="w-full"
                        ></Select>
                    </div>
                    <div class="flex flex-col gap-2 py-2 w-full">
                        <label for="previsao_alta">Previsão de alta?</label>
                        <InputSwitch 
                            id="previsao_alta" 
                            v-model="dadosGeraisStore.dadosGerais.previsao_alta" 
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_pai">Nome do pai</label>
                        <InputText 
                            id="nome_pai" 
                            v-model="dadosGeraisStore.dadosGerais.nome_pai" 
                            type="text" 
                            placeholder="Glauberthy glauberthy"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_mae">Nome da mãe</label>
                        <InputText 
                            id="nome_mae" 
                            v-model="dadosGeraisStore.dadosGerais.nome_mae" 
                            type="text" 
                            placeholder="Júlia silva"
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="convenio">Convenio médico</label>
                        <InputText 
                            id="convenio" 
                            v-model="dadosGeraisStore.dadosGerais.convenio" 
                            type="text" 
                            placeholder="Convenio médico xyz"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="hospital">Hospital de internação</label>
                        <InputText 
                            id="hospital" 
                            v-model="dadosGeraisStore.dadosGerais.hospital" 
                            type="text" 
                            placeholder="Português"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="data_admissao">Data de admissão</label>
                        <Calendar 
                            id="data_admissao" 
                            v-model="dadosGeraisStore.dadosGerais.data_admissao" 
                            dateFormat="dd/mm/yy" 
                            placeholder="dd/mm/aaaa" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="apartamento-hospital">Apartamento hospitalar</label>
                        <InputText 
                            id="apartamento_hospital" 
                            v-model="dadosGeraisStore.dadosGerais.apartamento" 
                            type="text" 
                            placeholder="Portugês"
                        />
                    </div>
                    <div class="flex flex-col gap-2 md:w-1/6">
                        <label for="carteirinha">Nº Carteirinha</label>
                        <InputText 
                            id="carteirinha" 
                            v-model="dadosGeraisStore.dadosGerais.carteirinha" 
                            type="text" 
                            placeholder="198310293"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="vencimento_carteirinha">Validade da carteirinha</label>
                        <Calendar 
                            id="vencimento_carteirinha" 
                            v-model="dadosGeraisStore.dadosGerais.vencimento_carteirinha" 
                            dateFormat="dd/mm/yy" 
                            placeholder="dd/mm/aaaa" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>
        </div>
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
                            v-model="enderecoStore.endereco.estado_geral" 
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
                            v-model="enderecoStore.endereco.numero_logradoro" 
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
    </Fluid>
</template>
```