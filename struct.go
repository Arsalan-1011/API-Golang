package main

type Admin struct {
	ID              int    `json:"id"`
	First_Name      string `json:"first_name"`
	Last_Name       string `json:"last_name"`
	Organization_ID int    `json:"organization_id"`
	Deleted         int    `json:"deleted"`
}

type Creative struct {
	ID     int `json:"id"`
	Height int `json:"height"`
	Width  int `json:"width"`
}

type DbUserConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type DbConfig struct {
	UserConfig DbUserConfig `json:"database"`
	Host       string       `json:"host"`
	Port       string       `json:"port"`
}

type Message struct {
	Msg string `json:"msg"`
}
