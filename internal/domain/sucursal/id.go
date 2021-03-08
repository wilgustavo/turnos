package sucursal

import "fmt"

type sucursalID struct {
	id string
}

// ID define operaciones sobre el ID de la sucursal
type ID interface {
	Value() string
}

// NewSucursalID crea un ID de sucursal
func NewSucursalID(id string) (ID, error) {
	if id == "" {
		return nil, fmt.Errorf("Id no puede ser vacío")
	}
	return &sucursalID{id: id}, nil
}

func (s sucursalID) Value() string {
	return s.id
}
