package sucursal

import "fmt"

type sucursalNombre struct {
	nombre string
}

// Nombre define operaciones sobre el nombre de la sucursal
type Nombre interface {
	Value() string
}

// NewSucursalNombre crea un nombre de sucursal
func NewSucursalNombre(nombre string) (Nombre, error) {
	if nombre == "" {
		return nil, fmt.Errorf("Nombre no puede estar vacío")
	}
	return sucursalNombre{nombre: nombre}, nil
}

func (n sucursalNombre) Value() string {
	return n.nombre
}
