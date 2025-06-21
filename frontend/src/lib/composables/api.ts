import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'
import type { MaybeRefOrGetter } from '@vueuse/core'
import { computed, toValue } from 'vue'

import { apiService } from '../apis/services'

// Query Keys
export const queryKeys = {
  messages: ['messages'] as const,
  messageDetail: (id: string) => ['messages', id] as const,
  userMessages: (userId?: string) => ['user', 'messages', userId] as const,
  userAchievements: (userId?: string) => ['user', 'achievements', userId] as const,
  achievements: ['achievements'] as const,
  userInfo: ['user', 'info'] as const,
}

// Messages API Hooks
export function useMessages() {
  return useQuery({
    queryKey: queryKeys.messages,
    queryFn: apiService.messages.getMessages,
    staleTime: 1000 * 60 * 5, // 5分間キャッシュ
  })
}

export function useMessageDetail(messageId: MaybeRefOrGetter<string>) {
  return useQuery({
    queryKey: computed(() => queryKeys.messageDetail(toValue(messageId))),
    queryFn: () => apiService.messages.getMessageDetail(toValue(messageId)),
    enabled: computed(() => !!toValue(messageId)),
  })
}

export function useCreateMessage() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({
      content,
      image,
      repliesTo,
    }: {
      content: string
      image?: File
      repliesTo?: string
    }) => apiService.messages.createMessage(content, image, repliesTo),
    onSuccess: () => {
      // メッセージ一覧を無効化して再取得
      queryClient.invalidateQueries({ queryKey: queryKeys.messages })
      // メッセージ詳細も無効化
      queryClient.invalidateQueries({ queryKey: ['messages'] })
    },
  })
}

export function useDeleteMessage() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (messageId: string) => apiService.messages.deleteMessage(messageId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.messages })
    },
  })
}

// User API Hooks
export function useUserInfo() {
  return useQuery({
    queryKey: queryKeys.userInfo,
    queryFn: apiService.user.getUserInfo,
    staleTime: 1000 * 60 * 10, // 10分間キャッシュ
  })
}

export function useUserMessages(userId?: MaybeRefOrGetter<string>) {
  return useQuery({
    queryKey: computed(() => queryKeys.userMessages(toValue(userId))),
    queryFn: () => {
      const id = toValue(userId)
      if (!id) throw new Error('userId is required')
      return apiService.user.getUserMessages(id)
    },
    enabled: computed(() => !!toValue(userId)),
  })
}

export function useUpdateUserInfo() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: apiService.user.updateUserInfo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.userInfo })
    },
  })
}

// Achievements API Hooks
export function useAchievements() {
  return useQuery({
    queryKey: queryKeys.achievements,
    queryFn: apiService.achievements.getAchievements,
    staleTime: 1000 * 60 * 15, // 15分間キャッシュ
  })
}

export function useUserAchievements(userId?: MaybeRefOrGetter<string>) {
  return useQuery({
    queryKey: computed(() => queryKeys.userAchievements(toValue(userId))),
    queryFn: () => {
      const id = toValue(userId)
      if (!id) throw new Error('userId is required')
      return apiService.achievements.getUserAchievements(id)
    },
    enabled: computed(() => !!toValue(userId)),
  })
}

export function useTryAchieve() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (achievementId: string) => apiService.achievements.tryAchieve(achievementId),
    onSuccess: () => {
      // ユーザーの実績リストを更新
      queryClient.invalidateQueries({ queryKey: queryKeys.userAchievements() })
    },
  })
}

// Reactions API Hooks
export function useAddReaction() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (messageId: string) => apiService.messages.addReaction(messageId),
    onSuccess: (_, messageId) => {
      // メッセージ詳細を無効化して再取得
      queryClient.invalidateQueries({ queryKey: queryKeys.messageDetail(messageId) })
      // メッセージ一覧も無効化
      queryClient.invalidateQueries({ queryKey: queryKeys.messages })
    },
  })
}

export function useRemoveReaction() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (messageId: string) => apiService.messages.removeReaction(messageId),
    onSuccess: (_, messageId) => {
      // メッセージ詳細を無効化して再取得
      queryClient.invalidateQueries({ queryKey: queryKeys.messageDetail(messageId) })
      // メッセージ一覧も無効化
      queryClient.invalidateQueries({ queryKey: queryKeys.messages })
    },
  })
}
