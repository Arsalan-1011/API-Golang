package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["ID"]

	res, err := dbm.Query("SELECT `ID`, `First_Name`, `Last_Name`, `Organization_ID`, `Deleted` FROM `userconfig` WHERE `ID`=?", id)

	if err != nil {
		msg := Message{Msg: "User Not Found"}
		json.NewEncoder(w).Encode(msg)
	}

	for res.Next() {
		var c Admin
		err := res.Scan(&c.ID, &c.First_Name, &c.Last_Name, &c.Organization_ID, &c.Deleted)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(c)
	}
	defer res.Close()

}

func getAllAdmins(w http.ResponseWriter, r *http.Request) {
	var admin []Admin
	res, err := dbm.Query("SELECT `ID`, `First_Name`, `Last_Name`, `Organization_ID`, `Deleted` FROM `userconfig`")
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var c Admin
		err := res.Scan(&c.ID, &c.First_Name, &c.Last_Name, &c.Organization_ID, &c.Deleted)

		if err != nil {
			log.Fatal(err)
		}
		admin = append(admin, c)

	}

	resp, _ := json.Marshal(admin)
	w.Write(resp)

}

func createAdmin(w http.ResponseWriter, r *http.Request) {
	var c Admin

	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Last_Name := c.Last_Name
	First_Name := c.First_Name
	id := c.ID
	Organization_ID := c.Organization_ID
	Deleted := c.Deleted

	res, err := dbm.Query("INSERT INTO `userconfig`(`ID`,`First_Name`, `Last_Name`, `Organization_ID`, `Deleted`) VALUES (?,?,?,?,?)", id, First_Name, Last_Name, Organization_ID, Deleted)

	if err != nil {
		msg := Message{Msg: "Not Created"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	msg := Message{Msg: "Created Successfully..."}

	json.NewEncoder(w).Encode(msg)
	defer res.Close()
}

func updateAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["ID"]

	res, err := dbm.Query("UPDATE `userconfig` SET `Deleted` = 0 WHERE `ID` = ?;", id)

	if err != nil {
		log.Fatal(err)
		msg := Message{Msg: "Not Updated"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	msg := Message{Msg: "Updated Successfully..."}
	json.NewEncoder(w).Encode(msg)
	defer res.Close()

}

func deleteAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["ID"]

	res, err := dbm.Query("DELETE FROM `admin` WHERE `ID`=?", id)
	if err != nil {
		msg := Message{Msg: "User Not Found"}
		json.NewEncoder(w).Encode(msg)
	}
	for res.Next() {
		var c Admin
		err := res.Scan(&c.ID, &c.First_Name, &c.Last_Name, &c.Organization_ID, &c.Deleted)

		if err != nil {
			log.Fatal(err)
			msg := Message{Msg: "Not Deleted"}
			json.NewEncoder(w).Encode(msg)
			return
		}
		msg := Message{Msg: "Deleted Successfully..."}
		json.NewEncoder(w).Encode(msg)
		defer res.Close()

	}
}

// Crerative
func AddCreative(w http.ResponseWriter, r *http.Request) {
	var c Creative
	rbody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(rbody, &c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Height := c.Height
	Width := c.Width

	res, err := dbm.Query("INSERT INTO `creative`(`height`,`width`) VALUES (?,?)", Height, Width)

	if err != nil {
		msg := Message{Msg: "Not Created"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	json.NewEncoder(w).Encode(c)
	defer res.Close()
}

func getCrearives(w http.ResponseWriter, r *http.Request) {
	var creat []Creative
	res, err := dbm.Query("SELECT `id`, `height`, `width` FROM `creative`")
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var c Creative
		err := res.Scan(&c.ID, &c.Height, &c.Width)

		if err != nil {
			log.Fatal(err)
		}
		creat = append(creat, c)

	}

	resp, _ := json.Marshal(creat)
	w.Write(resp)

}
