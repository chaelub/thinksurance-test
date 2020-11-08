FROM postgres:10-alpine
MAINTAINER chaelub <chaelub@gmail.com>

COPY init.sql /docker-entrypoint-initdb.d/


ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
