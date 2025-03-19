import { defineStore } from 'pinia';
import { getActions } from '../services/actionsService';

export const useActionsStore = defineStore('actions', {
    state: () => ({
        actions: [],
    }),
    actions: {
        async fetchActions() {
            try {
                const response = await getActions();
                this.actions = response.data;
            } catch (error) {
                console.error(error);
            }
        },
    },
});