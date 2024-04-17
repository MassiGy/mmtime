package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"mmtime/types"
	"mmtime/utils"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	filename := "/home/massigy/.config/mmtime/targets"

	// open the ~/.config/mmtime/targets file
	file, err := os.Open(filename)

	if err != nil {
		// create the file and exit
		err := os.WriteFile(
			filename,
			[]byte("# Add the applications that you want to track your usage time in, each application in a seperated line\n"),
			0666,
		)
		utils.Check(err)
		return
	}

	defer file.Close()

	// load the tokens from the file to a slice of tasks
	tasks := []types.Task{}

	chars, err := ioutil.ReadAll(file)
	utils.Check(err)

	fileContent := string(chars)
	tasksNames := strings.Split(fileContent, "\n")

	for _, name := range tasksNames {
		// ignore commented lines
		if strings.HasPrefix(name, "#") {
			continue
		}

		// ignore empty lines
		if len(name) <= 1 {
			continue
		}

		tasks = append(tasks, types.Task{
			Name: name,
		})
	}

	monitorTasks(&tasks)

	// setup a ticker to listen to
	tickCycle := 10 * time.Second
	ticker := time.NewTicker(tickCycle)

	// create a channel to which POSIX signals will be deleivered to
	// make it a buffered channel (manage only one signal at the time)
	sigPower := make(chan os.Signal, 1)
	sigContinue := make(chan os.Signal, 1)
	sigTerm := make(chan os.Signal, 1)

	// SIGPWR is sent when the battery|power level is low
	signal.Notify(sigPower, syscall.SIGPWR)

	signal.Notify(sigContinue, syscall.SIGCONT)
	signal.Notify(sigTerm, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-ticker.C:
			{
				monitorTasks(&tasks)
			}
		case <-sigPower:
			{
				// make the ticker.C = nil so as
				// its case will be ignored
				ticker.C = nil
			}
		case <-sigContinue:
			{
				// make the ticker.C valid again (!nil)
				ticker = time.NewTicker(tickCycle)
			}
		case <-sigTerm:
			{
				return
			}
		}
	}
}
func monitorTasks(tasks *[]types.Task) {

	shellCommand := "ps -axco command| tail -n +2|uniq"
	shellInterpreter := "bash"

	var cmdStdOut bytes.Buffer
	var cmdStdErr bytes.Buffer

	cmd := exec.Command(shellInterpreter, "-c", shellCommand)
	cmd.Stdout = &cmdStdOut
	cmd.Stderr = &cmdStdErr

	err := cmd.Run()
	utils.Check(err)

	if len(cmdStdErr.String()) >= 1 {
		panic(errors.New("could not fetch currently running processes"))
	}

	stdOutLines := strings.Split(cmdStdOut.String(), "\n")

	var nonInitializedTimeInstant time.Time

	for _, line := range stdOutLines {
		// every line is a process name

		for i, task := range *tasks {

			if strings.Compare(strings.ToLower(line), strings.ToLower(task.Name)) == 0 {

				if task.LaunchedAt.Equal(nonInitializedTimeInstant) {
					(*tasks)[i].LaunchedAt = time.Now()
					continue
				}

				(*tasks)[i].UsedFor = time.Since((*tasks)[i].LaunchedAt)
			}
		}
	}

	// log the tasks
	for _, task := range *tasks {
		fmt.Println(task)
	}
}
