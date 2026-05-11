# Daily Tracker CLI

A terminal-based health and productivity tracker written in Go.

The project is focused on:

* sleep tracking,
* mood and focus monitoring,
* medication tracking,
* statistics,
* data export,
* input validation.

---

# 🚀 Quick Start

Clone the repository:

```bash
git clone https://github.com/Judisk/daily-tracker-cli
cd daily-tracker-cli
```

Add a new record:

```bash
go run ./cmd --add
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

# 📝 Example

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

## Add a record

```bash
go run ./cmd --add
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

* Interactive CLI input
* Generic reusable input system
* Sleep tracking
* Sleep duration calculation
* Mood / energy / focus tracking
* Medication tracking
* Notes and side effects logging
* JSON storage
* CSV export
* Average statistics
* Last N records filtering
* Input validation
* Generic statistics functions
* Unit tests

---

# 🧠 Why this project?

This project was built to:

* practice Go,
* learn CLI application architecture,
* work with JSON and CSV,
* practice validation and generic functions,
* experiment with data tracking and analysis.

---

# 🛠 Tech Stack

* Go (standard library)
* JSON for storage
* CSV export

---

# 📌 Roadmap

* GUI version
* Record editing
* Record deletion
* Better sleep validation
* Charts and visual statistics
* SQLite support
* Configurable reminders

---

# 📄 License

MIT
