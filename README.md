## Twitter clone

> **Note**
> This is a work in progress application that aims to clone the basic functionalities of Twitter and to learn Go and Vue 3.

The frontend is built with Nuxt 3 and the backend is built with `chi` router, `gorm` ORM and a postgres database.

### Features

- [ ] User authentication
  - [x] Sign-up
  - [x] Login
  - [ ] Logout
- [ ] User profile
  - [ ] View own profile and others
  - [ ] View others profile
  - [ ] Edit own profile
- [ ] Tweet
  - [x] Create
  - [ ] Edit own tweets
  - [x] Display latest tweets on timeline
  - [ ] Like
  - [ ] Comment
  - [ ] Retweet
  - [ ] Bookmark
- [ ] Following
  - [ ] Follow users
  - [ ] Show followers
  - [ ] Show followings

### Getting started

- Go >=1.20
- Node.js >=18.x
- pnpm >=8.3.x

### Installation

Clone the repository

```sh
$ git clone https://github.com/0xYami/twitter
```

Setup environment variables, create the two following `.env` files and replace with your config.

For the frontend

```sh
# ./ui/.env
NUXT_PUBLIC_COOKIE_NAME=
NUXT_PUBLIC_SERVER_BASE_URL=
```

For the backend

```sh
# ./cmd/twitter/.env
ADDRESS=
PORT=
COOKIE_NAME=
TOKEN_SECRET=

PGHOST=
PGUSER=
PGPASSWORD=
PGDATABASE=
PGPORT=
```

To run the backend

```sh
$ cd ./cmd/twitter
$ go run .
```

To run the frontend

```sh
$ cd ./ui
$ pnpm dev
```

### License

This project is licensed under the [MIT License](https://opensource.org/license/mit/)
