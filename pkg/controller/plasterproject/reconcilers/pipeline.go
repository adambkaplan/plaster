package reconcilers

import (
	"context"
	"reflect"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	pipelinev1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/adambkaplan/plaster/pkg/apis/plaster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type PipelineReconciler struct {
	client client.Client
	scheme *runtime.Scheme
}

func NewPipelineReconciler(client client.Client, scheme *runtime.Scheme) *PipelineReconciler {
	return &PipelineReconciler{client: client, scheme: scheme}
}

func pipelineName(cr *v1alpha1.PlasterProject) string {
	return cr.Name + "-pipeline"
}

func mergePipeline(required, current *pipelinev1alpha1.Pipeline) *pipelinev1alpha1.Pipeline {
	merged := current.DeepCopy()
	merged.Labels = required.Labels
	merged.Spec = required.Spec
	return merged
}

func (p *PipelineReconciler) Reconcile(ctx context.Context, cr *v1alpha1.PlasterProject) (reconcile.Result, error) {
	expected, err := p.expected(cr)
	if err != nil {
		return reconcile.Result{}, err
	}
	current, err := p.current(ctx, cr)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating pipeline", "Pipeline.Namespace", expected.Namespace, "Pipeline.Name", expected.Name)
		return p.create(ctx, expected)
	} else if err != nil {
		return reconcile.Result{}, err
	}

	merged := mergePipeline(expected, current)
	if !reflect.DeepEqual(merged, current) {
		log.Info("Updating pipeline", "Pipeline.Namespace", merged.Namespace, "Pipeline.Name", merged.Name)
		return p.update(ctx, merged)
	}
	log.Info("Reconciled pipeline", "Pipeline.namespace", current.Namespace, "Pipeline.Name", current.Name)
	return reconcile.Result{}, nil
}

func (p *PipelineReconciler) expected(cr *v1alpha1.PlasterProject) (*pipelinev1alpha1.Pipeline, error) {
	labels := map[string]string{
		"plaster.pipelines.openshift.io/owner": cr.Name,
	}
	pipeline := &pipelinev1alpha1.Pipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pipelineName(cr),
			Namespace: cr.Namespace,
			Labels:    labels,
		},
	}
	// Set PlasterProject instance as the owner and controller
	if err := controllerutil.SetControllerReference(cr, pipeline, p.scheme); err != nil {
		return nil, err
	}
	return pipeline, nil
}

func (p *PipelineReconciler) current(ctx context.Context, cr *v1alpha1.PlasterProject) (*pipelinev1alpha1.Pipeline, error) {
	pipeline := &pipelinev1alpha1.Pipeline{}
	err := p.client.Get(ctx, types.NamespacedName{Name: pipelineName(cr), Namespace: cr.Namespace}, pipeline)
	return pipeline, err
}

func (p *PipelineReconciler) create(ctx context.Context, pipeline *pipelinev1alpha1.Pipeline) (reconcile.Result, error) {
	err := p.client.Create(ctx, pipeline)
	if err != nil {
		return reconcile.Result{}, err
	}
	return reconcile.Result{}, nil
}

func (p *PipelineReconciler) update(ctx context.Context, pipeline *pipelinev1alpha1.Pipeline) (reconcile.Result, error) {
	err := p.client.Update(ctx, pipeline)
	if err != nil {
		return reconcile.Result{}, err
	}
	return reconcile.Result{}, nil
}
