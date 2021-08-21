package tools

import "github.com/beego/beego/v2/server/web"

func GetRouter() web.LinkNamespace {
	ns := web.NSNamespace("/tools",
		web.NSRouter("/metric/", NewMetricController(), "get:Metrics"),
	)
	return ns
}
