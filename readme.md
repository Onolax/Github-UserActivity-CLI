# GitHub Activity Cli
This Go program fetches and displays recent events for a given GitHub user using the GitHub API. It shows the events in a human-readable format.
This is a roadmap.sh project (https://roadmap.sh/projects/github-user-activity)
## Features

- Fetches recent user events from GitHub API.
- Displays events in a user-friendly format, such as:
    - `Pushed 3 commits to repository-name`
    - `Opened a new issue in repository-name`
    - `Starred repository-name`


## Installation

1. Clone this repository:
   ```bash
   git clone <repository_url>
   ```
2. Build the binary:
    ```bash
   go build -o github-events
    ```
3. Run the app:
    ```bash
   ./github-events <username>
    ```
