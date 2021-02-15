package trays

import (
	"github.com/ArcaneDiver/your-tray/log"
	"github.com/getlantern/systray"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Tray struct {
	Name	string `yaml:"name,omitempty"`
	Icon	string `yaml:"icon"`
	Tooltip	string `yaml:"tooltip,omitempty"`
	Items	[]Item `yaml:"items,flow"`
}

func (t *Tray) Init(sync chan bool) {
	go systray.Run(t.run(sync), t.exit)
}

func (t *Tray) run(sync chan bool) func() {
	return func() {
		icon, err := t.readIcon()
		if err != nil {
			log.Log.Error(errors.Wrap(err, "Tray.run"))
			log.Log.Exit(1)
		}

		systray.SetTitle(t.Name)
		systray.SetIcon(icon)
		systray.SetTooltip(t.Tooltip)

		for idx, item := range t.Items {
			ch := t.Items[idx].Register()

			if !item.IsDynamic() {
				item := item
				go func() {
					for {
						<-ch
						if _, err := item.ExecCommand(); err != nil {
							log.Log.WithField("item", item.Name).Error(errors.Wrap(err, "Tray.run"))
						}
					}
				}()
			}

		}
		sync <- true
	}

}

func (t *Tray) exit()  {

}



func (t *Tray) readIcon() ([]byte, error)  {
	data, err := ioutil.ReadFile(t.Icon)
	if err != nil {
		return nil, errors.Wrap(err, "Tray.readIcon")
	}

	return data, nil
}