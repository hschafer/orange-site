<script setup>
import { ref } from 'vue'
import { readableTimestamp} from '../utils/utils.js'

defineProps({
  id: Number,
  rank: Number,
  title: String,
  author: String,
  url: String,
  createdOn: String,
  numComments: Number,
  showRank: {
    type: Boolean,
    default: true
  }
})

// TODO compute a preview URL
</script>

<template>
    <div class="post">
        <div class="postRank">
            <span v-if="showRank">{{ rank }}. </span>^
        </div>
        <div class="postInfo">
            <div class="postFirstRow">
                <a class="postTitle" v-bind:href="url">{{ title }}</a> <span class="postURL">(<a class="domain" v-bind:href="url">{{ url }}</a>)</span>
            </div>
            <div class="postSecondRow">
                X points by {{ author }} {{ readableTimestamp(createdOn) }} |
                <router-link :to="`/post/${id}`">{{ numComments }} comment<span v-if="numComments != 1">s</span></router-link>
            </div>
            <!-- TODO user link -->
        </div>
    </div>
</template>

<style scoped>
    .post {
        padding-top: 0.25em;
        display: flex;
        color: #828282
    }

    .post a {
        color: #828282;
        text-decoration: none;
    }

    a.postTitle {
        color: black;
    }

    .postInfo {
        padding-left: 0.4em;
    }

    .postSecondRow {
        font-size: small;
    }
</style>