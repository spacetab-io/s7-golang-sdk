package structsV052

type ItinReshopRS struct {
	Version  string `xml:"Version,attr"`
	Document *Document
	Response *ReshopResponse
	Errors   *Errors `xml:",omitempty"`
}

// GetPriceByPax func return passengers price object
func (r *ItinReshopRS) GetReshopPriceByPax(paxType string) *OfferPrice {

	for _, ReShopOffer := range r.Response.ReShopOffers.ReShopOffer {
		for _, OfferPrice := range ReShopOffer.ReShopPricedOffer.OfferPrice {
			for _, Association := range OfferPrice.RequestedDate.Associations {

				if getPaxByID(Association.AssociatedTraveler.TravelerReferences, r.Response.Passengers).PTC.Value == paxType {
					return OfferPrice
				}
			}
		}
	}

	return nil
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

func (r *ReshopResponse) GetTotalPrice() *DetailCurrencyPrice {

	return r.ReShopOffers.ReShopOffer[0].TotalPrice.DetailCurrencyPrice
}
