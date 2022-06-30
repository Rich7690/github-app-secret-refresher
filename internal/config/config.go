package config

import (
	"log"
	"os"
	"strconv"

	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/types"
	"github.com/joho/godotenv"
)

var (
	TokenProcessorType      types.TokenProcessorType
	GithubAppId             int
	GithubAppInstallationId int
	GithubAppPrivateKeyFile string
	GithubAppPrivateKey     string // this allows setting env vars with tools like Hashicorp Vault

	KubeConfigPath         string
	KubeSecretName         string
	KubeSecretNamespace    string
	KubeCombinedSecretName string
)

func Load() {
	godotenv.Load()

	TokenProcessorType = types.TokenProcessorType(os.Getenv("TOKEN_PROCESSOR_TYPE"))
	GithubAppId = getEnvAsInt("GITHUB_APP_ID")
	GithubAppInstallationId = getEnvAsInt("GITHUB_APP_INSTALLATION_ID")
	GithubAppPrivateKeyFile = os.Getenv("GITHUB_APP_PRIVATE_KEY_PATH")
	GithubAppPrivateKey = os.Getenv("GITHUB_APP_PRIVATE_KEY")

	KubeConfigPath = os.Getenv("KUBE_CONFIG_PATH")
	KubeSecretName = os.Getenv("KUBE_SECRET_NAME")
	KubeSecretNamespace = os.Getenv("KUBE_SECRET_NAMESPACE")
	KubeCombinedSecretName = os.Getenv("KUBE_COMBINED_SECRET_NAME")
	if KubeCombinedSecretName == "" {
		KubeCombinedSecretName = "combined" // default if not set
	}
}

func getEnvAsInt(envVar string) int {
	if val := os.Getenv(envVar); val != "" {
		intVal, err := strconv.Atoi(val)

		if err == nil {
			return intVal
		}

		log.Panicf("Environment variable %s is not an int", envVar)
	}

	log.Panicf("Environment variable %s is not an int", envVar)
	return 0
}
