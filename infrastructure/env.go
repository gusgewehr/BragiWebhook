package infrastructure

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
)

type Env struct {
	AmqpHost       string `mapstructure:"AMQP_HOST"`
	AmqpPort       int    `mapstructure:"AMQP_PORT"`
	AmqpUser       string `mapstructure:"AMQP_USER"`
	AmqpPassword   string `mapstructure:"AMQP_PASSWORD"`
	QueueName      string `mapstructure:"AMQP_QUEUE"`
	SubscribeToken string `mapstructure:"SUBSCRIBE_TOKEN"`
}

func NewEnv(envPath string, logger *zap.Logger) *Env {
	logger.Info("Iniciando leitura de variáveis de ambiente.")
	defer logger.Info("Leitura das variáveis de ambiente completada com sucesso.")

	env := Env{}

	v := viper.New()
	v.SetConfigFile(envPath)

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		logger.Fatal("Erro ao ler as configuracões de variáveis de ambiente.", zap.Error(err))
	}

	err = v.Unmarshal(&env)
	if err != nil {
		logger.Fatal("Erro ao fazer unmarshal das variáveis de ambiente.", zap.Error(err))
	}

	return &env
}
