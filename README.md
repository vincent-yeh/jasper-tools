# jasper-tools

This repository now includes a simple web-based leave management system written in Go.

The form collects an employee's ID, position, department, leave type, start and end dates, and the reason for the leave. Before saving, the information is previewed in a jQuery-powered dialog so you can confirm it.

## Running the leave app

1. Create a SQLite database file named `leaves.db` and create a table:

```sql
CREATE TABLE IF NOT EXISTS requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    emp_id TEXT,
    position TEXT,
    department TEXT,
    type TEXT,
    start_date TEXT,
    end_date TEXT,
    reason TEXT
);
```

2. Start the server:

```bash
cd leaveapp
go run main.go
```

Open `http://localhost:8080` in your browser to submit and view leave requests. All JavaScript (the full jQuery library and the dialog plugin) is served locally from the `static` directory. Download `jquery.min.js` from the jQuery website and place it in `leaveapp/static/`.
