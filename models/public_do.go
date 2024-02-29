package models

type ExtraData struct {
	Mnemonic   string `json:"mnemonic"`      //Mnemonic
	Seed       string `json:"seed"`          //Seed hex string
	WIF        string `json:"wif,omitempty"` //Wallet import format
	RootKey    string `json:"root_key"`      //BASE58 root key
	PrivateKey string `json:"private_key"`   //Private key hex string
	PublicKey  string `json:"public_key"`    //Public key hex string
}
