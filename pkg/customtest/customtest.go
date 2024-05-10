package customtest

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

var _ framework.PreFilterPlugin = &CustomTest{}

const (
	Name = "CustomTest"
)

type CustomTest struct {
	fh framework.Handle
}

// Name implements framework.Plugin.
func (ct *CustomTest) Name() string {
	return Name
}

func (ct *CustomTest) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) (*framework.PreFilterResult, *framework.Status) {
	klog.V(3).Info("Hello, custom test scheduler!")
	return nil, framework.NewStatus(framework.Success, "")
}

func (ct *CustomTest) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}

func New(_ context.Context, obj runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	ct := CustomTest{
		fh: handle,
	}
	return &ct, nil
}
