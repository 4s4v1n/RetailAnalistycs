FROM postgres:latest
WORKDIR /tmp/data

COPY ./data ./

COPY ./migrations/*up.sql /docker-entrypoint-initdb.d
EXPOSE 5432
