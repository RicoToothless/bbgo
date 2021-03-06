// Code generated by "callbackgen -type DepthFrame"; DO NOT EDIT.

package binance

import ()

func (f *DepthFrame) OnReady(cb func(snapshotDepth DepthEvent, bufEvents []DepthEvent)) {
	f.readyCallbacks = append(f.readyCallbacks, cb)
}

func (f *DepthFrame) EmitReady(snapshotDepth DepthEvent, bufEvents []DepthEvent) {
	for _, cb := range f.readyCallbacks {
		cb(snapshotDepth, bufEvents)
	}
}

func (f *DepthFrame) OnPush(cb func(e DepthEvent)) {
	f.pushCallbacks = append(f.pushCallbacks, cb)
}

func (f *DepthFrame) EmitPush(e DepthEvent) {
	for _, cb := range f.pushCallbacks {
		cb(e)
	}
}
