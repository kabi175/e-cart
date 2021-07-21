FROM node:alpine

WORKDIR /react/src/app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json .

COPY yarn.lock .

RUN yarn 

COPY . ./

EXPOSE 8080

CMD ["yarn", "start"]
