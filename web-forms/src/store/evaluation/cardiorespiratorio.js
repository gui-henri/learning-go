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

    usa_oxigenioterapia: false,
    fonte_oxigenio: null,
    frequencia_oxigenio: null,
    vezes_dia_oxigenio: null,
    modalidade_oxigenio: null,
    fluxo_oxigenio: null
}

export const useRespiratorioStore = defineStore('respiratorio', () => {
    const respiratorio = ref(respiratorioModel)
    return {
        respiratorio,
    }
}, {
    persist: true
})