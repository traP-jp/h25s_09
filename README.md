# h25s_09

## メンバー

- tidus (リーダー, バックエンド)
- ebisen(下っ端 フロント)
- quarantineeeeeeeeee (無職, バックエンド)
- you6174
- ebisen(部下 フロント)

## 開発ワークフロー

### バックエンド

`backend` ディレクトリ直下に `.env` ファイルが必要です.

デフォルト値は以下を推奨します.

```ini
NS_MARIADB_HOSTNAME=db
NS_MARIADB_PORT=3306
NS_MARIADB_DATABASE=h25s_09
NS_MARIADB_USER=user
NS_MARIADB_PASSWORD=password
```

また, 開発では Docker コンテナを利用することができます.
[Air](https://github.com/air-verse/air) を利用したホットリロードの設定をしているので, 変更を加えてもコンテナの再起動は不要です.

- [Taskfile](https://taskfile.dev) を利用する場合は, 以下のコマンドで開発用 Docker コンテナの起動と停止を行えます.

  ```shell
  task up
  task down
  ```

- [Docker Compose](https://docs.docker.com/compose/) を利用する場合は, 以下のコマンドで開発用 Docker コンテナの起動と停止を行えます.

  ```shell
  docker-compose up -d
  docker-compose down
  ```

Dockerfile は2つあるので, 用途に応じて使い分けてください.

- `dev.Dockerfile` -- 開発用. ホットリロード可. イメージサイズがだいぶでかい.
- `prod.Dockerfile` -- 本番用. ホットリロード不可. イメージサイズは軽量.
