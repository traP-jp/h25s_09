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
  imageId: string | null // OpenAPI仕様に合わせてimageIdフィールド
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

// モックデータ - より多くのメッセージを生成
const generateMockMessages = (count: number): Message[] => {
  const messages: Message[] = []
  for (let i = 1; i <= count; i++) {
    const reactions = {
      count: Math.floor(Math.random() * 10),
      myReaction: Math.random() > 0.5,
    }

    const message: Message = {
      id: `550e8400-e29b-41d4-a716-${i.toString().padStart(12, '0')}`,
      author: 'rei',
      content: `これは${i}番目のテストメッセージです。Infinite scrollのテストに使用されます。`,
      imageId: i % 5 === 0 ? `image-${i}` : null, // 5個おきに画像を追加
      reactions,
      replyCount: Math.floor(Math.random() * 3),
      createdAt: new Date(Date.now() - (count - i) * 60000).toISOString(), // 1分ずつ古い時刻
    }

    messages.push(message)
  }
  return messages
}

const mockMessages: Message[] = generateMockMessages(100) // 100個のメッセージを生成

const generateMockMessageDetails = (messages: Message[]): MessageDetail[] => {
  return messages.map((message) => ({
    id: message.id,
    author: message.author,
    content: message.content,
    imageId: message.imageId,
    reactions: message.reactions,
    createdAt: message.createdAt,
    replies:
      message.replyCount > 0
        ? [
            {
              id: `reply-${message.id}`,
              author: 'rei',
              content: `${message.content}への返信です。`,
              imageId: null,
              reactions: {
                count: Math.floor(Math.random() * 3),
                myReaction: Math.random() > 0.5,
              },
              createdAt: new Date(new Date(message.createdAt).getTime() + 300000).toISOString(), // 5分後
            },
          ]
        : [],
  }))
}

const mockMessageDetails: MessageDetail[] = generateMockMessageDetails(mockMessages)

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
      mswLog(
        'info',
        `ユーザー ${traqId} のメッセージをフィルタリング: ${filteredMessages.length}件`,
      )
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
              imageId: reply.imageId,
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

    // メッセージ本文が空で、かつ画像も添付されていない場合はエラー
    if ((!message || message.trim() === '') && !image) {
      return HttpResponse.json(
        { message: 'メッセージ本文または画像のいずれかが必要です' },
        { status: 400 },
      )
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
        imageId: image ? `image-${Date.now()}` : null,
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

  // 実績の作成
  http.post('/api/me/achievements', async ({ request }) => {
    const body = (await request.json()) as { name: string }

    if (!body.name) {
      return HttpResponse.json({ message: '実績名が必要です' }, { status: 400 })
    }

    if (body.name.length > 32) {
      return HttpResponse.json({ message: '実績名は32文字以内で入力してください' }, { status: 400 })
    }

    // 新しい実績を作成
    const newAchievement = {
      id: mockAchievements.length + 1,
      name: body.name,
      achievedAt: new Date().toISOString(),
    }

    mockAchievements.push(newAchievement)

    return HttpResponse.json(newAchievement, { status: 201 })
  }),

  // 実績一覧の取得
  http.get('/api/users/:traqId/achievements', ({ params }) => {
    const { traqId } = params

    // /me APIで取得されるcurrentUserのtraqId（モック環境では固定値）
    const currentUserTraqId = 'rei'

    if (traqId !== currentUserTraqId) {
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

    // まずメインメッセージから検索
    const messageDetail = mockMessageDetails.find((msg) => msg.id === id)
    let targetReply: Reply | null = null

    // メインメッセージが見つからない場合、リプライから検索
    if (!messageDetail) {
      for (const msg of mockMessageDetails) {
        const reply = msg.replies.find((r) => r.id === id)
        if (reply) {
          targetReply = reply
          break
        }
      }

      if (!targetReply) {
        return HttpResponse.json({ message: 'メッセージが見つかりません' }, { status: 404 })
      }
    }

    if (messageDetail) {
      // メインメッセージのリアクション処理
      if (!messageDetail.reactions.myReaction) {
        messageDetail.reactions.count += 1
        messageDetail.reactions.myReaction = true

        // 対応するMessageも更新
        const message = mockMessages.find((msg) => msg.id === id)
        if (message) {
          message.reactions = { ...messageDetail.reactions }
        }

        mswLog('info', `メッセージ ${id} にリアクションを追加`)
      } else {
        mswLog('warn', `メッセージ ${id} は既にリアクション済み`)
      }

      return HttpResponse.json(messageDetail.reactions, { status: 201 })
    } else if (targetReply) {
      // リプライのリアクション処理
      if (!targetReply.reactions.myReaction) {
        targetReply.reactions.count += 1
        targetReply.reactions.myReaction = true

        mswLog('info', `リプライ ${id} にリアクションを追加`)
      } else {
        mswLog('warn', `リプライ ${id} は既にリアクション済み`)
      }

      return HttpResponse.json(targetReply.reactions, { status: 201 })
    }
  }),

  // リアクション削除
  http.delete('/api/messages/:id/reactions', ({ params }) => {
    const { id } = params

    // まずメインメッセージから検索
    const messageDetail = mockMessageDetails.find((msg) => msg.id === id)
    let targetReply: Reply | null = null

    // メインメッセージが見つからない場合、リプライから検索
    if (!messageDetail) {
      for (const msg of mockMessageDetails) {
        const reply = msg.replies.find((r) => r.id === id)
        if (reply) {
          targetReply = reply
          break
        }
      }

      if (!targetReply) {
        return HttpResponse.json({ message: 'メッセージが見つかりません' }, { status: 404 })
      }
    }

    if (messageDetail) {
      // メインメッセージのリアクション削除
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
    } else if (targetReply) {
      // リプライのリアクション削除
      if (targetReply.reactions.myReaction) {
        targetReply.reactions.count = Math.max(0, targetReply.reactions.count - 1)
        targetReply.reactions.myReaction = false
      }

      mswLog('info', `リプライ ${id} のリアクションを削除`)
      return HttpResponse.json(targetReply.reactions)
    }
  }),
]
