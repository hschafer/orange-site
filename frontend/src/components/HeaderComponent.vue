<script setup>

</script>

<template>
    <div id="header">
        <router-link to="/"><div id="icon">H</div> <span id="title">HunterNews</span></router-link> <router-link to="/">posts</router-link> | <router-link to="/about">about</router-link>
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