# 生成されたAPI呼び出しコードの使用方法

## 概要

OpenAPIの定義から自動生成されたTypeScript APIクライアントの使用方法を説明します。

## 基本的な使い方

### 1. APIクライアントのインポート

```typescript
import { apiClient } from '@/lib/apis/client'
// または個別にインポート
import { MessagesApi, AchievementsApi } from '@/lib/apis/generated'
```

### 2. メッセージAPI

#### メッセージ一覧の取得

```typescript
// 基本的な使用方法
const messages = await apiClient.messages.messagesGet()

// パラメータ付きでの使用
const messages = await apiClient.messages.messagesGet(
  20, // limit
  0, // offset
  'tidus', // traqId (特定ユーザーのメッセージのみ)
)
```

#### メッセージの投稿

```typescript
// テキストのみ
await apiClient.messages.messagesPost('こんにちは！')

// 画像付き
const file = // FileオブジェクトまたはBlob
  await apiClient.messages.messagesPost(
    '画像付きメッセージです',
    undefined, // repliesTo (返信先ID)
    file, // image
  )

// 返信メッセージ
await apiClient.messages.messagesPost(
  'これは返信です',
  'parent-message-id', // 返信先のメッセージID
  undefined, // image
)
```

#### 特定メッセージの詳細取得

```typescript
const messageDetail = await apiClient.messages.messagesIdGet('message-id')
```

### 3. 実績API

#### 実績一覧の取得

```typescript
// 全ユーザーの実績
const achievements = await apiClient.achievements.achievementsGet()

// 特定ユーザーの実績
const userAchievements = await apiClient.achievements.achievementsGet('tidus')
```

#### 実績イベントの試行

```typescript
const result = await apiClient.achievements.tryAchieveIdPost('event-id')
console.log(result.data.dispatched) // イベントが発火したかどうか
```

### 4. 画像API

#### 画像の取得（主にimgタグのsrcで使用）

```typescript
// 直接URLとして使用
const imageUrl = `/api/images/${imageId}`

// または、Blobとして取得
const imageResponse = await apiClient.images.imagesIdGet(imageId)
const imageBlob = imageResponse.data
```

### 5. ユーザーAPI

#### 自分の情報取得

```typescript
const userInfo = await apiClient.user.meGet()
console.log(userInfo.data.traqId) // 自分のtraqID
```

## Vue.jsでの実用例

### Composition API使用例

```vue
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { apiClient } from '@/lib/apis/client'
import type { Message, Achievement } from '@/lib/apis/generated'

const messages = ref<Message[]>([])
const achievements = ref<Achievement[]>([])
const user = ref<string>('')

// データ取得
const fetchData = async () => {
  try {
    // 並列でデータを取得
    const [messagesRes, achievementsRes, userRes] = await Promise.all([
      apiClient.messages.messagesGet(10),
      apiClient.achievements.achievementsGet(),
      apiClient.user.meGet(),
    ])

    messages.value = messagesRes.data
    achievements.value = achievementsRes.data
    user.value = userRes.data.traqId
  } catch (error) {
    console.error('データ取得エラー:', error)
  }
}

onMounted(fetchData)
</script>
```

### Pinia Storeでの使用例

```typescript
// stores/messages.ts
import { defineStore } from 'pinia'
import { apiClient } from '@/lib/apis/client'
import type { Message } from '@/lib/apis/generated'

export const useMessagesStore = defineStore('messages', () => {
  const messages = ref<Message[]>([])
  const loading = ref(false)

  const fetchMessages = async () => {
    loading.value = true
    try {
      const response = await apiClient.messages.messagesGet()
      messages.value = response.data
    } catch (error) {
      console.error('メッセージ取得エラー:', error)
    } finally {
      loading.value = false
    }
  }

  const postMessage = async (content: string, image?: File) => {
    try {
      await apiClient.messages.messagesPost(content, undefined, image)
      await fetchMessages() // 再取得
    } catch (error) {
      console.error('メッセージ投稿エラー:', error)
      throw error
    }
  }

  return {
    messages: readonly(messages),
    loading: readonly(loading),
    fetchMessages,
    postMessage,
  }
})
```

## エラーハンドリング

```typescript
import { AxiosError } from 'axios'

try {
  const response = await apiClient.messages.messagesGet()
} catch (error) {
  if (error instanceof AxiosError) {
    switch (error.response?.status) {
      case 404:
        console.error('リソースが見つかりません')
        break
      case 400:
        console.error('リクエストが不正です')
        break
      case 500:
        console.error('サーバーエラーが発生しました')
        break
      default:
        console.error('予期しないエラー:', error.message)
    }
  } else {
    console.error('ネットワークエラー:', error)
  }
}
```

## カスタム設定

### 認証トークンの設定

```typescript
import { Configuration } from '@/lib/apis/generated'

const configurationWithAuth = new Configuration({
  basePath: '/api',
  accessToken: 'your-jwt-token',
  // または
  apiKey: 'your-api-key',
})
```

### インターセプターの追加

```typescript
import axios from 'axios'

// リクエストインターセプター
axios.interceptors.request.use(
  (config) => {
    // 認証トークンを自動付与
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error),
)

// レスポンスインターセプター
axios.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // 認証エラー時の処理
      window.location.href = '/login'
    }
    return Promise.reject(error)
  },
)
```

## 型安全性の活用

生成されたコードは完全に型安全です：

```typescript
import type { Message, MessageDetail, Achievement } from '@/lib/apis/generated'

// 型チェックされるため、プロパティの間違いやtypoを防げます
const handleMessage = (message: Message) => {
  console.log(message.id) // ✅ 正しい
  console.log(message.author) // ✅ 正しい
  console.log(message.content) // ✅ 正しい
  console.log(message.title) // ❌ コンパイルエラー（titleプロパティは存在しない）
}
```

## 注意点

1. **APIの再生成**: OpenAPI定義を変更した場合は、`npm run gen-api` を実行してクライアントコードを再生成してください
2. **型の更新**: 生成された型定義は自動的に更新されるため、既存のコードでTypeScriptエラーが発生する可能性があります
3. **デバッグ**: 生成されたコードは自動生成のため、直接編集せず、OpenAPI定義を修正してください
