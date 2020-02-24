# Golang Command Query Responsibility Segregation Example
CQRS Example with GoLang, MQTT and GraphQL


# How it works

I will assume you're know what CQRS is (if not see [https://martinfowler.com/bliki/CQRS.html](https://martinfowler.com/bliki/CQRS.html))

This example uses a MQTT Queue for async write calls to the database. Here the concept of a `write call` is anything that does an CREATE, UPDATE or DELETE query in the database.

![How it works](./assets/CQRS.svg "CQRS")

We have three GraphQL Mutations (write endpoints) and two GraphQL Queries (read endpoints):

* Mutations
    * AddContact
    * UpdateContact
    * DeleteContact
* Queries
    * GetContact
    * ListContacts

## Database Service

The database service provides two gRPC services: One for writing and one for reading to allow segregation. This can also be separated in two different services. Here they're in a single service to be able to use a Memory Storage Database.

![Database Service](./assets/Database%20Service.svg "Database Service")

At any given moment, the `Write Service` shouldn't access *any* of the `Database Read Endpoints` and the `Read Service` shouldn't access *any* of the `Database Write Endpoints` 


![Database RW](./assets/Database%20RW.svg "Database RW")


## Database Writes (Mutations)

The mutation calls adds a gRPC Compatible arguments to a MQTT Queue which topics maps as following:

*   AddContact    => `CONTACT/ADD`
*   UpdateContact => `CONTACT/UPDATE`
*   DeleteContact => `CONTACT/DELETE`    

The `Write Service` will listen for all these topics to perform the database writes as needed. To simulate a "slow" environment, a 3 second delay is added before actually committing the changes to the database.

![Write Service](./assets/Write%20Service.svg "Write Service")


## Database Reads (Queries)

The query calls receives a gRPC client and passes through a `Reader Service`. The `Reader Service` here only proxies the calls to the `Database Service` but in a real scenario the `Reader Service` would contain business logic.

All the requests are sent synchronously to the database. That means it will wait for the database service to return with the data that was asked by the user. 

![Read Service](./assets/Read%20Service.svg "Read Service")


# How to run

All needed stuff to run is in this repository. There are few services:

*   `cmd/database`   => Database Service. Uses lungo as mongodb compatible memory storage
*   `cmd/graphql`    => GraphQL gRPC Gateway. The endpoint exposed to the world
*   `cmd/mqttserver` => A MQTT Server using [github.com/fhmq/hmq](https://github.com/fhmq/hmq/broker). This is used as Queue Server
*   `cmd/reader`     => Read Service. This should provide a gRPC way to read stuff from the database
*   `cmd/writer`     => Write Service. This should listen to MQTT Topics and write stuff in the database

All of the services should be run, and with the exception of `mqttserver` can be run like this:
```bash
cd cmd/database
go run .
```

For the MQTT Server:

```bash
cd cmd/mqttserver
go run . -c hmq.config
```

Not much time was spent handling the gRPC / MQTT Connection failures, so to avoid any issues start the services in this order:

1. MQTT Server (`cmd/mqttserver`)
2. Database Server (`cmd/database`)
3. Read Service (`cmd/reader`)
4. Write Service (`cmd/writer`)
5. GraphQL gRPC Gateway (`cmd/graphql`)

After everything is running, you can open graphql playground at [http://localhost:8080](http://localhost:8080)


## Example Queries:

```graphql
query ListContacts {
  listContacts(count: 10) {
    id
    name
    last_updated
  }
}

mutation AddContact {
  addContact(
    name: "John HUEBR"
  ) {
    status
    message
  }
}

mutation UpdateContact {
  updateContact(
    id: "0797213a-d0d0-4221-ad6b-0184672ed7cd"
    name: "John HUEBR (2)"
  ) {
    status
    message
  }
}

mutation DeleteContact {
  deleteContact(
    id: "0797213a-d0d0-4221-ad6b-0184672ed7cd"
  ) {
    status
    message
  }
}
```

# Other stuff

## Update Protobuf Models

```bash
protoc -I protocol/ protocol/mycqrs.proto --go_out=plugins=grpc:protocol
```