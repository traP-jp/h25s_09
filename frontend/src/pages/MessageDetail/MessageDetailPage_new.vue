<script lang="ts" setup>
import UserIcon from '@/components/UserIcon.vue'
import { useMessageDetail } from '@/lib/composables'
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

// URLパラメータからメッセージIDを取得
const messageId = computed(() => route.params.id as string)

// メッセージ詳細取得
const { data: message, isLoading, error } = useMessageDetail(messageId)

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

// ユーザー詳細ページへの遷移
const goToUserDetail = (traqId: string) => {
  router.push(`/users/${traqId}`)
}

// 戻る
const goBack = () => {
  router.back()
}
</script>

<template>
  <div :class="$style.messageDetail">
    <div :class="$style.container">
      <!-- ヘッダー -->
      <header :class="$style.header">
        <button :class="$style.backButton" @click="goBack">← 戻る</button>
        <h1 :class="$style.title">メッセージ詳細</h1>
      </header>

      <!-- ローディング状態 -->
      <div v-if="isLoading" :class="$style.loading">
        <div :class="$style.spinner"></div>
        <p>メッセージを読み込み中...</p>
      </div>

      <!-- エラー状態 -->
      <div v-else-if="error" :class="$style.error">
        <div :class="$style.errorIcon">⚠️</div>
        <h2>メッセージの取得に失敗しました</h2>
        <p>しばらく待ってから再度お試しください。</p>
      </div>

      <!-- メッセージが見つからない -->
      <div v-else-if="!message" :class="$style.notFound">
        <div :class="$style.errorIcon">📭</div>
        <h2>メッセージが見つかりません</h2>
        <p>指定されたメッセージは削除されたか、存在しません。</p>
      </div>

      <!-- メッセージ内容 -->
      <main v-else :class="$style.content">
        <!-- メインメッセージ -->
        <article :class="$style.mainMessage">
          <div :class="$style.messageHeader">
            <UserIcon
              :traq-id="message.author"
              size="lg"
              clickable
              @click="goToUserDetail(message.author)"
            />
            <div :class="$style.messageInfo">
              <button :class="$style.authorName" @click="goToUserDetail(message.author)">
                @{{ message.author }}
              </button>
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
            <div v-if="message.imageId" :class="$style.imageContainer">
              <img
                :src="`/api/images/${message.imageId}`"
                :alt="'添付画像'"
                :class="$style.messageImage"
                loading="lazy"
              />
            </div>
          </div>

          <!-- リアクション表示（読み取り専用） -->
          <div v-if="message.reactions.count > 0" :class="$style.reactions">
            <div
              :class="[$style.reactionItem, { [$style.myReaction]: message.reactions.myReaction }]"
            >
              <span :class="$style.emoji">👍</span>
              <span :class="$style.count">{{ message.reactions.count }}</span>
            </div>
          </div>
        </article>

        <!-- 返信一覧 -->
        <section v-if="message.replies && message.replies.length > 0" :class="$style.replies">
          <h2 :class="$style.repliesTitle">返信 ({{ message.replies.length }})</h2>
          <div :class="$style.repliesList">
            <article v-for="reply in message.replies" :key="reply.id" :class="$style.reply">
              <div :class="$style.messageHeader">
                <UserIcon
                  :traq-id="reply.author"
                  size="md"
                  clickable
                  @click="goToUserDetail(reply.author)"
                />
                <div :class="$style.messageInfo">
                  <button :class="$style.authorName" @click="goToUserDetail(reply.author)">
                    @{{ reply.author }}
                  </button>
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
                <div v-if="reply.images" :class="$style.imageContainer">
                  <img
                    :src="`/api/images/${reply.images}`"
                    :alt="'添付画像'"
                    :class="$style.messageImage"
                    loading="lazy"
                  />
                </div>
              </div>

              <!-- 返信のリアクション表示（読み取り専用） -->
              <div v-if="reply.reactions.count > 0" :class="$style.reactions">
                <div
                  :class="[
                    $style.reactionItem,
                    { [$style.myReaction]: reply.reactions.myReaction },
                  ]"
                >
                  <span :class="$style.emoji">👍</span>
                  <span :class="$style.count">{{ reply.reactions.count }}</span>
                </div>
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
  background: none;
  border: none;
  color: var(--color-primary-600);
  cursor: pointer;
  font-size: 0.875rem;
  padding: 0.5rem;
  border-radius: 0.25rem;
  transition: background-color 0.2s ease;

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

.messageInfo {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.authorName {
  background: none;
  border: none;
  padding: 0;
  font-size: 1rem;
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

/* レスポンシブ対応 */
@media (max-width: 768px) {
  .container {
    padding: 0 0.5rem;
  }

  .mainMessage,
  .reply {
    padding: 1rem;
  }

  .reply {
    margin-left: 0.5rem;
  }
}
</style>
