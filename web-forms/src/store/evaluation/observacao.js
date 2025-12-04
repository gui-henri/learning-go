import { defineStore } from 'pinia';

export const useObservacoesStore = defineStore('observacoes', {
    state: () => ({
        observacoes: {
            texto: ''
        }
    }),
    actions: {
        clean() {
            this.observacoes = {
                texto: ''
            };
        }
    }
});