package main

type GameEngine struct {
	updates []IUpdateAble
}

type IUpdateAble interface {
	Update()
}

func (engine *GameEngine) Update() {
	for _, iupdateable := range engine.updates {
		iupdateable.Update()
	}
}
