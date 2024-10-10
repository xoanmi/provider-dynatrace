// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alerting "github.com/xoanmi/provider-dynatrace/internal/controller/alerting/alerting"
	dashboard "github.com/xoanmi/provider-dynatrace/internal/controller/dashboard/dashboard"
	events "github.com/xoanmi/provider-dynatrace/internal/controller/event/events"
	notification "github.com/xoanmi/provider-dynatrace/internal/controller/notification/notification"
	providerconfig "github.com/xoanmi/provider-dynatrace/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alerting.Setup,
		dashboard.Setup,
		events.Setup,
		notification.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
