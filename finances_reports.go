package appstore

//FinancialReport
type FinancialReport struct {
	StartDate                            CustomDate    `csv:"Start Date" json:"start_date"`                                                                     //This is the period start date, based on Apple’s fiscal calendar.
	EndDate                              CustomDate    `csv:"End Date" json:"end_date"`                                                                         //This is the period end date, also based on Apple’s fiscal calendar.
	UPC                                  string        `csv:"UPC" json:"upc"`                                                                                   //This field is not applicable to developers. This will display as blank.
	ISRCIsbn                             string        `csv:"ISRC / ISBN" json:"isrc_isbn"`                                                                     //For apps, this is your SKU. For details, see App information. For in app purchases, this is the product ID. See In-app purchase information for details.
	VendorIdentifier                     string        `csv:"Vendor Identifier" json:"vendor_identifier"`                                                       //This is the “SKU” that was provided for an app, or a “Product ID” provided for an in-app purchase.
	Quantity                             CustomInteger `csv:"Quantity" json:"quantity"`                                                                         //Aggregated number of units sold.
	PartnerShare                         CustomFloat64 `csv:"Partner Share" json:"partner_share"`                                                               //The proceeds you receive per unit. This is the Customer Price minus applicable taxes and Apple’s commission, per Schedule 2 of your Paid Applications agreement.
	ExtendedPartnerShare                 CustomFloat64 `csv:"Extended Partner Share" json:"extended_partner_share"`                                             //Quantity multiplied by Partner Share.
	PartnerShareCurrency                 string        `csv:"Partner Share Currency" json:"partner_share_currency"`                                             //Three-character ISO code for the currency of the amounts earned.
	SaleOrReturn                         string        `csv:"Sales or Return" json:"sales_or_return"`                                                           //S indicates a Sale, R indicates a Return
	AppleIdentifier                      CustomInteger `csv:"Apple Identifier" json:"apple_identifier"`                                                         //Apple ID, a unique identifier automatically generated for your app when you add the app to your account. You can view this property in the App Information section in App Store Connect. This identifier is also used in the URL for the App Store on desktop computers. You can’t edit this property.
	ArtistShowDeveloperAuthor            string        `csv:"Artist / Show / Developer / Author" json:"artist_show_developer_author"`                           //Your legal entity name.
	Title                                string        `csv:"Title" json:"title"`                                                                               //The name you entered for your app as described in App information.
	LabelStudioNetworkDeveloperPublisher string        `csv:"Label / Studio / Network / Developer / Publisher" json:"label_studio_network_developer_publisher"` //This field is not applicable to developers. This will display as blank.
	Grid                                 string        `csv:"Grid" json:"grid"`                                                                                 //This field is not applicable to developers. This will display as blank.
	ProductTypeIdentifier                string        `csv:"Product Type Identifier" json:"product_type_identifier"`                                           //The type of product purchased. See Product Type Identifiers for more information.
	ISANOtherIdentifier                  string        `csv:"ISAN / Other Identifier" json:"isan_other_identifier"`                                             //This field is not applicable to developers. This will display as blank.
	CountryOfSale                        string        `csv:"Country Of Sale" json:"country_of_sale"`                                                           //Two-character ISO code (such as US for the United States) that indicates the country or region for the App Store where the purchase occurred. This is based on the customer Apple ID country or region.
	PreOrderFlag                         string        `csv:"Pre-order Flag" json:"preorder_flag"`                                                              //“P” or null
	PromoCode                            string        `csv:"Promo Code" json:"promo_code"`                                                                     //If the transaction was part of a promotion, a gift, or was downloaded through the Volume Purchase Program for Education, this field will contain a value. This field is empty for all non-promotional items. For more information, see Promotional Codes.
	CustomerPrice                        CustomFloat64 `csv:"Customer Price" json:"customer_price"`                                                             //The price per unit billed to the customer, which you set for your app or in-app purchase in App Store Connect. *Customer price is inclusive of any applicable taxes we collect and remit per Schedule 2 of the Paid Applications agreement.
	CustomerCurrency                     string        `csv:"Customer Currency" json:"customer_currency"`                                                       //Three-character ISO code for the currency type paid by the customer. For example, USD for United States Dollar.
}