/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"reflect"
	"text/template"

	"github.com/pkg/errors"
	"github.com/vishvananda/netlink"

	corev1 "k8s.io/api/core/v1"
	utilsnet "k8s.io/utils/net"
)

/* cni config management */

// CNIConfigInputs is supplied to the CNI config template

type CNIConfigInputs struct {
	PodCIDRs      []string
	RangeStart    []string
	DefaultRoutes []string
	Mtu           int // maximum transmission unit
}

// ComputeCNIConfigInputs computes the template inputs for CNIConfigWriter
func ComputeCNIConfigInputs(node *corev1.Node) CNIConfigInputs {
	inputs := CNIConfigInputs{}
	PodCIDRs, _ := utilsnet.ParseCIDRs(node.Spec.PodCIDRs) // already validated
	for _, podCIDR := range PodCIDRs {
		inputs.PodCIDRs = append(inputs.PodCIDRs, podCIDR.String())

		// define the default route
		if utilsnet.IsIPv4CIDR(podCIDR) {
			inputs.DefaultRoutes = append(inputs.DefaultRoutes, "0.0.0.0/0")
		} else {
			inputs.DefaultRoutes = append(inputs.DefaultRoutes, "::/0")
		}

		// reserve the first IPs of the range
		size := utilsnet.RangeSize(podCIDR)
		podCapacity := node.Status.Capacity.Pods().Value()
		if podCapacity == 0 {
			podCapacity = 110
		}

		rangeStart := ""
		offset := size - podCapacity
		if offset > 10 { // reserve the first 10 addresses of the Pod range if there is capacity
			startAddress, err := utilsnet.GetIndexedIP(podCIDR, 10)
			if err == nil {
				rangeStart = startAddress.String()
			}

		}
		inputs.RangeStart = append(inputs.RangeStart, rangeStart)
	}
	return inputs
}
