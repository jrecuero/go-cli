package syntax

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jrecuero/go-cli/tools"
)

// Recorder represents recorded commands.
type Recorder struct {
	commands []interface{}
	enable   bool
}

// Start starts recording any command entered.
func (rec *Recorder) Start() error {
	rec.enable = true
	return nil
}

// Stop stops recording.
func (rec *Recorder) Stop() error {
	rec.enable = false
	length := len(rec.commands)
	if length > 0 {
		rec.commands = rec.commands[0 : length-1]
	}
	return nil
}

// Add adds a new line to the recorder.
func (rec *Recorder) Add(line interface{}) error {
	if rec.enable {
		rec.commands = append(rec.commands, line)
	}
	return nil
}

// Display displays all recorded commadnds.
func (rec *Recorder) Display() error {
	for i, line := range rec.commands {
		tools.ToDisplay("[%d]  %#v\n", i, line)
	}
	return nil
}

// Clean removes all recorded entries.
func (rec *Recorder) Clean() error {
	rec.commands = []interface{}{}
	return nil
}

// Save saves command recorded in the given filename.
func (rec *Recorder) Save(filename string, appendto bool) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return tools.ERROR(err, true, "Error saving recorder: %#v\n", err)
	}
	defer file.Close()
	for _, line := range rec.commands {
		file.WriteString(fmt.Sprintf("%v\n", line))
	}
	return nil
}

// Load loads recorded commands from the given filename.
func (rec *Recorder) Load(filename string, appendto bool) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return tools.ERROR(err, true, "Error loading recorder: %#v\n", err)
	}
	rec.commands = []interface{}{}
	for _, line := range strings.Split(string(data), "\n") {
		if len(strings.TrimSpace(line)) != 0 {
			rec.commands = append(rec.commands, line)
		}
	}
	return nil
}

// Play plays commands in the recorder.
func (rec *Recorder) Play(m *Matcher) error {
	var retcode error
	savedContext := *m.Ctx
	for _, line := range rec.commands {
		tools.ToDisplay("Playing %#v\n", line)
		if _, err := m.Execute(line); err != nil {
			retcode = tools.ERROR(err, false, "Playing recorder error: %#v\n", err)
			break
		}
	}
	m.Ctx = &savedContext
	return retcode
}

// NewRecorder creates a new recorder instance.
func NewRecorder() *Recorder {
	return &Recorder{}
}
