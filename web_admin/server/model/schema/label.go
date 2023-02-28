package schema

type Label struct {
}

func (uv Label) TableName() string {
    return "Label"
}
