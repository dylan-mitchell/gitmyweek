package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Commit struct {
	Date    string
	Message string
	Project string
}

var (
	mon = make(map[string][]Commit)
	tue = make(map[string][]Commit)
	wed = make(map[string][]Commit)
	thu = make(map[string][]Commit)
	fri = make(map[string][]Commit)
	sat = make(map[string][]Commit)
	sun = make(map[string][]Commit)
)

func exists(filePath string) bool {
	exists := true

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exists = false
	}

	return exists
}

func readFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func parseOutput(out, project string) []Commit {
	lines := strings.Split(out, "\n")
	var commits []Commit
	for _, line := range lines {
		if line != "" {
			msg := strings.SplitAfterN(line, ": ", 2)[1]
			msg = msg[:len(msg)-1]
			c := Commit{
				Date:    line[1:4],
				Message: msg,
				Project: project,
			}
			commits = append(commits, c)
			switch c.Date {
			case "Sun":
				sun[c.Project] = append(sun[c.Project], c)
			case "Mon":
				mon[c.Project] = append(mon[c.Project], c)
			case "Tue":
				tue[c.Project] = append(tue[c.Project], c)
			case "Wed":
				wed[c.Project] = append(wed[c.Project], c)
			case "Thu":
				thu[c.Project] = append(thu[c.Project], c)
			case "Fri":
				fri[c.Project] = append(fri[c.Project], c)
			case "Sat":
				sat[c.Project] = append(sat[c.Project], c)
			}
		}
	}

	return commits
}

func executeCommand(path, d, user, start, end string) ([]Commit, error) {
	// goExecutable, _ := exec.LookPath("git")
	// cmdTemplate := "git --git-dir=%s/%s/.git/ log --pretty=format:\"%s: %s\" --after=\"2020-04-06\" --until=\"2020-04-10\" --author=\"dylan-mitchell\""

	cmd := exec.Command("git", "--git-dir", fmt.Sprintf("%s/%s/.git/", path, d), "log", "--pretty=format:\"%ad: %s\"", "--after", start, "--until", end, "--author", user)
	// run command
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	commits := parseOutput(string(output), d)

	return commits, nil
}

func main() {
	dir := flag.String("dir", "", "Path to directory that hold all of your Git repos")
	user := flag.String("author", "", "Your Github username")
	startFlag := flag.String("start", "", "Start of the week")
	flag.Parse()
	if *dir == "" {
		log.Fatal("Please provide a directory to analyze")
	}

	if *user == "" {
		log.Fatal("Please provide an author")
	}

	now := time.Now()
	weekAgo := now.AddDate(0, 0, -7)

	start := weekAgo.Format("2006-01-02")
	end := time.Now().Format("2006-01-02")
	if *startFlag != "" {
		start = *startFlag
		layout := "2006-01-02"
		t, err := time.Parse(layout, start)
		if err != nil {
			log.Fatal("Error Parsing start, needs to be YYYY-MM-DD")
		}
		end = t.AddDate(0, 0, 7).Format("2006-01-02")
	}

	if !exists(*dir) {
		log.Fatal("Please provide a valid path.")
	}

	files := readFiles(*dir)
	var dirsToAnalyze []string
	for _, f := range files {
		if f.IsDir() {
			dirsToAnalyze = append(dirsToAnalyze, f.Name())
		}
	}

	var allCommits []Commit
	for _, d := range dirsToAnalyze {
		commits, _ := executeCommand(*dir, d, *user, start, end)
		if len(commits) != 0 {
			allCommits = append(allCommits, commits...)
		}
	}

	if len(mon) != 0 {
		fmt.Println("Monday")
	}
	for k, v := range mon {
		fmt.Println("\t" + k)
		for _, c := range v {
			fmt.Println("\t\t" + c.Message)
		}
	}

	if len(tue) != 0 {
		fmt.Println("Tuesday")
	}
	for k, v := range tue {
		fmt.Println("\t" + k)
		for _, c := range v {
			fmt.Println("\t\t" + c.Message)
		}
	}

	if len(wed) != 0 {
		fmt.Println("Wednesday")
	}
	for k, v := range wed {
		fmt.Println("\t" + k)
		for _, c := range v {
			fmt.Println("\t\t" + c.Message)
		}
	}

	if len(thu) != 0 {
		fmt.Println("Thursday")
	}
	for k, v := range thu {
		fmt.Println("\t" + k)
		for _, c := range v {
			fmt.Println("\t\t" + c.Message)
		}
	}

	if len(fri) != 0 {
		fmt.Println("Friday")
	}
	for k, v := range fri {
		fmt.Println("\t" + k)
		for _, c := range v {
			fmt.Println("\t\t" + c.Message)
		}
	}

	if len(sat) != 0 {
		fmt.Println("Saturday")
	}
	for k, v := range mon {
		fmt.Println("\t" + k)
		for _, c := range v {
			fmt.Println("\t\t" + c.Message)
		}
	}

	if len(sun) != 0 {
		fmt.Println("Sunday")
	}
	for k, v := range mon {
		fmt.Println("\t" + k)
		for _, c := range v {
			fmt.Println("\t\t" + c.Message)
		}
	}

}
