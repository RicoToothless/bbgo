// Code generated by "callbackgen -type WebSocketService"; DO NOT EDIT.

package max

import ()

func (s *WebSocketService) OnError(cb func(err error)) {
	s.errorCallbacks = append(s.errorCallbacks, cb)
}

func (s *WebSocketService) EmitError(err error) {
	for _, cb := range s.errorCallbacks {
		cb(err)
	}
}

func (s *WebSocketService) OnMessage(cb func(message []byte)) {
	s.messageCallbacks = append(s.messageCallbacks, cb)
}

func (s *WebSocketService) EmitMessage(message []byte) {
	for _, cb := range s.messageCallbacks {
		cb(message)
	}
}

func (s *WebSocketService) OnBookEvent(cb func(e BookEvent)) {
	s.bookEventCallbacks = append(s.bookEventCallbacks, cb)
}

func (s *WebSocketService) EmitBookEvent(e BookEvent) {
	for _, cb := range s.bookEventCallbacks {
		cb(e)
	}
}

func (s *WebSocketService) OnTradeEvent(cb func(e TradeEvent)) {
	s.tradeEventCallbacks = append(s.tradeEventCallbacks, cb)
}

func (s *WebSocketService) EmitTradeEvent(e TradeEvent) {
	for _, cb := range s.tradeEventCallbacks {
		cb(e)
	}
}

func (s *WebSocketService) OnErrorEvent(cb func(e ErrorEvent)) {
	s.errorEventCallbacks = append(s.errorEventCallbacks, cb)
}

func (s *WebSocketService) EmitErrorEvent(e ErrorEvent) {
	for _, cb := range s.errorEventCallbacks {
		cb(e)
	}
}

func (s *WebSocketService) OnSubscriptionEvent(cb func(e SubscriptionEvent)) {
	s.subscriptionEventCallbacks = append(s.subscriptionEventCallbacks, cb)
}

func (s *WebSocketService) EmitSubscriptionEvent(e SubscriptionEvent) {
	for _, cb := range s.subscriptionEventCallbacks {
		cb(e)
	}
}