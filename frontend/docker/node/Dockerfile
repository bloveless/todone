FROM node:10.14-alpine

RUN npm i -g npm yarn

USER node

RUN mkdir -p /home/node/app
WORKDIR /home/node/app

COPY --chown=node:node ./package.json ./yarn.lock /home/node/app/

RUN yarn install

COPY --chown=node:node . /home/node/app/

EXPOSE 8080

CMD ["yarn", "start"]
