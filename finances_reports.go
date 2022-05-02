package appstore

//FinancialReport
type FinancialReport struct {
	StartDate                            CustomDate    `csv:"Start Date" json:"start_date"`
	EndDate                              CustomDate    `csv:"End Date" json:"end_date"`
	UPC                                  string        `csv:"UPC" json:"upc"`
	ISRCIsbn                             string        `csv:"ISRC / ISBN" json:"isrc_isbn"`
	VendorIdentifier                     string        `csv:"Vendor Identifier" json:"vendor_identifier"`
	Quantity                             CustomInteger `csv:"Quantity" json:"quantity"`
	PartnerShare                         CustomFloat64 `csv:"Partner Share" json:"partner_share"`
	ExtendedPartnerShare                 CustomFloat64 `csv:"Extended Partner Share" json:"extended_partner_share"`
	PartnerShareCurrency                 string        `csv:"Partner Share Currency" json:"partner_share_currency"`
	SaleOrReturn                         string        `csv:"Sales or Return" json:"sales_or_return"`
	AppleIdentifier                      CustomInteger `csv:"Apple Identifier" json:"apple_identifier"`
	ArtistShowDeveloperAuthor            string        `csv:"Artist / Show / Developer / Author" json:"artist_show_developer_author"`
	Title                                string        `csv:"Title" json:"title"`
	LabelStudioNetworkDeveloperPublisher string        `csv:"Label / Studio / Network / Developer / Publisher" json:"label_studio_network_developer_publisher"`
	Grid                                 string        `csv:"Grid" json:"grid"`
	ProductTypeIdentifier                string        `csv:"Product Type Identifier" json:"product_type_identifier"`
	ISANOtherIdentifier                  string        `csv:"ISAN / Other Identifier" json:"isan_other_identifier"`
	CountryOfSale                        string        `csv:"Country Of Sale" json:"country_of_sale"`
	PreOrderFlag                         string        `csv:"Pre-order Flag" json:"preorder_flag"`
	PromoCode                            string        `csv:"Promo Code" json:"promo_code"`
	CustomerPrice                        CustomFloat64 `csv:"Customer Price" json:"customer_price"`
	CustomerCurrency                     string        `csv:"Customer Currency" json:"customer_currency"`
}
