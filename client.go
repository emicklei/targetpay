package targetpay

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const checkPaymentURL = "https://www.targetpay.com/api/sms-pincode"

const errorCodes = `000 OK
102 No layout code specified
103 No pincode specified
104 Pincode length incorrect
105 Internal Error: no connection to SMS gateway
106 Pincode already checked or not paid
107 Layoutcode unknown
108 No keyword specified
109 No country specified
110 No shortcode specified`

const NL = 31
const BE = 32

type Client struct {
	httpClient *http.Client
}

func NewClient() Client {
	c := Client{}
	c.httpClient = &http.Client{}
	return c
}

func (c Client) CheckPayment(rtlo int, code int, keyword string, shortcode int, country int, test bool) (bool, string) {
	query := url.Values{}
	query.Set("rtlo", strconv.Itoa(rtlo))
	if code < 100000 || code > 999999 {
		return false, "Pincode length incorrect"
	}
	query.Set("code", strconv.Itoa(code))
	query.Set("keyword", keyword)
	query.Set("shortcode", strconv.Itoa(shortcode))
	query.Set("country", strconv.Itoa(country))
	if test {
		query.Set("test", "1")
	}
	url := fmt.Sprintf("%s?%s", checkPaymentURL, query.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err.Error()
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err.Error()
	}
	return true, string(body)
}
