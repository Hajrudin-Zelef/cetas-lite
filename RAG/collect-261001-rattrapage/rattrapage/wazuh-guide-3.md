---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-3
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["open source", "agent", "agents", "apache", "arr"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [94, 234]
sha256: 5c67983f3e869fef170c965f22ebf176657ea673048c18f81e43e8f8680b3973
---

# Guide Wazuh — SIEM & XDR Open Source en production

83. [Ports et flux réseau : tableau de référence](#83-ports-et-flux-réseau--tableau-de-référence)
84. [Fichiers importants : où est quoi](#84-fichiers-importants--où-est-quoi)
85. [Glossaire](#85-glossaire)
86. [Quiz : 10 questions + réponses](#86-quiz--10-questions--réponses)
87. [Pour aller plus loin](#87-pour-aller-plus-loin)

---

## 1. Ce qu'est Wazuh et à quoi ça sert

Wazuh est une plateforme **open source** de sécurité qui combine :

- **SIEM** (*Security Information and Event Management*) : collecte et corrélation centralisée des logs et événements de sécurité.
- **XDR** (*Extended Detection and Response*) : détection et réponse sur les postes/serveurs via des agents.
- **FIM** (*File Integrity Monitoring*) : surveillance de l'intégrité des fichiers.
- **Détection de vulnérabilités** : inventaire des logiciels installés croisé avec les bases CVE.
- **Audit de conformité** : contrôles CIS, PCI-DSS, etc.
- **Réponse active** : blocage automatique (IP, compte, etc.).

Concrètement, vous installez un **agent** sur chaque machine (Linux, Windows, macOS), les agents envoient leurs événements à un **manager** central, le manager analyse (règles, decoders), stocke dans un **indexer**, et vous visualisez tout dans un **dashboard** web.

Cas d'usage typiques pour un service systèmes & énergies :

| Besoin | Ce que Wazuh apporte |
|---|---|
| Savoir qui s'est connecté où (SSH, RDP) | Alertes de connexion, géolocalisation, brute-force |
| Détecter un fichier modifié sur un serveur critique | FIM sur `/etc`, `C:\Windows\System32` |
| Suivre les CVE du parc | Inventaire logiciel + alertes CVE quotidiennes |
| Prouver la conformité (audit) | Rapports CIS, PCI-DSS exportables |
| Réagir vite la nuit | Active Response : bannir l'IP attaquante automatiquement |
| Centraliser les logs des équipements réseau | Collecte syslog (switchs, onduleurs, pare-feu) |

Wazuh ne remplace pas un antivirus ni un pare-feu : il les **complète** en donnant la vision centrale.

---

## 2. Architecture : manager, indexer, dashboard

Trois composants centraux, trois rôles distincts. Retenez ce schéma :

```
┌─────────────┐      ┌─────────────┐      ┌──────────────┐
│   AGENTS    │─────▶│   MANAGER   │─────▶│   INDEXER    │
│ (postes,    │ 1514 │ (analyse :  │ 9200 │ (stockage :  │
│  serveurs)  │/udp  │  règles,    │/tcp  │  OpenSearch) │
└─────────────┘      │  decoders,  │      └──────┬───────┘
                     │  FIM, SCA)  │             │
┌─────────────┐      └──────┬──────┘             │
│ SYSLOG      │─────▶│      │ Filebeat           │
│ (équipements│ 514  │      ▼                    ▼
│  réseau)    │/udp  │  ┌──────────────┐  ┌──────────────┐
└─────────────┘      │  │  DASHBOARD   │◀─│  (requêtes)  │
                     │  │ (IHM web :   │  │              │
                     │  │  443)        │  │              │
                     │  └──────────────┘  └──────────────┘
                     └─────────────────────────────┘
```

### 2.1 Le manager (`wazuh-manager`)

- Reçoit les événements des agents (port **1514**, UDP par défaut en 4.x ; TCP possible).
- Applique les **decoders** (parse les logs) puis les **règles** (détecte).
- Exécute les modules : FIM (syscheck), détection de vulnérabilités, SCA, rootcheck, active response.
- Expose l'**API** (port **55000**) pour le dashboard et l'automatisation.
- Envoie les alertes vers l'indexer via **Filebeat**.

### 2.2 L'indexer (`wazuh-indexer`)

- Basé sur **OpenSearch** : stocke et indexe les alertes + les événements.
- Port **9200** (API REST). En cluster : plusieurs nœuds.
- C'est lui qui grossit en disque : dimensionnez large (voir section 5).

### 2.3 Le dashboard (`wazuh-dashboard`)

- Basé sur **OpenSearch Dashboards** : interface web (port **443**).
- Affiche alertes, agents, vulnérabilités, conformité, MITRE ATT&CK.
- S'authentifie auprès de l'indexer ; ne stocke rien lui-même.

### 2.4 Filebeat, le tuyau discret

Filebeat lit `/var/ossec/logs/alerts/alerts.json` sur le manager et l'envoie à l'indexer. Si Filebeat est arrêté, **le manager continue d'analyser** mais le dashboard ne voit plus rien. Symptôme classique : « le dashboard est vide alors que les agents sont verts » → vérifiez Filebeat (voir section 70).

---

## 3. Concepts clés à maîtriser avant d'installer

| Concept | Définition courte | Où ça vit |
|---|---|---|
| **Agent** | Programme sur la machine surveillée | `/var/ossec` (Linux), `C:\Program Files (x86)\ossec-agent` (Windows) |
| **Decoder** | Parse un log brut → champs structurés | `/var/ossec/etc/decoders/local_decoder.xml` |
| **Règle** | Condition sur les champs → alerte + niveau | `/var/ossec/etc/rules/local_rules.xml` |
| **Niveau** | Sévérité 0–16 (voir section 30) | champ `rule.level` |
| **Groupe** | Ensemble d'agents partageant une conf | `/var/ossec/etc/shared/<groupe>/` |
| **FIM / syscheck** | Sommes de contrôle des fichiers surveillés | module `syscheck` |
| **SCA** | Contrôles de conformité (CIS...) | module `sca`, politiques `.yml` |
| **Active Response** | Script exécuté sur alerte | `/var/ossec/active-response/bin/` |
| **CDB list** | Liste de valeurs pour enrichir les règles | `/var/ossec/etc/lists/` |

Retenez surtout : **decoder d'abord, règle ensuite**. Un log non décodé ne déclenchera jamais de règle fine.

---

## 4. Licence, coût et positionnement vs concurrents

- **Licence :** GPLv2 pour le cœur (fork d'OSSEC), composants OpenSearch sous Apache 2.0 / SSPL selon version. Usage commercial autorisé, sans limite d'agents ni d'EPS dans la version communautaire.
- **Coût :** 0 € de licence. Le coût réel = **serveurs + disque + votre temps**.
- **Support :** communautaire (Slack, GitHub, forum) gratuit ; support commercial via Wazuh Inc. si besoin.

Comparatif rapide (ordre de grandeur, à valider selon votre contexte) :

| Solution | Licence | Agents illimités | Prise en main |
|---|---|---|---|
| Wazuh | Open source | Oui | Moyenne (ce guide aide) |
| Wazuh Cloud | Abonnement | Oui (facturé/agent) | Facile |
| Splunk | Propriétaire, au volume | Non (coût/volume) | Moyenne |
| Elastic Security | Freemium (limites) | Selon offre | Moyenne |
| Graylog | Open core | Selon offre | Facile |

Pour une petite équipe avec un budget serré et des compétences Linux, **Wazuh auto-hébergé** est imbattable en rapport fonctionnalités/coût.

---

## 5. Dimensionnement : CPU, RAM, disque

Le dimensionnement dépend de deux métriques : **nombre d'agents** et **EPS** (*events per second*, événements/seconde).

### 5.1 Repères (all-in-one, valeurs indicatives Wazuh 4.x)

| Agents | EPS approx. | CPU | RAM | Disque |
|---|---|---|---|---|
| < 25 | < 50 | 4 vCPU | 8 Go | 100 Go |
| 25 – 100 | 50 – 200 | 8 vCPU | 16 Go | 300 Go |
| 100 – 500 | 200 – 1000 | 16 vCPU | 32 Go | 1 To |
| 500+ | > 1000 | Distribué (voir section 12) | — | — |

> Ces chiffres sont des **ordres de grandeur**. Un parc Windows verbeux (audit avancé) génère 3 à 5× plus d'EPS qu'un parc Linux sobre. Mesurez vos EPS réels après 1 semaine (voir section 59) et ajustez.

### 5.2 Le disque : le poste critique

L'indexer stocke **chaque alerte** (et, selon réglage, chaque événement archivé). Règle pratique :

