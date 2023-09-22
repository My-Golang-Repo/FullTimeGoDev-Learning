package types

type OBUdata struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}
