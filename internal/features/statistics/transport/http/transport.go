package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/i-dontneedaname/study-todoapp/internal/core/domain"
	core_http_server "github.com/i-dontneedaname/study-todoapp/internal/core/transport/http/server"
)

type StatisticHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(ctx context.Context, userId *int, from *time.Time, to *time.Time) (domain.Statistics, error)
}

func NewStatisticHTTPHandler(statisticsService StatisticsService) *StatisticHTTPHandler {
	return &StatisticHTTPHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistic,
		},
	}
}
