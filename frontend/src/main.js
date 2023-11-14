import './style.css'

import { createRouter, createWebHistory } from 'vue-router'

import AboutComponent from './components/AboutComponent.vue'
import App from './App.vue'
import HomeComponent from './components/HomeComponent.vue'
import { createApp } from 'vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", component: HomeComponent},
        { path: "/about", component: AboutComponent},
    ]
});

const app = createApp(App);
app.use(router);
app.mount("#app");
