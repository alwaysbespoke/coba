/*
Copyright 2023 Always Bespoke LLC

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

package v1

import (
	v1 "github.com/alwaysbespoke/coba/pkg/crds/sip/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SIPMessageLister helps list SIPMessages.
// All objects returned here must be treated as read-only.
type SIPMessageLister interface {
	// List lists all SIPMessages in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.SIPMessage, err error)
	// SIPMessages returns an object that can list and get SIPMessages.
	SIPMessages(namespace string) SIPMessageNamespaceLister
	SIPMessageListerExpansion
}

// sIPMessageLister implements the SIPMessageLister interface.
type sIPMessageLister struct {
	indexer cache.Indexer
}

// NewSIPMessageLister returns a new SIPMessageLister.
func NewSIPMessageLister(indexer cache.Indexer) SIPMessageLister {
	return &sIPMessageLister{indexer: indexer}
}

// List lists all SIPMessages in the indexer.
func (s *sIPMessageLister) List(selector labels.Selector) (ret []*v1.SIPMessage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SIPMessage))
	})
	return ret, err
}

// SIPMessages returns an object that can list and get SIPMessages.
func (s *sIPMessageLister) SIPMessages(namespace string) SIPMessageNamespaceLister {
	return sIPMessageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SIPMessageNamespaceLister helps list and get SIPMessages.
// All objects returned here must be treated as read-only.
type SIPMessageNamespaceLister interface {
	// List lists all SIPMessages in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.SIPMessage, err error)
	// Get retrieves the SIPMessage from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.SIPMessage, error)
	SIPMessageNamespaceListerExpansion
}

// sIPMessageNamespaceLister implements the SIPMessageNamespaceLister
// interface.
type sIPMessageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SIPMessages in the indexer for a given namespace.
func (s sIPMessageNamespaceLister) List(selector labels.Selector) (ret []*v1.SIPMessage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SIPMessage))
	})
	return ret, err
}

// Get retrieves the SIPMessage from the indexer for a given namespace and name.
func (s sIPMessageNamespaceLister) Get(name string) (*v1.SIPMessage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("sipmessage"), name)
	}
	return obj.(*v1.SIPMessage), nil
}
