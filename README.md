# go-crud

POC — simple in-memory Go CRUD service. Experimental prototype for learning and API design; not production-ready.

## Overview
Small proof-of-concept project demonstrating an in-memory repository pattern, basic validation in a service layer, and simple CRUD operations for books.

## Status
Experimental / prototype.

## Requirements
- Go 1.25.3

## Quick start
Build and run:
```bash
go run ./...
```
## Project layout
- internal/domain — domain models
- internal/repo — repository implementations (including in-memory)
- internal/service — business logic / validation
- cmd / other folders — TODO: add server/router if needed

## Notes
- In-memory store; data is ephemeral.
- Intended for experimentation and learning; interfaces may change.
- Defaults: Go 1.25.3; see go.mod for dependencies.