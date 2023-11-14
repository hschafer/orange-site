<script setup>
import { ref } from 'vue'
import axios from 'axios'
import PostListing from '../components/PostListing.vue'

const posts = ref([]);

axios.get("http://localhost/api/posts").then(response => {
  posts.value = response.data;
});

</script>
<template>
    <div id="posts">
      <PostListing
        v-for="(post, index) in posts"
        :title="post.Title"
        :author="post.CreatorName"
        :rank="index+1"
        :url="post.Url"
        :createdOn="post.CreatedOn"
        :numComments="post.NumComments"
      />
    </div>
</template>

<style scoped>
  #posts {
    padding-top: 0.5em;
    padding-left: 1em;
  }
</style>