# Blog 博客系统

[English](README.md) | 中文

基于 Go + React + PostgreSQL 的现代化博客系统，Docker 一键部署。

## 技术栈

**后端：** Go 1.25 + Gin + PostgreSQL 18 + GORM + JWT

**前端：** React 19 + Vite + TailwindCSS + TypeScript

## 效果图
<img src="screenshot/img1.png" alt="首页" width="900">
<img src="screenshot/img2.png" alt="文章详情" width="900">
<img src="screenshot/img3.png" alt="AI聊天" width="900">
<img src="screenshot/img4.png" alt="后台概览" width="900">
<img src="screenshot/img5.png" alt="后台文章" width="900">
<img src="screenshot/img6.png" alt="媒体文件" width="900">

---

## 快速开始

### 前置要求

- Docker 20.10+
- Docker Compose 2.0+
- 2GB+ 内存

安装 Docker:
```bash
curl -fsSL https://get.docker.com | sh
sudo usermod -aG docker $USER
```

### 部署步骤

```bash
# 1. 克隆项目
git clone https://github.com/voocel/blog.git blog
cd blog

# 2. 初始化目录
./scripts/init.sh

# 3. 配置环境变量（必填）
cp .env.example .env
vim .env  # 设置 POSTGRES_PASSWORD

# 4. 启动服务
docker compose up -d

# 5. 创建管理员账号
docker compose run --rm backend ./blog -create-admin

# 6. 访问
# http://localhost 或 http://your-server-ip
```

**更新部署：**
```bash
git pull
docker compose up -d --build
docker compose up -d --build --no-deps backend frontend
```

**配置说明：**

根目录 `.env`（部署用）：
- `POSTGRES_PASSWORD`（必填）PostgreSQL 容器初始化密码

后端运行配置：
- `config/config.yaml`（后端单一配置源，包含 JWT/数据库/App/HTTP）

**前端构建可选项**
- 部署：在 `web/.env` 写 `VITE_API_KEY` / `VITE_DEFAULT_HOMEPAGE`，再 `docker compose build frontend && docker compose up -d frontend`
- 本地开发：在 `web/.env` 写 `VITE_API_KEY` / `VITE_DEFAULT_HOMEPAGE`（可选 `VITE_API_URL=http://localhost:8080/api/v1`），`npm run dev`

**配置机制：**
- 根目录 `.env.example`：部署环境变量模板（提交）
- 根目录 `.env`：部署环境变量（不提交，Compose 读取）
- `config/example.yaml`：后端模板（提交）
- `config/config.yaml`：后端配置（不提交，init.sh 自动创建）
- `web/.env.example`：前端本地开发模板（提交）
- `web/.env`：前端本地开发配置（不提交）
- 后端优先级：`config.yaml` > 代码默认值（不做后端环境变量覆盖）
- 请保持 `config/config.yaml` 中 `postgres.password` 与根 `.env` 的 `POSTGRES_PASSWORD` 一致

---

## HTTPS 配置

### 前提条件

- 拥有域名
- 域名已解析到服务器 IP

### 配置步骤

```bash
./scripts/setup-https.sh
```

按提示输入域名（可逗号分隔多域，如 `example.com,www.example.com`），脚本会自动：
1. 安装 acme.sh
2. 申请 Let's Encrypt 免费证书
3. 配置 Nginx HTTPS
4. 设置自动续期（90天有效期）

完成后访问：
- **HTTPS**: https://your-domain.com
- **HTTP**: http://your-domain.com（自动跳转 HTTPS）

**证书管理：**
```bash
# 查看证书
~/.acme.sh/acme.sh --list

# 手动续期
~/.acme.sh/acme.sh --renew -d your-domain.com
```

---

## 服务器迁移

### 备份数据（老服务器）

```bash
./scripts/backup.sh
```

生成 `backup_YYYYMMDD_HHMMSS.tar.gz`，包含：
- 数据库
- 静态文件（uploads、avatars）
- 配置文件（.env）
- SSL 证书
- GeoIP 数据库

### 恢复数据（新服务器）

```bash
# 1. 安装 Docker
curl -fsSL https://get.docker.com | sh

# 2. 克隆项目
git clone https://github.com/voocel/blog.git blog
cd blog
./scripts/init.sh

# 3. 上传备份文件
scp backup_*.tar.gz user@new-server:~/blog/

# 4. 恢复数据
./scripts/restore.sh  # 输入备份文件名
```

**迁移注意事项：**
- 迁移包大小：通常 20-200MB
- 不包含 Docker 镜像，新服务器会重新构建
- SSL 证书会自动迁移，需确保域名 DNS 已解析到新 IP

---

## 常用命令

```bash
# 查看服务状态
docker compose ps

# 查看日志
docker compose logs -f
docker compose logs -f backend
docker compose logs -f nginx

# 重启服务
docker compose restart

# 停止服务
docker compose stop

# 更新部署
git pull
docker compose up -d --build

# 进入容器
docker compose exec backend sh
docker compose exec postgres psql -U postgres blog

# 手动备份数据库
docker exec blog-postgres pg_dump -U postgres blog > backup.sql

# 手动恢复数据库
docker exec -i blog-postgres psql -U postgres blog < backup.sql
```

---

## 故障排查

### 服务无法启动

```bash
# 查看日志
docker compose logs -f

# 检查端口占用
netstat -tunlp | grep -E '80|443|8080|5432'

# 重新构建
docker compose down
docker compose up -d --build
```

### 数据库连接失败

```bash
# 检查数据库状态
docker exec blog-postgres pg_isready -U postgres

# 检查后端配置文件
docker compose exec backend cat /app/config/config.yaml

# 重启数据库
docker compose restart postgres
```

### 端口被占用

修改 `docker-compose.yml`:
```yaml
nginx:
  ports:
    - "8090:80"   # 改为其他端口
    - "8443:443"
```

### SSL 证书申请失败

可能原因：
1. 域名 DNS 未正确解析 - 运行 `ping your-domain.com` 检查
2. 80 端口未开放 - 运行 `ufw allow 80/tcp`
3. 防火墙阻止连接 - 运行 `ufw status` 检查

---

## 可选配置

### GeoIP 数据库

用于显示访客地理位置（可选）：

1. https://github.com/P3TERX/GeoLite.mmdb
2. 下载 GeoLite2-City.mmdb
3. 放到 `config/GeoLite2-City.mmdb`
4. 重启服务：`docker compose restart backend`

### 安全建议

1. 修改默认管理员密码
2. 使用强密码（.env 文件）
3. 配置 HTTPS
4. 配置防火墙
   ```bash
   ufw allow 80/tcp
   ufw allow 443/tcp
   ufw allow 22/tcp
   ufw enable
   ```
5. 定期备份数据
6. 定期更新系统和 Docker

---

## 开发

### 后端开发

```bash
go run cmd/blog/main.go
```

### 前端开发

```bash
cd web
npm install

# 配置前端环境变量（本地开发需要）
cp .env.example .env
# 编辑 .env，配置 VITE_API_URL=http://localhost:8080/api/v1
# 可选：VITE_DEFAULT_HOMEPAGE=blog

npm run dev
```

---

## License

Apache-2.0
