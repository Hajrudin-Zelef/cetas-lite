---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-4
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [235, 454]
sha256: 1f09157eb639de086dced4dcc9ec17b1afaa7a1fefa3ac0dbaf85cd9c44801b4
---

# Guide Wazuh — SIEM & XDR Open Source en production

```
Espace disque ≈ EPS × 86400 s × taille_moyenne_doc × jours_rétention × 1,3 (marge)
```

Exemple : 100 EPS × 86400 × 1,5 Ko × 90 jours × 1,3 ≈ **1,5 To**.

- Séparez les données sur un volume dédié (`/var/lib/wazuh-indexer`).
- Prévoyez la **rotation des index** (voir section 60) pour ne pas saturer.
- En production : **ne mettez jamais l'indexer sur la partition racine**.

### 5.3 RAM de l'indexer

OpenSearch utilise la moitié de la RAM pour le *heap* Java (max ~32 Go). Le reste sert au cache du système de fichiers. **Sous-dimensionner la RAM = indexer lent puis rouge.**

---

## 6. Prérequis système et réseau

### 6.1 Système (serveur central)

- Debian 11/12 ou Ubuntu 20.04/22.04/24.04, 64 bits.
- `curl`, `gnupg`, `lsb-release` installés.
- Nom d'hôte résolvable (FQDN) — les certificats s'appuient dessus.
- **Désactivez tout autre service sur les ports 1514, 1515, 443, 9200, 55000** avant install.

```bash
# Vérification rapide des ports avant installation
sudo ss -tulpn | grep -E '1514|1515|443|9200|55000' || echo "Ports libres : OK"
```

### 6.2 Réseau et pare-feu

Flux à autoriser (détail complet section 83) :

| Source → Destination | Port | Protocole | Usage |
|---|---|---|---|
| Agents → Manager | 1514 | UDP (ou TCP) | Remontée événements |
| Agents → Manager | 1515 | TCP | Inscription |
| Admin → Manager | 55000 | TCP | API |
| Manager → Indexer | 9200 | TCP | Envoi alertes |
| Admin → Dashboard | 443 | TCP | IHM web |
| Dashboard → Indexer | 9200 | TCP | Requêtes |

### 6.3 Temps et DNS

- **NTP synchronisé partout** (manager, agents, indexer) : des horloges décalées cassent la corrélation et les certificats.
- DNS direct + reverse corrects sur le manager.

```bash
timedatectl status | grep "synchronized"
# doit afficher : System clock synchronized: yes
```

---

## 7. Méthodes de déploiement : choisir la bonne

