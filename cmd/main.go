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
	tasks := getTasksFromConf()
	if len(tasks) == 0 {
		// targets file just got created
		return
	}

	monitorTasks(&tasks) // fst call to register tasks

	// setup a ticker to listen to
	tickCycle := 10 * time.Second
	ticker := time.NewTicker(tickCycle)

	// create a channel to which POSIX signals will be deleivered to
	// make it a buffered channel (manage only one signal at the time)
	sigPower := make(chan os.Signal, 1)
	sigContinue := make(chan os.Signal, 1)
	sigTerm := make(chan os.Signal, 1)
	sigUsr1 := make(chan os.Signal, 1) // will be used to update tasks slice with new config

	// SIGPWR=battery|power lvl low (shutdown|hibernation|suspend mode)
	signal.Notify(sigPower, syscall.SIGPWR)
	signal.Notify(sigContinue, syscall.SIGCONT)
	signal.Notify(sigTerm, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(sigUsr1, syscall.SIGUSR1)

	for {
		select {
		case <-ticker.C:
			{
				monitorTasks(&tasks)
			}
		case <-sigUsr1:
			{
				newConfigTasks := getTasksFromConf()

				for _, task := range newConfigTasks {

					alreadyMonitored := false
					lName := strings.ToLower(task.Name)

					for _, monitoredTask := range tasks {

						// only append non monitored yet tasks
						if strings.Compare(strings.ToLower(monitoredTask.Name), lName) == 0 {
							alreadyMonitored = true
						}
					}
					if !alreadyMonitored {
						tasks = append(tasks, types.Task{
							Name:       task.Name,
							LaunchedAt: time.Now(),
						})
					}
				}
			}
		case <-sigPower:
			{
				// make the ticker.C = nil so as
				// its case will be ignored
				ticker.C = nil

				// flush the current stats to db
				saveCurrentStats(tasks)

			}
		case <-sigContinue:
			{
				// make the ticker.C valid again (!nil)
				ticker = time.NewTicker(tickCycle)

				// update the tasks, this can be done after ticker re-initialization
				// since the tickCycle is long then the update job
				var nonInitializedTimeInstant time.Time
				for i, task := range tasks {

					// ignore non used tasks
					if task.LaunchedAt.Equal(nonInitializedTimeInstant) {
						continue
					}

					tasks[i].LaunchedAt = time.Now()
				}

			}
		case <-sigTerm:
			{
				return
			}
		}
	}
}

func getTasksFromConf() []types.Task {

	// load the tokens from the file to a slice of tasks
	tasks := []types.Task{}
	filename := "/home/massigy/.config/mmtime/targets"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	utils.Check(err)
	defer file.Close()

	chars, err := ioutil.ReadAll(file)
	utils.Check(err)

	fileContent := string(chars)

	if len(fileContent) <= 1 {
		file.WriteString("# Add the applications that you want to track your usage time in, each application in a seperated line\n")
		file.WriteString("\n")
		file.WriteString(os.Args[0]) //file just created, add curr process name to it, the user will change it as he wants
		tasks = append(tasks, types.Task{
			Name:       os.Args[0],
			LaunchedAt: time.Now(),
		})

		return tasks
	}

	tasksNames := strings.Split(fileContent, "\n")
	tasksNamesInStr := strings.Join(tasksNames, " ")

	for _, name := range tasksNames {
		// ignore commented lines
		if strings.HasPrefix(name, "#") {
			continue
		}

		// ignore empty lines
		if len(name) <= 1 {
			continue
		}

		// make sure to not allow duplicates
		if !strings.Contains(tasksNamesInStr, name) {
			continue
		}

		// remove current name from the lists of tasks to monitor
		tasksNamesInStr = strings.ReplaceAll(tasksNamesInStr, name, "")
		tasksNamesInStr = strings.ReplaceAll(tasksNamesInStr, strings.Title(name), "")
		tasksNamesInStr = strings.ReplaceAll(tasksNamesInStr, strings.ToUpper(name), "")
		tasksNamesInStr = strings.ReplaceAll(tasksNamesInStr, strings.ToLower(name), "")

		tasks = append(tasks, types.Task{
			Name: name,
		})
	}
	return tasks
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

	for i, task := range *tasks {

		(*tasks)[i].Running = false

		for _, line := range stdOutLines {
			// every line is a process name

			if strings.Compare(strings.ToLower(line), strings.ToLower(task.Name)) == 0 {

				(*tasks)[i].Running = true

				if task.LaunchedAt.Equal(nonInitializedTimeInstant) {
					(*tasks)[i].LaunchedAt = time.Now()
					continue
				}

				(*tasks)[i].UsedFor = time.Since((*tasks)[i].LaunchedAt)
			}
		}
	}

	// log the tasks
	fmt.Println("----------Logging all tasks---------------")
	for _, task := range *tasks {
		fmt.Println(task)
	}
	fmt.Println("------------------------------------------")
}

func saveCurrentStats(tasks []types.Task) {
	// db file
	filename := "/home/massigy/.local/share/mmtime/targets.stats.db"

	// open the file for RW(do not create to add the desc comment at first)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)

	utils.Check(err)
	defer file.Close()

	// read all the file's content
	chars, err := ioutil.ReadAll(file)
	utils.Check(err)

	fileContent := string(chars)
	if len(fileContent) <= 1 {
		// file just got created, write the dsc comment into it
		file.WriteString("# This file stores the mmtime stats according to your time usage of the applications that you specified in ~/.config/mmtime/targets\n")
		file.WriteString("# For every application/process just sum up the usage times for any given date to get the total usage during that day.\n")
		file.WriteString("\n")
		file.WriteString("Process\tDate\t\tUsage\n")
	}

	var nonInitializedTimeInstant time.Time

	// write the stats to the file
	for _, task := range tasks {

		// ignore task/process that did not run at all,
		if task.LaunchedAt.Equal(nonInitializedTimeInstant) {
			continue
		}

		// ignore task/process that've stopped running,
		if !task.Running {
			continue
		}

		file.WriteString(
			fmt.Sprintf("%s\t%v\t%v\n",
				task.Name,
				task.LaunchedAt.Format("2006-01-02"),
				task.UsedFor,
			),
		)
	}

}
