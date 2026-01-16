# GitHub User Activity CLI

A simple Command Line Interface (CLI) tool built with Go to fetch and display the recent activity of a GitHub user.

This project is a solution for the GitHub User Activity challenge from roadmap.sh.
Project URL: https://roadmap.sh/projects/github-user-activity

## Features

- Fetches recent events from the GitHub API for a specified user.
- Displays the event type and the repository name where the event occurred.
- Handles basic error cases (e.g., missing username argument).

## Prerequisites

- [Go](https://go.dev/dl/) (version 1.25 or higher recommended)

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd github-activity
   ```

2. Build the project:
   ```bash
   go build -o github-activity main.go
   ```

## Usage

Run the compiled binary with a GitHub username as the argument:

```bash
./github-activity <username>
```

### Example

```bash
./github-activity kamranahmedse
```

**Output:**
```
Event: PushEvent di Repo: kamranahmedse/developer-roadmap
Event: WatchEvent di Repo: kamranahmedse/developer-roadmap
...
```

## Project Structure

- `main.go`: Contains the main logic for fetching and parsing GitHub user events.
- `go.mod`: Go module definition.

## License

This project is open source and available under the [MIT License](LICENSE).
