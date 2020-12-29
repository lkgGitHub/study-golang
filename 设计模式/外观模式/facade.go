package facade

func NewCarModel() *CarModel {
	return &CarModel{}
}

type CarModel struct{}

func (c *CarModel) SetModel() {
	println("carModel - SetModel")
}

type CarEngine struct{}

func (c *CarEngine) SetEngine() {
	println("CarEngine - SetEngine")
}

type CarBody struct{}

func (c *CarBody) SetCarBody() {
	println("CarBody - SetCarBody")
}

type CarFacade struct {
	model  CarModel
	engine CarEngine
	body   CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}

func (c *CarFacade) CreateCompleteCar() {
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetCarBody()
}
