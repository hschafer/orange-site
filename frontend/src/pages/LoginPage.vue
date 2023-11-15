<script setup>
import { useStore } from 'vuex'
import AccountForm from '../components/AccountForm.vue'
</script>

<template>
    <div class="hello">
      <div v-if="this.$store.getters.isAuthenticated">
        <p>Already logged in! :)</p>
      </div>
      <div v-else>
        <h1>Login</h1>
        <AccountForm />
      </div>
  </div>
</template>

<style scoped>
</style>