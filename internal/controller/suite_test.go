/*
Copyright 2024.

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

package controller_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	ec2operatorv1alpha1 "github.com/cloud-infra-group/ec2-operator/api/v1alpha1"
)

var (
	k8sClient  client.Client        // Shared Kubernetes client for tests
	k8sManager ctrl.Manager         // Shared Kubernetes manager for controllers
	testEnv    *envtest.Environment // Test environment for controller-runtime
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	ctrl.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	// Initialize the test environment
	testEnv = &envtest.Environment{
		ControlPlaneStopTimeout: 90 * time.Second, // Extend the timeout for stopping the test environment
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "config", "crd", "bases"),
		},
		ErrorIfCRDPathMissing: true,

		// The BinaryAssetsDirectory is only required if you want to run the tests directly
		// without call the makefile target test. If not informed it will look for the
		// default path defined in controller-runtime which is /usr/local/kubebuilder/.
		// Note that you must have the required binaries setup under the bin directory to perform
		// the tests directly. When we run make test it will be setup and used automatically.
		BinaryAssetsDirectory: func() string {
			st, _ := filepath.Abs(filepath.Join("..", "..", "bin", "envtest", "k8s",
				fmt.Sprintf("1.31.0-%s-%s", runtime.GOOS, runtime.GOARCH)))
			return st
		}(),
	}

	// Start the environment and get the configuration
	cfg, err := testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	// Add API scheme
	err = ec2operatorv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	// Initialize Kubernetes manager
	k8sManager, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	Expect(err).NotTo(HaveOccurred())

	// Create a shared client
	k8sClient = k8sManager.GetClient()
	Expect(k8sClient).NotTo(BeNil())

	// Start the manager in a separate goroutine
	go func() {
		Expect(k8sManager.Start(ctrl.SetupSignalHandler())).To(Succeed())
	}()
})

var _ = AfterSuite(func() {
	Expect(testEnv.Stop()).To(Succeed())
})
