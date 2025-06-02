package appliances

type Microwave struct {
	typeName string
}

func (fridge *Microwave) Start() {
	fridge.typeName = "Microwave"
}

func (fridge *Microwave) GetPurpose() string {
	return fridge.typeName + " heat stuff up"
}
