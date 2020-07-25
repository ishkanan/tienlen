# Tiến lên (Thirteen)

This is an online-enabled version of Tiến lên ("Thirteen").

The server component is built in Golang, and the UI is built in TypeScript using the VueJS reactive framework.

# Quick up and run

Docker builds of `master` are available on [DockerHub](https://dockerhub.com/ishkanan/tienlen). To run the game:

```bash
$ docker pull ishkanan/tienlen:latest
$ docker run ishkanan/tienlen:latest -p 26000:26000
```

Or you can run the image using your favourite orchestration tool (e.g. `docker-compose`).

The game will be available at `http://localhost:26000` and other IPs the host has been assigned.

# Manual build and run

To build the game using the `Dockerfile`, simply do:

```bash
$ docker build -t <your repo>:<your tag> . --force-rm
```

To build and run the game directly, do the following:

1. Set up a [Golang environment](https://golang.org/doc/install)

2. Set up an [NPM environment](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)

3. Clone this repository into `<Go path>/src/github.com/ishkanan/tienlen`

4. Run the following commands:

```bash
$ cd api
$ go build -o ../tienlen-server main.go
$ cd ../ui
$ npm ci    // first run only
$ npm run build
```

5. And finally, start the game:

```bash
$ cd ..
$ ./tienlen-server -addr "0.0.0.0:26000" -ui "ui/dist"
```

The game will be available at `http://localhost:26000` and other IPs the host has been assigned.

# Development

As mentioned in the [Manual build and run](#manual-build-and-run) section, ensure you have your [Golang](https://golang.org/doc/install) and [NPM](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) environments set up, and that you've cloned this repository to `<Go path>/src/github.com/ishkanan/tienlen`.

To run the API server in development mode:

```bash
$ cd api
$ go run main.go
```

To run the UI in development mode, we use the [built-in Webpack DevServer](https://webpack.js.org/configuration/dev-server):

```bash
$ cd ui
$ npm ci    // first run only
$ npm run dev
```

The game will be available at `http://localhost:26000`, with the UI server proxying the API requests.

# Improvements

Feel free to submit PRs for changes, or fork to your heart's content.
