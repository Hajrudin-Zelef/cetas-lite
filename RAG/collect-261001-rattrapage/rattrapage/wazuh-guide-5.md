---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-5
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [455, 679]
sha256: 1db4cc9450536e95facc1f6da8a6eaf80b36ecb0c3d4010e8a0b042d7fc7669b
---

# Guide Wazuh — SIEM & XDR Open Source en production

Via le dashboard : *Security → Internal users → Create user*, puis assignez les rôles. Via API (voir section 58).

### 11.2 Régler le fuseau horaire

Dashboard → *Stack Management → Advanced Settings* → `dateFormat:tz` → `Europe/Paris` (ou votre TZ). Des alertes affichées en UTC alors que vos équipes pensent en heure locale = erreurs d'interprétation garanties.

---

## 12. Déploiement distribué : quand et comment

Passez en distribué quand : > 100–150 agents, besoin de HA, ou séparation réseau (DMZ).

### 12.1 Topologie type

```
                ┌──────────────────┐
                │  Load Balancer   │  (optionnel, ex. HAProxy)
                │  1514/udp 1515   │
                └───────┬──────────┘
                        │
           ┌────────────┼────────────┐
           ▼            ▼            ▼
     ┌──────────┐ ┌──────────┐ ┌──────────┐
     │ Manager  │ │ Manager  │ │ Manager  │  cluster Wazuh
     │    1     │ │    2     │ │    3     │
     └────┬─────┘ └────┬─────┘ └────┬─────┘
          │            │            │
          └────────────┼────────────┘
                       ▼ 9200
              ┌─────────────────┐
              │ Indexer cluster │  3 nœuds OpenSearch
              │ (dédiés)        │
              └────────┬────────┘
                       ▼
              ┌─────────────────┐
              │ Dashboard (1-2) │
              └─────────────────┘
```

### 12.2 Installation pas à pas (méthode paquets)

L'assistant gère aussi le distribué via un fichier de configuration. Procédure (4.x) :

```bash
# 1. Sur la première machine, générer la config
curl -sO https://packages.wazuh.com/4.12/wazuh-install.sh
sudo bash wazuh-install.sh --generate-config-files
# → crée wazuh-install-files.tar + config.yml à éditer
```

```yaml
# config.yml : déclarez chaque nœud
nodes:
  indexer:
    - name: indexer-1
      ip: 10.0.0.11
    - name: indexer-2
      ip: 10.0.0.12
    - name: indexer-3
      ip: 10.0.0.13
  server:
    - name: wazuh-1
      ip: 10.0.0.21
      node_type: master
    - name: wazuh-2
      ip: 10.0.0.22
      node_type: worker
  dashboard:
    - name: dashboard
      ip: 10.0.0.31
```

```bash
# 2. Générer certificats + mots de passe
sudo bash wazuh-install.sh --generate-config-files
# 3. Copier wazuh-install-files.tar sur chaque nœud, puis :
sudo bash wazuh-install.sh --indexer --node-name indexer-1      # sur chaque indexer
sudo bash wazuh-install.sh --server --node-name wazuh-1         # sur chaque manager
sudo bash wazuh-install.sh --dashboard --node-name dashboard    # sur le dashboard
# 4. Initialiser le cluster (sur un nœud indexer) :
sudo bash wazuh-install.sh --start-cluster
```

> Suivez **l'ordre** : indexers → démarrage cluster → managers → dashboard. Inverser l'ordre = erreurs de connexion TLS.

---

## 13. Déploiement conteneurs (Docker)

### 13.1 Quand l'utiliser

Lab, formation, tests de règles, CI. **Pas la production** : la documentation Wazuh elle-même recommande les paquets pour la prod (persistance des données, performances, support).

### 13.2 Lancement rapide (lab)

```bash
# Prérequis : Docker + Docker Compose
git clone https://github.com/wazuh/wazuh-docker.git
cd wazuh-docker/single-node
docker compose up -d
```

Accès : `https://localhost` (admin / `SecretPassword` par défaut — **changez-le**).

### 13.3 Persistance des données

Montez des volumes pour ne pas tout perdre au `docker compose down` :

```yaml
# extrait docker-compose.yml (single-node)
services:
  wazuh.manager:
    volumes:
      - wazuh_api_configuration:/var/ossec/api/configuration
      - wazuh_etc:/var/ossec/etc
      - wazuh_logs:/var/ossec/logs
      - wazuh_queue:/var/ossec/queue
      - wazuh_var_multigroups:/var/ossec/var/multigroups
```

