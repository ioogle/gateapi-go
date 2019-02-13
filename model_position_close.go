/*
 * Gate API v4
 *
 * APIv4 futures provides all sorts of futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 1.3.0
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type PositionClose struct {
	// Position close time
	Time float32 `json:"time,omitempty"`
	// Futures contract
	Contract string `json:"contract,omitempty"`
	// Position side, long or short
	Side string `json:"side,omitempty"`
	// PNL
	Pnl string `json:"pnl,omitempty"`
	// Text of close order
	Text string `json:"text,omitempty"`
}
