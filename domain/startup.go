package domain

import (
	"github.com/choby/go_ddd/domain/k8s_info"
	"github.com/choby/go_ddd/domain/testcmd"
	"github.com/choby/go_ddd/domain/trips"
)

func StartUp() (err error) {
	err = testcmd.StartUp()
	if err != nil {
		return
	}

	err = k8s_info.StartUp()
	if err != nil {
		return
	}

	err = trips.SingleTripsAgg.RegJob()
	if err != nil {
		return
	}
	return
}
