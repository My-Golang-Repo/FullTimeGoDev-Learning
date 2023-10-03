package types

type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuID"`
}

type OBUdata struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}
