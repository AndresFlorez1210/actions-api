import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../store/authStore';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import Actions from '../components/actions/Actions.vue';

const routes = [
    { path: '/', component: Login },
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    { path: '/actions', component: Actions, meta: { requiresAuth: true }},
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const authStore = useAuthStore();
    if(to.meta.requiresAuth && !authStore.token) {
        next({ name: 'login' });
    } else {
        next();
    }
});

export default router;