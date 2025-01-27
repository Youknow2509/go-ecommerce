package routers

import (
	"github.com/Youknow2509/go-ecommerce/internal/routers/manage"
	"github.com/Youknow2509/go-ecommerce/internal/routers/prometheus"
	"github.com/Youknow2509/go-ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
	Prometheus prometheus.PrometheusRouter
}

var RouterGroupApp = new(RouterGroup)