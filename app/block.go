package main

type Block struct {
	x, y      int
	isBlocked bool
}

type BlockController struct {
	blocks map[int]map[int]Block
	asdasd bool

	mapSizeX, mapSizeY int
}

func New(mapSizeX, mapSizeY int, engine *GameEngine) *BlockController {
	blockCtrl := &BlockController{mapSizeX: mapSizeX, mapSizeY: mapSizeY}
	engine.updates = append(engine.updates, blockCtrl)

	return blockCtrl
}

func (ctrl *BlockController) AddBlock(_x, _y int) {
	ctrl.blocks[_y][_x] = Block{x: _x, y: _y, isBlocked: true}
}

func (ctrl *BlockController) RemoveBlock(_x, _y int) {
	delete(ctrl.blocks[_y], _x)
}

func (ctrl *BlockController) Draw() string {
	uiString := ""
	for y := 0; y < ctrl.mapSizeY; y++ {
		for x := 0; x < ctrl.mapSizeX; x++ {
			uiString += " "
			if block, ok := ctrl.blocks[y][x]; ok {
				if block.isBlocked {
					uiString += "[X]"
					continue
				}
			}

			//original map
			uiString += "[ ]"
		}
		uiString += "\n"
	}
	return uiString
}

func (ctrl *BlockController) Update() {

}
