package tools

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"poly-bridge/basedef"
	"poly-bridge/conf"

	"github.com/beego/beego/v2/server/web"
)

var (
	prefix  string
	metrics *Metrics

	metricUpdates = make(chan []byte, 10)
)

func Init() {
	metrics = NewMetrics()
}

type Metric struct {
	Key   string
	Value interface{}
}

type Metrics struct {
	state map[string]string
	ch    chan Metric
}

func Record(value interface{}, key string, args ...interface{}) {
	select {
	case metrics.ch <- Metric{Key: fmt.Sprintf(key, args...), Value: value}:
	default:
	}
}

func (m *Metrics) start() {
	ticker := time.NewTicker(time.Second)
	for metric := range m.ch {
		m.state[fmt.Sprintf("%s.%s", prefix, metric.Key)] = fmt.Sprintf("%v", metric.Value)
		select {
		case <-ticker.C:
			bytes, _ := json.Marshal(m.state)
			metricUpdates <- bytes
		default:
		}
	}
}

func NewMetrics() *Metrics {
	prefix = fmt.Sprintf("%s.%s", conf.GlobalConfig.Server, basedef.ENV)
	m := &Metrics{state: map[string]string{}, ch: make(chan Metric, 1000)}
	go m.start()
	return m
}

type MetricController struct {
	sync.RWMutex
	state []byte
	web.Controller
}

func (c *MetricController) Metrics() {
	c.RLock()
	state := c.state
	c.RUnlock()
	c.ServeJSON()
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.Output.Body(state)
}

func (c *MetricController) start() {
	for update := range metricUpdates {
		c.Lock()
		c.state = update
		c.Unlock()
	}
}

func NewMetricController() *MetricController {
	c := &MetricController{state: []byte("{}")}
	go c.start()
	return c
}
