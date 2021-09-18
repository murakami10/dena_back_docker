## 初回手順

```
$ docker-compose build
$ docker-compose up
```

## go で新しい module を追加した時

```
docker exec -it <go-container-id> go mod tidy
```

## swagger 
```
docker-compose  -f docker-compose-swagger.yml up -d
```

## MySQL で DB やテーブルが生やされない時

`mysql/` ディレクトリを消す