> En lab c'est confortable ; en prod, préférez la méthode paquets (section 9 ou 12).

---

## 14. Cluster de managers : multi-nœuds

Le cluster Wazuh synchronise entre managers : règles, decoders, groupes d'agents, CDB lists.

### 14.1 Fichier `/var/ossec/etc/ossec.conf` (extrait cluster)

```xml
<cluster>
  <name>wazuh-production</name>
  <node_name>wazuh-1</node_name>
  <node_type>master</node_type>
  <key>MOT_DE_PASSE_PARTAGE_32_CARACTERES_MINIMUM</key>
  <port>1516</port>
  <bind_addr>0.0.0.0</bind_addr>
  <nodes>
      <node>10.0.0.21</node>
      <node>10.0.0.22</node>
  </nodes>
  <hidden>no</hidden>
  <disabled>no</disabled>
</cluster>
```

- **master** : 1 seul. C'est lui qui centralise la conf et la pousse aux workers.
- **workers** : reçoivent les agents (via LB) et synchronisent depuis le master.
- La clé `<key>` doit être **identique** sur tous les nœuds (32 caractères alphanumériques minimum).

### 14.2 Vérifier la synchronisation

```bash
sudo /var/ossec/bin/cluster_control -l          # liste les nœuds
sudo /var/ossec/bin/cluster_control -s           # statut de synchro
```

Si un worker affiche `synced: no` durablement : vérifiez la clé, le port 1516 entre managers, et les logs `/var/ossec/logs/cluster.log`.

---

## 15. Durcissement du serveur Wazuh (base)

Le serveur Wazuh est **la couronne** : qui le compromet voit tout le parc. Durcissez-le comme un bastion.

```bash
# 1. Mises à jour système (hors paquets Wazuh figés, cf. section 9)
sudo apt update && sudo apt upgrade -y

# 2. SSH : clés uniquement, pas de root
sudo sed -i 's/^#\?PermitRootLogin.*/PermitRootLogin no/' /etc/ssh/sshd_config
sudo sed -i 's/^#\?PasswordAuthentication.*/PasswordAuthentication no/' /etc/ssh/sshd_config
sudo systemctl reload ssh

# 3. Pare-feu restrictif : API et indexer réservés au réseau admin
sudo ufw allow from 10.0.10.0/24 to any port 55000 comment 'API Wazuh admins'
sudo ufw allow from 10.0.10.0/24 to any port 9200  comment 'Indexer admins'

# 4. Fail2ban sur SSH (et dashboard si exposé)
sudo apt install -y fail2ban
```

### 15.1 Durcir l'API Wazuh

`/var/ossec/api/configuration/api.yaml` :

```yaml
host: 127.0.0.1          # n'écouter qu'en local + reverse proxy, si possible
port: 55000
https:
  enabled: yes
  key: api/configuration/ssl/server.key
  cert: api/configuration/ssl/server.crt
access:
  max_login_attempts: 5
  block_time: 300
  max_request_per_minute: 300
```

Redémarrez le manager après modification : `sudo systemctl restart wazuh-manager`.

### 15.2 Comptes et mots de passe

- Changez `admin` dashboard (section 11), `wazuh-wui` (API dashboard) et `kibanaserver`.
- Mot de passe API dashboard : `/usr/share/wazuh-dashboard/data/wazuh/config/wazuh.yml` côté dashboard + utilisateur `wazuh-wui` côté manager — **les deux doivent correspondre**.

---

## 16. Certificats TLS : comprendre et régénérer

Wazuh 4.x chiffre : agents↔manager (clé symétrique par agent, voir section 21), manager↔indexer (TLS), dashboard↔indexer (TLS), navigateur↔dashboard (TLS).

### 16.1 Où sont les certificats (all-in-one)

| Usage | Emplacement |
|---|---|
| Indexer (nœud) | `/etc/wazuh-indexer/certs/` |
| Dashboard → indexer | `/etc/wazuh-dashboard/certs/` |
| Filebeat → indexer | `/etc/filebeat/certs/` |

### 16.2 Régénérer après changement de FQDN/IP

Si vous changez le nom ou l'IP du serveur, les certificats (liés au CN/SAN) deviennent invalides :

