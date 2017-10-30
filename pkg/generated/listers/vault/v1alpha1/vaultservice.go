/*
Copyright 2017 The vault-operator Authors

Commercial software license.
*/

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/coreos/vault-operator/pkg/apis/vault/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VaultServiceLister helps list VaultServices.
type VaultServiceLister interface {
	// List lists all VaultServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VaultService, err error)
	// VaultServices returns an object that can list and get VaultServices.
	VaultServices(namespace string) VaultServiceNamespaceLister
	VaultServiceListerExpansion
}

// vaultServiceLister implements the VaultServiceLister interface.
type vaultServiceLister struct {
	indexer cache.Indexer
}

// NewVaultServiceLister returns a new VaultServiceLister.
func NewVaultServiceLister(indexer cache.Indexer) VaultServiceLister {
	return &vaultServiceLister{indexer: indexer}
}

// List lists all VaultServices in the indexer.
func (s *vaultServiceLister) List(selector labels.Selector) (ret []*v1alpha1.VaultService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultService))
	})
	return ret, err
}

// VaultServices returns an object that can list and get VaultServices.
func (s *vaultServiceLister) VaultServices(namespace string) VaultServiceNamespaceLister {
	return vaultServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VaultServiceNamespaceLister helps list and get VaultServices.
type VaultServiceNamespaceLister interface {
	// List lists all VaultServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VaultService, err error)
	// Get retrieves the VaultService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VaultService, error)
	VaultServiceNamespaceListerExpansion
}

// vaultServiceNamespaceLister implements the VaultServiceNamespaceLister
// interface.
type vaultServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VaultServices in the indexer for a given namespace.
func (s vaultServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VaultService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultService))
	})
	return ret, err
}

// Get retrieves the VaultService from the indexer for a given namespace and name.
func (s vaultServiceNamespaceLister) Get(name string) (*v1alpha1.VaultService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vaultservice"), name)
	}
	return obj.(*v1alpha1.VaultService), nil
}