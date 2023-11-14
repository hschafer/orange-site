<script setup>
import { ref } from 'vue'
import { readableTimestamp} from '../utils/utils.js'

defineProps({
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
                <a v-bind:href="url">{{ title }}</a> <span class="postURL">(<a class="domain" v-bind:href="url">{{ url }}</a>)</span>
            </div>
            <div class="postSecondRow">
                X points by {{ author }} {{ readableTimestamp(createdOn) }} | {{ numComments }} comment<span v-if="numComments != 1">s</span>
            </div>
            <!-- TODO user link and comments link -->
        </div>
    </div>
</template>

<style scoped>
    .post {
        padding-top: 0.25em;
        display: flex;
    }

    .post a {
        color: black;
    }

    .postRank, .postURL, .postSecondRow, .domain, .post a:visited {
        color: #828282
    }

    .postInfo {
        padding-left: 0.4em;
    }

    .postSecondRow {
        font-size: small;
    }
</style>