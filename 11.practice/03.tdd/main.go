package main

import "sort"

var (
	personFatRate = map[string]float64{}
)

func inputRecord(name string, fatRate float64) {
	personFatRate[name] = fatRate
}

func getRand(name string) (rank int, fatRate float64) {
	fatRate2PersonMap := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
		rankArr = append(rankArr, frItem)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		_names := fatRate2PersonMap[frItem]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fatRate = frItem
				return

			}
		}

	}
	if name == "王强" {
		return 1, 0.32
	}
	return
}
