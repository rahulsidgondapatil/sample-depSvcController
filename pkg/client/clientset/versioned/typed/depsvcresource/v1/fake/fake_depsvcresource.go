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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	DepSvcResourcev1 "github.com/rahulsidgondapatil/sample-depSvcController/pkg/apis/DepSvcResource/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDepSvcResources implements DepSvcResourceInterface
type FakeDepSvcResources struct {
	Fake *FakeDepSvcResourceV1
	ns   string
}

var depsvcresourcesResource = schema.GroupVersionResource{Group: "DepSvcResource.com", Version: "v1", Resource: "depsvcresources"}

var depsvcresourcesKind = schema.GroupVersionKind{Group: "DepSvcResource.com", Version: "v1", Kind: "DepSvcResource"}

// Get takes name of the depSvcResource, and returns the corresponding depSvcResource object, and an error if there is any.
func (c *FakeDepSvcResources) Get(name string, options v1.GetOptions) (result *DepSvcResourcev1.DepSvcResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(depsvcresourcesResource, c.ns, name), &DepSvcResourcev1.DepSvcResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*DepSvcResourcev1.DepSvcResource), err
}

// List takes label and field selectors, and returns the list of DepSvcResources that match those selectors.
func (c *FakeDepSvcResources) List(opts v1.ListOptions) (result *DepSvcResourcev1.DepSvcResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(depsvcresourcesResource, depsvcresourcesKind, c.ns, opts), &DepSvcResourcev1.DepSvcResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &DepSvcResourcev1.DepSvcResourceList{ListMeta: obj.(*DepSvcResourcev1.DepSvcResourceList).ListMeta}
	for _, item := range obj.(*DepSvcResourcev1.DepSvcResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested depSvcResources.
func (c *FakeDepSvcResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(depsvcresourcesResource, c.ns, opts))

}

// Create takes the representation of a depSvcResource and creates it.  Returns the server's representation of the depSvcResource, and an error, if there is any.
func (c *FakeDepSvcResources) Create(depSvcResource *DepSvcResourcev1.DepSvcResource) (result *DepSvcResourcev1.DepSvcResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(depsvcresourcesResource, c.ns, depSvcResource), &DepSvcResourcev1.DepSvcResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*DepSvcResourcev1.DepSvcResource), err
}

// Update takes the representation of a depSvcResource and updates it. Returns the server's representation of the depSvcResource, and an error, if there is any.
func (c *FakeDepSvcResources) Update(depSvcResource *DepSvcResourcev1.DepSvcResource) (result *DepSvcResourcev1.DepSvcResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(depsvcresourcesResource, c.ns, depSvcResource), &DepSvcResourcev1.DepSvcResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*DepSvcResourcev1.DepSvcResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDepSvcResources) UpdateStatus(depSvcResource *DepSvcResourcev1.DepSvcResource) (*DepSvcResourcev1.DepSvcResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(depsvcresourcesResource, "status", c.ns, depSvcResource), &DepSvcResourcev1.DepSvcResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*DepSvcResourcev1.DepSvcResource), err
}

// Delete takes name of the depSvcResource and deletes it. Returns an error if one occurs.
func (c *FakeDepSvcResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(depsvcresourcesResource, c.ns, name), &DepSvcResourcev1.DepSvcResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDepSvcResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(depsvcresourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &DepSvcResourcev1.DepSvcResourceList{})
	return err
}

// Patch applies the patch and returns the patched depSvcResource.
func (c *FakeDepSvcResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *DepSvcResourcev1.DepSvcResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(depsvcresourcesResource, c.ns, name, pt, data, subresources...), &DepSvcResourcev1.DepSvcResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*DepSvcResourcev1.DepSvcResource), err
}
