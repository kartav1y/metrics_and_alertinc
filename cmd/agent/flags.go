package main

import (
	"flag"
	"os"
	"strconv"
)

type FlagNames struct {
	FlagRunAddr    string
	PollInterval   int64
	ReportInterval int64
}

var flagsName FlagNames

func parseFlags() {
	flag.StringVar(&flagsName.FlagRunAddr, "a", "localhost:8080", "address and port to run server")
	flag.Int64Var(&flagsName.PollInterval, "p", 2, "poll interval")
	flag.Int64Var(&flagsName.ReportInterval, "r", 10, "report interval")

	flag.Parse()
	//fmt.Println("Addr: ", flagsName.FlagRunAddr)
	//fmt.Println("pol: ", flagsName.PollInterval)
	//fmt.Println("rep: ", flagsName.ReportInterval)

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		//fmt.Println("Addr 1 : ", envRunAddr)
		envRunAddr = flagsName.FlagRunAddr
		//fmt.Println("Addr 2 : ", flagsName.FlagRunAddr)
	}

	if interval := os.Getenv("REPORT_INTERVAL"); interval != " " {
		//fmt.Println("Rep 1", interval)
		if value, err := strconv.ParseInt(interval, 10, 64); err == nil {
			flagsName.ReportInterval = value
			//fmt.Println("Rep 2", flagsName.ReportInterval)
		}
	}

	if interval := os.Getenv("POLL_INTERVAL"); interval != "" {
		//fmt.Println("Rep 1", interval)
		if value, err := strconv.ParseInt(interval, 10, 64); err == nil {
			flagsName.PollInterval = value
			//fmt.Println("Rep 2", flagsName.PollInterval)
		}
	}

	//if !strings.HasPrefix(flagsName.FlagRunAddr, "http") && !strings.HasPrefix(flagsName.FlagRunAddr, "https") {
	//	flagsName.FlagRunAddr = "http://" + flagsName.FlagRunAddr
	//}
}