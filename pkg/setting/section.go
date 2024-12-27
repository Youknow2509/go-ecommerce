package setting

// Struct Config
type Config struct {
	Server ServerSetting `mapstructure:"server"`
	MySQL MySQLSetting `mapstructure:mysql"`
}

// Server Struct
type ServerSetting struct {
	Port int `mapstructure:"port"`
}

// My SQL Struct
type MySQLSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname string `mapstructure:"dbname"`
	MaxIdleConns int `mapstructure:"maxIdleConns"`
	MaxOpenConns int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`
}
