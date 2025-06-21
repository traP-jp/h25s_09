import axios, {
  AxiosError,
  type AxiosInstance,
  type AxiosRequestConfig,
  type AxiosResponse,
} from 'axios'

// APIエラーの型定義
export interface ApiError {
  message: string
  status?: number
  code?: string
}

// HTTPクライアントの設定
class HttpClient {
  private client: AxiosInstance

  constructor(baseURL: string = '/api') {
    this.client = axios.create({
      baseURL,
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json',
      },
    })

    this.setupInterceptors()
  }

  private setupInterceptors() {
    // リクエストインターセプター
    this.client.interceptors.request.use(
      (config) => {
        // 認証トークンがあれば追加
        const token = localStorage.getItem('auth_token')
        if (token) {
          config.headers.Authorization = `Bearer ${token}`
        }
        return config
      },
      (error) => Promise.reject(error),
    )

    // レスポンスインターセプター
    this.client.interceptors.response.use(
      (response: AxiosResponse) => response,
      (error: AxiosError) => {
        const apiError: ApiError = {
          message: error.message || 'An unexpected error occurred',
          status: error.response?.status,
          code: error.code,
        }

        // エラーレスポンスからメッセージを抽出
        if (error.response?.data && typeof error.response.data === 'object') {
          const responseData = error.response.data as Record<string, unknown>
          if (responseData.message && typeof responseData.message === 'string') {
            apiError.message = responseData.message
          }
        }

        return Promise.reject(apiError)
      },
    )
  }

  async get<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
    const response = await this.client.get<T>(url, config)
    return response.data
  }

  async post<T>(url: string, data?: unknown, config?: AxiosRequestConfig): Promise<T> {
    const response = await this.client.post<T>(url, data, config)
    return response.data
  }

  async put<T>(url: string, data?: unknown, config?: AxiosRequestConfig): Promise<T> {
    const response = await this.client.put<T>(url, data, config)
    return response.data
  }

  async patch<T>(url: string, data?: unknown, config?: AxiosRequestConfig): Promise<T> {
    const response = await this.client.patch<T>(url, data, config)
    return response.data
  }

  async delete<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
    const response = await this.client.delete<T>(url, config)
    return response.data
  }
}

export const httpClient = new HttpClient()
export default httpClient
