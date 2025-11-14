import { createApp } from 'vue';
import App from './App.vue';
import router from './router';

import Aura from '@primeuix/themes/aura';
import PrimeVue from 'primevue/config';
import ConfirmationService from 'primevue/confirmationservice';
import ToastService from 'primevue/toastservice';
import { definePreset } from '@primeuix/themes';

import '@/assets/tailwind.css';
import '@/assets/styles.scss';

const app = createApp(App);

app.use(router);

const Interne = definePreset(Aura, {
    semantic: {
        primary: { 0: '#ffffff', 50: '#ffeeef', 100: '#ffdadb', 200: '#ffbbbd', 300: '#ff8b8f', 400: '#ff4950', 500: '#ff111a', 600: '#ff0009', 700: '#e70008', 800: '#be0007', 900: '#a80006', 950: '#560003' }
    }
});

app.use(PrimeVue, {
    theme: {
        preset: Interne,
        options: {
            darkModeSelector: '.app-dark'
        }
    }
});
app.use(ToastService);
app.use(ConfirmationService);

app.mount('#app');
