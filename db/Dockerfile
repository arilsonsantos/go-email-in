FROM migrate/migrate:latest

# Copy all db files
COPY ./migrations /migrations

ENTRYPOINT [ "migrate", "-path", "/migrations", "-database"]
CMD ["postgresql://postgres:123@db:5432/go_campaigns?sslmode=disable up"]
