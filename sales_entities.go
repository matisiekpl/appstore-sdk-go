package appstore_sdk

//Aggregated sales and download data for your apps and In-App Purchases
type SalesReportSale struct {
	Provider              string        `csv:"Provider" json:"provider"`                               //The service provider in your reports (typically Apple).
	ProviderCountry       string        `csv:"Provider Country" json:"provider_country"`               //The service provider country code (typically U.S.).
	SKU                   string        `csv:"SKU" json:"sku"`                                         //A product identifier provided by you during app setup.
	Developer             string        `csv:"Developer" json:"developer"`                             //Provided by you during the initial account setup.
	Title                 string        `csv:"Title" json:"title"`                                     //Provided by you during app setup.
	Version               string        `csv:"SalesReportVersion" json:"version"`                      //Provided by you during app setup.
	ProductTypeIdentifier string        `csv:"Product Type Identifier" json:"product_type_identifier"` //Defines the type of transaction (for example, initial download, update, and so on). For more information, see Product Type Identifiers.
	Units                 CustomInteger `csv:"Units" json:"units"`                                     //The aggregated number of units. Negative values indicate refunds, or CMB credits for previously purchased apps when CMB column shows ‘CMB-C’.
	DeveloperProceeds     CustomFloat64 `csv:"Developer Proceeds" json:"developer_proceeds"`           //The amount you receive per unit. This is the Customer Price minus applicable taxes and Apple’s commission, per Schedule 2 of your Paid Applications agreement.
	BeginDate             CustomDate    `csv:"Begin Date" json:"begin_date"`                           //Start date of report.
	EndDate               CustomDate    `csv:"End Date" json:"end_date"`                               //End date of report.
	CustomerCurrency      string        `csv:"Customer Currency" json:"customer_currency"`             //Three-character ISO code indicating the customer’s currency. For more information, see Currency codes.
	CountryCode           string        `csv:"Country Code" json:"country_code"`                       //Two-character ISO country code indicating the App Store territory for the purchase. For more information, see Financial Report Regions and Currencies.
	CurrencyOfProceeds    string        `csv:"Currency of Proceeds" json:"currency_of_proceeds"`       //The currency in which your proceeds are earned. For more information, see Currency codes.
	AppleIdentifier       CustomInteger `csv:"Apple Identifier" json:"apple_identifier"`               //The Apple ID for your app.
	CustomerPrice         CustomFloat64 `csv:"Customer Price" json:"customer_price"`                   //The price per unit billed to the customer, which you set for your app or in-app purchase in App Store Connect. *Customer price is inclusive of any applicable taxes we collect and remit per Schedule 2 of the Paid Applications agreement. Negative values indicate refunds, or CMB credits for previously purchased apps when CMB column shows ‘CMB-C’.
	PromoCode             string        `csv:"Promo Code" json:"promo_code"`                           //If the transaction was part of a promotion, this field will contain a value. This field is empty for all non-promotional items. For more information, see Promotional Codes.
	ParentIdentifier      string        `csv:"Parent Identifier" json:"parent_identifier"`             //In-App Purchases will show the SKU of the associated app.
	Subscription          string        `csv:"Subscription" json:"subscription"`                       //Defines whether an auto-renewable subscription is new or a renewal.
	Period                string        `csv:"Period" json:"period"`                                   //Defines the duration of an auto-renewable subscription purchase. Values include: 7 days, 1 month, 2 months, 3 months, 6 months, and 1 year.
	Category              string        `csv:"Category" json:"category"`                               //Indicates the category of the app, such as Games.
	CMB                   string        `csv:"CMB" json:"cmb"`                                         //If the transaction involves a “completed” app bundle, this field will contain a value of “CMB.” App credits for completed bundles will show a value of “CMB-C.” Otherwise this field is blank.
	Device                string        `csv:"Device" json:"device"`                                   //List of platforms that your app supports: iOS, tvOS, iOS and tvOS, or macOS.
	SupportedPlatforms    string        `csv:"Supported Platforms" json:"supported_platforms"`         // Type of device used for purchase or redownload: iPhone, iPad, Apple TV, iPod touch, or Desktop.
	ProceedsReason        string        `csv:"Proceeds Reason" json:"proceeds_reason"`                 //For Renew events, if the price is preserved then this field equals “Yes”. Otherwise it is blank.
	PreservedPricing      string        `csv:"Preserved Pricing" json:"preserved_pricing"`             //If a subscription has been active for more than a year then you receive 85% of the customer price, minus applicable taxes, and this field equals “Rate After One Year." Otherwise, you receive 70% and the field is blank.
	Client                string        `csv:"Client" json:"client"`                                   //Indicates where the purchase happened: App Store for iMessage, News, or blank.
	OrderType             string        `csv:"Order Type" json:"order_type"`                           //For introductory offers or subscription offers, indicates what type of transaction this line item is: Pay Up Front or Pay As You Go. For pre-orders, indicates whether a purchase originated from a Pre-Order. For promotional offers, the field will populate the Order ID.
}

