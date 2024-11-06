package config

import (
	"flag"
	"os"
)

var options struct {
	flagRunAddr    string
	flagRunAccrual string
	databaseDsn    string
}

func ParseFlags() {
	flag.StringVar(&options.flagRunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&options.flagRunAccrual, "r", "http://localhost:9090", "address and port to run accrual system")
	flag.StringVar(&options.databaseDsn, "d", "host=localhost user=test password=password dbname=gophemart sslmode=disable", "pgsql data source name")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		options.flagRunAddr = envRunAddr
	}

	if envRunAccrual := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envRunAccrual != "" {
		options.flagRunAccrual = envRunAccrual
	}

	if databaseDsn := os.Getenv("DATABASE_URI"); databaseDsn != "" {
		options.databaseDsn = databaseDsn
	}
}

func GetHost() string {
	return options.flagRunAddr
}

func GetAccrualHost() string {
	return options.flagRunAccrual
}

func GetDsn() string {
	return options.databaseDsn
}

func GetJWTSecret() string {
	if jwtSecret := os.Getenv("DATABASE_URI"); jwtSecret != "" {
		return jwtSecret
	}

	return "1801f0c0cad004cc87de53eeb864b2589c605fcfcd42e13d564e7d93e21a4a6b"
}
