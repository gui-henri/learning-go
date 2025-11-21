import { defineStore } from 'pinia'
import { ref } from 'vue';

export const exameFisicoModel = {
    estado_geral: null,
    avaliacao_locomotora: null,
    possui_dreno: false,
    local_dreno: null,
    data_implantacao_dreno: null,
    curativos_dreno: null,

    possui_acesso: false
}

export const useExameFisicoStore = defineStore('exameFisico', () => {
    
    const exameFisico = ref(exameFisicoModel)

    return {
        exameFisico,
    }
}, {
    persist: true
})