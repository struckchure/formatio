package lib

import "sync"

var mutexMap sync.Map

type IMutext interface {
	CreateMutext(name string) *sync.Mutex
}

type Mutext struct{}

// CreateMutext implements IMutext.
func (m *Mutext) CreateMutext(name string) *sync.Mutex {
	mutex, ok := mutexMap.Load(name)
	if !ok {
		newMutex := &sync.Mutex{}
		mutex, _ = mutexMap.LoadOrStore(name, newMutex)
	}
	return mutex.(*sync.Mutex)
}

func NewMutext() IMutext {

	return &Mutext{}
}
