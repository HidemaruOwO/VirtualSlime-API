# VirtualSlime-API

## 用法

### 設定

    git clone https://github.com/HidemaruOwO/VirtualSlime-API
    cd VirtualSlime-API

    # Edit .env file
    vim .env

### 測試開發（本機：3001）

```bash
VIRTUALSLIME_DIR=~/Code/VirtualSlime DEBUG=true go run src/main.go
```

## 使用產品

### 建造

```bash
go build -ldflags="-s -w" -trimpath
mv main virtualslime
```

### 開始

```bash
APP_ENV=production ./virtualslime
```

## 環境

-   `DOMAIN`：（需要）字串值：`Your domain name (ex. v-sli.me)`
-   `APP_ENV`：（產品只需）字串值：`production`
-   `PORT`：（可選）int16 值：`App listen port (default: 3000)`

### 為了發展

-   `VIRTUALSLIME_DIR`：（開發只需） string 值：`VirtualSlime directory (ex. ~/Code/VirtualSlime)`

## API文件

-   /v1/帖子


    GET api.v-sli.me/v1/posts?q=<SearchWord>

`?q`: 搜尋文字 :`string+string+string....`
