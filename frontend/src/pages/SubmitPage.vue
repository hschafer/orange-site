<script setup>
import axios from 'axios'
import { ref } from 'vue'
import { useRouter } from 'vue-router';
import { useStore } from 'vuex'

const title = ref(null);
const url = ref(null);

const store = useStore();
const router = useRouter();

function submit() {
  if (store.getters.isAuthenticated) {
    if (!title.value) {
      alert("Must include title");
      return;
    }
    if (!url.value) {
      alert("Must include link");  // TODO validation based on regex?
      return;
    }

    axios.post("/api/post/new", {
      "title": title.value,
      "url": url.value
    }, {
      "headers": {
        "Authorization": `Bearer ${store.getters.getToken}`
      }
    }).then(() => {
      console.log("Link submitted");
      router.push("/")
    }).catch((error) => {
      console.log("Error posting link", error)
    });
  } else {
    console.log("Need to be signed in to submit a link");
  }
}
</script>

<template>
    <div class="hello">
      <h1>Submit New Link</h1>

      <div v-if="this.$store.getters.isAuthenticated">
        <p>Submit a new link</p>

        <div>
          <label for="title">Post Title: </label>
          <input v-model="title" />
        </div>
        <div>
          <label for="title">URL: </label>
          <input v-model="url" />
        </div>

        <button @click="submit">Submit</button>

      </div>
      <div v-else>
        <p>Must be logged in to submit a link!</p>
      </div>
    </div>
</template>

<style scoped>
  input {
    margin-bottom: 1em;
    width: 40vw;
  }

  label {
    min-width: 6em;
    display: inline-block;
    font-weight:bold;
  }
</style>