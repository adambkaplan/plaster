package v1alpha1

import (
	buildv1 "github.com/openshift/api/build/v1"
	operatorv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PlasterProjectSpec defines the desired state of PlasterProject
type PlasterProjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Source is the source code used to run your Plaster project build.
	Source *PlasterSource `json:"source"`

	// Build contains the specifications on how to build your Plaster project.
	Build *PlasterBuild `json:"build"`

	// Tests lists the tests to run against your source code.
	// +optional
	Tests []PlasterTest `json:"tests,omitempty"`
}

type PlasterSourceType string

const (
	// GitPlasterSourceType represents source from a Git repository.
	GitPlasterSourceType = "Git"
)

type PlasterSource struct {
	// Type is the source type for the Plaster project.
	Type PlasterSourceType `json:"type"`

	// Git is a source from a Git repostority.
	// +optional
	Git *buildv1.GitBuildSource `json:"git,omitempty"`
}

type PlasterBuildType string

const (
	DockerfileBuildType = "Dockerfile"
	LanguageBuildType   = "Language"
)

type PlasterBuild struct {
	// Type is the type of build to use
	Type PlasterBuildType `json:"type"`

	// Dockerfile builds use a provided Dockerfile to build an application
	// +optional
	Dockerfile *PlasterDockerfileBuild `json:"dockerfile,omitempty"`

	// Language builds use a pre-baked set of languages to build an application
	// +optional
	Language *PlasterLanguageBuild `json:"language,omitempty"`
}

type PlasterDockerfileBuild struct {
	// Dockerfile is the docker file to use. If empty, Plaster assumes you want to build with a `Dockerfile` in the context root.
	Dockerfile string `json:"dockerfile,omitempty"`
}

type PlasterLanguageBuild struct {
	// Name is the language framework to use to build your application
	Name string `json:"name"`
	// Optmimized indicates if the deployed application should be optimized for running in production
	Optimized bool `json:"optimized"`
}

type PlasterTest struct {
	// Name is the name of the test
	Name string `json:"name"`
	// Command is the command and arguments to run in the test
	// +optional
	Command []string `json:"command,omitempty"`
}

// PlasterProjectStatus defines the observed state of PlasterProject
type PlasterProjectStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	ObservedGeneration int64                     `json:"observedGeneration,omitempty"`
	Conditions         []PlasterProjectCondition `json:"conditions,omitempty"`
}

type PlasterProjectCondition struct {
	Type               string                     `json:"type"`
	Status             operatorv1.ConditionStatus `json:"status"`
	LastTransitionTime metav1.Time                `json:"lastTransitionTime,omitempty"`
	Reason             string                     `json:"reason,omitempty"`
	Message            string                     `json:"reason,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlasterProject is the Schema for the plasterprojects API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=plasterprojects,scope=Namespaced
type PlasterProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlasterProjectSpec   `json:"spec,omitempty"`
	Status PlasterProjectStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlasterProjectList contains a list of PlasterProject
type PlasterProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlasterProject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PlasterProject{}, &PlasterProjectList{})
}
