import { defineStore } from 'pinia';

export const useNutricionalStore = defineStore('nutricional', {
    state: () => ({
        nutricional: {
    
            alimentacao_oral: false,
            alimentacao_enteral: false,
            alimentacao_parenteral: false,
            sonda: false,
            botton: false,
            suplemento: false,

            via_enteral: null,
            via_parenteral: null,
            adaptador_sonda: null,
            restricao_alimentar: '',
            tipo_dieta: null,
            forma_administracao: null,
            marca_bomba: null,

            data_ultima_troca: null,
            gavando: '',
            volume_diario: '',
            qual_suplemento: ''
        }
    }),
    actions: {
        clean() {
            this.nutricional = {
                alimentacao_oral: false,
                alimentacao_enteral: false,
                alimentacao_parenteral: false,
                sonda: false,
                botton: false,
                suplemento: false,
                via_enteral: null,
                via_parenteral: null,
                adaptador_sonda: null,
                restricao_alimentar: '',
                tipo_dieta: null,
                forma_administracao: null,
                marca_bomba: null,
                data_ultima_troca: null,
                gavando: '',
                volume_diario: '',
                qual_suplemento: ''
            };
        }
    }
});