// This file was automatically generated by lister-gen

package internalversion

import (
	machine "github.com/gardener/node-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AzureMachineClassLister helps list AzureMachineClasses.
type AzureMachineClassLister interface {
	// List lists all AzureMachineClasses in the indexer.
	List(selector labels.Selector) (ret []*machine.AzureMachineClass, err error)
	// Get retrieves the AzureMachineClass from the index for a given name.
	Get(name string) (*machine.AzureMachineClass, error)
	AzureMachineClassListerExpansion
}

// azureMachineClassLister implements the AzureMachineClassLister interface.
type azureMachineClassLister struct {
	indexer cache.Indexer
}

// NewAzureMachineClassLister returns a new AzureMachineClassLister.
func NewAzureMachineClassLister(indexer cache.Indexer) AzureMachineClassLister {
	return &azureMachineClassLister{indexer: indexer}
}

// List lists all AzureMachineClasses in the indexer.
func (s *azureMachineClassLister) List(selector labels.Selector) (ret []*machine.AzureMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.AzureMachineClass))
	})
	return ret, err
}

// Get retrieves the AzureMachineClass from the index for a given name.
func (s *azureMachineClassLister) Get(name string) (*machine.AzureMachineClass, error) {
	key := &machine.AzureMachineClass{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("azuremachineclass"), name)
	}
	return obj.(*machine.AzureMachineClass), nil
}
