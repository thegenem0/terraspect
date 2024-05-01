FROM node:18-alpine AS dev

WORKDIR /app

COPY package.json package-lock.json* yarn.lock* ./

RUN yarn install

COPY . .

EXPOSE 5173

CMD ["yarn", "dev"]