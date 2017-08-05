package tui

type UI interface {
	SetTheme(p *Theme)
	SetKeybinding(k interface{}, fn func())
	SetFocusChain(ch FocusChain)
	Run() error
	Repaint()
	Quit()
}

func New(root Widget) UI {
	tcellui, _ := newTcellUI(root)
	return tcellui
}
