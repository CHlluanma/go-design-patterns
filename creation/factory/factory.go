package factory

type Factory interface {
	FactoryMethod(owner string) Product
}

type ConcreteFactory struct {
}

func (cf *ConcreteFactory) FactoryMethod(owner string) Product {
	p := &ConcreteProduct{}
	return p
}

type Product interface {
	Use() error
}

type ConcreteProduct struct{}

func (p *ConcreteProduct) Use() error {
	return nil
}
