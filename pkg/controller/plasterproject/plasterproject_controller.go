package plasterproject

import (
	"context"

	"github.com/adambkaplan/plaster/pkg/controller/plasterproject/reconcilers"

	plasterv1alpha1 "github.com/adambkaplan/plaster/pkg/apis/plaster/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	pipelinev1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
)

var log = logf.Log.WithName("controller_plasterproject")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new PlasterProject Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	client := mgr.GetClient()
	scheme := mgr.GetScheme()

	return &ReconcilePlasterProject{client: client,
		scheme:   scheme,
		pipeline: reconcilers.NewPipelineReconciler(client, scheme),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("plasterproject-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource PlasterProject
	err = c.Watch(&source.Kind{Type: &plasterv1alpha1.PlasterProject{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		log.Error(err, "Failed to watch PlasterProject")
		return err
	}

	// Watch for changes to secondary resource Pipeline and requeue the owner PlasterProject
	// - tekton.dev/v1alpha1 Pipeline
	err = c.Watch(&source.Kind{Type: &pipelinev1alpha1.Pipeline{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &plasterv1alpha1.PlasterProject{},
	})
	if err != nil {
		log.Error(err, "Failed to watch Pipelines")
		return err
	}

	return nil
}

// blank assignment to verify that ReconcilePlasterProject implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcilePlasterProject{}

// ReconcilePlasterProject reconciles a PlasterProject object
type ReconcilePlasterProject struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme

	pipeline reconcilers.Reconciler
}

// Reconcile reads that state of the cluster for a PlasterProject object and makes changes based on the state read
// and what is in the PlasterProject.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcilePlasterProject) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling PlasterProject")

	// Fetch the PlasterProject instance
	ctx := context.Background()
	instance := &plasterv1alpha1.PlasterProject{}
	err := r.client.Get(ctx, request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	pipelineReconcile, err := r.pipeline.Reconcile(ctx, instance)
	if err != nil {
		return pipelineReconcile, err
	}

	return reconcile.Result{}, nil
}
