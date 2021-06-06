package cmd

import (
	"api-server/database"
	"fmt"
)

// Migrate は DB マイグレーションを行う
func Migrate() {
	fmt.Println("DBマイグレーションを開始します")
	connector := database.NewConnector()
	db, err := connector.Connect()
	if err != nil {
		fmt.Println("DB接続に失敗しました")
		fmt.Println(err)
		return
	}

	/*
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("マイグレーションに失敗しました")
		fmt.Println(err)
		return
	}
	*/

	fmt.Println("マイグレーションが完了しました")
}
