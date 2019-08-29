// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/onap/v1alpha1.CollectdGlobal":       schema_pkg_apis_onap_v1alpha1_CollectdGlobal(ref),
		"./pkg/apis/onap/v1alpha1.CollectdGlobalSpec":   schema_pkg_apis_onap_v1alpha1_CollectdGlobalSpec(ref),
		"./pkg/apis/onap/v1alpha1.CollectdGlobalStatus": schema_pkg_apis_onap_v1alpha1_CollectdGlobalStatus(ref),
		"./pkg/apis/onap/v1alpha1.CollectdPlugin":       schema_pkg_apis_onap_v1alpha1_CollectdPlugin(ref),
		"./pkg/apis/onap/v1alpha1.CollectdPluginSpec":   schema_pkg_apis_onap_v1alpha1_CollectdPluginSpec(ref),
		"./pkg/apis/onap/v1alpha1.CollectdPluginStatus": schema_pkg_apis_onap_v1alpha1_CollectdPluginStatus(ref),
	}
}

func schema_pkg_apis_onap_v1alpha1_CollectdGlobal(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CollectdGlobal is the Schema for the collectdglobals API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/onap/v1alpha1.CollectdGlobalSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/onap/v1alpha1.CollectdGlobalStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/onap/v1alpha1.CollectdGlobalSpec", "./pkg/apis/onap/v1alpha1.CollectdGlobalStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_onap_v1alpha1_CollectdGlobalSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CollectdGlobalSpec defines the desired state of CollectdGlobal",
				Properties: map[string]spec.Schema{
					"globalOptions": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configMap": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"globalOptions"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_onap_v1alpha1_CollectdGlobalStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CollectdGlobalStatus defines the observed state of CollectdGlobal",
				Properties: map[string]spec.Schema{
					"collectdAgents": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html CollectdAgents are the collectd pods in the Daemonset Status can be one of \"\", Created, Deleting, Applied, Deprecated",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"status"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_onap_v1alpha1_CollectdPlugin(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CollectdPlugin is the Schema for the collectdplugins API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/onap/v1alpha1.CollectdPluginSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/onap/v1alpha1.CollectdPluginStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/onap/v1alpha1.CollectdPluginSpec", "./pkg/apis/onap/v1alpha1.CollectdPluginStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_onap_v1alpha1_CollectdPluginSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CollectdPluginSpec defines the desired state of CollectdPlugin",
				Properties: map[string]spec.Schema{
					"pluginName": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pluginConf": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"pluginName", "pluginConf"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_onap_v1alpha1_CollectdPluginStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CollectdPluginStatus defines the observed state of CollectdPlugin",
				Properties: map[string]spec.Schema{
					"collectdAgents": {
						SchemaProps: spec.SchemaProps{
							Description: "CollectdAgents are the collectd pods in the Daemonset Status can be one of \"\", Created, Deleting, Applied, Deprecated",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"status"},
			},
		},
		Dependencies: []string{},
	}
}
