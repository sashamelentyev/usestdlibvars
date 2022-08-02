package rpc

var _ = "/_goRPC_" // want `"/_goRPC_" can be replaced by rpc.DefaultRPCPath`

func _() {
	_ = "/debug/rpc" // want `"/debug/rpc" can be replaced by rpc.DefaultDebugPath`
}
