package setting

// Struct Config
type Config struct {
	Server ServerSetting `mapstructure:"server"`
	MySQL  MySQLSetting  `mapstructure:mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
}

// Redis Struct
type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// Server Struct
type ServerSetting struct {
	Port int `mapstructure:"port"`
}

// My SQL Struct
type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

// Logger Struct
type LoggerSetting struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"file_log_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}
