package s7_api_sdk

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
	TotalPrice        *TotalPrice
	ReShopPricedOffer *PricedOffer
}

type DetailsPrice struct {
	Detail *DetalPrice
}

type DetalPrice struct {
	SubTotal    *Total
	Application string
}

type Surcharges struct {
	Surcharge *Surcharge
}

type Surcharge struct {
	Total *Total
}
