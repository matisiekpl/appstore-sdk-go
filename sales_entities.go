package appstore_sdk

type SalesReportSale struct {
	Provider              string `csv:"Provider" json:"provider"`
	ProviderCountry       string `csv:"Provider Country" json:"provider_country"`
	SKU                   string `csv:"SKU" json:"sku"`
	Developer             string `csv:"Developer" json:"developer"`
	Title                 string `csv:"Title" json:"title"`
	Version               string `csv:"SalesReportVersion" json:"version"`
	ProductTypeIdentifier string `csv:"Product Type Identifier" json:"product_type_identifier"`
	Units                 string `csv:"Units" json:"units"`
	DeveloperProceeds     string `csv:"Developer Proceeds" json:"developer_proceeds"`
	BeginDate             string `csv:"Begin Date" json:"begin_date"`
	EndDate               string `csv:"End Date" json:"end_date"`
	CustomerCurrency      string `csv:"Customer Currency" json:"customer_currency"`
	CountryCode           string `csv:"Country Code" json:"country_code"`
	CurrencyOfProceeds    string `csv:"Currency of Proceeds" json:"currency_of_proceeds"`
	AppleIdentifier       string `csv:"Apple Identifier" json:"apple_identifier"`
	CustomerPrice         string `csv:"Customer Price" json:"customer_price"`
	PromoCode             string `csv:"Promo Code" json:"promo_code"`
	ParentIdentifier      string `csv:"Parent Identifier" json:"parent_identifier"`
	Subscription          string `csv:"Subscription" json:"subscription"`
	Period                string `csv:"Period" json:"period"`
	Category              string `csv:"Category" json:"category"`
	CMB                   string `csv:"CMB" json:"cmb"`
	Device                string `csv:"Device" json:"device"`
	SupportedPlatforms    string `csv:"Supported Platforms" json:"supported_platforms"`
	ProceedsReason        string `csv:"Proceeds Reason" json:"proceeds_reason"`
	PreservedPricing      string `csv:"Preserved Pricing" json:"preserved_pricing"`
	Client                string `csv:"Client" json:"client"`
	OrderType             string `csv:"Order Type" json:"order_type"`
}

type SalesReportSubscription struct {
	AppName                                        string `csv:"App Name" json:"app_name"`
	AppAppleID                                     string `csv:"App Apple ID" json:"app_apple_id"`
	SubscriptionName                               string `csv:"Subscription Name" json:"subscription_name"`
	SubscriptionAppleID                            string `csv:"Subscription Apple ID" json:"subscription_apple_id"`
	StandardSubscriptionDuration                   string `csv:"Standard Subscription Duration" json:"standard_subscription_duration"`
	PromotionalOfferName                           string `csv:"Promotional Offer Name" json:"promotional_offer_name"`
	PromotionalOfferID                             string `csv:"Promotional Offer ID" json:"promotional_offer_id"`
	CustomerPrice                                  string `csv:"Customer Price" json:"customer_price"`
	CustomerCurrency                               string `csv:"Customer Currency" json:"customer_currency"`
	DeveloperProceeds                              string `csv:"Developer Proceeds" json:"developer_proceeds"`
	ProceedsCurrency                               string `csv:"Proceeds Currency" json:"proceeds_currency"`
	PreservedPricing                               string `csv:"Preserved Pricing" json:"preserved_pricing"`
	ProceedsReason                                 string `csv:"Proceeds Reason" json:"proceeds_reason"`
	Client                                         string `csv:"Client" json:"client"`
	Device                                         string `csv:"Device" json:"device"`
	State                                          string `csv:"State" json:"state"`
	Country                                        string `csv:"Country" json:"country"`
	ActiveStandardPriceSubscriptions               string `csv:"Active Standard Price Subscriptions" json:"active_standard_price_subscriptions"`
	ActiveFreeTrialIntroductoryOfferSubscriptions  string `csv:"Active Free Trial Introductory Offer Subscriptions" json:"active_free_trial_introductory_offer_subscriptions"`
	ActivePayUpFrontIntroductoryOfferSubscriptions string `csv:"Active Pay Up Front Introductory Offer Subscriptions" json:"active_pay_up_front_introductory_offer_subscriptions"`
	ActivePayAsYouGoIntroductoryOfferSubscriptions string `csv:"Active Pay As You Go Introductory Offer Subscriptions" json:"active_pay_as_you_go_introductory_offer_subscriptions"`
	FreeTrialPromotionalOfferSubscriptions         string `csv:"Free Trial Promotional Offer Subscriptions" json:"free_trial_promotional_offer_subscriptions"`
	PayUpFrontPromotionalOfferSubscriptions        string `csv:"Pay Up Front Promotional Offer Subscriptions" json:"pay_up_front_promotional_offer_subscriptions"`
	PayAsYouGoPromotionalOfferSubscriptions        string `csv:"Pay As You Go Promotional Offer Subscriptions" json:"pay_as_you_go_promotional_offer_subscriptions"`
	MarketingOptIns                                string `csv:"Marketing Opt-Ins" json:"marketing_opt_ins"`
	BillingRetry                                   string `csv:"Billing Retry" json:"billing_retry"`
	GracePeriod                                    string `csv:"Grace Period" json:"grace_period"`
}

