package trays

import (
	"bytes"
	"github.com/ArcaneDiver/your-tray/log"
	"github.com/getlantern/systray"
	"github.com/pkg/errors"
	"io/ioutil"
	"os/exec"
	"text/template"
)

type Item struct {
	Name 		string `yaml:"text"`
	Command		string `yaml:"command"`
	Tooltip		string `yaml:"tooltip"`
	Type		string `yaml:"type"`
	MenuItem	*systray.MenuItem
}

const (
	TypeCommand string = "command"
	TypeData	string = "data"
)

func (i *Item) Register() chan struct{}  {
	i.MenuItem = systray.AddMenuItem(i.Name, i.Tooltip)

	if i.IsDynamic() {
		i.MenuItem.Disable()
	}

	return i.MenuItem.ClickedCh
}

func (i *Item) IsValidType() bool {
	switch i.Type {
	case TypeCommand, TypeData:
		return true
	default:
		return false
	}
}

func (i *Item) IsDynamic() bool {
	switch i.Type {
	case TypeData:
		return true
	default:
		return false
	}
}

func (i *Item) HandleTick() error {
	output, err := i.ExecCommand()
	if err != nil {
		return errors.Wrap(err, "Item.HandleTick")
	}

	var parsedTemplate bytes.Buffer
	tmpl, err := template.New("data").Parse(i.Name)
	tmpl.Execute(&parsedTemplate, map[string]string {
		"output": *output,
	})

	i.MenuItem.SetTitle(parsedTemplate.String())

	return nil
}

func (i *Item) ExecCommand() (*string, error) {
	cmd := exec.Command("bash", "-c", i.Command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "Item.ExecCommand")
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, errors.Wrap(err, "Item.ExecCommand")
	}

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "Item.ExecCommand")
	}

	dataOut, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, errors.Wrap(err, "Item.ExecCommand")
	}
	outAsString := string(dataOut)

	dataErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return nil, errors.Wrap(err, "Item.ExecCommand")
	}
	errAsString := string(dataErr)

	if err := cmd.Wait(); err != nil {
		logEntry := log.Log.WithFields(map[string]interface{}{
			"command": i.Command,
			"type":    "stderr",
		})
		if len(dataErr) > 0 {
			logEntry.Error(errAsString)
		} else {
			logEntry.Error()
		}

		return nil, errors.Wrap(err, "Item.ExecCommand")
	}

	logEntry := log.Log.WithFields(map[string]interface{}{
		"command": i.Command,
		"type":    "stdout",
	})
	if len(dataOut) > 0 {
		logEntry.Debug(outAsString)
	} else {
		logEntry.Debug()
	}

	return &outAsString, nil
}