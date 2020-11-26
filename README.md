<h1 align="center">Passwd Box</h1>
<p align="center">The sync 🔁 server for passwd</p>

## Building

### Requirements

- Golang
- Git

### Running locally

- Clone this repo wherever you want
- Copy the `.env.sample` file to `.env`
- Change the variables inside `.env`
- `go run cmd/passwdbox/main.go`

### Building a binary

- Clone this repo
- Copy the `.env.sample` file to `.env`
- Change the variables inside `.env`
- `go build -o passwdbox cmd/passwdbox/main.go`
- `./passwdbox`

### Running without `.env`

- Build a binary
- Set the environment variables for `SECRET_KEY`, `JWT_SECRET`, `MAX_USERS` and `LISTEN_ADDRESS`
- Run the binary using `./passwdbox -use-env=false`
