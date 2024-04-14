# Pomodoro CLI

Pomodoro CLI is a simple yet powerful command-line tool designed to help you manage your time efficiently using the Pomodoro Technique. Built in Go and leveraging the Bubble Tea library, it provides a clean, interactive interface to conduct focused work and break sessions.

## Demo
![Demo](https://github.com/vnsonvo/pomodoro-cli/assets/78811434/5c755feb-34ec-4357-84c7-22b639d1a58b)

## Features
- Customizable Work and Break Lengths: Specify the duration of each session directly through the command line.
- Interactive Text-Based UI: Utilize an interactive textarea for command input and see live session updates.
- Session Persistence: Sessions are saved in db.json, allowing you to review your productivity data over time.
- Session Review: Retrieve and display sessions from any specific date.

## Installation

To run this application, you need to have Go installed on your machine. Download [Go](https://go.dev/doc/install) if you haven't installed it yet.

1. Clone the repository:
```bash
git clone https://github.com/vnsonvo/pomodoro-cli.git
cd pomodoro-cli
```

2. Build the application:
```bash
go build -o pomodoro
```

3. Run the application:
```bash
./pomodoro
```

## Usage

The application supports various commands to manage your Pomodoro sessions:

- **Start a Work Session**:
  - s <minutes>: Start a work session for <minutes> minutes. Default is 25 minutes if no time is specified.
- **Start a Break**:
  - b <minutes>: Start a break for <minutes> minutes. Default is 5 minutes if no time is specified.
- **List Today's Completed Sessions**:
  - l: Lists all sessions completed today.
- **List Sessions on a Specific Date**:
  - l YYYY-MM-DD: Lists all sessions completed on YYYY-MM-DD.
- **Quit**:
  - q: Exit the application.

### Example Commands
```bash
s 50         # Starts a 50-minute work session
b 10         # Starts a 10-minute break
l            # Lists today's completed sessions
l 2023-09-30 # Lists sessions from September 30, 2023
q            # Quits the application
```

### Improvement

- Starting and Ending Sounds: We've added a sound that plays when you start and finish each session.
- Notifications: get a popup message when your session is over. This makes sure we don't have to keep checking the time ourselves.

### Contributing

Contributions are welcome! Please feel free to submit a pull request.

