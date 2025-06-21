import { AchievementsApi, Configuration, ImagesApi, MessagesApi, UserApi } from './generated'

// APIのベースURL設定
const configuration = new Configuration({
  basePath: '/api', // バックエンドのAPIパス
})

// 各APIクラスのインスタンスを作成
export const messagesApi = new MessagesApi(configuration)
export const achievementsApi = new AchievementsApi(configuration)
export const imagesApi = new ImagesApi(configuration)
export const userApi = new UserApi(configuration)

// 使いやすいAPIクライアント
export const apiClient = {
  messages: messagesApi,
  achievements: achievementsApi,
  images: imagesApi,
  user: userApi,
}

export default apiClient
