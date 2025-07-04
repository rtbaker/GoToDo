# gotodo

A simple front end to the todo api using Vue 3/Vite. Needs a designer to make it look good :-).

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

Edit `.env.developement` and set the `VITE_API_URL` env var to the backend endpoint, e.g.

```
VITE_API_URL=http://192.168.178.217:9090
```

NOTE: we use the machine IP address here rather than localhost so that we don't get into
third party cookie hell when testing from non localhost.

```sh
npm run dev
```

Browse to: `http://192.168.178.217:5173` or whatever your IP address is.

### Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
