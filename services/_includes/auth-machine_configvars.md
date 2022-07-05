## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_TRACING_ENABLED<br/>AUTH_MACHINE_TRACING_ENABLED | bool | false | Activates tracing.|
| OCIS_TRACING_TYPE<br/>AUTH_MACHINE_TRACING_TYPE | string |  | The type of tracing. Defaults to "", which is the same as "jaeger". Allowed tracing types are "jaeger" and "" as of now.|
| OCIS_TRACING_ENDPOINT<br/>AUTH_MACHINE_TRACING_ENDPOINT | string |  | The endpoint of the tracing agent.|
| OCIS_TRACING_COLLECTOR<br/>AUTH_MACHINE_TRACING_COLLECTOR | string |  | The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces. Only used if the tracing endpoint is unset.|
| OCIS_LOG_LEVEL<br/>AUTH_MACHINE_LOG_LEVEL | string |  | The log level. Valid values are: "panic", "fatal", "error", "warn", "info", "debug", "trace".|
| OCIS_LOG_PRETTY<br/>AUTH_MACHINE_LOG_PRETTY | bool | false | Activates pretty log output.|
| OCIS_LOG_COLOR<br/>AUTH_MACHINE_LOG_COLOR | bool | false | Activates colorized log output.|
| OCIS_LOG_FILE<br/>AUTH_MACHINE_LOG_FILE | string |  | The path to the log file. Activates logging to this file if set.|
| AUTH_MACHINE_DEBUG_ADDR | string | 127.0.0.1:9167 | Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed.|
| AUTH_MACHINE_DEBUG_TOKEN | string |  | Token to secure the metrics endpoint|
| AUTH_MACHINE_DEBUG_PPROF | bool | false | Enables pprof, which can be used for profiling|
| AUTH_MACHINE_DEBUG_ZPAGES | bool | false | Enables zpages, which can be used for collecting and viewing in-memory traces.|
| AUTH_MACHINE_GRPC_ADDR | string | 127.0.0.1:9166 | The bind address of the GRPC service.|
| AUTH_MACHINE_GRPC_PROTOCOL | string | tcp | The transport protocol of the grpc service.|
| OCIS_JWT_SECRET<br/>AUTH_MACHINE_JWT_SECRET | string |  | The secret to mint and validate jwt tokens.|
| REVA_GATEWAY | string | 127.0.0.1:9142 | The CS3 gateway endpoint.|
| AUTH_MACHINE_SKIP_USER_GROUPS_IN_TOKEN | bool | false | Disables the encoding of the user's group memberships in the reva access token. This reduces the token size, especially when users are members of a large number of groups.|
| OCIS_MACHINE_AUTH_API_KEY<br/>AUTH_MACHINE_API_KEY | string |  | Machine auth API key used for validating requests from other services when impersonating users.|