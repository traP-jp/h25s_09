import { http, HttpResponse } from 'msw'
import { delay, mswLog } from './utils'

// 型定義
interface Message {
  id: string
  author: {
    traqId: string
  }
  content: string
  images: string | null
  reactions: {
    count: number
    myReaction: boolean
  }
  createdAt: string
}

interface MessageDetail extends Message {
  replies: Message[]
}

// モックデータ
const mockMessages: Message[] = [
  {
    id: '550e8400-e29b-41d4-a716-446655440001',
    author: {
      traqId: 'testuser1',
    },
    content: 'こんにちは！これは最初のテストメッセージです。',
    images: null,
    reactions: {
      count: 3,
      myReaction: true,
    },
    createdAt: '2025-06-21T10:00:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440002',
    author: {
      traqId: 'testuser2',
    },
    content: 'こんにちは！画像付きのメッセージです。',
    images: '550e8400-e29b-41d4-a716-446655440010',
    reactions: {
      count: 1,
      myReaction: false,
    },
    createdAt: '2025-06-21T10:30:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440003',
    author: {
      traqId: 'testuser3',
    },
    content: 'これは返信メッセージの例です。',
    images: null,
    reactions: {
      count: 0,
      myReaction: false,
    },
    createdAt: '2025-06-21T11:00:00Z',
  },
]

const mockMessageDetails: MessageDetail[] = [
  {
    id: '550e8400-e29b-41d4-a716-446655440001',
    author: {
      traqId: 'testuser1',
    },
    content: 'こんにちは！これは最初のテストメッセージです。',
    images: null,
    reactions: {
      count: 3,
      myReaction: true,
    },
    createdAt: '2025-06-21T10:00:00Z',
    replies: [
      {
        id: '550e8400-e29b-41d4-a716-446655440004',
        author: {
          traqId: 'testuser2',
        },
        content: 'このメッセージへの返信です！',
        images: null,
        reactions: {
          count: 1,
          myReaction: false,
        },
        createdAt: '2025-06-21T10:15:00Z',
      },
    ],
  },
]

const mockAchievements = [
  {
    id: 1,
    name: '初回投稿',
    achievedAt: '2025-06-21T09:00:00Z',
  },
  {
    id: 2,
    name: '10回投稿',
    achievedAt: '2025-06-21T10:00:00Z',
  },
  {
    id: 3,
    name: '初回リアクション',
    achievedAt: '2025-06-21T10:30:00Z',
  },
]

export const handlers = [
  // メッセージ一覧の取得
  http.get('/api/messages', async ({ request }) => {
    mswLog('info', 'GET /api/messages リクエストを受信')

    // リアルなAPI応答を模擬するため少し遅延を追加
    await delay(300)

    const url = new URL(request.url)
    const limit = Number(url.searchParams.get('limit')) || 20
    const offset = Number(url.searchParams.get('offset')) || 0
    const traqId = url.searchParams.get('traqId')

    let filteredMessages = mockMessages

    if (traqId) {
      filteredMessages = mockMessages.filter((message) => message.author.traqId === traqId)
      mswLog('info', `ユーザー ${traqId} のメッセージをフィルタリング`)
    }

    const paginatedMessages = filteredMessages.slice(offset, offset + limit)

    mswLog('info', `${paginatedMessages.length}件のメッセージを返却`)
    return HttpResponse.json(paginatedMessages)
  }),

  // メッセージの投稿
  http.post('/api/messages', async ({ request }) => {
    const formData = await request.formData()
    const message = formData.get('message') as string
    const image = formData.get('image') as File

    if (!message || message.trim() === '') {
      return HttpResponse.json({ message: 'メッセージ本文が空です' }, { status: 400 })
    }

    const newMessage: Message = {
      id: `550e8400-e29b-41d4-a716-${Date.now()}`,
      author: {
        traqId: 'currentuser',
      },
      content: message,
      images: image ? `image-${Date.now()}` : null,
      reactions: {
        count: 0,
        myReaction: false,
      },
      createdAt: new Date().toISOString(),
    }

    // モックデータに追加
    mockMessages.unshift(newMessage)

    return HttpResponse.json(newMessage, { status: 201 })
  }),

  // メッセージ詳細の取得
  http.get('/api/messages/:id', ({ params }) => {
    const { id } = params
    const messageDetail = mockMessageDetails.find((msg) => msg.id === id)

    if (!messageDetail) {
      return HttpResponse.json({ message: 'メッセージが見つかりません' }, { status: 404 })
    }

    return HttpResponse.json(messageDetail)
  }),

  // リアクションの追加/削除
  http.post('/api/messages/:id/reactions', ({ params }) => {
    const { id } = params
    const message = mockMessages.find((msg) => msg.id === id)

    if (!message) {
      return HttpResponse.json({ message: 'メッセージが見つかりません' }, { status: 404 })
    }

    // リアクションの状態を切り替え
    if (message.reactions.myReaction) {
      message.reactions.count -= 1
      message.reactions.myReaction = false
    } else {
      message.reactions.count += 1
      message.reactions.myReaction = true
    }

    return HttpResponse.json({ success: true })
  }),

  // 画像のアップロード
  http.post('/api/images', async ({ request }) => {
    const formData = await request.formData()
    const image = formData.get('image') as File

    if (!image) {
      return HttpResponse.json({ message: '画像が指定されていません' }, { status: 400 })
    }

    return HttpResponse.json({ imageId: `image-${Date.now()}` }, { status: 201 })
  }),

  // 画像の取得
  http.get('/api/images/:id', ({ params }) => {
    const { id } = params

    // 実際のプロジェクトでは、適切な画像データを返す
    return HttpResponse.json({ url: `https://via.placeholder.com/400x300?text=Image+${id}` })
  }),

  // 実績一覧の取得
  http.get('/api/achievements', ({ request }) => {
    const url = new URL(request.url)
    const traqId = url.searchParams.get('traqId')

    if (traqId && traqId !== 'currentuser') {
      // 他のユーザーの実績を返す（少なめ）
      return HttpResponse.json([mockAchievements[0]])
    }

    return HttpResponse.json(mockAchievements)
  }),

  // 実績達成の試行
  http.post('/api/try-achieve/:id', ({ params }) => {
    const { id } = params
    const achievementId = Number(id)

    // すでに達成済みの実績かチェック
    const isAlreadyAchieved = mockAchievements.some((a) => a.id === achievementId)

    if (isAlreadyAchieved) {
      return HttpResponse.json({ message: 'すでに達成済みの実績です' }, { status: 400 })
    }

    // 新しい実績を追加
    const newAchievement = {
      id: achievementId,
      name: `実績${achievementId}`,
      achievedAt: new Date().toISOString(),
    }

    mockAchievements.push(newAchievement)

    return HttpResponse.json({ success: true })
  }),

  // ユーザー情報の取得
  http.get('/api/user', () => {
    return HttpResponse.json({
      traqId: 'currentuser',
    })
  }),

  // ヘルスチェック
  http.get('/api/health', () => {
    return HttpResponse.json({ status: 'ok' })
  }),
]
