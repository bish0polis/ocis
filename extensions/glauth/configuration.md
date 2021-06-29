---
title: "Configuration"
date: "2021-06-29T07:12:51+0000"
weight: 20
geekdocRepo: https://github.com/owncloud/ocis
geekdocEditPath: edit/master/glauth/templates
geekdocFilePath: CONFIGURATION.tmpl
---

{{< toc >}}

## Configuration

### Configuration using config files

Out of the box extensions will attempt to read configuration details from:

```console
/etc/ocis
$HOME/.ocis
./config
```

For this configuration to be picked up, have a look at your extension `root` command and look for which default config name it has assigned. *i.e: ocis-glauth reads `glauth.json | yaml | toml ...`*.

So far we support the file formats `JSON` and `YAML`, if you want to get a full example configuration just take a look at [our repository](https://github.com/owncloud/ocis/tree/master/glauth/config), there you can always see the latest configuration format. These example configurations include all available options and the default values. The configuration file will be automatically loaded if it's placed at `/etc/ocis/ocis.yml`, `${HOME}/.ocis/ocis.yml` or `$(pwd)/config/ocis.yml`.

### Environment variables

If you prefer to configure the service with environment variables you can see the available variables below.

If multiple variables are listed for one option, they are in order of precedence. This means the leftmost variable will always win if given.

### Commandline flags

If you prefer to configure the service with commandline flags you can see the available variables below. Command line flags are only working when calling the subcommand directly.

### glauth health

Check health status

Usage: `glauth health [command options] [arguments...]`





-debug-addr |  $GLAUTH_DEBUG_ADDR
: Address to debug endpoint. Default: `flags.OverrideDefaultString(cfg.Debug.Addr, "0.0.0.0:9129")`.

































### glauth ocis-glauth

Serve GLAuth API for oCIS

Usage: `glauth ocis-glauth [command options] [arguments...]`


-log-level |  $GLAUTH_LOG_LEVEL , $OCIS_LOG_LEVEL
: Set logging level.


-log-pretty |  $GLAUTH_LOG_PRETTY , $OCIS_LOG_PRETTY
: Enable pretty logging.


-log-color |  $GLAUTH_LOG_COLOR , $OCIS_LOG_COLOR
: Enable colored logging.


































### glauth server

Start integrated server

Usage: `glauth server [command options] [arguments...]`






-log-file |  $GLAUTH_LOG_FILE , $OCIS_LOG_FILE
: Enable log to file.


-config-file |  $GLAUTH_CONFIG_FILE
: Path to config file. Default: `flags.OverrideDefaultString(cfg.File, "")`.


-tracing-enabled |  $GLAUTH_TRACING_ENABLED
: Enable sending traces.


-tracing-type |  $GLAUTH_TRACING_TYPE
: Tracing backend type. Default: `flags.OverrideDefaultString(cfg.Tracing.Type, "jaeger")`.


-tracing-endpoint |  $GLAUTH_TRACING_ENDPOINT
: Endpoint for the agent. Default: `flags.OverrideDefaultString(cfg.Tracing.Endpoint, "")`.


-tracing-collector |  $GLAUTH_TRACING_COLLECTOR
: Endpoint for the collector. Default: `flags.OverrideDefaultString(cfg.Tracing.Collector, "")`.


-tracing-service |  $GLAUTH_TRACING_SERVICE
: Service name for tracing. Default: `flags.OverrideDefaultString(cfg.Tracing.Service, "glauth")`.


-debug-addr |  $GLAUTH_DEBUG_ADDR
: Address to bind debug server. Default: `flags.OverrideDefaultString(cfg.Debug.Addr, "0.0.0.0:9129")`.


-debug-token |  $GLAUTH_DEBUG_TOKEN
: Token to grant metrics access. Default: `flags.OverrideDefaultString(cfg.Debug.Token, "")`.


-debug-pprof |  $GLAUTH_DEBUG_PPROF
: Enable pprof debugging.


-debug-zpages |  $GLAUTH_DEBUG_ZPAGES
: Enable zpages debugging.


-role-bundle-id |  $GLAUTH_ROLE_BUNDLE_ID
: roleid used to make internal grpc requests. Default: `flags.OverrideDefaultString(cfg.RoleBundleUUID, "71881883-1768-46bd-a24d-a356a2afdf7f")`.


-ldap-addr |  $GLAUTH_LDAP_ADDR
: Address to bind ldap server. Default: `flags.OverrideDefaultString(cfg.Ldap.Address, "0.0.0.0:9125")`.


-ldap-enabled |  $GLAUTH_LDAP_ENABLED
: Enable ldap server. Default: `flags.OverrideDefaultBool(cfg.Ldap.Enabled, true)`.


-ldaps-addr |  $GLAUTH_LDAPS_ADDR
: Address to bind ldap server. Default: `flags.OverrideDefaultString(cfg.Ldaps.Address, "0.0.0.0:9126")`.


-ldaps-enabled |  $GLAUTH_LDAPS_ENABLED
: Enable ldap server. Default: `flags.OverrideDefaultBool(cfg.Ldaps.Enabled, true)`.


-ldaps-cert |  $GLAUTH_LDAPS_CERT
: path to ldaps certificate in PEM format. Default: `flags.OverrideDefaultString(cfg.Ldaps.Cert, path.Join(pkgos.MustUserConfigDir("ocis", "ldap"), "ldap.crt"))`.


-ldaps-key |  $GLAUTH_LDAPS_KEY
: path to ldaps key in PEM format. Default: `flags.OverrideDefaultString(cfg.Ldaps.Key, path.Join(pkgos.MustUserConfigDir("ocis", "ldap"), "ldap.key"))`.


-backend-basedn |  $GLAUTH_BACKEND_BASEDN
: base distinguished name to expose. Default: `flags.OverrideDefaultString(cfg.Backend.BaseDN, "dc=example,dc=org")`.


-backend-name-format |  $GLAUTH_BACKEND_NAME_FORMAT
: name attribute for entries to expose. typically cn or uid. Default: `flags.OverrideDefaultString(cfg.Backend.NameFormat, "cn")`.


-backend-group-format |  $GLAUTH_BACKEND_GROUP_FORMAT
: name attribute for entries to expose. typically ou, cn or dc. Default: `flags.OverrideDefaultString(cfg.Backend.GroupFormat, "ou")`.


-backend-ssh-key-attr |  $GLAUTH_BACKEND_SSH_KEY_ATTR
: ssh key attribute for entries to expose. Default: `flags.OverrideDefaultString(cfg.Backend.SSHKeyAttr, "sshPublicKey")`.


-backend-datastore |  $GLAUTH_BACKEND_DATASTORE
: datastore to use as the backend. one of accounts, ldap or owncloud. Default: `flags.OverrideDefaultString(cfg.Backend.Datastore, "accounts")`.


-backend-insecure |  $GLAUTH_BACKEND_INSECURE
: Allow insecure requests to the datastore. Default: `flags.OverrideDefaultBool(cfg.Backend.Insecure, false)`.


-backend-use-graphapi |  $GLAUTH_BACKEND_USE_GRAPHAPI
: use Graph API, only for owncloud datastore. Default: `flags.OverrideDefaultBool(cfg.Backend.UseGraphAPI, true)`.


-fallback-basedn |  $GLAUTH_FALLBACK_BASEDN
: base distinguished name to expose. Default: `flags.OverrideDefaultString(cfg.Fallback.BaseDN, "dc=example,dc=org")`.


-fallback-name-format |  $GLAUTH_FALLBACK_NAME_FORMAT
: name attribute for entries to expose. typically cn or uid. Default: `flags.OverrideDefaultString(cfg.Fallback.NameFormat, "cn")`.


-fallback-group-format |  $GLAUTH_FALLBACK_GROUP_FORMAT
: name attribute for entries to expose. typically ou, cn or dc. Default: `flags.OverrideDefaultString(cfg.Fallback.GroupFormat, "ou")`.


-fallback-ssh-key-attr |  $GLAUTH_FALLBACK_SSH_KEY_ATTR
: ssh key attribute for entries to expose. Default: `flags.OverrideDefaultString(cfg.Fallback.SSHKeyAttr, "sshPublicKey")`.


-fallback-datastore |  $GLAUTH_FALLBACK_DATASTORE
: datastore to use as the fallback. one of accounts, ldap or owncloud. Default: `flags.OverrideDefaultString(cfg.Fallback.Datastore, "")`.


-fallback-insecure |  $GLAUTH_FALLBACK_INSECURE
: Allow insecure requests to the datastore. Default: `flags.OverrideDefaultBool(cfg.Fallback.Insecure, false)`.


-fallback-use-graphapi |  $GLAUTH_FALLBACK_USE_GRAPHAPI
: use Graph API, only for owncloud datastore. Default: `flags.OverrideDefaultBool(cfg.Fallback.UseGraphAPI, true)`.
