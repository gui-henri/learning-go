import { defineStore } from 'pinia'
import { ref } from 'vue';

const responsavelBase = {
    nome: null,
    parentesco: null,
    email: null,
    telefone: null,
    forma_contato: null,
    numero_prioritario: false
}

export const contatoModel = {
    telefone_residencial: null,
    telefone_paciente: null,
    email_paciente: null,

    responsaveis: [{ ...responsavelBase }]
}

export const useContatoStore = defineStore('contato', () => {

    const contato = ref(JSON.parse(JSON.stringify(contatoModel)))

    function adicionarResponsavel() {
        contato.value.responsaveis.push({ ...responsavelBase })
    }

    function removerResponsavel(index) {
        if (contato.value.responsaveis.length > 1) {
            contato.value.responsaveis.splice(index, 1)
        }
    }

    return {
        contato,
        adicionarResponsavel,
        removerResponsavel
    }
}, {
    persist: true
})