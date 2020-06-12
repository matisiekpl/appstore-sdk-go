package appstore

type SalesReportSale struct {
	Provider              string `csv:"Provider"`
	ProviderCountry       string `csv:"Provider Country"`
	SKU                   string `csv:"SKU"`
	Developer             string `csv:"Developer"`
	Title                 string `csv:"Title"`
	Version               string `csv:"Version"`
	ProductTypeIdentifier string `csv:"Product Type Identifier"`
	Units                 string `csv:"Units"`
	DeveloperProceeds     string `csv:"Developer Proceeds"`
	BeginDate             string `csv:"Begin Date"`
	EndDate               string `csv:"End Date"`
	CustomerCurrency      string `csv:"Customer Currency"`
	CountryCode           string `csv:"Country Code"`
	CurrencyOfProceeds    string `csv:"Currency of Proceeds"`
	AppleIdentifier       string `csv:"Apple Identifier"`
	CustomerPrice         string `csv:"Customer Price"`
	PromoCode             string `csv:"Promo Code"`
	ParentIdentifier      string `csv:"Parent Identifier"`
	Subscription          string `csv:"Subscription"`
	Period                string `csv:"Period"`
	Category              string `csv:"Category"`
	CMB                   string `csv:"CMB"`
	Device                string `csv:"Device"`
	SupportedPlatforms    string `csv:"Supported Platforms"`
	ProceedsReason        string `csv:"Proceeds Reason"`
	PreservedPricing      string `csv:"Preserved Pricing"`
	Client                string `csv:"Client"`
	OrderType             string `csv:"Order Type"`
}

type SalesReportSubscription struct {
	AppName                                        string `csv:"App Name"`
	AppAppleID                                     string `csv:"App Apple ID"`
	SubscriptionName                               string `csv:"Subscription Name"`
	SubscriptionAppleID                            string `csv:"Subscription Apple ID"`
	StandardSubscriptionDuration                   string `csv:"Standard Subscription Duration"`
	PromotionalOfferName                           string `csv:"Promotional Offer Name"`
	PromotionalOfferID                             string `csv:"Promotional Offer ID"`
	CustomerPrice                                  string `csv:"Customer Price"`
	CustomerCurrency                               string `csv:"Customer Currency"`
	DeveloperProceeds                              string `csv:"Developer Proceeds"`
	ProceedsCurrency                               string `csv:"Proceeds Currency"`
	PreservedPricing                               string `csv:"Preserved Pricing"`
	ProceedsReason                                 string `csv:"Proceeds Reason"`
	Client                                         string `csv:"Client"`
	Device                                         string `csv:"Device"`
	State                                          string `csv:"State"`
	Country                                        string `csv:"Country"`
	ActiveStandardPriceSubscriptions               string `csv:"Active Standard Price Subscriptions"`
	ActiveFreeTrialIntroductoryOfferSubscriptions  string `csv:"Active Free Trial Introductory Offer Subscriptions"`
	ActivePayUpFrontIntroductoryOfferSubscriptions string `csv:"Active Pay Up Front Introductory Offer Subscriptions"`
	ActivePayAsYouGoIntroductoryOfferSubscriptions string `csv:"Active Pay As You Go Introductory Offer Subscriptions"`
	FreeTrialPromotionalOfferSubscriptions         string `csv:"Free Trial Promotional Offer Subscriptions"`
	PayUpFrontPromotionalOfferSubscriptions        string `csv:"Pay Up Front Promotional Offer Subscriptions"`
	PayAsYouGoPromotionalOfferSubscriptions        string `csv:"Pay As You Go Promotional Offer Subscriptions"`
	MarketingOptIns                                string `csv:"Marketing Opt-Ins"`
	BillingRetry                                   string `csv:"Billing Retry"`
	GracePeriod                                    string `csv:"Grace Period"`
}

type SalesReportSubscriptionEvent struct {
	EventDate                    string `csv:"Event Date"`
	Event                        string `csv:"Event"`
	AppName                      string `csv:"App Name"`
	AppAppleID                   string `csv:"App Apple ID"`
	SubscriptionName             string `csv:"Subscription Name"`
	SubscriptionAppleID          string `csv:"Subscription Apple ID"`
	SubscriptionGroupID          string `csv:"Subscription Group ID"`
	StandardSubscriptionDuration string `csv:"Standard Subscription Duration"`
	PromotionalOfferName         string `csv:"Promotional Offer Name"`
	PromotionalOfferID           string `csv:"Promotional Offer ID"`
	SubscriptionOfferType        string `csv:"Subscription Offer Type"`
	SubscriptionOfferDuration    string `csv:"Subscription Offer Duration"`
	MarketingOptIn               string `csv:"Marketing Opt-In"`
	MarketingOptInDuration       string `csv:"Marketing Opt-In Duration"`
	PreservedPricing             string `csv:"Preserved Pricing"`
	ProceedsReason               string `csv:"Proceeds Reason"`
	ConsecutivePaidPeriods       string `csv:"Consecutive Paid Periods"`
	OriginalStartDate            string `csv:"Original Start Date"`
	Client                       string `csv:"Client"`
	Device                       string `csv:"Device"`
	State                        string `csv:"State"`
	Country                      string `csv:"Country"`
	PreviousSubscriptionName     string `csv:"Previous Subscription Name"`
	PreviousSubscriptionAppleID  string `csv:"Previous Subscription Apple ID"`
	DaysBeforeCanceling          string `csv:"Days Before Canceling"`
	CancellationReason           string `csv:"Cancellation Reason"`
	DaysCanceled                 string `csv:"Days Canceled"`
	Quantity                     string `csv:"Quantity"`
}

type SalesReportSubscriber struct {
	EventDate                    string `csv:"Event Date"`
	AppName                      string `csv:"App Name"`
	AppAppleID                   string `csv:"App Apple ID"`
	SubscriptionName             string `csv:"Subscription Name"`
	SubscriptionAppleID          string `csv:"Subscription Apple ID"`
	SubscriptionGroupID          string `csv:"Subscription Group ID"`
	StandardSubscriptionDuration string `csv:"Standard Subscription Duration"`
	PromotionalOfferName         string `csv:"Promotional Offer Name"`
	PromotionalOfferID           string `csv:"Promotional Offer ID"`
	SubscriptionOfferType        string `csv:"Subscription Offer Type"`
	SubscriptionOfferDuration    string `csv:"Subscription Offer Duration"`
	MarketingOptInDuration       string `csv:"Marketing Opt-In Duration"`
	CustomerPrice                string `csv:"Customer Price"`
	CustomerCurrency             string `csv:"Customer Currency"`
	DeveloperProceeds            string `csv:"Developer Proceeds"`
	ProceedsCurrency             string `csv:"Proceeds Currency"`
	PreservedPricing             string `csv:"Preserved Pricing"`
	ProceedsReason               string `csv:"Proceeds Reason"`
	Client                       string `csv:"Client"`
	Country                      string `csv:"Country"`
	SubscriberID                 string `csv:"Subscriber ID"`
	SubscriberIDReset            string `csv:"Subscriber ID Reset"`
	Refund                       string `csv:"Refund"`
	PurchaseDate                 string `csv:"Purchase Date"`
	Units                        string `csv:"Units"`
}
