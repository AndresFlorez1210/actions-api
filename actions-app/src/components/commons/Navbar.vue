<template>
    <header class="bg-teal-400">
        <nav class="flex justify-between w-full bg-gray-500 text-white p-4">
            <a href="#"><span class="font-semibold text-xl tracking-tight">Actions app</span></a>
            <div class="md:items-center md:w-auto flex">
                <div class="flex text-sm">
                    <a class="p-2 ml-2 bg-white text-teal-500 font-semibold leading-none border border-gray-100 rounded hover:border-transparent hover:bg-gray-100 cursor-pointer" @click="logout">Cerrar sesión</a>
                </div>
            </div>
        </nav>
    </header>
</template>

<script>
import { useAuthStore } from '../../store/authStore';
import { useRouter } from 'vue-router';

export default {
    name: 'NavBar',
    setup() {
        const router = useRouter();
        const authStore = useAuthStore();

        const logout = () => {
            authStore.logout();
            router.push('/login');
        };

        return { logout };
    }
}
</script>