# VirtualSlime-API

## Usage

### Setup

```
git clone https://github.com/HidemaruOwO/VirtualSlime-API
cd VirtualSlime-API

# Edit .env file
vim .env
```

### Test dev (localhost:3001)

```bash
VIRTUALSLIME_DIR=~/Code/VirtualSlime DEBUG=true go run src/main.go
```

## Use Product

### Build

```bash
go build -ldflags="-s -w" -trimpath
mv main virtualslime
```

### Start

```bash
APP_ENV=production ./virtualslime
```

## Environment

- `DOMAIN`: (need) string Value: `Your domain name (ex. v-sli.me)`
- `APP_ENV`: (product only need) string Value: `production`
- `PORT`: (optional) int16 Value: `App listen port (default: 3000)`
- `DEBUG`: (optional) bool Value: `Make Debug Mode (ex. true or false, default: false)`

### For development

- `VIRTUALSLIME_DIR`: (development only need) string Value: `VirtualSlime directory (ex. ~/Code/VirtualSlime)`

## API Documents

- /v1/posts

```
GET api.v-sli.me/v1/posts?q=<SearchWord>
```

`?q` : Search text : `string+string+string....`
