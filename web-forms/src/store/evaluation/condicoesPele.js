import { defineStore } from 'pinia';
import { ref } from 'vue';

const curativoBase = {
    local_curativo: '',
    tipo_cobertura: '',
    tamanho_curativo: '',
    qtd_curativo: '',
    frequencia_troca: '',
    materiais_curativo: ''
};

export const useCondicoesPeleStore = defineStore('condicoesPele', () => {

    const condicoesPele = ref({
        condicao_cutanea_mucosa: null,
        realiza_curativo: false,
        curativos: [] 
    });

    function adicionarCurativo() {

        condicoesPele.value.curativos.push({ ...curativoBase });
    }

    function removerCurativo(index) {
        if (condicoesPele.value.curativos.length > 0) {
            condicoesPele.value.curativos.splice(index, 1);
        }
    }

    return {
        condicoesPele,
        adicionarCurativo,
        removerCurativo
    };
}, {
    persist: true
});