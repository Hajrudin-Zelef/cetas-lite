---
id: collect-261001-general-networking/general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026-3
title: "Redémarre l'agent et relance les scans SCA"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention", "mai", "valuation"]
source: docs/RAG/collect-261001-general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [151, 218]
sha256: 40c1cb48b8891bdbd8c8b3bc5091a09e704a7bf55870bc1299600d709ffeaac7
---

# Redémarre l'agent et relance les scans SCA

La surveillance d’intégrité des fichiers (File Integrity Monitoring) détecte toute création, modification ou suppression dans les répertoires sensibles – un signal fort de compromission ou d’action malveillante. C’est aussi une exigence explicite de nombreux référentiels de conformité (PCI DSS, ISO 27001). Sur l’agent `web01`, éditez `/var/ossec/etc/ossec.conf` et adaptez le bloc `<syscheck>` :

```
<syscheck>
  <disabled>no</disabled>
  <frequency>43200</frequency>
  <directories check_all="yes" realtime="yes">/etc,/usr/bin,/usr/sbin</directories>
  <directories check_all="yes" whodata="yes">/home/admin/.ssh,/var/www</directories>
  <ignore>/etc/mtab</ignore>
  <ignore>/etc/random-seed</ignore>
  <alert_new_files>yes</alert_new_files>
</syscheck>
```
L’attribut `realtime="yes"` déclenche une alerte instantanée grâce à l’API inotify du noyau ; `whodata="yes"` va plus loin en identifiant *quel utilisateur* a modifié le fichier, via le sous-système d’audit Linux. Redémarrez l’agent pour appliquer :

`sudo systemctl restart wazuh-agent`
Testez immédiatement en modifiant un fichier surveillé, par exemple `sudo touch /etc/test-fim.conf`. En quelques secondes, une alerte remonte dans le tableau de bord, module *File Integrity Monitoring*, avec le chemin, l’utilisateur, l’empreinte avant/après et l’horodatage. Attention toutefois : activer le temps réel sur d’énormes arborescences (comme `/var` entier) génère une charge CPU importante – ciblez les répertoires réellement sensibles.

## Étape 6 – Activer la détection de vulnérabilités

Le module de détection de vulnérabilités de Wazuh croise l’inventaire logiciel remonté par chaque agent avec des flux de CVE (bases NVD et catalogues éditeurs). Il signale ainsi les paquets obsolètes ou vulnérables sur l’ensemble du parc, sans installer de scanner tiers. Depuis la refonte du moteur (versions 4.8 et suivantes), le module a continué de gagner en précision : la version 4.11.0 (20 février 2025) a amélioré la détection de vulnérabilités et la fiabilité de l’inventaire logiciel, tandis que la 4.12.0 (7 mai 2025) a ajouté le support des composants centraux sur architecture ARM et migré vers OpenSearch 2.19.1. La configuration se fait côté manager dans `/var/ossec/etc/ossec.conf` :

```
<vulnerability-detection>
  <enabled>yes</enabled>
  <index-status>yes</index-status>
  <feed-update-interval>60m</feed-update-interval>
</vulnerability-detection>
```
Redémarrez le manager avec `sudo systemctl restart wazuh-manager`. Au premier lancement, Wazuh télécharge les flux de vulnérabilités (plusieurs minutes) puis évalue l’inventaire de chaque agent. Les résultats apparaissent dans le module *Vulnerability Detection*, classés par sévérité (Critique, Élevée, Moyenne, Faible) avec le CVE, le paquet concerné et la version corrective recommandée.

C’est un atout majeur pour la conformité NIS 2, qui impose une gestion active des vulnérabilités. Vous obtenez, sans licence supplémentaire, une cartographie continue de votre dette de sécurité, exportable en rapport. Couplez-le au module SCA (étape suivante) pour une vision complète : « quels logiciels sont vulnérables » d’un côté, « quelles configurations sont dangereuses » de l’autre.

## Étape 7 – Évaluer la configuration avec le module SCA (CIS)

