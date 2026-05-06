# Daily Tracker CLI

Simple CLI tool to track your daily state: mood, energy, focus — and pills.

---

## 🚀 Quick Start

git clone https://github.com/Judisk/daily-tracker-cli  
cd daily-tracker-cli  
go run ./cmd --add

Example:

$ go run ./cmd --add  
Mood (0-5): 4  
Energy (0-5): 3  
Focus (0-5): 5  
Pills (0-50): 6  
Warning pills running low (6 left)  
Saved ✅

---

## 📊 Stats

go run ./cmd --stats

Last N days:

go run ./cmd --stats --last 7

Example output:

Records used: 7  
Average mood:   3.71  
Average energy: 3.42  
Average focus:  4.00  

---

## ⚙️ Usage

### Add record

go run ./cmd --add

Or with flags:

go run ./cmd --add --mood 4 --energy 3 --focus 5 --pills 10

---

### Flags

--mood    Mood level (0–5)  
--energy  Energy level (0–5)  
--focus   Focus level (0–5)  
--pills   Pills left (0–50)  
--stats   Show statistics  
--last    Last N records  

---

## 📁 Data Storage

Data is stored locally in:

data/data.csv

Example:

Date,Mood,Energy,Focus,Pills  
2026-05-06,4,3,5,10  

---

## ✨ Features

- CLI-based daily tracking  
- Interactive input or flags  
- CSV storage  
- Average statistics  
- Last N records filtering  
- Input validation  
- Pills low warning  

---

## 🧠 Why this project?

This project was built to:

- practice Go (CLI, file IO, validation)  
- track personal daily metrics  
- experiment with simple data analysis  

---

## 🛠 Tech Stack

- Go (standard library)  
- CSV for storage  

---

## 📌 Roadmap

- Add pills consumption tracking  
- Show pills stats  
- Add edit/delete commands  
- JSON export  
- Tests  

---

## 📄 License

MIT
