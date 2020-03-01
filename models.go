package main

// BranchComparison represents the result of comparing two branches
type BranchComparison struct {
	Message  string
	Status   string `json:"status"`
	AheadBy  int    `json:"ahead_by"`
	BehindBy int    `json:"behind_by"`
}

// Repository models a github repository to get stats for
type Repository struct {
	Name               string   `json:"name"`
	UserInterfaceURL   string   `json:"userInterfaceUrl"`
	RepositoryAPIURL   string   `json:"repositoryApiUrl"`
	BaseBranch         string   `json:"baseBranch"`
	ComparisonBranches []string `json:"comparisonBranches"`
}

// Configuration holds all of the repositories and their configurations
type Configuration struct {
	Repositories []Repository `json:"repositories"`
}
