import './style.css'

import { createRouter, createWebHistory } from 'vue-router'

import AboutPage from './pages/AboutPage.vue'
import App from './App.vue'
import HomePage from './pages/HomePage.vue'
import PostPage from './pages/PostPage.vue'
import axios from 'axios'
import { createApp } from 'vue'
import store from './store/'

//axios.defaults.withCredentials = true
axios.defaults.baseURL = 'http://localhost/';  // TODO

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", component: HomePage},
        { path: "/post/:id", component: PostPage },
        { path: "/about", component: AboutPage},

    ]
});

const app = createApp(App);
app.use(router);
app.use(store)
app.mount("#app");