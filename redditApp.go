package main

// Importing packages.
import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
)

const baseURL = "http://www.reddit.com/r/"

// Each item has a lot of JSON fields. We only need four of them.
// We must use the JSON alias.
type Item struct {
  Author string `json:"author"`
  Score int     `json:"score"`
  URL string    `json:"url"`
  Title string  `json:"title"`
}

// We need to catch the single JSON response.
type response struct {
  // We have only one data field.
  Data struct {
    // But we hace a children, who has many Items.
    Children []struct {
      Content Item `json:"data"`
    } `json:"children"`
  } `json:"data"`
}

// Show the response's content.
func (result response) showContent() {
  // One block per entry.
  for _, elem := range result.Data.Children {
    //My try to make it look "friendlier".
    fmt.Printf("\u2502 Title: %s.\n", elem.Content.Title)
    fmt.Printf("\u251C\u2500 Author: %s.\n", elem.Content.Author)
    fmt.Printf("\u251C\u2500 Score: %d.\n", elem.Content.Score)
    fmt.Printf("\u251C\u2500 Available at: %s.\n\u2514\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\n", elem.Content.URL)
  }
}

// Checks if there is an error.
func isError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Generate the URL for the query.
func generateURL(search string) string {
  return baseURL + search + ".json"
}

// Make the petition to the API and get a response.
func makePetition(search string) {
	// First, we need to get the URL to the request.
	url := generateURL(search)
	// Then, we make an HTTP GET to the URL and get a response.
	resp, err := http.Get(url)
	//Check error...
	isError(err)
  //Check if the response is correct.
  if resp.StatusCode != http.StatusOK {
    log.Fatal("Something went wrong with code: ", resp.Status)
  }
	// And close when everything has ended.
	defer resp.Body.Close()
	// Create a new response struct.
	results := new(response)
  // Decode the stream from de body of the response.
	err = json.NewDecoder(resp.Body).Decode(results)
	// Check for errors.
	isError(err)
	// Show the answer.
	results.showContent()
}


func main() {
	// Get the arguments.
	searches := os.Args[1:]
	// For each one, make a request and show the results.
	for _, val := range searches {
    // Showing a header per search.
    fmt.Printf("\n\nResults for: %s:\n|\n", val)
    // Make the petition.
		makePetition(val)
	}
}
