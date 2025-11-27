import { defineStore } from 'pinia'
import { ref } from 'vue';

const antimicrobianosBase = {
    nome_antimicrobiano: null,
    data_inicio: null,
    data_termino: null
}

export const exameFisicoModel = {
    estado_geral: null,
    avaliacao_locomotora: null,
    possui_dreno: false,
    local_dreno: null,
    data_implantacao_dreno: null,
    curativos_dreno: null,
    possui_acesso: false,
    antimicrobianos: []
}

export const useExameFisicoStore = defineStore('exameFisico', () => {

    const exameFisico = ref(exameFisicoModel)
    function adicionarAntimicrobiano() {
        exameFisico.value.antimicrobianos.push({ ...antimicrobianosBase })
    }

    function removerAntimicrobiano(index) {
        if (exameFisico.value.antimicrobianos.length > 0) {
            exameFisico.value.antimicrobianos.splice(index, 1)
        }
    }
    return {
        exameFisico,
        adicionarAntimicrobiano,
        removerAntimicrobiano
    }
}, {
    persist: true
})