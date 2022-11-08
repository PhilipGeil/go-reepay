package go_reepay

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CreateChargeSession creates a new charge returning the id and url for checkout if successful.
func (r *Reepay) CreateChargeSession(charge ChargeDTO) (chargeResponse *ChargeResponse, chargeError *ChargeError, err error) {
	charge.AcceptURL = &r.SuccessURL
	charge.CancelURL = &r.CancelURL

	j, err := charge.Marshal()
	if err != nil {
		return
	}

	url := "https://checkout-api.reepay.com/v1/session/charge/"

	body := strings.NewReader(string(j))

	req, _ := http.NewRequest("POST", url, body)

	keyEnc := base64.StdEncoding.EncodeToString([]byte(r.Key))

	req.Header.Add("Authorization", keyEnc)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode == 200 {
		err = json.Unmarshal(b, &chargeResponse)
	} else {
		err = json.Unmarshal(b, &chargeError)
	}
	return
}

func (r *Reepay) SettleChargeSession(charge SettleDTO, handle string) (settleResponse SettleResponse, chargeError *ChargeError, err error) {
	url := "https://api.reepay.com/v1/charge/" + handle + "/settle"

	j, err := charge.Marshal()
	if err != nil {
		return
	}

	body := strings.NewReader(string(j))

	req, _ := http.NewRequest("POST", url, body)

	keyEnc := base64.StdEncoding.EncodeToString([]byte(r.Key))

	req.Header.Add("Authorization", keyEnc)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode == 200 {
		err = json.Unmarshal(b, &settleResponse)
	} else {
		err = json.Unmarshal(b, &chargeError)
	}
	return

}

func (r *Reepay) CancelChargeSession(handle string) (settleResponse SettleResponse, chargeError *ChargeError, err error) {
	url := "https://api.reepay.com/v1/charge/" + handle + "/cancel"

	req, _ := http.NewRequest("POST", url, nil)

	keyEnc := base64.StdEncoding.EncodeToString([]byte(r.Key))

	req.Header.Add("Authorization", keyEnc)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode == 200 {
		err = json.Unmarshal(b, &settleResponse)
	} else {
		err = json.Unmarshal(b, &chargeError)
	}
	return

}

func (r *Reepay) RefundChargeSession(charge RefundDTO) (refundResponse *RefundResponse, chargeError *ChargeError, err error) {

	j, err := charge.Marshal()
	if err != nil {
		return
	}

	url := "https://api.reepay.com/v1/refund"

	body := strings.NewReader(string(j))

	req, _ := http.NewRequest("POST", url, body)

	keyEnc := base64.StdEncoding.EncodeToString([]byte(r.Key))

	req.Header.Add("Authorization", keyEnc)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode == 200 {
		err = json.Unmarshal(b, &refundResponse)
	} else {
		err = json.Unmarshal(b, &chargeError)
	}
	return
}

func (r *Reepay) GetChargeSession(handle string) (chargeResponse SettleResponse, chargeError *ChargeError, err error) {
	url := "https://api.reepay.com/v1/charge/" + handle

	req, _ := http.NewRequest("GET", url, nil)
	keyEnc := base64.StdEncoding.EncodeToString([]byte(r.Key))

	req.Header.Add("Authorization", keyEnc)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode == 200 {
		err = json.Unmarshal(b, &chargeResponse)
	} else {
		err = json.Unmarshal(b, &chargeError)
	}
	return
}
