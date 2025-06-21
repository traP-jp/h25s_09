import { useInfiniteQuery } from '@tanstack/vue-query'
import { useIntersectionObserver } from '@vueuse/core'
import { computed, ref, toValue, type MaybeRefOrGetter } from 'vue'
import { apiService } from '../apis/services'
import { queryKeys } from './api'

interface UseInfiniteMessagesOptions {
  limit?: number
  traqId?: string
  includeReplies?: boolean
}

export function useInfiniteMessages(options: UseInfiniteMessagesOptions = {}) {
  const limit = options.limit || 20

  const { data, fetchNextPage, hasNextPage, isFetchingNextPage, isLoading, error, refetch } =
    useInfiniteQuery({
      queryKey: [...queryKeys.messages, 'infinite', options],
      queryFn: ({ pageParam = 0 }) =>
        apiService.messages.getMessages({
          limit,
          offset: pageParam,
          traqId: options.traqId,
          includeReplies: options.includeReplies,
        }),
      initialPageParam: 0,
      getNextPageParam: (lastPage, allPages) => {
        // 最後のページが空、または件数がlimitより少ない場合は終了
        if (!lastPage || lastPage.length < limit) {
          return undefined
        }
        // 次のページのオフセットを計算
        return allPages.length * limit
      },
      staleTime: 1000 * 60 * 5, // 5分間キャッシュ
    })

  // 全ページのメッセージを平坦化
  const messages = computed(() => {
    if (!data.value) return []
    return data.value.pages.flat()
  })

  // Intersection Observer用のターゲット要素のref
  const target = ref<HTMLElement>()

  // 画面下部に到達したら次のページを読み込む
  const { stop } = useIntersectionObserver(
    target,
    ([{ isIntersecting }]) => {
      if (isIntersecting && hasNextPage.value && !isFetchingNextPage.value) {
        fetchNextPage()
      }
    },
    {
      threshold: 0.1,
    },
  )

  return {
    messages,
    isLoading,
    error,
    hasNextPage,
    isFetchingNextPage,
    fetchNextPage,
    refetch,
    target,
    stop,
  }
}

interface UseInfiniteUserMessagesOptions {
  limit?: number
  includeReplies?: boolean
}

export function useInfiniteUserMessages(
  userId: MaybeRefOrGetter<string>,
  options: UseInfiniteUserMessagesOptions = {},
) {
  const limit = options.limit || 20

  const { data, fetchNextPage, hasNextPage, isFetchingNextPage, isLoading, error, refetch } =
    useInfiniteQuery({
      queryKey: computed(() => [...queryKeys.userMessages(toValue(userId)), 'infinite', options]),
      queryFn: ({ pageParam = 0 }) => {
        const id = toValue(userId)
        if (!id) throw new Error('userId is required')
        return apiService.messages.getMessages({
          limit,
          offset: pageParam,
          traqId: id,
          includeReplies: options.includeReplies,
        })
      },
      initialPageParam: 0,
      getNextPageParam: (lastPage, allPages) => {
        if (!lastPage || lastPage.length < limit) {
          return undefined
        }
        return allPages.length * limit
      },
      enabled: computed(() => !!toValue(userId)),
      staleTime: 1000 * 60 * 5,
    })

  const messages = computed(() => {
    if (!data.value) return []
    return data.value.pages.flat()
  })

  const target = ref<HTMLElement>()

  const { stop } = useIntersectionObserver(
    target,
    ([{ isIntersecting }]) => {
      if (isIntersecting && hasNextPage.value && !isFetchingNextPage.value) {
        fetchNextPage()
      }
    },
    {
      threshold: 0.1,
    },
  )

  return {
    messages,
    isLoading,
    error,
    hasNextPage,
    isFetchingNextPage,
    fetchNextPage,
    refetch,
    target,
    stop,
  }
}

interface UseInfiniteAchievementsOptions {
  limit?: number
  traqId?: string
}

export function useInfiniteAchievements(options: UseInfiniteAchievementsOptions = {}) {
  const { data, fetchNextPage, hasNextPage, isFetchingNextPage, isLoading, error, refetch } =
    useInfiniteQuery({
      queryKey: [...queryKeys.achievements, 'infinite', options],
      queryFn: ({ pageParam = 0 }) => {
        // 現在のAPIはページネーションをサポートしていないため、
        // 最初のページのみ取得してシミュレート
        if (pageParam > 0) return []
        return apiService.achievements.getAchievements()
      },
      initialPageParam: 0,
      getNextPageParam: () => {
        // 現在のAPIでは常に全データを返すため、次のページはない
        return undefined
      },
      staleTime: 1000 * 60 * 15, // 15分間キャッシュ
    })

  const achievements = computed(() => {
    if (!data.value) return []
    return data.value.pages.flat()
  })

  const target = ref<HTMLElement>()

  const { stop } = useIntersectionObserver(
    target,
    ([{ isIntersecting }]) => {
      if (isIntersecting && hasNextPage.value && !isFetchingNextPage.value) {
        fetchNextPage()
      }
    },
    {
      threshold: 0.1,
    },
  )

  return {
    achievements,
    isLoading,
    error,
    hasNextPage,
    isFetchingNextPage,
    fetchNextPage,
    refetch,
    target,
    stop,
  }
}

interface UseInfiniteUserAchievementsOptions {
  limit?: number
}

export function useInfiniteUserAchievements(
  userId: MaybeRefOrGetter<string>,
  options: UseInfiniteUserAchievementsOptions = {},
) {
  const { data, fetchNextPage, hasNextPage, isFetchingNextPage, isLoading, error, refetch } =
    useInfiniteQuery({
      queryKey: computed(() => [
        ...queryKeys.userAchievements(toValue(userId)),
        'infinite',
        options,
      ]),
      queryFn: ({ pageParam = 0 }) => {
        const id = toValue(userId)
        if (!id) throw new Error('userId is required')
        // 現在のAPIはページネーションをサポートしていないため、
        // 最初のページのみ取得してシミュレート
        if (pageParam > 0) return []
        return apiService.achievements.getUserAchievements(id)
      },
      initialPageParam: 0,
      getNextPageParam: () => {
        // 現在のAPIでは常に全データを返すため、次のページはない
        return undefined
      },
      enabled: computed(() => !!toValue(userId)),
      staleTime: 1000 * 60 * 15,
    })

  const achievements = computed(() => {
    if (!data.value) return []
    return data.value.pages.flat()
  })

  const target = ref<HTMLElement>()

  const { stop } = useIntersectionObserver(
    target,
    ([{ isIntersecting }]) => {
      if (isIntersecting && hasNextPage.value && !isFetchingNextPage.value) {
        fetchNextPage()
      }
    },
    {
      threshold: 0.1,
    },
  )

  return {
    achievements,
    isLoading,
    error,
    hasNextPage,
    isFetchingNextPage,
    fetchNextPage,
    refetch,
    target,
    stop,
  }
}
