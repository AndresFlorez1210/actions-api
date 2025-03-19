<template>
    <div>
        <h1>Listado de acciones</h1>
        <button @click="logout">Cerrar sesión</button>
        <ul>
            <li v-for="action in actionsStore.actions" :key="action.ticker">
                <h2>{{ action.action }}</h2>
                <p>{{ action.company }}</p>
                <p>{{ action.brokerage }}</p>
                <p>{{ action.rating_from }}</p>
                <p>{{ action.rating_to }}</p>
                <p>{{ action.time }}</p>
            </li>
        </ul>
    </div>
</template>

<script>
import { onMounted } from 'vue';
import { useAuthStore } from '../store/authStore';
import { useActionsStore } from '../store/actionsStore';
import { useRouter } from 'vue-router';

export default {
    name: 'ActionsList',
    setup() {
        const actionsStore = useActionsStore();
        const router = useRouter();
        const authStore = useAuthStore();

        onMounted(async () => {
            await actionsStore.fetchActions();
        });

        const logout = () => {
            authStore.logout();
            router.push('/login');
        };

        return { actionsStore, logout };  
    }
}
</script>