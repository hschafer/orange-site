<script setup>
import { ref } from 'vue'
import { useStore } from 'vuex'
import { useRoute, useRouter } from 'vue-router'

const props = defineProps({
    message: String,
    action: String,
    buttonLabel: String,
})

const store = useStore();
const route = useRoute();
const router = useRouter();

const username = ref(null);
const password = ref(null);
const revealPassword = ref(false);
const errorMessage = ref("");

// TODO Add ability to hit enter on inputs to submit?
async function submit() {
  if (!username.value) {
    alert("Username must not be empty");
  } else if (!password.value) {
    alert("Password must not be empty");
  } else {
    await store.dispatch(props.action, {
      "username": username.value,
      "password": password.value,
    }).catch((error) => {
      console.log("Login/Register error", error);
      if (error?.response?.data?.Message) {
        errorMessage.value =  error.response.data.Message
      } else {
        errorMessage.value = `Unable to ${action}`
      }
    });

    // TODO navigate to account page if registration?
    if (store.getters.isAuthenticated) {
        // Redirect if we logged in successfully
        router.push(route.query.redirectFrom || "/")
    }
  }
}
</script>

<template>
    <p>{{ message }}</p>
    <div>
      <label for="username">Username: </label>
      <input v-model="username" />
    </div>
    <div>
      <label for="password">Password: </label>
      <input :type="revealPassword ? 'text' : 'password'"  v-model="password" />
      <input class="revealPassword" type="checkbox" v-model="revealPassword"> Show Password
    </div>
    <!-- TODO Login messages are linked between instances -->
    <button @click="submit">{{ buttonLabel }}</button> <span class="errorMessage" v-if="errorMessage">{{ errorMessage }}</span>
</template>

<style scoped>
  label {
    display: inline-block;
    min-width: 6em;
    font-weight:bold;
  }

  p {
    margin-top: 0;
  }

  input {
    margin-bottom: 1em;
  }

  .revealPassword {
    margin-left: 1em;
  }

  .errorMessage {
    color: red;
    margin-left: 2.4em;
  }
</style>