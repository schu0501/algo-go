package database

import "fmt"

type MySQL struct {
	ConnString string
}

func (db *MySQL) Save(data string) error {
	fmt.Printf("【MySQL层】正在将数据 '%s' 写入数据库...\n\n", data)
	return nil
}
