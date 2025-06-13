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

Open `http://localhost:8080` in your browser to submit and view leave requests. jQuery is loaded from the official CDN, while the custom dialog plugin is served from the `static` directory.
