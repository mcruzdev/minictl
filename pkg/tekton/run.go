package tekton

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RunPipeline(application, org string) {
	var pr *v1beta1.PipelineRun

	objMeta := metav1.ObjectMeta{
		Namespace: "tekton-pipelines",
	}
	pr = &v1beta1.PipelineRun{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "tekton.dev/v1beta1",
			Kind:       "PipelineRun",
		},
		ObjectMeta: objMeta,
		Spec: v1beta1.PipelineRunSpec{
			PipelineRef: &v1beta1.PipelineRef{Name: "create-version"},
			Params: v1beta1.Params{
				{
					Name: "application",
					Value: v1beta1.ParamValue{
						Type:      v1beta1.ParamTypeString,
						StringVal: application,
					},
				},
				{
					Name: "org",
					Value: v1beta1.ParamValue{
						Type:      v1beta1.ParamTypeString,
						StringVal: org,
					},
				},
			},
		},
	}

	payload, err := json.Marshal(pr)
	if err != nil {
		log.Fatal("Error while marshalling the PipelineRun request")
	}

	url := fmt.Sprintf("https://%s/api/v1/namespaces/tekton-pipelines/pipelineruns", os.Getenv("MINI_PLATFORM_CLUSTER_IP"))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("Error while creating request to mini-platform cluster")
	}

	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Error while requesting mini-platform cluster. %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Println(fmt.Sprintf("Was not possible to create the pipelinerun resource, status code %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error while reading response body")
	}

	log.Println(string(body))

}
