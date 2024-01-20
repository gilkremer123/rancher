/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/project.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type PipelineSettingHandler func(string, *v3.PipelineSetting) (*v3.PipelineSetting, error)

type PipelineSettingController interface {
	generic.ControllerMeta
	PipelineSettingClient

	OnChange(ctx context.Context, name string, sync PipelineSettingHandler)
	OnRemove(ctx context.Context, name string, sync PipelineSettingHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() PipelineSettingCache
}

type PipelineSettingClient interface {
	Create(*v3.PipelineSetting) (*v3.PipelineSetting, error)
	Update(*v3.PipelineSetting) (*v3.PipelineSetting, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.PipelineSetting, error)
	List(namespace string, opts metav1.ListOptions) (*v3.PipelineSettingList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.PipelineSetting, err error)
}

type PipelineSettingCache interface {
	Get(namespace, name string) (*v3.PipelineSetting, error)
	List(namespace string, selector labels.Selector) ([]*v3.PipelineSetting, error)

	AddIndexer(indexName string, indexer PipelineSettingIndexer)
	GetByIndex(indexName, key string) ([]*v3.PipelineSetting, error)
}

type PipelineSettingIndexer func(obj *v3.PipelineSetting) ([]string, error)

type pipelineSettingController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewPipelineSettingController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) PipelineSettingController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &pipelineSettingController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromPipelineSettingHandlerToHandler(sync PipelineSettingHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.PipelineSetting
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.PipelineSetting))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *pipelineSettingController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.PipelineSetting))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdatePipelineSettingDeepCopyOnChange(client PipelineSettingClient, obj *v3.PipelineSetting, handler func(obj *v3.PipelineSetting) (*v3.PipelineSetting, error)) (*v3.PipelineSetting, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *pipelineSettingController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *pipelineSettingController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *pipelineSettingController) OnChange(ctx context.Context, name string, sync PipelineSettingHandler) {
	c.AddGenericHandler(ctx, name, FromPipelineSettingHandlerToHandler(sync))
}

func (c *pipelineSettingController) OnRemove(ctx context.Context, name string, sync PipelineSettingHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromPipelineSettingHandlerToHandler(sync)))
}

func (c *pipelineSettingController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *pipelineSettingController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *pipelineSettingController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *pipelineSettingController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *pipelineSettingController) Cache() PipelineSettingCache {
	return &pipelineSettingCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *pipelineSettingController) Create(obj *v3.PipelineSetting) (*v3.PipelineSetting, error) {
	result := &v3.PipelineSetting{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *pipelineSettingController) Update(obj *v3.PipelineSetting) (*v3.PipelineSetting, error) {
	result := &v3.PipelineSetting{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *pipelineSettingController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *pipelineSettingController) Get(namespace, name string, options metav1.GetOptions) (*v3.PipelineSetting, error) {
	result := &v3.PipelineSetting{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *pipelineSettingController) List(namespace string, opts metav1.ListOptions) (*v3.PipelineSettingList, error) {
	result := &v3.PipelineSettingList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *pipelineSettingController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *pipelineSettingController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.PipelineSetting, error) {
	result := &v3.PipelineSetting{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type pipelineSettingCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *pipelineSettingCache) Get(namespace, name string) (*v3.PipelineSetting, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.PipelineSetting), nil
}

func (c *pipelineSettingCache) List(namespace string, selector labels.Selector) (ret []*v3.PipelineSetting, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.PipelineSetting))
	})

	return ret, err
}

func (c *pipelineSettingCache) AddIndexer(indexName string, indexer PipelineSettingIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.PipelineSetting))
		},
	}))
}

func (c *pipelineSettingCache) GetByIndex(indexName, key string) (result []*v3.PipelineSetting, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.PipelineSetting, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.PipelineSetting))
	}
	return result, nil
}
