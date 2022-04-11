/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/linkerd/linkerd2/controller/gen/apis/policy/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkAuthenticationLister helps list NetworkAuthentications.
// All objects returned here must be treated as read-only.
type NetworkAuthenticationLister interface {
	// List lists all NetworkAuthentications in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkAuthentication, err error)
	// NetworkAuthentications returns an object that can list and get NetworkAuthentications.
	NetworkAuthentications(namespace string) NetworkAuthenticationNamespaceLister
	NetworkAuthenticationListerExpansion
}

// networkAuthenticationLister implements the NetworkAuthenticationLister interface.
type networkAuthenticationLister struct {
	indexer cache.Indexer
}

// NewNetworkAuthenticationLister returns a new NetworkAuthenticationLister.
func NewNetworkAuthenticationLister(indexer cache.Indexer) NetworkAuthenticationLister {
	return &networkAuthenticationLister{indexer: indexer}
}

// List lists all NetworkAuthentications in the indexer.
func (s *networkAuthenticationLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkAuthentication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkAuthentication))
	})
	return ret, err
}

// NetworkAuthentications returns an object that can list and get NetworkAuthentications.
func (s *networkAuthenticationLister) NetworkAuthentications(namespace string) NetworkAuthenticationNamespaceLister {
	return networkAuthenticationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkAuthenticationNamespaceLister helps list and get NetworkAuthentications.
// All objects returned here must be treated as read-only.
type NetworkAuthenticationNamespaceLister interface {
	// List lists all NetworkAuthentications in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkAuthentication, err error)
	// Get retrieves the NetworkAuthentication from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NetworkAuthentication, error)
	NetworkAuthenticationNamespaceListerExpansion
}

// networkAuthenticationNamespaceLister implements the NetworkAuthenticationNamespaceLister
// interface.
type networkAuthenticationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkAuthentications in the indexer for a given namespace.
func (s networkAuthenticationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkAuthentication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkAuthentication))
	})
	return ret, err
}

// Get retrieves the NetworkAuthentication from the indexer for a given namespace and name.
func (s networkAuthenticationNamespaceLister) Get(name string) (*v1alpha1.NetworkAuthentication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkauthentication"), name)
	}
	return obj.(*v1alpha1.NetworkAuthentication), nil
}