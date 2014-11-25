package command_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/dockpit/pit/command"
)

func TestBuildingOfStates(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	out := bytes.NewBuffer(nil)
	cmd := command.NewBuild(out)

	AssertCommand(t, cmd, []string{
		"-examples", filepath.Join(wd, "..", ".dockpit", "examples"),
		"-states", filepath.Join(wd, "..", ".dockpit", "states"),
	}, `(?s)built.*successful`, out)
}