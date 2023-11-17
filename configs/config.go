package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Issuer       string `mapstructure:"issuer"`
	PostgresUrl  string `mapstructure:"postgres_url"`
	SymmetricKey string `mapstructure:"symmetric_key"`

	ManticoreDB     Database `mapstructure:"manticore_db"`
	ScyllaDB        Database `mapstructure:"scylla_db"`
	ElasticSearchDB Database `mapstructure:"elastic_search_db"`

	UserServiceEndpoint     Endpoint `mapstructure:"user_service_endpoint"`
	AnalyticServiceEndpoint Endpoint `mapstructure:"analytic_service_endpoint"`
	SearchServiceEndpoint   Endpoint `mapstructure:"search_service_endpoint"`
	GatewayServiceEndpoint  Endpoint `mapstructure:"gateway_service_endpoint"`
	StorageServiceEndpoint  Endpoint `mapstructure:"storage_service_endpoint"`
	GeneralServiceEndpoint  Endpoint `mapstructure:"general_service_endpoint"`
	MovieServiceEndpoint    Endpoint `mapstructure:"movie_service_endpoint"`
	TvSeriesServiceEndpoint Endpoint `mapstructure:"tv_series_service_endpoint"`
	DirectorServiceEndpoint Endpoint `mapstructure:"director_service_endpoint"`
	PartnerServiceEndpoint  Endpoint `mapstructure:"partner_service_endpoint"`

	TmDBEndpoint Endpoint `mapstructure:"tmdb_endpoint"`

	AdminAccount Account `mapstructure:"admin"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("env")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read config file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to read config file: %w", err)
	}

	return config, nil
}

type Account struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Email    string `mapstructure:"email"`
}
