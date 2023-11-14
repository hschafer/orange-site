<script setup>
import HeaderComponent from './components/HeaderComponent.vue'
</script>

<template>
  <HeaderComponent></HeaderComponent>
  <div id="appBody">
    <router-view></router-view>
  </div>
</template>