package main

import (
	"chaelub/thinksurance/exchange-rates/client"
	"chaelub/thinksurance/exchange-rates/client/operations"
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/strfmt"
	table "github.com/tatsushid/go-prettytable"
	"time"
)

var (
	customBaseCurrency = "USD"
	customSymbols      = []string{"EUR", "HUF"}
)

func main() {

	// Fetching latest exchange rates
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Default.Operations.GetLatest(&operations.GetLatestParams{
		Base:    &customBaseCurrency,
		Symbols: customSymbols,
		Context: ctx,
	})
	if err != nil {
		panic(err)
	}

	respMap, ok := resp.Payload.Rates.(map[string]interface{})
	if !ok {
		panic("the Latest result is not a map[string]interface{}")
	}
	fmt.Println("\nLatest exchange rates:")
	err = printResult(resp.Payload.Base, time.Time(resp.Payload.Date), respMap)
	if err != nil {
		panic(err)
	}

	// Fetching historical exchange rates
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	end := strfmt.Date(time.Now())
	start := strfmt.Date(time.Now().AddDate(0, 0, -3))
	respHistory, err := client.Default.Operations.GetHistory(&operations.GetHistoryParams{
		EndAt:   &end,
		StartAt: &start,
		Symbols: customSymbols,
		Base:    &customBaseCurrency,
		Context: ctx,
	})
	if err != nil {
		panic(err)
	}
	respMap, ok = respHistory.Payload.Rates.(map[string]interface{})
	if !ok {
		panic("the HistoryRates result is not a map[string]interface{}")
	}

	fmt.Println("\nHistorical Exchange rates:")
	err = printResult(respHistory.Payload.Base, time.Time(respHistory.Payload.Date), respMap)
	if err != nil {
		panic(err)
	}

	// Fetching exchange rates for special date
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	date := strfmt.Date(time.Now().AddDate(-1, 0, 0))
	respDate, err := client.Default.Operations.GetDate(&operations.GetDateParams{
		Date:    date,
		Context: ctx,
	})
	if err != nil {
		panic(err)
	}

	respMap, ok = respDate.Payload.Rates.(map[string]interface{})
	if !ok {
		panic("the RatesForDate result is not a map[string]interface{}")
	}

	fmt.Println("\nExchange rates for special date:")
	err = printResult(respDate.Payload.Base, time.Time(respDate.Payload.Date), respMap)
	if err != nil {
		panic(err)
	}
}

func printResult(base string, date time.Time, ratesObj map[string]interface{}) error {

	tbl, _ := table.NewTable([]table.Column{
		{
			Header: "Date",
		},
		{
			Header: "Currency",
		},
		{
			Header: "Rate",
		},
	}...)
	tbl.Separator = " | "
	if date.IsZero() {
		// got historical data
		for date, rates := range ratesObj {
			ratesMap, ok := rates.(map[string]interface{})
			if !ok {
				return errors.New("can't convert nested rates struct to map[string]interface{}")
			}
			for currency, rate := range ratesMap {
				err := tbl.AddRow(date, fmt.Sprintf("%s->%s", base, currency), rate)
				if err != nil {
					return err
				}
			}
		}
	} else {
		for currency, rate := range ratesObj {
			err := tbl.AddRow(
				date.Format("2006-01-02"),
				fmt.Sprintf("%s->%s",
					base,
					currency),
				rate)
			if err != nil {
				return err
			}
		}
	}
	_, err := tbl.Print()
	return err
}
