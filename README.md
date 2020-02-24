# go-cqrs-example
CQRS Example with GoLang, MQTT and GraphQL


# Generate Protobuf Models

```bash
protoc -I protocol/ protocol/mycqrs.proto --go_out=plugins=grpc:protocol
```