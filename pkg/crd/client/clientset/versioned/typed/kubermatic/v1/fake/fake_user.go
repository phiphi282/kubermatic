// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUsers implements UserInterface
type FakeUsers struct {
	Fake *FakeKubermaticV1
}

var usersResource = schema.GroupVersionResource{Group: "kubermatic.k8s.io", Version: "v1", Resource: "users"}

var usersKind = schema.GroupVersionKind{Group: "kubermatic.k8s.io", Version: "v1", Kind: "User"}

// Get takes name of the user, and returns the corresponding user object, and an error if there is any.
func (c *FakeUsers) Get(name string, options v1.GetOptions) (result *kubermaticv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(usersResource, name), &kubermaticv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.User), err
}

// List takes label and field selectors, and returns the list of Users that match those selectors.
func (c *FakeUsers) List(opts v1.ListOptions) (result *kubermaticv1.UserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(usersResource, usersKind, opts), &kubermaticv1.UserList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.UserList{ListMeta: obj.(*kubermaticv1.UserList).ListMeta}
	for _, item := range obj.(*kubermaticv1.UserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested users.
func (c *FakeUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(usersResource, opts))
}

// Create takes the representation of a user and creates it.  Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Create(user *kubermaticv1.User) (result *kubermaticv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(usersResource, user), &kubermaticv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.User), err
}

// Update takes the representation of a user and updates it. Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Update(user *kubermaticv1.User) (result *kubermaticv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(usersResource, user), &kubermaticv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.User), err
}

// Delete takes name of the user and deletes it. Returns an error if one occurs.
func (c *FakeUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(usersResource, name), &kubermaticv1.User{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(usersResource, listOptions)

	_, err := c.Fake.Invokes(action, &kubermaticv1.UserList{})
	return err
}

// Patch applies the patch and returns the patched user.
func (c *FakeUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kubermaticv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(usersResource, name, pt, data, subresources...), &kubermaticv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.User), err
}