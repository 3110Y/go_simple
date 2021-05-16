FROM golang:latest

WORKDIR $GOPATH/src

ARG UID=1000
ARG GID=1000
ARG APP_ENV='dev'
ENV UID=${UID}
ENV GID=${GID}
ENV APP_ENV=${APP_ENV}

RUN getent passwd $UID || (groupadd --gid $GID user && useradd --uid $UID --gid user --shell /bin/bash --create-home user)
USER $UID