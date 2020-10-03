# post-delivey system

## Prerequisites (besides Git, of course)
- Docker Desktop or any Docker compatible app/manager
- Go Tools

## CI
If the build/test of a PR fails, it won't be allowed to be merged to master. Go to [actions] tab for info about the CI execution.

## Testing
From your CI or the console use for executing the developer tests:

```bash
go test github.com/rjar2020/post-delivery/service
```

## Running the app

### Building the docker image
To build post delivery system Ubuntu (Groovy) image, ***from the project root*** use:

```bash
docker build -t postback-delivery .
```

When the image is built, you can work there:

```bash
#Running a container of postback-delivery for working on it
docker run -it -p 4000:4000 -p 8080:8080 --entrypoint bash postback-delivery
#Start the app
./home/rjar/start-postback-in-image
```
Use ***Control + c*** to stop the app and exit to leave the container

Or you can run a postback-delivery based container to use it from your host:

```bash
#To see the output of the app (Preferred for troubleshooting and also demo)
docker run -p 4000:4000 -p 8080:8080 postback-delivery
#In background mode
docker run -d -p 4000:4000 -p 8080:8080 postback-delivery
```

Once you are here, POST localhost:4000/postback endpoint can be used to produce messages to kafka. See [postman/delivery-postback.postman_collection.json]

Also PHP Ingester form will be available in localhost:8080 and will allow to deliver postbacks to kafka.

Other useful commands:
```bash
#Finding the id of the postback-delivery based container 
docker ps
#Pause the container
docker pause <container_id>
#Resume the container
docker unpause <container_id>
```

The image is available in https://hub.docker.com/repository/docker/rjar2020/postback-delivery
Please request collaborator access if you want to work with it.

For collaborators/mantainers, to publish the latest version of the image:
```bash
docker tag postback-delivery:latest rjar2020/postback-delivery:latest
docker push rjar2020/postback-delivery:latest
```

### Go and Kafka Local / Dev env
To start post delivery system, ***from the project root*** use:

- For spinning up all the dependencies (kafka platform, etc)
```bash
./load-and-execute-dependencies
```
You can access confluentinc/cp-enterprise-control-center via http://localhost:9021/ to visualize kafka stats
Including: topics, messages, consumer groups, etc.

- For starting the app
```bash
go run github.com/rjar2020/post-delivery
```
Once you are here, POST /postback endpoint can be used to produce messages to kafka. See [postman/delivery-postback.postman_collection.json]

Use ***Control + c*** to stop the app

- For stopping all the dependencies plus deleting related containers 
```bash
./stop-and-delete-dependencies
```
Once this is exectuted the state persisted in kafka is gone.

## Notes

### to-do
- Observability (Prometheus, Grafana, etc.)
- Kafka config or using admin client at topic creation to include a proper replication factor. As we are starting 2 brokers, should be a replication factor of 3.
- Github CD to build and push image to Dockerhub

### to-do from code review

- The interpretation of the project was a little flawed, we expected a simple PHP server accepting requests and pushing messages to a Kafka (or Redis) instance, and a Go application pulling from that queue and sending out requests to 3rd parties based on the original request. We allow for some interpretations but reaching out for clarity if the instructions aren't understood is the correct path.
```
- Create a PHP controller instead of the form.
- Nice to have: try again a PHP framework/lib/component that integrates with kafka locally
```
- I'd expect to see the actual application(s) being started in docker-compose
```
- Include slim dockerfile for local go related components and start it in docker compose along the other infra components.
- Include slim dockerfile for local PHP related components and start it in docker compose along the other infra components.
- Nice to have: multiple kafka brokers in local docker compose
```
- There is a certificate missing to be able to run locally (./resources/my_Cert.crt)
```
- Generate image out of the VPN without needing the certs (A nice to have as apparently, as according to the feedback the Ubuntu image wasn't that important)
```

### Log
- In PROD, postback endpoint and postback consumer should live in separate containers, kafka should have its own cluster (proper backups if needed, rack awareness config, etc) and instead of 8 consumers in one Pod 8 Pods with one consumer each should be started.
- Stressing out kafka config, it needs way more work to be prod ready, but this could be done independently from the Go service/components.
- PHP and kafka integration existing projects/modules usage is not straight forward and requires a lot of infrastructure/plugins, so I decided to create a Go endpoint to facilitate this integration.
- Also PHP tools are messy, but maybe this is just a personal opinion comming from someone trying Go tools (I really like this self-contained ecosystem) and relying on the JVM ecosystem most of the time (Yes, variety, not self-contained like Golang but with clear conventions to follow)
- Despite I really rely on TDD to orient my design and improve speed of development, when experimenting a new tools ecosystem, I had lean to implement first. So, much likely a huge refactor it's gonna be needed in the after-match.
- I couldn't make docker-compose work in the ubuntu image, something is going on with Python, so I decided to install plain kafka infra.
- Certs and docker images are always fun!

[postman/delivery-postback.postman_collection.json]: https://github.com/rjar2020/post-delivery/blob/master/postman/delivery-postback.postman_collection.json
[actions]: https://github.com/rjar2020/post-delivery/actions
