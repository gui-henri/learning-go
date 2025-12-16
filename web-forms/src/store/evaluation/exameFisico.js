import { defineStore } from 'pinia'
import { ref } from 'vue';

const antimicrobianosBase = {
    nome_antimicrobiano: null,
    data_inicio: null,
    data_termino: null
}

const drenosBase = {
    local_dreno: null,
    data_implantacao_dreno: null,
    curativos_dreno: null,
}

export const exameFisicoModel = {
    estado_geral: null,
    avaliacao_locomotora: null,
    drenos: [],
    possui_acesso: false,
    antimicrobianos: []
}

export const useExameFisicoStore = defineStore('exameFisico', () => {
    const exameFisico = ref(JSON.parse(JSON.stringify(exameFisicoModel)))

    function adicionarDreno() {
        exameFisico.value.drenos.push({ ...drenosBase })
    }

    function removerDreno(index) {
        if (exameFisico.value.drenos.length > 0) {
            exameFisico.value.drenos.splice(index, 1)
        }
    }

    function adicionarAntimicrobiano() {
        exameFisico.value.antimicrobianos.push({ ...antimicrobianosBase })
    }

    function removerAntimicrobiano(index) {
        if (exameFisico.value.antimicrobianos.length > 0) {
            exameFisico.value.antimicrobianos.splice(index, 1)
        }
    }

    function reset() {
        exameFisico.value = JSON.parse(JSON.stringify(exameFisicoModel))
    }

    return {
        exameFisico,
        adicionarAntimicrobiano,
        removerAntimicrobiano,
        adicionarDreno,
        removerDreno,
        reset
    }
}, {
    persist: true
})