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

                    <div class="flex flex-col gap-2 w-full">
                        <label for="nome_paciente">Nome completo</label>
                        <InputText id="nome_paciente" v-model="store.evaluation.nome_paciente" type="text" placeholder="Glauberthy júnior"/>
                    </div>
                </div>

                <!-- LINHA 2 -->

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2  md:w-1/5">
                        <label for="genero">Gênero</label>
                        <Select id="genero" v-model="store.evaluation.genero" :options="genero" optionLabel="name" placeholder="Selecione" class="w-full"></Select>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/5">
                        <label for="nascimento">Data de nascimento</label>
                        <Calendar id="nascimento" v-model="store.evaluation.dat_nascimento" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/5">
                        <label for="cpf">CPF</label>
                        <InputText id="cpf" v-model="store.evaluation.cpf" placeholder="123.456.789-00"class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/5">
                        <label for="rg">RG</label>
                        <InputText id="rg" v-model="store.evaluation.rg" placeholder="12.123-456" class="w-full"/>
                    </div>
                </div>

                <!-- LINHA 3 -->

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="peso">Peso</label>
                        <InputGroup>
                            <InputText id="peso" v-model="store.evaluation.peso" type="text" maxlength="3" placeholder="64"/>
                            <InputGroupAddon>kg</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="altura">Altura</label>
                        <InputGroup>
                            <InputText id="altura" v-model="store.evaluation.altura" type="text" maxlength="3" placeholder="181"/>
                            <InputGroupAddon>cm</InputGroupAddon>
                        </InputGroup>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
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
                    <div class="flex flex-col gap-2 py-2 w-full m"> 
                        <label for="previsao_alta">Previsão de alta?</label>
                        <InputSwitch id="previsao_alta" v-model="store.evaluation.previsao_alta" />
                    </div>
                </div>

                <!-- LINHA 4 -->

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full">
                        <label for="nome_pai">Nome do pai</label>
                        <InputText id="nome_pai" v-model="store.evaluation.nome_pai" type="text" placeholder="Glauberthy glauberthy"/>
                    </div>
                </div>

                <!-- LINHA 5 -->

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full">
                        <label for="nome_mae">Nome da mãe</label>
                        <InputText id="nome_mae" v-model="store.evaluation.nome_mae" type="text" placeholder="Júlia silva"/>
                    </div>
                </div>

                <!-- LINHA 6 -->

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="convenio">Convenio médico</label>
                        <InputText id="convenio" v-model="store.evaluation.convenio" type="text" placeholder="Convenio médico xyz"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/2" >
                        <label for="hospital">Hospital de internação</label>
                        <InputText id="hospital" v-model="store.evaluation.hospital" type="text" placeholder="Portugês"/>
                    </div>
                </div>

                <!-- LINHA 7 -->

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/5">
                        <label for="data_admissao">Data de admissão</label>
                        <Calendar id="data_admissao" v-model="store.evaluation.data_admissao" dateFormat="dd/mm/yy" placeholder="dd/mm/aaaa" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/5" >
                        <label for="apartamento-hospital">Apartamento hospitalar</label>
                        <InputText id="apartamento_hospital" v-model="store.evaluation.apartamento" type="text" placeholder="Portugês"/>
                    </div>
                    <div class="flex flex-col gap-2 md:w-1/5" >
                        <label for="carteirinha">Nº Carteirinha</label>
                        <InputText id="carteirinha" v-model="store.evaluation.carteirinha" type="text" placeholder="198310293"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/5">
                        <label for="vencimento_carteirinha">Data de validade da Carteirinha</label>
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

                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="cep">CEP</label>
                        <InputText id="cep" v-model="store.evaluation.cep" placeholder="55050-550"class="w-full"/>
                    </div>

                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="estado">Estado</label>
                        <InputText id="estado" v-model="store.evaluation.estado" placeholder="Pernambuco"class="w-full"/>
                    </div>

                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="cidade">Cidade</label>
                        <InputText id="cidade" v-model="store.evaluation.cidade" placeholder="Recife"class="w-full"/>
                    </div>

                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="bairro">Bairro</label>
                        <InputText id="bairro" v-model="store.evaluation.bairro" placeholder="Boa vista"class="w-full"/>
                    </div>

                </div>
                <div class="flex flex-col md:flex-row gap-4">

                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="logradoro">Logradoro</label>
                        <InputText id="logradoro" v-model="store.evaluation.logradoro" placeholder="Agamenon magalhães" class="w-full"/>
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="numero-logradoro">Número</label>
                        <InputText id="numero-logradoro" v-model="store.evaluation.numero_logradoro" placeholder="23-A" class="w-full"/>
                    </div>
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

    </Fluid>
</template>