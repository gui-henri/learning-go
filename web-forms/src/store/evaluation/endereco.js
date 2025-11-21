import { defineStore } from 'pinia'
import { ref } from 'vue';

export const enderecoModel = {
    cep: null,
    estado: null,
    cidade: null,
    bairro: null,
    logradouro: null,
    numero: null,
    complemento: null,
    ponto_referencia: null
}

export const useEnderecoStore = defineStore('endereco', () => {
    
    const endereco = ref(enderecoModel)

    return {
        endereco    
    }
}, {
    persist: true
})