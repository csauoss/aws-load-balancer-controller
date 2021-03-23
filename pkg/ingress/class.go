package ingress

import (
	networking "k8s.io/api/networking/v1beta1"
	elbv2api "sigs.k8s.io/aws-load-balancer-controller/apis/elbv2/v1beta1"
)

// ClassConfiguration contains configurations for IngressClass
type ClassConfiguration struct {
	// The IngressClass for Ingress if any.
	IngClass *networking.IngressClass
	// The IngressClassParams for Ingress if any.
	IngClassParams *elbv2api.IngressClassParams
}
