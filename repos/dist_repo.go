package repos

import (
	"fmt"

	"github.com/limnrai/limnr-data/models"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func GetDistRand(dist interface{ Rand() float64 }, size int) models.Distribution {
	var fakeDistribution models.Distribution

	for i := 0; i < size; i++ {
		fakeDistribution.Observation = append(fakeDistribution.Observation, dist.Rand())
	}

	mean, std := stat.MeanStdDev(fakeDistribution.Observation, nil)
	meanErr := stat.StdErr(std, float64(len(fakeDistribution.Observation)))
	fmt.Printf("mean= %1.1f ± %0.1v\n", mean, meanErr)

	return fakeDistribution
}

func GetFakeNormalDistributionDataset(request *models.NormalDistributionRequest) (error, models.Distribution) {
	dist := distuv.Normal{
		Mu:    request.Mu,
		Sigma: request.Sigma,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeBinomialDistributionDataset(request *models.BinomialDistributionRequest) (error, models.Distribution) {
	dist := distuv.Binomial{
		N: request.N,
		P: request.P,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeStudentsTDistributionDataset(request *models.StudentsTDistributionRequest) (error, models.Distribution) {
	dist := distuv.StudentsT{
		Mu:    request.Mu,
		Sigma: request.Sigma,
		Nu:    request.Nu,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeAlphaStableDistributionDataset(request *models.AlphaStableDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.AlphaStable{
		Alpha: request.Alpha,
		Beta:  request.Beta,
		C:     request.Beta,
		Mu:    request.Mu,
		Src:   src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeBernoulliDistributionDataset(request *models.BernoulliDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.Bernoulli{
		P:   request.P,
		Src: src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeBetaDistributionDataset(request *models.BetaDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.Beta{
		Alpha: request.Alpha,
		Beta:  request.Beta,
		Src:   src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeChiDistributionDataset(request *models.ChiDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.Chi{
		K:   request.K,
		Src: src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeChiSquaredDistributionDataset(request *models.ChiSquaredDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.ChiSquared{
		K:   request.K,
		Src: src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeExponentialDistributionDataset(request *models.ExponentialDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.Exponential{
		Rate: request.Rate,
		Src:  src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeFDistributionDataset(request *models.FDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.F{
		D1:  request.D1,
		D2:  request.D2,
		Src: src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeGammaDistributionDataset(request *models.GammaDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.Gamma{
		Alpha: request.Alpha,
		Beta:  request.Beta,
		Src:   src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeGumbelRightDistributionDataset(request *models.GumbelRightDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.GumbelRight{
		Mu:   request.Mu,
		Beta: request.Beta,
		Src:  src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeInverseGammaDistributionDataset(request *models.InverseGammaDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.InverseGamma{
		Alpha: request.Alpha,
		Beta:  request.Beta,
		Src:   src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeLaplaceDistributionDataset(request *models.LaplaceDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.Laplace{
		Mu:    request.Mu,
		Scale: request.Scale,
		Src:   src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeLogNormalDistributionDataset(request *models.LogNormalDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.LogNormal{
		Mu:    request.Mu,
		Sigma: request.Sigma,
		Src:   src,
	}
	return nil, GetDistRand(dist, request.Size)
}

func GetFakeParetoDistributionDataset(request *models.ParetoDistributionRequest) (error, models.Distribution) {
	src := rand.New(rand.NewSource(1))
	dist := distuv.Pareto{
		Xm:    request.Xm,
		Alpha: request.Alpha,
		Src:   src,
	}
	return nil, GetDistRand(dist, request.Size)
}
