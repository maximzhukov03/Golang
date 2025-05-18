package config

type Config struct {
    DBPath      string `env:"DB_PATH" default:"./data/messenger.db"`
    MinIOEndpoint string `env:"MINIO_ENDPOINT" default:"localhost:9000"`
    MinIOAccessKey string `env:"MINIO_ACCESS_KEY"`
    MinIOSecretKey string `env:"MINIO_SECRET_KEY"`
    JWTSecret     string `env:"JWT_SECRET"`
}

func Load() *Config {
    // Загрузка из .env или переменных окружения
    return &Config{
        JWTSecret: "your-secret-key",
    }
}