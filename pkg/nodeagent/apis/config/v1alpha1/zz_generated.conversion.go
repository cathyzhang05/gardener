//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	v3 "github.com/Masterminds/semver/v3"
	config "github.com/gardener/gardener/pkg/nodeagent/apis/config"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentbaseconfig "k8s.io/component-base/config"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*APIServer)(nil), (*config.APIServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_APIServer_To_config_APIServer(a.(*APIServer), b.(*config.APIServer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.APIServer)(nil), (*APIServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_APIServer_To_v1alpha1_APIServer(a.(*config.APIServer), b.(*APIServer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BootstrapConfiguration)(nil), (*config.BootstrapConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BootstrapConfiguration_To_config_BootstrapConfiguration(a.(*BootstrapConfiguration), b.(*config.BootstrapConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.BootstrapConfiguration)(nil), (*BootstrapConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_BootstrapConfiguration_To_v1alpha1_BootstrapConfiguration(a.(*config.BootstrapConfiguration), b.(*BootstrapConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerConfiguration)(nil), (*config.ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(a.(*ControllerConfiguration), b.(*config.ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControllerConfiguration)(nil), (*ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(a.(*config.ControllerConfiguration), b.(*ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeAgentConfiguration)(nil), (*config.NodeAgentConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NodeAgentConfiguration_To_config_NodeAgentConfiguration(a.(*NodeAgentConfiguration), b.(*config.NodeAgentConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeAgentConfiguration)(nil), (*NodeAgentConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeAgentConfiguration_To_v1alpha1_NodeAgentConfiguration(a.(*config.NodeAgentConfiguration), b.(*NodeAgentConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OperatingSystemConfigControllerConfig)(nil), (*config.OperatingSystemConfigControllerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_OperatingSystemConfigControllerConfig_To_config_OperatingSystemConfigControllerConfig(a.(*OperatingSystemConfigControllerConfig), b.(*config.OperatingSystemConfigControllerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.OperatingSystemConfigControllerConfig)(nil), (*OperatingSystemConfigControllerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_OperatingSystemConfigControllerConfig_To_v1alpha1_OperatingSystemConfigControllerConfig(a.(*config.OperatingSystemConfigControllerConfig), b.(*OperatingSystemConfigControllerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Server)(nil), (*config.Server)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Server_To_config_Server(a.(*Server), b.(*config.Server), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Server)(nil), (*Server)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Server_To_v1alpha1_Server(a.(*config.Server), b.(*Server), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServerConfiguration)(nil), (*config.ServerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(a.(*ServerConfiguration), b.(*config.ServerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ServerConfiguration)(nil), (*ServerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(a.(*config.ServerConfiguration), b.(*ServerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TokenControllerConfig)(nil), (*config.TokenControllerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_TokenControllerConfig_To_config_TokenControllerConfig(a.(*TokenControllerConfig), b.(*config.TokenControllerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TokenControllerConfig)(nil), (*TokenControllerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TokenControllerConfig_To_v1alpha1_TokenControllerConfig(a.(*config.TokenControllerConfig), b.(*TokenControllerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TokenSecretSyncConfig)(nil), (*config.TokenSecretSyncConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_TokenSecretSyncConfig_To_config_TokenSecretSyncConfig(a.(*TokenSecretSyncConfig), b.(*config.TokenSecretSyncConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TokenSecretSyncConfig)(nil), (*TokenSecretSyncConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TokenSecretSyncConfig_To_v1alpha1_TokenSecretSyncConfig(a.(*config.TokenSecretSyncConfig), b.(*TokenSecretSyncConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_APIServer_To_config_APIServer(in *APIServer, out *config.APIServer, s conversion.Scope) error {
	out.Server = in.Server
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	return nil
}

// Convert_v1alpha1_APIServer_To_config_APIServer is an autogenerated conversion function.
func Convert_v1alpha1_APIServer_To_config_APIServer(in *APIServer, out *config.APIServer, s conversion.Scope) error {
	return autoConvert_v1alpha1_APIServer_To_config_APIServer(in, out, s)
}

func autoConvert_config_APIServer_To_v1alpha1_APIServer(in *config.APIServer, out *APIServer, s conversion.Scope) error {
	out.Server = in.Server
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	return nil
}

// Convert_config_APIServer_To_v1alpha1_APIServer is an autogenerated conversion function.
func Convert_config_APIServer_To_v1alpha1_APIServer(in *config.APIServer, out *APIServer, s conversion.Scope) error {
	return autoConvert_config_APIServer_To_v1alpha1_APIServer(in, out, s)
}

func autoConvert_v1alpha1_BootstrapConfiguration_To_config_BootstrapConfiguration(in *BootstrapConfiguration, out *config.BootstrapConfiguration, s conversion.Scope) error {
	out.KubeletDataVolumeSize = (*int64)(unsafe.Pointer(in.KubeletDataVolumeSize))
	return nil
}

// Convert_v1alpha1_BootstrapConfiguration_To_config_BootstrapConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_BootstrapConfiguration_To_config_BootstrapConfiguration(in *BootstrapConfiguration, out *config.BootstrapConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_BootstrapConfiguration_To_config_BootstrapConfiguration(in, out, s)
}

func autoConvert_config_BootstrapConfiguration_To_v1alpha1_BootstrapConfiguration(in *config.BootstrapConfiguration, out *BootstrapConfiguration, s conversion.Scope) error {
	out.KubeletDataVolumeSize = (*int64)(unsafe.Pointer(in.KubeletDataVolumeSize))
	return nil
}

// Convert_config_BootstrapConfiguration_To_v1alpha1_BootstrapConfiguration is an autogenerated conversion function.
func Convert_config_BootstrapConfiguration_To_v1alpha1_BootstrapConfiguration(in *config.BootstrapConfiguration, out *BootstrapConfiguration, s conversion.Scope) error {
	return autoConvert_config_BootstrapConfiguration_To_v1alpha1_BootstrapConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_OperatingSystemConfigControllerConfig_To_config_OperatingSystemConfigControllerConfig(&in.OperatingSystemConfig, &out.OperatingSystemConfig, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_TokenControllerConfig_To_config_TokenControllerConfig(&in.Token, &out.Token, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in, out, s)
}

func autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	if err := Convert_config_OperatingSystemConfigControllerConfig_To_v1alpha1_OperatingSystemConfigControllerConfig(&in.OperatingSystemConfig, &out.OperatingSystemConfig, s); err != nil {
		return err
	}
	if err := Convert_config_TokenControllerConfig_To_v1alpha1_TokenControllerConfig(&in.Token, &out.Token, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration is an autogenerated conversion function.
func Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_NodeAgentConfiguration_To_config_NodeAgentConfiguration(in *NodeAgentConfiguration, out *config.NodeAgentConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	out.LogFormat = in.LogFormat
	if err := Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	if in.Debugging != nil {
		in, out := &in.Debugging, &out.Debugging
		*out = new(componentbaseconfig.DebuggingConfiguration)
		if err := configv1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Debugging = nil
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := Convert_v1alpha1_APIServer_To_config_APIServer(&in.APIServer, &out.APIServer, s); err != nil {
		return err
	}
	out.Bootstrap = (*config.BootstrapConfiguration)(unsafe.Pointer(in.Bootstrap))
	if err := Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(&in.Controllers, &out.Controllers, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_NodeAgentConfiguration_To_config_NodeAgentConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_NodeAgentConfiguration_To_config_NodeAgentConfiguration(in *NodeAgentConfiguration, out *config.NodeAgentConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeAgentConfiguration_To_config_NodeAgentConfiguration(in, out, s)
}

func autoConvert_config_NodeAgentConfiguration_To_v1alpha1_NodeAgentConfiguration(in *config.NodeAgentConfiguration, out *NodeAgentConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	out.LogFormat = in.LogFormat
	if err := Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	if in.Debugging != nil {
		in, out := &in.Debugging, &out.Debugging
		*out = new(configv1alpha1.DebuggingConfiguration)
		if err := configv1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Debugging = nil
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := Convert_config_APIServer_To_v1alpha1_APIServer(&in.APIServer, &out.APIServer, s); err != nil {
		return err
	}
	out.Bootstrap = (*BootstrapConfiguration)(unsafe.Pointer(in.Bootstrap))
	if err := Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(&in.Controllers, &out.Controllers, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_NodeAgentConfiguration_To_v1alpha1_NodeAgentConfiguration is an autogenerated conversion function.
func Convert_config_NodeAgentConfiguration_To_v1alpha1_NodeAgentConfiguration(in *config.NodeAgentConfiguration, out *NodeAgentConfiguration, s conversion.Scope) error {
	return autoConvert_config_NodeAgentConfiguration_To_v1alpha1_NodeAgentConfiguration(in, out, s)
}

func autoConvert_v1alpha1_OperatingSystemConfigControllerConfig_To_config_OperatingSystemConfigControllerConfig(in *OperatingSystemConfigControllerConfig, out *config.OperatingSystemConfigControllerConfig, s conversion.Scope) error {
	out.SyncPeriod = (*v1.Duration)(unsafe.Pointer(in.SyncPeriod))
	out.SyncJitterPeriod = (*v1.Duration)(unsafe.Pointer(in.SyncJitterPeriod))
	out.SecretName = in.SecretName
	out.KubernetesVersion = (*v3.Version)(unsafe.Pointer(in.KubernetesVersion))
	return nil
}

// Convert_v1alpha1_OperatingSystemConfigControllerConfig_To_config_OperatingSystemConfigControllerConfig is an autogenerated conversion function.
func Convert_v1alpha1_OperatingSystemConfigControllerConfig_To_config_OperatingSystemConfigControllerConfig(in *OperatingSystemConfigControllerConfig, out *config.OperatingSystemConfigControllerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_OperatingSystemConfigControllerConfig_To_config_OperatingSystemConfigControllerConfig(in, out, s)
}

func autoConvert_config_OperatingSystemConfigControllerConfig_To_v1alpha1_OperatingSystemConfigControllerConfig(in *config.OperatingSystemConfigControllerConfig, out *OperatingSystemConfigControllerConfig, s conversion.Scope) error {
	out.SyncPeriod = (*v1.Duration)(unsafe.Pointer(in.SyncPeriod))
	out.SyncJitterPeriod = (*v1.Duration)(unsafe.Pointer(in.SyncJitterPeriod))
	out.SecretName = in.SecretName
	out.KubernetesVersion = (*v3.Version)(unsafe.Pointer(in.KubernetesVersion))
	return nil
}

// Convert_config_OperatingSystemConfigControllerConfig_To_v1alpha1_OperatingSystemConfigControllerConfig is an autogenerated conversion function.
func Convert_config_OperatingSystemConfigControllerConfig_To_v1alpha1_OperatingSystemConfigControllerConfig(in *config.OperatingSystemConfigControllerConfig, out *OperatingSystemConfigControllerConfig, s conversion.Scope) error {
	return autoConvert_config_OperatingSystemConfigControllerConfig_To_v1alpha1_OperatingSystemConfigControllerConfig(in, out, s)
}

func autoConvert_v1alpha1_Server_To_config_Server(in *Server, out *config.Server, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_v1alpha1_Server_To_config_Server is an autogenerated conversion function.
func Convert_v1alpha1_Server_To_config_Server(in *Server, out *config.Server, s conversion.Scope) error {
	return autoConvert_v1alpha1_Server_To_config_Server(in, out, s)
}

func autoConvert_config_Server_To_v1alpha1_Server(in *config.Server, out *Server, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_config_Server_To_v1alpha1_Server is an autogenerated conversion function.
func Convert_config_Server_To_v1alpha1_Server(in *config.Server, out *Server, s conversion.Scope) error {
	return autoConvert_config_Server_To_v1alpha1_Server(in, out, s)
}

func autoConvert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in *ServerConfiguration, out *config.ServerConfiguration, s conversion.Scope) error {
	out.HealthProbes = (*config.Server)(unsafe.Pointer(in.HealthProbes))
	out.Metrics = (*config.Server)(unsafe.Pointer(in.Metrics))
	return nil
}

// Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in *ServerConfiguration, out *config.ServerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in, out, s)
}

func autoConvert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *config.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	out.HealthProbes = (*Server)(unsafe.Pointer(in.HealthProbes))
	out.Metrics = (*Server)(unsafe.Pointer(in.Metrics))
	return nil
}

// Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration is an autogenerated conversion function.
func Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *config.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_TokenControllerConfig_To_config_TokenControllerConfig(in *TokenControllerConfig, out *config.TokenControllerConfig, s conversion.Scope) error {
	out.SyncConfigs = *(*[]config.TokenSecretSyncConfig)(unsafe.Pointer(&in.SyncConfigs))
	out.SyncPeriod = (*v1.Duration)(unsafe.Pointer(in.SyncPeriod))
	return nil
}

// Convert_v1alpha1_TokenControllerConfig_To_config_TokenControllerConfig is an autogenerated conversion function.
func Convert_v1alpha1_TokenControllerConfig_To_config_TokenControllerConfig(in *TokenControllerConfig, out *config.TokenControllerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_TokenControllerConfig_To_config_TokenControllerConfig(in, out, s)
}

func autoConvert_config_TokenControllerConfig_To_v1alpha1_TokenControllerConfig(in *config.TokenControllerConfig, out *TokenControllerConfig, s conversion.Scope) error {
	out.SyncConfigs = *(*[]TokenSecretSyncConfig)(unsafe.Pointer(&in.SyncConfigs))
	out.SyncPeriod = (*v1.Duration)(unsafe.Pointer(in.SyncPeriod))
	return nil
}

// Convert_config_TokenControllerConfig_To_v1alpha1_TokenControllerConfig is an autogenerated conversion function.
func Convert_config_TokenControllerConfig_To_v1alpha1_TokenControllerConfig(in *config.TokenControllerConfig, out *TokenControllerConfig, s conversion.Scope) error {
	return autoConvert_config_TokenControllerConfig_To_v1alpha1_TokenControllerConfig(in, out, s)
}

func autoConvert_v1alpha1_TokenSecretSyncConfig_To_config_TokenSecretSyncConfig(in *TokenSecretSyncConfig, out *config.TokenSecretSyncConfig, s conversion.Scope) error {
	out.SecretName = in.SecretName
	out.Path = in.Path
	return nil
}

// Convert_v1alpha1_TokenSecretSyncConfig_To_config_TokenSecretSyncConfig is an autogenerated conversion function.
func Convert_v1alpha1_TokenSecretSyncConfig_To_config_TokenSecretSyncConfig(in *TokenSecretSyncConfig, out *config.TokenSecretSyncConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_TokenSecretSyncConfig_To_config_TokenSecretSyncConfig(in, out, s)
}

func autoConvert_config_TokenSecretSyncConfig_To_v1alpha1_TokenSecretSyncConfig(in *config.TokenSecretSyncConfig, out *TokenSecretSyncConfig, s conversion.Scope) error {
	out.SecretName = in.SecretName
	out.Path = in.Path
	return nil
}

// Convert_config_TokenSecretSyncConfig_To_v1alpha1_TokenSecretSyncConfig is an autogenerated conversion function.
func Convert_config_TokenSecretSyncConfig_To_v1alpha1_TokenSecretSyncConfig(in *config.TokenSecretSyncConfig, out *TokenSecretSyncConfig, s conversion.Scope) error {
	return autoConvert_config_TokenSecretSyncConfig_To_v1alpha1_TokenSecretSyncConfig(in, out, s)
}
