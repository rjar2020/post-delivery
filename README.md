# post-delivey system

## Prerequisites (besides Git, of course)
- Docker Desktop or any Docker compatible app/manager
- Go Tools

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
Once you are here, POST /postback endpoint can be used to produce messages to kafka (see postman/delivery-postback.postman_collection.json)

Use ***Control + c*** to stop the app

- For stopping all the dependencies plus deleting related containers 
```bash
./stop-and-delete-me
```
Once this is exectuted the state persisted in kafka is gone.


## Notes

### to-do
- PHP ingester to hit POST /postback endpoint
- Complete business logic
- Kafka config improvements
- Observability

### Log
- PHP and kafka integration existing projects/plugins usage is not straight forward and requires a lot of infrastructure/plugins, so I decided to create a Go endpoint to facilitate this integration.
- Also PHP tools are messy, but maybe this is just a personal opinion comming for someone trying Go tools (I really like this ecosystem) and relying on the JVM ecosystem most of the time (Yes, variety, not self-contained like Golang but with clear conventions to follow)
- Despite I really rely on TDD to orient my design and improve speed of development, when experimenting a new tools ecosystem, I have lean to implement first. So, much likely a huge refactor it's gonna be needed in the after-match.