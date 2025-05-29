package remote

type DeviceMarketName struct {
	OPPO    string
	HONOR   string
	XIAOMI  string
	IQOO    string
	HUAWEI  string
	ONEPLUS string
	REDMI   string
}

var MarketNames = DeviceMarketName{
	OPPO:    "ro.oppo.market.name",
	HONOR:   "ro.config.marketing_name",
	XIAOMI:  "ro.product.marketname",
	IQOO:    "ro.vivo.market.name",
	HUAWEI:  "ro.config.marketing_name",
	ONEPLUS: "ro.vendor.oplus.market.name",
	REDMI:   "ro.product.vendor.model",
}
