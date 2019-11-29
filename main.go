package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Tag struct {
	username string `json:"username"`
	password int  `json:"password"`
}
func main() {
	//连接mysql
	db, err := sql.Open("mysql", "root:yyhaungjq@(127.0.0.1)test")
	if err != nil {
		fmt.Println("连接数据失败")
	} else {
		fmt.Println("连接数据成功")
	}
	defer db.Close() //关闭

	//增加数据到数据库中
	insert, err := db.Query("INSERT INTO TEST VALUES ('Anselker',1234567)")
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	//删除数据
	//delete,err:=db.Query("delete from test")
	//if err!=nil{
	//	panic(err.Error())
	//}
	//defer  delete.Close()

	//更改数据
	//func updateDB(db *sql.DB)  {
	//	stmt, err := db.Prepare("UPDATE test SET name = 'aaa' WHERE id = 1")
	//	if err != nil{
	//		log.Fatal(err)
	//	}
	//	stmt.Exec();
	//}


	//查找数据


	func selectDB(db *sql.DB)  {

		stmt, err := db.Query("select * from test;")

		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		for stmt.Next(){
			var Password Tag
			var Username Tag
			err := stmt.Scan(&Password,&Username)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(Password,Username)

		}

	}
}