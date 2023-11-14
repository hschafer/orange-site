<script setup>
    defineProps({
      comment: Object
    })
</script>
<template>
    <p>{{ comment.Comment }}</p>

    <div v-if="comment.Children">
        <Comment
            v-for="(child, _) in comment.Children"
            :comment="child"
        />
    </div>
</template>

<style scoped>
</style>