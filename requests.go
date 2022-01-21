package erh

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	host = "https://api.exchangerate.host"
)

//Arg contain formatted argument
type Arg struct {
	key   string
	value string
}

//Key returns key of the argument
func (a *Arg) Key() string {
	return a.key
}

//Value returns value of the argument
func (a *Arg) Value() string {
	return a.value
}

//String stringer interface
func (a *Arg) String() string {
	return fmt.Sprintf("%s=%s", a.key, a.value)
}

//ArgBase add base currency argument
func ArgBase(base string) Arg {
	return Arg{key: "base", value: base}
}

//ArgPlaces add currency places argument
func ArgPlaces(places int) Arg {
	return Arg{key: "places", value: strconv.Itoa(places)}
}

//ArgAmount add currency amount argument
func ArgAmount(amount float64) Arg {
	return Arg{key: "amount", value: strconv.FormatFloat(amount, 'f', -1, 64)}
}

//ArgSymbols add currency symbols argument
func ArgSymbols(symbols []string) Arg {
	return Arg{key: "symbols", value: strings.Join(symbols, ",")}
}

//argsURLEncoded convert slice of arguments to url encoded string
func argsURLEncoded(args []Arg) string {
	uv := url.Values{}
	for _, arg := range args {
		uv.Add(arg.Key(), arg.Value())
	}
	return uv.Encode()
}

//ConvertCtx convert currency request with context
func ConvertCtx(ctx context.Context, from, to string, amount float64, args ...Arg) (ConvertResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/convert?from=%s&to=%s&amount=%f&%s",
			host,
			from,
			to,
			amount,
			argsURLEncoded(args),
		),
		nil,
	)
	if err != nil {
		return ConvertResponse{}, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req.WithContext(ctx))
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
			argsURLEncoded(args),
		),
		nil,
	)
	if err != nil {
		return HistoricalResponse{}, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req.WithContext(ctx))
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
			argsURLEncoded(args),
		),
		nil,
	)
	if err != nil {
		return TimeSeriesResponse{}, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req.WithContext(ctx))
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

	client := http.DefaultClient
	resp, err := client.Do(req.WithContext(ctx))
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
			argsURLEncoded(args),
		),
		nil,
	)
	if err != nil {
		return LatestResponse{}, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req.WithContext(ctx))
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
