package main

import (
	"fmt"
	"time"
	"github.com/denisovaalena0604-max/CLI-utility/bins"
// "github.com/denisovaalena0604-max/CLI-utility/api"
// "github.com/denisovaalena0604-max/CLI-utility/file"
// "github.com/denisovaalena0604-max/CLI-utility/storage"
)



func NewBin(id string, private bool, name string) bins.Bin {
	return bins.Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
func NewBinList() bins.BinList {
	return bins.BinList{
		Bins: []bins.Bin{},
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
