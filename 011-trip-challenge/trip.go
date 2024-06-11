package main

import "fmt"   

type Ticket struct {
	src string
	dst string
}

type Trip struct {
	Tickets []Ticket
}

func (t *Trip) FirstCity() string {
	firstCity := ""
	dstMap := make(map[string]struct{})
	for _, v := range t.Tickets {
		dstMap[v.dst] = struct{}{}
	}

	
	for _, v := range t.Tickets {
		if _, ok := dstMap[v.src] ; !ok {
			firstCity = v.src
			break
		}
	}
	return firstCity
} 


func (t *Trip) Itinerary() []string {
	ticketsMap := make(map[string]string)
	for _, v := range t.Tickets {
		ticketsMap[v.src] = v.dst
	}

	firstCity := t.FirstCity()
	result := []string{firstCity}
	curr := firstCity
	for {
		destination, ok := ticketsMap[curr]
		if !ok {
			break
		}
		result = append(result, destination)
		curr = destination
	}
	return result
} 

func main() {
	trip := Trip{ 
		Tickets: []Ticket{
			{"SP", "MT"},
			{"RJ", "MG"},
			{"RS", "SP"},
			{"MT", "RJ"},
		},
	}

	firstCity := trip.FirstCity()
	fmt.Println(firstCity)

	itinerary := trip.Itinerary()
	fmt.Println(itinerary)
}

