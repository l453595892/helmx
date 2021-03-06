package kubetypes

type SecretVolumeSource struct {
	SecretName string `toml:"secretName,omitempty" json:"secretName,omitempty" yaml:"secretName,omitempty"`
}

type ConfigMapVolumeSource struct {
	KubeLocalObjectReference `yaml:",inline"`
}

type PersistentVolumeClaimVolumeSource struct {
	ClaimName string `toml:"claimName" json:"claimName" yaml:"claimName"`
	ReadOnly  bool   `toml:"readOnly" json:"readOnly"  yaml:"readOnly"`
}

type HostPathVolumeSource struct {
	Path string `toml:"path" json:"path" yaml:"path"`
	Type string `toml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}
