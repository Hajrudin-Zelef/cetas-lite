---
id: collect-261001-general-networking/general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026-6
title: "Redémarre l'agent et relance les scans SCA"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "agents", "open source"]
source: docs/RAG/collect-261001-general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [360, 383]
sha256: be62341b6a422b619a6d46f364c85481d908d4b00dcdfdfafc3fe18aa157e7e6
---

# Redémarre l'agent et relance les scans SCA

Les composants centraux (manager, indexeur, tableau de bord) nécessitent un serveur Linux 64 bits. En revanche, l’agent Wazuh s’installe sur Windows (via un MSI), macOS, Solaris, AIX et HP-UX, en plus de Linux. Vous pouvez donc surveiller un parc mixte Windows/Linux depuis un unique serveur Wazuh.

### Wazuh aide-t-il à la conformité NIS 2 ?

Oui, directement. Wazuh couvre la détection d’incidents, la journalisation, la gestion des vulnérabilités et la production de preuves horodatées – autant d’exigences de NIS 2. Étant auto-hébergé, il garantit aussi la souveraineté des journaux, un atout pour le RGPD. Il ne remplace pas une démarche de conformité globale mais en constitue un pilier technique majeur.

### Faut-il un agent sur chaque machine ?

L’agent offre la visibilité la plus riche (FIM, SCA, inventaire, réponse active). Toutefois, Wazuh peut aussi ingérer des données sans agent via Syslog ou API, pour les équipements réseau, pare-feu ou services cloud qui n’acceptent pas d’agent. La stratégie idéale combine les deux approches.

### Quelle est la configuration matérielle minimale ?

Pour un serveur tout-en-un supervisant jusqu’à 25 agents : 4 vCPU, 8 Gio de RAM et 50 Go de disque, sous Ubuntu, RHEL ou Amazon Linux 64 bits. L’agent, lui, ne consomme qu’environ 35 Mo de RAM en moyenne, ce qui le rend déployable même sur des systèmes très contraints.

### Couverture associée

- CrowdSec : bloquer les attaques en 12 étapes – le complément idéal de Wazuh pour la mise en liste noire collaborative.
- Directive NIS 2 : 15 000 entités visées – le cadre réglementaire qui rend un SIEM incontournable.
- Tutoriel WireGuard : VPN Linux en 12 étapes – pour chiffrer les flux entre agents distants et manager.
- Tutoriel Keycloak 26 : SSO et OIDC – pour sécuriser l’accès au tableau de bord Wazuh.
- Vaultwarden : auto-héberger Bitwarden – dans la même logique d’outils de sécurité souverains.
- pfSense vs OPNsense 2026 – pour alimenter Wazuh avec les journaux de votre pare-feu.

Explorez l’ensemble de nos guides dans la rubrique Cybersécurité. Avec ce déploiement Wazuh opérationnel – installation, agents, FIM, détection de vulnérabilités, réponse active et conformité NIS 2 – vous disposez désormais d’un SIEM open source de niveau professionnel, souverain et sans coût de licence. Le code source et les dernières évolutions sont disponibles sur le dépôt GitHub officiel de Wazuh.
