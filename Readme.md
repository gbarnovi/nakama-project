# Game logic implementation using Nakama

This project involves the implementation of game logic using Nakama, an open-source game engine for the backend.

The task at hand is to create an RPC (Remote Procedure Call) function that accepts a payload with optional parameters, including type, version, and hash. The default values for these parameters are type=core, version=1.0.0, and hash=null.

Overall, this project aims to facilitate game logic implementation using Nakama by providing an RPC function that handles file reading, database storage, and hash calculation based on the provided payload.

# Project overview

I chose to utilize the Nakama Storage engine for data storage due to its simplicity and ease of use. Instead of using hashed values as key names, I decided to adopt a `<type>/<version>` naming convention for the keys. This choice allows for easier identification and management of data within the Nakama dashboard.

To optimize efficiency and minimize costly write operations, the system only writes data to the storage if the specified key name does not already exist for the user. This approach helps avoid unnecessary expenses and ensures that data is only written when necessary.

To ensure that the RPC functionality is working correctly, I have made two JSON files. These files act as examples that can be used to check if the RPC functions are functioning properly.

```
core/1.0.0.json
extension/3.0.0.json
```
# Folder structure

``` tree 
├── Dockerfile
├── Readme.md
├── data
│   ├── core
│   │   └── 1.0.0.json
│   └── extension
│       └── 3.0.0.json
├── docker-compose.yml
├── errors
│   └── error.go
├── go.mod
├── go.sum
├── local.yml
├── main.go
├── mockgen.go
├── mocks
│   ├── nakama_module.go
│   └── runtime_logger.go
├── rpc.go
├── rpc_models.go
├── rpc_test.go
├── storage.go
└── vendor
```

# Running the project
To run the project, navigate to the project's root directory and execute the following command:
```
docker-compose up
```
Visit Nakama dashboard at: http://127.0.0.1:7351

`Username: admin`

`Password: password`

Once logged in, go to the API Explorer page and select the `custom_rpc_function` from the dropdown menu.

Enter `00000000-0000-0000-0000-000000000000` as the user ID (this is the default user ID provided). This will allow you to interact with the desired custom RPC function.

# Tests

To simulate Nakama interfaces for testing purposes, I used [mockery](https://github.com/vektra/mockery). It allowed me to create mock versions of Nakama interfaces, making it easier to test and verify their functionality.

To run the tests, simply execute the following command from the project's root directory:

```
go test
```

# Things to improve

* Test coverage
* Input validation

Additionally, if the task were to have more specific requirements regarding the user ID, I would improve the logging by adding the userID to the log fields. This would make it easier to identify any potential issues.