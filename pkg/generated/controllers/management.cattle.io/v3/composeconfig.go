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
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
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

type ComposeConfigHandler func(string, *v3.ComposeConfig) (*v3.ComposeConfig, error)

type ComposeConfigController interface {
	generic.ControllerMeta
	ComposeConfigClient

	OnChange(ctx context.Context, name string, sync ComposeConfigHandler)
	OnRemove(ctx context.Context, name string, sync ComposeConfigHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() ComposeConfigCache
}

type ComposeConfigClient interface {
	Create(*v3.ComposeConfig) (*v3.ComposeConfig, error)
	Update(*v3.ComposeConfig) (*v3.ComposeConfig, error)
	UpdateStatus(*v3.ComposeConfig) (*v3.ComposeConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.ComposeConfig, error)
	List(opts metav1.ListOptions) (*v3.ComposeConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ComposeConfig, err error)
}

type ComposeConfigCache interface {
	Get(name string) (*v3.ComposeConfig, error)
	List(selector labels.Selector) ([]*v3.ComposeConfig, error)

	AddIndexer(indexName string, indexer ComposeConfigIndexer)
	GetByIndex(indexName, key string) ([]*v3.ComposeConfig, error)
}

type ComposeConfigIndexer func(obj *v3.ComposeConfig) ([]string, error)

type composeConfigController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewComposeConfigController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ComposeConfigController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &composeConfigController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromComposeConfigHandlerToHandler(sync ComposeConfigHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.ComposeConfig
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.ComposeConfig))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *composeConfigController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.ComposeConfig))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateComposeConfigDeepCopyOnChange(client ComposeConfigClient, obj *v3.ComposeConfig, handler func(obj *v3.ComposeConfig) (*v3.ComposeConfig, error)) (*v3.ComposeConfig, error) {
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

func (c *composeConfigController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *composeConfigController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *composeConfigController) OnChange(ctx context.Context, name string, sync ComposeConfigHandler) {
	c.AddGenericHandler(ctx, name, FromComposeConfigHandlerToHandler(sync))
}

func (c *composeConfigController) OnRemove(ctx context.Context, name string, sync ComposeConfigHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromComposeConfigHandlerToHandler(sync)))
}

func (c *composeConfigController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *composeConfigController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *composeConfigController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *composeConfigController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *composeConfigController) Cache() ComposeConfigCache {
	return &composeConfigCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *composeConfigController) Create(obj *v3.ComposeConfig) (*v3.ComposeConfig, error) {
	result := &v3.ComposeConfig{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *composeConfigController) Update(obj *v3.ComposeConfig) (*v3.ComposeConfig, error) {
	result := &v3.ComposeConfig{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *composeConfigController) UpdateStatus(obj *v3.ComposeConfig) (*v3.ComposeConfig, error) {
	result := &v3.ComposeConfig{}
	return result, c.client.UpdateStatus(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *composeConfigController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *composeConfigController) Get(name string, options metav1.GetOptions) (*v3.ComposeConfig, error) {
	result := &v3.ComposeConfig{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *composeConfigController) List(opts metav1.ListOptions) (*v3.ComposeConfigList, error) {
	result := &v3.ComposeConfigList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *composeConfigController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *composeConfigController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.ComposeConfig, error) {
	result := &v3.ComposeConfig{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type composeConfigCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *composeConfigCache) Get(name string) (*v3.ComposeConfig, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.ComposeConfig), nil
}

func (c *composeConfigCache) List(selector labels.Selector) (ret []*v3.ComposeConfig, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ComposeConfig))
	})

	return ret, err
}

func (c *composeConfigCache) AddIndexer(indexName string, indexer ComposeConfigIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.ComposeConfig))
		},
	}))
}

func (c *composeConfigCache) GetByIndex(indexName, key string) (result []*v3.ComposeConfig, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.ComposeConfig, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.ComposeConfig))
	}
	return result, nil
}

type ComposeConfigStatusHandler func(obj *v3.ComposeConfig, status v3.ComposeStatus) (v3.ComposeStatus, error)

type ComposeConfigGeneratingHandler func(obj *v3.ComposeConfig, status v3.ComposeStatus) ([]runtime.Object, v3.ComposeStatus, error)

func RegisterComposeConfigStatusHandler(ctx context.Context, controller ComposeConfigController, condition condition.Cond, name string, handler ComposeConfigStatusHandler) {
	statusHandler := &composeConfigStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromComposeConfigHandlerToHandler(statusHandler.sync))
}

func RegisterComposeConfigGeneratingHandler(ctx context.Context, controller ComposeConfigController, apply apply.Apply,
	condition condition.Cond, name string, handler ComposeConfigGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &composeConfigGeneratingHandler{
		ComposeConfigGeneratingHandler: handler,
		apply:                          apply,
		name:                           name,
		gvk:                            controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterComposeConfigStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type composeConfigStatusHandler struct {
	client    ComposeConfigClient
	condition condition.Cond
	handler   ComposeConfigStatusHandler
}

func (a *composeConfigStatusHandler) sync(key string, obj *v3.ComposeConfig) (*v3.ComposeConfig, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type composeConfigGeneratingHandler struct {
	ComposeConfigGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *composeConfigGeneratingHandler) Remove(key string, obj *v3.ComposeConfig) (*v3.ComposeConfig, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.ComposeConfig{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *composeConfigGeneratingHandler) Handle(obj *v3.ComposeConfig, status v3.ComposeStatus) (v3.ComposeStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.ComposeConfigGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
