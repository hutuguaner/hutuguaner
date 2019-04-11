package message

import(
	"net/http"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"encoding/json"
)

func Putmsg(w http.ResponseWriter,r *http.Request){
	fmt.Println(" Putmsg ")
	r.ParseForm()

	var msg string
	msg = r.Form["msg"][0]

	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/hutuguaner")
	defer db.Close()

	if err!=nil{
		fmt.Println(err)
		panic(err)
	}

	stmt,err := db.Prepare("insert msg(msg)values(?)")
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}

	res,err := stmt.Exec(msg)

	id,err := res.LastInsertId()
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(id)
	stmt.Close()

}

func Getmsg(w http.ResponseWriter,r *http.Request){
	fmt.Println(" Getmsg ")
	db,err :=sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/hutuguaner")
	var msgjs []MsgJ
	if err!=nil{
		fmt.Println(err)
		
	}

	defer db.Close()

	rows,err := db.Query("select * from msg")
	if err!=nil{
		fmt.Println(err)
		
	}

	for rows.Next(){
		var msg string
		rows.Scan(&msg)
		msgj := MsgJ{msg}
		msgjs = append(msgjs,msgj)
	}

	responsej :=ResponseJ{0,msgjs}
	js,err := json.Marshal(responsej)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(js)

	defer rows.Close()


}

type ResponseJ struct{
	Code int `json:"code"`
	Data []MsgJ `json:"data"`
}

type MsgJ struct{
	Msg string `json:"msg"`
}

func Delmsg(w http.ResponseWriter,r *http.Request){
	

}