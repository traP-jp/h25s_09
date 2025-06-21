# MSW (Mock Service Worker) 設定

このプロジェクトでは、フロントエンド開発時にバックエンドAPIをモックするためにMSWを使用しています。

## 設定ファイル

- `src/mocks/handlers.ts`: APIエンドポイントのモックハンドラー
- `src/mocks/browser.ts`: ブラウザ環境用のMSW設定
- `src/mocks/server.ts`: Node.js環境（テスト用）のMSW設定
- `src/mocks/utils.ts`: MSW用のユーティリティ関数
- `src/mocks/index.ts`: エクスポート用のインデックスファイル

## 環境変数

### 開発環境 (`.env.development`)

```
VITE_ENABLE_MSW=true
```

### 本番環境 (`.env.production`)

```
VITE_ENABLE_MSW=false
```

## 使用方法

### 開発環境での使用

開発サーバーを起動すると、自動的にMSWが有効になります：

```bash
npm run dev
```

ブラウザの開発者ツールのコンソールで、MSWの動作ログを確認できます。

### MSWの有効/無効の切り替え

環境変数 `VITE_ENABLE_MSW` を `false` に設定することで、MSWを無効にできます。

### モックデータの編集

`src/mocks/handlers.ts` でモックデータとレスポンスを編集できます。

## 実装されているAPIエンドポイント

### Messages API

- `GET /api/messages` - メッセージ一覧の取得
- `POST /api/messages` - メッセージの投稿
- `GET /api/messages/:id` - メッセージ詳細の取得
- `POST /api/messages/:id/reactions` - リアクションの追加/削除

### Images API

- `POST /api/images` - 画像のアップロード
- `GET /api/images/:id` - 画像の取得

### Achievements API

- `GET /api/achievements` - 実績一覧の取得
- `POST /api/try-achieve/:id` - 実績達成の試行

### User API

- `GET /api/user` - ユーザー情報の取得

### Health Check

- `GET /api/health` - ヘルスチェック

## デバッグ

MSWの動作を確認するには：

1. ブラウザの開発者ツールを開く
2. コンソールタブを確認
3. `[MSW]` プレフィックスのログメッセージを確認

## カスタマイズ

### 新しいエンドポイントの追加

1. `src/mocks/handlers.ts` に新しいハンドラーを追加
2. 必要に応じてモックデータを追加
3. エラーハンドリングを実装

### レスポンス遅延の調整

`src/mocks/utils.ts` の `delay` 関数を使用してレスポンス遅延を調整できます。

### エラーレスポンスのテスト

`src/mocks/utils.ts` の `maybeError` 関数を使用してランダムエラーを発生させることができます。
