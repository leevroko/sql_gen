# Example
PG_PW=postgres_password go run cmd/sql_gen/main.go \
    --db-host localhost \
    --db-port 5430 \
    --db-user postgres_user \
    --db-name postgres_db \
    --db-password-env PG_PW \
    -l Debug \
    docs/config.yaml
