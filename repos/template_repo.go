package repos

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type IntSeries struct {
	values []int
}

type FloatSeries struct {
	values []float64
}

type StringSeries struct {
	values []string
}

func generateSyntheticDataWithoutMeta(df dataframe.DataFrame, size int) dataframe.DataFrame {
	dtypes := df.Types()
	columns := df.Names()

	fmt.Println(df.Col(columns[0]))
	fmt.Println(df.Records()[1])

	var s []series.Series
	for j := 0; j < len(columns); j++ {

		switch dtypes[j] {
		case series.String:
			var strSeries StringSeries
			for i := 0; i < size; i++ {
				strSeries.values = append(strSeries.values, GetFakeString(len(df.Col(columns[j]).MinStr()), len(df.Col(columns[j]).MaxStr())))
			}
			if len(strSeries.values) > 0 {
				s = append(s, series.New(strSeries.values, series.String, columns[j]))
			}
		case series.Int:
			var intSeries IntSeries
			for i := 0; i < size; i++ {
				intSeries.values = append(intSeries.values, GetFakeInt(int(df.Col(columns[j]).Min()), int(df.Col(columns[j]).Max())))
			}
			if len(intSeries.values) > 0 {
				s = append(s, series.New(intSeries.values, series.Int, columns[j]))
			}
		case series.Float:
			var floaSeries FloatSeries
			for i := 0; i < size; i++ {
				floaSeries.values = append(floaSeries.values, GetFakeFloat(df.Col(columns[j]).Min(), df.Col(columns[j]).Max()))
			}
			if len(floaSeries.values) > 0 {
				s = append(s, series.New(floaSeries.values, series.Float, columns[j]))
			}
		}

	}

	return dataframe.New(s...)
}
