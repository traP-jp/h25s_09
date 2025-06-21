<script lang="ts" setup>
import type { Message } from '@/lib/apis/generated'
import { useAddReaction, useRemoveReaction } from '@/lib/composables'
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import UserIcon from './UserIcon.vue'

interface Props {
  /** メッセージデータ */
  message: Message
}

const props = defineProps<Props>()

const router = useRouter()

// リアクション機能
const addReactionMutation = useAddReaction()
const removeReactionMutation = useRemoveReaction()

// リアクション切り替え処理
const toggleReaction = async () => {
  try {
    if (props.message.reactions.myReaction) {
      await removeReactionMutation.mutateAsync(props.message.id)
    } else {
      await addReactionMutation.mutateAsync(props.message.id)
    }
  } catch (error) {
    console.error('Failed to toggle reaction:', error)
  }
}

// フォーマット済み作成日時
const formattedCreatedAt = computed(() => {
  const date = new Date(props.message.createdAt)
  const now = new Date()
  const diffInHours = (now.getTime() - date.getTime()) / (1000 * 60 * 60)

  if (diffInHours < 24) {
    return date.toLocaleTimeString('ja-JP', {
      hour: '2-digit',
      minute: '2-digit',
    })
  } else {
    return date.toLocaleDateString('ja-JP', {
      month: 'short',
      day: 'numeric',
    })
  }
})

// メッセージ詳細ページへの遷移
const goToDetail = () => {
  router.push(`/messages/${props.message.id}`)
}

// ユーザー詳細ページへの遷移
const goToUserDetail = (traqId: string) => {
  router.push(`/users/${traqId}`)
}
</script>

<template>
  <article :class="$style.messageItem">
    <div :class="$style.messageHeader">
      <UserIcon :traq-id="message.author" size="md" clickable @click="goToUserDetail" />
      <div :class="$style.messageInfo">
        <button :class="$style.authorName" @click="goToUserDetail(message.author)">
          @{{ message.author }}
        </button>
        <time :class="$style.timestamp" :datetime="message.createdAt">
          {{ formattedCreatedAt }}
        </time>
      </div>
    </div>

    <div :class="$style.messageContent">
      <p :class="$style.messageText">
        {{ message.content }}
      </p>

      <!-- 画像表示 -->
      <div v-if="message.imageId" :class="$style.imageContainer">
        <img
          :src="`/api/images/${message.imageId}`"
          :alt="'添付画像'"
          :class="$style.messageImage"
          loading="lazy"
        />
      </div>
    </div>

    <div :class="$style.messageActions">
      <!-- リアクションボタン -->
      <button
        :class="[
          $style.actionButton,
          $style.reactionButton,
          { [$style.active]: message.reactions.myReaction },
        ]"
        @click="toggleReaction"
        :disabled="addReactionMutation.isPending.value || removeReactionMutation.isPending.value"
      >
        <span :class="$style.emoji">👍</span>
        <span :class="$style.count">
          {{ message.reactions.count }}
        </span>
      </button>

      <!-- 返信ボタン -->
      <button :class="[$style.actionButton, $style.replyButton]" @click="goToDetail">
        <span :class="$style.icon">💬</span>
        <span v-if="message.replyCount > 0" :class="$style.count">
          {{ message.replyCount }}
        </span>
      </button>
    </div>
  </article>
</template>

<style lang="scss" module>
.messageItem {
  padding: 1rem;
  background-color: var(--color-surface);
  border-bottom: 1px solid var(--color-border-light);
  transition: background-color 0.2s ease;

  &:hover {
    background-color: var(--color-surface-variant);
  }
}

.messageHeader {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.5rem;
}

.messageInfo {
  display: flex;
  flex-direction: column;
  gap: 0.125rem;
}

.authorName {
  background: none;
  border: none;
  padding: 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--color-primary-600);
  cursor: pointer;
  text-align: left;

  &:hover {
    color: var(--color-primary-700);
    text-decoration: underline;
  }
}

.timestamp {
  font-size: 0.75rem;
  color: var(--color-text-secondary);
}

.messageContent {
  margin-bottom: 0.75rem;
}

.messageText {
  margin: 0;
  line-height: 1.5;
  color: var(--color-text-primary);
  white-space: pre-wrap;
  word-wrap: break-word;
}

.imageContainer {
  margin-top: 0.75rem;
}

.messageImage {
  max-width: 100%;
  max-height: 20rem;
  border-radius: 0.5rem;
  border: 1px solid var(--color-border-light);
}

.messageActions {
  display: flex;
  gap: 0.5rem;
  align-items: center;
}

.actionButton {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.25rem 0.5rem;
  background: none;
  border: 1px solid transparent;
  border-radius: 1rem;
  font-size: 0.75rem;
  color: var(--color-text-secondary);
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover {
    background-color: var(--color-surface-variant);
    border-color: var(--color-border-light);
  }

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
}

.reactionButton {
  &.active {
    background-color: var(--color-primary-50);
    border-color: var(--color-primary-200);
    color: var(--color-primary-700);

    [data-theme='dark'] & {
      background-color: var(--color-primary-900);
      border-color: var(--color-primary-800);
      color: var(--color-primary-300);
    }
  }
}

.emoji,
.icon {
  font-size: 1rem;
  line-height: 1;
}

.count {
  font-size: 0.75rem;
  font-weight: 500;
  min-width: 1rem;
  text-align: center;
}
</style>
