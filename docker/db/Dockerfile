FROM postgres:alpine
COPY ./db/scripts/init-tables.sql /docker-entrypoint-initdb.d/init.sql
EXPOSE 5432