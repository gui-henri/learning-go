import { defineStore } from 'pinia'
import { ref } from 'vue';

const alergiaBase = {
    tipo_alergia: null,
    quais_alergias: null,
    precaucao: null,
    cuidados_paliativos: null
}

export const segurancaModel = {
    alergias: [],
}

export const useSegurancaStore = defineStore('seguranca', () => {
    const seguranca = ref(JSON.parse(JSON.stringify(segurancaModel)))

    function adicionarAlergia() {
        seguranca.value.alergias.push({ ...alergiaBase })
    }

    function removerAlergia(index) {
        if (seguranca.value.alergias.length > 0) {
            seguranca.value.alergias.splice(index, 1)
        }
    }

    function reset() {
        seguranca.value = JSON.parse(JSON.stringify(segurancaModel))
    }

    return {
        seguranca,
        adicionarAlergia,
        removerAlergia,
        reset
    }
}, {
    persist: true
})