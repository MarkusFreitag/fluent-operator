/*

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha2 "kubesphere.io/fluentbit-operator/api/generated/clientset/versioned/typed/fluentbitoperator/v1alpha2"
)

type FakeFluentbitoperatorV1alpha2 struct {
	*testing.Fake
}

func (c *FakeFluentbitoperatorV1alpha2) Filters(namespace string) v1alpha2.FilterInterface {
	return &FakeFilters{c, namespace}
}

func (c *FakeFluentbitoperatorV1alpha2) FluentBits(namespace string) v1alpha2.FluentBitInterface {
	return &FakeFluentBits{c, namespace}
}

func (c *FakeFluentbitoperatorV1alpha2) FluentBitConfigs(namespace string) v1alpha2.FluentBitConfigInterface {
	return &FakeFluentBitConfigs{c, namespace}
}

func (c *FakeFluentbitoperatorV1alpha2) Inputs(namespace string) v1alpha2.InputInterface {
	return &FakeInputs{c, namespace}
}

func (c *FakeFluentbitoperatorV1alpha2) Outputs(namespace string) v1alpha2.OutputInterface {
	return &FakeOutputs{c, namespace}
}

func (c *FakeFluentbitoperatorV1alpha2) Parsers(namespace string) v1alpha2.ParserInterface {
	return &FakeParsers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeFluentbitoperatorV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
