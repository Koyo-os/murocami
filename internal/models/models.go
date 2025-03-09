package models

type GitHubWebhookPayload struct {
    Ref        string `json:"ref"`
    Repository struct {
        CloneURL string `json:"clone_url"`
    } `json:"repository"`
}