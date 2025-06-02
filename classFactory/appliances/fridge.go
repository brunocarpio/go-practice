package appliances

type Fridge struct {
	typeName string
}

func (fridge *Fridge) Start() {
	fridge.typeName = "Fridge"
}

func (fridge *Fridge) GetPurpose() string {
	return fridge.typeName + " cools stuff down"
}
