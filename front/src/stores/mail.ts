import { defineStore, acceptHMRUpdate } from "pinia";

export const useMailStore = defineStore({
    id: 'mail',
    state: () => ({
    }),
    getters: {},
    actions: {
    }
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useMailStore, import.meta.hot));
}