Le module SCA (Security Configuration Assessment) audite automatiquement vos systèmes par rapport aux référentiels de durcissement CIS (Center for Internet Security). Il vérifie des centaines de points de contrôle : permissions de fichiers, paramètres SSH, politiques de mot de passe, services inutiles, etc. Wazuh embarque des politiques prêtes à l’emploi pour Ubuntu, Debian, RHEL, Windows Server (avec une politique SCA dédiée à Windows Server 2025 introduite dans la 4.14.2 du 14 janvier 2026) et macOS.

Le module est activé par défaut ; ses résultats sont visibles dans le tableau de bord, module *Security Configuration Assessment*, pour chaque agent. Chaque politique affiche un score de conformité (par exemple « 68 % – 214 réussis, 99 échoués ») et détaille chaque contrôle échoué avec la remédiation exacte. Pour forcer une réévaluation immédiate depuis l’agent :

```
# Redémarre l'agent et relance les scans SCA
sudo systemctl restart wazuh-agent
# Lister les politiques SCA disponibles
ls /var/ossec/ruleset/sca/
```
Traitez en priorité les contrôles « Élevés » et les échecs liés à SSH, aux comptes et aux permissions. Chaque correction améliore mécaniquement votre score et réduit votre surface d’attaque. Un serveur passant de 60 % à 90 % de conformité CIS, c’est concrètement des dizaines de portes d’entrée fermées aux attaquants.

## Comprendre les niveaux de règles et la corrélation Wazuh

Avant de déclencher une détection, il faut saisir comment Wazuh « pense ». Chaque événement remonté par un agent suit une chaîne de traitement : un *décodeur* extrait les champs pertinents du journal brut (IP source, utilisateur, action), puis le moteur de *règles* évalue ces champs. Si une règle correspond, Wazuh lui attribue un **niveau de gravité** et, au-delà d’un seuil, génère une alerte. Comprendre ces niveaux est indispensable pour régler le bruit et prioriser les réponses.

Les règles Wazuh sont classées sur une échelle de 0 à 15. Plus le niveau est élevé, plus l’événement est grave. Voici la grille de lecture généralement retenue par les analystes.

| Niveau | Gravité | Signification | Exemple typique | 
|---|---|---|---|
| 0 à 3 | Faible | Événements informatifs ou autorisés | Connexion réussie, notification système | 
| 4 à 7 | Moyenne | Erreurs et anomalies mineures | Échec d’authentification isolé, erreur applicative | 
| 8 à 11 | Élevée | Activité suspecte, attaque probable | Force brute, modification d’un fichier critique | 
| 12 à 15 | Critique | Compromission probable ou attaque réussie | Rootkit détecté, élévation de privilèges | 

La vraie force de Wazuh réside dans la **corrélation**. Un unique échec de connexion SSH (règle 5716, niveau 5) n’a rien d’alarmant. Mais une règle *composite*, dotée d’attributs `frequency` et `timeframe`, surveille la répétition : si le même IP source déclenche par exemple huit échecs en 120 secondes, Wazuh remonte automatiquement l’événement au niveau 10 avec la règle 5712 « brute force ». C’est ce mécanisme de fenêtre temporelle qui transforme un flot d’événements anodins en une alerte de sécurité exploitable – et c’est exactement ce que nous allons déclencher à l’étape suivante.

Conseil pratique : par défaut, seules les alertes de niveau 3 et plus sont indexées. Ajustez ce seuil (`<alerts><log_alert_level>`) selon votre tolérance au bruit. Trop bas, vous noyez les analystes ; trop haut, vous manquez des signaux faibles. Un réglage à 5 ou 7 constitue souvent un bon compromis pour démarrer, quitte à l’affiner en observant vos propres volumes d’alertes.

## Étape 8 – Détecter une attaque SSH par force brute

Passons au cœur du projet : la détection en temps réel. Wazuh embarque des milliers de règles prêtes à l’emploi, dont un jeu complet pour `sshd`. Par défaut, l’agent surveille déjà `/var/log/auth.log` (Ubuntu/Debian) ou `/var/log/secure` (RHEL). Vérifiez la présence de ce bloc dans l’`ossec.conf` de l’agent :

