import { defineStore } from 'pinia'
import { ref } from 'vue';

export const historicoModel = {
    historia_doenca_atual: null,
    anamnese_enfermagem: null,
    plano_terapeutico: null
}

export const useHistoricoStore = defineStore('historico', () => {
    
    const historico = ref(historicoModel)

    return {
        historico,
    }
}, {
    persist: true
})