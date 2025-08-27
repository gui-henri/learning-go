FROM node:current-alpine3.22

WORKDIR /frontend

EXPOSE 3000

COPY . .

RUN npm install

CMD ["npm", "run", "dev"]