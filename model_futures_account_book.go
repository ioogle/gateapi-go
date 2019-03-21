/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.6.0
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesAccountBook struct {
	// Change time
	Time float32 `json:"time,omitempty"`
	// Change amount
	Change string `json:"change,omitempty"`
	// Balance after change
	Balance string `json:"balance,omitempty"`
	// Changing Type  - dnw: Deposit & Withdraw - pnl: Profit & Loss by reducing position - fee: Trading fee - refr: Referrer rebate - fund: Funding
	Type string `json:"type,omitempty"`
	// Comment
	Text string `json:"text,omitempty"`
}
