# gohost-server

## 使い方

### 推奨環境

mysqlのdockerコンテナを立ち上げた状態でgoをローカルにホストする

```shell
$ docker-compose up
```

### 立ち上げ

```shell
$ go run ./cmd/gohost-server/main.go
```

### APIのcall

#### register

```shell
$ curl -XPOST \
  -H "content-type: application/json" \
  -d '{"first_name": "dena", "last_name": "gohost", "email": "gohost@dena.co.jp", "password": "passw0rd"}' \
  localhost:5050/register
```

#### login

```shell
$ curl -XPOST \
  -H "content-type: application/json" \
  -d '{"email": "gohost@dena.co.jp", "password": "passw0rd"}' \
  --dump-header - \
  localhost:5050/login
```

#### spots

```shell
$ curl \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  "localhost:5050/spots?date=2022-09-10&limit=3"
```

```shell
$ curl \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  "localhost:5050/spots/7e684568-5ecf-4c65-9360-3c8b771b21e7"
```

```shell
$ curl -XPOST \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  -d '{"date": "2022-09-10"}' \
  "localhost:5050/spots/7e684568-5ecf-4c65-9360-3c8b771b21e7/entry"
```

#### logout

```shell
$ curl -XPOST \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  localhost:5050/logout
```


## 開発手順

### API仕様書を書く

1. `./spec/openapi.yaml` にAPI仕様を書く
2. `make gen-api` を実行する

### データベーススキーマを書く

1. `./ddl/*` にDBスキーマを書く
2. `make db-migrate` でマイグレーションする
3. `make gen-db` でdaoファイルを生成する

### handlerを記述する

1. `./internal/handler/` にエンドポイントごとの関数を所定のインターフェースに則って実装する
2. 特に、リクエストやレスポンスのパースの処理をhandlerに記述し、メインのロジックは `./internal/service` に記述する
3. `./internal/service` ではメインのロジックやデータアクセスを記述する

## Makefile

### sqlからmigration

( `docker-compose up` をした状態で)

```shell
$ make db-migrate
```

### migrationの結果からdaocoreのコード自動生成

```shell
$ make gen-db
```

### openapi schemaからmockの自動生成

```shell
$ make gen-api
```

### openapi から HTML のレンダー

```shell
$ make render-api
```
