---
id: collect-261001-meraki/meraki/ciscodevnet-terraform-provider-meraki-acf740ab
title: "ciscodevnet-terraform-provider-meraki-acf740ab"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ciscodevnet-terraform-provider-meraki-acf740ab.md
source_anchor: ""
source_lines: [1, 30]
sha256: 8d2ce126a8073f1393939a8ab29cacfb46c14ff1bbcd40f30c4f481d6179cee4
---

# ciscodevnet-terraform-provider-meraki-acf740ab

The Meraki provider provides resources to interact with a Cisco Meraki Dashboard. It communicates with Meraki Dashboard via the REST API.
All resources and data sources have been tested with the following API releases.
| API | Version | 
|---|---|
| Meraki Dashboard | 1.67.0 | 
Documentation: https://registry.terraform.io/providers/CiscoDevNet/meraki/latest
As this is a community-driven project, support is provided by the community. If you encounter issues or have questions, please use the following resources:
- GitHub Issues: Report bugs or request features on our GitHub Issues page.
- Discussion Forums: Engage with other users and contributors on GitHub Discussions.
We welcome contributions from the community! If you'd like to contribute, please follow our contribution guidelines. Whether it's reporting bugs, suggesting features, or submitting pull requests, your involvement helps improve the project for everyone.
- Clone the repository
- Enter the repository directory
- Build the provider using the Go install command:
go install
This provider uses Go modules. Please see the Go documentation for the most up to date information about using Go modules.
To add a new dependency github.com/author/dependency to your Terraform provider:
go get github.com/author/dependency
go mod tidy
Then commit the changes to go.mod and go.sum.
This Terraform Provider is available to install automatically via terraform init. If you're building the provider, follow the instructions to
install it as a plugin.
After placing it into your plugins directory,  run terraform init to initialize it.
Additional documentation, including available resources and their arguments/attributes can be found on the Terraform documentation website.
If you wish to work on the provider, you'll first need Go installed on your machine (see Requirements above).
To compile the provider, run go install. This will build the provider and put the provider binary in the $GOPATH/bin directory.
More information about how the code generation works can be found in the contribution guide.
In order to run the full suite of acceptance tests, run make test. Make sure the respective environment variables are set (e.g., MERAKI_API_KEY).
Note: Acceptance tests create real resources.
make test
More information about how the acceptance tests work can be found in the contribution guide.
