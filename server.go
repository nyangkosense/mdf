package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mdf/format"
	"mdf/utils"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gomarkdown/markdown"
)

var selectedTheme string
var filePath string
var showCheatsheet bool
var showHelp bool

func main() {
	flag.StringVar(&filePath, "filepath", "", "Path to the Markdown file")
	flag.StringVar(&selectedTheme, "theme", "github-dark", "Theme to use (github-dark, github-light, dracula, nord)")
	flag.BoolVar(&showCheatsheet, "cheatsheet", false, "Display Markdown syntax cheatsheet")
	flag.BoolVar(&showHelp, "help", false, "Display help information")
	flag.Parse()

	if showHelp {
		utils.PrintHelp()
		os.Exit(0)
	}

	if showCheatsheet {
		fmt.Println(utils.Cheatsheet)
		os.Exit(0)
	}

	if filePath == "" {
		fmt.Println("Please provide a filepath using the --filepath flag")
		fmt.Println("Use --help for more information")
		os.Exit(1)
	}

	if _, ok := utils.Themes[selectedTheme]; !ok {
		fmt.Println("Invalid theme. Using default theme (github-dark)")
		selectedTheme = "github-dark"
	}
	fmt.Printf("Using theme: %s\n", selectedTheme) // Add this line

	http.HandleFunc("/", handleMarkdown)
	http.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) {
		format.ServeHTML(w, r, selectedTheme)
	})

	go func() {
		fmt.Println("Server is running on http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	time.Sleep(100 * time.Millisecond) // give the server a moment to start
	openBrowser("http://localhost:8080/view")

	select {}
}

func handleMarkdown(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	html := markdown.ToHTML(content, nil, nil)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(html)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
