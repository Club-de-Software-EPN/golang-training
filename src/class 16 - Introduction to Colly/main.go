package main

import(
	"github.com/gocolly/colly"
	"github.com/alejandrollanganate/golang-training/csv"
)

const url = "deceac-el.espe.edu.ec"
const objective = "/personal-docente-e-investigadores/"

func main(){
	// manages the network communication
	collector := colly.NewCollector(
		// the domain used by the Collector
		colly.AllowedDomains(url),
	)

	var data [][]string
	data = append(data, [][]string{{"Professor", "Email", "UrlPhotoProfie"}}...)

	collector.OnHTML(".elementor-widget-wrap", func(htmlElement *colly.HTMLElement){
		professorName := htmlElement.ChildText("h5")
		email := htmlElement.ChildText("span[style='color: #008000; font-family: Agency FB,serif;']")
		photo := htmlElement.ChildAttrs("img", "src")
		if professorName != "" {
			if email == "" {
				email = "empty"
			}
			data = append(data, [][]string{{professorName,email, photo[0]}}...)
		}
	})

	collector.Visit("https://"+url+objective)
	csv.SaveDataOnCSVFile(data, "dataESPE")
}