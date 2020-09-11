# post-delivey system

## Prerequisites (besides Git, of course)
- Docker Desktop or any Docker compatible app/manager
- Go Tools

## Testing
From your CI or the console use for executing the developer tests:

```bash
go test github.com/rjar2020/post-delivery/tests
```

## Running the app

### Local / Dev env
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

### Building the docker image
To build post delivery system Ubuntu (Groovy) image, ***from the project root*** use:

```bash
docker build -t postback-delivery .
```

When the image is built, you can work there:

```bash
#Running a container of postback-delivery for working on it
docker run -it --entrypoint bash postback-delivery
#Start the app
./home/rjar/start-postback-in-image
```
Use ***Control + c*** to stop the app and exit to leave the container

Or you can run a postback-delivery based container to use it from your host:

```bash
#To see the output of the app (Preferred for troubleshooting and also demo)
docker run -p 4000:4000 postback-delivery
#In background mode
docker run -d -p 4000:4000 postback-delivery
```

Once you are here, POST localhost:4000/postback endpoint can be used to produce messages to kafka. See [postman/delivery-postback.postman_collection.json]

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

## Notes

### to-do
- PHP ingester to hit POST /postback endpoint
- Observability
- Kafka config improvements for HA, and start many instances of the app to make the most of it.
- Github CI to run the tests, build and push image to Dockerhub

### Log
- In PROD, postback endpoint and postback consumer should live in separate containers. Also kafka should have its own cluster (proper backups if needed, etc)
- PHP and kafka integration existing projects/modules usage is not straight forward and requires a lot of infrastructure/plugins, so I decided to create a Go endpoint to facilitate this integration.
- Also PHP tools are messy, but maybe this is just a personal opinion comming from someone trying Go tools (I really like this self-contained ecosystem) and relying on the JVM ecosystem most of the time (Yes, variety, not self-contained like Golang but with clear conventions to follow)
- Despite I really rely on TDD to orient my design and improve speed of development, when experimenting a new tools ecosystem, I had lean to implement first. So, much likely a huge refactor it's gonna be needed in the after-match.
- I couldn't make docker-compose work in the ubuntu image, something is going on with Python, so I decided to install plain kafka infra.
- Certs and docker images are always fun!
