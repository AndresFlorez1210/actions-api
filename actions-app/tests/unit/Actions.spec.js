import { mount } from '@vue/test-utils'
import { createTestingPinia } from '@pinia/testing'
import { useActionsStore } from '@/store/actionsStore'
import { useAuthStore } from '@/store/authStore'
import Actions from '@/components/actions/Actions.vue'
import { nextTick } from 'vue'
import { describe, it, expect, beforeEach, jest } from '@jest/globals'
import { createRouter, createWebHistory } from 'vue-router'

describe('Actions.vue', () => {
    let wrapper
    let actionsStore
    let authStore
    let router

    const mockActions = [
        {
            ticker: 'AAPL1',
            company: 'Apple Inc',
            brokerage: 'Morgan Stanley',
            rating_from: 'Buy',
            rating_to: 'Buy',
            target_from: '$150.00',
            target_to: '$180.00'
        },
        {
            ticker: 'GOOGL1',
            company: 'Google Inc',
            brokerage: 'JP Morgan',
            rating_from: 'Buy',
            rating_to: 'Buy',
            target_from: '$2500.00',
            target_to: '$3000.00'
        }
    ]

    const mockBestActions = [
        {
            ticker: 'MSFT1',
            company: 'Microsoft Corp',
            brokerage: 'Goldman Sachs',
            rating_from: 'Buy',
            rating_to: 'Buy',
            target_from: '$300.00',
            target_to: '$350.00'
        },
        {
            ticker: 'AMZN1',
            company: 'Amazon Inc',
            brokerage: 'JP Morgan',
            rating_from: 'Buy',
            rating_to: 'Buy',
            target_from: '$2500.00',
            target_to: '$3000.00'
        },
        {
            ticker: 'TSLA1',
            company: 'Tesla Inc',
            brokerage: 'Morgan Stanley',
            rating_from: 'Buy',
            rating_to: 'Buy',
            target_from: '$150.00',
            target_to: '$180.00'
        }
    ]

    beforeEach(async () => {
        router = createRouter({
            history: createWebHistory(),
            routes: [
                {
                    path: '/',
                    name: 'Home',
                    component: {}
                },
                {
                    path: '/login',
                    name: 'Login',
                    component: {}
                }
            ]
        })
        
        wrapper = mount(Actions, {
            global: {
                plugins: [
                    createTestingPinia({
                        createSpy: jest.fn
                    }),
                    router
                ]
            }
        })
        actionsStore = useActionsStore()
        authStore = useAuthStore()
        
        actionsStore.actions = mockActions
        actionsStore.bestActions = mockBestActions

        await router.isReady()
    })

    it('renders the component', () => {
        expect(wrapper.exists()).toBe(true)
    })

    it('displays the correct number of actions in the table', () => {
        const rows = wrapper.findAll('tbody tr')
        expect(rows).toHaveLength(2)
    })

    it('displays pagination information correctly', () => {
        const paginationText = wrapper.find('.text-xs.xs\\:text-sm.text-gray-900')
        expect(paginationText.text()).toContain('Página 1 de 1')
        expect(paginationText.text()).toContain('2 acciones en total')
    })

    it('displays best actions when data is available', () => {
        const bestActionCards = wrapper.findAll('.md\\:w-4\\/12')
        expect(bestActionCards).toHaveLength(3)
        
        expect(wrapper.text()).toContain('MSFT1')
        expect(wrapper.text()).toContain('AMZN1')
        expect(wrapper.text()).toContain('TSLA1')
    })

    it('handles pagination next button click', async () => {
        actionsStore.actions = [...mockActions, ...mockActions, ...mockActions] // 6 items
        await nextTick()

        const nextButton = wrapper.find('button:last-child')
        await nextButton.trigger('click')

        expect(wrapper.vm.currentPage).toBe(2)
    })

    it('handles pagination prev button click', async () => {
        actionsStore.actions = [...mockActions, ...mockActions, ...mockActions]
        await nextTick()

        await wrapper.vm.nextPage()
        
        const prevButton = wrapper.find('button:first-child')
        await prevButton.trigger('click')

        expect(wrapper.vm.currentPage).toBe(1)
    })

    it('fetches actions on mount', () => {
        expect(actionsStore.fetchActions).toHaveBeenCalled()
        expect(actionsStore.fetchBestActions).toHaveBeenCalled()
    })

    it('handles logout', async () => {
        const logoutButton = wrapper.find('a.cursor-pointer')
        await logoutButton.trigger('click')

        expect(authStore.logout).toHaveBeenCalled()
        await router.push('/login')
        expect(router.currentRoute.value.path).toBe('/login')
    })

    it('disables prev button on first page', () => {
        const prevButton = wrapper.find('button:first-child')
        expect(prevButton.attributes('disabled')).toBeDefined()
    })

    it('disables next button on last page', async () => {
        while (wrapper.vm.currentPage < wrapper.vm.totalPages) {
            await wrapper.vm.nextPage()
        }

        const nextButton = wrapper.find('button:last-child')
        expect(nextButton.attributes('disabled')).toBeDefined()
    })
}) 