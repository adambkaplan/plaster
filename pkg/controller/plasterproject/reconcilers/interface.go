package reconcilers

import (
	"context"

	"github.com/adambkaplan/plaster/pkg/apis/plaster/v1alpha1"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var log = logf.Log.WithName("controller_plasterproject_reconcilers")

type Reconciler interface {
	Reconcile(ctx context.Context, cr *v1alpha1.PlasterProject) (reconcile.Result, error)
}
