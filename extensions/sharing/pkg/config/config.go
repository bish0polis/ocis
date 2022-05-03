package config

import "github.com/owncloud/ocis/ocis-pkg/shared"

type Config struct {
	*shared.Commons `yaml:"-"`
	Service         Service  `yaml:"-"`
	Tracing         *Tracing `yaml:"tracing"`
	Logging         *Logging `yaml:"log"`
	Debug           Debug    `yaml:"debug"`
	Supervised      bool     `yaml:"-"`

	GRPC GRPCConfig `yaml:"grpc"`

	TokenManager *TokenManager `yaml:"token_manager"`
	Reva         *Reva         `yaml:"reva"`

	SkipUserGroupsInToken bool                 `yaml:"skip_user_groups_in_token"`
	UserSharingDriver     string               `yaml:"user_sharing_driver"`
	UserSharingDrivers    UserSharingDrivers   `yaml:"user_sharin_drivers"`
	PublicSharingDriver   string               `yaml:"public_sharing_driver"`
	PublicSharingDrivers  PublicSharingDrivers `yaml:"public_sharing_drivers"`
	Events                Events               `yaml:"events"`
}
type Tracing struct {
	Enabled   bool   `yaml:"enabled" env:"OCIS_TRACING_ENABLED;SHARING_TRACING_ENABLED" desc:"Activates tracing."`
	Type      string `yaml:"type" env:"OCIS_TRACING_TYPE;SHARING_TRACING_TYPE"`
	Endpoint  string `yaml:"endpoint" env:"OCIS_TRACING_ENDPOINT;SHARING_TRACING_ENDPOINT" desc:"The endpoint to the tracing collector."`
	Collector string `yaml:"collector" env:"OCIS_TRACING_COLLECTOR;SHARING_TRACING_COLLECTOR"`
}

type Logging struct {
	Level  string `yaml:"level" env:"OCIS_LOG_LEVEL;SHARING_LOG_LEVEL" desc:"The log level."`
	Pretty bool   `yaml:"pretty" env:"OCIS_LOG_PRETTY;SHARING_LOG_PRETTY" desc:"Activates pretty log output."`
	Color  bool   `yaml:"color" env:"OCIS_LOG_COLOR;SHARING_LOG_COLOR" desc:"Activates colorized log output."`
	File   string `yaml:"file" env:"OCIS_LOG_FILE;SHARING_LOG_FILE" desc:"The target log file."`
}

type Service struct {
	Name string `yaml:"-"`
}

type Debug struct {
	Addr   string `yaml:"addr" env:"SHARING_DEBUG_ADDR"`
	Token  string `yaml:"token" env:"SHARING_DEBUG_TOKEN"`
	Pprof  bool   `yaml:"pprof" env:"SHARING_DEBUG_PPROF"`
	Zpages bool   `yaml:"zpages" env:"SHARING_DEBUG_ZPAGES"`
}

type GRPCConfig struct {
	Addr     string `yaml:"addr" env:"SHARING_GRPC_ADDR" desc:"The address of the grpc service."`
	Protocol string `yaml:"protocol" env:"SHARING_GRPC_PROTOCOL" desc:"The transport protocol of the grpc service."`
}

type UserSharingDrivers struct {
	JSON UserSharingJSONDriver
	SQL  UserSharingSQLDriver
	CS3  UserSharingCS3Driver
}

type UserSharingJSONDriver struct {
	File string `env:"SHARING_USER_JSON_FILE"`
}

type UserSharingSQLDriver struct {
	DBUsername                 string `env:"SHARING_USER_SQL_USERNAME"`
	DBPassword                 string `env:"SHARING_USER_SQL_PASSWORD"`
	DBHost                     string `env:"SHARING_USER_SQL_HOST"`
	DBPort                     int    `env:"SHARING_USER_SQL_PORT"`
	DBName                     string `env:"SHARING_USER_SQL_NAME"`
	PasswordHashCost           int
	EnableExpiredSharesCleanup bool
	JanitorRunInterval         int
	UserStorageMountID         string
}

type UserSharingCS3Driver struct {
	ProviderAddr     string
	SystemUserID     string `env:"SYSTEM_USER_ID;SHARING_USER_SYSTEM_USER_ID"`
	SystemUserIDP    string `env:"SYSTEM_USER_IDP;SHARING_USER_SYSTEM_USER_IDP"`
	SystemAuthAPIKey string `env:"SYSTEM_AUTH_API_KEY;SHARING_USER_SYSTEM_AUTH_API_KEY"`
}

type PublicSharingDrivers struct {
	JSON PublicSharingJSONDriver
	SQL  PublicSharingSQLDriver
	CS3  PublicSharingCS3Driver
}

type PublicSharingJSONDriver struct {
	File string
}

type PublicSharingSQLDriver struct {
	DBUsername                 string
	DBPassword                 string
	DBHost                     string
	DBPort                     int
	DBName                     string
	PasswordHashCost           int
	EnableExpiredSharesCleanup bool
	JanitorRunInterval         int
	UserStorageMountID         string
}

type PublicSharingCS3Driver struct {
	ProviderAddr     string
	SystemUserID     string `env:"SYSTEM_USER_ID;SHARING_PUBLIC_SYSTEM_USER_ID"`
	SystemUserIDP    string `env:"SYSTEM_USER_IDP;SHARING_PUBLIC_SYSTEM_USER_IDP"`
	SystemAuthAPIKey string `env:"SYSTEM_AUTH_API_KEY;SHARING_PUBLIC_SYSTEM_AUTH_API_KEY"`
}

type Events struct {
	Addr      string
	ClusterID string
}
