# post-delivey system

## Prerequisites (besides Git, of course)
- Docker Desktop or any Docker compatible app/manager
- Go Tools

## Running the app

### Local / Dev env
To start post delivery system, ***from the project root*** use:

- For spinning up all the dependencies (kafka platform, etc) and start the app
```bash
./load-and-execute-dependencies
```
You can access confluentinc/cp-enterprise-control-center via http://localhost:9021/ to visualize kafka stats
Including: topics, messages, consumer groups, etc.

- For stopping the app and its dependencies plus deleting related containers 
```bash
./stop-and-delete-me
```
Once this is exectuted the state persisted in kafka is gone.

- For starting the app
```bash
go run github.com/rjar2020/post-delivery
```
Use ***Control + c*** to stop the app