package market

type Trade struct {
	ID     int     `json:"id"`
	Market int     `json:"market"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	IsBuy  bool    `json:"is_buy"`
}

type Market struct {
	count            			int
	buyCount 				 			int
	TotalVolume      			float64 `json:"total_volume"`
	TotalPrice       			float64 `json:"total_price"`
	totalPriceVolume    	float64
	MeanVolume       			float64 `json:"mean_volume"`
	MeanPrice        			float64 `json:"mean_price"`
	WeightedAvgPrice 	 		float64 `json:"volume_weighted_avg_price"`
	PercentageOrders 			float64 `json:"percentage_orders"`
}

type Markets map[int]*Market

func NewMarket() (*Market, error) {
	m := &Market{}
	return m, nil
}

func (m *Market) UpdateMarket(trade *Trade) (*Market, error) {
	var err error
	m.count =+ m.count + 1
	m.TotalVolume =+ trade.Volume
	m.TotalPrice =+ trade.Price
	m.totalPriceVolume =+ (trade.Price * trade.Volume)

	if (trade.IsBuy) {
		m.buyCount =+ 1
	}

	m.MeanVolume, err = m.getMeanVolume()
	if err != nil {
		return nil, err
	}

	m.MeanPrice, err = m.getMeanPrice()
	if err != nil {
		return nil, err
	}

	m.WeightedAvgPrice, err = m.getWeigtedVolumeAvgPrice()
	if err != nil {
		return nil, err
	}

	m.PercentageOrders, err = m.getPercentageOrders()
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Mean price per market
func (m *Market) getMeanPrice() (float64, error) {
	total := m.TotalPrice / float64(m.count)

	return total, nil
}

// Mean volume per market
func (m *Market) getMeanVolume() (float64, error) {
	total := m.TotalVolume / float64(m.count)

	return total, nil
}

// Volume-weighted average price per market
func (m *Market) getWeigtedVolumeAvgPrice() (float64, error) {
	total := (m.totalPriceVolume) / float64(m.TotalVolume)

	return total, nil
}

// Percentage buy orders per market
func (m *Market) getPercentageOrders() (float64, error) {
	total := (m.buyCount / m.count) * 100

	return float64(total), nil
}
