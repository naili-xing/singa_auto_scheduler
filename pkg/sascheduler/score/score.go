package score

import (
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

func Score(state *framework.CycleState, node *nodeinfo.NodeInfo) (int64, error) {
	score, err := CalculateCollectScore(state, node)
	if err != nil {
		return 0, err
	}
	return score, nil
}
