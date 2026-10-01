---
id: collect-261001-rattrapage/rattrapage/ansible-guide-2
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [195, 463]
sha256: c8fa3e25df0be6bf59e1755022a1e5ee603472208d553e4804f73a3be4942533
---

# Guide Ansible — Automatisation système en production

- [ ] `host_key_checking = True` (sauf bootstrap initial documenté)
- [ ] `forks` adapté au parc (20–50 pour 200 nœuds, voir section 50)
- [ ] `log_path` défini et le dossier `logs/` exclu de Git via `.gitignore`
- [ ] `private_key_file` pointe vers une clé dédiée à Ansible, pas votre clé perso
- [ ] `stdout_callback = yaml` pour des logs exploitables

## 8. Arborescence projet recommandée

```text
ansible-projet/
├── ansible.cfg
├── inventaires/
│   ├── production.ini
│   ├── staging.ini
│   └── dynamique/
│       ├── proxmox.yml          # plugin d'inventaire Proxmox
│       └── aws_ec2.yml          # plugin d'inventaire AWS
├── group_vars/
│   ├── all.yml                  # variables globales
│   ├── webservers.yml
│   └── dbservers.yml
├── host_vars/
│   └── db-01.yml
├── roles/
│   └── mon_role/
│       ├── tasks/main.yml
│       ├── handlers/main.yml
│       ├── templates/
│       ├── files/
│       ├── vars/main.yml
│       ├── defaults/main.yml
│       └── meta/main.yml
├── playbooks/
│   ├── site.yml
│   ├── webservers.yml
│   └── durcissement_ssh.yml
├── collections/
│   └── requirements.yml
├── logs/                        # .gitignore
├── .gitignore
└── README.md
```

## 9. Inventaire statique au format INI

```ini
# inventaires/production.ini
[webservers]
web-01 ansible_host=192.168.10.11
web-02 ansible_host=192.168.10.12

[dbservers]
db-01 ansible_host=192.168.10.21 ansible_user=dba

[proxmox]
pve-01 ansible_host=192.168.10.2
pve-02 ansible_host=192.168.10.3

# Groupe de groupes
[production:children]
webservers
dbservers
proxmox

# Variables de groupe
[production:vars]
ansible_ssh_private_key_file=~/.ssh/id_ed25519_ansible
ntp_server=192.168.10.1
```

Vérifier l'inventaire :

```bash
ansible-inventory -i inventaires/production.ini --list
ansible-inventory -i inventaires/production.ini --graph
```

## 10. Inventaire statique au format YAML

Plus lisible dès que les variables se multiplient. Équivalent exact du INI précédent :

```yaml
# inventaires/production.yml
all:
  children:
    production:
      children:
        webservers:
          hosts:
            web-01:
              ansible_host: 192.168.10.11
            web-02:
              ansible_host: 192.168.10.12
        dbservers:
          hosts:
            db-01:
              ansible_host: 192.168.10.21
              ansible_user: dba
        proxmox:
          hosts:
            pve-01:
              ansible_host: 192.168.10.2
            pve-02:
              ansible_host: 192.168.10.3
      vars:
        ansible_ssh_private_key_file: ~/.ssh/id_ed25519_ansible
        ntp_server: 192.168.10.1
```

## 11. Groupes, groupes enfants et variables d'inventaire

