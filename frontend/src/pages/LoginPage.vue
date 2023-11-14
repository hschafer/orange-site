<script setup>
import { ref } from 'vue'
import axios from 'axios'

const username = ref(null);
const password = ref(null);
const revealPassword = ref(false);
// TODO Should this be a component to reduce Register redundancy?
function submit() {
  if (!username.value) {
    alert("Username must not be empty");
  } else if (!password.value) {
    alert("Password must not be empty");
  } else {
    // TODO more password validation:?
    axios.post("/api/login", {
      "username": username.value,
      "password": password.value,  // TODO Encrypt?
    }).then((response) => {
      console.log(response)
    })

  }
}

</script>
<template>
    <div class="hello">
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
      <button @click="submit">Log in</button>
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

</style>