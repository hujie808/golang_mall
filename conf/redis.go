package conf

//redis

type RdsConfig struct {
	Host      string
	Port      int
	User      string
	Psw       string
	IsRunning bool
}

var RdsCacheList = []RdsConfig{
	{
		Host:      "127.0.0.1",
		Port:      6379,
		User:      "",
		Psw:       "",
		IsRunning: true,
	},
}

var RdsCache RdsConfig = RdsCacheList[0]
