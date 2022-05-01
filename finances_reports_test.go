package appstore

//func Test_Finances_FinancialReport_MarshalJson(t *testing.T) {
//	reportData, _ := ioutil.ReadFile("stubs/reports/finances/financial.tsv")
//	r := bytes.NewReader(reportData)
//	decoder, _ := NewLineSkipDecoder(r, 3)
//	reports := []*FinancialReport{}
//	err := gocsv.UnmarshalDecoder(decoder, &reports)
//	fmt.Println(err)
//	assert.Error(t, err)
//	//expected := `{"provider":"APPLE","provider_country":"RU","sku":"","developer":"","title":"Foo","preorder_start_date":"2020-10-05","preorder_end_date":"2020-10-05","ordered":10.199999809265137,"canceled":5.5,"cumulative_ordered":10,"cumulative_canceled":12,"start_date":"2020-10-05","end_date":"2020-10-05","country_code":"RU","apple_identifier":1234567890,"device":"iPhone","supported_platforms":"iOS","category":"Lifestyle","client":"foo"}`
//	//data, _ := json.Marshal(reports[0])
//	//assert.Equal(t, expected, string(data))
//}
