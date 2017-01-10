package charta

import (
	"net/http"
	"reflect"
	"testing"

	"k8s.io/client-go/kubernetes"
)

func TestGetInternalCV(t *testing.T) {
	tests := []struct {
		name string
		want *[]InternalCV
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInternalCV(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInternalCV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetInternalCV(t *testing.T) {
	type args struct {
		icv *[]InternalCV
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetInternalCV(tt.args.icv)
		})
	}
}

func TestGetKubeClient(t *testing.T) {
	tests := []struct {
		name string
		want *kubernetes.Clientset
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKubeClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKubeClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetKubeClient(t *testing.T) {
	type args struct {
		client *kubernetes.Clientset
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetKubeClient(tt.args.client)
		})
	}
}

func TestGetClusterView(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetClusterView(tt.args.w, tt.args.req)
		})
	}
}

func TestGetNamespaces(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNamespaces(tt.args.w, tt.args.req)
		})
	}
}

func TestGetNamespaceView(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNamespaceView(tt.args.w, tt.args.req)
		})
	}
}
