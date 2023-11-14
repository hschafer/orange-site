<script setup>
import { ref } from 'vue'
import axios from 'axios'
import PostListing from '../components/PostListing.vue'
import CommentBox from '../components/CommentBox.vue'
import Comment from '../components/Comment.vue'
import { useRoute } from 'vue-router'

const route = useRoute();
const post = ref(null);
const comments = ref([]);

axios.get(`http://localhost/api/post/${route.params.id}`).then(response => {
  post.value = response.data;
});

axios.get(`http://localhost/api/post/${route.params.id}/comments`).then(response => {
  comments.value = response.data;
});
</script>

<template>
  <div id="container">
    <PostListing
      :title="post.Title"
      :author="post.CreatorName"
      :rank="index+1"
      :url="post.Url"
      :createdOn="post.CreatedOn"
      :numComments="post.NumComments"
      :showRank="false"
    />

    <CommentBox></CommentBox>

    <div id="comments">
      <Comment
        v-for="(comment, index) in comments"
        :comment="comment"
        :indent="0"
      />
    </div>
  </div>
</template>

<style scoped>
  #container {
    margin-left: 1em;
    margin-top: 0.5em;
  }
</style>