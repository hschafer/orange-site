<script setup>
import { ref } from 'vue';
import { readableTimestamp} from '../utils/utils.js'
import CommentBox from './CommentBox.vue';

defineProps({
    comment: Object,
    postID: Number,
    onComment: Function
})

const showCommentBox = ref(false);

function toggleShowComment() {
    showCommentBox.value = !showCommentBox.value;
}
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

            <div v-if="showCommentBox">
                <CommentBox
                    :postID="postID"
                    :onComment="onComment"
                    :parentID="comment.CommentID"
                />
            </div>

            <span class="clickable" @click="toggleShowComment">{{ showCommentBox ? 'close' : 'reply' }}</span>

            <div class="children" v-if="comment.Children">
                <Comment
                    v-for="(child, _) in comment.Children"
                    :comment="child"
                    :postID="postID"
                    :onComment="onComment"
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

    .clickable {
        color: black;
        cursor: pointer;
        text-decoration: underline;
        font-size: small;
    }
</style>