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