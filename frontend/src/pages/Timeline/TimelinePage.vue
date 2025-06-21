<script lang="ts" setup>
import MessageForm from '@/components/MessageForm.vue'
import MessageList from '@/components/MessageList.vue'
import { useGlobalError } from '@/lib/composables'

const { setError } = useGlobalError()

// メッセージ投稿成功時の処理
const handlePostSuccess = () => {
  // 投稿成功時の処理（必要に応じて通知等）
  console.log('Message posted successfully')
}

// メッセージ投稿エラー時の処理
const handlePostError = (error: Error) => {
  setError({
    message: error.message,
    status: 500,
  })
}
</script>

<template>
  <div :class="$style.timeline">
    <div :class="$style.container">
      <MessageForm @success="handlePostSuccess" @error="handlePostError" />
      <MessageList :include-replies="false" />
    </div>
  </div>
</template>

<style lang="scss" module>
.timeline {
  max-width: 800px;
  margin: 0 auto;
  width: 100%;
}

.container {
  display: flex;
  flex-direction: column;
  gap: 0; /* MessageFormがmargin-bottomを持っているため */
}

@media (max-width: 640px) {
  .timeline {
    max-width: 100%;
  }
}
</style>
