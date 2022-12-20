package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Params struct {
	Interval int `json:"interval"`
	Duration int `json:"duration"`
}

type Config struct {
	AppName   string `json:"appName"`
	CfgParams Params `json:"params"`
}

var cfgData Config

func loadConfig() (context.Context, context.CancelFunc, time.Duration) {
	configPath := "./config.json"
	var interval, dur *int

	configFile, err := os.Open(configPath)

	if err == nil {
		log.Print("Successfully Opened config file:", configPath)

		defer configFile.Close()

		byteValue, _ := ioutil.ReadAll(configFile)

		err = json.Unmarshal(byteValue, &cfgData)
		interval = &cfgData.CfgParams.Interval
		dur = &cfgData.CfgParams.Duration
	} else {
		interval = flag.Int("interval", 20, "moving interval [sec]")
		dur = flag.Int("duration", -1, "moving duration [sec]")
		flag.Parse()

	}

	if *interval <= 0 {
		*interval = 1
	}

	itv := time.Second * time.Duration(*interval)

	log.Print("Config params are: interval = ", *interval, ", duration = ", *dur)

	if *dur <= 0 {
		ctx, cancel := context.WithCancel(context.Background())
		return ctx, cancel, itv
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*dur))
	return ctx, cancel, itv
}
