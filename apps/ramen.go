// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package apps

import (
	"log"

	ramenapi "github.com/ramendr/ramen/api/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	if err := ramenapi.AddToScheme(scheme.Scheme); err != nil {
		log.Fatal(err)
	}
}
