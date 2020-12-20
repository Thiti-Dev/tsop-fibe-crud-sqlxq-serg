package database

import "fmt"

// DoCrudOperationWithTargetQuery will take a set of SQL statement and do the operation and return whether it successes or not
func DoCrudOperationWithTargetQuery(querystr string) error{
	db,err := OpenPgSQL()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = db.Exec(querystr)
	return err
}