type SalesReportSubscriptionEvent struct {
	EventDate                    string `csv:"Event Date" json:"event_date"`
	Event                        string `csv:"Event" json:"event"`
	AppName                      string `csv:"App Name" json:"app_name"`
	AppAppleID                   string `csv:"App Apple ID" json:"app_apple_id"`
	SubscriptionName             string `csv:"Subscription Name" json:"subscription_name"`
	SubscriptionAppleID          string `csv:"Subscription Apple ID" json:"subscription_apple_id"`
	SubscriptionGroupID          string `csv:"Subscription Group ID" json:"subscription_group_id"`
	StandardSubscriptionDuration string `csv:"Standard Subscription Duration" json:"standard_subscription_duration"`
	PromotionalOfferName         string `csv:"Promotional Offer Name" json:"promotional_offer_name"`
	PromotionalOfferID           string `csv:"Promotional Offer ID" json:"promotional_offer_id"`
	SubscriptionOfferType        string `csv:"Subscription Offer Type" json:"subscription_offer_type"`
	SubscriptionOfferDuration    string `csv:"Subscription Offer Duration" json:"subscription_offer_duration"`
	MarketingOptIn               string `csv:"Marketing Opt-In" json:"marketing_opt_in"`
	MarketingOptInDuration       string `csv:"Marketing Opt-In Duration" json:"marketing_opt_in_duration"`
	PreservedPricing             string `csv:"Preserved Pricing" json:"preserved_pricing"`
	ProceedsReason               string `csv:"Proceeds Reason" json:"proceeds_reason"`
	ConsecutivePaidPeriods       string `csv:"Consecutive Paid Periods" json:"consecutive_paid_periods"`
	OriginalStartDate            string `csv:"Original Start Date" json:"original_start_date"`
	Client                       string `csv:"Client" json:"client"`
	Device                       string `csv:"Device" json:"device"`
	State                        string `csv:"State" json:"state"`
	Country                      string `csv:"Country" json:"country"`
	PreviousSubscriptionName     string `csv:"Previous Subscription Name" json:"previous_subscription_name"`
	PreviousSubscriptionAppleID  string `csv:"Previous Subscription Apple ID" json:"previous_subscription_apple_id"`
	DaysBeforeCanceling          string `csv:"Days Before Canceling" json:"days_before_canceling"`
	CancellationReason           string `csv:"Cancellation Reason" json:"cancellation_reason"`
	DaysCanceled                 string `csv:"Days Canceled" json:"days_canceled"`
	Quantity                     string `csv:"Quantity" json:"quantity"`
}

type SalesReportSubscriber struct {
	EventDate                    string `csv:"Event Date" json:"event_date"`
	AppName                      string `csv:"App Name" json:"app_name"`
	AppAppleID                   string `csv:"App Apple ID" json:"app_apple_id"`
	SubscriptionName             string `csv:"Subscription Name" json:"subscription_name"`
	SubscriptionAppleID          string `csv:"Subscription Apple ID" json:"subscription_apple_id"`
	SubscriptionGroupID          string `csv:"Subscription Group ID" json:"subscription_group_id"`
	StandardSubscriptionDuration string `csv:"Standard Subscription Duration" json:"standard_subscription_duration"`
	PromotionalOfferName         string `csv:"Promotional Offer Name" json:"promotional_offer_name"`
	PromotionalOfferID           string `csv:"Promotional Offer ID" json:"promotional_offer_id"`
	SubscriptionOfferType        string `csv:"Subscription Offer Type" json:"subscription_offer_type"`
	SubscriptionOfferDuration    string `csv:"Subscription Offer Duration" json:"subscription_offer_duration"`
	MarketingOptInDuration       string `csv:"Marketing Opt-In Duration" json:"marketing_opt_in_duration"`
	CustomerPrice                string `csv:"Customer Price" json:"customer_price"`
	CustomerCurrency             string `csv:"Customer Currency" json:"customer_currency"`
	DeveloperProceeds            string `csv:"Developer Proceeds" json:"developer_proceeds"`
	ProceedsCurrency             string `csv:"Proceeds Currency" json:"proceeds_currency"`
	PreservedPricing             string `csv:"Preserved Pricing" json:"preserved_pricing"`
	ProceedsReason               string `csv:"Proceeds Reason" json:"proceeds_reason"`
	Client                       string `csv:"Client" json:"client"`
	Country                      string `csv:"Country" json:"country"`
	SubscriberID                 string `csv:"Subscriber ID" json:"subscriber_id"`
	SubscriberIDReset            string `csv:"Subscriber ID Reset" json:"subscriber_id_reset"`
	Refund                       string `csv:"Refund" json:"refund"`
	PurchaseDate                 string `csv:"Purchase Date" json:"purchase_date"`
	Units                        string `csv:"Units" json:"units"`
}
