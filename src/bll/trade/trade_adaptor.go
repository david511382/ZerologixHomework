package trade

import (
	"time"
	"zerologix-homework/src/pkg/constant"
	"zerologix-homework/src/pkg/timeutil"
	errUtil "zerologix-homework/src/pkg/util/error"
	"zerologix-homework/src/repo/redis/schema/trade"
	orderRds "zerologix-homework/src/repo/redis/schema/trade/order"
)

type TradeStorage struct {
	timeUtil timeutil.ITime
	tradeRds *trade.Database
}

func NewTradeStorage(
	timeUtil timeutil.ITime,
	tradeRds *trade.Database,
) *TradeStorage {
	return &TradeStorage{
		timeUtil: timeUtil,
		tradeRds: tradeRds,
	}
}

func (l *TradeStorage) LockID() (
	resultErr error,
) {
	var reTryDelay time.Duration = 10
	for isTrying := true; isTrying; {
		isTrying = false
		timestamp := l.timeUtil.Now().UnixNano()
		if errInfo := l.tradeRds.OrderIDLocker.SetNX(timestamp, time.Second*10); errInfo != nil {
			if errUtil.Equal(errInfo, constant.ErrInfoRedisNotChange) {
				time.Sleep(time.Millisecond)
				isTrying = true
				reTryDelay += 10
			} else {
				resultErr = errInfo
				return
			}
		}
	}

	return
}

func (l *TradeStorage) ReleaseIDLock() {
	_, _ = l.tradeRds.OrderIDLocker.Del()
}

func (l *TradeStorage) LoadID() (
	id uint,
	resultErr error,
) {
	idP, errInfo := l.tradeRds.OrderID.Get()
	if errInfo != nil {
		resultErr = errInfo
		return
	}

	if idP == nil {
		id = 0
	} else {
		id = *idP
	}

	return
}

func (l *TradeStorage) UpdateID(id uint) (
	resultErr error,
) {
	if errInfo := l.tradeRds.OrderID.Set(id, 0); errInfo != nil {
		resultErr = errInfo
		return
	}

	return
}

func (l *TradeStorage) Load(
	price float64,
) (
	storeOrders []*Order,
	resultErr error,
) {
	price_ordersMap, err := l.tradeRds.Order.Read(price)
	if err != nil {
		resultErr = err
		return
	}
	for _, v := range price_ordersMap[price] {
		storeOrders = append(storeOrders, l.parseModelToOrder(price, v))
	}

	return
}

func (l *TradeStorage) LockOrders(
	price float64,
) (
	resultErr error,
) {
	var reTryDelay time.Duration = 10
	for isTrying := true; isTrying; {
		isTrying = false
		timestamp := l.timeUtil.Now().UnixNano()
		if errInfo := l.tradeRds.OrderLocker.HSetNx(price, timestamp); errInfo != nil {
			if errUtil.Equal(errInfo, constant.ErrInfoRedisNotChange) {
				time.Sleep(time.Millisecond * reTryDelay)
				isTrying = true
				reTryDelay += 10
			} else {
				resultErr = errInfo
				return
			}
		}
	}

	return
}

func (l *TradeStorage) ReleaseOrdersLock(
	price float64,
) {
	_, _ = l.tradeRds.OrderLocker.HDel(price)
}

func (l *TradeStorage) Set(
	price float64,
	orders []*Order,
) (
	resultErr error,
) {
	var storeOrders []*orderRds.Model
	for _, order := range orders {
		storeOrders = append(storeOrders, l.parseOrderToModel(order))
	}

	if errInfo := l.tradeRds.Order.HSet(price, storeOrders); errInfo != nil && !errUtil.Equal(errInfo, constant.ErrInfoRedisNotChange) {
		resultErr = errInfo
		return
	}

	return
}

func (l *TradeStorage) Delete(
	price float64,
) (
	resultErr error,
) {
	if _, errInfo := l.tradeRds.Order.HDel(price); errInfo != nil && !errUtil.Equal(errInfo, constant.ErrInfoRedisNotChange) {
		resultErr = errInfo
		return
	}

	return
}

func (l *TradeStorage) parseOrderToModel(order *Order) *orderRds.Model {
	return &orderRds.Model{
		ID:        order.ID,
		OrderType: uint8(order.OrderType),
		Quantity:  order.Quantity,
		Timestamp: order.Timestamp,
	}
}

func (l *TradeStorage) parseModelToOrder(price float64, order *orderRds.Model) *Order {
	return &Order{
		ID:        order.ID,
		OrderType: OrderType(order.OrderType),
		Quantity:  order.Quantity,
		Price:     price,
		Timestamp: order.Timestamp,
	}
}
