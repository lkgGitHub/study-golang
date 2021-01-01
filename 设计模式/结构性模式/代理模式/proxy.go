package proxys

type IObject interface {
	objOo(action string)
}

type Object struct {
	action string
}

func (o *Object) objDo(action string) {
	println("I can ", action)
}

type ProxyObject struct {
	object *Object
}

func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = &Object{}
	}
	if action == "Run" {
		 p.object.objDo(action)
	}
}