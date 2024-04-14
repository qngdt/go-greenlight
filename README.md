# The Greenlight Project

Following the famous book [Let’s Go Further](https://lets-go-further.alexedwards.net/)
---
### This repository uses some different libraries compared to the book:
- [pgx](https://github.com/jackc/pgx) instead of [pq](https://github.com/lib/pq) for PostgreSQL: `pg` is unmaintained and `pgx` is a more modern and faster alternative.
- [net/http](https://pkg.go.dev/net/http) instead of [httprouter](https://github.com/julienschmidt/httprouter) for the HTTP router: `net/http` is a standard library and is enough with the enhancements of Go 1.22
- [dbmate](https://github.com/amacneil/dbmate) instead of [migrate](https://github.com/golang-migrate/migrate) for database migrations: `dbmate` is a more modern alternative and is easier to use.
---
**Update: Finished April 14, 2024**
