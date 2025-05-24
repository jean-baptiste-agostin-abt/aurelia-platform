# Aurelia Platform

This repository contains a minimal full-stack setup for the Aurelia project.
It includes a Go backend and a React + TypeScript frontend with Docker support.

## Requirements
- Go 1.24
- Node 18
- Docker

## Setup
Create a `.env` file at the project root to override environment variables if needed.
Important variables:
- `DB_HOST` (default `db`)
- `DB_USER` (default `root`)
- `DB_PASS` (default `password`)
- `DB_NAME` (default `aurelia`)
- `JWT_SECRET` (default `secret`)

### Using Make
- `make build` – build Docker images
- `make test` – run backend and frontend tests
- `make up` – start services with docker-compose
- `make down` – stop services

### API Endpoints
- `POST /auth/signup` – create user
- `POST /auth/login` – login and receive JWT
- `POST /families` – create family
- `POST /capsules` – create capsule
- `GET /capsules/:id` – get capsule
- `GET /feed` – list events
- `GET /events` – list events
- `POST /legacyguard/trigger` – trigger legacy guard

## Development
Run `make up` and visit the frontend at `http://localhost:3000` and backend at `http://localhost:8080`.
