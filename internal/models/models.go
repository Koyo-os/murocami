package models

type WebhookPayload struct {
    Ref        string     `json:"ref"`
    Repository Repository `json:"repository"`
    Commits    []Commit   `json:"commits"`
}

type Repository struct {
    Name     string `json:"name"`
    HTMLURL  string `json:"html_url"`
    CloneURL string `json:"clone_url"`
}

type Commit struct {
    ID      string `json:"id"`
    Message string `json:"message"`
}

