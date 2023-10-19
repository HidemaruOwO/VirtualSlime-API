# VirtualSlime-API

## 用法

### 设置

    git clone https://github.com/HidemaruOwO/VirtualSlime-API
    cd VirtualSlime-API

    # Edit .env file
    vim .env

### 测试开发（本地主机：3001）

```bash
VIRTUALSLIME_DIR=~/Code/VirtualSlime DEBUG=true go run src/main.go
```

## 使用产品

### 建造

```bash
go build -ldflags="-s -w" -trimpath
mv main virtualslime
```

### 开始

```bash
APP_ENV=production ./virtualslime
```

## 环境

-   `DOMAIN`：（需要）字符串值：`Your domain name (ex. v-sli.me)`
-   `APP_ENV`：（产品只需）字符串值：`production`
-   `PORT`：（可选）int16 值：`App listen port (default: 3000)`

### 为了发展

-   `VIRTUALSLIME_DIR`：（开发只需） string 值：`VirtualSlime directory (ex. ~/Code/VirtualSlime)`

## API文档

-   /v1/帖子


    GET api.v-sli.me/v1/posts?q=<SearchWord>

`?q`: 搜索文字 :`string+string+string....`
