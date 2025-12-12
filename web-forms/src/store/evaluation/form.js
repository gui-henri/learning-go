import { defineStore } from 'pinia'
import { ref } from 'vue';

/*
importante para quando for transformar em TypeScript: 

avaliationFormStoreModel pode ser null ou pode seguir esta scruct:

type AvaliationFormOptions struct {
    Version            int
    DadosGerais        FormMetadata `json:"dados_gerais"`
    Contato            FormMetadata `json:"contato"`
    DadosClinicos      FormMetadata `json:"dados_clinicos"`
    CardioRespiratorio FormMetadata `json:"cardio_respiratorio"`
    Nutricional        FormMetadata `json:"nutricional"`
    Eliminacoes        FormMetadata `json:"eliminacoes"`
    Avaliacao          FormMetadata `json:"avaliacao"`
}
*/
export const avaliationFormStoreModel = null

export const useAvaliationForm = defineStore('avaliationForm', () => {
    const avaliationForm = ref(avaliationFormStoreModel)
    const avaliationId = ref(null)

    return {
        avaliationForm,
        avaliationId
    }
}, {
    persist: true
})