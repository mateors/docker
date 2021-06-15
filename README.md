# Docker App Development & Command Practice

## What is Docker?
> A platform for building,running and shipping applications. in a consistent manner so if your application works on your development machine it can run and function the same way on other machine.

> **Container:** An isolated environment for running an application.

> **Container states** – A container can be in one of four states: created,running, paused, exited, restarting.

## Docker Image: 
> Docker images are The blueprints of our application which form the basis of containers. We use docker pull command to download an image.

* A standalone, executable package that can be run in a container.

* A Docker image is a binary that includes all of the requirements for running a single Docker container, as well as metadata describing its needs and capabilities.

* An image includes everything that is needed to run an application, including the application's executable code, any software on which the application depends, and any required configuration settings. You can build your own images (using a Dockerfile) or use images that have been built by others and then made available in a registry (such as Docker Hub).

* To build an image from a Dockerfile you use the docker build command.

* To run an image in a container you use the docker run command.


> **Containers** - Created from Docker images and run the actual application. After downloading the image We create a container using docker run command. A list of running containers can be seen using the docker ps command.

> All possible commands are listed in the following google docs

[Docker docs](https://docs.google.com/document/d/1aXqP3HGMoaD-zOfmsQFjbuMB13LMNEVg1bZNuBtM-wM/edit?usp=sharing)

