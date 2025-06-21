<script lang="ts" setup>
import type { Message } from '@/lib/apis/generated'
import { useAddReaction, useRemoveReaction } from '@/lib/composables'
import { formatFullDateTime, formatRelativeTime } from '@/lib/utils/format'
import { Icon } from '@iconify/vue'
import { computed } from 'vue'
import { RouterLink } from 'vue-router'
import UserIcon from './UserIcon.vue'

interface Props {
  /** メッセージデータ */
  message: Message
  /** 返信メッセージかどうか */
  isReply?: boolean
}

const props = defineProps<Props>()

// リアクション機能
const addReactionMutation = useAddReaction()
const removeReactionMutation = useRemoveReaction()

// リアクション切り替え処理
const toggleReaction = async () => {
  try {
    // reactionsプロパティの存在チェック
    if (!props.message.reactions || typeof props.message.reactions.myReaction !== 'boolean') {
      console.error('Invalid reactions data:', props.message.reactions)
      return
    }

    console.log(
      'Toggling reaction for message:',
      props.message.id,
      'current myReaction:',
      props.message.reactions.myReaction,
    )

    if (props.message.reactions.myReaction) {
      console.log('Removing reaction...')
      await removeReactionMutation.mutateAsync(props.message.id)
    } else {
      console.log('Adding reaction...')
      await addReactionMutation.mutateAsync(props.message.id)
    }

    console.log('Reaction toggle successful')
  } catch (error) {
    console.error('Failed to toggle reaction:', error)
    // エラーの詳細をログに出力
    if (error instanceof Error) {
      console.error('Error message:', error.message)
      console.error('Error stack:', error.stack)
    }
  }
}

// フォーマット済み作成日時
const formattedCreatedAt = computed(() => formatRelativeTime(props.message.createdAt))
const fullDateTime = computed(() => formatFullDateTime(props.message.createdAt))

// 画像読み込みエラーハンドリング
const onImageError = (event: Event) => {
  const img = event.target as HTMLImageElement
  img.style.display = 'none'
  console.warn('Failed to load image:', img.src)
}
</script>

<template>
  <article :class="[$style.messageItem, { [$style.reply]: isReply }]" role="article">
    <div v-if="isReply" :class="$style.replyIndicator">
      <Icon icon="mdi:reply" :class="$style.replyIcon" />
    </div>
    <RouterLink
      :to="`/messages/${message.id}`"
      :class="$style.messageLink"
      :aria-label="`${message.author}のメッセージ詳細を表示: ${message.content.slice(0, 50)}${message.content.length > 50 ? '...' : ''}`"
    >
      <div :class="$style.messageHeader">
        <RouterLink
          :to="`/users/${message.author}`"
          :class="$style.userIconLink"
          :aria-label="`${message.author}のプロフィールを表示`"
          @click.stop
        >
          <UserIcon
            :traq-id="message.author"
            size="md"
            :aria-label="`${message.author}のプロフィール`"
          />
        </RouterLink>
        <div :class="$style.messageInfo">
          <RouterLink
            :to="`/users/${message.author}`"
            :class="$style.authorName"
            :aria-label="`${message.author}のプロフィールを表示`"
            @click.stop
          >
            @{{ message.author }}
          </RouterLink>
          <time :class="$style.timestamp" :datetime="message.createdAt" :title="fullDateTime">
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
            @error="onImageError"
          />
        </div>
      </div>

      <div :class="$style.messageActions" role="group" aria-label="メッセージの操作">
        <!-- リアクションボタン -->
        <button
          v-if="message.reactions"
          :class="[
            $style.actionButton,
            $style.reactionButton,
            { [$style.active]: message.reactions.myReaction },
            {
              [$style.error]:
                addReactionMutation.isError.value || removeReactionMutation.isError.value,
            },
          ]"
          @click.stop.prevent="toggleReaction"
          :disabled="addReactionMutation.isPending.value || removeReactionMutation.isPending.value"
          :aria-label="`${message.reactions.myReaction ? 'いいねを取り消す' : 'いいねする'} (現在 ${message.reactions.count} 件)`"
          :aria-pressed="message.reactions.myReaction"
          :title="
            addReactionMutation.isError.value || removeReactionMutation.isError.value
              ? 'リアクションでエラーが発生しました'
              : undefined
          "
        >
          <Icon icon="mdi:heart" :class="$style.emoji" aria-hidden="true" />
          <span :class="$style.count" aria-label="いいね数">
            {{ message.reactions.count }}
          </span>
        </button>

        <!-- 返信ボタン -->
        <RouterLink
          :to="`/messages/${message.id}`"
          :class="[$style.actionButton, $style.replyButton]"
          :aria-label="`返信する${message.replyCount > 0 ? ` (${message.replyCount} 件の返信)` : ''}`"
          @click.stop
        >
          <Icon icon="mdi:reply" :class="$style.icon" aria-hidden="true" />
          <span v-if="message.replyCount > 0" :class="$style.count" aria-label="返信数">
            {{ message.replyCount }}
          </span>
        </RouterLink>
      </div>
    </RouterLink>
  </article>
</template>

<style lang="scss" module>
.messageItem {
  background-color: var(--color-surface);
  border: 1px solid var(--color-border-light);
  border-radius: 0.5rem;
  position: relative;

  &.reply {
    margin-left: 2rem;
    background-color: var(--color-surface-variant);
    border-left: 3px solid var(--color-primary-300);

    [data-theme='dark'] & {
      background-color: var(--color-background-soft);
      border-left-color: var(--color-primary-600);
    }
  }
}

.replyIndicator {
  position: absolute;
  top: 0.5rem;
  left: -1.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 1.5rem;
  height: 1.5rem;
  background-color: var(--color-primary-100);
  border-radius: 50%;

  [data-theme='dark'] & {
    background-color: var(--color-primary-800);
  }
}

.replyIcon {
  font-size: 0.75rem;
  color: var(--color-primary-600);

  [data-theme='dark'] & {
    color: var(--color-primary-300);
  }
}

.messageLink {
  display: block;
  padding: 1rem;
  text-decoration: none;
  color: inherit;
  transition: all 0.2s ease;

  &:hover {
    background-color: var(--color-surface-variant);
    box-shadow: 0 2px 8px var(--color-shadow-light);
    transform: translateY(-1px);
    border-color: var(--color-border-medium);
  }

  &:focus {
    outline: 2px solid var(--color-primary-500);
    outline-offset: 2px;
    background-color: var(--color-surface-variant);
  }

  &:active {
    transform: translateY(0);
    box-shadow: 0 1px 4px var(--color-shadow-light);
  }
}

.messageHeader {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.5rem;
}

.userIconLink {
  display: flex;
  align-items: center;
  border-radius: 50%;
  transition: transform 0.2s ease;

  &:hover {
    transform: scale(1.05);
  }

  &:focus {
    outline: 2px solid var(--color-primary-500);
    outline-offset: 2px;
  }
}

.messageInfo {
  display: flex;
  flex-direction: column;
  gap: 0.125rem;
}

.authorName {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--color-primary-600);
  text-decoration: none;

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
  text-decoration: none;
  position: relative;
  z-index: 1;

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

  &.error {
    background-color: var(--color-error-50);
    border-color: var(--color-error-200);
    color: var(--color-error-700);

    [data-theme='dark'] & {
      background-color: var(--color-error-900);
      border-color: var(--color-error-800);
      color: var(--color-error-300);
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