Variables de connexion les plus utiles (à mettre dans l'inventaire ou `group_vars`) :

| Variable | Rôle |
|---|---|
| `ansible_host` | IP ou FQDN réel (le nom d'inventaire peut être logique) |
| `ansible_user` | Utilisateur SSH |
| `ansible_port` | Port SSH (défaut 22) |
| `ansible_ssh_private_key_file` | Clé privée dédiée |
| `ansible_become` / `ansible_become_user` | Élévation de privilèges par hôte/groupe |
| `ansible_python_interpreter` | Chemin Python cible (`/usr/bin/python3`) |
| `ansible_ssh_common_args` | Arguments SSH supplémentaires (ProxyJump…) |

Exemple bastion (hôtes derrière un jump host) :

```ini
[isoles]
srv-dmz-01 ansible_host=10.0.5.11

[isoles:vars]
ansible_ssh_common_args=-o ProxyJump=deploy@bastion.example.lan
```

## 12. Bonnes pratiques d'inventaire statique

- [ ] Un fichier par environnement (`production.ini`, `staging.ini`) — jamais de mélange.
- [ ] Les secrets **jamais** dans l'inventaire : utilisez `host_vars/` + Vault (section 46).
- [ ] Noms d'hôtes logiques (`web-01`) + `ansible_host` pour l'IP : l'IP peut changer, le rôle reste.
- [ ] Versionnez l'inventaire en Git, sauf les fichiers contenant des données sensibles.
- [ ] Validez après chaque modification : `ansible-inventory --graph` puis un `ping` ciblé.

## 13. Inventaire dynamique : plugin AWS EC2

L'inventaire dynamique interroge l'API du fournisseur à chaque exécution : les hôtes
apparaissent/disparaissent sans édition manuelle.

```yaml
# inventaires/dynamique/aws_ec2.yml
plugin: amazon.aws.aws_ec2
regions:
  - eu-west-3
# Ne prendre que les instances démarrées
filters:
  instance-state-name: running
# Construire des groupes à partir des tags
keyed_groups:
  - key: tags.Role
    prefix: role
    separator: ""
  - key: tags.Environment
    prefix: env
    separator: ""
# Variables d'hôte depuis les métadonnées AWS
hostvars_prefix: aws_
compose:
  ansible_host: public_ip_address | default(private_ip_address)
```

```bash
pipx inject ansible --include-apps boto3 botocore   # dépendances AWS
ansible-inventory -i inventaires/dynamique/aws_ec2.yml --graph
```

## 14. Inventaire dynamique : plugin Proxmox

Indispensable quand le parc virtualisé bouge souvent. Nécessite la collection
`community.general` et le module Python `proxmoxer`.

```yaml
# inventaires/dynamique/proxmox.yml
plugin: community.general.proxmox
url: https://pve-01.example.lan:8006
user: ansible@pve
password: "{{ vault_pve_api_password }}"   # Vault, jamais en clair !
validate_certs: false   # true en production avec un certificat valide
# Grouper par statut, template, type de nœud...
keyed_groups:
  - key: proxmox_status
    prefix: pve_status
  - key: proxmox_type
    prefix: pve_type
compose:
  ansible_host: proxmox_ipconfig0 | regex_replace('.*ip=([^,]+).*', '\\1')
```

```bash
ansible-galaxy collection install community.general
pipx inject ansible proxmoxer requests
ansible-inventory -i inventaires/dynamique/proxmox.yml --list | head -50
```

> **Avertissement** : `validate_certs: false` n'est acceptable qu'en labo. En production,
> déployez un certificat Let's Encrypt ou interne sur Proxmox et passez à `true`.

## 15. Inventaire dynamique : script personnalisé

Quand aucun plugin n'existe (CMDB maison, NetBox sans plugin…), un script exécutable
qui imprime du JSON `--list` fait l'affaire. Exemple minimal :

```python
#!/usr/bin/env python3
"""inventaires/dynamique/cmdb.py — inventaire depuis la CMDB interne."""
import json, sys

def main():
    inventory = {
        "webservers": {
            "hosts": ["web-01", "web-02"],
            "vars": {"http_port": 80},
        },
        "_meta": {
            "hostvars": {
                "web-01": {"ansible_host": "192.168.10.11"},
                "web-02": {"ansible_host": "192.168.10.12"},
            }
        },
    }
    if "--list" in sys.argv:
        print(json.dumps(inventory))
    elif "--host" in sys.argv:
        print(json.dumps({}))

if __name__ == "__main__":
    main()
```

```bash
chmod +x inventaires/dynamique/cmdb.py
ansible-inventory -i inventaires/dynamique/cmdb.py --graph
```

## 16. Première commande ad-hoc

Les commandes ad-hoc exécutent **un module** sans playbook — parfait pour l'exploitation
quotidienne et les vérifications rapides.

```bash
# Ping Ansible (pas ICMP) sur tout le parc
ansible all -m ansible.builtin.ping

# Temps de disponibilité des serveurs web
ansible webservers -m ansible.builtin.command -a "uptime"

# Espace disque avec sudo
ansible all -b -m ansible.builtin.command -a "df -h /"

# Redémarrer un service sur les DB
ansible dbservers -b -m ansible.builtin.service -a "name=postgresql state=restarted"
```

Options ad-hoc à connaître : `-b` (become), `-K` (demande le mot de passe sudo),
`-f 20` (forks), `-l web-01` (limiter à un hôte), `--check` (dry-run).

## 17. Dix commandes ad-hoc utiles au quotidien

```bash
# 1. Qui est connecté ?
ansible all -m ansible.builtin.command -a "who"