//Total number of Active Subscriptions, Subscriptions with Introductory Prices, and Marketing Opt-Ins for your auto-renewable subscriptions.
type SalesReportSubscription struct {
	AppName                                        string        `csv:"App Name" json:"app_name"`
	AppAppleID                                     CustomInteger `csv:"App Apple ID" json:"app_apple_id"`
	SubscriptionName                               string        `csv:"Subscription Name" json:"subscription_name"`
	SubscriptionAppleID                            CustomInteger `csv:"Subscription Apple ID" json:"subscription_apple_id"`
	SubscriptionGroupID                            CustomInteger `csv:"Subscription Group ID" json:"subscription_group_id"`
	StandardSubscriptionDuration                   string        `csv:"Standard Subscription Duration" json:"standard_subscription_duration"`
	PromotionalOfferName                           string        `csv:"Promotional Offer Name" json:"promotional_offer_name"`
	PromotionalOfferID                             string        `csv:"Promotional Offer ID" json:"promotional_offer_id"`
	CustomerPrice                                  CustomFloat64 `csv:"Customer Price" json:"customer_price"`
	CustomerCurrency                               string        `csv:"Customer Currency" json:"customer_currency"`
	DeveloperProceeds                              CustomFloat64 `csv:"Developer Proceeds" json:"developer_proceeds"`
	ProceedsCurrency                               string        `csv:"Proceeds Currency" json:"proceeds_currency"`
	PreservedPricing                               string        `csv:"Preserved Pricing" json:"preserved_pricing"`
	ProceedsReason                                 string        `csv:"Proceeds Reason" json:"proceeds_reason"`
	Client                                         string        `csv:"Client" json:"client"`
	Device                                         string        `csv:"Device" json:"device"`
	State                                          string        `csv:"State" json:"state"`
	Country                                        string        `csv:"Country" json:"country"`
	ActiveStandardPriceSubscriptions               CustomInteger `csv:"Active Standard Price Subscriptions" json:"active_standard_price_subscriptions"`
	ActiveFreeTrialIntroductoryOfferSubscriptions  CustomInteger `csv:"Active Free Trial Introductory Offer Subscriptions" json:"active_free_trial_introductory_offer_subscriptions"`
	ActivePayUpFrontIntroductoryOfferSubscriptions CustomInteger `csv:"Active Pay Up Front Introductory Offer Subscriptions" json:"active_pay_up_front_introductory_offer_subscriptions"`
	ActivePayAsYouGoIntroductoryOfferSubscriptions CustomInteger `csv:"Active Pay As You Go Introductory Offer Subscriptions" json:"active_pay_as_you_go_introductory_offer_subscriptions"`
	FreeTrialPromotionalOfferSubscriptions         CustomInteger `csv:"Free Trial Promotional Offer Subscriptions" json:"free_trial_promotional_offer_subscriptions"`
	PayUpFrontPromotionalOfferSubscriptions        CustomInteger `csv:"Pay Up Front Promotional Offer Subscriptions" json:"pay_up_front_promotional_offer_subscriptions"`
	PayAsYouGoPromotionalOfferSubscriptions        CustomInteger `csv:"Pay As You Go Promotional Offer Subscriptions" json:"pay_as_you_go_promotional_offer_subscriptions"`
	MarketingOptIns                                CustomInteger `csv:"Marketing Opt-Ins" json:"marketing_opt_ins"`
	BillingRetry                                   CustomInteger `csv:"Billing Retry" json:"billing_retry"`
	GracePeriod                                    CustomInteger `csv:"Grace Period" json:"grace_period"`
}

