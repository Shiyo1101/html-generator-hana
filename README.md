# Hana道場HP用 HTML（CSS）ジェネレータ

## 開発

### Dockerイメージの作成

```
docker-compose build
```
or
```
docker compose up
```

### Dockerコンテナの立ち上げ

```
docker-compose up -d
```
or
```
docker compose up
```

### コンテナに入る

```
docker exec -it html-generator-hana-go /bin/sh
```

## main.goのクロスコンパイル

### Linux Amd64

```
GOOS=linux GOARCH=arm64 go build -o [生成ファイル名].exe main.go
```

### MacOS Apple Silicon（Arm64）
```
GOOS=darwin GOARCH=arm64 go build -o [生成ファイル名].exe main.go
```

### Windows Intel,AMD CPU（Amd64）

```
GOOS=windows GOARCH=amd64 go build -o [生成ファイル名].exe main.go
```

