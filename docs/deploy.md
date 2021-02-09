# How to deploy

### Requirements

- A server with persistent storage
- 1 or more CPU cores
- ~20MB free RAM

### Steps to deploy

- Build the server ([building.md](./building.md))
- Make data directories with `mkdir -pv data/uploads` in the directory where you will run the binary
- Copy `.env.sample` as `.env` (can be skipped if using env variables)
- Set the variables in `.env` (can be skipped if using env variables)
  - `MAX_USERS` - Maximum number of users registered on an instance
  - `SECRET_KEY` - A key shared with all the users. It should be unique.
    This is required to do any operation on the server.
    Without this, server will reject the request with a `401 Unauthorized` HTTP code
  - `JWT_SECRET` - This is used to sign the short lived authentication tokens and therefore be long and random. You can generate a random secret using the command `openssl rand -hex 64`
  - `LISTEN_ADDRESS` - The HTTP address to start the server on
- Run the server with or without `-use-env=false/true` depending on your configuration.

### Recommendations

- When running outside your local network, always use a reverse proxy with HTTPs
- Regularly take offsite backups of the `data` directory. This directory contains the `sqlite` database and the encrypted data for each user

### Caveats

- Currently heroku based deployments are not supported due to the lack of persistent storage
- Only DB supported as of [9th feb 2021] is SQLite
