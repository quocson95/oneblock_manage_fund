package models

type BinanceWalletBalance []BinanceWalletBalanceElement

type BinanceWalletBalanceElement struct {
	Activate   bool   `json:"activate"`
	Balance    string `json:"balance"`
	WalletName string `json:"walletName"`
}

type BinanceUserAsset []BinanceUserAssetElement

type BinanceUserAssetElement struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}
