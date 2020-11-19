package jenkinsapi

import (
	"log"
	"os"
	"time"
)

type HTTPTimeout struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	HeaderTimeout    time.Duration
	LongTimeout      time.Duration
	IdleConnTimeout  time.Duration
}

type HTTPMaxConns struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
}

type Config struct {
	Endpoint      string
	User          string
	ApiToken      string
	RetryTimes    uint         // Retry count by default it's 5.
	UserAgent     string       // SDK name/version/system information
	IsDebug       bool         // Enable debug mode. Default is false.
	Timeout       uint         // Timeout in seconds. By default it's 60.
	IsCname       bool         // If cname is in the endpoint.
	HTTPTimeout   HTTPTimeout  // HTTP timeout
	HTTPMaxConns  HTTPMaxConns // Http max connections
	IsUseProxy    bool         // Flag of using proxy.
	ProxyHost     string       // Flag of using proxy host.
	IsAuthProxy   bool         // Flag of needing authentication.
	ProxyUser     string       // Proxy user
	ProxyPassword string       // Proxy password
	Logger        *log.Logger  // For write log
}

func getDefaultConf() *Config {
	return &Config{
		Endpoint:   "",
		User:       "",
		ApiToken:   "",
		RetryTimes: 5,
		UserAgent:  userAgent(),
		IsDebug:    false,
		Timeout:    60,
		IsCname:    false,
		HTTPTimeout: HTTPTimeout{
			ConnectTimeout:   time.Second * 30,
			ReadWriteTimeout: time.Second * 60,
			HeaderTimeout:    time.Second * 60,
			LongTimeout:      time.Second * 300,
			IdleConnTimeout:  time.Second * 60,
		},
		HTTPMaxConns: HTTPMaxConns{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
		},
		IsUseProxy:    false,
		ProxyHost:     "",
		IsAuthProxy:   false,
		ProxyUser:     "",
		ProxyPassword: "",
		Logger:        log.New(os.Stdout, "", log.LstdFlags),
	}
}
