package trays

import (
	"github.com/ArcaneDiver/your-tray/log"
	"github.com/getlantern/systray"
	"github.com/pkg/errors"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Tray struct {
	Name	string `yaml:"name"`
	Icon	string `yaml:"icon"`
	Tooltip	string `yaml:"tooltip"`
	Items	[]Item `yaml:"items,flow"`
}

func (t *Tray) Init() {
	systray.Run(t.run, t.exit)
}

func (t *Tray) run() {
	icon, err := t.readIcon()
	if err != nil {
		log.Log.Error(errors.Wrap(err, "Tray.run"))
		log.Log.Exit(1)
	}

	systray.SetTitle(t.Name)
	systray.SetIcon(icon)
	systray.SetTooltip(t.Tooltip)

	for _, item := range t.Items {
		ch := item.Register()

		item := item
		go func() {
			for {
				<-ch
				if err := t.execCommand(item.Command); err != nil {
					log.Log.WithField("item", item.Name).Error(errors.Wrap(err, "Tray.run"))
				}
			}
		}()
	}
}

func (t *Tray) exit()  {

}

func (t *Tray) execCommand(command string) error {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "Tray.execCommand")
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return errors.Wrap(err, "Tray.execCommand")
	}

	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "Tray.execCommand")
	}

	dataOut, err := ioutil.ReadAll(stdout)
	if err != nil {
		return errors.Wrap(err, "Tray.execCommand")
	}

	dataErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return errors.Wrap(err, "Tray.execCommand")
	}

	if err := cmd.Wait(); err != nil {
		logEntry := log.Log.WithFields(map[string]interface{}{
			"command": command,
			"type":    "stderr",
		})
		if len(dataErr) > 0 {
			logEntry.Error(string(dataErr))
		} else {
			logEntry.Error()
		}

		return errors.Wrap(err, "Tray.execCommand")
	}

	logEntry := log.Log.WithFields(map[string]interface{}{
		"command": command,
		"type":    "stdout",
	})
	if len(dataOut) > 0 {
		logEntry.Info(string(dataOut))
	} else {
		logEntry.Info()
	}

	return nil
}

func (t *Tray) readIcon() ([]byte, error)  {
	data, err := ioutil.ReadFile(t.Icon)
	if err != nil {
		return nil, errors.Wrap(err, "Tray.readIcon")
	}

	return data, nil
}