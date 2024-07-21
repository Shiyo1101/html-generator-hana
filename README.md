# Hana道場HP用 HTML（CSS）ジェネレータ

各種プラットフォームで動作するコマンドラインツールです。

## 開発

### Dockerイメージの作成

```sh
docker-compose build
```
または
```sh
docker compose build
```

### Dockerコンテナの立ち上げ

```sh
docker-compose up -d
```
または
```sh
docker compose up
```

### コンテナに入る

```sh
docker exec -it html-generator-hana /bin/sh
```

## main.goのクロスコンパイル

### Linux Amd64

```sh
GOOS=linux GOARCH=arm64 go build -o [生成ファイル名].exe main.go
```

### MacOS Apple Silicon（Arm64）

```sh
GOOS=darwin GOARCH=arm64 go build -o [生成ファイル名].exe main.go
```

### Windows Intel, AMD CPU（Amd64）

```sh
GOOS=windows GOARCH=amd64 go build -o [生成ファイル名].exe main.go
```