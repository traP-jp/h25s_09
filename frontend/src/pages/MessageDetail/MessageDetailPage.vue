<script lang="ts" setup>
import ReplyForm from '@/components/ReplyForm.vue'
import UserIcon from '@/components/UserIcon.vue'
import { useAddReaction, useMessageDetail, useRemoveReaction } from '@/lib/composables'
import { Icon } from '@iconify/vue'
import { computed, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()

// URLパラメータからメッセージIDを取得
const messageId = computed(() => route.params.id as string)

// メッセージ詳細取得
const { data: message, isLoading, error, refetch } = useMessageDetail(messageId)

// リプライフォーム表示状態
const showReplyForm = ref(false)

// リアクション機能
const addReactionMutation = useAddReaction()
const removeReactionMutation = useRemoveReaction()

// メインメッセージのリアクション切り替え処理
const toggleMainReaction = async () => {
  if (!message.value) return

  try {
    if (message.value.reactions.myReaction) {
      await removeReactionMutation.mutateAsync(message.value.id)
    } else {
      await addReactionMutation.mutateAsync(message.value.id)
    }
  } catch (error) {
    console.error('Failed to toggle reaction:', error)
  }
}

// 返信のリアクション切り替え処理
const toggleReplyReaction = async (replyId: string, myReaction: boolean) => {
  try {
    if (myReaction) {
      await removeReactionMutation.mutateAsync(replyId)
    } else {
      await addReactionMutation.mutateAsync(replyId)
    }
  } catch (error) {
    console.error('Failed to toggle reply reaction:', error)
  }
}

// リプライ成功時の処理
const handleReplySuccess = () => {
  showReplyForm.value = false
  // メッセージ詳細を再取得してリプライを反映
  refetch()
}

// リプライエラー時の処理
const handleReplyError = (error: Error) => {
  console.error('Failed to create reply:', error)
  // エラーハンドリング（トーストなどでユーザーに通知）
}

// フォーマット済み作成日時
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('ja-JP', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}
</script>

<template>
  <div :class="$style.messageDetail">
    <div :class="$style.container">
      <!-- ヘッダー -->
      <header :class="$style.header">
        <RouterLink to="/timeline" :class="$style.backButton">
          <Icon icon="mdi:arrow-left" />
          戻る
        </RouterLink>
        <h1 :class="$style.title">メッセージ詳細</h1>
      </header>

      <!-- ローディング状態 -->
      <div v-if="isLoading" :class="$style.loading">
        <div :class="$style.spinner"></div>
        <p>メッセージを読み込み中...</p>
      </div>

      <!-- エラー状態 -->
      <div v-else-if="error" :class="$style.error">
        <Icon icon="mdi:alert-circle" :class="$style.errorIcon" />
        <h2>メッセージの取得に失敗しました</h2>
        <p>しばらく待ってから再度お試しください。</p>
      </div>

      <!-- メッセージが見つからない -->
      <div v-else-if="!message" :class="$style.notFound">
        <Icon icon="mdi:message-off-outline" :class="$style.errorIcon" />
        <h2>メッセージが見つかりません</h2>
        <p>指定されたメッセージは削除されたか、存在しません。</p>
      </div>

      <!-- メッセージ内容 -->
      <main v-else :class="$style.content">
        <!-- メインメッセージ -->
        <article :class="$style.mainMessage">
          <div :class="$style.messageHeader">
            <RouterLink :to="`/users/${message.author}`" :class="$style.userIconLink">
              <UserIcon :traq-id="message.author" size="lg" />
            </RouterLink>
            <div :class="$style.messageInfo">
              <RouterLink :to="`/users/${message.author}`" :class="$style.authorName">
                @{{ message.author }}
              </RouterLink>
              <time :class="$style.timestamp" :datetime="message.createdAt">
                {{ formatDate(message.createdAt) }}
              </time>
            </div>
          </div>

          <div :class="$style.messageContent">
            <p :class="$style.messageText">
              {{ message.content }}
            </p>

            <!-- 画像表示 -->
            <div
              v-if="message.imageId && message.imageId !== '00000000-0000-0000-0000-000000000000'"
              :class="$style.imageContainer"
            >
              <img
                :src="`/api/images/${message.imageId}`"
                :alt="'添付画像'"
                :class="$style.messageImage"
                loading="lazy"
              />
            </div>
          </div>

          <!-- リアクションボタン -->
          <div :class="$style.messageActions">
            <button
              v-if="message.reactions"
              :class="[$style.reactionButton, { [$style.active]: message.reactions.myReaction }]"
              @click="toggleMainReaction"
              :disabled="
                addReactionMutation.isPending.value || removeReactionMutation.isPending.value
              "
              :aria-label="`${message.reactions.myReaction ? 'いいねを取り消す' : 'いいねする'} (現在 ${message.reactions.count} 件)`"
            >
              <Icon icon="mdi:heart" :class="$style.emoji" />
              <span :class="$style.count">{{ message.reactions.count }}</span>
            </button>

            <!-- リプライボタン -->
            <button
              :class="[$style.replyToggleButton, { [$style.active]: showReplyForm }]"
              @click="showReplyForm = !showReplyForm"
              :aria-label="showReplyForm ? 'リプライを閉じる' : 'リプライする'"
            >
              <Icon icon="mdi:reply" :class="$style.icon" />
              <span>リプライ</span>
            </button>
          </div>
        </article>

        <!-- リプライフォーム -->
        <section v-if="showReplyForm" :class="$style.replyFormSection">
          <h3 :class="$style.replyFormTitle">
            <Icon icon="mdi:reply" :class="$style.replyIcon" />
            返信を書く
          </h3>
          <ReplyForm
            :replies-to="message.id"
            :auto-focus="true"
            @success="handleReplySuccess"
            @error="handleReplyError"
            @cancel="showReplyForm = false"
          />
        </section>

        <!-- 返信一覧 -->
        <section v-if="message.replies && message.replies.length > 0" :class="$style.replies">
          <h2 :class="$style.repliesTitle">返信 ({{ message.replies.length }})</h2>
          <div :class="$style.repliesList">
            <article v-for="reply in message.replies" :key="reply.id" :class="$style.reply">
              <div :class="$style.messageHeader">
                <RouterLink :to="`/users/${reply.author}`" :class="$style.userIconLink">
                  <UserIcon :traq-id="reply.author" size="md" />
                </RouterLink>
                <div :class="$style.messageInfo">
                  <RouterLink :to="`/users/${reply.author}`" :class="$style.authorName">
                    @{{ reply.author }}
                  </RouterLink>
                  <time :class="$style.timestamp" :datetime="reply.createdAt">
                    {{ formatDate(reply.createdAt) }}
                  </time>
                </div>
              </div>

              <div :class="$style.messageContent">
                <p :class="$style.messageText">
                  {{ reply.content }}
                </p>

                <!-- 返信の画像表示 -->
                <div
                  v-if="reply.images && reply.images !== '00000000-0000-0000-0000-000000000000'"
                  :class="$style.imageContainer"
                >
                  <img
                    :src="`/api/images/${reply.imageId}`"
                    :alt="'添付画像'"
                    :class="$style.messageImage"
                    loading="lazy"
                  />
                </div>
              </div>

              <!-- 返信のリアクションボタン -->
              <div :class="$style.messageActions">
                <button
                  v-if="reply.reactions"
                  :class="[$style.reactionButton, { [$style.active]: reply.reactions.myReaction }]"
                  @click="toggleReplyReaction(reply.id, reply.reactions.myReaction)"
                  :disabled="
                    addReactionMutation.isPending.value || removeReactionMutation.isPending.value
                  "
                  :aria-label="`${reply.reactions.myReaction ? 'いいねを取り消す' : 'いいねする'} (現在 ${reply.reactions.count} 件)`"
                >
                  <Icon icon="mdi:heart" :class="$style.emoji" />
                  <span :class="$style.count">{{ reply.reactions.count }}</span>
                </button>
              </div>
            </article>
          </div>
        </section>
      </main>
    </div>
  </div>
</template>

<style lang="scss" module>
.messageDetail {
  min-height: 100vh;
  background-color: var(--color-background);
}

.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 0 1rem;
}

