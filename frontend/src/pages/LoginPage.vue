<script setup>
import AccountForm from '../components/AccountForm.vue'
</script>

<template>
    <div class="hello">
      <div v-if="this.$store.getters.isAuthenticated">
        <p>Already logged in! :)</p>
      </div>
      <div v-else>
        <h1>Login</h1>
        <p></p>
        <AccountForm
          message="Log in to an existing account"
          action="login"
          buttonLabel="Log in"
        />

        <h1>Register</h1>
        <AccountForm
          message="Don't have an account? Create a new one for free!"
          action="register"
          buttonLabel="Register"
        />
      </div>
  </div>
</template>

<style scoped>
  p {
    margin-top: 0em;
  }
</style>