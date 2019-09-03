# semp-client

Go client for the [Solace](https://solace.com/) [SEMPv2 API](https://docs.solace.com/SEMP/Using-SEMP.htm). Generated using [go-swagger](https://goswagger.io) on the swagger definition:

```sh
go-swagger generate client -f specs/semp-v2_10-swagger-config-edited.yaml
```

The spec can be obtained from [the Solace website](https://solace.com/downloads/) with a command like
```
curl -L -v -o specs/semp-v2_12-swagger-config.yaml https://products.solace.com/download/PUBSUB_SEMPV2_SCHEMA_YAML
```
