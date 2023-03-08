package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

func connect() {

	//get the config values
	file, err := ioutil.ReadFile("Config.json")
	if nil != err {
		//("Error %s", err.Error())
		return
	}
	var config DbConfig
	err = json.Unmarshal(file, &config)

	if nil != err {
		//("Error %s", err.Error())
		return
	}

	//use the values now from config
	log.Println(config)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.UserConfig.User, config.UserConfig.Password, config.Host, config.Port, config.UserConfig.Database))
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("connected")
	}
	dbm = db

}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

// var db_connect []Creative

// if err != nil {
// 	log.Fatal(err)
// }

// var c Creative
// err := db.Scan(&c.ID, &c.Height, &c.Width)
// db, err := sql.Open("mysql", "arsalan:root@tcp(127.0.0.1:3306)/test?parseTime=true")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db_connect = append(db_connect, c)

// }

// a,_ := json.Marshal(db_connect)
// print(a)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// func connect() Config {
// 	var config Config
// 	defer configFile.Close()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	jsonParser := json.NewDecoder(configFile)
// 	jsonParser.Decode(&config)
// 	return config
// }
