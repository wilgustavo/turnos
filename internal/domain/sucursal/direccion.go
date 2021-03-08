package sucursal

import "fmt"

type sucursalDireccion struct {
	direccion string
}

// Direccion define operaciones sobre la direccion de la sucursal
type Direccion interface {
	Value() string
}

// NewSucursalDireccion crea una direccion de sucursal
func NewSucursalDireccion(direccion string) (Direccion, error) {
	if direccion == "" {
		return nil, fmt.Errorf("direccion no puede esta vacia")
	}
	return sucursalDireccion{direccion: direccion}, nil
}

func (d sucursalDireccion) Value() string {
	return d.direccion
}
