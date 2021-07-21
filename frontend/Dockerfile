FROM node:13.12.0-alpine

WORKDIR /react/src/app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./
COPY yarn.lock ./

RUN yarn install --silent
RUN yarn install react-scripts@3.4.1 -g --silent

COPY . .

EXPOSE 8080

# start app
CMD ["npm", "start"]
