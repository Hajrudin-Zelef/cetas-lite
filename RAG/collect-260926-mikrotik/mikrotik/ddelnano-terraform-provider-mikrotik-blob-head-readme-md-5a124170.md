---
id: collect-260926-mikrotik/mikrotik/ddelnano-terraform-provider-mikrotik-blob-head-readme-md-5a124170
title: "custom.tfrc"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/ddelnano-terraform-provider-mikrotik-blob-head-readme-md-5a124170.md
source_anchor: ""
source_lines: [1, 49]
sha256: 4dc365c26b8ae7ba17d4b963723cd875d9c1ead1ba0ea8617cec2d2174d76dfa
---

# custom.tfrc

This is a terraform provider for managing resources on your RouterOS device. To see what resources and data sources are supported, please see the documentation on the terraform registry.
You can discuss any issues you have or feature requests in Discord.
If you get value out this project and want to show your support you can find me on patreon.
Requirements:
To build the provider with make:
$ make build
which creates a terraform-provider-mikrotik binary in repository's root folder.
or build with go compiler:
$ go build -o terraform-provider-mikrotik
To use locally built provider, Terraform should be aware of its binary.
It could be done with custom CLI config file:
# custom.tfrc
provider_installation {
    dev_overrides {
        "ddelnano/mikrotik" = "/path/to/clones/repository/terraform-provider-mikrotik"
    }
    direct {}
}
The dev_overrides section is available since Terraform 0.14.
Finally, tell Terraform CLI to use custom confiuration by exporting environment variable:
$ export TF_CLI_CONFIG_FILE=path/to/custom.tfrc
NOTE: with dev_overrides it is not possible to run terraform init (see official docs) so you should immediately use terraform plan and terraform apply without initializing.
- RouterOS. See which versions are supported by what is tested in CI
- Terraform 0.12+
For code generation of boilerplate code, see codegen Readme
The provider is tested with Terraform's acceptance testing framework. As long as you have a RouterOS device you should be able to run them. Please be aware it will create resources on your device! Code that is accepted by the project will not be destructive for anything existing on your router but be careful when changing test code!
In order to run the tests you will need to set the following environment variables:
export MIKROTIK_HOST=router-hostname:8728
export MIKROTIK_USER=username
# Please be aware this will put your password in your bash history and is not safe
export MIKROTIK_PASSWORD=password
After those environment variables are set you can run the tests with the following command:
make testacc
If you do not have MikroTik hardware or virtual machine with pre-installed RouterOS, you still have a way to run tests locally.
To make this happen, install Docker on your developer machine, and run from the root of the repository:
$ make routeros
It will start RouterOS container locally and make its API server available at 127.0.0.1:8728
Just export connection settings
export MIKROTIK_HOST=127.0.0.1:8728
export MIKROTIK_USER=admin
export MIKROTIK_PASSWORD=""
and you are ready to run tests with
$ make test
You can use specific RouterOS version by passing ROUTEROS_VERSION argument
$ make routeros ROUTEROS_VERSION="6.49beta54"
or even
$ make routeros ROUTEROS_VERSION=latest
To cleanup everything, just run:
$ make routeros-clean
