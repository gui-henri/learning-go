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
    
    const endereco = ref(JSON.parse(JSON.stringify(enderecoModel)))

    function reset() {
        endereco.value = JSON.parse(JSON.stringify(enderecoModel))
    }

    return {
        endereco,
        reset
    }
}, {
    persist: true
})