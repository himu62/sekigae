.PHONY: more \
	docker/build docker/start docker/stop docker/clean docker/bash \
	db/create db/migrate db/clean \
	server server/test server/build server/deps \
	client/build client/watch client/deps

# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
more:
	@awk 'BEGIN {FS = ":.*?## "} /^[\$$\(\)\/a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

docker/build: Dockerfile docker-compose.yml       ## MySQLコンテナを作成
	docker-compose build

docker/start:                                     ## MySQLコンテナを起動
	docker-compose up -d

docker/stop:                                      ## MySQLコンテナを停止
	docker-compose down

docker/clean:                                     ## MySQLコンテナを削除
	docker-compose rm

docker/bash:                                      ## MySQLコンテナのコンソール起動
	docker exec -it $(shell docker-compose ps -q) bash

db/create:                                        ## DBを作成
	mysql --protocol tcp -u root -ppassword -e 'CREATE DATABASE sekigae DEFAULT COLLATE utf8mb4_unicode_ci'

db/migrate:                                       ## DBをマイグレーション
	sql-migrate up

db/clean:                                         ## DBを削除
	sql-migrate down

server:                                           ## サーバを起動
	go run sekigae.go

server/test:                                      ## サーバをテスト
	go test -v ./...

server/build:                                     ## サーバをビルド
	go build -o sekigae sekigae.go

server/deps:                                      ## サーバの依存パッケージをダウンロード
	go get -d -t -v ./...
	go get -v github.com/rubenv/sql-migrate/...

client/build:                                     ## クライアントをビルド
	NODE_ENV=production ./node_modules/.bin/webpack

client/watch:                                     ## クライアントを自動コンパイル
	NODE_ENV=development ./node_modules/.bin/webpack --watch

client/deps:                                      ## クライアントの依存パッケージをダウンロード
	npm install
	./node_modules/.bin/typings install
