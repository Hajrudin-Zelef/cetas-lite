---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-guides-csv-users-md-3b740f34
title: "paultyng-terraform-provider-unifi-blob-head-docs-guides-csv-users-md-3b740f34"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-guides-csv-users-md-3b740f34.md
source_anchor: ""
source_lines: [1, 21]
sha256: 35e6d87ed80f87bf8a33b4ca83b86ca469302a68b79f1490afd500aec7c221b2
---

# paultyng-terraform-provider-unifi-blob-head-docs-guides-csv-users-md-3b740f34

| subcategory |  | 
|---|---|
| page_title | Manage Users/Clients in a CSV - Unifi Provider | 
| description | An example of using a CSV to manage all of your users of your network. | 
Given a CSV file with the following content:
mac,name,note
01:23:45:67:89:AB,My Device,custom note
You could create/manage a unifi_user for every row/MAC address in the CSV with the following config:
locals {
  userscsv = csvdecode(file("${path.module}/users.csv"))
  users    = { for user in local.userscsv : user.mac => user }
}
resource "unifi_user" "user" {
  for_each = local.users
  mac  = each.key
  name = each.value.name
  # append an optional additional note
  note = trimspace("${each.value.note}\n\nmanaged by TF")
  allow_existing         = true
  skip_forget_on_destroy = true
}
