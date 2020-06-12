package connect

// Connect represents a connection to a server or database.
type Connect struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}
