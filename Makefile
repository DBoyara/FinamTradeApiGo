proto:
	# - Update contracts from https://github.com/FinamWeb/trade-api-docs/tree/master/contracts
	#
	# - Change path to googleapis folder
	# https://github.com/googleapis/googleapis
	protoc  -I /Users/evgeny/Projects/googleapis \
        --go_out=. \
        --go-grpc_out=. \
        --go_opt=paths=source_relative \
        --go-grpc_opt=paths=source_relative \
        --proto_path=contracts \
        --go_opt=Mproto/tradeapi/v1/candles.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mproto/tradeapi/v1/common.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mproto/tradeapi/v1/events.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mproto/tradeapi/v1/orders.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mproto/tradeapi/v1/portfolios.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mproto/tradeapi/v1/security.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mproto/tradeapi/v1/stops.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mgrpc/tradeapi/v1/candles_grpc.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mgrpc/tradeapi/v1/events_grpc.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mgrpc/tradeapi/v1/orders_grpc.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mgrpc/tradeapi/v1/portfolios_grpc.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mgrpc/tradeapi/v1/securities_grpc.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        --go_opt=Mgrpc/tradeapi/v1/stops_grpc.proto=github.com/evsamsonov/FinamTradeGo/v2/tradeapi \
        proto/tradeapi/v1/candles.proto \
        proto/tradeapi/v1/common.proto \
        proto/tradeapi/v1/events.proto \
        proto/tradeapi/v1/orders.proto \
        proto/tradeapi/v1/portfolios.proto \
        proto/tradeapi/v1/security.proto \
        proto/tradeapi/v1/stops.proto \
        grpc/tradeapi/v1/candles.proto \
        grpc/tradeapi/v1/events.proto \
        grpc/tradeapi/v1/orders.proto \
        grpc/tradeapi/v1/portfolios.proto \
        grpc/tradeapi/v1/securities.proto \
        grpc/tradeapi/v1/stops.proto