package models

type Distribution struct {
	Observation []float64 `json:"observation"`
}

type BaseDistributionRequest struct {
	Size int `json:"size"`
}

type NormalDistributionRequest struct {
	BaseDistributionRequest
	Mu    float64 `json:"mu"`
	Sigma float64 `json:"sigma"`
}

type BinomialDistributionRequest struct {
	BaseDistributionRequest
	N float64 `json:"n"`
	P float64 `json:"p"`
}

type StudentsTDistributionRequest struct {
	BaseDistributionRequest
	Mu    float64 `json:"mu"`
	Sigma float64 `json:"sigma"`
	Nu    float64 `json:"nu"`
}

type AlphaStableDistributionRequest struct {
	BaseDistributionRequest
	Alpha float64 `json:"alpha"`
	Beta  float64 `json:"beta"`
	C     float64 `json:"c"`
	Mu    float64 `json:"mu"`
}

type BernoulliDistributionRequest struct {
	BaseDistributionRequest
	P float64 `json:"p"`
}

type BetaDistributionRequest struct {
	BaseDistributionRequest
	Alpha float64 `json:"alpha"`
	Beta  float64 `json:"beta"`
}

type ChiDistributionRequest struct {
	BaseDistributionRequest
	K float64 `json:"k"`
}

type ChiSquaredDistributionRequest struct {
	BaseDistributionRequest
	K float64 `json:"k"`
}

type ExponentialDistributionRequest struct {
	BaseDistributionRequest
	Rate float64 `json:"rate"`
}

type FDistributionRequest struct {
	BaseDistributionRequest
	D1 float64 `json:"d1"`
	D2 float64 `json:"d2"`
}

type GammaDistributionRequest struct {
	BaseDistributionRequest
	Alpha float64 `json:"alpha"`
	Beta  float64 `json:"beta"`
}

type GumbelRightDistributionRequest struct {
	BaseDistributionRequest
	Beta float64 `json:"beta"`
	Mu   float64 `json:"mu"`
}

type InverseGammaDistributionRequest struct {
	BaseDistributionRequest
	Alpha float64 `json:"alpha"`
	Beta  float64 `json:"beta"`
}

type LaplaceDistributionRequest struct {
	BaseDistributionRequest
	Mu    float64 `json:"mu"`
	Scale float64 `json:"scale"`
}

type LogNormalDistributionRequest struct {
	BaseDistributionRequest
	Mu    float64 `json:"mu"`
	Sigma float64 `json:"sigma"`
}

type ParetoDistributionRequest struct {
	BaseDistributionRequest
	Xm    float64 `json:"xm"`
	Alpha float64 `json:"alpha"`
}