| Méthode | Quand l'utiliser | Limites |
|---|---|---|
| **All-in-one** (assistant d'install) | < 100 agents, POC, PME | Ne scale pas au-delà de ~500 agents |
| **Distribué** (paquets pas à pas) | > 100 agents, HA, séparation des rôles | Plus complexe à maintenir |
| **Docker** (conteneurs) | Lab, tests, CI | **Déconseillé en production** par Wazuh (persistance, perf) |
| **Cluster multi-managers** | Continuité si un manager tombe | Nécessite un répartiteur de charge |

**Recommandation pour démarrer :** all-in-one sur une VM dédiée. C'est ce que décrit la section 9. Vous migrerez vers du distribué le jour où les métriques (section 59) le justifient.

---

## 8. All-in-one : le déploiement le plus simple

Wazuh fournit un **assistant d'installation** (`wazuh-install.sh`) qui installe et configure manager + indexer + dashboard + Filebeat en une commande, avec génération des certificats.

Principe :

```bash
curl -sO https://packages.wazuh.com/4.12/wazuh-install.sh
sudo bash wazuh-install.sh -a
```

L'option `-a` = all-in-one. L'assistant :
1. Ajoute le dépôt Wazuh.
2. Installe `wazuh-manager`, `wazuh-indexer`, `wazuh-dashboard`, `filebeat`.
3. Génère les certificats TLS (nœuds + admin).
4. Démarre les services dans le bon ordre.
5. Affiche les identifiants du dashboard.

> ⚠️ **Ne lancez jamais l'assistant deux fois sans désinstaller** (`-u`) entre les deux : vous corrompriez les certificats. En cas d'échec, désinstallez proprement puis recommencez.

---

## 9. Installation pas à pas sur Debian/Ubuntu (all-in-one)

### Étape 1 — Préparer le système

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y curl gnupg lsb-release ca-certificates
sudo hostnamectl set-hostname wazuh-soc       # adaptez
echo "127.0.0.1 wazuh-soc" | sudo tee -a /etc/hosts
```

Vérifiez que le FQDN se résout :

```bash
hostname -f
getent hosts "$(hostname -f)"
```

### Étape 2 — Télécharger l'assistant (version 4.x)

```bash
cd /tmp
curl -sO https://packages.wazuh.com/4.12/wazuh-install.sh
# Remplacez 4.12 par la dernière 4.x disponible sur https://documentation.wazuh.com
```

> Vérifiez toujours la dernière version mineure 4.x dans la documentation officielle avant de lancer : les correctifs de sécurité sortent régulièrement.

### Étape 3 — Lancer l'installation

```bash
sudo bash wazuh-install.sh -a
```

Durée indicative : 5 à 15 minutes selon la machine et le réseau. **Ne l'interrompez pas** (Ctrl+C en plein milieu = état incohérent → désinstallation requise).

### Étape 4 — Noter les identifiants affichés

En fin d'installation, l'assistant affiche quelque chose comme :

```
INFO: --- Summary ---
INFO: You can now start using your Wazuh cluster with:
INFO:   Wazuh dashboard: https://wazuh-soc
INFO:     User: admin
INFO:     Password: <mot-de-passe-généré>
```

**Copiez immédiatement ce mot de passe** dans votre coffre (Bitwarden, KeePass...). Il n'est affiché qu'une fois.

### Étape 5 — Sécuriser l'accès initial

```bash
# Changez le mot de passe admin dès la première connexion (dashboard > menu utilisateur)
# Puis régénérez proprement si besoin via le script fourni :
sudo tar -xzf wazuh-install-files.tar -C /tmp   # archive créée par l'assistant
```

L'archive `wazuh-install-files.tar` contient les certificats et mots de passe : **protégez-la** (droits 600, sauvegarde chiffrée) puis supprimez-la du serveur une fois archivée ailleurs.

### Étape 6 — Ouvrir le pare-feu (exemple UFW)

```bash
sudo ufw allow 1514/udp comment 'Wazuh agents'
sudo ufw allow 1515/tcp comment 'Wazuh enrollment'
sudo ufw allow 443/tcp  comment 'Wazuh dashboard'
sudo ufw allow 55000/tcp comment 'Wazuh API (restreindre aux admins)'
sudo ufw enable
```

> L'API (55000) et l'indexer (9200) ne doivent être exposés qu'aux IP d'administration, jamais à tout le parc.

### Étape 7 — Figer les versions (éviter les mises à jour surprises)

```bash
# Empêcher apt de mettre à jour Wazuh hors maintenance planifiée
sudo apt-mark hold wazuh-manager wazuh-indexer wazuh-dashboard filebeat
```

Les mises à jour se font **planifiées** (voir section 65), jamais au fil de l'eau.

---

## 10. Vérification post-installation

```bash
# 1. Les 4 services tournent ?
sudo systemctl is-active wazuh-manager wazuh-indexer wazuh-dashboard filebeat

# 2. Le manager écoute les agents ?
sudo ss -tulpn | grep -E '1514|1515'

# 3. L'indexer répond ? (admin / mot de passe dashboard, -k car autosigné)
curl -k -u admin:'<MOT_DE_PASSE>' https://localhost:9200

# 4. Santé du cluster indexer
curl -k -u admin:'<MOT_DE_PASSE>' 'https://localhost:9200/_cluster/health?pretty'

# 5. L'API du manager répond ?
curl -k -u wazuh-wui:'<MOT_DE_PASSE_API>' https://localhost:55000/security/user/authenticate -X GET
```

Résultat attendu pour la santé : `"status" : "green"` (ou `yellow` en mono-nœud, c'est normal : voir section 60).

Checklist de validation :

- [ ] Les 4 services sont `active`
- [ ] `curl` sur `:9200` répond 200
- [ ] Le dashboard s'affiche sur `https://<ip-ou-fqdn>`
- [ ] La page *Agents* du dashboard est accessible (0 agent pour l'instant, c'est normal)
- [ ] `wazuh-install-files.tar` sauvegardé hors serveur

---

## 11. Accéder au dashboard : première connexion

1. Ouvrez `https://<ip-ou-fqdn-du-serveur>` (acceptez l'avertissement de certificat autosigné, ou installez votre CA — voir section 16).
2. Connectez-vous avec `admin` / mot de passe généré.
3. **Changez immédiatement le mot de passe** : menu haut-droit → *Security* → utilisateurs → `admin`.

### 11.1 Créer des utilisateurs (principe du moindre privilège)

Ne travaillez jamais au quotidien avec `admin`. Créez :

| Utilisateur | Rôle | Usage |
|---|---|---|
| `admin` | Administrateur | Urgences, mises à jour |
| `soc-analyste` | `wazuh-ui-user` + lecture alertes | Analyse quotidienne |
| `soc-responsable` | Rôle avec active response | Validation des réponses |

