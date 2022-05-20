### Docker

---

##### Some core concepts

---

1. Docker Clients : After starting , docker will run as a daemon. And you can use ` docker + command ` to operate it . 
2. Image and container : In local, docker has some image and you can create corresponding container from images. Image can be created by Docker File, it can be seen as a root files system.

3. Registry : remote repository, provide users with some ready-made image so that we can directly pull them.
4. Private Registry : your own registry build by yourself. 

---

##### some commands

---

###### system:

- Start/Stop docker : ``systemctl start/stop docker``

- Check status : ``systemctl status docker``

###### Image : 

- Check local images : ``docker images`` , option -q to only show id.
- Search image : ``docker search xxx``
- Download image : ``docker pull xxx``
- Delete image : ``docker rmi Imageid`` , delete all images : ``docker rmi $(docker images -q)``

###### Container :

- start a container : `docker run -i`

  - -i : always run the container, without it container will close once there is no client connect.
  - -t : request a terminal for container.
  - -d: run in the background.
  - --name : the name of container.

  Example : ``docker run -it -name=test centos:7 /bin/bash`` ( /bin/bash is the command who will be executed when container be created, here is to open a terminal )

  If you create a container in with parameter -d, use command ``docker exec`` enter it.

  - exit : quit container.

- show containers : ``docker ps -a``

- Stop container: `docker stop XXX`

---

##### Data volume

---

###### Use parameter -v to add a data volume :

Example : ``docker run -it -name=test -v /root/data /root/data_container centos:7 /bin/bash``

###### data volume container to realize data volume between two containers :

Create a new container :

``docker run -it --name=c3 -v /volume centos:7 /bin/bash``

Then use parameter --volumes-from to add other container to this data volume container.

``docker run -it --name=c1 --volumes-from c3 centos:7 /bin/bash``

``docker run -it --name=c2 --volumes-from c3 centos:7 /bin/bash``

###### Use parameter -p to add a Port mapping :

``docker run -id -p 3307:3306 --name my_mysql -v $PWD/logs:/logs mysql:5.6``

---

##### Dockerfile 

---

###### keywords :

- From : base image
- Label : some comments
- ENV : environment variables
- Run : run some commands
- Expose : expose some ports
- CMD : command will be executed when container starts
- COPY/ADD : copy/add a file 
- VOLUME : add a data volume
- WORKDIR : the path when you enter in a container

---

##### Docker Compose

---

Docker compose is a way to create a cluster of service by several containers in one time.

First, create a file named ``docker-compose.yml`` and that use command ``docker-compose up``.

The format of the content in ``docker-compose.yml`` is similar with docker file.

Example:

```shell
version: "3.7"

services:
  app:
    image: node:12-alpine
    command: sh -c "yarn install && yarn run dev"
    ports:
      - 3000:3000
    working_dir: /app
    volumes:
      - ./:/app
    environment:
      MYSQL_HOST: mysql
      MYSQL_USER: root
      MYSQL_PASSWORD: secret
      MYSQL_DB: todos

  mysql:
    image: mysql:5.7
    volumes:
      - todo-mysql-data:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: todos

volumes:
  todo-mysql-data:
```

Some key words : 

- Image : refer a image
- command : = run in docker file
- Ports : = -p option when you run docker to start a container
- working_dir = workdir in docker file
- Volumes : = -v option when you run docker to start a container
- environment : environment variables