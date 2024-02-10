package daos

import "database/sql"

type GenericDao struct {
	dbConn sql.DB
}

func (g GenericDao) SaveToDB(in interface{}) error {
	// Fire a DB query to save the deal
	return nil
}
