---
id: etape6-phasee2-ansible-nornir-terraform/04-versions-pinned-binary-pinned-in-mise-toml-opentofu-1-12-6/overview
title: "versions pinned; binary pinned in mise.toml (opentofu = \"1.12.6\")"
domain: versions-pinned-binary-pinned-in-mise-toml-opentofu-1-12-6
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [360, 378]
section: "versions pinned; binary pinned in mise.toml (opentofu = \"1.12.6\")"
sha256: 30b6be7c068bf4fb67726a05cf2e53acf041162e25f0f001dd600967cf741511
---

# versions pinned; binary pinned in mise.toml (opentofu = "1.12.6")
terraform {
  required_providers {
    netbox = { source = "e-breuninger/netbox", version = "~> 4.1.0" }
  }
  backend "s3" {} # native locking, no DynamoDB needed (OpenTofu 1.10+)
}

variable "netbox_token" {
  type      = string
  sensitive = true
  ephemeral = true # never persisted in state (OpenTofu 1.11+)
}

provider "netbox" {
  server_url = "https://netbox.example.com"
  api_token  = var.netbox_token
}

