package order

type GetOrderResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         *struct {
		AccountId        int    `json:"account-id"`
		Amount           string `json:"amount"`
		Id               int64  `json:"id"`
		ClientOrderId    string `json:"client-order-id"`
		Symbol           string `json:"symbol"`
		Price            string `json:"price"`
		CreatedAt        int64  `json:"created-at"`
		CanceledAt       int64  `json:"canceled-at,omitempty"`
		FinishedAt       int64  `json:"finished-at,omitempty"`
		Type             string `json:"type"`
		FilledAmount     string `json:"field-amount"`
		FilledCashAmount string `json:"field-cash-amount"`
		FilledFees       string `json:"field-fees"`
		Source           string `json:"source"`
		State            string `json:"state"`
		StopPrice        string `json:"stop_price,omitempty"`
		Operator         string `json:"operator,omitempty"`
	}
}
