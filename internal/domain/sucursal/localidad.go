package sucursal

import "fmt"

type localidadSucursal struct {
	localidad string
}

// Localidad define operaciones sobre una localidad de sucursal
type Localidad interface {
	Value() string
}

// NewSucursalLocalidad crea una Lucalidad para una sucursal
func NewSucursalLocalidad(localidad string) (Localidad, error) {
	if localidad == "" {
		return nil, fmt.Errorf("Localidad no puede estar vacia")
	}
	return localidadSucursal{localidad: localidad}, nil
}

func (l localidadSucursal) Value() string {
	return l.localidad
}
