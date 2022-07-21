# URL shortener project

---

simple project with web client gui, take a long url and generate the short url with two hour expiration.

## Built using rich tech-stack:

---

- Api-server built using [Go-fiber](https://gofiber.io/).
- [Redis](https://redis.io/) as cache.
- [Docker](https://docker.com/) for managing application.

## Installation & setup :-

---

- Go,Docker,Docker compose & Make should be pre-installed.
- Clone the repository: `git clone https://github.com/erfidev/url_shortener`.
- Run the command `make init` (this will install go modules).
- Create a new file .env under root directory & copy the env variables from .env.example
- Run the command `make run_dependencies` (this will start the required docker containers - redis).
- Run the command `docker ps` to ensure all the four containers are up & running.
- Open a new terminal & run the command `make run` to spin up the server.

# Contributing beers:

---

- Performance improvements, bug fixes, better design approaches are welcome. Please discuss any changes by raising an issue, beforehand.

# Maintainer :sunglasses:

---

By Erfan Hanifezade 2022 jul

[Linkedin](https://www.linkedin.com/in/erfan-hanifezade-07239b201/) <br>
[Email](erfanhanifezade@gmail.com)
