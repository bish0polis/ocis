## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| AUTH_BASIC_DEBUG_ADDR | string | 127.0.0.1:9147 | |
| AUTH_BASIC_DEBUG_TOKEN | string |  | |
| AUTH_BASIC_DEBUG_PPROF | bool | false | |
| AUTH_BASIC_DEBUG_ZPAGES | bool | false | |
| AUTH_BASIC_GRPC_ADDR | string | 127.0.0.1:9146 | The address of the grpc service.|
| AUTH_BASIC_GRPC_PROTOCOL | string | tcp | The transport protocol of the grpc service.|
| AUTH_BASIC_AUTH_PROVIDER | string | ldap | The auth provider which should be used by the service|
| AUTH_BASIC_JSON_PROVIDER_FILE | string |  | The file to which the json provider writes the data.|
| LDAP_URI;AUTH_BASIC_LDAP_URI | string | ldaps://localhost:9126 | |
| LDAP_CACERT;AUTH_BASIC_LDAP_CACERT | string | ~/.ocis/ldap/ldap.crt | |
| LDAP_INSECURE;AUTH_BASIC_LDAP_INSECURE | bool | false | |
| LDAP_BIND_DN;AUTH_BASIC_LDAP_BIND_DN | string | cn=reva,ou=sysusers,dc=ocis,dc=test | |
| LDAP_BIND_PASSWORD;AUTH_BASIC_LDAP_BIND_PASSWORD | string | reva | |
| LDAP_USER_BASE_DN;AUTH_BASIC_LDAP_USER_BASE_DN | string | dc=ocis,dc=test | |
| LDAP_GROUP_BASE_DN;AUTH_BASIC_LDAP_GROUP_BASE_DN | string | dc=ocis,dc=test | |
| LDAP_USERFILTER;AUTH_BASIC_LDAP_USERFILTER | string |  | |
| LDAP_GROUPFILTER;AUTH_BASIC_LDAP_USERFILTER | string |  | |
| LDAP_USER_OBJECTCLASS;AUTH_BASIC_LDAP_USER_OBJECTCLASS | string | posixAccount | |
| LDAP_GROUP_OBJECTCLASS;AUTH_BASIC_LDAP_GROUP_OBJECTCLASS | string | posixGroup | |
| LDAP_LOGIN_ATTRIBUTES;AUTH_BASIC_LDAP_LOGIN_ATTRIBUTES |  | [cn mail] | |
| OCIS_URL;AUTH_BASIC_IDP_URL | string | https://localhost:9200 | |
| LDAP_USER_SCHEMA_ID;AUTH_BASIC_LDAP_USER_SCHEMA_ID | string | ownclouduuid | |
| LDAP_USER_SCHEMA_ID_IS_OCTETSTRING;AUTH_BASIC_LDAP_USER_SCHEMA_ID_IS_OCTETSTRING | bool | false | |
| LDAP_USER_SCHEMA_MAIL;AUTH_BASIC_LDAP_USER_SCHEMA_MAIL | string | mail | |
| LDAP_USER_SCHEMA_DISPLAYNAME;AUTH_BASIC_LDAP_USER_SCHEMA_DISPLAYNAME | string | displayname | |
| LDAP_USER_SCHEMA_USERNAME;AUTH_BASIC_LDAP_USER_SCHEMA_USERNAME | string | cn | |
| LDAP_GROUP_SCHEMA_ID;AUTH_BASIC_LDAP_GROUP_SCHEMA_ID | string | cn | |
| LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING;AUTH_BASIC_LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING | bool | false | |
| LDAP_GROUP_SCHEMA_MAIL;AUTH_BASIC_LDAP_GROUP_SCHEMA_MAIL | string | mail | |
| LDAP_GROUP_SCHEMA_DISPLAYNAME;AUTH_BASIC_LDAP_GROUP_SCHEMA_DISPLAYNAME | string | cn | |
| LDAP_GROUP_SCHEMA_GROUPNAME;AUTH_BASIC_LDAP_GROUP_SCHEMA_GROUPNAME | string | cn | |
| LDAP_GROUP_SCHEMA_MEMBER;AUTH_BASIC_LDAP_GROUP_SCHEMA_MEMBER | string | cn | |