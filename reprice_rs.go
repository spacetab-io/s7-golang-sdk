package sdk

type ItinReshopRS struct {
	Version  string `xml:"Version,attr"`
	Document *Document
	Response *Response
	Errors   *Errors
}

type ReShopOffers struct {
	ReShopOffer []*ReShopOffer
}

type ReShopOffer struct {
	ID                string `xml:"ObjectKey,attr"`
	OfferID           *OfferID
	ReShopPricedOffer *PricedOffer
}
