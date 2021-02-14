package trays

import "github.com/getlantern/systray"

type Item struct {
	Name 	string `yaml:"text"`
	Command	string `yaml:"command"`
	Tooltip	string `yaml:"tooltip"`
}

func (i *Item) Register() chan struct{}  {
	item := systray.AddMenuItem(i.Name, i.Tooltip)

	return item.ClickedCh
}
