package sucursal

import "fmt"

type sucursal struct {
	id        ID
	nombre    Nombre
	direccion Direccion
	localidad Localidad
}

// Sucursal define operaciones sobre una sucursal
type Sucursal interface {
	ID() string
	Nombre() string
	Direccion() string
	Localidad() string
}

// Repository repositorio de Sucursal
type Repository interface {
	Save(sucursal Sucursal) error
}

// NewSucursal crea una nueva sucursal
func NewSucursal(id ID, nombre Nombre, direccion Direccion, localidad Localidad) (Sucursal, error) {
	if id == nil {
		return nil, fmt.Errorf("Id no puede ser nulo")
	}
	if nombre == nil {
		return nil, fmt.Errorf("nombre no puede ser nulo")
	}
	if direccion == nil {
		return nil, fmt.Errorf("localidad no puede ser nulo")
	}
	if localidad == nil {
		return nil, fmt.Errorf("localidad no puede ser nulo")
	}
	return &sucursal{id: id, nombre: nombre, direccion: direccion, localidad: localidad}, nil
}

func (s *sucursal) ID() string {
	return s.id.Value()
}

func (s *sucursal) Nombre() string {
	return s.nombre.Value()
}

func (s *sucursal) Direccion() string {
	return s.direccion.Value()
}

func (s *sucursal) Localidad() string {
	return s.localidad.Value()
}
