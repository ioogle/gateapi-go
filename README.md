# Go API client for gateapi

APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.

## Specific note for 4.8.0

**BREAKING** change:

4.8.0 add new support with different settle currency for futures API(BTC is the only one allowed before), which makes ALL methods in FuturesApi REQUIRE an additional `settle` parameter.

But previous `/futures/xxx` APIs are still preserved for compatibility usage(will be treated as BTC), so if one of the following condition is met:

- Changing all your futures method call to include `settle` is not a big issue for you
- You need to use futures settled in non-BTC currency

then you'd better move to 4.8.0 and changes all your futures method call to pass in `settle` parameter. Otherwise, you can stick to version<=4.7.3,
but will not receive any future API upgrade support

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.14.0
- Package version: 4.14.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.gate.io/page/contacts](https://www.gate.io/page/contacts)

## Installation

Install the following dependencies:
```
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under the working directory(e.g. $GOPATH location), rename it into `gateapi` and in your own project add the following in import:
```golang
import "gateapi"
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Go code:

```golang

package main

import (
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against other hosts
    // client.ChangeBasePath("https://some-other-host")
    client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
    api := client.DeliveryApi
    
    settle := "usdt"; // string - Settle currency
    
    orderId := "12345"; // string - ID returned on order successfully being created
    result, _, err := api.CancelDeliveryOrder(nil, settle, orderId)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(result)
    }
}

```

## Documentation for API Endpoints

All URIs are relative to *https://api.gateio.ws/api/v4*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DeliveryApi* | [**CancelDeliveryOrder**](docs/DeliveryApi.md#canceldeliveryorder) | **Delete** /delivery/{settle}/orders/{order_id} | Cancel a single order
*DeliveryApi* | [**CancelDeliveryOrders**](docs/DeliveryApi.md#canceldeliveryorders) | **Delete** /delivery/{settle}/orders | Cancel all &#x60;open&#x60; orders matched
*DeliveryApi* | [**CancelPriceTriggeredDeliveryOrder**](docs/DeliveryApi.md#cancelpricetriggereddeliveryorder) | **Delete** /delivery/{settle}/price_orders/{order_id} | Cancel a single order
*DeliveryApi* | [**CancelPriceTriggeredDeliveryOrderList**](docs/DeliveryApi.md#cancelpricetriggereddeliveryorderlist) | **Delete** /delivery/{settle}/price_orders | Cancel all open orders
*DeliveryApi* | [**CreateDeliveryOrder**](docs/DeliveryApi.md#createdeliveryorder) | **Post** /delivery/{settle}/orders | Create a futures order
*DeliveryApi* | [**CreatePriceTriggeredDeliveryOrder**](docs/DeliveryApi.md#createpricetriggereddeliveryorder) | **Post** /delivery/{settle}/price_orders | Create a price-triggered order
*DeliveryApi* | [**GetDeliveryContract**](docs/DeliveryApi.md#getdeliverycontract) | **Get** /delivery/{settle}/contracts/{contract} | Get a single contract
*DeliveryApi* | [**GetDeliveryOrder**](docs/DeliveryApi.md#getdeliveryorder) | **Get** /delivery/{settle}/orders/{order_id} | Get a single order
*DeliveryApi* | [**GetDeliveryPosition**](docs/DeliveryApi.md#getdeliveryposition) | **Get** /delivery/{settle}/positions/{contract} | Get single position
*DeliveryApi* | [**GetMyDeliveryTrades**](docs/DeliveryApi.md#getmydeliverytrades) | **Get** /delivery/{settle}/my_trades | List personal trading history
*DeliveryApi* | [**GetPriceTriggeredDeliveryOrder**](docs/DeliveryApi.md#getpricetriggereddeliveryorder) | **Get** /delivery/{settle}/price_orders/{order_id} | Get a single order
*DeliveryApi* | [**ListDeliveryAccountBook**](docs/DeliveryApi.md#listdeliveryaccountbook) | **Get** /delivery/{settle}/account_book | Query account book
*DeliveryApi* | [**ListDeliveryAccounts**](docs/DeliveryApi.md#listdeliveryaccounts) | **Get** /delivery/{settle}/accounts | Query futures account
*DeliveryApi* | [**ListDeliveryCandlesticks**](docs/DeliveryApi.md#listdeliverycandlesticks) | **Get** /delivery/{settle}/candlesticks | Get futures candlesticks
*DeliveryApi* | [**ListDeliveryContracts**](docs/DeliveryApi.md#listdeliverycontracts) | **Get** /delivery/{settle}/contracts | List all futures contracts
*DeliveryApi* | [**ListDeliveryInsuranceLedger**](docs/DeliveryApi.md#listdeliveryinsuranceledger) | **Get** /delivery/{settle}/insurance | Futures insurance balance history
*DeliveryApi* | [**ListDeliveryLiquidates**](docs/DeliveryApi.md#listdeliveryliquidates) | **Get** /delivery/{settle}/liquidates | List liquidation history
*DeliveryApi* | [**ListDeliveryOrderBook**](docs/DeliveryApi.md#listdeliveryorderbook) | **Get** /delivery/{settle}/order_book | Futures order book
*DeliveryApi* | [**ListDeliveryOrders**](docs/DeliveryApi.md#listdeliveryorders) | **Get** /delivery/{settle}/orders | List futures orders
*DeliveryApi* | [**ListDeliveryPositionClose**](docs/DeliveryApi.md#listdeliverypositionclose) | **Get** /delivery/{settle}/position_close | List position close history
*DeliveryApi* | [**ListDeliveryPositions**](docs/DeliveryApi.md#listdeliverypositions) | **Get** /delivery/{settle}/positions | List all positions of a user
*DeliveryApi* | [**ListDeliverySettlements**](docs/DeliveryApi.md#listdeliverysettlements) | **Get** /delivery/{settle}/settlements | List settlement history
*DeliveryApi* | [**ListDeliveryTickers**](docs/DeliveryApi.md#listdeliverytickers) | **Get** /delivery/{settle}/tickers | List futures tickers
*DeliveryApi* | [**ListDeliveryTrades**](docs/DeliveryApi.md#listdeliverytrades) | **Get** /delivery/{settle}/trades | Futures trading history
*DeliveryApi* | [**ListPriceTriggeredDeliveryOrders**](docs/DeliveryApi.md#listpricetriggereddeliveryorders) | **Get** /delivery/{settle}/price_orders | List all auto orders
*DeliveryApi* | [**UpdateDeliveryPositionLeverage**](docs/DeliveryApi.md#updatedeliverypositionleverage) | **Post** /delivery/{settle}/positions/{contract}/leverage | Update position leverage
*DeliveryApi* | [**UpdateDeliveryPositionMargin**](docs/DeliveryApi.md#updatedeliverypositionmargin) | **Post** /delivery/{settle}/positions/{contract}/margin | Update position margin
*DeliveryApi* | [**UpdateDeliveryPositionRiskLimit**](docs/DeliveryApi.md#updatedeliverypositionrisklimit) | **Post** /delivery/{settle}/positions/{contract}/risk_limit | Update position risk limit
*FuturesApi* | [**CancelFuturesOrder**](docs/FuturesApi.md#cancelfuturesorder) | **Delete** /futures/{settle}/orders/{order_id} | Cancel a single order
*FuturesApi* | [**CancelFuturesOrders**](docs/FuturesApi.md#cancelfuturesorders) | **Delete** /futures/{settle}/orders | Cancel all &#x60;open&#x60; orders matched
*FuturesApi* | [**CancelPriceTriggeredOrder**](docs/FuturesApi.md#cancelpricetriggeredorder) | **Delete** /futures/{settle}/price_orders/{order_id} | Cancel a single order
*FuturesApi* | [**CancelPriceTriggeredOrderList**](docs/FuturesApi.md#cancelpricetriggeredorderlist) | **Delete** /futures/{settle}/price_orders | Cancel all open orders
*FuturesApi* | [**CreateFuturesOrder**](docs/FuturesApi.md#createfuturesorder) | **Post** /futures/{settle}/orders | Create a futures order
*FuturesApi* | [**CreatePriceTriggeredOrder**](docs/FuturesApi.md#createpricetriggeredorder) | **Post** /futures/{settle}/price_orders | Create a price-triggered order
*FuturesApi* | [**GetFuturesContract**](docs/FuturesApi.md#getfuturescontract) | **Get** /futures/{settle}/contracts/{contract} | Get a single contract
*FuturesApi* | [**GetFuturesOrder**](docs/FuturesApi.md#getfuturesorder) | **Get** /futures/{settle}/orders/{order_id} | Get a single order
*FuturesApi* | [**GetMyTrades**](docs/FuturesApi.md#getmytrades) | **Get** /futures/{settle}/my_trades | List personal trading history
*FuturesApi* | [**GetPosition**](docs/FuturesApi.md#getposition) | **Get** /futures/{settle}/positions/{contract} | Get single position
*FuturesApi* | [**GetPriceTriggeredOrder**](docs/FuturesApi.md#getpricetriggeredorder) | **Get** /futures/{settle}/price_orders/{order_id} | Get a single order
*FuturesApi* | [**ListFuturesAccountBook**](docs/FuturesApi.md#listfuturesaccountbook) | **Get** /futures/{settle}/account_book | Query account book
*FuturesApi* | [**ListFuturesAccounts**](docs/FuturesApi.md#listfuturesaccounts) | **Get** /futures/{settle}/accounts | Query futures account
*FuturesApi* | [**ListFuturesCandlesticks**](docs/FuturesApi.md#listfuturescandlesticks) | **Get** /futures/{settle}/candlesticks | Get futures candlesticks
*FuturesApi* | [**ListFuturesContracts**](docs/FuturesApi.md#listfuturescontracts) | **Get** /futures/{settle}/contracts | List all futures contracts
*FuturesApi* | [**ListFuturesFundingRateHistory**](docs/FuturesApi.md#listfuturesfundingratehistory) | **Get** /futures/{settle}/funding_rate | Funding rate history
*FuturesApi* | [**ListFuturesInsuranceLedger**](docs/FuturesApi.md#listfuturesinsuranceledger) | **Get** /futures/{settle}/insurance | Futures insurance balance history
*FuturesApi* | [**ListFuturesOrderBook**](docs/FuturesApi.md#listfuturesorderbook) | **Get** /futures/{settle}/order_book | Futures order book
*FuturesApi* | [**ListFuturesOrders**](docs/FuturesApi.md#listfuturesorders) | **Get** /futures/{settle}/orders | List futures orders
*FuturesApi* | [**ListFuturesTickers**](docs/FuturesApi.md#listfuturestickers) | **Get** /futures/{settle}/tickers | List futures tickers
*FuturesApi* | [**ListFuturesTrades**](docs/FuturesApi.md#listfuturestrades) | **Get** /futures/{settle}/trades | Futures trading history
*FuturesApi* | [**ListLiquidates**](docs/FuturesApi.md#listliquidates) | **Get** /futures/{settle}/liquidates | List liquidation history
*FuturesApi* | [**ListPositionClose**](docs/FuturesApi.md#listpositionclose) | **Get** /futures/{settle}/position_close | List position close history
*FuturesApi* | [**ListPositions**](docs/FuturesApi.md#listpositions) | **Get** /futures/{settle}/positions | List all positions of a user
*FuturesApi* | [**ListPriceTriggeredOrders**](docs/FuturesApi.md#listpricetriggeredorders) | **Get** /futures/{settle}/price_orders | List all auto orders
*FuturesApi* | [**UpdatePositionLeverage**](docs/FuturesApi.md#updatepositionleverage) | **Post** /futures/{settle}/positions/{contract}/leverage | Update position leverage
*FuturesApi* | [**UpdatePositionMargin**](docs/FuturesApi.md#updatepositionmargin) | **Post** /futures/{settle}/positions/{contract}/margin | Update position margin
*FuturesApi* | [**UpdatePositionRiskLimit**](docs/FuturesApi.md#updatepositionrisklimit) | **Post** /futures/{settle}/positions/{contract}/risk_limit | Update position risk limit
*MarginApi* | [**CancelLoan**](docs/MarginApi.md#cancelloan) | **Delete** /margin/loans/{loan_id} | Cancel lending loan
*MarginApi* | [**CreateLoan**](docs/MarginApi.md#createloan) | **Post** /margin/loans | Lend or borrow
*MarginApi* | [**GetLoan**](docs/MarginApi.md#getloan) | **Get** /margin/loans/{loan_id} | Retrieve one single loan detail
*MarginApi* | [**GetLoanRecord**](docs/MarginApi.md#getloanrecord) | **Get** /margin/loan_records/{loan_record_id} | Get one single loan record
*MarginApi* | [**ListFundingAccounts**](docs/MarginApi.md#listfundingaccounts) | **Get** /margin/funding_accounts | Funding account list
*MarginApi* | [**ListFundingBook**](docs/MarginApi.md#listfundingbook) | **Get** /margin/funding_book | Order book of lending loans
*MarginApi* | [**ListLoanRecords**](docs/MarginApi.md#listloanrecords) | **Get** /margin/loan_records | List repayment records of specified loan
*MarginApi* | [**ListLoanRepayments**](docs/MarginApi.md#listloanrepayments) | **Get** /margin/loans/{loan_id}/repayment | List loan repayment records
*MarginApi* | [**ListLoans**](docs/MarginApi.md#listloans) | **Get** /margin/loans | List all loans
*MarginApi* | [**ListMarginAccounts**](docs/MarginApi.md#listmarginaccounts) | **Get** /margin/accounts | Margin account list
*MarginApi* | [**ListMarginCurrencyPairs**](docs/MarginApi.md#listmargincurrencypairs) | **Get** /margin/currency_pairs | List all supported currency pairs supported in margin trading
*MarginApi* | [**MergeLoans**](docs/MarginApi.md#mergeloans) | **Post** /margin/merged_loans | Merge multiple lending loans
*MarginApi* | [**RepayLoan**](docs/MarginApi.md#repayloan) | **Post** /margin/loans/{loan_id}/repayment | Repay a loan
*MarginApi* | [**UpdateLoan**](docs/MarginApi.md#updateloan) | **Patch** /margin/loans/{loan_id} | Modify a loan
*MarginApi* | [**UpdateLoanRecord**](docs/MarginApi.md#updateloanrecord) | **Patch** /margin/loan_records/{loan_record_id} | Modify a loan record
*SpotApi* | [**CancelBatchOrders**](docs/SpotApi.md#cancelbatchorders) | **Post** /spot/cancel_batch_orders | Cancel a batch of orders with an ID list
*SpotApi* | [**CancelOrder**](docs/SpotApi.md#cancelorder) | **Delete** /spot/orders/{order_id} | Cancel a single order
*SpotApi* | [**CancelOrders**](docs/SpotApi.md#cancelorders) | **Delete** /spot/orders | Cancel all &#x60;open&#x60; orders in specified currency pair
*SpotApi* | [**CreateBatchOrders**](docs/SpotApi.md#createbatchorders) | **Post** /spot/batch_orders | Create a batch of orders
*SpotApi* | [**CreateOrder**](docs/SpotApi.md#createorder) | **Post** /spot/orders | Create an order
*SpotApi* | [**GetCurrencyPair**](docs/SpotApi.md#getcurrencypair) | **Get** /spot/currency_pairs/{currency_pair} | Get detail of one single order
*SpotApi* | [**GetOrder**](docs/SpotApi.md#getorder) | **Get** /spot/orders/{order_id} | Get a single order
*SpotApi* | [**ListCandlesticks**](docs/SpotApi.md#listcandlesticks) | **Get** /spot/candlesticks | Market candlesticks
*SpotApi* | [**ListCurrencyPairs**](docs/SpotApi.md#listcurrencypairs) | **Get** /spot/currency_pairs | List all currency pairs supported
*SpotApi* | [**ListMyTrades**](docs/SpotApi.md#listmytrades) | **Get** /spot/my_trades | List personal trading history
*SpotApi* | [**ListOrderBook**](docs/SpotApi.md#listorderbook) | **Get** /spot/order_book | Retrieve order book
*SpotApi* | [**ListOrders**](docs/SpotApi.md#listorders) | **Get** /spot/orders | List orders
*SpotApi* | [**ListSpotAccounts**](docs/SpotApi.md#listspotaccounts) | **Get** /spot/accounts | List spot accounts
*SpotApi* | [**ListTickers**](docs/SpotApi.md#listtickers) | **Get** /spot/tickers | Retrieve ticker information
*SpotApi* | [**ListTrades**](docs/SpotApi.md#listtrades) | **Get** /spot/trades | Retrieve market trades
*WalletApi* | [**GetDepositAddress**](docs/WalletApi.md#getdepositaddress) | **Get** /wallet/deposit_address | Generate currency deposit address
*WalletApi* | [**ListDeposits**](docs/WalletApi.md#listdeposits) | **Get** /wallet/deposits | Retrieve deposit records
*WalletApi* | [**ListSubAccountTransfers**](docs/WalletApi.md#listsubaccounttransfers) | **Get** /wallet/sub_account_transfers | Transfer records between main and sub accounts
*WalletApi* | [**ListWithdrawals**](docs/WalletApi.md#listwithdrawals) | **Get** /wallet/withdrawals | Retrieve withdrawal records
*WalletApi* | [**Transfer**](docs/WalletApi.md#transfer) | **Post** /wallet/transfers | Transfer between accounts
*WalletApi* | [**TransferWithSubAccount**](docs/WalletApi.md#transferwithsubaccount) | **Post** /wallet/sub_account_transfers | Transfer between main and sub accounts
*WithdrawalApi* | [**Withdraw**](docs/WithdrawalApi.md#withdraw) | **Post** /withdrawals | Withdraw


## Documentation For Models

 - [BatchOrder](docs/BatchOrder.md)
 - [CancelOrder](docs/CancelOrder.md)
 - [CancelOrderResult](docs/CancelOrderResult.md)
 - [Contract](docs/Contract.md)
 - [CurrencyPair](docs/CurrencyPair.md)
 - [DeliveryContract](docs/DeliveryContract.md)
 - [DeliverySettlement](docs/DeliverySettlement.md)
 - [DepositAddress](docs/DepositAddress.md)
 - [FundingAccount](docs/FundingAccount.md)
 - [FundingBookItem](docs/FundingBookItem.md)
 - [FundingRateRecord](docs/FundingRateRecord.md)
 - [FuturesAccount](docs/FuturesAccount.md)
 - [FuturesAccountBook](docs/FuturesAccountBook.md)
 - [FuturesCandlestick](docs/FuturesCandlestick.md)
 - [FuturesInitialOrder](docs/FuturesInitialOrder.md)
 - [FuturesLiquidate](docs/FuturesLiquidate.md)
 - [FuturesOrder](docs/FuturesOrder.md)
 - [FuturesOrderBook](docs/FuturesOrderBook.md)
 - [FuturesOrderBookItem](docs/FuturesOrderBookItem.md)
 - [FuturesPriceTrigger](docs/FuturesPriceTrigger.md)
 - [FuturesPriceTriggeredOrder](docs/FuturesPriceTriggeredOrder.md)
 - [FuturesTicker](docs/FuturesTicker.md)
 - [FuturesTrade](docs/FuturesTrade.md)
 - [GateErrorResponse](docs/GateErrorResponse.md)
 - [InsuranceRecord](docs/InsuranceRecord.md)
 - [LedgerRecord](docs/LedgerRecord.md)
 - [Loan](docs/Loan.md)
 - [LoanPatch](docs/LoanPatch.md)
 - [LoanRecord](docs/LoanRecord.md)
 - [MarginAccount](docs/MarginAccount.md)
 - [MarginAccountCurrency](docs/MarginAccountCurrency.md)
 - [MarginCurrencyPair](docs/MarginCurrencyPair.md)
 - [MyFuturesTrade](docs/MyFuturesTrade.md)
 - [Order](docs/Order.md)
 - [OrderBook](docs/OrderBook.md)
 - [Position](docs/Position.md)
 - [PositionClose](docs/PositionClose.md)
 - [PositionCloseOrder](docs/PositionCloseOrder.md)
 - [RepayRequest](docs/RepayRequest.md)
 - [Repayment](docs/Repayment.md)
 - [SpotAccount](docs/SpotAccount.md)
 - [SubAccountTransfer](docs/SubAccountTransfer.md)
 - [Ticker](docs/Ticker.md)
 - [Trade](docs/Trade.md)
 - [Transfer](docs/Transfer.md)
 - [TriggerOrderResponse](docs/TriggerOrderResponse.md)


## Author

support@mail.gate.io

