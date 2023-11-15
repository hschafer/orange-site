<script setup>
import { useStore } from 'vuex'
import { useRoute, useRouter } from 'vue-router'

const store = useStore();
const route = useRoute();
const router = useRouter();

function login() {
    console.log("HELLO", route.path)
    router.push({
        path: "/login",
        query: {
            redirectFrom: route.path
        }
    });
}

function logout() {
    store.dispatch('logout');
}

</script>

<template>
    <div id="header">
        <router-link to="/"><div id="icon">H</div> <span id="title">HunterNews</span></router-link>
        <router-link to="/">posts</router-link> |
        <router-link to="/about">about</router-link> |
        <router-link to="/submit">submit</router-link> |
        <a @click="login" href="" v-if="!this.$store.getters.isAuthenticated">login</a><a @click="logout" href="" v-else>logout</a>
    </div>
</template>

<style scoped>
    #header {
        width: 100%;
        text-align: left;
        background-color: #F86602;
        padding: 0.75em 1em 0.5em 1em;
    }

    #header a {
        color: black;
        text-decoration: none;
    }

    #header a:first-child {
        margin-right: 0.75em;
    }


    #title {
        font-weight: bold;
    }

    #icon {
        display: inline-block;
        width: 24px;
        height: 24px;
        padding: 0.1em;
        text-align: center;
        margin-right: 0.2em;
        color: white;
        border: white;
        border-width: 1px;
        border-style: solid;
    }
</style>