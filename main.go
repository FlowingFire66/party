package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/FlowingFire66/party/controller"
	"github.com/FlowingFire66/party/logger"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AutotaskRequest struct {
	RequestID string     `json:"requestid"`
	Clone     CloneModel `json:"clone"`
	Push      PushModel  `json:"push"`
}

type CloneModel struct {
	//TODO
	//"Method": string `json:"ceph"`
	RequestID   string `json:"requestid"`
	CallbackURL string `json:"callbackurl"`
}

type PushModel struct {
	RequestID   string `json:"requestiD"`
	CallbackURL string `json:"callbackuRL"`
	IP          string `json:"remoteip"`
	Port        int    `json:"remoteport"`
	User        string `json:"user"`
}

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(50)
	db.Ping()
}
func main() {
	defer func() {
		if p := recover(); p != nil {
			logger.Log.Error("Recovered panic: %s\n", p)
		}
	}()
	logger.Log.Info("in main args:%v", os.Args)

	logger.Log.Error("eerror %v", "error")
	//设置访问路由
	http.HandleFunc("/", controller.QryUser)
	http.HandleFunc("/pool", pool)
	a := 100
	b := 100
	c := a - b
	_ = a / c

	//设置访问的ip和端口
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
func test(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	//defer fmt.Fprintf(w, "ok\n")

	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	println("json:", string(body))

	var a AutotaskRequest
	if err = json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return
	}
	fmt.Print(a)
	w.Write(body)
}
func pool(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT user_id FROM user limit 1")
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	fmt.Println(record)
	fmt.Fprintln(w, "finish")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
