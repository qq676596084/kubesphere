// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "istio.io/client-go/pkg/apis/rbac/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterRbacConfigLister helps list ClusterRbacConfigs.
type ClusterRbacConfigLister interface {
	// List lists all ClusterRbacConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterRbacConfig, err error)
	// Get retrieves the ClusterRbacConfig from the index for a given name.
	Get(name string) (*v1alpha1.ClusterRbacConfig, error)
	ClusterRbacConfigListerExpansion
}

// clusterRbacConfigLister implements the ClusterRbacConfigLister interface.
type clusterRbacConfigLister struct {
	indexer cache.Indexer
}

// NewClusterRbacConfigLister returns a new ClusterRbacConfigLister.
func NewClusterRbacConfigLister(indexer cache.Indexer) ClusterRbacConfigLister {
	return &clusterRbacConfigLister{indexer: indexer}
}

// List lists all ClusterRbacConfigs in the indexer.
func (s *clusterRbacConfigLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterRbacConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterRbacConfig))
	})
	return ret, err
}

// Get retrieves the ClusterRbacConfig from the index for a given name.
func (s *clusterRbacConfigLister) Get(name string) (*v1alpha1.ClusterRbacConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterrbacconfig"), name)
	}
	return obj.(*v1alpha1.ClusterRbacConfig), nil
}
