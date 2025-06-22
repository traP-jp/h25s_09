import type { Achievement, Message, MessageDetail, Reactions, UserInfo } from './generated'
import httpClient from './http-client'

// Messages API
export const messagesService = {
  async getMessages(options?: {
    limit?: number
    offset?: number
    traqId?: string
    includeReplies?: boolean
  }): Promise<Message[]> {
    const params = new URLSearchParams()

    if (options?.limit !== undefined) {
      params.append('limit', options.limit.toString())
    }
    if (options?.offset !== undefined) {
      params.append('offset', options.offset.toString())
    }
    if (options?.traqId) {
      params.append('traqId', options.traqId)
    }
    // includeRepliesは明示的にfalseの場合も送信する
    if (options?.includeReplies !== undefined) {
      params.append('includeReplies', options.includeReplies.toString())
    }

    const queryString = params.toString()
    const url = queryString ? `/messages?${queryString}` : '/messages'

    console.log('API Request URL:', url) // デバッグ用

    return httpClient.get<Message[]>(url)
  },

  async getMessageDetail(id: string): Promise<MessageDetail> {
    return httpClient.get<MessageDetail>(`/messages/${id}`)
  },

  async createMessage(content: string, image?: File, repliesTo?: string): Promise<MessageDetail> {
    const formData = new FormData()
    formData.append('message', content)

    if (image) {
      formData.append('image', image)
    }

    if (repliesTo) {
      formData.append('repliesTo', repliesTo)
    }

    return httpClient.post<MessageDetail>('/messages', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })
  },

  async deleteMessage(id: string): Promise<void> {
    return httpClient.delete(`/messages/${id}`)
  },

  async addReaction(messageId: string): Promise<Reactions> {
    return httpClient.post<Reactions>(`/messages/${messageId}/reactions`)
  },

  async removeReaction(messageId: string): Promise<Reactions> {
    return httpClient.delete<Reactions>(`/messages/${messageId}/reactions`)
  },
}

// User API
export const userService = {
  async getUserInfo(): Promise<UserInfo> {
    return httpClient.get<UserInfo>('/me')
  },

  async getUserMessages(userId: string, includeReplies?: boolean): Promise<Message[]> {
    // OpenAPI仕様では/messagesエンドポイントでtraqIdパラメータを使用
    const params = new URLSearchParams({ traqId: userId })
    if (includeReplies !== undefined) {
      params.append('includeReplies', includeReplies.toString())
    }
    return httpClient.get<Message[]>(`/messages?${params.toString()}`)
  },

  async updateUserInfo(userInfo: Partial<UserInfo>): Promise<UserInfo> {
    return httpClient.patch<UserInfo>('/me', userInfo)
  },
}

// Achievements API
export const achievementsService = {
  async getAchievements(): Promise<Achievement[]> {
    return httpClient.get<Achievement[]>('/achievements')
  },

  async getUserAchievements(userId: string): Promise<Achievement[]> {
    // OpenAPI仕様では/achievementsエンドポイントでtraqIdパラメータを使用
    return httpClient.get<Achievement[]>(`/achievements?traqId=${userId}`)
  },

  async createAchievement(data: { name: string }): Promise<Achievement> {
    return httpClient.post<Achievement>('/me-achievements', data)
  },
}

// Images API
export const imagesService = {
  async getImage(imageId: string): Promise<{ url: string }> {
    return httpClient.get<{ url: string }>(`/images/${imageId}`)
  },
}

// 統合APIサービス
export const apiService = {
  messages: messagesService,
  user: userService,
  achievements: achievementsService,
  images: imagesService,
}
