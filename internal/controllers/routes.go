package controllers

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/payment/invoice/{id}", server.CreateInvoice).Methods("GET")
	server.Router.HandleFunc("/payment/invoice", server.GetInvoice).Methods("GET")
	server.Router.HandleFunc("/payment/invoice/{id}", server.GetInvoiceById).Methods("GET")

	server.Router.HandleFunc("/payment/deduction", server.CreateDeduction).Methods("POST")
	server.Router.HandleFunc("/payment/deduction", server.GetDeduction).Methods("GET")
	server.Router.HandleFunc("/payment/deduction", server.GetDeductionById).Methods("GET")

	server.Router.HandleFunc("/payment/threshold", server.CreateThreshold).Methods("POST")
	server.Router.HandleFunc("/payment/threshold", server.GetThreshold).Methods("GET")
	server.Router.HandleFunc("/payment/threshold/{id}", server.GetThresholdById).Methods("GET")

	server.Router.HandleFunc("/payment/promocode", server.CreatePromoCode).Methods("POST")
	server.Router.HandleFunc("/payment/promocode", server.GetPromoCode).Methods("GET")
	server.Router.HandleFunc("/payment/promocode/{id}", server.GetPromoCodeById).Methods("GET")

	server.Router.HandleFunc("/payment/gateway", server.CreateGateway).Methods("POST")
	server.Router.HandleFunc("/payment/gateway", server.GetGateway).Methods("GET")
	server.Router.HandleFunc("/payment/gateway/{id}", server.GetGatewayById).Methods("GET")

	server.Router.HandleFunc("/payment/paymenthistory", server.CreatePaymentHistory).Methods("POST")
	server.Router.HandleFunc("/payment/paymenthistory", server.GetPaymentHistory).Methods("GET")
	server.Router.HandleFunc("/payment/paymenthistory/{id}", server.GetPaymentHistoryById).Methods("GET")

	server.Router.HandleFunc("/payment/invoiceitems", server.CreateInvoiceItems).Methods("POST")
	server.Router.HandleFunc("/payment/invoiceitems", server.GetInvoiceItems).Methods("GET")
	server.Router.HandleFunc("/payment/invoiceitems/{id}", server.GetInvoiceItemsById).Methods("GET")

	server.Router.HandleFunc("/payment/transaction", server.CreateTransaction).Methods("POST")
	server.Router.HandleFunc("/payment/transaction", server.GetTransaction).Methods("GET")
	server.Router.HandleFunc("/payment/transaction/{id}", server.GetTransactionById).Methods("GET")

}
