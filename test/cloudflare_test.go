package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"io/ioutil"
	"log"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCloudflare(t *testing.T) {
	t.Parallel()

	exampleFolder := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/cloudflare")

	defer test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, exampleFolder)
		terraform.Destroy(t, terraformOptions)

		keyPair := test_structure.LoadEc2KeyPair(t, exampleFolder)
		aws.DeleteEC2KeyPair(t, keyPair)
	})

	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions, keyPair := configureCloudflareOptions(t, exampleFolder)
		test_structure.SaveTerraformOptions(t, exampleFolder, terraformOptions)
		test_structure.SaveEc2KeyPair(t, exampleFolder, keyPair)

		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, exampleFolder)

		testCloudflareEndpoints(t, terraformOptions)
	})
}


func configureCloudflareOptions(t *testing.T, exampleFolder string) (*terraform.Options, *aws.Ec2Keypair) {

	uniqueID := random.UniqueId()
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	keyPairName := fmt.Sprintf("ci-node-%s", uniqueID)
	keyPair := aws.CreateAndImportEC2KeyPair(t, awsRegion, keyPairName)

	privateKeyPath := path.Join(exampleFolder, keyPairName)

	err := ioutil.WriteFile(privateKeyPath, []byte(keyPair.PrivateKey), 0600)
	if err != nil {
		panic(err)
	}

	log.Printf("Key saved to: %s", privateKeyPath)

	terraformOptions := &terraform.Options{
		TerraformDir: exampleFolder,
		Vars: map[string]interface{}{
			"aws_region":    awsRegion,
			"public_key":    keyPair.PublicKey,
			"private_key_path": privateKeyPath,
			"create_ansible": true,
			"root_domain_name": "insight-infra.net",
		},
	}

	return terraformOptions, keyPair
}

func testCloudflareEndpoints(t *testing.T, terraformOptions *terraform.Options) {

	fqdn := terraform.Output(t, terraformOptions, "fqdn")

	expectedStatus := "200"
	body := strings.NewReader(`{"id":1, "jsonrpc":"2.0", "method":"system_health", "params":[]}`)
	url := fmt.Sprintf("http://%s:9933", fqdn)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	description := fmt.Sprintf("curl to LB %s with error command", fqdn)
	maxRetries := 60
	timeBetweenRetries := 2 * time.Second

	retry.DoWithRetry(t, description, maxRetries, timeBetweenRetries, func() (string, error) {

		outputStatus, _, err := http_helper.HTTPDoE(t, "POST", url, body, headers, nil)

		if err != nil {
			return "", err
		}

		if strings.TrimSpace(strconv.Itoa(outputStatus)) != expectedStatus {
			return "", fmt.Errorf("expected SSH command to return '%s' but got '%s'", expectedStatus, strconv.Itoa(outputStatus))
		}

		return "", nil
	})
}
