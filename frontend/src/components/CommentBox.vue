<script setup>
import axios from 'axios';
import { ref } from 'vue'
import { useStore } from 'vuex'

const props = defineProps({
  postID: Number,
  onComment: Function,
  parentID: {
    type: [Number, null],
    default: null
  }
});
const store = useStore();

const comment = ref("");

function addComment() {
  if (store.getters.isAuthenticated) {
    axios.post("/api/comment/new", {
      "comment": comment.value,
      "postID": props.postID,
      "parentID": props.parentID,
    }, {
      "headers": {
        "Authorization": `Bearer ${store.getters.getToken}`
      }
    }).then(() => {
      console.log("Comment posted");
      comment.value = "";  // Clear the text box
      props.onComment();
    }).catch((error) => {
      console.log("Error posting comment", error)
    });
  } else {
    console.log("Need to be signed in to comment");
  }
}
</script>

<template>
    <div class="commentBox">
      <textarea
        :disabled="!this.$store.getters.isAuthenticated"
        v-model="comment"
        :placeholder="this.$store.getters.isAuthenticated ? '' : 'Must be logged in to comment' ">
      </textarea>
      <button :disabled="!this.$store.getters.isAuthenticated" @click="addComment">add comment</button>
    </div>
</template>

<style scoped>
  .commentBox {
    padding-top: 0.5em;
    display: flex;
    flex-direction: column;
  }
  .commentBox textarea {
    width: 50vw;
    height: 10vh;
    min-height: 5em;
  }

  .commentBox button {
    margin-top: 1em;
    width: fit-content;
  }

</style>