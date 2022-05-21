package main

//https://clickup.com/api

//https://docs.github.com/en/rest/guides/basics-of-authentication
//https://github.com/google/go-github

//https://docs.github.com/en/developers/apps/building-oauth-apps/authorizing-oauth-apps
//https://github.com/settings/applications/new

func RegisterHook(hooktype string) (err error) {
	switch hooktype {
	case "in_todo": // Missing part
		// When a task is moved in progress, create branch from develop (CU-taskID_TaskTitle)
		break
	case "in_testing": //testing
		// When a task branch is merged in develop, move card to testing
		break
	case "in_prod": //complete: Might be missing
		// When develop ahead is merged in master, move card that have the same sprint number tag to complete
		break
	}
	return nil
}

func main() {
	RegisterHook("in_todo")
	RegisterHook("in_testing")
	RegisterHook("in_prod")
}
