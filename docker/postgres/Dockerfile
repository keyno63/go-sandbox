FROM postgres:13-alpine
ENV LANG ja_JP.UTF-8
ENV POSTGRES_DB development
ENV POSTGRES_USER keyno63
ENV POSTGRES_PASSWORD password1
COPY sql/*.sql /docker-entrypoint-initdb.d/
