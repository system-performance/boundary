// Code generated by "make api"; DO NOT EDIT.
package targets

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/kr/pretty"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/scopes"
)

type Target struct {
	Id          string            `json:"id,omitempty"`
	ScopeId     string            `json:"scope_id,omitempty"`
	Scope       *scopes.ScopeInfo `json:"scope,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	CreatedTime time.Time         `json:"created_time,omitempty"`
	UpdatedTime time.Time         `json:"updated_time,omitempty"`
	Version     uint32            `json:"version,omitempty"`
	Type        string            `json:"type,omitempty"`
	HostSetIds  []string          `json:"host_set_ids,omitempty"`
	HostSets    []*HostSet        `json:"host_sets,omitempty"`
	DefaultPort uint32            `json:"default_port,omitempty"`
}

type Client struct {
	client *api.Client
}

func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}

func (c *Client) Create2(ctx context.Context, resourceType string, scopeId string, opt ...Option) (*Target, *api.Error, error) {
	if scopeId == "" {
		return nil, nil, fmt.Errorf("empty scopeId value passed into Create request")
	}
	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}
	if resourceType == "" {
		return nil, nil, fmt.Errorf("empty resourceType value passed into Create request")
	} else {
		opts.postMap["type"] = resourceType
	}

	opts.postMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "POST", "targets", opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(Target)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) Read2(ctx context.Context, targetId string, opt ...Option) (*Target, *api.Error, error) {
	if targetId == "" {
		return nil, nil, fmt.Errorf("empty  targetId value passed into Read request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("targets/%s", targetId), nil, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(Target)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) Update2(ctx context.Context, targetId string, version uint32, opt ...Option) (*Target, *api.Error, error) {
	if targetId == "" {
		return nil, nil, fmt.Errorf("empty targetId value passed into Update request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingApiErr, existingErr := c.Read2(ctx, targetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("targets/%s", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(Target)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) Delete2(ctx context.Context, targetId string, opt ...Option) (bool, *api.Error, error) {
	if targetId == "" {
		return false, nil, fmt.Errorf("empty targetId value passed into Delete request")
	}
	if c.client == nil {
		return false, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("targets/%s", targetId), nil, apiOpts...)
	if err != nil {
		return false, nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return false, nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	type deleteResponse struct {
		Existed bool
	}
	target := &deleteResponse{}
	apiErr, err := resp.Decode(target)
	if err != nil {
		return false, nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return false, apiErr, nil
	}
	return target.Existed, apiErr, nil
}

func (c *Client) List2(ctx context.Context, scopeId string, opt ...Option) ([]*Target, *api.Error, error) {
	if scopeId == "" {
		return nil, nil, fmt.Errorf("empty scopeId value passed into List request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())
	opts.queryMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "GET", "targets", nil, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	type listResponse struct {
		Items []*Target
	}
	target := &listResponse{}
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target.Items, apiErr, nil
}

func (c *Client) AddHostSets2(ctx context.Context, targetId string, version uint32, hostSetIds []string, opt ...Option) (*Target, *api.Error, error) {
	if targetId == "" {
		return nil, nil, fmt.Errorf("empty targetId value passed into AddHostSets request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into AddHostSets request")
		}
		existingTarget, existingApiErr, existingErr := c.Read2(ctx, targetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.postMap["version"] = version

	if len(hostSetIds) > 0 {
		opts.postMap["host_set_ids"] = hostSetIds
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("targets/%s:add-host-sets", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating AddHostSets request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during AddHostSets call: %w", err)
	}

	target := new(Target)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding AddHostSets response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) SetHostSets2(ctx context.Context, targetId string, version uint32, hostSetIds []string, opt ...Option) (*Target, *api.Error, error) {
	if targetId == "" {
		return nil, nil, fmt.Errorf("empty targetId value passed into SetHostSets request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into SetHostSets request")
		}
		existingTarget, existingApiErr, existingErr := c.Read2(ctx, targetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.postMap["version"] = version

	if len(hostSetIds) > 0 {
		opts.postMap["host_set_ids"] = hostSetIds
	} else if hostSetIds != nil {
		// In this function, a non-nil but empty list means clear out
		opts.postMap["host_set_ids"] = nil
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("targets/%s:set-host-sets", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating SetHostSets request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during SetHostSets call: %w", err)
	}

	target := new(Target)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding SetHostSets response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) RemoveHostSets2(ctx context.Context, targetId string, version uint32, hostSetIds []string, opt ...Option) (*Target, *api.Error, error) {
	if targetId == "" {
		return nil, nil, fmt.Errorf("empty targetId value passed into RemoveHostSets request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into RemoveHostSets request")
		}
		existingTarget, existingApiErr, existingErr := c.Read2(ctx, targetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.postMap["version"] = version

	if len(hostSetIds) > 0 {
		opts.postMap["host_set_ids"] = hostSetIds
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("targets/%s:remove-host-sets", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating RemoveHostSets request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during RemoveHostSets call: %w", err)
	}

	target := new(Target)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding RemoveHostSets response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}