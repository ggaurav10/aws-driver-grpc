// Code generated by lister-gen. DO NOT EDIT.

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/gardener/aws-driver-grpc/pkg/apis/machine/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AWSMachineClassLister helps list AWSMachineClasses.
type AWSMachineClassLister interface {
	// List lists all AWSMachineClasses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AWSMachineClass, err error)
	// AWSMachineClasses returns an object that can list and get AWSMachineClasses.
	AWSMachineClasses(namespace string) AWSMachineClassNamespaceLister
	AWSMachineClassListerExpansion
}

// aWSMachineClassLister implements the AWSMachineClassLister interface.
type aWSMachineClassLister struct {
	indexer cache.Indexer
}

// NewAWSMachineClassLister returns a new AWSMachineClassLister.
func NewAWSMachineClassLister(indexer cache.Indexer) AWSMachineClassLister {
	return &aWSMachineClassLister{indexer: indexer}
}

// List lists all AWSMachineClasses in the indexer.
func (s *aWSMachineClassLister) List(selector labels.Selector) (ret []*v1alpha1.AWSMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSMachineClass))
	})
	return ret, err
}

// AWSMachineClasses returns an object that can list and get AWSMachineClasses.
func (s *aWSMachineClassLister) AWSMachineClasses(namespace string) AWSMachineClassNamespaceLister {
	return aWSMachineClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AWSMachineClassNamespaceLister helps list and get AWSMachineClasses.
type AWSMachineClassNamespaceLister interface {
	// List lists all AWSMachineClasses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AWSMachineClass, err error)
	// Get retrieves the AWSMachineClass from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AWSMachineClass, error)
	AWSMachineClassNamespaceListerExpansion
}

// aWSMachineClassNamespaceLister implements the AWSMachineClassNamespaceLister
// interface.
type aWSMachineClassNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AWSMachineClasses in the indexer for a given namespace.
func (s aWSMachineClassNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AWSMachineClass, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSMachineClass))
	})
	return ret, err
}

// Get retrieves the AWSMachineClass from the indexer for a given namespace and name.
func (s aWSMachineClassNamespaceLister) Get(name string) (*v1alpha1.AWSMachineClass, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsmachineclass"), name)
	}
	return obj.(*v1alpha1.AWSMachineClass), nil
}
