FROM node:alpine

EXPOSE 3000

# create destination directory
RUN mkdir -p /usr/src/nuxt-app
WORKDIR /usr/src/nuxt-app

# update and install dependency
RUN apk update && apk upgrade
#RUN apk add git

# copy the app, note .dockerignore
COPY . /usr/src/nuxt-app/
RUN npm install
RUN npm run build

CMD [ "npm", "start" ]

#Para construir contenedor correr:
# docker build -t middle_front .
# docker network create a_default
# docker run -d --name a-frontend --network a_default -p 3000:3000 middle_front
