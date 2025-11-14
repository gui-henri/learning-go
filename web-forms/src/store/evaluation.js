import { defineStore } from 'pinia'
import { ref } from 'vue';

export const useEvaluationStore = defineStore('evaluation', () => {
    const evaluation = ref({
        genero: null,
        complexidade: null,
        peso: null,
        altura: null,
        nome_paciente: null,
        idade: null,
        previsao_alta: null
    })

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

    return {
        evaluation,
        setGenero,
        setAltura,
        setComplexidade,
        setIdade,
        setNomePaciente,
        setPeso,
        setPrevisaoAlta
    }
}, {
    persist: true
})