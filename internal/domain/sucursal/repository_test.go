package sucursal_test

import "github.com/wilgustavo/turnos/internal/domain/sucursal"

type repositoryMock struct {
	llamadas map[string]int
}

func NewRepositoryMock() *repositoryMock {
	llamadas := make(map[string]int)
	llamadas["Save"] = 0
	return &repositoryMock{llamadas: llamadas}
}

func (m *repositoryMock) Save(sucursal sucursal.Sucursal) error {
	m.llamadas["Save"]++
	return nil
}

func (m *repositoryMock) Llamadas(metodo string) int {
	return m.llamadas[metodo]
}
