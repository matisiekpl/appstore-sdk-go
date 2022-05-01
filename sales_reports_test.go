package appstore

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_SalesReports_SalesReport_MarshalJson(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/sales.tsv")
	reports := []*SalesReport{}
	_ = UnmarshalCSV(reportData, &reports)
	expected := `{"provider":"APPLE","provider_country":"US","sku":"foo.bar.baz","developer":" ","name":"","title":"FooBarTitle","version":"","product_type_identifier":"IAY","units":12,"developer_proceeds":209.3000030517578,"begin_date":"2020-10-05","end_date":"2020-10-05","customer_currency":"RUB","country_code":"RU","currency_of_proceeds":"RUB","apple_identifier":1234567890,"customer_price":299,"promo_code":" ","parent_identifier":"foo.bar.baz","subscription":"Renewal","period":"7 Days","category":"Lifestyle","cmb":"","device":"iPhone","supported_platforms":"iOS","proceeds_reason":" ","preserved_pricing":"Yes","client":" ","order_type":" "}`
	data, _ := json.Marshal(reports[0])
	assert.Equal(t, expected, string(data))
}

func Test_SalesReports_SubscriptionsReport_MarshalJson(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/subscriptions.tsv")
	reports := []*SubscriptionsReport{}
	_ = UnmarshalCSV(reportData, &reports)
	expected := `{"app_name":"FooBarApp","app_apple_id":1234567890,"subscription_name":"foo.bar.baz","subscription_apple_id":1234567890,"subscription_group_id":1234567890,"standard_subscription_duration":"1 Year","promotional_offer_name":" ","promotional_offer_id":" ","customer_price":2950,"customer_currency":"RUB","developer_proceeds":2065,"proceeds_currency":"RUB","preserved_pricing":"","proceeds_reason":"","client":"","device":"iPhone","state":" ","country":"RU","active_standard_price_subscriptions":20,"active_free_trial_introductory_offer_subscriptions":0,"active_pay_up_front_introductory_offer_subscriptions":0,"active_pay_as_you_go_introductory_offer_subscriptions":0,"free_trial_promotional_offer_subscriptions":0,"pay_up_front_promotional_offer_subscriptions":0,"pay_as_you_go_promotional_offer_subscriptions":0,"marketing_opt_ins":0,"billing_retry":0,"grace_period":0,"free_trial_offer_code_subscriptions":0,"pay_up_front_offer_code_subscriptions":0,"pay_as_you_go_offer_code_subscriptions":0,"subscribers":""}`
	data, _ := json.Marshal(reports[0])
	//fmt.Println(string(data))
	assert.Equal(t, expected, string(data))
}

func Test_SalesReports_SubscriptionsEventsReport_MarshalJson(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/subscriptions-events.tsv")
	reports := []*SubscriptionsEventsReport{}
	_ = UnmarshalCSV(reportData, &reports)
	expected := `{"event_date":"2020-10-06","event":"Renew","app_name":"AppFooBar","app_apple_id":1234567890,"subscription_name":"foo.bar.baz","subscription_apple_id":1234567890,"subscription_group_id":1234567890,"standard_subscription_duration":"7 Days","promotional_offer_name":" ","promotional_offer_id":" ","subscription_offer_type":"","subscription_offer_duration":"","marketing_opt_in":"","marketing_opt_in_duration":" ","preserved_pricing":"","proceeds_reason":"","consecutive_paid_periods":11,"original_start_date":"2020-07-25","client":"","device":"iPhone","state":" ","country":"RU","previous_subscription_name":"","previous_subscription_apple_id":0,"days_before_canceling":0,"cancellation_reason":" ","days_canceled":0,"quantity":1}`
	data, _ := json.Marshal(reports[0])
	//fmt.Println(string(data))
	assert.Equal(t, expected, string(data))
}

func Test_SalesReports_SubscribersReport_MarshalJson(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/subscribers.tsv")
	reports := []*SubscribersReport{}
	_ = UnmarshalCSV(reportData, &reports)
	expected := `{"event_date":"2020-10-05","app_name":"FooBarApp","app_apple_id":1234567890,"subscription_name":"foo.bar.baz","subscription_apple_id":1234567890,"subscription_group_id":1234567890,"standard_subscription_duration":"7 Days","introductory_price_type":"","promotional_offer_name":"","promotional_offer_id":"","subscription_offer_name":"","subscription_offer_type":"","subscription_offer_duration":"","marketing_opt_in_duration":"","customer_price":4.489999771118164,"customer_currency":"USD","developer_proceeds":3.1500000953674316,"proceeds_currency":"USD","preserved_pricing":" ","proceeds_reason":" ","client":" ","country":"UA","subscriber_id":1234567890000,"subscriber_id_reset":"","refund":"","purchase_date":"","units":1}`
	data, _ := json.Marshal(reports[0])
	assert.Equal(t, expected, string(data))
}

func Test_SalesReports_PreOrdersReport_MarshalJson(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/preorders.tsv")
	reports := []*PreOrdersReport{}
	_ = UnmarshalCSV(reportData, &reports)
	expected := `{"provider":"APPLE","provider_country":"RU","sku":"","developer":"","title":"Foo","preorder_start_date":"2020-10-05","preorder_end_date":"2020-10-05","ordered":10.199999809265137,"canceled":5.5,"cumulative_ordered":10,"cumulative_canceled":12,"start_date":"2020-10-05","end_date":"2020-10-05","country_code":"RU","apple_identifier":1234567890,"device":"iPhone","supported_platforms":"iOS","category":"Lifestyle","client":"foo"}`
	data, _ := json.Marshal(reports[0])
	assert.Equal(t, expected, string(data))
}
