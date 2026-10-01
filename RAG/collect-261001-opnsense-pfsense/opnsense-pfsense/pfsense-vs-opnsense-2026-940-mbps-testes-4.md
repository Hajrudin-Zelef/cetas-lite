---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/pfsense-vs-opnsense-2026-940-mbps-testes-4
title: "pfsense-vs-opnsense-2026-940-mbps-testes"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "open source"]
source: docs/RAG/collect-261001-opnsense-pfsense/pfsense-vs-opnsense-2026-940-mbps-testes.md
source_anchor: ""
source_lines: [141, 197]
sha256: 196f2c8cd3fc21f69ede76e3e6b38cc239ca049d467c4848c7812a2c6ac23c86
---

# pfsense-vs-opnsense-2026-940-mbps-testes

**Patrick Kennedy (ServeTheHome)** a publié en février 2026 un test croisé sur l’appliance Netgate 6100 et la Deciso DEC2750, concluant : « Les deux atteignent 9 Gbit/s en routage. Le choix se résume désormais à votre relation contractuelle préférée – Netgate aux États-Unis ou Deciso aux Pays-Bas – et à votre tolérance pour les mises à jour fréquentes. »

Sur Reddit, le subreddit *r/homelab* (3,2 millions de membres) compte en avril 2026 environ **62 % de mentions OPNsense** contre 38 % pfSense dans les fils de discussion sur les pare-feu – une inversion historique de la tendance qui prévalait avant 2022. Le subreddit *r/PFSENSE* (90 000 membres) reste fidèle à pfSense mais reconnaît dans son wiki officiel qu’« OPNsense est un excellent choix alternatif ».

**Jeff Geerling**, ingénieur Ansible et créateur YouTube reconnu (810 000 abonnés), recommande OPNsense pour ses pipelines automation : « L’API REST native d’OPNsense est tellement plus propre que celle de pfSense que j’ai migré tous mes playbooks Ansible en deux soirées. » Cette qualité d’API est un point clé pour les administrateurs qui pratiquent l’infrastructure as code avec Ansible.

## Cinq exemples concrets de déploiement en 2026

**1. Académie de Versailles (Éducation nationale, France)** – En février 2026, l’académie a migré 47 sites scolaires de pfSense CE 2.6 vers OPNsense 25.1 dans le cadre de la directive de souveraineté numérique. Argument décisif : licence BSD-2-Clause, hébergement Deciso aux Pays-Bas et compatibilité avec La Suite Numérique d’Etat. Coût total de la migration : 0 € en licences, 184 000 € en matériel Deciso DEC2700.

**2. Banque Postale Filiale Asset Management (France)** – Conserve pfSense Plus 24.03 sur 12 sites, justifié par le contrat de support TAC Premium signé avec Netgate Europe en 2023. Le contrat de cinq ans expire en 2028, après quoi la banque évalue OPNsense Business Edition.

**3. Université de Munich (LMU Munich)** – Déploiement homogène d’OPNsense sur 23 segments de campus depuis 2022, motivé par l’*open source policy* du Land de Bavière. La LMU contribue activement au plugin *os-collectd* et a versé 18 000 € à Deciso en mécénat sur 2025.

**4. CHU de Bordeaux (santé publique)** – Pare-feu segmentation DMZ médicale sous pfSense Plus 24.03 sur 4 appliances Netgate 8300 redondantes (HA cluster), choix justifié par le support 24×7 critique pour la continuité des soins. Dépenses 2026 : 31 200 $ (4 x 3 499 $ + 4 x 799 $ x 5 ans).

**5. OVHcloud (datacenters Roubaix et Strasbourg)** – Utilise OPNsense en tant que *customer edge firewall* sur l’offre Public Cloud Gateway depuis 2024. Plus de 14 000 instances OPNsense en production dans l’écosystème Cloud Souverain européen, choisies pour leur licence BSD permissive permettant la commercialisation du service managé.

## Cas d’usage et recommandations 2026

Voici les sept scénarios les plus fréquents et nos recommandations :

