<script setup>
import { ref } from 'vue';
import { useEvaluationStore } from '@/store/evaluation'

const store = useEvaluationStore();

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

const opcoes_estado_geral = ref([
    { name: 'Bom', code: 'bom' },
    { name: 'Regular', code: 'regular' },
    { name: 'Grave', code: 'grave' },
    { name: 'Gravíssimo', code: 'gravissimo' },
    { name: 'Decaído', code: 'decaido' }
]);


</script>

<template>

    <Fluid >
        <!-- ID's e labels estão criados com a seguinte lógica: 
         tudo em minusculo
         se caso mais de uma palavra, ignora-se (de, do , da etc)
         ignora-se acentos -->


        <!--CARD DADOS PESSOAIS -->

        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">

            <!-- AJUSTAR TODAS AS LINHAS PARA COLUNA EM CASO DE QUEBRA -->

            <div class="flex flex-col gap-4 w-full">

                <!-- TÍTULO DO CARD -->

                <h4 class="font-semibold text-xl">Dados Gerais</h4>

                <!-- LINHA 1 -->

                <div class="flex flex-col md:flex-row gap-4">

                    <div class="flex flex-col gap-2  w-full md:w-1/2">
                        <label for="nome_paciente">Nome completo</label>
                        <InputText id="nome_paciente" v-model="store.evaluation.nome_paciente" type="text" placeholder="Glauberthy júnior"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="genero">Gênero</label>
                        <Select id="genero" v-model="store.evaluation.genero" :options="genero" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>
                     <div class="flex flex-col gap-2 w-full  md:w-1/8">
                        <label for="nascimento">Data de nascimento</label>
                        <Calendar id="nascimento" v-model="store.evaluation.dat_nascimento" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full  md:w-1/9">
                        <label for="cpf">CPF</label>
                        <InputText id="cpf" v-model="store.evaluation.cpf" placeholder="123.456.789-00" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/10 ">
                        <label for="rg">RG</label>
                        <InputText id="rg" v-model="store.evaluation.rg" placeholder="12.123-456" class="w-full"/>
                    </div>
                </div>
                <!-- LINHA 2 -->

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/7">
                        <label for="peso">Peso</label>
                        <InputGroup>
                            <InputText id="peso" v-model="store.evaluation.peso" type="text" maxlength="3" placeholder="64"/>
                            <InputGroupAddon>kg</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/7">
                        <label for="altura">Altura</label>
                        <InputGroup>
                            <InputText id="altura" v-model="store.evaluation.altura" type="text" maxlength="3" placeholder="181"/>
                            <InputGroupAddon>cm</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/7">
                        <label for="age">Idade</label>
                        <InputGroup>
                            <InputText id="idade" v-model="store.evaluation.idade" type="text" maxlength="3" placeholder="32"/>
                            <InputGroupAddon>anos</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="complexidade">Complexidade</label>
                        <Select id="complexidade" v-model="store.evaluation.complexidade" :options="complexidade" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>
                    <div class="flex flex-col gap-2 py-2 w-full"> 
                        <label for="previsao_alta">Previsão de alta?</label>
                        <InputSwitch id="previsao_alta" v-model="store.evaluation.previsao_alta" />
                    </div>
                    

                </div>

                <!-- LINHA 4 -->

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_pai">Nome do pai</label>
                        <InputText id="nome_pai" v-model="store.evaluation.nome_pai" type="text" placeholder="Glauberthy glauberthy"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_mae">Nome da mãe</label>
                        <InputText id="nome_mae" v-model="store.evaluation.nome_mae" type="text" placeholder="Júlia silva"/>
                    </div>


                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="convenio">Convenio médico</label>
                        <InputText id="convenio" v-model="store.evaluation.convenio" type="text" placeholder="Convenio médico xyz"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6" >
                        <label for="hospital">Hospital de internação</label>
                        <InputText id="hospital" v-model="store.evaluation.hospital" type="text" placeholder="Português"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="data_admissao">Data de admissão</label>
                        <Calendar id="data_admissao" v-model="store.evaluation.data_admissao" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6" >
                        <label for="apartamento-hospital">Apartamento hospitalar</label>
                        <InputText id="apartamento_hospital" v-model="store.evaluation.apartamento" type="text" placeholder="Portugês"/>
                    </div>
                    <div class="flex flex-col gap-2 md:w-1/6" >
                        <label for="carteirinha">Nº Carteirinha</label>
                        <InputText id="carteirinha" v-model="store.evaluation.carteirinha" type="text" placeholder="198310293"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="vencimento_carteirinha">Validade da carteirinha</label>
                        <Calendar id="vencimento_carteirinha" v-model="store.evaluation.vencimento_carteirinha" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full"/>
                    </div>
                </div>
            </div>
        </div>
        <!--CARD ENDEREÇO -->

        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">

            <!-- AJUSTAR TODAS AS LINHAS PARA COLUNA EM CASO DE QUEBRA -->

            <div class="flex flex-col gap-4 w-full">

                <!-- TÍTULO DO CARD -->

                <h4 class="font-semibold text-xl">Endereço</h4>

                <!-- LINHA 1 -->
                <div class="flex flex-col md:flex-row gap-4">

                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="cep">CEP</label>
                        <InputText id="cep" v-model="store.evaluation.cep" placeholder="55050-550"class="w-full"/>
                    </div>

                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="estado">Estado</label>
                        <InputText id="estado" v-model="store.evaluation.estado" placeholder="Pernambuco"class="w-full"/>
                    </div>

                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="cidade">Cidade</label>
                        <InputText id="cidade" v-model="store.evaluation.cidade" placeholder="Recife"class="w-full"/>
                    </div>

                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="bairro">Bairro</label>
                        <InputText id="bairro" v-model="store.evaluation.bairro" placeholder="Boa vista"class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="logradoro">Logradoro</label>
                        <InputText id="logradoro" v-model="store.evaluation.logradoro" placeholder="Agamenon magalhães" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="numero-logradoro">Número</label>
                        <InputText id="numero-logradoro" v-model="store.evaluation.numero_logradoro" placeholder="23-A" class="w-full"/>
                    </div>

                </div>
                <!-- LINHA 2 -->
                <div class="flex flex-col md:flex-row gap-4">
                     <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="complemento">Complemento</label>
                        <InputText id="complemento" v-model="store.evaluation.complemento" placeholder="Prédio azul" class="w-full"/>
                    </div>
                     <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="ponto_referencia">Ponto de referência</label>
                        <InputText id="ponto_referencia" v-model="store.evaluation.ponto_referencia" placeholder="perto do posto de saúde" class="w-full"/>
                     </div>
                </div>
            </div>
        </div>
        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">

            <!-- AJUSTAR TODAS AS LINHAS PARA COLUNA EM CASO DE QUEBRA -->

            <div class="flex flex-col gap-4 w-full">
                <!-- TÍTULO DO CARD -->
                <h4 class="font-semibold text-xl">Contato</h4>

                <!-- LINHA 1 -->
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="telefone_residencial">Telefone residencial</label>
                        <InputText id="telefone_residencial" v-model="store.evaluation.telefone_residencial" placeholder="(81) 9 9999-9999"class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="telefone_paciente">Telefone paciente</label>
                        <InputText id="telefone_paciente" v-model="store.evaluation.telefone_paciente" placeholder="(81) 9 9999-9999"class="w-full"/>
                    </div>

                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="email_paciente">E-mail paciente</label>
                        <InputText id="email_paciente" v-model="store.evaluation.email_paciente" placeholder="exemplo@dominio.com" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_responsavel">Nome do responsável</label>
                        <InputText id="nome_responsavel" v-model="store.evaluation.nome_responsavel" placeholder="João albério" class="w-full"/>
                    </div>
                </div>
                <!-- LINHA 2 -->
                <div class="flex flex-col md:flex-row gap-4">
                     <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="parentesco_responsavel">Parentesco do responsável</label>
                        <InputText id="parentesco_responsavel" v-model="store.evaluation.parentesco_responsavel" placeholder="Pai, mãe, irmão etc" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="email_responsavel">E-mail do responsável</label>
                        <InputText id="email_responsavel" v-model="store.evaluation.email_responsavel" placeholder="responsavel@gmail.com" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="telefone_responsavel">Telefone do responsável</label>
                        <InputText id="telefone_responsavel" v-model="store.evaluation.telefone_responsavel" placeholder="(81) 9 9999-9999" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2  md:w-1/6">
                        <label for="forma_contato">Forma de contato</label>
                        <Select id="forma_contato" v-model="store.evaluation.forma_contato" :options="forma_contato" optionLabel="name" placeholder="Selecione forma de contato" class="w-full"></Select>
                    </div>
                    <div class="flex flex-col gap-2 py-2 w-"> 
                        <label for="previsao_alta">Número prioritário</label>
                        <InputSwitch id="previsao_alta" v-model="store.evaluation.previsao_alta" />
                    </div>
                </div>



            </div>
        </div>
    

        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                <h4 class="font-semibold text-xl">Cuidadores</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="medico_solicitante">Médico responsável (solicitante)</label>
                        <InputText id="medico_solicitante" v-model="store.evaluation.medico_solicitante" placeholder="Nome do médico" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="contato_medico">Contato do médico</label>
                        <InputText id="contato_medico" v-model="store.evaluation.contato_medico" placeholder="(XX) 9 9999-9999" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="possui_cuidador">Possui cuidador?</label>
                        <InputSwitch id="possui_cuidador" v-model="store.evaluation.possui_cuidador" optionLabel="name" placeholder="Selecione" class="w-full"></InputSwitch>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_cuidador">Nome do cuidador</label>
                        <InputText id="nome_cuidador" v-model="store.evaluation.nome_cuidador" placeholder="Nome completo" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="contato_cuidador">Contato do cuidador</label>
                        <InputText id="contato_cuidador" v-model="store.evaluation.contato_cuidador" placeholder="(XX) 9 9999-9999" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="turno_cuidador">Turno</label>
                        <Select id="turno_cuidador" v-model="store.evaluation.turno_cuidador" :options="turnos_cuidador" optionLabel="name" placeholder="Selecione o turno" class="w-full"></Select>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="precisa_treinamento">Precisa de treinamento?</label>
                        <InputSwitch id="precisa_treinamento" v-model="store.evaluation.precisa_treinamento" placeholder="Selecione" class="w-full"></InputSwitch>
                    </div>
                    <div class="flex flex-col gap-2 w-full flex-1 md:w-9/10">
                        <label for="obs_treinamento">Observações do treinamento</label>
                        <InputText id="obs_treinamento" v-model="store.evaluation.obs_treinamento" placeholder="Descreva a necessidade..." class="w-full"/>
                    </div>
                </div>

            </div>
        </div>

        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                <h4 class="font-semibold text-xl">Alergias e Precauções</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/10">
                        <label for="alergia_relatada">Alergia relatada?</label>
                        <InputSwitch id="alergia_relatada" v-model="store.evaluation.alergia_relatada" class="w-full"></InputSwitch>
                    </div>
                    
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="tipo_alergia">Tipo de alergia</label>
                        <Select id="tipo_alergia" v-model="store.evaluation.tipo_alergia" :options="tipo_alergia" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>

                    <div class="flex flex-col gap-2 w-full flex-1">
                        <label for="quais_alergias">Quais alergias?</label>
                        <InputText id="quais_alergias" v-model="store.evaluation.quais_alergias" placeholder="Descreva as alergias..." class="w-full"/>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="precaucao">Precaução</label>
                        <Select id="precaucao" v-model="store.evaluation.precaucao" :options="tipo_precaucao" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>

                    <div class="flex flex-col gap-2 w-full flex-1">
                        <label for="cuidados_paliativos">Cuidados Paliativos</label>
                        <InputText id="cuidados_paliativos" v-model="store.evaluation.cuidados_paliativos" placeholder="Descreva os cuidados..." class="w-full"/>
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
                        id="historia_doenca_atual" v-model="store.evaluation.historia_doenca_atual" rows="4" placeholder="Descreva a história clínica mais atual..." class="w-full" autoResize/>
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="anamnese_enfermagem">Anamnese de enfermagem</label>
                    <Textarea id="anamnese_enfermagem" v-model="store.evaluation.anamnese_enfermagem" rows="4" placeholder="Dados sobre o histórico de saúde..." class="w-full" autoResize/>
                </div>
                <div class="flex flex-col gap-2 w-full">
                    <label for="plano_terapeutico">Plano terapêutico</label>
                    <Textarea id="plano_terapeutico" v-model="store.evaluation.plano_terapeutico" rows="4" placeholder="Planejamento do cuidado..." class="w-full" autoResize/> 
                </div>
            </div>
        </div>
        <div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                <h4 class="font-semibold text-xl">Estado Geral e Dispositivos</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="estado_geral">Estado geral</label>
                        <Select id="estado_geral" v-model="store.evaluation.estado_geral" :options="opcoes_estado_geral" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="possui_dreno">Drenos?</label>
                        <InputSwitch id="possui_dreno" v-model="store.evaluation.possui_dreno" class="w-full"></InputSwitch>
                    </div>
                </div>
            </div>
        </div>
    








        

        
    </Fluid>
</template>