package workflow

import (
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	wfclientset "github.com/argoproj/argo/pkg/client/clientset/versioned"
	"github.com/argoproj/argo/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"github.com/argoproj/argo/workflow/common"
	"github.com/argoproj/pkg/json"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	v1alpha1.WorkflowInterface
}

func NewClient(namespace string, configPath ...string) (client *Client, err error) {
	var config *rest.Config
	if len(configPath) == 0 {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", configPath[0])
	}
	if err != nil {
		return
	}

	wfclient := wfclientset.NewForConfigOrDie(config).ArgoprojV1alpha1().Workflows(namespace)
	client = &Client{WorkflowInterface: wfclient}

	return
}

func unmarshalWorkflows(wfBytes []byte, strict bool) (wfs []wfv1.Workflow, err error) {
	var wf wfv1.Workflow
	var jsonOpts []json.JSONOpt
	if strict {
		jsonOpts = append(jsonOpts, json.DisallowUnknownFields)
	}
	err = json.Unmarshal(wfBytes, &wf, jsonOpts...)
	if err == nil {
		return []wfv1.Workflow{wf}, nil
	}
	wfs, err = common.SplitWorkflowYAMLFile(wfBytes, strict)
	if err == nil {
		return
	}

	return
}

func (c *Client) Submit(wfBytes []byte, strict bool) (err error) {
	workflows, err := unmarshalWorkflows(wfBytes, strict)
	if err == nil {
		return err
	}

	for _, wf := range workflows {
		if err != nil {
			return err
		}

		_, err = c.Create(&wf)
		if err != nil {
			return err
		}
	}

	return nil
}
