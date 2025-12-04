import { defineStore } from 'pinia';

export const useObservacoesStore = defineStore('observacoes', {
    persist: true,
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