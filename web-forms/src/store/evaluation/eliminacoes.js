import { defineStore } from 'pinia';

export const useEliminacoesStore = defineStore('eliminacoes', {
    persist: true,
    state: () => ({
        eliminacoes: {
            funcao_intestinal: null,
            via_evacuacao_fisiologica: false,
            via_evacuacao_estomia: false,
            tipo_estomia: null,
            placa_estomia: '',
            bolsa_estomia: '',

            diurese: null,
            num_snd_diurese: null,
            usa_fralda: false,
            trocas_fralda_dia: '',
            sva: false,

            cateterismo_intermitente: false,
            vezes_cateterismo_dia: '',

            svd: false,
            num_svd: '',
            vias_svd: null,
            irrigacao_svd: false,
            data_troca_svd: null,
            volume_diurese_svd: null,

            cistostomia: false,
            num_cistostomia: '',
            vias_cistostomia: null,
            irrigacao_cistostomia: false,
            qtd_irrigacao_cistostomia: '',
            medicacao_irrigacao_cistostomia: false,
            nome_medicacao_cistostomia: '',
            volume_diurese_cistostomia: null,

            preservativo: false,
            volume_diurese_preservativo: null
        }
    }),
    actions: {
        reset() {
            this.eliminacoes = {
                funcao_intestinal: null,
                via_evacuacao_fisiologica: false,
                via_evacuacao_estomia: false,
                tipo_estomia: null,
                placa_estomia: '',
                bolsa_estomia: '',
                diurese: null,
                usa_fralda: false,
                trocas_fralda_dia: '',
                sva: false,
                cateterismo_intermitente: false,
                vezes_cateterismo_dia: '',
                svd: false,
                num_svd: '',
                vias_svd: null,
                irrigacao_svd: false,
                data_troca_svd: null,
                volume_diurese_svd: null,
                cistostomia: false,
                num_cistostomia: '',
                vias_cistostomia: null,
                irrigacao_cistostomia: false,
                qtd_irrigacao_cistostomia: '',
                medicacao_irrigacao_cistostomia: false,
                nome_medicacao_cistostomia: '',
                volume_diurese_cistostomia: null,
                preservativo: false,
                volume_diurese_preservativo: null
            };
        }
    }
});