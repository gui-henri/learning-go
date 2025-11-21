import { defineStore } from 'pinia'
import { ref } from 'vue';

export const cuidadoresModel = {
    medico_solicitante: null,
    contato_medico: null,

    possui_cuidador: false,
    nome_cuidador: null,
    contato_cuidador: null,
    turno_cuidador: null,

    precisa_treinamento: false,
    obs_treinamento: null
}

export const useCuidadoresStore = defineStore('cuidadores', () => {
    
    const cuidadores = ref(cuidadoresModel)

    return {
        cuidadores,
    }
}, {
    persist: true
})