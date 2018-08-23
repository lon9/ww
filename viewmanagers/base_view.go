package viewmanagers

type ViewManager interface {
	GetName() string
	GetEditable() bool
}

type BaseView struct {
	Name     string
	Editable bool
}

func NewBaseView(name string, editable bool) *BaseView {
	return &BaseView{
		name,
		editable,
	}
}

func (b *BaseView) GetName() string {
	return b.Name
}

func (b *BaseView) GetEditable() bool {
	return b.Editable
}