.header {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem 0;
  border-bottom: 1px solid var(--color-border-light);
  margin-bottom: 2rem;
}

.backButton {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  color: var(--color-primary-600);
  font-size: 0.875rem;
  padding: 0.5rem;
  border-radius: 0.25rem;
  transition: background-color 0.2s ease;
  text-decoration: none;

  &:hover {
    background-color: var(--color-primary-50);
  }
}

.title {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-text-primary);
}

.loading,
.error,
.notFound {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 4rem 2rem;
  text-align: center;
  color: var(--color-text-secondary);
}

.loading {
  gap: 1rem;
}

.error,
.notFound {
  gap: 0.5rem;
}

.spinner {
  width: 2rem;
  height: 2rem;
  border: 2px solid var(--color-border-light);
  border-top: 2px solid var(--color-primary-500);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.errorIcon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.error h2,
.notFound h2 {
  margin: 0 0 0.5rem 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-text-primary);
}

.error p,
.notFound p {
  margin: 0;
  color: var(--color-text-secondary);
}

.content {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.mainMessage,
.reply {
  background-color: var(--color-surface);
  border-radius: 0.75rem;
  padding: 1.5rem;
  box-shadow: 0 1px 3px var(--color-shadow-light);
}

.messageHeader {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 1rem;
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
  gap: 0.25rem;
}

.authorName {
  font-size: 1rem;
  font-weight: 600;
  color: var(--color-primary-600);
  text-decoration: none;

  &:hover {
    color: var(--color-primary-700);
    text-decoration: underline;
  }
}

.timestamp {
  font-size: 0.875rem;
  color: var(--color-text-secondary);
}

.messageContent {
  margin-bottom: 1rem;
}

.messageText {
  margin: 0;
  line-height: 1.6;
  color: var(--color-text-primary);
  white-space: pre-wrap;
  word-wrap: break-word;
}

.imageContainer {
  margin-top: 1rem;
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
  margin-top: 1rem;
}

.reactionButton {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.5rem 0.75rem;
  background-color: var(--color-surface-variant);
  border: 1px solid var(--color-border-light);
  border-radius: 1rem;
  font-size: 0.875rem;
  color: var(--color-text-secondary);
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover {
    background-color: var(--color-primary-50);
    border-color: var(--color-primary-200);
    color: var(--color-primary-600);
  }

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

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

.reactions {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}

.reactionItem {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.25rem 0.5rem;
  background-color: var(--color-surface-variant);
  border: 1px solid var(--color-border-light);
  border-radius: 1rem;
  font-size: 0.75rem;

  &.myReaction {
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

.emoji {
  font-size: 1rem;
  line-height: 1;
}

.count {
  font-weight: 500;
  min-width: 1rem;
  text-align: center;
}

.replies {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.repliesTitle {
  margin: 0 0 1rem 0;
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--color-text-primary);
}

.repliesList {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.reply {
  margin-left: 1rem;
  border-left: 3px solid var(--color-primary-200);
}

.replyToggleButton {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.5rem 0.75rem;
  background-color: var(--color-surface-variant);
  border: 1px solid var(--color-border-light);
  border-radius: 1rem;
  font-size: 0.875rem;
  color: var(--color-text-secondary);
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover {
    background-color: var(--color-primary-50);
    border-color: var(--color-primary-200);
    color: var(--color-primary-600);
  }

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

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

.replyFormSection {
  background-color: var(--color-surface);
  border-radius: 0.75rem;
  padding: 1.5rem;
  box-shadow: 0 1px 3px var(--color-shadow-light);
  border: 1px solid var(--color-border-light);
}

.replyFormTitle {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin: 0 0 1rem 0;
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--color-text-primary);
}

.replyIcon {
  font-size: 1.25rem;
  color: var(--color-primary-600);
}

.icon {
  font-size: 1rem;
  line-height: 1;
}

.replyFormSection {
  margin-top: 2rem;
}

/* レスポンシブ対応 */
@media (max-width: 768px) {
  .container {
    padding: 0 0.5rem;
  }

  .mainMessage,
  .reply,
  .replyFormSection {
    padding: 1rem;
  }

  .reply {
    margin-left: 0.5rem;
  }

  .messageActions {
    flex-wrap: wrap;
    gap: 0.375rem;
  }

  .replyToggleButton {
    font-size: 0.8rem;
    padding: 0.375rem 0.625rem;
  }

  .replyFormTitle {
    font-size: 1rem;
  }
}
</style>
