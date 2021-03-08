package mysql

import (
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	"github.com/wilgustavo/turnos/internal/domain/sucursal"
)

const tableName = "sucursal"

type sucursalTable struct {
	ID        string `db:"id"`
	Nombre    string `db:"nombre"`
	Direccion string `db:"direccion"`
	Localidad string `db:"localidad"`
}

// SucursalRepository datos del repositorio de sucursal
type SucursalRepository struct {
	db *sql.DB
}

// NewSucursalRepository crea un repositorio de sucursal con MySQL
func NewSucursalRepository(db *sql.DB) *SucursalRepository {
	return &SucursalRepository{db: db}
}

// Save graba una nueva sucursal
func (r *SucursalRepository) Save(sucursal sucursal.Sucursal) error {
	sqlStruct := sqlbuilder.NewStruct(new(sucursalTable))
	query, args := sqlStruct.InsertInto(tableName, sucursalTable{
		ID:        sucursal.ID(),
		Nombre:    sucursal.Nombre(),
		Direccion: sucursal.Direccion(),
		Localidad: sucursal.Localidad(),
	}).Build()

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("No se ha podido persistir los datos de la sucursal: %v", err)
	}
	return nil
}
