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
	baseUrl = "https://api.exchangerate.host"
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

//Client is API client
type Client struct {
	client *http.Client
}

//NewClient create new API client with default http.Client
func NewClient() *Client {
	return &Client{client: http.DefaultClient}
}

//SetHttpClient set custom http client
func (c *Client) SetHttpClient(client *http.Client) {
	if client == nil {
		return
	}
	c.client = client
}

//ConvertCtx convert currency request with context
func (c *Client) ConvertCtx(ctx context.Context, from, to string, amount float64, args ...Arg) (ConvertResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/convert?from=%s&to=%s&amount=%f&%s",
			baseUrl,
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

	resp, err := c.client.Do(req.WithContext(ctx))
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
func (c *Client) Convert(from, to string, amount float64, args ...Arg) (ConvertResponse, error) {
	return c.ConvertCtx(context.Background(), from, to, amount, args...)
}

//HistoricalCtx historical currency with context
func (c *Client) HistoricalCtx(ctx context.Context, date time.Time, args ...Arg) (HistoricalResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/%s?%s",
			baseUrl,
			date.Format("2006-01-02"),
			argsURLEncoded(args),
		),
		nil,
	)
	if err != nil {
		return HistoricalResponse{}, err
	}

	resp, err := c.client.Do(req.WithContext(ctx))
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
func (c *Client) Historical(date time.Time, args ...Arg) (HistoricalResponse, error) {
	return c.HistoricalCtx(context.Background(), date, args...)
}

//TimeSeriesCtx time series request with context
func (c *Client) TimeSeriesCtx(ctx context.Context, start, end time.Time, args ...Arg) (TimeSeriesResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/timeseries?start_date=%s&end_date=%s&%s",
			baseUrl,
			start.Format("2006-01-02"),
			end.Format("2006-01-02"),
			argsURLEncoded(args),
		),
		nil,
	)
	if err != nil {
		return TimeSeriesResponse{}, err
	}

	resp, err := c.client.Do(req.WithContext(ctx))
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
func (c *Client) TimeSeries(start, end time.Time, args ...Arg) (TimeSeriesResponse, error) {
	return c.TimeSeriesCtx(context.Background(), start, end, args...)
}

//SymbolsCtx returns available symbols
func (c *Client) SymbolsCtx(ctx context.Context) (SymbolsResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/symbols",
			baseUrl,
		),
		nil,
	)
	if err != nil {
		return SymbolsResponse{}, err
	}

	resp, err := c.client.Do(req.WithContext(ctx))
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
func (c *Client) Symbols() (SymbolsResponse, error) {
	return c.SymbolsCtx(context.Background())
}

//LatestCtx request latest rates with context
func (c *Client) LatestCtx(ctx context.Context, args ...Arg) (LatestResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/latest?%s",
			baseUrl,
			argsURLEncoded(args),
		),
		nil,
	)
	if err != nil {
		return LatestResponse{}, err
	}

	resp, err := c.client.Do(req.WithContext(ctx))
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
func (c *Client) Latest(args ...Arg) (LatestResponse, error) {
	return c.LatestCtx(context.Background(), args...)
}
