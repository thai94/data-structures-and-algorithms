package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FoodRespData struct {
	City          string `json:"city,omitempty"`
	Name          string `json:"name,omitempty"`
	EstimatedCost int32  `json:"estimated_cost,omitempty"`
}

type FoodResp struct {
	Page       int32           `json:"page,omitempty"`
	PerPage    int32           `json:"per_page,omitempty"`
	Total      int32           `json:"total,omitempty"`
	TotalPages int32           `json:"total_pages,omitempty"`
	Data       []*FoodRespData `json:"data"`
}

func getFood(url string) *FoodResp {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var f FoodResp
	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &f
}

func getRelevantFoodOutlets(city string, maxCost int32) []string {

	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/food_outlets?city=%s&page=%d", city, 1)
	resp := getFood(url)
	totalPages := resp.TotalPages

	var ret []string
	for _, data := range resp.Data {
		if data.EstimatedCost <= maxCost {
			ret = append(ret, data.Name)
		}
	}

	for i := int32(2); i <= totalPages; i++ {
		url = fmt.Sprintf("https://jsonmock.hackerrank.com/api/food_outlets?city=%s&page=%d", city, i)
		resp = getFood(url)
		for _, data := range resp.Data {
			if data.EstimatedCost <= maxCost {
				ret = append(ret, data.Name)
			}
		}
	}

	return ret
}

type Capital struct {
	Capital string `json:"capital"`
}

type CapitalResp struct {
	Data []*Capital `json:"data"`
}

func getCapitalCity(country string) string {

	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/countries?name=%s", country)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var f CapitalResp
	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if f.Data != nil && len(f.Data) == 1 {
		return f.Data[0].Capital
	}

	return "-1"
}

func main() {
	fmt.Println(getCapitalCity("Afghanistan"))
}
