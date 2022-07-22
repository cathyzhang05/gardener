// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backupbucket

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/gardener/gardener/pkg/logger"
	"github.com/gardener/gardener/test/framework"

	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func TestBackupBucket(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Extensions Controller BackupBucket Integration Test Suite")
}

var (
	ctx        = context.Background()
	err        error
	log        logr.Logger
	testEnv    *envtest.Environment
	restConfig *rest.Config
)

var _ = BeforeSuite(func() {
	logf.SetLogger(logger.MustNewZapLogger(logger.DebugLevel, logger.FormatJSON, zap.WriteTo(GinkgoWriter)))
	log = logf.Log.WithName("backupbucket-test")

	By("starting test environment")
	extensionsCRDs := filepath.Join("..", "..", "..", "..", "..", "pkg", "operation", "botanist", "component", "extensions", "crds", "templates")
	testEnv = &envtest.Environment{
		CRDInstallOptions: envtest.CRDInstallOptions{
			Paths: []string{
				filepath.Join(extensionsCRDs, "crd-extensions.gardener.cloud_backupbuckets.yaml"),
				filepath.Join(extensionsCRDs, "crd-extensions.gardener.cloud_clusters.yaml"),
			},
		},
		ErrorIfCRDPathMissing: true,
	}
	restConfig, err = testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(restConfig).ToNot(BeNil())

})

var _ = AfterSuite(func() {
	By("running cleanup actions")
	framework.RunCleanupActions()

	By("stopping test environment")
	Expect(testEnv.Stop()).To(Succeed())
})
