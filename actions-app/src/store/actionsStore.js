import { defineStore } from 'pinia';
import { getActions, getBestActions } from '../services/actionsService';

export const useActionsStore = defineStore('actions', {
    state: () => ({
        actions: [],
        bestActions: [],
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
        async fetchBestActions() {
            try {
                const response = await getBestActions();
                this.bestActions = response.data;
            } catch (error) {
                console.error(error);
            }
        },
    },
});