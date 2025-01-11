package main

import "fmt"

func describeEvent(event Event) {
	switch event.Type {
	case "PushEvent":
		commitCount := len(event.Payload.Commits)
		fmt.Printf("- Pushed %d commit(s) to %s\n", commitCount, event.Repo.Name)
	case "IssuesEvent":
		fmt.Printf("- Opened a new issue in %s\n", event.Repo.Name)
	case "WatchEvent":
		fmt.Printf("- Starred %s\n", event.Repo.Name)
	case "ForkEvent":
		fmt.Printf("- Forked %s\n", event.Repo.Name)
	case "CreateEvent":
		fmt.Printf("- Created a new repository or branch in %s\n", event.Repo.Name)
	default:
		fmt.Printf("- Performed an event of type %s in %s\n", event.Type, event.Repo.Name)
	}
}
