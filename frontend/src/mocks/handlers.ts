import { http, HttpResponse } from 'msw'
import { delay, mswLog } from './utils'

// 型定義
interface Message {
  id: string
  author: string // OpenAPI仕様に合わせてstringに変更
  content: string
  imageId: string | null // OpenAPI仕様に合わせてimageIdに変更
  reactions: {
    count: number
    myReaction: boolean
  }
  replyCount: number // OpenAPI仕様に合わせて追加
  createdAt: string
}

interface Reply {
  id: string
  author: string
  content: string
  images: string | null // OpenAPI仕様に合わせてimagesフィールド
  reactions: {
    count: number
    myReaction: boolean
  }
  createdAt: string
}

interface MessageDetail {
  id: string
  author: string
  content: string
  imageId: string | null
  reactions: {
    count: number
    myReaction: boolean
  }
  replies: Reply[]
  createdAt: string
}

// モックデータ
const mockMessages: Message[] = [
  {
    id: '550e8400-e29b-41d4-a716-446655440001',
    author: 'rei',
    content: 'こんにちは！これは最初のテストメッセージです。',
    imageId: null,
    reactions: {
      count: 3,
      myReaction: true,
    },
    replyCount: 1,
    createdAt: '2025-06-21T10:00:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440002',
    author: 'rei',
    content: 'こんにちは！画像付きのメッセージです。',
    imageId: '550e8400-e29b-41d4-a716-446655440010',
    reactions: {
      count: 1,
      myReaction: false,
    },
    replyCount: 0,
    createdAt: '2025-06-21T10:30:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440003',
    author: 'rei',
    content: 'これは返信メッセージの例です。',
    imageId: null,
    reactions: {
      count: 0,
      myReaction: false,
    },
    replyCount: 0,
    createdAt: '2025-06-21T11:00:00Z',
  },
]

