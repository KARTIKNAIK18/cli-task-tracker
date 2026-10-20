# Go Task Tracker CLI

A simple command-line task tracker built with **Go** that allows users to manage their daily tasks. It supports adding, updating, deleting, listing, and completing tasks, with persistent storage using a JSON file.

## Features

- Add new tasks with unique IDs.
- List all tasks.
- Update task descriptions.
- Delete tasks by ID.
- Mark tasks as completed.
- List only completed tasks.
- Store task data in `tasks.json` so it persists between executions.
- Track task creation and update timestamps.

## Technologies Used

- **Go** — Application logic and CLI handling
- **JSON** — Data storage
- **OS package** — Command-line arguments and file operations
- **Encoding/JSON package** — JSON encoding and decoding
- **Time package** — Task timestamps


### 1. Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/task-tracker.git
cd task-tracker
```

Replace `YOUR_USERNAME` with your GitHub username.

### 2. Run the application

```bash
go run .
```

## Usage

### Add a task

```bash
go run . add "Learn Go"
```

### List all tasks

```bash
go run . list
```

### Update a task

```bash
go run . update 1 "Learn Go file handling"
```

### Delete a task

```bash
go run . delete 1
```

### Mark a task as completed

```bash
go run . done 1
```

### List completed tasks

```bash
go run . list-done
```

**Note:** Task IDs are used to identify tasks. Replace `1` in the examples with the relevant task ID.

## Data Storage

Tasks are stored in `tasks.json` in JSON format. Each task contains:

| Field | Description |
|---|---|
| `id` | Unique task identifier |
| `description` | Task description |
| `status` | Completion status (`true` or `false`) |
| `createdAt` | Task creation timestamp |
| `updatedAt` | Last update timestamp, or `null` if never updated |

## Learning Objectives

This project demonstrates practical use of:

- Go structs and slices
- Functions and pointers
- Command-line arguments using `os.Args`
- JSON encoding and decoding
- File handling and persistence
- Error handling
- Working with timestamps

## Future Improvements

- Add task priorities and due dates.
- Add filtering by pending and completed status.
- Improve command-line input validation.
- Add unit tests.

## License

This project is open source and available under the MIT License, if you choose to distribute it under that license.