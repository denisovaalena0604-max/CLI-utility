package bins

import "time"

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}
type BinList struct {
	Bins []Bin
}