package main

import (
	"io"
	"net/http"

	"github.com/tenhan/gostresslib/measurer"
)

func main() {
	m := measurer.NewJobMeasurer()
	total := 10000
	concurrency := 500
	m.Run(total, concurrency, []string{"response_size(byte)"}, func(num int, metric *measurer.JobMetric) error {
		resp, err := http.Get("http://127.0.0.1:8000/ping")
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		metric.SetMetricsValue(float64(len(bytes)))
		return nil
	}).Print()
}
