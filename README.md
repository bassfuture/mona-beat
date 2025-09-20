# NFT捕捉游戏项目 - 后端服务

一个基于区块链的NFT捕捉游戏后端服务，提供API接口供前端调用，处理捕捉逻辑并与智能合约交互，自动铸造NFT并发送到用户钱包地址。

## 项目结构

```
nft-capture-game/
├── backend/            # Go后端服务
│   ├── cmd/server/     # 服务器入口
│   ├── internal/       # 内部包
│   │   ├── api/        # API路由和处理器
│   │   ├── service/    # 业务逻辑服务
│   │   ├── repository/ # 数据访问层
│   │   ├── model/      # 数据模型
│   │   └── config/     # 配置管理
│   ├── pkg/            # 公共包
│   │   ├── database/   # 数据库连接
│   │   ├── redis/      # Redis连接
│   │   └── blockchain/ # 区块链交互
│   └── migrations/     # 数据库迁移
├── contracts/          # 智能合约
│   ├── contracts/      # Solidity合约
│   ├── scripts/        # 部署脚本
│   └── test/           # 合约测试
└── docs/               # 项目文档
    ├── 需求文档.md
    └── 技术开发文档.md
```

## 技术栈

### 后端
- Go 1.21 + Gin框架
- PostgreSQL + Redis
- GORM + Wire
- Zap日志 + Viper配置

### 智能合约
- Solidity 0.8.19
- OpenZeppelin
- Hardhat
- ERC-721标准

## 快速开始

### 环境要求
- Node.js 18+
- Go 1.21+
- PostgreSQL 15+
- Redis 7+

### 安装依赖

#### 后端
```bash
cd backend
go mod download
```

#### 智能合约
```bash
cd contracts
npm install
```

### 运行项目

#### 启动后端服务
```bash
cd backend
go run cmd/server/main.go
```

#### 部署智能合约
```bash
cd contracts
npx hardhat deploy --network localhost
```

#### 使用Docker运行完整服务
```bash
docker-compose up -d
```

## 开发指南

详细的开发指南请参考：
- [需求文档](./docs/需求文档.md)
- [技术开发文档](./docs/技术开发文档.md)

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。