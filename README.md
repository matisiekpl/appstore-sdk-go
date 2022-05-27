# AppStore Connect API SDK GO (Unofficial)

[![Build Status](https://app.travis-ci.com/Kachit/appstore-sdk-go.svg?branch=master)](https://app.travis-ci.com/Kachit/appstore-sdk-go)
[![Codecov](https://codecov.io/gh/Kachit/appstore-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/appstore-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/appstore-sdk-go)](https://goreportcard.com/report/github.com/kachit/appstore-sdk-go)
[![Version](https://img.shields.io/github/go-mod/go-version/Kachit/appstore-sdk-go)](https://go.dev/doc/go1.14)
[![Release](https://img.shields.io/github/v/release/Kachit/appstore-sdk-go.svg)](https://github.com/Kachit/appstore-sdk-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/appstore-sdk-go/blob/master/LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#third-party-apis) 

## Description
Unofficial Golang SDK for AppStore Connect API

## API documentation
https://developer.apple.com/documentation/appstoreconnectapi/download_finance_reports

https://developer.apple.com/documentation/appstoreconnectapi/download_sales_and_trends_reports

## Download
```shell
go get -u github.com/kachit/appstore-sdk-go
```

## Usage
```go
package main

import (
    "fmt"
    "time"
    appstore_sdk "github.com/kachit/appstore-sdk-go/v1"
)

func main(){
    cfg := appstore_sdk.NewConfig("Issuer Id", "Key Id", "Vendor No", "path/to/your/private.key")
    client := appstore_sdk.NewClientFromConfig(cfg, nil)
    
    //Build auth token
    err := client.Init()
    if err != nil {
        fmt.Printf("Wrong API client init " + err.Error())
        panic(err)
    }
}
```

### Get sales reports
```go
ctx := context.Background()
date, _ := time.Parse("2006-01-02", "2020-05-05")
filter := appstore_sdk.NewSalesReportsFilter()
filter.SubTypeSummary().Version10().Daily().SetReportDate(date)

result, resp, err := client.SalesReports().GetSalesReports(ctx, filter)
if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(resp)

//Dump result
fmt.Println(result.Data[0].Provider)
fmt.Println(result.Data[0].ProviderCountry)
fmt.Println(result.Data[0].SKU)
fmt.Println(result.Data[0].Developer)
fmt.Println(result.Data[0].Title)
fmt.Println(result.Data[0].Version)
fmt.Println(result.Data[0].ProductTypeIdentifier)
fmt.Println(result.Data[0].Units.Value())
fmt.Println(result.Data[0].AppleIdentifier.Value())
fmt.Println(result.Data[0].DeveloperProceeds.Value())
fmt.Println(result.Data[0].BeginDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].CustomerCurrency)
fmt.Println(result.Data[0].CountryCode)
fmt.Println(result.Data[0].CurrencyOfProceeds)
fmt.Println(result.Data[0].AppleIdentifier.Value())
fmt.Println(result.Data[0].CustomerPrice.Value())
fmt.Println(result.Data[0].PromoCode)
fmt.Println(result.Data[0].ParentIdentifier)
fmt.Println(result.Data[0].Subscription)
fmt.Println(result.Data[0].Period)
fmt.Println(result.Data[0].Category)
fmt.Println(result.Data[0].CMB)
fmt.Println(result.Data[0].Device)
fmt.Println(result.Data[0].SupportedPlatforms)
fmt.Println(result.Data[0].ProceedsReason)
fmt.Println(result.Data[0].PreservedPricing)
fmt.Println(result.Data[0].Client)
fmt.Println(result.Data[0].OrderType)
```

### Get subscriptions reports
```go
ctx := context.Background()
date, _ := time.Parse("2006-01-02", "2020-05-05")
filter := appstore_sdk.NewSubscriptionsReportsFilter()
filter.SubTypeSummary().Version12().Daily().SetReportDate(date)

result, resp, err := client.SalesReports().GetSubscriptionsReports(ctx, filter)
if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(resp)

//Dump result
fmt.Println(result.Data[0].AppName)
fmt.Println(result.Data[0].AppAppleID.Value())
fmt.Println(result.Data[0].SubscriptionName)
fmt.Println(result.Data[0].SubscriptionAppleID.Value())
fmt.Println(result.Data[0].SubscriptionGroupID.Value())
fmt.Println(result.Data[0].StandardSubscriptionDuration)
fmt.Println(result.Data[0].PromotionalOfferName)
fmt.Println(result.Data[0].PromotionalOfferID)
fmt.Println(result.Data[0].CustomerPrice.Value())
fmt.Println(result.Data[0].CustomerCurrency)
fmt.Println(result.Data[0].DeveloperProceeds.Value())
fmt.Println(result.Data[0].ProceedsCurrency)
fmt.Println(result.Data[0].PreservedPricing)
fmt.Println(result.Data[0].ProceedsReason)
fmt.Println(result.Data[0].Client)
fmt.Println(result.Data[0].Device)
fmt.Println(result.Data[0].State)
fmt.Println(result.Data[0].Country)
fmt.Println(result.Data[0].ActiveStandardPriceSubscriptions.Value())
fmt.Println(result.Data[0].ActiveFreeTrialIntroductoryOfferSubscriptions.Value())
fmt.Println(result.Data[0].ActivePayUpFrontIntroductoryOfferSubscriptions.Value())
fmt.Println(result.Data[0].ActivePayAsYouGoIntroductoryOfferSubscriptions.Value())
fmt.Println(result.Data[0].FreeTrialPromotionalOfferSubscriptions.Value())
fmt.Println(result.Data[0].PayUpFrontPromotionalOfferSubscriptions.Value())
fmt.Println(result.Data[0].PayAsYouGoPromotionalOfferSubscriptions.Value())
fmt.Println(result.Data[0].MarketingOptIns.Value())
fmt.Println(result.Data[0].BillingRetry.Value())
fmt.Println(result.Data[0].GracePeriod.Value())
```

### Get subscriptions events reports
```go
ctx := context.Background()
date, _ := time.Parse("2006-01-02", "2020-05-05")
filter := appstore_sdk.NewSubscriptionsEventsReportsFilter()
filter.SubTypeSummary().Version12().Daily().SetReportDate(date)

result, resp, err := client.SalesReports().GetSubscriptionsEventsReports(ctx, filter)
if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(resp)

//Dump result
fmt.Println(result.Data[0].EventDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].Event)
fmt.Println(result.Data[0].AppName)
fmt.Println(result.Data[0].AppAppleID.Value())
fmt.Println(result.Data[0].SubscriptionName)
fmt.Println(result.Data[0].SubscriptionAppleID.Value())
fmt.Println(result.Data[0].SubscriptionGroupID.Value())
fmt.Println(result.Data[0].StandardSubscriptionDuration)
fmt.Println(result.Data[0].PromotionalOfferName)
fmt.Println(result.Data[0].PromotionalOfferID)
fmt.Println(result.Data[0].SubscriptionOfferType)
fmt.Println(result.Data[0].SubscriptionOfferDuration)
fmt.Println(result.Data[0].MarketingOptIn)
fmt.Println(result.Data[0].MarketingOptInDuration)
fmt.Println(result.Data[0].PreservedPricing)
fmt.Println(result.Data[0].ProceedsReason)
fmt.Println(result.Data[0].ConsecutivePaidPeriods.Value())
fmt.Println(result.Data[0].OriginalStartDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].Client)
fmt.Println(result.Data[0].Device)
fmt.Println(result.Data[0].State)
fmt.Println(result.Data[0].Country)
fmt.Println(result.Data[0].PreviousSubscriptionName)
fmt.Println(result.Data[0].PreviousSubscriptionAppleID.Value())
fmt.Println(result.Data[0].DaysBeforeCanceling.Value())
fmt.Println(result.Data[0].CancellationReason)
fmt.Println(result.Data[0].DaysCanceled.Value())
fmt.Println(result.Data[0].Quantity.Value())
```

### Get subscribers reports
```go
ctx := context.Background()
date, _ := time.Parse("2006-01-02", "2020-05-05")
filter := appstore_sdk.NewSubscribersReportsFilter()
filter.SubTypeDetailed().Version12().Daily().SetReportDate(date)

result, resp, err := client.SalesReports().GetSubscribersReports(ctx, filter)
if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(resp)

//Dump result
fmt.Println(result.Data[0].EventDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].AppName)
fmt.Println(result.Data[0].AppAppleID.Value())
fmt.Println(result.Data[0].SubscriptionName)
fmt.Println(result.Data[0].SubscriptionAppleID.Value())
fmt.Println(result.Data[0].SubscriptionGroupID.Value())
fmt.Println(result.Data[0].StandardSubscriptionDuration)
fmt.Println(result.Data[0].PromotionalOfferName)
fmt.Println(result.Data[0].PromotionalOfferID)
fmt.Println(result.Data[0].SubscriptionOfferType)
fmt.Println(result.Data[0].SubscriptionOfferDuration)
fmt.Println(result.Data[0].MarketingOptInDuration)
fmt.Println(result.Data[0].CustomerPrice.Value())
fmt.Println(result.Data[0].CustomerCurrency)
fmt.Println(result.Data[0].DeveloperProceeds.Value())
fmt.Println(result.Data[0].ProceedsCurrency)
fmt.Println(result.Data[0].PreservedPricing)
fmt.Println(result.Data[0].ProceedsReason)
fmt.Println(result.Data[0].Client)
fmt.Println(result.Data[0].Country)
fmt.Println(result.Data[0].SubscriberID.Value())
fmt.Println(result.Data[0].SubscriberIDReset)
fmt.Println(result.Data[0].Refund)
fmt.Println(result.Data[0].PurchaseDate.Value())
fmt.Println(result.Data[0].Units.Value())
```

### Get preorders reports
```go
ctx := context.Background()
date, _ := time.Parse("2006-01-02", "2020-05-05")
filter := appstore_sdk.NewPreOrdersReportsFilter()
filter.SubTypeSummary().Version10().Daily().SetReportDate(date)

result, resp, err := client.SalesReports().GetPreOrdersReports(ctx, filter)
if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(resp)

//Dump result
fmt.Println(result.Data[0].Provider)
fmt.Println(result.Data[0].ProviderCountry)
fmt.Println(result.Data[0].Title)
fmt.Println(result.Data[0].SKU)
fmt.Println(result.Data[0].Developer)
fmt.Println(result.Data[0].PreOrderStartDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].PreOrderEndDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].Ordered.Value())
fmt.Println(result.Data[0].Canceled.Value())
fmt.Println(result.Data[0].CumulativeOrdered.Value())
fmt.Println(result.Data[0].CumulativeCanceled.Value())
fmt.Println(result.Data[0].StartDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].CountryCode)
fmt.Println(result.Data[0].AppleIdentifier.Value())
fmt.Println(result.Data[0].Category)
fmt.Println(result.Data[0].Device)
fmt.Println(result.Data[0].SupportedPlatforms)
fmt.Println(result.Data[0].Client)
fmt.Println(result.Data[0].ProviderCountry)
```

### Get financial reports
```go
ctx := context.Background()
date, _ := time.Parse("2006-01-02", "2020-05-05")
filter :=  appstore_sdk.NewFinancesReportsFilter()
filter.SetReportDate(date).SetRegionCode("US")

result, resp, err := client.FinancesReports().GetFinancialReports(ctx, filter)
if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(resp)

//Dump result
fmt.Println(result.Data[0].StartDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
fmt.Println(result.Data[0].UPC)
fmt.Println(result.Data[0].ISRCIsbn)
fmt.Println(result.Data[0].VendorIdentifier)
fmt.Println(result.Data[0].Quantity.Value())
fmt.Println(result.Data[0].PartnerShare.Value())
fmt.Println(result.Data[0].ExtendedPartnerShare.Value())
fmt.Println(result.Data[0].PartnerShareCurrency)
fmt.Println(result.Data[0].SaleOrReturn)
fmt.Println(result.Data[0].AppleIdentifier.Value())
fmt.Println(result.Data[0].ArtistShowDeveloperAuthor)
fmt.Println(result.Data[0].Title)
fmt.Println(result.Data[0].LabelStudioNetworkDeveloperPublisher)
fmt.Println(result.Data[0].Grid)
fmt.Println(result.Data[0].ProductTypeIdentifier)
fmt.Println(result.Data[0].ISANOtherIdentifier)
fmt.Println(result.Data[0].CountryOfSale)
fmt.Println(result.Data[0].PreOrderFlag)
fmt.Println(result.Data[0].PromoCode)
fmt.Println(result.Data[0].CustomerPrice.Value())
fmt.Println(result.Data[0].CustomerCurrency)
```