package main

import "fmt"

type TopNAd struct {
	AdId         uint64
	FilterReason string
}
type PosData struct {
	TopNAd     []*TopNAd
	FilteredAd []*TopNAd
}

func (p *PosData) FilterAds(shouldFilter func(nAd *TopNAd) bool, reason string) {
	newAds := make([]*TopNAd, 0, len(p.TopNAd))

	for _, topAd := range p.TopNAd {
		if !shouldFilter(topAd) {
			newAds = append(newAds, topAd)
		}
	}

	p.AddNotChosenFilteredAds(p.TopNAd, newAds, reason)
	p.TopNAd = newAds
}
func (p *PosData) AddNotChosenFilteredAds(candidateAds []*TopNAd, chosenAds []*TopNAd, filterReason string) {
	chosenSet := make(map[uint64]struct{})
	for _, topAd := range chosenAds {
		chosenSet[topAd.AdId] = struct{}{}
	}

	for _, cAd := range candidateAds {
		if _, ok := chosenSet[cAd.AdId]; !ok {
			p.AddFilteredAd(cAd, filterReason)
		}
	}
}
func (p *PosData) AddFilteredAd(topAd *TopNAd, filterReason string) {
	topAd.FilterReason = filterReason
	p.FilteredAd = append(p.FilteredAd, topAd)
}

type MixerContext struct {
	PosData []*PosData
}

func main() {
	topNad := make([]*TopNAd, 0, 4)
	for i := 1; i < 4; i++ {
		topNad = append(topNad, &TopNAd{AdId: uint64(i), FilterReason: "ceshi"})
	}
	fmt.Printf("%v \n", topNad)
	posData := make([]*PosData, 0, 4)
	posData = append(posData, &PosData{TopNAd: topNad})
	ctx := &MixerContext{PosData: posData}
	fmt.Printf("%v \n", ctx.PosData[0].TopNAd)
	for _, pos := range ctx.PosData {
		pos.FilterAds(func(nAd *TopNAd) bool {
			return nAd.AdId == 2
		}, "111111")
	}
	fmt.Printf("%v 3333 \n", ctx.PosData[0].TopNAd)
	for _, datum := range ctx.PosData {
		fmt.Printf("%v 111 \n", datum.TopNAd)
		fmt.Printf("%v 222 \n", datum.FilteredAd)

	}
	fmt.Printf("%v \n", ctx.PosData[0].TopNAd)

}
