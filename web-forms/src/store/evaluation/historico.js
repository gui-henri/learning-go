import { defineStore } from 'pinia'
import { ref } from 'vue';

export const historicoModel = {
    historia_doenca_atual: null,
    anamnese_enfermagem: null,
    plano_terapeutico: null
}

export const useHistoricoStore = defineStore('historico', () => {
    
    const historico = ref(JSON.parse(JSON.stringify(historicoModel)))

    function reset() {
        historico.value = JSON.parse(JSON.stringify(historicoModel))
    }

    return {
        historico,
        reset
    }
}, {
    persist: true
})