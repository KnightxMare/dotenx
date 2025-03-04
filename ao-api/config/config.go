package config

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		App       App
		Database  Database
		Secrets   Secrets
		Endpoints Endpoints
		Queue     Queue
		Redis     Redis
	}

	App struct {
		Port            string `envconfig:"AOA_APP_PORT" default:"3004"`
		CheckTrigger    string `envconfig:"AOA_APP_CHECK_TRIGGER"`
		AccountId       string `envconfig:"AOA_APP_ACCOUNT_ID"`
		Environment     string `envconfig:"AOA_APP_ENV"`
		FileSharing     string `envconfig:"AOA_FILE_DIR"`
		SessionDuration string `envconfig:"AOA_SESSION_DURATION"`
		AllowedOrigins  string `envconfig:"AOA_APP_ALLOWED_ORIGINS" default:"*"`
		RunLocally      bool   `envconfig:"AOA_APP_RUN_LOCALLY" default:"true"`
	}

	Queue struct {
		BULL string `envconfig:"AOA_BULL_URL"`
	}

	Endpoints struct {
		AoApi      string `envconfig:"AOA_AO_API_URL"`
		AoApiLocal string `envconfig:"AOA_AO_API_LOCAL_URL"`
		Core       string `envconfig:"AOA_CORE_API_URL"`
		UI         string `envconfig:"AOA_UI_URL"`
		UILocal    string `envconfig:"AOA_UI_LOCAL_URL"`
		Admin      string `envconfig:"AOA_ADMIN_URL"`
	}

	Database struct {
		Host     string `envconfig:"AOA_DATABASE_HOST"`
		Port     int    `envconfig:"AOA_DATABASE_PORT"`
		User     string `envconfig:"AOA_DATABASE_USER"`
		Password string `envconfig:"AOA_DATABASE_PASSWORD"`
		DbName   string `envconfig:"AOA_DATABASE_DBNAME"`
		Extras   string `envconfig:"AOA_DATABASE_EXTRAS"`
		Driver   string `envconfig:"AOA_DATABASE_DRIVER" default:"postgres"`
	}

	Redis struct {
		Host string `envconfig:"AOA_REDIS_HOST"`
		Port int    `envconfig:"AOA_REDIS_PORT"`
	}

	Secrets struct {
		AuthServerJwtSecret  string `envconfig:"AOA_AUTH_SERVER_JWT_SECRET"`
		AppName              string `envconfig:"AOA_APP_NAME"`
		AppSecret            string `envconfig:"AOA_APP_SECRET"`
		CookieSecret         string `envconfig:"AOA_Cookie_SECRET"`
		Encryption           string `envconfig:"AOA_ENCRYPTION_SECRET"`
		SessionAuthSecret    string `envconfig:"AOA_SESSION_AUTH_SECRET"`
		SessionEncryptSecret string `envconfig:"AOA_SESSION_ENCRYPT_SECRET"`
		RunnerToken          string `envconfig:"AOA_RUNNER_TOKEN_SECRET"`
		CodeChallenge        string `envconfig:"AOA_CODE_CHALLENGE"`
	}
)

var Configs Config

func Load() error {
	err := envconfig.Process("", &Configs)
	return err
}
