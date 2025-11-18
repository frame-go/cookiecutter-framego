FROM golang:{{ cookiecutter.go_version }}-bookworm

COPY . src/{{ cookiecutter.project_name }}
RUN cd src/{{ cookiecutter.project_name }} && make build

FROM debian:bookworm-slim

ENV APP_NAME app
WORKDIR /app/{{ cookiecutter.project_name }}
RUN apt update && apt install -y ca-certificates
COPY --from=0 /go/src/{{ cookiecutter.project_name }}/bin/{{ cookiecutter.app_package_name }} /bin/app
CMD app
