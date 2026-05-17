# Daily Tracker

A health and productivity tracker written in Go.

The project includes:

* CLI interface,
* GUI interface built with Fyne,
* sleep tracking,
* mood and focus monitoring,
* medication tracking,
* statistics,
* CSV export,
* validation system.

---

# 🚀 Quick Start

Clone the repository:

```bash
git clone https://github.com/Judisk/daily-tracker-cli
cd daily-tracker-cli
```

Run CLI mode:

```bash
go run ./cmd --add
```

Run GUI mode:

```bash
go run ./cmd --gui
```

Show statistics:

```bash
go run ./cmd --stats
```

Export data to CSV:

```bash
go run ./cmd --export
```

---

# 🖥 GUI

The project includes a desktop GUI built with Fyne.

Current GUI features:

* record creation,
* real-time validation,
* CSV export,
* reusable field system,
* configurable form architecture.

---

# 📝 CLI Example

```text
$ go run ./cmd --add

Went to bed -> 23:10
Fell asleep -> 23:40
Woke up -> 07:30
Sleep quality 0-5 -> 4

Mood 0-5 -> 3
Energy 0-5 -> 2
Focus 0-5 -> 4

Took meds -> 08:00
Pills 0-50 -> 12

Notes -> felt productive today
Side Effects -> none

Saved ✅
```

---

# 📊 Statistics

```bash
go run ./cmd --stats
```

Last N records:

```bash
go run ./cmd --stats --last 7
```

Example output:

```text
Records used: 7
Average sleep quality: 4.14
Average mood:          3.71
Average energy:        3.42
Average focus:         4.00
```

---

# 📤 Export

Export all saved records into CSV format:

```bash
go run ./cmd --export
```

Generated file:

```text
data/data.csv
```

---

# ⚙️ Available Commands

## Add a record (CLI)

```bash
go run ./cmd --add
```

## Launch GUI

```bash
go run ./cmd --gui
```

## Show statistics

```bash
go run ./cmd --stats
```

## Show statistics for last N records

```bash
go run ./cmd --stats --last 7
```

## Export JSON data to CSV

```bash
go run ./cmd --export
```

---

# 📁 Data Storage

Data is stored locally in JSON format:

```text
data/data.json
```

CSV export:

```text
data/data.csv
```

---

# ✨ Features

* CLI application
* Desktop GUI application
* Reusable validation system
* Reusable field configuration system
* Sleep tracking
* Sleep duration calculation
* Mood / energy / focus tracking
* Medication tracking
* JSON storage
* CSV export
* Statistics calculation
* Last N records filtering
* Real-time GUI validation
* Modular package structure
* Unit tests

---

# 🧠 Why this project?

This project was built to:

* practice Go,
* learn application architecture,
* learn GUI development in Go,
* work with JSON and CSV,
* practice validation systems,
* experiment with reusable abstractions,
* improve software design skills.

---

# 🛠 Tech Stack

* Go
* Fyne GUI toolkit
* JSON storage
* CSV export
* Go standard library

---

# 📌 Roadmap

* Record editing
* Record deletion
* Better sleep validation
* Charts and visual statistics
* SQLite support
* Configurable reminders
* Theme support
* Better GUI layout
* Notifications

---

# 📄 License

MIT
