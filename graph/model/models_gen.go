// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CryptoTransaction struct {
	ID       string  `json:"id"`
	From     *User   `json:"from"`
	FromID   string  `json:"from_id"`
	To       *User   `json:"to"`
	ToID     string  `json:"to_id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Datetime string  `json:"datetime"`
	Status   string  `json:"status"`
}

type Giftcard struct {
	ID       string  `json:"id"`
	Wallet   *Wallet `json:"wallet"`
	WalletID string  `json:"wallet_id"`
	Number   string  `json:"number"`
	Pin      string  `json:"pin"`
	Brand    string  `json:"brand"`
	Balance  int     `json:"balance"`
	Exp      string  `json:"exp"`
	Tradable bool    `json:"tradable"`
}

type GiftcardTransaction struct {
	ID       string `json:"id"`
	From     *User  `json:"from"`
	FromID   string `json:"from_id"`
	To       *User  `json:"to"`
	ToID     string `json:"to_id"`
	CardID   string `json:"card_id"`
	Datetime string `json:"datetime"`
	Status   string `json:"status"`
}

type User struct {
	ID          string  `json:"id"`
	Wallet      *Wallet `json:"wallet"`
	WalletID    string  `json:"wallet_id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Password    string  `json:"password"`
}

type Wallet struct {
	ID           string      `json:"id"`
	User         *User       `json:"user"`
	UserID       string      `json:"user_id"`
	BtcAddress   string      `json:"BTC_address"`
	BtcBalance   float64     `json:"BTC_balance"`
	EthAddress   string      `json:"ETH_address"`
	EthBalance   float64     `json:"ETH_balance"`
	ScashBalance float64     `json:"scash_balance"`
	Giftcards    []*Giftcard `json:"giftcards"`
}
