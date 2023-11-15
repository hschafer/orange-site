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
const loadedPost = ref(false);

const postID = parseInt(route.params.id);

axios.get(`/api/post/${postID}`).then(response => {
  loadedPost.value = true;
  post.value = response.data;
});

function loadComments() {
  axios.get(`/api/post/${postID}/comments`).then(response => {
    comments.value = response.data;
  });
}

loadComments();
</script>

<template>
  <div id="container">
    <PostListing
      :id="postID"
      :title="post.Title"
      :author="post.CreatorName"
      :rank="-1"
      :url="post.Url"
      :createdOn="post.CreatedOn"
      :numComments="post.NumComments"
      :showRank="false"
      v-if="loadedPost"
    />

    <CommentBox
      :postID="postID"
      :onComment="loadComments"
    />

    <div id="comments">
      <Comment
        v-for="(comment, index) in comments"
        :comment="comment"
        :postID="postID"
        :onComment="loadComments"
      />
    </div>
  </div>
</template>

<style scoped>
</style>