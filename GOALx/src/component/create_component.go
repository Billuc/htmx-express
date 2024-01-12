package component

func CreateComponent(name string) *Component {
	return &Component{
		Name: name,
	}
}
