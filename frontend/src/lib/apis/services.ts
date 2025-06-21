import type { Achievement, Message, MessageDetail, Reply, UserInfo } from './generated'
import httpClient from './http-client'

// Messages API
export const messagesService = {
  async getMessages(): Promise<Message[]> {
    return httpClient.get<Message[]>('/messages')
  },

  async getMessageDetail(id: string): Promise<MessageDetail> {
    return httpClient.get<MessageDetail>(`/messages/${id}`)
  },

  async createMessage(content: string): Promise<Message> {
    return httpClient.post<Message>('/messages', { content })
  },

  async deleteMessage(id: string): Promise<void> {
    return httpClient.delete(`/messages/${id}`)
  },

  async createReply(messageId: string, content: string): Promise<Reply> {
    return httpClient.post<Reply>(`/messages/${messageId}/replies`, { content })
  },

  async addReaction(messageId: string, emoji: string): Promise<void> {
    return httpClient.post(`/messages/${messageId}/reactions`, { emoji })
  },

  async removeReaction(messageId: string, emoji: string): Promise<void> {
    return httpClient.delete(`/messages/${messageId}/reactions/${emoji}`)
  },
}

// User API
export const userService = {
  async getUserInfo(): Promise<UserInfo> {
    return httpClient.get<UserInfo>('/user')
  },

  async getUserMessages(userId?: string): Promise<Message[]> {
    const endpoint = userId ? `/user/${userId}/messages` : '/user/messages'
    return httpClient.get<Message[]>(endpoint)
  },

  async updateUserInfo(userInfo: Partial<UserInfo>): Promise<UserInfo> {
    return httpClient.patch<UserInfo>('/user', userInfo)
  },
}

// Achievements API
export const achievementsService = {
  async getAchievements(): Promise<Achievement[]> {
    return httpClient.get<Achievement[]>('/achievements')
  },

  async getUserAchievements(userId?: string): Promise<Achievement[]> {
    const endpoint = userId ? `/user/${userId}/achievements` : '/user/achievements'
    return httpClient.get<Achievement[]>(endpoint)
  },

  async tryAchieve(achievementId: string): Promise<{ success: boolean; message: string }> {
    return httpClient.post(`/achievements/${achievementId}/try`)
  },
}

// Images API
export const imagesService = {
  async uploadImage(file: File): Promise<{ url: string }> {
    const formData = new FormData()
    formData.append('image', file)

    return httpClient.post<{ url: string }>('/images', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })
  },
}

// 統合APIサービス
export const apiService = {
  messages: messagesService,
  user: userService,
  achievements: achievementsService,
  images: imagesService,
}
