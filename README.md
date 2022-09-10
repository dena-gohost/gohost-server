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

get

```shell
$ curl -XGET \
  -H "content-type: application/json" \
  localhost:5050/register
```

post

```shell
$ curl -XPOST \
  -H "content-type: application/json" \
  -d '{"first_name": "DeNA", "last_name": "ハッカソン", "user_name": "denason", "email": "hackathon@dena.ac.jp", "password": "passw0rd", "university_id": "06faba8c-5d9f-495c-a9d7-3bc82a040398", "birth_date": "1999-09-10", "year": 4, "gender_id": "1", "icon_url": "https://www.photo-ac.com/assets/img/ai_model_v2/model_6.png", "instagram_id": "dummy"}' \
  localhost:5050/register
```

#### login

```shell
$ curl -XPOST \
  -H "content-type: application/json" \
  -d '{"email": "hackathon@dena.ac.jp", "password": "passw0rd"}' \
  --dump-header - \
  localhost:5050/login
```

`Set-Cookie: ` 以降をコピーする. 変数 `COOKIE`に代入する

macなら次の通り

```shell
$ COOKIE="$(pbpaste)"
```

#### spots

心霊スポット一覧

```shell
$ curl \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  "localhost:5050/spots?date=2022-09-10&limit=3"
```

心霊スポット詳細

```shell
$ curl \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  "localhost:5050/spots/7ed0c28c-890a-4c3e-8c0d-6db93969168c"
```

エントリー

```shell
$ curl -XPOST \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  -d '{"date": "2022-09-10"}' \
  "localhost:5050/spots/7ed0c28c-890a-4c3e-8c0d-6db93969168c/entry"
```

#### plan

プラン確認

```shell
$ curl -XGET \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  localhost:5050/plan
```

プランキャンセル

```shell
$ curl -XPOST \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  localhost:5050/plan/cancel
```

プラン完了(GET)

```shell
$ curl -XGET \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  localhost:5050/plan/finish
```

プラン完了(POST)

```shell
$ curl -XPOST \
  -H "cookie: ${COOKIE}" \
  -H "content-type: application/json" \
  -d '[{"user_id": "aeb4238a-0582-408f-a68e-fc3f8163fc87", "like": true}, {"user_id": "1f4e3db9-e6ec-4028-befe-4f92c2bb2051", "like": false}]' \
  localhost:5050/plan/finish
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
