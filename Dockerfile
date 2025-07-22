FROM golang:1.22.3-alpine

RUN apk update && apk add --no-cache curl bash sudo aws-cli

RUN curl -1sLf 'https://dl.cloudsmith.io/public/infisical/infisical-cli/setup.alpine.sh' | sudo -E codename=v3.9 bash

RUN apk update && sudo apk add infisical

WORKDIR /code/

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64

ARG infisical_token
ARG infisical_project_id
ARG infisical_env

ARG aws_access_key_id
ARG aws_secret_access_key

RUN aws configure set aws_access_key_id ${aws_access_key_id}
RUN aws configure set aws_secret_access_key ${aws_secret_access_key}

ENV INFISICAL_TOKEN=${infisical_token}
RUN infisical export --projectId=${infisical_project_id} --env=${infisical_env} --format=dotenv > .env

RUN source .env && go run github.com/steebchen/prisma-client-go migrate deploy
RUN go run github.com/steebchen/prisma-client-go generate

RUN go mod tidy
RUN go build -o main

CMD [ "./main" ]
