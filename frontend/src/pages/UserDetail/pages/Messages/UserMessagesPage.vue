<script lang="ts" setup>
import MessageList from '@/components/MessageList.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

// URLパラメータからユーザーIDを取得
const userId = computed(() => route.params.traqId as string)
</script>

<template>
  <div :class="$style.userMessagesPage">
    <MessageList :include-replies="false" :use-infinite-scroll="true" :user-id="userId" />
  </div>
</template>

<style lang="scss" module>
.userMessagesPage {
  padding: 1rem;
  min-height: 100%;
}
</style>
