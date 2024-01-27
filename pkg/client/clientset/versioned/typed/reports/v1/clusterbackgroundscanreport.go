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

package v1

import (
	"context"
	"time"

	v1 "github.com/kyverno/kyverno/api/reports/v1"
	scheme "github.com/kyverno/kyverno/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterBackgroundScanReportsGetter has a method to return a ClusterBackgroundScanReportInterface.
// A group's client should implement this interface.
type ClusterBackgroundScanReportsGetter interface {
	ClusterBackgroundScanReports() ClusterBackgroundScanReportInterface
}

// ClusterBackgroundScanReportInterface has methods to work with ClusterBackgroundScanReport resources.
type ClusterBackgroundScanReportInterface interface {
	Create(ctx context.Context, clusterBackgroundScanReport *v1.ClusterBackgroundScanReport, opts metav1.CreateOptions) (*v1.ClusterBackgroundScanReport, error)
	Update(ctx context.Context, clusterBackgroundScanReport *v1.ClusterBackgroundScanReport, opts metav1.UpdateOptions) (*v1.ClusterBackgroundScanReport, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterBackgroundScanReport, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterBackgroundScanReportList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterBackgroundScanReport, err error)
	ClusterBackgroundScanReportExpansion
}

// clusterBackgroundScanReports implements ClusterBackgroundScanReportInterface
type clusterBackgroundScanReports struct {
	client rest.Interface
}

// newClusterBackgroundScanReports returns a ClusterBackgroundScanReports
func newClusterBackgroundScanReports(c *ReportsV1Client) *clusterBackgroundScanReports {
	return &clusterBackgroundScanReports{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterBackgroundScanReport, and returns the corresponding clusterBackgroundScanReport object, and an error if there is any.
func (c *clusterBackgroundScanReports) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterBackgroundScanReport, err error) {
	result = &v1.ClusterBackgroundScanReport{}
	err = c.client.Get().
		Resource("clusterbackgroundscanreports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterBackgroundScanReports that match those selectors.
func (c *clusterBackgroundScanReports) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterBackgroundScanReportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterBackgroundScanReportList{}
	err = c.client.Get().
		Resource("clusterbackgroundscanreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterBackgroundScanReports.
func (c *clusterBackgroundScanReports) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterbackgroundscanreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterBackgroundScanReport and creates it.  Returns the server's representation of the clusterBackgroundScanReport, and an error, if there is any.
func (c *clusterBackgroundScanReports) Create(ctx context.Context, clusterBackgroundScanReport *v1.ClusterBackgroundScanReport, opts metav1.CreateOptions) (result *v1.ClusterBackgroundScanReport, err error) {
	result = &v1.ClusterBackgroundScanReport{}
	err = c.client.Post().
		Resource("clusterbackgroundscanreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterBackgroundScanReport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterBackgroundScanReport and updates it. Returns the server's representation of the clusterBackgroundScanReport, and an error, if there is any.
func (c *clusterBackgroundScanReports) Update(ctx context.Context, clusterBackgroundScanReport *v1.ClusterBackgroundScanReport, opts metav1.UpdateOptions) (result *v1.ClusterBackgroundScanReport, err error) {
	result = &v1.ClusterBackgroundScanReport{}
	err = c.client.Put().
		Resource("clusterbackgroundscanreports").
		Name(clusterBackgroundScanReport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterBackgroundScanReport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterBackgroundScanReport and deletes it. Returns an error if one occurs.
func (c *clusterBackgroundScanReports) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterbackgroundscanreports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterBackgroundScanReports) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterbackgroundscanreports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterBackgroundScanReport.
func (c *clusterBackgroundScanReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterBackgroundScanReport, err error) {
	result = &v1.ClusterBackgroundScanReport{}
	err = c.client.Patch(pt).
		Resource("clusterbackgroundscanreports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}