- **Homelab débutant (1 à 5 utilisateurs)** –*OPNsense Community* sur un Mini PC Topton 4-port à 220 €. UI moderne, 80+ plugins, mises à jour bi-hebdomadaires. Verdict : choisir OPNsense.
- **Homelab avancé (3 à 10 utilisateurs, 1 Gbit/s fibre)** –*pfSense CE 2.7.2* sur un Protectli FW6E pour bénéficier de pfBlockerNG Devel. Verdict : choisir pfSense.
- **PME 50-200 postes** –*OPNsense Business Edition* sur Deciso DEC2700, 159 €/an de support et licence BSD. Verdict : choisir OPNsense.
- **PME sous contrat existant Netgate** –*pfSense Plus 24.x* sur Netgate 4100/6100, conserver l’écosystème. Verdict : choisir pfSense.
- **Administration publique européenne** –*OPNsense Community ou Business* , hébergement Deciso UE, conformité NIS 2 et souveraineté. Verdict : choisir OPNsense.
- **Datacenter multi-tenant 10 Gbit/s+** –*pfSense Plus* sur Netgate 8300 pour le tuning netmap natif et le support 24×7. Verdict : choisir pfSense Plus.
- **Pipeline DevOps + Ansible/Terraform** –*OPNsense Community* pour l’API REST native exhaustive. Verdict : choisir OPNsense.

## Guide de migration : passer de pfSense à OPNsense en 8 étapes

La migration d’un déploiement pfSense vers OPNsense est techniquement possible, OPNsense fournissant un importateur officiel pour les sauvegardes XML pfSense. Voici la procédure standard sur une appliance de remplacement :

- **Étape 1** – Sauvegarder la configuration pfSense actuelle :*Diagnostics → Backup & Restore → Download configuration as XML* .
- **Étape 2** – Préparer l’appliance OPNsense (Deciso DEC ou Mini PC). Télécharger l’image VGA 25.7 sur opnsense.org et flasher la clé USB avec dd ou Rufus.
- **Étape 3** – Installer OPNsense (mode UFS ou ZFS), connecter sur l’IP par défaut 192.168.1.1, login root / opnsense.
- **Étape 4** – Importer la configuration pfSense :*System → Configuration → Backups → Restore → Import pfSense XML* . L’importeur supporte les règles, NAT, alias, certificats TLS et VPN site-to-site IPsec/OpenVPN. WireGuard et certains packages tiers (pfBlockerNG) doivent être reconfigurés manuellement.
- **Étape 5** – Vérifier la cohérence des règles via*Firewall → Rules → Inspect* . Comparer avec la sortie`pfctl -sr` de l’ancien pfSense.
- **Étape 6** – Installer les plugins équivalents :*System → Firmware → Plugins* . Pour pfBlockerNG, l’équivalent OPNsense est*os-bind* +*os-bsdinstaller* ou*Adguard Home* .
- **Étape 7** – Test de bascule en heures creuses : déconnecter le pfSense, brancher l’OPNsense sur le même WAN/LAN, vérifier la connectivité.
- **Étape 8** – Conserver l’ancien pfSense en standby pendant 30 jours minimum. Activer le monitoring SNMP et Telegraf vers une stack Grafana ou Datadog pour comparer les métriques avant/après.

Sur un parc moyen de 5 appliances, la migration prend 8 à 12 heures-ingénieur, avec une fenêtre de bascule réelle de 30 à 45 minutes par site. À noter : la migration inverse (OPNsense vers pfSense) n’est pas officiellement supportée et exige une reconfiguration manuelle complète.

## Forces et faiblesses : la liste pour/contre

### pfSense : pour et contre

**Avantages** : pfBlockerNG Devel reste la référence pour les blocklists DNS et GeoIP. Documentation très exhaustive sur docs.netgate.com. Communauté historique, base utilisateurs estimée à 3 millions de déploiements actifs. Support commercial 24×7 disponible via Netgate. Tuning netmap natif pour les charges 10 Gbit/s+. Appliances Netgate certifiées et largement disponibles auprès des intégrateurs.

**Inconvénients** : licence Plus propriétaire depuis 2021. Cycle de mises à jour annuel, lent pour les correctifs FreeBSD-SA. Interface Bootstrap 3 datée. API REST native uniquement sur Plus 23.09+. Société Netgate basée aux États-Unis, soumise au CLOUD Act. Pas d’AmneziaWG. Tailscale via package non officiel. Outils de visualisation traffic shaping moins avancés qu’OPNsense.

### OPNsense : pour et contre

**Avantages** : licence BSD-2-Clause intégralement open source. Cycle de mises à jour bi-hebdomadaire, intégration FreeBSD-SA en 4,2 jours en moyenne. Interface MVC Phalcon moderne et accessible. API REST native exhaustive depuis 2018. Société Deciso basée aux Pays-Bas, juridiction européenne. Support natif d’AmneziaWG, Tailscale, Zenarmor. Traffic shaping pipes/queues plus performant en QoS. Reporting Insight intégré sans plugin. Support communautaire bilingue (français inclus).

