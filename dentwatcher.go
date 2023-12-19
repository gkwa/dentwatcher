package dentwatcher

import (
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"os"
)

type Options struct {
	LogFormat  string
	LogLevel   string
	FolderPath string
}

func Execute() int {
	options := parseArgs()

	logger, err := getLogger(options.LogLevel, options.LogFormat)
	if err != nil {
		slog.Error("getLogger", "error", err)
		return 1
	}

	slog.SetDefault(logger)

	run(options)
	return 0
}

func parseArgs() Options {
	options := Options{}

	flag.StringVar(&options.LogLevel, "log-level", "info", "Log level (debug, info, warn, error), defult: info")
	flag.StringVar(&options.LogFormat, "log-format", "", "Log format (text or json)")
	flag.StringVar(&options.FolderPath, "folder", "", "Specify the folder path")

	flag.Parse()

	return options
}

const tasksJSONTemplate = `
{
  "version": "2.0.0",
  "presentation": {
    "echo": false,
    "reveal": "always",
    "focus": true,
	"open": true,
    "panel": "dedicated",
    "showReuseMessage": true
  },
  "tasks": [
    {
      "label": "Terminatel",
      "dependsOn": [
        "Terminal"
      ],
      "runOptions": {
        "runOn": "folderOpen"
      }
    },
    {
      "label": "Terminal",
      "type": "shell",
      "command": "bash",
      "options": {
        "shell": {
          "executable": "bash",
		  "args": ["-l", "-c", "bash"]
        }
      },
      "isBackground": true,
      "problemMatcher": []
    }
  ]
}
`

func run(options Options) {
	folderPath := options.FolderPath

	// Check if folder path is provided
	if folderPath == "" {
		fmt.Println("Please provide a folder path using the -folder flag.")
		os.Exit(1)
	}

	// Create .vscode folder
	vscodePath := folderPath + "/.vscode"
	err := os.MkdirAll(vscodePath, 0o755)
	if err != nil {
		fmt.Printf("Error creating .vscode folder: %v\n", err)
		os.Exit(1)
	}

	// Create tasks.json file using template
	filePath := vscodePath + "/tasks.json"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating tasks.json file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Use template to generate tasks.json content
	tmpl, err := template.New("tasksJSON").Parse(tasksJSONTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, nil)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("tasks.json file successfully created at %s\n", filePath)
}
