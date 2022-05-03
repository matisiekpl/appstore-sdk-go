package appstore

//SalesReport Aggregated sales and download data for your apps and In-App Purchases
type SalesReport struct {
	Provider              string        `csv:"Provider" json:"provider"`                               //The service provider in your reports (typically Apple).
	ProviderCountry       string        `csv:"Provider Country" json:"provider_country"`               //The service provider country code (typically U.S.).
	SKU                   string        `csv:"SKU" json:"sku"`                                         //A product identifier provided by you during app setup.
	Developer             string        `csv:"Developer" json:"developer"`                             //Provided by you during the initial account setup.
	Name                  string        `csv:"Name" json:"name"`                                       //Provided by you during app setup.
	Title                 string        `csv:"Title" json:"title"`                                     //Provided by you during app setup.
	Version               string        `csv:"SalesReportVersion" json:"version"`                      //Provided by you during app setup.
	ProductTypeIdentifier string        `csv:"Product Type Identifier" json:"product_type_identifier"` //Defines the type of transaction (for example, initial download, update, and so on). For more information, see Product Type Identifiers.
	Units                 CustomFloat64 `csv:"Units" json:"units"`                                     //The aggregated number of units. Negative values indicate refunds, or CMB credits for previously purchased apps when CMB column shows ‘CMB-C’.
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

//SubscriptionsReport Total number of Active Subscriptions, Subscriptions with Introductory Prices, and Marketing Opt-Ins for your auto-renewable subscriptions.
type SubscriptionsReport struct {
	AppName                                        string        `csv:"App Name" json:"app_name"`                                                                                           //Title of your subscription’s parent app.
	AppAppleID                                     CustomInteger `csv:"App Apple ID" json:"app_apple_id"`                                                                                   //Apple ID of your subscription’s parent app.
	SubscriptionName                               string        `csv:"Subscription Name" json:"subscription_name"`                                                                         //Title of your subscription.
	SubscriptionAppleID                            CustomInteger `csv:"Subscription Apple ID" json:"subscription_apple_id"`                                                                 //Apple ID of your subscription.
	SubscriptionGroupID                            CustomInteger `csv:"Subscription Group ID" json:"subscription_group_id"`                                                                 //Your subscription’s Group ID (formerly Family ID).
	StandardSubscriptionDuration                   string        `csv:"Standard Subscription Duration" json:"standard_subscription_duration"`                                               //Duration of the standard subscription: 7 Days, 1 Month, 2 Months, 3 Months, 6 Months, or 1 Year.
	PromotionalOfferName                           string        `csv:"Promotional Offer Name" json:"promotional_offer_name"`                                                               //Retail Price displayed on the App Store and charged to the customer.
	PromotionalOfferID                             string        `csv:"Promotional Offer ID" json:"promotional_offer_id"`                                                                   //Three-character ISO code indicating the customer’s currency. For more information, see Currency codes.
	CustomerPrice                                  CustomFloat64 `csv:"Customer Price" json:"customer_price"`                                                                               //The proceeds for each subscription.
	CustomerCurrency                               string        `csv:"Customer Currency" json:"customer_currency"`                                                                         //The currency in which your proceeds are earned. For more information, see Currency codes.
	DeveloperProceeds                              CustomFloat64 `csv:"Developer Proceeds" json:"developer_proceeds"`                                                                       //For Renew events, if the price is preserved then this field equals “Yes”. Otherwise, it is blank.
	ProceedsCurrency                               string        `csv:"Proceeds Currency" json:"proceeds_currency"`                                                                         //For Renew events, if the subscription has been active for more than a year then you receive 85% of the customer price, minus applicable taxes, and this field equals “Rate After One Year”. Otherwise, you receive 70% and the field is blank.
	PreservedPricing                               string        `csv:"Preserved Pricing" json:"preserved_pricing"`                                                                         //The promotional offer reference name or the offer code reference name used in App Store Connect when setting up the subscription offer.
	ProceedsReason                                 string        `csv:"Proceeds Reason" json:"proceeds_reason"`                                                                             //An identifier that you set for your subscription offers in App Store Connect. For Promotional Offers this is the value entered in the Promotional Offer Reference Name field when setting up the offer. For one-time use offer codes, this is the value entered in the Offer Code Reference Name field when setting up the offer. For custom offer codes this is the code shared with your users.
	Client                                         string        `csv:"Client" json:"client"`                                                                                               //State field in the address submitted by the customer when signing up for their Apple ID. This field is not validated and may be blank.
	Device                                         string        `csv:"Device" json:"device"`                                                                                               //Two-character ISO country code indicating the App Store territory for the purchase. For more information, see Financial Report Regions and Currencies.
	State                                          string        `csv:"State" json:"state"`                                                                                                 //Type of device used for subscription purchase: iPhone, iPad, Apple TV, iPod touch, or Desktop.
	Country                                        string        `csv:"Country" json:"country"`                                                                                             //If the subscription was purchased from News then this field equals “News”. Otherwise, it is blank.
	ActiveStandardPriceSubscriptions               CustomInteger `csv:"Active Standard Price Subscriptions" json:"active_standard_price_subscriptions"`                                     //Total number of auto-renewable standard paid subscriptions currently active, excluding free trials, subscription offers, introductory offers, and marketing opt-ins. Subscriptions are active during the period for which the customer has paid without cancellation.
	ActiveFreeTrialIntroductoryOfferSubscriptions  CustomInteger `csv:"Active Free Trial Introductory Offer Subscriptions" json:"active_free_trial_introductory_offer_subscriptions"`       //Total number of introductory offer subscriptions currently in a free trial.
	ActivePayUpFrontIntroductoryOfferSubscriptions CustomInteger `csv:"Active Pay Up Front Introductory Offer Subscriptions" json:"active_pay_up_front_introductory_offer_subscriptions"`   //Total number of introductory offer subscriptions currently with a pay up front introductory price.
	ActivePayAsYouGoIntroductoryOfferSubscriptions CustomInteger `csv:"Active Pay As You Go Introductory Offer Subscriptions" json:"active_pay_as_you_go_introductory_offer_subscriptions"` //Total number of introductory offer subscriptions currently with a pay as you go introductory price.
	FreeTrialPromotionalOfferSubscriptions         CustomInteger `csv:"Free Trial Promotional Offer Subscriptions" json:"free_trial_promotional_offer_subscriptions"`                       //Total number of promotional offers currently in a free trial.
	PayUpFrontPromotionalOfferSubscriptions        CustomInteger `csv:"Pay Up Front Promotional Offer Subscriptions" json:"pay_up_front_promotional_offer_subscriptions"`                   //Total number of promotional offers with a pay up front promotional price.
	PayAsYouGoPromotionalOfferSubscriptions        CustomInteger `csv:"Pay As You Go Promotional Offer Subscriptions" json:"pay_as_you_go_promotional_offer_subscriptions"`                 //Total number of promotional offers with a pay as you go promotional price.
	MarketingOptIns                                CustomInteger `csv:"Marketing Opt-Ins" json:"marketing_opt_ins"`                                                                         //Total number of subscriptions currently in a marketing opt-in bonus period.
	BillingRetry                                   CustomInteger `csv:"Billing Retry" json:"billing_retry"`                                                                                 //Total number of subscriptions in the Billing Retry status. This indicates that the App Store is still attempting to automatically renew a subscription when billing issues arise (for example, an expired credit card). Available in reports for dates starting December 7, 2017.
	GracePeriod                                    CustomInteger `csv:"Grace Period" json:"grace_period"`                                                                                   //Total number of subscriptions in the Billing Grace Period state. This indicates that subscriber can continue accessing your content for a certain period of time (6 or 16 days) while Apple attempts to correct the billing issue.
	FreeTrialOfferCodeSubscriptions                CustomInteger `csv:"Free Trial Offer Code Subscriptions" json:"free_trial_offer_code_subscriptions"`                                     //Total number of offer code subscriptions currently in a free trial.
	PayUpFrontOfferCodeSubscriptions               CustomInteger `csv:"Pay Up Front Offer Code Subscriptions" json:"pay_up_front_offer_code_subscriptions"`                                 //Total number of offer code subscriptions with a pay up front offer price.
	PayAsYouGoOfferCodeSubscriptions               CustomInteger `csv:"Pay As You Go Offer Code Subscriptions" json:"pay_as_you_go_offer_code_subscriptions"`                               //Total number of offer code subscriptions with a pay as you go offer price.
	Subscribers                                    string        `csv:"Subscribers" json:"subscribers"`                                                                                     //The number of subscribers who have access to the auto-renewable subscription including entitled family members. Note that this field is only populated when the record represents more than 3 subscriptions. Learn more.
}

//SubscriptionsEventsReport Aggregated data about subscriber activity, including upgrades, renewals, and introductory price conversions
type SubscriptionsEventsReport struct {
	EventDate                    CustomDate    `csv:"Event Date" json:"event_date"`                                         //Date the event occurred.
	Event                        string        `csv:"Event" json:"event"`                                                   //Type of event that occurred. For more information, see Subscription Events.
	AppName                      string        `csv:"App Name" json:"app_name"`                                             //Title of your subscription’s parent app.
	AppAppleID                   CustomInteger `csv:"App Apple ID" json:"app_apple_id"`                                     //Apple ID of your subscription’s parent app.
	SubscriptionName             string        `csv:"Subscription Name" json:"subscription_name"`                           //Title of your subscription.
	SubscriptionAppleID          CustomInteger `csv:"Subscription Apple ID" json:"subscription_apple_id"`                   //Apple ID of your subscription.
	SubscriptionGroupID          CustomInteger `csv:"Subscription Group ID" json:"subscription_group_id"`                   //Your subscription’s Group ID (formerly Family ID).
	StandardSubscriptionDuration string        `csv:"Standard Subscription Duration" json:"standard_subscription_duration"` //Duration of the standard subscription: 7 Days, 1 Month, 2 Months, 3 Months, 6 Months, or 1 Year.
	SubscriptionOfferType        string        `csv:"Subscription Offer Type" json:"subscription_offer_type"`               //Type of introductory price: Pay Up Front, Pay As You Go, or Free Trial
	SubscriptionOfferDuration    string        `csv:"Subscription Offer Duration" json:"subscription_offer_duration"`       //Duration of the introductory price if applicable. For example: 3 Days, 7 Days, 2 Weeks, 1 Month, 2 Months, 3 Months, 6 Months, or 1 Year.
	MarketingOptIn               string        `csv:"Marketing Opt-In" json:"marketing_opt_in"`                             //If the subscription included a marketing opt-in then this field equals “Yes”. Otherwise, it is blank.
	MarketingOptInDuration       string        `csv:"Marketing Opt-In Duration" json:"marketing_opt_in_duration"`           //Duration of the opt-in if applicable: 7 Days, 1 Month, 2 Months, 3 Months, 6 Months, or 1 Year.
	PreservedPricing             string        `csv:"Preserved Pricing" json:"preserved_pricing"`                           //For Renew events, if the price is preserved then this field equals “Yes”. Otherwise, it is blank.
	ProceedsReason               string        `csv:"Proceeds Reason" json:"proceeds_reason"`                               //For Renew events, if the subscription has been active for more than a year then you receive 85% of the customer price, minus applicable taxes, and this field equals “Rate After One Year”. Otherwise, you receive 70% and the field is blank.
	PromotionalOfferName         string        `csv:"Promotional Offer Name" json:"promotional_offer_name"`                 //The Promotional Offer Reference Name used in App Store Connect when setting up the offer.
	PromotionalOfferID           string        `csv:"Promotional Offer ID" json:"promotional_offer_id"`                     //An identifier that you set for your subscription offers in App Store Connect. For Promotional Offers this is the value entered in the Promotional Offer Reference Name field when setting up the offer. For one-time use offer codes, this is the value entered in the Offer Code Reference Name field when setting up the offer. For custom offer codes this is the code shared with your users.
	ConsecutivePaidPeriods       CustomInteger `csv:"Consecutive Paid Periods" json:"consecutive_paid_periods"`             //The total number of paid periods that the subscription has been active without cancellation. This does not include free trials, marketing opt-in bonus periods, or grace periods.
	OriginalStartDate            CustomDate    `csv:"Original Start Date" json:"original_start_date"`                       //Date of the initial subscription purchase.
	Device                       string        `csv:"Device" json:"device"`                                                 //Type of device used for initial subscription purchase: iPhone, iPad, Apple TV, iPod touch, or Desktop.
	Client                       string        `csv:"Client" json:"client"`                                                 //If the subscription was purchased from News then this field equals “News”. Otherwise, it is blank.
	State                        string        `csv:"State" json:"state"`                                                   //State field in the address submitted by the customer when signing up for their Apple ID. This field is not validated and may be blank.
	Country                      string        `csv:"Country" json:"country"`                                               //Two-character ISO country code indicating the App Store territory for the purchase. For more information, see Financial Report Regions and Currencies
	PreviousSubscriptionName     string        `csv:"Previous Subscription Name" json:"previous_subscription_name"`         //For upgrade, downgrade, and crossgrade events, the title of the previous subscription.
	PreviousSubscriptionAppleID  CustomInteger `csv:"Previous Subscription Apple ID" json:"previous_subscription_apple_id"` //For upgrade, downgrade, and crossgrade events, the Apple ID of the previous subscription.
	DaysBeforeCanceling          CustomInteger `csv:"Days Before Canceling" json:"days_before_canceling"`                   //For cancel events, the number of days from the start date to when a subscriber canceled, which could be in the middle of the period. This only applies to cancel events where cancellation reason equals ‘canceled.' Otherwise, it is blank.
	CancellationReason           string        `csv:"Cancellation Reason" json:"cancellation_reason"`                       //Reason for a cancellation: Billing issue, Price increase, Canceled, Removed from Sale, or Other. For more information, see Cancellation Reasons.
	DaysCanceled                 CustomInteger `csv:"Days Canceled" json:"days_canceled"`                                   //For reactivate events, the number of days ago that the subscriber canceled.
	Quantity                     CustomInteger `csv:"Quantity" json:"quantity"`                                             //Number of events with the same values for the other fields.
}

//SubscribersReport Transaction-level data about subscriber activity using randomly generated Subscriber IDs.
type SubscribersReport struct {
	EventDate                    CustomDate    `csv:"Event Date" json:"event_date"`                                         //Date the event occurred.
	AppName                      string        `csv:"App Name" json:"app_name"`                                             //Title of your subscription’s parent app.
	AppAppleID                   CustomInteger `csv:"App Apple ID" json:"app_apple_id"`                                     //Apple ID of your subscription’s parent app.
	SubscriptionName             string        `csv:"Subscription Name" json:"subscription_name"`                           //Title of your subscription.
	SubscriptionAppleID          CustomInteger `csv:"Subscription Apple ID" json:"subscription_apple_id"`                   //Apple ID of your subscription.
	SubscriptionGroupID          CustomInteger `csv:"Subscription Group ID" json:"subscription_group_id"`                   //Your subscription’s Group ID (formerly Family ID).
	StandardSubscriptionDuration string        `csv:"Standard Subscription Duration" json:"standard_subscription_duration"` //Duration of the standard subscription: 7 Days, 1 Month, 2 Months, 3 Months, 6 Months, or 1 Year.
	SubscriptionOfferName        string        `csv:"Subscription Offer Name" json:"subscription_offer_name"`               //The promotional offer reference name or the offer code reference name used in App Store Connect when setting up the subscription offer.
	PromotionalOfferID           string        `csv:"Promotional Offer ID" json:"promotional_offer_id"`                     //A code that you create for customers to enter and redeem the subscription offer.
	IntroductoryPriceType        string        `csv:"Introductory Price Type" json:"introductory_price_type"`               //Type of introductory price: Pay Up Front, Pay As You Go, or Free Trial
	PromotionalOfferName         string        `csv:"Promotional Offer Name" json:"promotional_offer_name"`                 //The Promotional Offer Reference Name used in App Store Connect when setting up the Offer.
	SubscriptionOfferDuration    string        `csv:"Subscription Offer Duration" json:"subscription_offer_duration"`       //Duration of the introductory price if applicable: For example, 3 Days, 1 Week, 2 Weeks, 1 Month, 2 Months, 3 Months, 6 Months, or 1 Year.
	SubscriptionOfferType        string        `csv:"Subscription Offer Type" json:"subscription_offer_type"`               //The promotional offer reference name or the offer code reference name used in App Store Connect when setting up the subscription offer.
	MarketingOptInDuration       string        `csv:"Marketing Opt-In Duration" json:"marketing_opt_in_duration"`           //Duration of the marketing opt-in if applicable: 7 Days, 1 Month, 2 Months, 3 Months, 6 Months, or 1 Year.
	CustomerPrice                CustomFloat64 `csv:"Customer Price" json:"customer_price"`                                 //The price of your auto-renewable subscription.
	CustomerCurrency             string        `csv:"Customer Currency" json:"customer_currency"`                           //Three-character ISO code indicating the customer’s currency. For more information, see Currency codes.
	DeveloperProceeds            CustomFloat64 `csv:"Developer Proceeds" json:"developer_proceeds"`                         //The proceeds for each item delivered.
	ProceedsCurrency             string        `csv:"Proceeds Currency" json:"proceeds_currency"`                           //The currency in which your proceeds are earned. For more information, see Currency codes.
	PreservedPricing             string        `csv:"Preserved Pricing" json:"preserved_pricing"`                           //For renewals, if the price is preserved then this field equals “Yes”. Otherwise, it is blank.
	ProceedsReason               string        `csv:"Proceeds Reason" json:"proceeds_reason"`                               //If a subscription has been active for more than a year then you receive 85% of the customer price, minus applicable taxes, and this field equals “Rate After One Year.” Otherwise, you receive 70% and the field is blank.
	Client                       string        `csv:"Client" json:"client"`                                                 //If the subscription was purchased from News then this field equals “News”. Otherwise, it is blank.
	Country                      string        `csv:"Country" json:"country"`                                               //Two-character ISO country code indicating the App Store territory for the purchase. For more information, see Financial Report Regions and Currencies.
	SubscriberID                 CustomInteger `csv:"Subscriber ID" json:"subscriber_id"`                                   //The randomly generated Subscriber ID that is unique to each customer and developer.
	SubscriberIDReset            string        `csv:"Subscriber ID Reset" json:"subscriber_id_reset"`                       //If a customer cancels all of their subscriptions with you and does not resubscribe within 180 days, the Subscriber ID will be deleted. If the same customer resubscribes after 180 days, then we create a new Subscriber ID and this field equals “Yes.” Otherwise, it is blank.Subscriber IDs are reset when an app is transferred to another developer account.
	Refund                       string        `csv:"Refund" json:"refund"`                                                 //For full or partial refunds, this field equals “Yes.” Otherwise, it is blank.
	PurchaseDate                 CustomDate    `csv:"Purchase Date" json:"purchase_date"`                                   //For refunds, the date of the original purchase.
	Units                        CustomInteger `csv:"Units" json:"units"`                                                   //The aggregated number of units.
}

//PreOrdersReport Aggregated data for your apps made available for pre-order, including the number of units ordered and canceled by customers.
type PreOrdersReport struct {
	Provider           string        `csv:"Provider" json:"provider"`                        //Service provider in your reports (typically Apple).
	ProviderCountry    string        `csv:"Provider Country" json:"provider_country"`        //Service provider country code (typically U.S.).
	SKU                string        `csv:"SKU" json:"sku"`                                  //Product identifier provided by you during app setup.
	Developer          string        `csv:"Developer" json:"developer"`                      //Provided by you during the initial account setup.
	Title              string        `csv:"Title" json:"title"`                              //Provided by you during app setup.
	PreOrderStartDate  CustomDate    `csv:"Pre-Order Start Date" json:"preorder_start_date"` //Date the app becomes available for pre-order.
	PreOrderEndDate    CustomDate    `csv:"Pre-Order End Date" json:"preorder_end_date"`     //Last date the app is available for pre-order, after which the app is available for sale.
	Ordered            CustomFloat64 `csv:"Ordered" json:"ordered"`                          //Aggregated number of pre-orders for the period.
	Canceled           CustomFloat64 `csv:"Canceled" json:"canceled"`                        //Aggregated number of canceled pre-orders for the period.
	CumulativeOrdered  CustomFloat64 `csv:"Cumulative Ordered" json:"cumulative_ordered"`    //Total number of pre-orders since the start of the pre-order period.
	CumulativeCanceled CustomFloat64 `csv:"Cumulative Canceled" json:"cumulative_canceled"`  //Total number of canceled pre-orders since the start of the pre-order period.
	StartDate          CustomDate    `csv:"Start Date" json:"start_date"`                    //Start date of report.
	EndDate            CustomDate    `csv:"End Date" json:"end_date"`                        //End date of report.
	CountryCode        string        `csv:"Country Code" json:"country_code"`                //Two-character ISO country code indicating the App Store territory of the Pre-Order. For more information, see Financial Report Regions and Currencies.
	AppleIdentifier    CustomInteger `csv:"Apple Identifier" json:"apple_identifier"`        //Apple ID for your app.
	Device             string        `csv:"Device" json:"device"`                            //Type of device used for purchase: iPhone, iPad, Apple TV, iPod touch, or Desktop.
	SupportedPlatforms string        `csv:"Supported Platforms" json:"supported_platforms"`  //List of platforms that your app supports: iOS, tvOS, iOS and tvOS, or macOS.
	Category           string        `csv:"Category" json:"category"`                        //Indicates the category of the app, such as Games.
	Client             string        `csv:"Client" json:"client"`                            //Indicates where the purchase happened: App Store, App Store for iMessage, or blank.
}

//SubscriptionsOffersRedemptionReport
type SubscriptionsOffersRedemptionReport struct {
	Date                CustomDate    `csv:"Date" json:"date"`                                   //The date the redemptions occurred.
	AppName             string        `csv:"App Name" json:"app_name"`                           //Title of your subscription’s parent app.
	AppAppleID          CustomInteger `csv:"App Apple ID" json:"app_apple_id"`                   //Apple ID of your subscription’s parent app.
	SubscriptionName    string        `csv:"Subscription Name" json:"subscription_name"`         //Title of your subscription.
	SubscriptionAppleID CustomInteger `csv:"Subscription Apple ID" json:"subscription_apple_id"` //Apple ID of your subscription.
	OfferReferenceName  string        `csv:"Offer Reference Name" json:"offer_reference_name"`   //The Offer Reference Name used in App Store Connect when setting up the offer.
	OfferCode           string        `csv:"Offer Code" json:"offer_code"`                       //The custom code you created in App Store Connect. In the case of one-time use code redemptions, this field is blank.
	Territory           string        `csv:"Territory" json:"territory"`                         //Two-character ISO country code indicating the App Store territory.
	Redemptions         CustomInteger `csv:"Redemptions" json:"redemptions"`                     //Number of redemptions
}
