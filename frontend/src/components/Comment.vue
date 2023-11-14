<script setup>
import { readableTimestamp} from '../utils/utils.js'

defineProps({
    comment: Object,
})
</script>

<template>
    <div class="comment">
        <div class="upvote">^</div>
        <div class="commentWrapper">
            <div class="commentMetadata">
                {{ comment.Username }} posted {{ readableTimestamp(comment.CreatedOn) }}
            </div>
            <p>{{ comment.Comment }}</p>

            <!-- TODO link to comment page-->
            <router-link to="/">reply</router-link>

            <div class="children" v-if="comment.Children">
                <Comment
                    v-for="(child, _) in comment.Children"
                    :comment="child"
                />
            </div>
        </div>
    </div>

</template>

<style scoped>
    .comment {
        display: flex;
        align-items: flex-start;
        margin-top: 0.75em;
    }

    .comment p {
        margin-top: 0.1em;
        margin-bottom: 0.1em;
    }

    .comment a {
        font-size: small;
    }

    .upvote {
        margin-right: 0.4em;
    }

    .upvote, .commentMetadata {
        color: #828282;
    }

    .commentMetadata {
        font-size: small;
    }

    .children {
        padding-left: 2em;
        border-left: 1px solid #828282;
    }
</style>