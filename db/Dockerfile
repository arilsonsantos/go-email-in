FROM migrate/migrate:latest

# Copy all db files
COPY ./migrations /migrations

ENTRYPOINT [ "migrate", "-path", "/migrations", "-database"]
CMD ["postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable up"]
