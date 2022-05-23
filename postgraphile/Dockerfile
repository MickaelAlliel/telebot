FROM node:16.15.0-alpine
WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn

COPY . ./
RUN yarn build

EXPOSE 5000

ENTRYPOINT [ "node" ]
CMD ["dist/server.js"]