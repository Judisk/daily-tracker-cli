package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Judisk/daily-tracker-cli/internal/cli"
	"github.com/Judisk/daily-tracker-cli/internal/export"
	"github.com/Judisk/daily-tracker-cli/internal/gui"
	"github.com/Judisk/daily-tracker-cli/internal/stats"
)

func main() {

	var reader = bufio.NewReader(os.Stdin)
	addFlag := flag.Bool("add", false, "Add new record")
	statsFlag := flag.Bool("stats", false, "Show stats")
	last := flag.Int("last", 0, "Show last N days")
	exportFlag := flag.Bool("export", false, "Export to CSV")
	guiFlag := flag.Bool("gui", false, "gui version")

	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Printf("Unexpected arguments: %v\n", flag.Args())
		fmt.Println("Use --help to see available flags.")
		return
	}

	activeModes := countActiveModes(*addFlag, *statsFlag, *exportFlag, *guiFlag)

	if activeModes > 1 {
		fmt.Println("Choose only one: --add, --stats, --export, or --gui")
		return
	}

	switch {

	case *statsFlag:
		stats.RunStats(*last)

	case *addFlag:
		if err := cli.Add(reader); err != nil {
			fmt.Println(err)
			return
		}

	case *exportFlag:
		if err := export.ExportJsonToCsv(); err != nil {
			fmt.Println(err)
			return
		}

	case *guiFlag:
		gui.RunGui()

	default:
		fmt.Println("Usage:")
		fmt.Println("  --add    Add new record")
		fmt.Println("  --stats  Show stats")
		fmt.Println("  --export Export to CSV")
		fmt.Println("  --stats --last 7")
		fmt.Println("  --gui    Launch GUI")

		return
	}

}

func countActiveModes(modes ...bool) int {
	count := 0
	for _, mode := range modes {
		if mode {
			count++
		}
	}
	return count
}
