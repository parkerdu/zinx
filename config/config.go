package config

import (
	"encoding/xml"
	"os"
	"zinx/zinterface"
)

type conf struct {
	/*
		server
	*/
	TcpServer zinterface.IServer

	Server ServerCfg
	/*
		zinx
	*/
	Version        string
	MaxConn        int
	MaxPackageSize int
}

type ServerCfg struct {
	XMLName xml.Name `xml:"server"`
	Name    string   `json:"name" xml:"name"`
	Host    string   `json:"host" xml:"host"`
	Port    int      `json:"port" xml:"port"`
}

var cfg *conf

func init() {
	cfg = &conf{
		TcpServer: nil,
		Server: ServerCfg{

			Name: "V0.4",
			Host: "0.0.0.0",
			Port: 8999,
		},
		Version:        "v0.4",
		MaxConn:        3,
		MaxPackageSize: 512,
	}

	//file, err := os.Open("config.xml")
	data, err := os.ReadFile("D:\\github\\zinx\\config.xml")
	if err != nil {
		panic(err)
	}
	if err := xml.Unmarshal(data, &cfg.Server); err != nil {
		panic(err)
	}
}

func Server() ServerCfg {
	return cfg.Server
}

func MaxPackageSize() int {
	return cfg.MaxPackageSize
}
