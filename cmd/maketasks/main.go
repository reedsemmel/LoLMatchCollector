package main

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	Rank int
	Div  int
	Page int
}

// Estimated, conservative estimates of page counts current according to op.gg
var Tasks []Task = []Task{
	{Rank: 0, Div: 0, Page: 2},
	{Rank: 1, Div: 0, Page: 3},
	{Rank: 2, Div: 0, Page: 10},
	{Rank: 3, Div: 0, Page: 21},
	{Rank: 3, Div: 1, Page: 21},
	{Rank: 3, Div: 2, Page: 21},
	{Rank: 3, Div: 3, Page: 51},
	{Rank: 4, Div: 0, Page: 151},
	{Rank: 4, Div: 1, Page: 121},
	{Rank: 4, Div: 2, Page: 151},
	{Rank: 4, Div: 3, Page: 401},
	{Rank: 5, Div: 0, Page: 291},
	{Rank: 5, Div: 1, Page: 401},
	{Rank: 5, Div: 2, Page: 411},
	{Rank: 5, Div: 3, Page: 1001},
}

func main() {
	out := make([]Task, 0)
	// For apex tiers, get all
	for i := 0; i < 3; i++ {
		for j := 1; j <= Tasks[i].Page; j++ {
			new_task := Tasks[i]
			new_task.Page = j
			out = append(out, new_task)
		}
	}
	// For the rest, only look at every 10 pages
	for i := 3; i < len(Tasks); i++ {
		for j := 1; j <= Tasks[i].Page; j = j + 10 {
			new_task := Tasks[i]
			new_task.Page = j
			out = append(out, new_task)
		}
	}

	fmt.Println(len(out))

	dat1, _ := json.Marshal(out[0:80])
	dat2, _ := json.Marshal(out[80:160])
	dat3, _ := json.Marshal(out[160:240])
	dat4, _ := json.Marshal(out[240:320])

	fmt.Printf("%v\n\n\n", string(dat1))
	fmt.Printf("%v\n\n\n", string(dat2))
	fmt.Printf("%v\n\n\n", string(dat3))
	fmt.Printf("%v\n\n\n", string(dat4))
}
