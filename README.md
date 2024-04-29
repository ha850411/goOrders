# 使用說明

## 使用技術
<img src="/static/images/sourceImg/Golang.png" alt="Golang" width="100" />
<img src="/static/images/sourceImg/Vue.png" alt="Vue" width="100" />
<img src="/static/images/sourceImg/HTML.png" alt="HTML" width="100" />

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
