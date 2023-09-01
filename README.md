# 使用說明

## 複製一份 config.yaml 並修改 db 連線資訊
```
cp conf/config.example.yaml conf/config.yaml
```

## 匯入 db
```
conf/sql/db-schema.sql
```

## 透過 docker-compose 啟動
```
cd this_project
docker-compose up -d --build
```
