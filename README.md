# boxd

**Status: Work in Progress**

This is a personal project showcasing how I write Go code for building a RESTful API. The project is **not yet complete** and should be considered a work in progress. It is intended to provide an example of my approach to structuring Go code, handling HTTP requests, and organizing the backend logic.

## Overview

The project aims to create a simple API server using Go (Golang). Although some functionality is implemented, **this is not a production-ready project**. Features, tests, and optimizations are still being developed.

## Features (so far)
- Basic RESTful API setup using Gorilla Mux
- User authentication:
  - User login with JWT authentication
  - User registration with hashed password
- PostgreSQL integration for persistent storage with database migrations
- Redis integration for caching or session management
- Graceful shutdown with context and signal handling

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/soufianiso/boxd.git
   cd boxd
   make migrate 
   make watch
