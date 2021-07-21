package main

import "time"

type Common struct {
	ChargingMod string //付费模式，预付费和后付费
	Region      string
	Az          string
	CreateTime  time.Time
}

type yunzhuji struct {
	guige string
}

type yunzhujiku struct {
	dbType string //代表数据库是哪一种
}