const mockMessageDetails: MessageDetail[] = [
  {
    id: '550e8400-e29b-41d4-a716-446655440001',
    author: 'rei',
    content: 'こんにちは！これは最初のテストメッセージです。',
    imageId: null,
    reactions: {
      count: 3,
      myReaction: true,
    },
    createdAt: '2025-06-21T10:00:00Z',
    replies: [
      {
        id: '550e8400-e29b-41d4-a716-446655440004',
        author: 'rei',
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
    const includeReplies = url.searchParams.get('includeReplies') === 'true'

    // デバッグログ
    mswLog(
      'info',
      `Parameters: limit=${limit}, offset=${offset}, traqId=${traqId}, includeReplies=${includeReplies}`,
    )

    let filteredMessages = mockMessages

    if (traqId) {
      filteredMessages = mockMessages.filter((message) => message.author === traqId)
      mswLog('info', `ユーザー ${traqId} のメッセージをフィルタリング`)
    }

    // includeRepliesがtrueの場合は、返信も含めたメッセージを返す
    if (includeReplies) {
      // 返信も含めたメッセージ配列を作成
      const messagesWithReplies: Message[] = []

      filteredMessages.forEach((message) => {
        messagesWithReplies.push(message)

        // 対応するMessageDetailから返信を取得
        const messageDetail = mockMessageDetails.find((md) => md.id === message.id)
        if (messageDetail && messageDetail.replies.length > 0) {
          // 返信をMessageの形式に変換して追加
          messageDetail.replies.forEach((reply) => {
            const replyAsMessage: Message = {
              id: reply.id,
              author: reply.author,
              content: reply.content,
              imageId: reply.images,
              reactions: reply.reactions,
              replyCount: 0, // 返信の返信は現在サポートしていない
              createdAt: reply.createdAt,
            }
            messagesWithReplies.push(replyAsMessage)
          })
        }
      })

      const paginatedMessages = messagesWithReplies.slice(offset, offset + limit)
      mswLog('info', `${paginatedMessages.length}件のメッセージ（返信含む）を返却`)
      return HttpResponse.json(paginatedMessages)
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
    const repliesTo = formData.get('repliesTo') as string

    if (!message || message.trim() === '') {
      return HttpResponse.json({ message: 'メッセージ本文が空です' }, { status: 400 })
    }

    // /me APIで取得されるcurrentUserのtraqId（モック環境では固定値）
    const currentUserTraqId = 'rei'

    const newMessageId = `550e8400-e29b-41d4-a716-${Date.now()}`

    if (repliesTo) {
      // リプライの場合はMessageDetailのrepliesに追加
      const parentMessage = mockMessageDetails.find((msg) => msg.id === repliesTo)
      if (!parentMessage) {
        return HttpResponse.json({ message: '返信先のメッセージが見つかりません' }, { status: 400 })
      }

      const newReply: Reply = {
        id: newMessageId,
        author: currentUserTraqId,
        content: message,
        images: image ? `image-${Date.now()}` : null,
        reactions: {
          count: 0,
          myReaction: false,
        },
        createdAt: new Date().toISOString(),
      }

      parentMessage.replies.push(newReply)

      // 親メッセージのreplyCountを更新
      const parentInMessages = mockMessages.find((msg) => msg.id === repliesTo)
      if (parentInMessages) {
        parentInMessages.replyCount += 1
      }

      // MessageDetailを返却
      const newMessageDetail: MessageDetail = {
        id: newMessageId,
        author: currentUserTraqId,
        content: message,
        imageId: image ? `image-${Date.now()}` : null,
        reactions: {
          count: 0,
          myReaction: false,
        },
        replies: [],
        createdAt: new Date().toISOString(),
      }

      mswLog('info', `リプライメッセージを作成: ${newMessageId}`)
      return HttpResponse.json(newMessageDetail, { status: 201 })
    } else {
      // 通常のメッセージ投稿
      const newMessage: Message = {
        id: newMessageId,
        author: currentUserTraqId,
        content: message,
        imageId: image ? `image-${Date.now()}` : null,
        reactions: {
          count: 0,
          myReaction: false,
        },
        replyCount: 0,
        createdAt: new Date().toISOString(),
      }

      // モックデータに追加
      mockMessages.unshift(newMessage)

      // MessageDetailも作成
      const newMessageDetail: MessageDetail = {
        ...newMessage,
        replies: [],
      }
      mockMessageDetails.unshift(newMessageDetail)

      mswLog('info', `新しいメッセージを作成: ${newMessageId}`)
      return HttpResponse.json(newMessageDetail, { status: 201 })
    }
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

  // 画像の取得
  http.get('/api/images/:id', ({ params }) => {
    const { id } = params

    // 実際のプロジェクトでは、適切な画像データを返す
    // モック環境では画像URLを返す
    return new Response(
      JSON.stringify({ url: `https://via.placeholder.com/400x300?text=Image+${id}` }),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        },
      },
    )
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

    // OpenAPI仕様に合わせてdispatchedフィールドを返す
    return HttpResponse.json({ dispatched: true })
  }),

  // 実績一覧の取得
  http.get('/api/achievements', ({ request }) => {
    const url = new URL(request.url)
    const traqId = url.searchParams.get('traqId')

    // /me APIで取得されるcurrentUserのtraqId（モック環境では固定値）
    const currentUserTraqId = 'rei'

    if (traqId && traqId !== currentUserTraqId) {
      // 他のユーザーの実績を返す（少なめ）
      return HttpResponse.json([mockAchievements[0]])
    }

    return HttpResponse.json(mockAchievements)
  }),

  // 自身の情報を取得
  http.get('/api/me', () => {
    return HttpResponse.json({
      traqId: 'rei',
    })
  }),

  // リアクション追加
  http.post('/api/messages/:id/reactions', ({ params }) => {
    const { id } = params
    const messageDetail = mockMessageDetails.find((msg) => msg.id === id)

    if (!messageDetail) {
      return HttpResponse.json({ message: 'メッセージが見つかりません' }, { status: 404 })
    }

    // リアクションを追加/切り替え
    if (messageDetail.reactions.myReaction) {
      // 既にリアクション済みの場合は削除
      messageDetail.reactions.count = Math.max(0, messageDetail.reactions.count - 1)
      messageDetail.reactions.myReaction = false
    } else {
      // リアクションを追加
      messageDetail.reactions.count += 1
      messageDetail.reactions.myReaction = true
    }

    // 対応するMessageも更新
    const message = mockMessages.find((msg) => msg.id === id)
    if (message) {
      message.reactions = { ...messageDetail.reactions }
    }

    mswLog('info', `メッセージ ${id} のリアクションを切り替え`)
    return HttpResponse.json(messageDetail.reactions, { status: 201 })
  }),

  // リアクション削除
  http.delete('/api/messages/:id/reactions', ({ params }) => {
    const { id } = params
    const messageDetail = mockMessageDetails.find((msg) => msg.id === id)

    if (!messageDetail) {
      return HttpResponse.json({ message: 'メッセージが見つかりません' }, { status: 404 })
    }

    // リアクションを削除
    if (messageDetail.reactions.myReaction) {
      messageDetail.reactions.count = Math.max(0, messageDetail.reactions.count - 1)
      messageDetail.reactions.myReaction = false

      // 対応するMessageも更新
      const message = mockMessages.find((msg) => msg.id === id)
      if (message) {
        message.reactions = { ...messageDetail.reactions }
      }
    }

    mswLog('info', `メッセージ ${id} のリアクションを削除`)
    return HttpResponse.json(messageDetail.reactions)
  }),
]
