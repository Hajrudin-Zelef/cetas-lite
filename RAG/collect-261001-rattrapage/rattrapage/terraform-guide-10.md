---
id: collect-261001-rattrapage/rattrapage/terraform-guide-10
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident", "memory"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [2200, 2449]
sha256: 0c099706e24eeab4f44b2ce814f30787bba4520c2ab9f0d380950c78f8205c66
---

# Terraform — Guide complet : Infrastructure as Code en production

```hcl
# variables.tf
variable "proxmox_endpoint"  { type = string }
variable "proxmox_api_token" { type = string, sensitive = true }
variable "vm_count"    { type = number, default = 10 }
variable "base_vmid"   { type = number, default = 200 }
variable "subnet"      { type = string, default = "10.20.0.0/24" }
variable "gateway"     { type = string, default = "10.20.0.1" }
variable "nodes"       { type = list(string), default = ["pve1", "pve2"] }
variable "template_id" { type = number, default = 9000 }  # template debian-12-cloudinit
variable "ssh_key"     { type = string, sensitive = true }
```

```hcl
# main.tf
provider "proxmox" {
  endpoint  = var.proxmox_endpoint
  api_token = var.proxmox_api_token
  insecure  = true   # ⚠️ labo avec CA autosignée ; false en prod avec CA valide
}

locals {
  vms = {
    for i in range(var.vm_count) : format("web-%02d", i + 1) => {
      node = var.nodes[i % length(var.nodes)]       # alternance pve1/pve2
      vmid = var.base_vmid + i
      ip   = cidrhost(var.subnet, 11 + i)
    }
  }
}

resource "proxmox_virtual_environment_file" "cloud_config" {
  for_each     = local.vms
  content_type = "snippets"
  datastore_id = "local"
  node_name    = each.value.node
  source_raw {
    file_name = "${each.key}-cloud-config.yaml"
    data = templatefile("${path.module}/templates/cloud-init.tftpl", {
      hostname   = each.key
      admin_user = "admin"
      ssh_key    = var.ssh_key
    })
  }
}

resource "proxmox_virtual_environment_vm" "web" {
  for_each  = local.vms
  name      = each.key
  node_name = each.value.node
  vm_id     = each.value.vmid
  tags      = ["terraform", "prod", "web", "farm"]

  clone {
    vm_id = var.template_id
    full  = true
  }
  cpu { cores = 4 }
  memory { dedicated = 8192 }
  disk {
    datastore_id = "ceph-vm"
    interface    = "scsi0"
    size         = 40
  }
  network_device { bridge = "vmbr0", vlan_id = 20 }

  initialization {
    datastore_id      = "local-lvm"
    user_data_file_id = proxmox_virtual_environment_file.cloud_config[each.key].id
    ip_config {
      ipv4 { address = "${each.value.ip}/24", gateway = var.gateway }
    }
  }

  lifecycle {
    replace_triggered_by = [proxmox_virtual_environment_file.cloud_config[each.key].id]
  }
}
```

```hcl
# outputs.tf
output "vm_inventory" {
  description = "Nom, nœud, VMID et IP de chaque VM"
  value = {
    for k, vm in proxmox_virtual_environment_vm.web : k => {
      node = vm.node_name
      vmid = vm.vm_id
      ip   = local.vms[k].ip
    }
  }
}
```

```bash
terraform init
terraform plan -out=tfplan        # vérifier : 10 VM + 10 snippets = 20 ajouts
terraform apply tfplan
terraform output -json vm_inventory | jq .
```

---

## 66. Cas pratique 2 : reprise après sinistre

Scénario : le nœud `pve1` est perdu (incendie). Les VM critiques doivent renaître sur `pve2`/`pve3` depuis le code + les sauvegardes.

**Pré-requis (à préparer AVANT le sinistre)** :
- [ ] Code Terraform versionné (Git) décrivant tout le parc
- [ ] State distant (S3 versionné) — **hors du site sinistré** si possible
- [ ] Sauvegardes VM (Proxmox Backup Server) répliquées hors site
- [ ] `terraform.tfvars` de prod sauvegardé chiffré (SOPS)

**Procédure** :

```bash
# 1. Sur le poste de secours, cloner le code
git clone https://git.entreprise.lan/infra/proxmox-prod.git && cd proxmox-prod/envs/prod

# 2. Restaurer les variables chiffrées
sops --decrypt secrets.tfvars.enc > /tmp/secrets.tfvars

# 3. Pointer vers le cluster de secours (nouveau endpoint)
export TF_VAR_proxmox_endpoint="https://pve2-secours.lan:8006/"
export TF_VAR_proxmox_api_token="..."

# 4. Init + plan : Terraform détecte les VM manquantes et propose de les recréer
terraform init
terraform plan -out=tfplan   # RELIRE : que des créations attendues

# 5. Restaurer les données depuis PBS AVANT ou APRÈS selon la stratégie :
#    Option A (simple) : Terraform recrée des VM vierges, puis restauration PBS par-dessus
#    Option B : adapter le code pour cloner depuis les sauvegardes si le provider le permet
terraform apply tfplan
```

