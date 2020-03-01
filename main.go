package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var configuration Configuration = LoadConfiguration("config.json")

	for i := 0; i < len(configuration.Repositories); i++ {
		var repository Repository = configuration.Repositories[i]
		name := repository.Name
		apiURL := repository.RepositoryAPIURL
		baseBranch := repository.BaseBranch
		comparisonBranches := repository.ComparisonBranches

		for j := 0; j < len(comparisonBranches); j++ {
			comparisonBranch := comparisonBranches[j]
			result := compareBranchToBase(apiURL, baseBranch, comparisonBranch)
			result.Message = fmt.Sprintf("[%s] Comparing %s to %s", name, comparisonBranch, baseBranch)
			prettyPrintJSON(result)
		}
	}
}

func compareBranchToBase(apiURL string, baseBranch string, comparisonBranch string) BranchComparison {
	requestURL := fmt.Sprintf("%s/compare/%s...%s", apiURL, baseBranch, comparisonBranch)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var comparison BranchComparison

	err = json.NewDecoder(resp.Body).Decode(&comparison)
	if err != nil {
		panic(err)
	}

	return comparison
}

func prettyPrintJSON(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	fmt.Printf("%s\n", string(prettyJSON))
}
