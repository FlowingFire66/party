package controller

import (
	"encoding/json"
	"fmt"
	"github.com/FlowingFire66/party/model"
	"github.com/FlowingFire66/party/service"
	"io/ioutil"
	"net/http"
)

type UserController struct {
}

func QryUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	println("json:", string(body))

	var a model.UserQry
	if err = json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return
	}

	//fmt.Println("json:", string(b))
	user := service.QryUser(a.UserId)

	user2, err := json.Marshal(user)
	if err != nil {

		fmt.Println("Umarshal failed:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//fmt.Print(a)
	w.Write(user2)
}
