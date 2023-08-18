# VirtualSlime-API

## Usage

### Setup

```
git clone https://github.com/HidemaruOwO/VirtualSlime
git submodule update --init

yarn install
yarn run cache-posts

cd VirtualSlime-API
yarn install

yarn build
NODE_ENV=production yarn start
```

### Test dev (localhost:3001)

```bash
yarn dev
```

## Use Product

### Build

```bash
yarn build
```

### Start

```bash
NODE_ENV=production yarn start
```

## API Documents

- /api/posts

```
GET api.v-sli.me/api/posts?q=<Query>
```

`?q` : Search text : `string+string+string....`
