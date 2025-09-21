<div align="center">
<h1>Clover</h1>
<img src="./assets/profile.png" width="300" />
<p>
<i>
Easy to run and use, self-hosted web profile & link shortener
</i>
</p>
</div>

## Configuration

Typical `config.json` file looks like this:

```jsonc
{
  "address": ":80",
  "enableTls": false, // put server.crt and server.key in the same folder and change address to :443 to enable TLS
  "avatar": "https://.../avatar.png", // avatar URL
  "name": "My Cool Name",
  "bio": "My Cool Bio",
  "links": [
    // Button example
    {
      "name": "Telegram",
      "url": "https://t.me/...",
      "symlink": "telegram", // used in URL shortener. In this example, https://<your-host>/telegram will redirect to a specific URL
      "background": "#000", // background color
      "foreground": "#fff", // text color
      "icon": "telegram" // Font Awesome brand icon name
    },

    // Label example
    {
      "isLabel": true,
      "name": "MUSIC"
    }
  ]
}
```

## Run using Docker

_Without TLS:_

```sh
$ docker run -d \
  -p 80:80 \
  -v ./config.json:/clover/config.json \
  --name clover eliva1e/clover
```

_With TLS:_

```sh
$ docker run -d \
  -p 443:443 \
  -v ./config.json:/clover/config.json \
  -v ./server.key:/clover/server.key \
  -v ./server.crt:/clover/server.crt \
  --name clover eliva1e/clover
```

## Static page

You can generate static HTML page to use it in any hosting service (e.g. Vercel, GitHub Pages, nginx) using Clover CLI. Please note that URL shortener won't work with static page.

Download the latest Clover CLI in [Releases](https://github.com/eliva1e/clover/releases) or build it yourself and run:

```sh
$ ./clover_cli -config <path-to-config>
```

## TODO

1. Use default net/http instead of chi
2. Add URL shortener for static page
