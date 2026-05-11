package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Judisk/daily-tracker-cli/internal/app"
	"github.com/Judisk/daily-tracker-cli/internal/export"
	"github.com/Judisk/daily-tracker-cli/internal/stats"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	add := flag.Bool("add", false, "Add new record")
	statsFlag := flag.Bool("stats", false, "Show stats")
	last := flag.Int("last", 0, "Show last N days")
	exportFlag := flag.Bool("export", false, "Export to CSV")

	flag.Parse()

	if *add && *statsFlag || *add && *exportFlag || *statsFlag && *exportFlag || *add && *statsFlag && *exportFlag {
		fmt.Println("Choose only one: --add or --stats or --export")
		return
	}

	if !*add && !*statsFlag && !*exportFlag {
		fmt.Println("Usage:")
		fmt.Println("  --add    Add new record")
		fmt.Println("  --stats  Show stats")
		fmt.Println("  --export Export to CSV")
		fmt.Println("  --stats --last 7")

		return
	}
	if *statsFlag {
		stats.RunStats(*last)
	} else if *add {
		app.RunAdd(reader)
	} else {
		if err := export.ExportJsonToCsv(); err != nil {
			fmt.Println(err)
			return
		}
	}

}
