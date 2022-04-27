## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| USERS_DEBUG_ADDR | string | 127.0.0.1:9145 | |
| USERS_DEBUG_TOKEN | string |  | |
| USERS_DEBUG_PPROF | bool | false | |
| USERS_DEBUG_ZPAGES | bool | false | |
| USERS_GRPC_ADDR | string | 127.0.0.1:9144 | The address of the grpc service.|
| USERS_GRPC_PROTOCOL | string | tcp | The transport protocol of the grpc service.|
| LDAP_URI;USERS_LDAP_URI | string | ldaps://localhost:9126 | |
| LDAP_CACERT;USERS_LDAP_CACERT | string | ~/.ocis/ldap/ldap.crt | |
| LDAP_INSECURE;USERS_LDAP_INSECURE | bool | false | |
| LDAP_BIND_DN;USERS_LDAP_BIND_DN | string | cn=reva,ou=sysusers,dc=ocis,dc=test | |
| LDAP_BIND_PASSWORD;USERS_LDAP_BIND_PASSWORD | string | reva | |
| LDAP_USER_BASE_DN;USERS_LDAP_USER_BASE_DN | string | dc=ocis,dc=test | |
| LDAP_GROUP_BASE_DN;USERS_LDAP_GROUP_BASE_DN | string | dc=ocis,dc=test | |
| LDAP_USERFILTER;USERS_LDAP_USERFILTER | string |  | |
| LDAP_GROUPFILTER;USERS_LDAP_USERFILTER | string |  | |
| LDAP_USER_OBJECTCLASS;USERS_LDAP_USER_OBJECTCLASS | string | posixAccount | |
| LDAP_GROUP_OBJECTCLASS;USERS_LDAP_GROUP_OBJECTCLASS | string | posixGroup | |
| LDAP_LOGIN_ATTRIBUTES;USERS_LDAP_LOGIN_ATTRIBUTES |  | [cn mail] | |
| OCIS_URL;USERS_IDP_URL | string | https://localhost:9200 | |
| LDAP_USER_SCHEMA_ID;USERS_LDAP_USER_SCHEMA_ID | string | ownclouduuid | |
| LDAP_USER_SCHEMA_ID_IS_OCTETSTRING;USERS_LDAP_USER_SCHEMA_ID_IS_OCTETSTRING | bool | false | |
| LDAP_USER_SCHEMA_MAIL;USERS_LDAP_USER_SCHEMA_MAIL | string | mail | |
| LDAP_USER_SCHEMA_DISPLAYNAME;USERS_LDAP_USER_SCHEMA_DISPLAYNAME | string | displayname | |
| LDAP_USER_SCHEMA_USERNAME;USERS_LDAP_USER_SCHEMA_USERNAME | string | cn | |
| LDAP_GROUP_SCHEMA_ID;USERS_LDAP_GROUP_SCHEMA_ID | string | cn | |
| LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING;USERS_LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING | bool | false | |
| LDAP_GROUP_SCHEMA_MAIL;USERS_LDAP_GROUP_SCHEMA_MAIL | string | mail | |
| LDAP_GROUP_SCHEMA_DISPLAYNAME;USERS_LDAP_GROUP_SCHEMA_DISPLAYNAME | string | cn | |
| LDAP_GROUP_SCHEMA_GROUPNAME;USERS_LDAP_GROUP_SCHEMA_GROUPNAME | string | cn | |
| LDAP_GROUP_SCHEMA_MEMBER;USERS_LDAP_GROUP_SCHEMA_MEMBER | string | cn | |