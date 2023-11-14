import './style.css'

import { createRouter, createWebHistory } from 'vue-router'

import AboutPage from './pages/AboutComponent.vue'
import App from './App.vue'
import HomePage from './pages/HomeComponent.vue'
import { createApp } from 'vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", component: HomePage},
        { path: "/about", component: AboutPage},
    ]
});

const app = createApp(App);
app.use(router);
app.mount("#app");
