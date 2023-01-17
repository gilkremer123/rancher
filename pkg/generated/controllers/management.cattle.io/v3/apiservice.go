/*
Copyright 2023 Rancher Labs, Inc.

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

type APIServiceHandler func(string, *v3.APIService) (*v3.APIService, error)

type APIServiceController interface {
	generic.ControllerMeta
	APIServiceClient

	OnChange(ctx context.Context, name string, sync APIServiceHandler)
	OnRemove(ctx context.Context, name string, sync APIServiceHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() APIServiceCache
}

type APIServiceClient interface {
	Create(*v3.APIService) (*v3.APIService, error)
	Update(*v3.APIService) (*v3.APIService, error)
	UpdateStatus(*v3.APIService) (*v3.APIService, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.APIService, error)
	List(opts metav1.ListOptions) (*v3.APIServiceList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.APIService, err error)
}

type APIServiceCache interface {
	Get(name string) (*v3.APIService, error)
	List(selector labels.Selector) ([]*v3.APIService, error)

	AddIndexer(indexName string, indexer APIServiceIndexer)
	GetByIndex(indexName, key string) ([]*v3.APIService, error)
}

type APIServiceIndexer func(obj *v3.APIService) ([]string, error)

type aPIServiceController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewAPIServiceController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) APIServiceController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &aPIServiceController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromAPIServiceHandlerToHandler(sync APIServiceHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.APIService
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.APIService))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *aPIServiceController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.APIService))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateAPIServiceDeepCopyOnChange(client APIServiceClient, obj *v3.APIService, handler func(obj *v3.APIService) (*v3.APIService, error)) (*v3.APIService, error) {
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

func (c *aPIServiceController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *aPIServiceController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *aPIServiceController) OnChange(ctx context.Context, name string, sync APIServiceHandler) {
	c.AddGenericHandler(ctx, name, FromAPIServiceHandlerToHandler(sync))
}

func (c *aPIServiceController) OnRemove(ctx context.Context, name string, sync APIServiceHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromAPIServiceHandlerToHandler(sync)))
}

func (c *aPIServiceController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *aPIServiceController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *aPIServiceController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *aPIServiceController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *aPIServiceController) Cache() APIServiceCache {
	return &aPIServiceCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *aPIServiceController) Create(obj *v3.APIService) (*v3.APIService, error) {
	result := &v3.APIService{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *aPIServiceController) Update(obj *v3.APIService) (*v3.APIService, error) {
	result := &v3.APIService{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *aPIServiceController) UpdateStatus(obj *v3.APIService) (*v3.APIService, error) {
	result := &v3.APIService{}
	return result, c.client.UpdateStatus(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *aPIServiceController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *aPIServiceController) Get(name string, options metav1.GetOptions) (*v3.APIService, error) {
	result := &v3.APIService{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *aPIServiceController) List(opts metav1.ListOptions) (*v3.APIServiceList, error) {
	result := &v3.APIServiceList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *aPIServiceController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *aPIServiceController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.APIService, error) {
	result := &v3.APIService{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type aPIServiceCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *aPIServiceCache) Get(name string) (*v3.APIService, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.APIService), nil
}

func (c *aPIServiceCache) List(selector labels.Selector) (ret []*v3.APIService, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.APIService))
	})

	return ret, err
}

func (c *aPIServiceCache) AddIndexer(indexName string, indexer APIServiceIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.APIService))
		},
	}))
}

func (c *aPIServiceCache) GetByIndex(indexName, key string) (result []*v3.APIService, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.APIService, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.APIService))
	}
	return result, nil
}

type APIServiceStatusHandler func(obj *v3.APIService, status v3.APIServiceStatus) (v3.APIServiceStatus, error)

type APIServiceGeneratingHandler func(obj *v3.APIService, status v3.APIServiceStatus) ([]runtime.Object, v3.APIServiceStatus, error)

func RegisterAPIServiceStatusHandler(ctx context.Context, controller APIServiceController, condition condition.Cond, name string, handler APIServiceStatusHandler) {
	statusHandler := &aPIServiceStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromAPIServiceHandlerToHandler(statusHandler.sync))
}

func RegisterAPIServiceGeneratingHandler(ctx context.Context, controller APIServiceController, apply apply.Apply,
	condition condition.Cond, name string, handler APIServiceGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &aPIServiceGeneratingHandler{
		APIServiceGeneratingHandler: handler,
		apply:                       apply,
		name:                        name,
		gvk:                         controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterAPIServiceStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type aPIServiceStatusHandler struct {
	client    APIServiceClient
	condition condition.Cond
	handler   APIServiceStatusHandler
}

func (a *aPIServiceStatusHandler) sync(key string, obj *v3.APIService) (*v3.APIService, error) {
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

type aPIServiceGeneratingHandler struct {
	APIServiceGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *aPIServiceGeneratingHandler) Remove(key string, obj *v3.APIService) (*v3.APIService, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.APIService{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *aPIServiceGeneratingHandler) Handle(obj *v3.APIService, status v3.APIServiceStatus) (v3.APIServiceStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.APIServiceGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
