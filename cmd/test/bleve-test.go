package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
)

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}

	// index some data
	var article string = "Bayesian hierarchical modeling. From Wikipedia. Bayesian hierarchical modelling is a statistical model written in multiple levels (hierarchical form) that estimates the parameters of the posterior distribution using the Bayesian method.[1] The sub-models combine to form the hierarchical model, and the Bayes’ theorem is used to integrate them with the observed data, and account for all the uncertainty that is present. The result of this integration is the posterior distribution, also known as the updated probability estimate, as additional evidence on the prior distribution is acquired. Frequentist statistics, the more popular foundation of statistics, has been known to contradict Bayesian statistics due to its treatment of the parameters as a random variable, and its use of subjective information in establishing assumptions on these parameters.[2] However, Bayesians argue that relevant information regarding decision making and updating beliefs cannot be ignored and that hierarchical modeling has the potential to overrule classical methods in applications where respondents give multiple observational data. Moreover, the model has proven to be robust, with the posterior distribution less sensitive to the more flexible hierarchical priors. Hierarchical modeling is used when information is available on several different levels of observational units. The hierarchical form of analysis and organization helps in the understanding of multiparameter problems and also plays an important role in developing computational strategies."
	index.Index(article, data)

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}
