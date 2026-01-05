# GoAway - URL Shortener

## Overview
I developed this project to explore Go, HTTP protocols, Redis, PostgreSQL, and microservice architecture. It’s a learning project, so it might not be perfect and may contain some bugs =)

## Tech Stack
- **Language**: Go 1.2x  
- **HTTP Framework**: Gin Gonic  
- **Database**: PostgreSQL (GORM ORM)  
- **Cache / Sessions**: Redis
- **Containerization**: Docker, Docker Compose  

## Features
- Short URL generation.
- Redirection with click tracking (via goroutines).
- User dashboard: list links, view stats, delete links.
- User authorization via Redis (sessions).
- Access control: users can only manage their own links.

## Deployment
### 1. Environment Setup
```bash
cp .env.example .env
```

### 2. Run Infrastructure
```bash
docker-compose -f configs/docker-compose.yaml up --build
```

Database migrations and required services are applied automatically on startup.

## API Specification
### Public Routes
- `GET /:id` — Redirects to the original URL and increments the click counter.
- `POST /reg` — Registration in service.
- `POST /login` — Login in service.

### Protected Routes (Authentication Required)
- `POST /u/logout` — Logout from service.
- `POST /u/new` — Create a new shortened URL. 
- `GET /u/links` — Retrieve all links owned by the authenticated user.  
- `GET /u/link` — Get detailed information and statistics for a specific link.
- `DELETE /u/link` — Soft delete link from database.

## TODO
- Swagger
- Graceful Shutdown support
- Redis L2 caching for popular links
- Role-Based Access Control (RBAC) and admin interface
- Custom TTL per link
- Better stats (User-Agent, Country, Referer etc.)
- Frontend