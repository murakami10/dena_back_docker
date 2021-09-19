## 初回手順

- app/.env.defaultをapp/.envに変更し、適切な値を設定する.
- app_chat/.env.defaultをapp_chat/.envに変更し、適切な値に変更する.

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
docker-compose -f docker-compose-swagger.yml up -d
```

## for production 
```
docker-compose -f docker-compose-release.yml up -d
```

## MySQL で DB やテーブルが生やされない時

`mysql/` ディレクトリを消す
