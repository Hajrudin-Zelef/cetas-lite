---
id: collect-261001-fortinet/fortinet/fortinetdev-terraform-provider-fortios-blob-head-readme-md-d575ac25
title: "fortinetdev-terraform-provider-fortios-blob-head-readme-md-d575ac25"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/fortinetdev-terraform-provider-fortios-blob-head-readme-md-d575ac25.md
source_anchor: ""
source_lines: [1, 15]
sha256: 3ef44575b7acb7babe74d67ff71508b90091a04888bffee1b25bfadff619282f
---

# fortinetdev-terraform-provider-fortios-blob-head-readme-md-d575ac25

- Website: https://www.terraform.io
- Mailing list: Google Groups
- Terraform 0.12.x +
- Go Minimum version 1.24.1, or 1.23.7 if using the 1.23 branch (to build the provider plugin)
- The provider can cover FortiOS 6.0, 6.2, 6.4, 7.0, 7.2, 7.4, and 7.6 versions, the configuration of all parameters should be based on the relevant FortiOS version manual. For FortiManager, the support on this provider is deprecated, please use FortiManager Terraform provider instead.
- 
Clone repository to: $GOPATH/src/github.com/terraform-providers/terraform-provider-fortios .$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers $ git clone git@github.com:fortinetdev/terraform-provider-fortios
- 
Enter the provider directory and build the provider. $ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-fortios $ make build
If you're building the provider, follow the instructions to install it as a plugin. After placing it into your plugins directory,  run terraform init to initialize it.
$ terraform init
If you wish to work on the provider, you'll first need Go installed on your machine (version 1.13+ is required). You'll also need to correctly setup a GOPATH, as well as adding $GOPATH/bin to your $PATH.
To compile the provider, run make build. This will build the provider and put the provider binary in the $GOPATH/bin directory.
$ make build
...
