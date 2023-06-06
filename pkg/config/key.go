package config

type Key string

const (
	AppPort                       Key = "app.port"
	AppTimezone                   Key = "app.timezone"
	AppServiceName                Key = "app.serviceName"
	AppDevMode                    Key = "app.devMode"
	PostgresHost                  Key = "database.postgres.host"
	PostgresUser                  Key = "database.postgres.user"
	PostgresPassword              Key = "database.postgres.password"
	PostgresDBName                Key = "database.postgres.dbName"
	PostgresPort                  Key = "database.postgres.port"
	PostgresSSLMode               Key = "database.postgres.sslMode"
	PostgresConnectionMaxLifetime Key = "database.postgres.connectionMaxLifetime"
	PostgresConnectionMaxOpen     Key = "database.postgres.connectionMaxOpen"
	PostgresConnectionMaxIdle     Key = "database.postgres.connectionMaxIdle"
	PostgresLogLevel              Key = "database.postgres.logLevel"
)
