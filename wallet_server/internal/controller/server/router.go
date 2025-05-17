package server

import (
	GetBalanceInteractor "wallet_server/internal/usecase/getBalanceInteractor"
	GetTransactionsHistoryInteractor "wallet_server/internal/usecase/getTransactionsHistoryInteractor"
	SendCurrencyInteractor "wallet_server/internal/usecase/sendCurrencyInteractor"
)

type Router struct {
	UCGetBalance             *GetBalanceInteractor.UseCase
	UCGetTransactionsHistory *GetTransactionsHistoryInteractor.UseCase
	UCSendCurrency           *SendCurrencyInteractor.UseCase
}

func NewRouter(
	getBalance *GetBalanceInteractor.UseCase,
	getTXHIstory *GetTransactionsHistoryInteractor.UseCase,
	sendCurrency *SendCurrencyInteractor.UseCase,
) *Router {
	return &Router{UCGetBalance: getBalance, UCGetTransactionsHistory: getTXHIstory, UCSendCurrency: sendCurrency}
}
