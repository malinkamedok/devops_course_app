package weather

type Stats struct {
	Average float64 `json:"average"`
	Median  float64 `json:"median"`
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
}
