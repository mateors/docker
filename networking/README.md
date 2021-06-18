# Docker Networking
> Docker networking allows you to attach a container to as many networks as you like. You can also attach an already running container.

> `docker exec -it -u root <containerid> sh`\
> `ping webapp`\
![ping-container](../screenshots/ping-container.png)

> `docker info`\
![docker_network_driver](../screenshots/docker_network_driver.png)

> `docker network ls`\
![docker_default_network_driver](../screenshots/docker_default_network_driver.png)

> `docker run -it alpine:latest sh`\
![docker_run_new_container_interactive_mode](../screenshots/docker_run_new_container_interactive_mode.png)


> `docker run -d --network host --name mynginx nginx:latest` (Not working in windows)\
> `docker stop mynginx`