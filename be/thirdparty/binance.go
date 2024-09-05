package thirdparty

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"oneblock_manage_fund/config"
	"oneblock_manage_fund/models"
	"strconv"
	"strings"
	"time"
)

func BinanceApi(key, secret, method string, path string, query map[string]string, bodyReq []byte) ([]byte, int, error) {
	binPath := fmt.Sprintf("https://api.binance.com/%s", path)
	req, _ := http.NewRequest(method, binPath, bytes.NewReader(bodyReq))
	req.Header.Set("X-MBX-APIKEY", key)
	if query == nil {
		query = map[string]string{}
	}
	query["timestamp"] = strconv.FormatInt(binTime().UnixMilli(), 10)
	reqQuery := req.URL.Query()
	for k, v := range query {
		reqQuery.Add(k, v)
	}
	if len(reqQuery) > 0 {
		reqQuery.Add("signature", hmacUtil([]byte(secret), []byte(reqQuery.Encode())))
	}
	req.URL.RawQuery = reqQuery.Encode()
	fmt.Printf("do [%s] URL: [%s]", method, req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, 0, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}

func UserAsset(asserts ...string) (models.BinanceUserAsset, error) {
	// query["asset"] = "USDT"
	var query map[string]string
	if len(asserts) > 0 {
		query = make(map[string]string)
		query["asset"] = strings.Join(asserts, ",")
	}
	data, _, err := BinanceApi(config.Get().Binance.ID, config.Get().Binance.Secret, http.MethodPost, "sapi/v3/asset/getUserAsset", query, nil)
	if err != nil {
		return nil, err
	}
	s := make(models.BinanceUserAsset, 0)
	err = json.Unmarshal(data, &s)
	return s, err
}

func UserBalance() (models.BinanceWalletBalance, error) {
	data, _, err := BinanceApi(config.Get().Binance.ID, config.Get().Binance.Secret, http.MethodGet, "sapi/v1/asset/wallet/balance", nil, nil)
	if err != nil {
		return nil, err
	}
	s := make(models.BinanceWalletBalance, 0)
	err = json.Unmarshal(data, &s)
	return s, err
}

type ServerTime struct {
	ServerTime int64 `json:"serverTime"`
}

func binTime() time.Time {
	url := "https://api.binance.com/api/v3/time"

	// Send a GET request to the Binance API to get server time
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching server time:", err)
		return time.Time{}
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return time.Time{}
	}

	// Parse the JSON response into a ServerTime struct
	var serverTime ServerTime
	if err := json.Unmarshal(responseBody, &serverTime); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return time.Time{}
	}

	// Get the server time from the response
	binanceServerTime := time.Unix(serverTime.ServerTime/1000, 0)

	// Synchronize local system time with Binance server time
	localTime := time.Now()
	timeDifference := binanceServerTime.Sub(localTime)
	synchronizedTime := localTime.Add(timeDifference)
	return synchronizedTime
}

func hmacUtil(secretKey, message []byte) string {
	h := hmac.New(sha256.New, secretKey)

	// Write the message to the hasher
	h.Write(message)

	// Get the result of the HMAC computation
	hmacResult := h.Sum(nil)

	// Convert the result to a hexadecimal string
	hmacHex := hex.EncodeToString(hmacResult)

	return hmacHex
}
