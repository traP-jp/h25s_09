<script lang="ts" setup>
import MessageList from '@/components/MessageList.vue'
import { useInfiniteUserMessages } from '@/lib/composables'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

// URLパラメータからユーザーIDを取得
const userId = computed(() => route.params.traqId as string)

// ユーザーメッセージ取得（infinite scroll対応）
const { messages, isLoading, error } = useInfiniteUserMessages(userId, {
  includeReplies: false,
})
</script>

<template>
  <div :class="$style.userMessagesPage">
    <MessageList
      :messages="messages"
      :loading="isLoading"
      :error="error"
      :include-replies="false"
      :use-infinite-scroll="true"
    />
  </div>
</template>

<style lang="scss" module>
.userMessagesPage {
  padding: 1rem;
  min-height: 100%;
}
</style>
