package appliances

type Stove struct {
	typeName string
}

func (fridge *Stove) Start() {
	fridge.typeName = "Stove"
}

func (fridge *Stove) GetPurpose() string {
	return fridge.typeName + " cooks food"
}
