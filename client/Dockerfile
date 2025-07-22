FROM node:18-alpine as builder

WORKDIR /app

COPY package.json package.json
COPY yarn.lock yarn.lock

RUN yarn

COPY . .

RUN yarn build

FROM node:18-alpine

WORKDIR /app

COPY --from=builder /app/dist dist

RUN yarn global add pm2

EXPOSE 3000

CMD ["pm2", "serve", "/app/dist/", "--spa", "--port", "3000", "--no-daemon"]
