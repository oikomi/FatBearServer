package autoload

type Security struct {
	SuperToken  string `mapstructure:"super_token" json:"super_token" yaml:"super_token"`
	PasswordKey string `mapstructure:"password_key" json:"password_key" yaml:"password_key"`
}
