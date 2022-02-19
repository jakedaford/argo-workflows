// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"

	workflowtemplate "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflowtemplate"
)

// WorkflowTemplateServiceClient is an autogenerated mock type for the WorkflowTemplateServiceClient type
type WorkflowTemplateServiceClient struct {
	mock.Mock
}

// CreateWorkflowTemplate provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowTemplateServiceClient) CreateWorkflowTemplate(ctx context.Context, in *workflowtemplate.WorkflowTemplateCreateRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowTemplate, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.WorkflowTemplate
	if rf, ok := ret.Get(0).(func(context.Context, *workflowtemplate.WorkflowTemplateCreateRequest, ...grpc.CallOption) *v1alpha1.WorkflowTemplate); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.WorkflowTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflowtemplate.WorkflowTemplateCreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkflowTemplate provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowTemplateServiceClient) DeleteWorkflowTemplate(ctx context.Context, in *workflowtemplate.WorkflowTemplateDeleteRequest, opts ...grpc.CallOption) (*workflowtemplate.WorkflowTemplateDeleteResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *workflowtemplate.WorkflowTemplateDeleteResponse
	if rf, ok := ret.Get(0).(func(context.Context, *workflowtemplate.WorkflowTemplateDeleteRequest, ...grpc.CallOption) *workflowtemplate.WorkflowTemplateDeleteResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workflowtemplate.WorkflowTemplateDeleteResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflowtemplate.WorkflowTemplateDeleteRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowTemplate provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowTemplateServiceClient) GetWorkflowTemplate(ctx context.Context, in *workflowtemplate.WorkflowTemplateGetRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowTemplate, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.WorkflowTemplate
	if rf, ok := ret.Get(0).(func(context.Context, *workflowtemplate.WorkflowTemplateGetRequest, ...grpc.CallOption) *v1alpha1.WorkflowTemplate); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.WorkflowTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflowtemplate.WorkflowTemplateGetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LintWorkflowTemplate provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowTemplateServiceClient) LintWorkflowTemplate(ctx context.Context, in *workflowtemplate.WorkflowTemplateLintRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowTemplate, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.WorkflowTemplate
	if rf, ok := ret.Get(0).(func(context.Context, *workflowtemplate.WorkflowTemplateLintRequest, ...grpc.CallOption) *v1alpha1.WorkflowTemplate); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.WorkflowTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflowtemplate.WorkflowTemplateLintRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkflowTemplates provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowTemplateServiceClient) ListWorkflowTemplates(ctx context.Context, in *workflowtemplate.WorkflowTemplateListRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowTemplateList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.WorkflowTemplateList
	if rf, ok := ret.Get(0).(func(context.Context, *workflowtemplate.WorkflowTemplateListRequest, ...grpc.CallOption) *v1alpha1.WorkflowTemplateList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.WorkflowTemplateList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflowtemplate.WorkflowTemplateListRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWorkflowTemplate provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowTemplateServiceClient) UpdateWorkflowTemplate(ctx context.Context, in *workflowtemplate.WorkflowTemplateUpdateRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowTemplate, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.WorkflowTemplate
	if rf, ok := ret.Get(0).(func(context.Context, *workflowtemplate.WorkflowTemplateUpdateRequest, ...grpc.CallOption) *v1alpha1.WorkflowTemplate); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.WorkflowTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflowtemplate.WorkflowTemplateUpdateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
