package conf

//mysql配置

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Psw       string
	Database  string
	IsRunning bool //服务器状态
}

var DbMasterList = []DbConfig{
	{
		Host:      "39.96.160.74",
		Port:      3306,
		User:      "ccl123bat23E",
		Psw:       "cclbat32123EQ*",
		Database:  "golang_mall",
		IsRunning: false,
	},
}

var DbMaster DbConfig = DbMasterList[0]
