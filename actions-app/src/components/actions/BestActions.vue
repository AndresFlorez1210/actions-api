<template>
    <section class="pb-10 bg-blueGray-200 -mt-24" v-if="bestActions && bestActions.length >= 3">
        <h1 class="text-3xl mt-24 text-center">Mejores acciones</h1>
        <div class="container mx-auto px-4">
            <div class="flex flex-wrap">
                <div v-for="(action, index) in bestActions.slice(0,3)" :key="action.ticker" 
                     :class="['w-full md:w-4/12 px-4 text-center', index === 0 ? 'lg:pt-12 pt-6' : index === 1 ? '' : 'pt-6']">
                    <div class="relative flex flex-col min-w-0 break-words bg-white w-full mb-8 shadow-lg rounded-lg">
                        <div class="px-4 py-5 flex-auto">
                            <div class="text-white p-3 text-center inline-flex items-center justify-center w-12 h-12 mb-5 shadow-lg rounded-full"
                                 :class="index === 0 ? 'bg-red-400' : index === 1 ? 'bg-lightBlue-400' : 'bg-emerald-400'">
                                <i :class="index === 0 ? 'fas fa-award' : index === 1 ? 'fas fa-retweet' : 'fas fa-fingerprint'"></i>
                            </div>
                            <h6 class="text-xl font-semibold">{{ action.ticker }}</h6>
                            <p class="mt-2 mb-4 text-blueGray-500">COMPANY: {{ action.company }}</p>
                            <p class="mt-2 mb-4 text-blueGray-500">BROKERAGE: {{ action.brokerage }}</p>
                            <p class="mt-2 mb-4 text-blueGray-500">RATING FROM: {{ action.rating_from }}</p>
                            <p class="mt-2 mb-4 text-blueGray-500">RATING TO: {{ action.rating_to }}</p>
                            <p class="mt-2 mb-4 text-blueGray-500">TARGET FROM: {{ action.target_from }}</p>
                            <p class="mt-2 mb-4 text-blueGray-500">TARGET TO: {{ action.target_to }}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<script>
export default {
    name: 'BestActions',
    props: {
        bestActions: {
            type: Array,
            required: true
        }
    }
}
</script>