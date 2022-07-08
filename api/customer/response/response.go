package response

import "api-redeem-point/business/customer"

type Login struct {
	Code     string `json:"code"`
	Messages string `json:"messages"`
	Results  customer.ResponseLogin
}
