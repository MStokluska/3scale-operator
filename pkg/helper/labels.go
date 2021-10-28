package helper

import (
	"github.com/3scale/3scale-operator/pkg/3scale/amp/product"
)

type ComponentType string

const (
	ApplicationType    ComponentType = "application"
	InfrastructureType ComponentType = "infrastructure"
)

func MeteringLabels(componentName string, componentType ComponentType) map[string]string {
	return map[string]string{
		"com.company":   "Red_Hat",
		"rht.prod_name": "Red_Hat_Integration",
		// How to make this version match RHI version? Env var and data from env var?
		"rht.prod_ver":  "2021.Q4",
		"rht.comp":      "3scale",
		"rht.comp_ver":  product.ThreescaleRelease,
		"rht.subcomp":   componentName,
		"rht.subcomp_t": string(componentType),
	}
}
