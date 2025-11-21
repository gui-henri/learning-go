import { defineStore } from 'pinia'
import { ref } from 'vue';

export const segurancaModel = {
    alergia_relatada: false,
    tipo_alergia: null,
    quais_alergias: null,
    precaucao: null,
    cuidados_paliativos: null
}

export const useSegurancaStore = defineStore('seguranca', () => {
    
    const seguranca = ref(segurancaModel)

    return {
        seguranca,
    }
}, {
    persist: true
})