**Leçons** :
- Le code seul ne restaure pas les **données** : la sauvegarde (PBS) reste indispensable.
- Tester la procédure 1×/an (exercice PRA) : un plan jamais testé est un vœu pieux.
- `prevent_destroy` sur les BDD : en cas de vrai sinistre, c'est une protection, pas un obstacle — on le retire consciemment si besoin.

---

## 67. Cas pratique 3 : réseau complet avec VLAN et cloud-init

```hcl
# VMs sur 3 VLANs : 10 (DMZ web), 20 (applicatif), 30 (BDD)
locals {
  vlans = { web = 10, app = 20, db = 30 }
  tiers = {
    "web-01" = { vlan = "web", ip = "10.10.10.11/24", node = "pve1" }
    "web-02" = { vlan = "web", ip = "10.10.10.12/24", node = "pve2" }
    "app-01" = { vlan = "app", ip = "10.10.20.11/24", node = "pve1" }
    "db-01"  = { vlan = "db",  ip = "10.10.30.11/24", node = "pve2" }
  }
}

resource "proxmox_virtual_environment_vm" "tiers" {
  for_each  = local.tiers
  name      = each.key
  node_name = each.value.node
  tags      = ["terraform", "prod", each.value.vlan]

  clone { vm_id = var.template_id, full = true }
  cpu { cores = each.value.vlan == "db" ? 8 : 4 }
  memory { dedicated = each.value.vlan == "db" ? 16384 : 8192 }
  disk {
    datastore_id = "ceph-vm"
    interface    = "scsi0"
    size         = each.value.vlan == "db" ? 100 : 40
  }
  network_device {
    bridge  = "vmbr0"
    vlan_id = local.vlans[each.value.vlan]
  }
  initialization {
    datastore_id = "local-lvm"
    ip_config {
      ipv4 {
        address = each.value.ip
        gateway = "10.10.${local.vlans[each.value.vlan]}.1"
      }
    }
  }
}
```

> 💡 La passerelle est dérivée du VLAN (`10.10.<vlan>.1`) : un seul endroit à maintenir. Les bridges VLAN-aware (`bridge-vlan-aware yes`) doivent être configurés côté Proxmox au préalable.

---

## 68. Cas pratique 4 : cluster à 3 nœuds avec anti-affinité

Pour la HA, les 3 VM d'un service ne doivent pas être sur le même nœud Proxmox :

```hcl
locals {
  ha_nodes = ["pve1", "pve2", "pve3"]
  cluster_vms = {
    for i in range(3) : format("kafka-%02d", i + 1) => {
      node = local.ha_nodes[i % length(local.ha_nodes)]
      vmid = 300 + i
      ip   = cidrhost("10.30.0.0/24", 11 + i)
    }
  }
}

resource "proxmox_virtual_environment_vm" "kafka" {
  for_each  = local.cluster_vms
  name      = each.key
  node_name = each.value.node   # pve1, pve2, pve3 : anti-affinité naturelle
  vm_id     = each.value.vmid
  tags      = ["terraform", "prod", "kafka", "ha"]

  clone { vm_id = var.template_id, full = true }
  cpu { cores = 4 }
  memory { dedicated = 8192 }
  disk { datastore_id = "ceph-vm", interface = "scsi0", size = 60 }
  network_device { bridge = "vmbr0", vlan_id = 30 }

  initialization {
    datastore_id = "local-lvm"
    ip_config { ipv4 { address = "${each.value.ip}/24", gateway = "10.30.0.1" } }
  }

  # Garde-fou : jamais de destruction accidentelle du cluster
  lifecycle { prevent_destroy = true }
}

output "kafka_placement" {
  value = { for k, vm in proxmox_virtual_environment_vm.kafka : k => vm.node_name }
}
```

> 💡 Complétez avec les groupes HA de Proxmox (`ha-manager`) côté cluster pour le redémarrage automatique en cas de panne de nœud — Terraform pose les VM, Proxmox HA les protège à l'exécution.

---

## 69. Supervision et alerting autour de Terraform

Terraform ne supervise pas l'infra à l'exécution : il faut l'entourer.

| Couche | Outil | Rôle |
|---|---|---|
| Drift | `plan -refresh-only` nightly (section 56) | détecter les changements hors code |
| Métriques Proxmox | Prometheus + `pve-exporter` | CPU/RAM/disque des nœuds et VM |
| Logs | Loki / journald centralisé | corréler un apply avec un incident |
| Alerting | Alertmanager / mail | dérive, échec de pipeline, verrou bloqué |
| Audit | `TF_LOG` + logs CI archivés | qui a appliqué quoi, quand |

