package erh

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	host = "https://api.exchangerate.host"
)

type Arg struct {
	data string
}

func (a *Arg) String() string {
	return a.data
}

func ArgBase(base string) Arg {
	return Arg{data: fmt.Sprintf("base=%s", base)}
}

func ArgPlaces(places int) Arg {
	return Arg{data: fmt.Sprintf("places=%d", places)}
}

func ArgAmount(amount float64) Arg {
	return Arg{data: fmt.Sprintf("amount=%f", amount)}
}

func ArgSymbols(symbols []string) Arg {
	return Arg{data: fmt.Sprintf("symbols=%s", strings.Join(symbols, ","))}
}

func argsToString(args []Arg) string {
	var argStrings []string
	for _, a := range args {
		argStrings = append(argStrings, a.String())
	}
	return strings.Join(argStrings, "&")
}

func ConvertCtx(ctx context.Context, from, to string, amount float64, args ...Arg) (ConvertResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/convert?from=%s&to=%s&amount=%f%s",
			host,
			from,
			to,
			amount,
			argsToString(args),
		),
		nil,
	)
	if err != nil {
		return ConvertResponse{}, err
	}
	req.WithContext(ctx)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return ConvertResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return ConvertResponse{}, errors.New(resp.Status)
	}
	var response ConvertResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return ConvertResponse{}, err
	}
	return response, nil
}

func Convert(from, to string, amount float64, args ...Arg) (ConvertResponse, error) {
	return ConvertCtx(context.Background(), from, to, amount, args...)
}

func HistoricalCtx(ctx context.Context, date time.Time, args ...Arg) (HistoricalResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/%s?%s",
			host,
			date.Format("2006-01-02"),
			argsToString(args),
		),
		nil,
	)
	if err != nil {
		return HistoricalResponse{}, err
	}
	req.WithContext(ctx)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return HistoricalResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return HistoricalResponse{}, errors.New(resp.Status)
	}

	var response HistoricalResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return HistoricalResponse{}, err
	}
	return response, nil
}

func Historical(date time.Time, args ...Arg) (HistoricalResponse, error) {
	return HistoricalCtx(context.Background(), date, args...)
}

func TimeSeriesCtx(ctx context.Context, start, end time.Time, amount float64, args ...Arg) (TimeSeriesResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/timeseries?start_date=%s&end_date=%s&%s",
			host,
			start.Format("2006-01-02"),
			end.Format("2006-01-02"),
			argsToString(args),
		),
		nil,
	)
	if err != nil {
		return TimeSeriesResponse{}, err
	}
	req.WithContext(ctx)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return TimeSeriesResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return TimeSeriesResponse{}, errors.New(resp.Status)
	}

	var response TimeSeriesResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return TimeSeriesResponse{}, err
	}
	return response, nil
}

func TimeSeries(ctx context.Context, start, end time.Time, amount float64, args ...Arg) (TimeSeriesResponse, error) {
	return TimeSeriesCtx(context.Background(), start, end, amount, args...)
}

func SymbolsCtx(ctx context.Context) (SymbolsResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/symbols",
			host,
		),
		nil,
	)
	if err != nil {
		return SymbolsResponse{}, err
	}
	req.WithContext(ctx)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return SymbolsResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return SymbolsResponse{}, errors.New(resp.Status)
	}

	var response SymbolsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return SymbolsResponse{}, err
	}
	return response, nil
}

func Symbols() (SymbolsResponse, error) {
	return SymbolsCtx(context.Background())
}
