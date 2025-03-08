package config

var Config AppConfig

type AppConfig struct {
	Port                  int      `json:"port"`
	AppName               string   `json:"appName"`
	AppEnv                string   `json:"appEnv"`
	SignatureKey          string   `json:"signatureKey"`
	Database              Database `json:"database"`
	RateLimiterMaxRequest float64  `json:"rateLimiterRequest"`
	RateLimiterTimeSecond int      `json:"rateLimiterTimeSecond"`
	JwtSecretKey          string   `json:"jwtSecretKey"`
	JwtExpirationTime     int      `json:"jwtExpirationTime"`
}

type Database struct {
	Host                  string `json:"database"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnection     int    `json:"maxOpenConnection"`
	MaxLifetimeConnection int    `json:"maxLifetimeConnection"`
	MaxIdleConnection     int    `json:"maxIdleConnection"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

func Init() {
	// err :=
	// TODO ADD INIT CONFIG
}
