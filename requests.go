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

//Arg contain formatted argument
type Arg struct {
	data string
}

func (a *Arg) String() string {
	return a.data
}

//ArgBase add base currency argument
func ArgBase(base string) Arg {
	return Arg{data: fmt.Sprintf("base=%s", base)}
}

//ArgPlaces add currency places argument
func ArgPlaces(places int) Arg {
	return Arg{data: fmt.Sprintf("places=%d", places)}
}

//ArgAmount add currency amount argument
func ArgAmount(amount float64) Arg {
	return Arg{data: fmt.Sprintf("amount=%f", amount)}
}

//ArgSymbols add currency symbols argument
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

//ConvertCtx convert currency request with context
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

//Convert with context.Background
func Convert(from, to string, amount float64, args ...Arg) (ConvertResponse, error) {
	return ConvertCtx(context.Background(), from, to, amount, args...)
}

//HistoricalCtx historical currency with context
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

//Historical with context.Background
func Historical(date time.Time, args ...Arg) (HistoricalResponse, error) {
	return HistoricalCtx(context.Background(), date, args...)
}

//TimeSeriesCtx time series request with context
func TimeSeriesCtx(ctx context.Context, start, end time.Time, args ...Arg) (TimeSeriesResponse, error) {
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

//TimeSeries with context.Background
func TimeSeries(start, end time.Time, args ...Arg) (TimeSeriesResponse, error) {
	return TimeSeriesCtx(context.Background(), start, end, args...)
}

//SymbolsCtx returns available symbols
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

//Symbols with context.Background
func Symbols() (SymbolsResponse, error) {
	return SymbolsCtx(context.Background())
}

//LatestCtx request latest rates with context
func LatestCtx(ctx context.Context, args ...Arg) (LatestResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/latest?%s",
			host,
			argsToString(args),
		),
		nil,
	)
	if err != nil {
		return LatestResponse{}, err
	}
	req.WithContext(ctx)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return LatestResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return LatestResponse{}, errors.New(resp.Status)
	}

	var response LatestResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return LatestResponse{}, err
	}
	return response, nil
}

//Latest returns latest rates with context.Background
func Latest(args ...Arg) (LatestResponse, error) {
	return LatestCtx(context.Background(), args...)
}
