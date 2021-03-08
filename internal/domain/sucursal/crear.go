package sucursal

// CrearSucursalService Servicio para crear una sucursal
type CrearSucursalService func(id string, nombre string, direccion string, localidad string) error

// NewCrearSucursalService crea un servicio para crear una sucursal
func NewCrearSucursalService(repository Repository) CrearSucursalService {
	return func(id string, nombre string, direccion string, localidad string) error {
		sucursalID, err := NewSucursalID(id)
		if err != nil {
			return err
		}
		sucursalNombre, err := NewSucursalNombre(nombre)
		if err != nil {
			return err
		}
		sucursalDireccion, err := NewSucursalDireccion(direccion)
		if err != nil {
			return err
		}
		sucursalLocalidad, err := NewSucursalLocalidad(localidad)
		if err != nil {
			return err
		}

		sucursal, err := NewSucursal(sucursalID, sucursalNombre, sucursalDireccion, sucursalLocalidad)

		return repository.Save(sucursal)
	}
}
