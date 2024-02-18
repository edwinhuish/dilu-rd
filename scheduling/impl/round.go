package impl

import (
	"github.com/baowk/dilu-rd/models"

	"go.uber.org/zap"
)

type RoundHandler struct {
	cur    map[string]int
	logger *zap.SugaredLogger
}

func NewRoundHandler(logger *zap.SugaredLogger) *RoundHandler {
	return &RoundHandler{
		cur:    make(map[string]int),
		logger: logger,
	}
}

func (r *RoundHandler) GetServiceNode(nodes []*models.ServiceNode, name string) *models.ServiceNode {
	if len(nodes) == 0 {
		return nil
	}
	for i := 0; i < len(nodes); i++ {
		if idx, ok := r.cur[name]; ok {
			useIdx := idx % len(nodes)
			r.cur[name] = useIdx + 1
			if nodes[useIdx].Enable() {
				return nodes[useIdx]
			}
		} else {
			r.cur[name] = 0
			if nodes[0].Enable() {
				return nodes[0]
			}
		}
	}
	return nil
}
