package k8s

import (
	"github.com/choby/go_ddd/domain/k8s_info"
	"github.com/choby/go_ddd/infrastructure/logs"
	v1 "k8s.io/api/core/v1"
)

func PodList(req PodsReq) (pods *v1.PodList, err error) {
	pods, err = k8s_info.SingleK8sInfoAgg.PodList()
	if err != nil {
		logs.Error("SingleK8sInfoAgg PodList ERR:%v", err)
		return
	}
	return
}
