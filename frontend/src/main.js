import './style.css'

import { createRouter, createWebHistory } from 'vue-router'

import AboutPage from './pages/AboutPage.vue'
import App from './App.vue'
import HomePage from './pages/HomePage.vue'
import PostPage from './pages/PostPage.vue'
import VueDemoPage from './pages/VueDemoPage.vue'
import { createApp } from 'vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", component: HomePage},
        { path: "/post/:id", component: PostPage },
        { path: "/about", component: AboutPage},
        { path: "/vue", component: VueDemoPage},

    ]
});

const app = createApp(App);
app.use(router);
app.mount("#app");
