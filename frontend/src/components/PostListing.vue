<script setup>
import { ref } from 'vue'

defineProps({
  rank: Number,
  title: String,
  author: String,
  url: String,
})

// TODO compute a preview URL
</script>

<template>
    <div class="post">
        <div class="postRank">
            {{ rank }}. ^
        </div>
        <div class="postInfo">
            <div class="postFirstRow">
                {{ title }} <span class="postURL">(<a href="{{ url }}">{{ url }}</a>)</span>
            </div>
            <div class="postSecondRow">
                X points by {{ author }} Y time ago | Z comments
            </div>
        </div>
    </div>
</template>

<style scoped>
    .post {
        padding-top: 0.25em;
        display: flex;
    }

    .postRank, .postURL, .postSecondRow, .post a{
        color: #828282
    }

    .postInfo {
        padding-left: 0.4em;
    }

    .postSecondRow {
        font-size: small;
    }
</style>