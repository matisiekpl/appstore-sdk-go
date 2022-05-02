package appstore

import (
	"bytes"
	"encoding/json"
	"github.com/gocarina/gocsv"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_Finances_FinancialReport_MarshalJson(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/finances/financial.tsv")
	r := bytes.NewReader(reportData)
	decoder, err := NewLineSkipDecoder(r)
	reports := []*FinancialReport{}
	err = gocsv.UnmarshalDecoder(decoder, &reports)
	assert.NoError(t, err)
	expected := `{"start_date":"2020-10-05","end_date":"2021-10-05","upc":"","isrc_isbn":"","vendor_identifier":"foo.bar.baz","quantity":1,"partner_share":3.1500000953674316,"extended_partner_share":3.1500000953674316,"partner_share_currency":"USD","sales_or_return":"S","apple_identifier":1234567890,"artist_show_developer_author":"","title":"foo.bar.baz","label_studio_network_developer_publisher":"","grid":"","product_type_identifier":"IAY","isan_other_identifier":"","country_of_sale":"US","preorder_flag":"","promo_code":"","customer_price":4.489999771118164,"customer_currency":"USD"}`
	data, _ := json.Marshal(reports[0])
	assert.Equal(t, expected, string(data))
}
