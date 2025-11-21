import { defineStore } from 'pinia'
import { ref } from 'vue';

export const contatoModel = {
    telefone_residencial: null,
    telefone_paciente: null,
    email_paciente: null,
    
    nome_responsavel: null,
    parentesco_responsavel: null,
    email_responsavel: null,
    telefone_responsavel: null,
    
    forma_contato: null,
    numero_prioritario: false
}

export const useContatoStore = defineStore('contato', () => {
    
    const contato = ref(JSON.parse(JSON.stringify(contatoModel)))

    function limparContato() {
        contato.value = JSON.parse(JSON.stringify(contatoModel))
    }

    return {
        contato,
        limparContato
    }
}, {
    persist: true
})