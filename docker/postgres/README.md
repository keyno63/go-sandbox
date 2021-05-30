# Postgres Dcoker

Postgres 用の Dcoker を準備するための環境  

## Dockerfile のビルド

```shell
docekr build -t keyno63/postgres .
```

## Docker の起動

 ```shell
docker run -p 15432:5432 --name [container name] keyno63/postgres
```

// [container name] には一意となる、任意のコンテナ名をつけてください。  
// 設定するのは管理を容易にするためです

## postgres への接続
### psql を用いて

```shell
psql -l 127.0.0.1  -p 15432 -U keyno63 development
```

password を聞かれますが、Dockerfile 内の設定を確認してください。  
