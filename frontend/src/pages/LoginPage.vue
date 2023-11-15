<script setup>
import { ref } from 'vue'
import {useStore} from 'vuex'

const store = useStore();

const username = ref(null);
const password = ref(null);
const revealPassword = ref(false);

// TODO Should this be a component to reduce Register redundancy?
// TODO Add ability to hit enter on inputs to submit?
function submit() {
  if (!username.value) {
    alert("Username must not be empty");
  } else if (!password.value) {
    alert("Password must not be empty");
  } else {
    store.dispatch("login", {
      "username": username.value,
      "password": password.value,
    })
  }
}

</script>
<template>
    <div class="hello">
      <div v-if="this.$store.getters.isAuthenticated">
        <p>Already logged in! :)</p>
      </div>
      <div v-else>
        <h1>Login</h1>
        <div>
          <label for="username">Username: </label>
          <input v-model="username" />
        </div>
        <div>
          <label for="password">Password: </label>
          <input :type="revealPassword ? 'text' : 'password'"  v-model="password" />
          <input class="revealPassword" type="checkbox" v-model="revealPassword"> Show Password
        </div>
        <button @click="submit">Log in</button> <span class="loginMessage" v-if="this.$store.getters.getLoginMessage">{{ this.$store.getters.getLoginMessage }}</span>
      </div>
  </div>
</template>

<style scoped>
label {
  display: inline-block;
  min-width: 6em;
}

input {
  margin-bottom: 1em;
}

.revealPassword {
  margin-left: 1em;
}

.loginMessage {
  color: red;
  margin-left: 2.4em;
}

</style>