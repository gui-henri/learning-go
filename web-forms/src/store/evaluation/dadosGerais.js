import { defineStore } from 'pinia'
import { ref } from 'vue';

export const dadosGeraisModel = {
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

export const useDadosGeraisStore = defineStore('dadosGerais', () => {
    const dadosGerais = ref(JSON.parse(JSON.stringify(dadosGeraisModel)))
    function reset() {
        dadosGerais.value = JSON.parse(JSON.stringify(dadosGeraisModel))
    }
    return {
        dadosGerais,
        reset
    }
}, {
    persist: true
})