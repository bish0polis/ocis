package config

import (
	"context"

	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
)

type Config struct {
	Commons *shared.Commons `yaml:"-"` // don't use this directly as configuration for a service
	Service Service         `yaml:"-"`
	Tracing *Tracing        `yaml:"tracing"`
	Log     *Log            `yaml:"log"`
	Debug   Debug           `yaml:"debug"`

	GRPC GRPCConfig `yaml:"grpc"`

	TokenManager *TokenManager `yaml:"token_manager"`
	Reva         *Reva         `yaml:"reva"`
	Events       Events        `yaml:"events"`

	SkipUserGroupsInToken bool `yaml:"skip_user_groups_in_token" env:"SHARING_SKIP_USER_GROUPS_IN_TOKEN"`

	UserSharingDriver    string               `yaml:"user_sharing_driver" env:"SHARING_USER_DRIVER"`
	UserSharingDrivers   UserSharingDrivers   `yaml:"user_sharing_drivers"`
	PublicSharingDriver  string               `yaml:"public_sharing_driver" env:"SHARING_PUBLIC_DRIVER"`
	PublicSharingDrivers PublicSharingDrivers `yaml:"public_sharing_drivers"`

	Supervised bool            `yaml:"-"`
	Context    context.Context `yaml:"-"`
}
type Tracing struct {
	Enabled   bool   `yaml:"enabled" env:"OCIS_TRACING_ENABLED;SHARING_TRACING_ENABLED" desc:"Activates tracing."`
	Type      string `yaml:"type" env:"OCIS_TRACING_TYPE;SHARING_TRACING_TYPE" desc:"The type of tracing. Defaults to \"\", which is the same as \"jaeger\". Allowed tracing types are \"jaeger\" and \"\" as of now."`
	Endpoint  string `yaml:"endpoint" env:"OCIS_TRACING_ENDPOINT;SHARING_TRACING_ENDPOINT" desc:"The endpoint of the tracing agent."`
	Collector string `yaml:"collector" env:"OCIS_TRACING_COLLECTOR;SHARING_TRACING_COLLECTOR" desc:"The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces. Only used if the tracing endpoint is unset."`
}

type Log struct {
	Level  string `yaml:"level" env:"OCIS_LOG_LEVEL;SHARING_LOG_LEVEL" desc:"The log level. Valid values are: \"panic\", \"fatal\", \"error\", \"warn\", \"info\", \"debug\", \"trace\"."`
	Pretty bool   `yaml:"pretty" env:"OCIS_LOG_PRETTY;SHARING_LOG_PRETTY" desc:"Activates pretty log output."`
	Color  bool   `yaml:"color" env:"OCIS_LOG_COLOR;SHARING_LOG_COLOR" desc:"Activates colorized log output."`
	File   string `yaml:"file" env:"OCIS_LOG_FILE;SHARING_LOG_FILE" desc:"The path to the log file. Activates logging to this file if set."`
}

type Service struct {
	Name string `yaml:"-"`
}

type Debug struct {
	Addr   string `yaml:"addr" env:"SHARING_DEBUG_ADDR" desc:"Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed."`
	Token  string `yaml:"token" env:"SHARING_DEBUG_TOKEN" desc:"Token to secure the metrics endpoint"`
	Pprof  bool   `yaml:"pprof" env:"SHARING_DEBUG_PPROF" desc:"Enables pprof, which can be used for profiling"`
	Zpages bool   `yaml:"zpages" env:"SHARING_DEBUG_ZPAGES" desc:"Enables zpages, which can be used for collecting and viewing in-memory traces."`
}

type GRPCConfig struct {
	Addr      string `yaml:"addr" env:"SHARING_GRPC_ADDR" desc:"The address of the grpc service."`
	Namespace string `yaml:"-"`
	Protocol  string `yaml:"protocol" env:"SHARING_GRPC_PROTOCOL" desc:"The transport protocol of the grpc service."`
}

type UserSharingDrivers struct {
	JSON        UserSharingJSONDriver        `yaml:"json"`
	CS3         UserSharingCS3Driver         `yaml:"cs3"`
	OwnCloudSQL UserSharingOwnCloudSQLDriver `yaml:"owncloudsql"`

	SQL UserSharingSQLDriver `yaml:"sql,omitempty"` // not supported by the oCIS product, therefore not part of docs
}

type UserSharingJSONDriver struct {
	File string `yaml:"file" env:"SHARING_USER_JSON_FILE"`
}

