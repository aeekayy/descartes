FROM node:16-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY tsconfig*.json ./
COPY src src
RUN npm run build

FROM node:16-alpine
ENV NODE_ENV=production
RUN apk add --no-cache tini
WORKDIR /usr/src/app
COPY package*.json ./
RUN chown -R node:node .
USER node
RUN npm install
COPY --from=builder /app/lib/ lib/
ENTRYPOINT [ "/sbin/tini","--", "node", "lib/worker.js" ]