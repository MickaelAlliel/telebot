FROM node:16.15.0-alpine as build

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn

COPY . ./
RUN yarn lint:ci
RUN yarn build

FROM nginx:alpine as release
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]