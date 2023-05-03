package creation_type_mode

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost   string
	DBPort   int
	DBName   string
	LogLevel string
}

func LoadConfig() (*Config, error) {
	//从环境变量中读取配置
	loglevel := os.Getenv("LOG_LEVEL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// 解析端口号
	var port int
	if dbPort != "" {
		_, err := fmt.Sscanf(dbPort, "%d", &port)
		if err != nil {
			return nil, err
		}
	}
	return &Config{
		DBHost:   dbHost,
		DBPort:   port,
		DBName:   dbName,
		LogLevel: loglevel,
	}, nil
}

func CallConfigurePatternLowLevel() {
	// 加载配置
	config, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading config: ", err)
		return
	}
	// 输出配置信息
	fmt.Println("DBHost", config.DBHost)
	fmt.Println("DBPort", config.DBPort)
	fmt.Println("DBName", config.DBName)
	fmt.Println("DBLogLevel", config.LogLevel)

}
