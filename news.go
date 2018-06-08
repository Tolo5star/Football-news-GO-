package main 
import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml"
"strings"
)
type NewsMap struct {
	key string
	Location string
}

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`   //a slice of string type
}

type news struct {
	Titles []string `xml:"url>news>title"`
	Key []string `xml:"url>news>keywords"`
	Time []string `xml:"url>news>publication_date"`
}

func main() {
	var s news
	var search string
	//news_map= make(map[string])
	resp,_ := http.Get("http://www.goal.com/en-us/google-news/page")//second thing is error (as using _ will not give the error of unsuse variable) 
	bytes ,_ := ioutil.ReadAll(resp.Body)
	/*string_body := string(bytes)
	fmt.Println(string_body)*/

	resp.Body.Close()

	
	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)
	//fmt.Println(s.Key[2])
	/*for _,i := range s.Titles { //range gives index and values at that index
		fmt.Printf("\n%s",i)

	}*/

	/*for _,i := range s.Key { //range gives index and values at that index
		fmt.Printf("\n%s",i)

	}*/
	fmt.Println("Results from Goal.com")
	fmt.Println("Enter query (first letter capital): ")
	fmt.Scanln(&search)
   

	for i:= 0 ; i<len(s.Key);i++ {
		s.Time[i]=s.Time[i][:10]
		if strings.Contains(s.Key[i],search){
			fmt.Printf("\n %d) %s  %s ",i,s.Titles[i],s.Time[i])
		}
	}
	
}