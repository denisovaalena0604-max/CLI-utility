package main

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}
type BinList struct {
	Bins []Bin
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
func NewBinList() BinList {
	return BinList{
		Bins: []Bin{},
	}
}
func main() {
	binList := NewBinList()

	
	bin1 := NewBin("001", false, "Первый")
	bin2 := NewBin("002", true, "Второй")

	
	binList.Bins = append(binList.Bins, bin1, bin2)

	
	fmt.Printf("Всего бинов: %d\n", len(binList.Bins))

	for i, bin := range binList.Bins {
		fmt.Printf("%d. ID=%s, Name=%s, Private=%t, Created=%s\n",
			i+1, bin.ID, bin.Name, bin.Private, bin.CreatedAt.Format("15:04:05"))
	}
}
