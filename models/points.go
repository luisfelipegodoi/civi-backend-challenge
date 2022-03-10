package models

type Points struct {
	X        int `json:"X"`
	Y        int `json:"Y"`
	Distance int `json:"Distance"`
}

type DistanceClassify []*Points

func (d DistanceClassify) Len() int { return len(d) }

func (d DistanceClassify) Less(i, j int) bool { return d[i].Distance < d[j].Distance }

func (d DistanceClassify) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
