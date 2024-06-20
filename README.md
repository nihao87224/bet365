# bet365
A gambling app with features such as recording gambling topics and automatically transferring stakes.

## description
It is made up of three components:
* Frontend App: The frontend is an H5 app that runs on the WeChat platform. It serves as the entry point for the entire system, allowing you to create gambling items, join existing ones, and review the rewards of the gambling activities you have participated in.
* Backend Service: The backend service provides various APIs to support the frontend functionalities, such as creating bet orders, settling bets, and other related operations.
* Blockchain layer: Records all bet details, including topics, participants, and results.

## directory structural
* bootstrap: Initialization files.
* config: Configuration files.
* database: SQL scripts.
* docs: Documentation include system architecture, features, and database design.
* routers: url routing.
* weChat: WeChat mini-program service, business logic.
* test: Test files.
* main.go: Main entry file.
* go.mod: Go module file.

## start and verify the project
1. Start the service: `go run main.go`
2. Verify the service: `curl http://127.0.0.1:20191/api/v1/home/news?page=2&limit=52&newsType=portal`