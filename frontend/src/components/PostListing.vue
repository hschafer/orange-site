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
        <div class="postFirstRow">
            <span class="postRank">{{ rank }}. ^</span> {{ title }} <span class="postURL">(<a href="{{ url }}">{{ url }}</a>)</span>
        </div>
        <div class="postSecondRow">
            X points by {{ author }} Y time ago | Z comments
        </div>
    </div>
</template>

<style scoped>
    .post {
        padding-top: 0.25em;
    }

    .postRank, .postURL, .postSecondRow, .post a{
        color: #828282
    }

    .postSecondRow {
        font-size: small;
    }
</style>