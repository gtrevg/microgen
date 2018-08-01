// Code generated by microgen 0.9.0. DO NOT EDIT.

package transport

import generated "github.com/devimteam/microgen/examples/generated"

type (
	UppercaseRequest struct {
		StringsMap map[string]string `json:"strings_map"`
	}
	UppercaseResponse struct {
		Ans string `json:"ans"`
	}

	CountRequest struct {
		Text   string `json:"text"`
		Symbol string `json:"symbol"`
	}
	CountResponse struct {
		Count     int   `json:"count"`
		Positions []int `json:"positions"`
	}

	TestCaseRequest struct {
		Comments []*generated.Comment `json:"comments"`
	}
	TestCaseResponse struct {
		Tree map[string]int `json:"tree"`
	}

	// Formal exchange type, please do not delete.
	DummyMethodRequest struct{}
	// Formal exchange type, please do not delete.
	DummyMethodResponse struct{}
)
