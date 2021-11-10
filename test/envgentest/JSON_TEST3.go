package envgentest

type JSON_TEST3 struct {
	Category struct {
		ID       int64  `json:"ID"`
		Name     string `json:"Name"`
		Products []struct {
			Name string `json:"Name"`
			ID   string `json:"_id"`
		} `json:"Products"`
		Tags []float64 `json:"Tags"`
	} `json:"Category"`
	R  []float64 `json:"_r"`
	ID int64     `json:"id"`
}
