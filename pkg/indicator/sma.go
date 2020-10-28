package indicator

import (
	"time"

	"github.com/c9s/bbgo/pkg/store"
	"github.com/c9s/bbgo/pkg/types"
)

type Float64Slice []float64

func (s *Float64Slice) Push(v float64) {
	*s = append(*s, v)
}

var zeroTime time.Time

type SMA struct {
	Interval types.Interval
	Window   int
	Values   Float64Slice
	EndTime  time.Time
}

func (inc *SMA) calculateAndUpdate(kLines []types.KLine) {
	if len(kLines) < inc.Window {
		return
	}

	var index = len(kLines) - 1
	var kline = kLines[index]

	if inc.EndTime != zeroTime && kline.EndTime.Before(inc.EndTime) {
		return
	}

	var recentK = kLines[index-(inc.Window-1) : index+1]
	var sma = calculateSMA(recentK)
	inc.Values.Push(sma)
	inc.EndTime = kLines[index].EndTime
}

func (inc *SMA) BindMarketDataStore(store *store.MarketDataStore) {
	store.OnKLineUpdate(func(kline types.KLine) {
		// kline guard
		if inc.Interval != kline.Interval {
			return
		}

		if kLines, ok := store.KLinesOfInterval(kline.Interval); ok {
			inc.calculateAndUpdate(kLines)
		}
	})
}

func calculateSMA(kLines []types.KLine) float64 {
	sum := 0.0
	length := len(kLines)
	for _, k := range kLines {
		sum += k.Close
	}

	avg := sum / float64(length)
	return avg
}