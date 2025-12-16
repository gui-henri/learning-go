import { defineStore } from 'pinia'
import { ref } from 'vue';

export const respiratorioModel = {
    padrao_respiratorio: null,
    via_aerea: null,

    numero_tubo: null,
    data_orotraqueal: null,

    suporte_ventilatorio: null,

    modo_vni: null,
    frequencia_vni: null,
    fio2_vni: null,
    mascara_vni: null,
    tamanho_mascara_vni: null,

    modo_avm: null,
    frequencia_avm: null,
    fio2_avm: null,
    mascara_avm: null,
    tamanho_mascara_avm: null,

    usa_oxigenioterapia: false,
    fonte_oxigenio: null,
    frequencia_oxigenio: null,
    vezes_dia_oxigenio: null,
    modalidade_oxigenio: null,
    fluxo_oxigenio: null
}

export const useRespiratorioStore = defineStore('respiratorio', () => {
    const respiratorio = ref(JSON.parse(JSON.stringify(respiratorioModel)))
    function reset() {
        respiratorio.value = JSON.parse(JSON.stringify(respiratorioModel))
    }

    return {
        respiratorio,
        reset,
    }
}, {
    persist: true
})