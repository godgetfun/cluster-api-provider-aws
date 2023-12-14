/*
Copyright 2019 The Kubernetes Authors.

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

// Package mock_services provides a way to generate mock services for the cloud provider.
// Run go generate to regenerate this mock. //nolint:revive
//
//go:generate ../../../../hack/tools/bin/mockgen -destination ec2_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services EC2Interface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt ec2_interface_mock.go > _ec2_interface_mock.go && mv _ec2_interface_mock.go ec2_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination reconcile_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services MachinePoolReconcileInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt reconcile_interface_mock.go > _reconcile_interface_mock.go && mv _reconcile_interface_mock.go reconcile_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination secretsmanager_machine_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services SecretInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt secretsmanager_machine_interface_mock.go > _secretsmanager_machine_interface_mock.go && mv _secretsmanager_machine_interface_mock.go secretsmanager_machine_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination objectstore_machine_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services ObjectStoreInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt objectstore_machine_interface_mock.go > _objectstore_machine_interface_mock.go && mv _objectstore_machine_interface_mock.go objectstore_machine_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination autoscaling_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services ASGInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt autoscaling_interface_mock.go > _autoscaling_interface_mock.go && mv _autoscaling_interface_mock.go autoscaling_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination elb_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services ELBInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt elb_interface_mock.go > _elb_interface_mock.go && mv _elb_interface_mock.go elb_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination network_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services NetworkInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt network_interface_mock.go > _network_interface_mock.go && mv _network_interface_mock.go network_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination security_group_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services SecurityGroupInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt security_group_interface_mock.go > _security_group_interface_mock.go && mv _security_group_interface_mock.go security_group_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination aws_node_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services AWSNodeInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt aws_node_interface_mock.go > _aws_node_interface_mock.go && mv _aws_node_interface_mock.go aws_node_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination iam_authenticator_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services IAMAuthenticatorInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt iam_authenticator_interface_mock.go > _iam_authenticator_interface_mock.go && mv _iam_authenticator_interface_mock.go iam_authenticator_interface_mock.go"
//go:generate ../../../../hack/tools/bin/mockgen -destination kube_proxy_interface_mock.go -package mock_services sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services KubeProxyInterface
//go:generate /usr/bin/env bash -c "cat ../../../../hack/boilerplate/boilerplate.generatego.txt kube_proxy_interface_mock.go > _kube_proxy_interface_mock.go && mv _kube_proxy_interface_mock.go kube_proxy_interface_mock.go"
package mock_services //nolint:stylecheck
