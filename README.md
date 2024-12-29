# Media Website

This is a simple media hosting website, like Getty or Shutterstock, written in React and Go.

## Development

- For running with build: `docker compose -f docker-compose.local.yaml up --build`
- For running with cache: `docker compose -f docker-compose.local.yaml up`
- Go to [http://localhost:3000](http://localhost:3000) in a browser.

### Frontend Non-Docker Instructions

- Install [Node.js](https://nodejs.org/en) version **23.1.0**
- Run `npm install` in the **/frontend** directory.
- Run `npm start` in the **/frontend** directory.

### Backend Non-Docker Instructions

- Install [Go](https://go.dev/dl/) version **1.23.4**
- Install [Air](https://github.com/air-verse/air) using Go `go install github.com/air-verse/air@latest`
- Run each service in the **/backend** directory using the `air` command

## Contributing

### Style Rules

- Line break at the end of files
- 2 space indenting
