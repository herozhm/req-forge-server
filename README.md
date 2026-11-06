# req-forge-server

> ReqForge 的 Go 后端 —— 把「需求拆解历史」从浏览器 localStorage 升级到服务端持久化。
> 练手定位:**前端(React/req-forge)+ 后端(Go/Gin/GORM/SQLite)的完整全栈闭环**。

## 技术栈

- **Go + Gin** —— HTTP 服务
- **GORM + SQLite**(glebarez/sqlite 纯 Go 驱动,零 CGO)
- 四个 REST 接口:保存 / 列表 / 详情 / 删除

## 快速开始

```bash
# 1. 装依赖(需要 Go 1.22+,未装先 brew install go)
go mod tidy

# 2. 启动(监听 :8080,自动建 req-forge.db)
go run .

# 3. 冒烟测试(另开终端)
curl -X POST http://localhost:8080/api/decompositions \
  -H "Content-Type: application/json" \
  -d '{"requirement":"给净化器加滤芯提醒","result":"{\"summary\":\"test\"}","model":"gpt-4o"}'

curl http://localhost:8080/api/decompositions
```

## 接口

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/decompositions` | 保存(requirement/result/model) |
| GET | `/api/decompositions` | 列表(最近 50 条) |
| GET | `/api/decompositions/:id` | 详情 |
| DELETE | `/api/decompositions/:id` | 删除 |

## 目录结构

```
main.go                      入口 + 路由 + CORS
internal/
├── handler/decomposition.go HTTP 处理器(CRUD)
├── model/decomposition.go    GORM 模型
└── store/db.go               SQLite 初始化 + AutoMigrate
```

## 与 req-forge 前端的对接(下一步)

前端拆解完成后调 `POST /api/decompositions` 保存,加一个「历史」面板调 `GET` 列表——
把拆解结果从「一次性」变成「可回看、可复用」的项目资产。

## 学习要点(边做边理解)

- Go 的 `error` 显式返回 vs JS 的 try/catch
- goroutine 并发(本项目未用到,但这是 Go 与 Node 的最大差异,后续 WebSocket 长连接服务会用到)
- GORM AutoMigrate:结构体 → 表结构的声明式映射
- 纯 Go 的 SQLite 驱动选择(避开 CGO 的坑)
