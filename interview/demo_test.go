package interview

type MyService struct {
	isOpen bool
}

func (s *MyService) Start() {

	if s.isOpen {
		return
	}

	//初始化资源

	s.isOpen = true

}

func (s *MyService) Stop() {

	if !s.isOpen {
		return
	}

	//释放资源

	s.isOpen = false

}

func (s *MyService) Handle() {

	if !s.isOpen {
		return
	}

	//处理

}
