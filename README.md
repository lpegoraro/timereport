# Timereport

## GoLang HTTP multi container service  POC
- 1 Container with the HTTP API and validations
- 1 Container with the Storage functionality
- Communication with protoc and Channels

## Use Docker to compile the protoc

```
docker pull namely/protoc-all
docker run -v `pwd`:/defs namely/protoc-all -f reportapi.proto -l go
```

