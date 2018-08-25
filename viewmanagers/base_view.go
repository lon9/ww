package viewmanagers

// ViewManager is struct for view manager
type ViewManager interface {
	GetName() string
	GetEditable() bool
}

// BaseView is base struct for views
type BaseView struct {
	Name     string
	Editable bool
}

// NewBaseView is constructor
func NewBaseView(name string, editable bool) *BaseView {
	return &BaseView{
		name,
		editable,
	}
}

// GetName returns name
func (b *BaseView) GetName() string {
	return b.Name
}

// GetEditable returns editable
func (b *BaseView) GetEditable() bool {
	return b.Editable
}
