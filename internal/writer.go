package internal

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type K8SECRETCONSTRUCT struct {
	ApiVersion string
	Kind       string
	Metadata   METADATA
	Data       map[string]string
	Type       string
}

type METADATA struct {
	Name      string
	Namespace string
}

func WriteSecretYAMLFile(coreValues map[string]string, name string, namespace string, secretType string) {

	secret := K8SECRETCONSTRUCT{
		ApiVersion: "v1",
		Kind:       "Secret",
		Metadata: METADATA{
			Name:      name + "-secret",
			Namespace: namespace,
		},
		Type: secretType,
		Data: coreValues,
	}

	out, err := yaml.Marshal(secret)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(name+"-secret.yaml", out, 0644)

}
