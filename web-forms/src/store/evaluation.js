import { defineStore } from 'pinia'
import { ref } from 'vue';

export const evaluationModel = {
        genero: null,
        complexidade: null,
        peso: null,
        altura: null,
        nome_paciente: null,
        idade: null,
        previsao_alta: null,
        dat_nascimento: null,
        cpf: null,
        rg: null,
        nome_pai: null,
        nome_mae: null,
        convenio: null,
        hospital: null,
        dat_admissao: null,
        apartamento: null,
        carteirinha: null,
        vencimento_carteirinha: null,
    }

export const useEvaluationStore = defineStore('evaluation', () => {
    const evaluation = ref(evaluationModel)

    function setGenero(genero) {
        evaluation.value.genero = genero
    }
    
    function setComplexidade(complexidade) {
        evaluation.value.complexidade = complexidade
    }

    function setPeso(peso) {
        evaluation.value.peso = peso
    }

    function setAltura(altura) {
        evaluation.value.altura = altura
    }

    function setNomePaciente(nomePaciente) {
        evaluation.value.nome_paciente = nomePaciente
    }

    function setIdade(idade) {
        evaluation.value.idade = idade
    }

    function setPrevisaoAlta(previsaoAlta) {
        evaluation.value.previsao_alta = previsaoAlta
    }

    function setDatNascimento(dat_nascimento) {
        evaluation.value.dat_nascimento = dat_nascimento
    }

    function setCpf(cpf) {
        evaluation.value.cpf = cpf
    }

    function setRg(rg) {
        evaluation.value.rg = rg
    }

    function setNomePai(nome_pai) {
        evaluation.value.nome_pai = nome_pai
    }

    function setNomeMae(nome_mae) {
        evaluation.value.nome_mae = nome_mae
    }

    function setConvenio(convenio) {
        evaluation.value.convenio = convenio
    }

    function setHospital(hospital) {
        evaluation.value.hospital = hospital
    }

    function setDatAdmissao(dat_admissao) {
        evaluation.value.dat_admissao = dat_admissao
    }

    function setApartamento(apartamento) {
        evaluation.value.apartamento = apartamento
    }

    function setCarteirinha(carteirinha) {
        evaluation.value.carteirinha = carteirinha
    }

    function setVencimentoCarteirinha(vencimento_carteirinha) {
        evaluation.value.vencimento_carteirinha = vencimento_carteirinha
    }


    return {
        evaluation,
        setGenero,
        setAltura,
        setComplexidade,
        setIdade,
        setNomePaciente,
        setPeso,
        setPrevisaoAlta,
        setDatNascimento,
        setCpf,
        setRg,
        setNomePai,
        setNomeMae,
        setConvenio,
        setHospital,
        setDatAdmissao,
        setApartamento,
        setCarteirinha,
        setVencimentoCarteirinha
    }
}, {
    persist: true
})