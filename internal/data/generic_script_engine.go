package data

import (
	"encoding/json"
	"fmt"
	"github.com/securenative/packman/internal/etc"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type genericScriptEngine struct {
	command string
}

func NewGenericScriptEngine(command string) *genericScriptEngine {
	return &genericScriptEngine{command: command}
}

func (this *genericScriptEngine) Run(scriptPath string, flags map[string]string) (map[string]interface{}, error) {
	flagsFile := pwdPath(scriptPath, "~flags.json")
	replyFile := pwdPath(scriptPath, "~reply.json")
	err := etc.WriteFile(flagsFile, flags, etc.JsonEncoder)

	mainCommand, args, err := splitCommand(this.command)
	if err != nil {
		return nil, err
	}

	var cmdArgs []string
	cmdArgs = append(cmdArgs, args...)
	cmdArgs = append(cmdArgs, scriptPath)
	cmdArgs = append(cmdArgs, flagsFile)
	cmdArgs = append(cmdArgs, replyFile)

	cmd := exec.Command(mainCommand, cmdArgs...)
	etc.PrintInfo(fmt.Sprintf("Running %s script file with: '%s'\n", scriptPath, cmd.String()))
	result, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	etc.PrintResponse(string(result) + "\n")
	etc.PrintSuccess("Script was run successfully.\n")
	etc.PrintInfo("Trying to read reply file: %s...\n", replyFile)
	content, err := etc.ReadFile(replyFile)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal([]byte(content), &out)
	if err != nil {
		return nil, err
	}

	os.Remove(flagsFile)
	os.Remove(replyFile)

	return out, nil
}

func pwdPath(scriptPath string, newName string) string {
	fileName := filepath.Base(scriptPath)
	scriptFolder := strings.ReplaceAll(scriptPath, fileName, "")
	return filepath.Join(scriptFolder, newName)
}

func splitCommand(command string) (string, []string, error) {
	parts := strings.Split(command, " ")
	if parts != nil && len(parts) > 0 {
		return parts[0], parts[1:], nil
	}
	return "", nil, fmt.Errorf("cannot parse command %s, the command syntax should be as follows: 'commnad arg1 arg2 arg3 ...'", command)
}
