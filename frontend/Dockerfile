FROM node:20.8
WORKDIR /frontend

COPY package.json vite.config.ts ./
RUN npm install
COPY . .