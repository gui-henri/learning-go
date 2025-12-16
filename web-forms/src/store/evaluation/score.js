import { defineStore } from 'pinia';

export const useScoreStore = defineStore('score', {
    state: () => ({
        score: {
     
            diagnostico_primario: '',
            diagnostico_secundario: '',
            domicilio_risco: false,
            impedimento_deslocamento: false,

            alimentacao_parenteral: null,
            aspiracao_traqueo: null,
            ventilacao_mecanica: null,
            medicacao_parenteral: null,


            banho: null,
            vestir: null,
            higiene_pessoal: null,
            transferencia: null,
            continencia: null,
            alimentacao: null,
            atividades: null,

    
            estado_nutricional: null,
            via_alimentacao_medicacao: null,
            internacoes_ultimo_ano: null,
            aspiracao_vias_aereas: null,
            lesoes: null,
            via_medicacoes: null,
            exercicios_ventilatorios: null,
            uso_oxigenoterapia: null,
            nivel_consciencia: null
        }
    }),
    actions: {
        reset() {
            this.score = {
                diagnostico_primario: '',
                diagnostico_secundario: '',
                domicilio_risco: false,
                impedimento_deslocamento: false,
                alimentacao_parenteral: null,
                aspiracao_traqueo: null,
                ventilacao_mecanica: null,
                medicacao_parenteral: null,
                banho: null,
                vestir: null,
                higiene_pessoal: null,
                transferencia: null,
                continencia: null,
                alimentacao: null,
                estado_nutricional: null,
                via_alimentacao_medicacao: null,
                internacoes_ultimo_ano: null,
                aspiracao_vias_aereas: null,
                lesoes: null,
                via_medicacoes: null,
                exercicios_ventilatorios: null,
                uso_oxigenoterapia: null,
                nivel_consciencia: null
            };
        }
    }
});