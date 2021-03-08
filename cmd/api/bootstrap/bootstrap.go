package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wilgustavo/turnos/internal/platform/server"
	"github.com/wilgustavo/turnos/internal/platform/storage/mysql"
)

const (
	host       = "localhost"
	port       = 8080
	dbUser     = "user"
	dbPassword = "password"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "turnos"
)

// Run the server
func Run() error {
	mySQLURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mySQLURI)
	if err != nil {
		return err
	}
	sucursalRepository := mysql.NewSucursalRepository(db)
	srv := server.New(host, port, sucursalRepository)
	return srv.Run()
}