//Aggregated data about subscriber activity, including upgrades, renewals, and introductory price conversions
type SalesReportSubscriptionEvent struct {
	EventDate                    CustomDate    `csv:"Event Date" json:"event_date"`
	Event                        string        `csv:"Event" json:"event"`
	AppName                      string        `csv:"App Name" json:"app_name"`
	AppAppleID                   CustomInteger `csv:"App Apple ID" json:"app_apple_id"`
	SubscriptionName             string        `csv:"Subscription Name" json:"subscription_name"`
	SubscriptionAppleID          CustomInteger `csv:"Subscription Apple ID" json:"subscription_apple_id"`
	SubscriptionGroupID          CustomInteger `csv:"Subscription Group ID" json:"subscription_group_id"`
	StandardSubscriptionDuration string        `csv:"Standard Subscription Duration" json:"standard_subscription_duration"`
	PromotionalOfferName         string        `csv:"Promotional Offer Name" json:"promotional_offer_name"`
	PromotionalOfferID           string        `csv:"Promotional Offer ID" json:"promotional_offer_id"`
	SubscriptionOfferType        string        `csv:"Subscription Offer Type" json:"subscription_offer_type"`
	SubscriptionOfferDuration    string        `csv:"Subscription Offer Duration" json:"subscription_offer_duration"`
	MarketingOptIn               string        `csv:"Marketing Opt-In" json:"marketing_opt_in"`
	MarketingOptInDuration       string        `csv:"Marketing Opt-In Duration" json:"marketing_opt_in_duration"`
	PreservedPricing             string        `csv:"Preserved Pricing" json:"preserved_pricing"`
	ProceedsReason               string        `csv:"Proceeds Reason" json:"proceeds_reason"`
	ConsecutivePaidPeriods       CustomInteger `csv:"Consecutive Paid Periods" json:"consecutive_paid_periods"`
	OriginalStartDate            CustomDate    `csv:"Original Start Date" json:"original_start_date"`
	Client                       string        `csv:"Client" json:"client"`
	Device                       string        `csv:"Device" json:"device"`
	State                        string        `csv:"State" json:"state"`
	Country                      string        `csv:"Country" json:"country"`
	PreviousSubscriptionName     string        `csv:"Previous Subscription Name" json:"previous_subscription_name"`
	PreviousSubscriptionAppleID  string        `csv:"Previous Subscription Apple ID" json:"previous_subscription_apple_id"`
	DaysBeforeCanceling          string        `csv:"Days Before Canceling" json:"days_before_canceling"`
	CancellationReason           string        `csv:"Cancellation Reason" json:"cancellation_reason"`
	DaysCanceled                 string        `csv:"Days Canceled" json:"days_canceled"`
	Quantity                     CustomInteger `csv:"Quantity" json:"quantity"`
}

//Transaction-level data about subscriber activity using randomly generated Subscriber IDs.
type SalesReportSubscriber struct {
	EventDate                    CustomDate    `csv:"Event Date" json:"event_date"`
	AppName                      string        `csv:"App Name" json:"app_name"`
	AppAppleID                   CustomInteger `csv:"App Apple ID" json:"app_apple_id"`
	SubscriptionName             string        `csv:"Subscription Name" json:"subscription_name"`
	SubscriptionAppleID          CustomInteger `csv:"Subscription Apple ID" json:"subscription_apple_id"`
	SubscriptionGroupID          CustomInteger `csv:"Subscription Group ID" json:"subscription_group_id"`
	StandardSubscriptionDuration string        `csv:"Standard Subscription Duration" json:"standard_subscription_duration"`
	PromotionalOfferName         string        `csv:"Promotional Offer Name" json:"promotional_offer_name"`
	PromotionalOfferID           string        `csv:"Promotional Offer ID" json:"promotional_offer_id"`
	SubscriptionOfferType        string        `csv:"Subscription Offer Type" json:"subscription_offer_type"`
	SubscriptionOfferDuration    string        `csv:"Subscription Offer Duration" json:"subscription_offer_duration"`
	MarketingOptInDuration       string        `csv:"Marketing Opt-In Duration" json:"marketing_opt_in_duration"`
	CustomerPrice                CustomFloat64 `csv:"Customer Price" json:"customer_price"`
	CustomerCurrency             string        `csv:"Customer Currency" json:"customer_currency"`
	DeveloperProceeds            CustomFloat64 `csv:"Developer Proceeds" json:"developer_proceeds"`
	ProceedsCurrency             string        `csv:"Proceeds Currency" json:"proceeds_currency"`
	PreservedPricing             string        `csv:"Preserved Pricing" json:"preserved_pricing"`
	ProceedsReason               string        `csv:"Proceeds Reason" json:"proceeds_reason"`
	Client                       string        `csv:"Client" json:"client"`
	Country                      string        `csv:"Country" json:"country"`
	SubscriberID                 CustomInteger `csv:"Subscriber ID" json:"subscriber_id"`
	SubscriberIDReset            string        `csv:"Subscriber ID Reset" json:"subscriber_id_reset"`
	Refund                       string        `csv:"Refund" json:"refund"`
	PurchaseDate                 CustomDate    `csv:"Purchase Date" json:"purchase_date"`
	Units                        CustomInteger `csv:"Units" json:"units"`
}
