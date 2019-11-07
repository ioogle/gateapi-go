/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesAccount struct {
	// Total assets, total = position_margin + order_margin + available
	Total string `json:"total,omitempty"`
	// Unrealized PNL
	UnrealisedPnl string `json:"unrealised_pnl,omitempty"`
	// Position margin
	PositionMargin string `json:"position_margin,omitempty"`
	// Order margin of unfinished orders
	OrderMargin string `json:"order_margin,omitempty"`
	// Available balance to transfer out or trade
	Available string `json:"available,omitempty"`
	// POINT amount
	Point string `json:"point,omitempty"`
	// Settle currency
	Currency string `json:"currency,omitempty"`
}
