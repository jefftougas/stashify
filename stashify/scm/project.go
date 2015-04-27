package scm

type Project interface {
	Create(name string)
	List()
}
