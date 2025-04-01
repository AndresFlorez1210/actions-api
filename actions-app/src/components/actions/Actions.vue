<template>
  <div class="w-full h-screen">
      <Navbar />
      
      <main class="flex justify-center items-center">
          <h1 class="text-3xl mt-10 text-center">Listado de acciones</h1>
      </main>

      <ActionTable :actions="actionsStore.actions" />
      <BestActions :bestActions="bestActionsInfo" />
  </div>
</template>

<script>
import { onMounted, computed } from 'vue';
import { useActionsStore } from '../../store/actionsStore';
import Navbar from '../commons/Navbar.vue';
import ActionTable from './ActionTable.vue';
import BestActions from './BestActions.vue';

export default {
  name: 'ActionsList',
  components: {
      Navbar,
      ActionTable,
      BestActions
  },
  setup() {
      const actionsStore = useActionsStore();
      
      const bestActionsInfo = computed(() => {
          return actionsStore.bestActions;
      });

      onMounted(async () => {
          await actionsStore.fetchActions();
          await actionsStore.fetchBestActions();
      });

      return { 
          actionsStore,
          bestActionsInfo,
      };
  }
}
</script>