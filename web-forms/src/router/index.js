import AppLayout from '@/layout/AppLayout.vue';
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: AppLayout,
            children: [
                {
                path: '/',
                name: 'home',
                component: () => import('@/views/Home.vue')
                },
                {
                    path: '/avaliacao',
                    name: 'formulario_avaliacao',
                    component: () => import('@/views/Evaluation.vue')
                }
            ]
        }
    ]
});

export default router;
