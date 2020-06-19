# Tiến lên (Thirteen)

This is an online-enabled version of Tiến lên ("Thirteen").

The server component is built in Golang, and the UI is built in TypeScript using the VueJS reactive framework.

# How to run it out of the box

This section assumes you have your Golang and NPM environments set up and ready to go.

## Development mode

To run the server:

```bash
$ cd api
$ go run main.go
```

And the UI:

```bash
$ cd ui
$ npm install // first run only
$ npm run dev
```

By default, the application will be available at `http://127.0.0.1:26000`.

## "Production" mode

First, build the artifacts:

```bash
$ cd api
$ go build -o ../tienlen-server main.go
$ cd ../ui
$ npm install // first run only
$ npm run build
```

Then run the server (which serves the UI files):

```bash
$ cd ..
$ ./tienlen-server -addr "0.0.0.0:26000" -ui "ui/dist"
```

The application will be available at `http://127.0.0.1:26000` and other IP address(es) assigned to the host.

# Things to improve

- vote to kick a player
- kick newly connected clients who take too long to self-identify
- move timer
- different card faces
- lobby password
- text chat
- more unit tests
