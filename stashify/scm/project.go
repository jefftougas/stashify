package scm

type Project interface {
	Create(name string)
	Request(resource string)
	List()
}
