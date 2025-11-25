import { defineStore } from 'pinia';

export const useCondicoesPeleStore = defineStore('condicoesPele', {
    state: () => ({
        condicoesPele: {
            condicao_cutanea_mucosa: null,
            realiza_curativo: false,
            local_curativo: '',
            tipo_cobertura: '',
            tamanho_curativo: '',
            qtd_curativo: '',
            frequencia_troca: '',
            materiais_curativo: ''
        }
    }),
    actions: {
        clean() {
            this.condicoesPele = {
                condicao_cutanea_mucosa: null,
                realiza_curativo: false,
                local_curativo: '',
                tipo_cobertura: '',
                tamanho_curativo: '',
                qtd_curativo: '',
                frequencia_troca: '',
                materiais_curativo: ''
            };
        }
    }
});