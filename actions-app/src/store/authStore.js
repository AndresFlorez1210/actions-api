import { defineStore } from 'pinia';
import { login, register, logout } from '../services/authService';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: localStorage.getItem('token') || null,
    }),
    actions: {
        async login(credentials) {
            try {
                const response = await login(credentials);
                this.token = response.data.token;
                localStorage.setItem('token', this.token);
                return true;
            } catch (error) {
                console.error(error);
                return false;
            }
        },
        async register(userData) {
            try {
                const response = await register(userData);
                this.token = response.data.token;
                localStorage.setItem('token', this.token);
                return true;
            } catch (error) {
                console.error(error);
                return false;
            }
        },
        async logout() {
            this.token = null;
            logout();
        },
    },
});