type UserSharingSQLDriver struct {
	DBUsername                 string `yaml:"db_username"`
	DBPassword                 string `yaml:"db_password"`
	DBHost                     string `yaml:"db_host"`
	DBPort                     int    `yaml:"db_port"`
	DBName                     string `yaml:"db_name"`
	PasswordHashCost           int    `yaml:"password_hash_cost"`
	EnableExpiredSharesCleanup bool   `yaml:"enable_expired_shares_cleanup"`
	JanitorRunInterval         int    `yaml:"janitor_run_interval"`
	UserStorageMountID         string `yaml:"user_storage_mount_id"`
}

type UserSharingOwnCloudSQLDriver struct {
	DBUsername         string `yaml:"db_username" env:"SHARING_USER_OWNCLOUDSQL_DB_USERNAME"`
	DBPassword         string `yaml:"db_password" env:"SHARING_USER_OWNCLOUDSQL_DB_PASSWORD"`
	DBHost             string `yaml:"db_host" env:"SHARING_USER_OWNCLOUDSQL_DB_HOST"`
	DBPort             int    `yaml:"db_port" env:"SHARING_USER_OWNCLOUDSQL_DB_PORT"`
	DBName             string `yaml:"db_name" env:"SHARING_USER_OWNCLOUDSQL_DB_NAME"`
	UserStorageMountID string `yaml:"user_storage_mount_id" env:"SHARING_USER_OWNCLOUDSQL_USER_STORAGE_MOUNT_ID"`
}

type UserSharingCS3Driver struct {
	ProviderAddr     string `yaml:"provider_addr" env:"SHARING_USER_CS3_PROVIDER_ADDR"`
	SystemUserID     string `yaml:"system_user_id" env:"OCIS_SYSTEM_USER_ID;SHARING_USER_CS3_SYSTEM_USER_ID"`
	SystemUserIDP    string `yaml:"system_user_idp" env:"OCIS_SYSTEM_USER_IDP;SHARING_USER_CS3_SYSTEM_USER_IDP"`
	SystemUserAPIKey string `yaml:"system_user_api_key" env:"OCIS_SYSTEM_USER_API_KEY;SHARING_USER_CS3_SYSTEM_USER_API_KEY"`
}

type PublicSharingDrivers struct {
	JSON PublicSharingJSONDriver `yaml:"json"`
	CS3  PublicSharingCS3Driver  `yaml:"cs3"`

	SQL PublicSharingSQLDriver `yaml:"sql,omitempty"` // not supported by the oCIS product, therefore not part of docs
}

type PublicSharingJSONDriver struct {
	File string `yaml:"file" env:"SHARING_PUBLIC_JSON_FILE"`
}

type PublicSharingSQLDriver struct {
	DBUsername                 string `yaml:"db_username"`
	DBPassword                 string `yaml:"db_password"`
	DBHost                     string `yaml:"db_host"`
	DBPort                     int    `yaml:"db_port"`
	DBName                     string `yaml:"db_name"`
	PasswordHashCost           int    `yaml:"password_hash_cost"`
	EnableExpiredSharesCleanup bool   `yaml:"enable_expired_shares_cleanup"`
	JanitorRunInterval         int    `yaml:"janitor_run_interval"`
	UserStorageMountID         string `yaml:"user_storage_mount_id"`
}

type PublicSharingCS3Driver struct {
	ProviderAddr     string `yaml:"provider_addr" env:"SHARING_PUBLIC_CS3_PROVIDER_ADDR"`
	SystemUserID     string `yaml:"system_user_id" env:"OCIS_SYSTEM_USER_ID;SHARING_PUBLIC_CS3_SYSTEM_USER_ID"`
	SystemUserIDP    string `yaml:"system_user_idp" env:"OCIS_SYSTEM_USER_IDP;SHARING_PUBLIC_CS3_SYSTEM_USER_IDP"`
	SystemUserAPIKey string `yaml:"system_user_api_key" env:"OCIS_SYSTEM_USER_API_KEY;SHARING_USER_CS3_SYSTEM_USER_API_KEY"`
}

type Events struct {
	Addr      string `yaml:"endpoint" env:"SHARING_EVENTS_ENDPOINT" desc:"the address of the streaming service"`
	ClusterID string `yaml:"cluster" env:"SHARING_EVENTS_CLUSTER" desc:"the clusterID of the streaming service. Mandatory when using nats"`
}