---
id: collect-261001-general-networking/general-networking/clamav-antivirus-open-source-linux-10-etapes-2026-5
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/clamav-antivirus-open-source-linux-10-etapes-2026.md
source_anchor: ""
source_lines: [372, 414]
sha256: 48dab26f342b70ae368cf855b7364ce867fca7d10e35d7a8288a6597d5f1d1e9
---

# Debian / Ubuntu

Trois points de vigilance pour un dossier de conformité solide : documentez la politique de mise à jour des versions (branche LTS choisie, date de migration prévue avant l’échéance de fin de support), tracez la surveillance des vulnérabilités connues avec preuve d’application des correctifs (comme le passage aux versions 1.5.4 ou 1.4.6 d’août 2026), et conservez les logs de détection pendant une durée cohérente avec votre politique de rétention interne. Ces trois éléments, combinés à une authentification multifacteur sur les accès administratifs et une protection contre le brute-force via un outil comme Fail2ban, forment une base défendable devant un auditeur NIS2.

## Astuces avancées pour optimiser ClamAV en production

Une fois l’installation de base stabilisée, plusieurs réglages permettent d’affiner les performances et la couverture de détection sur un serveur à fort trafic.

Sur les serveurs multi-cœurs, augmentez progressivement `MaxThreads` par palier de deux en surveillant la charge CPU et mémoire avec `htop`, plutôt que de fixer une valeur arbitraire élevée d’emblée. Pour les environnements où le scan d’archives imbriquées ralentit le traitement, ajustez `MaxRecursion` dans clamd.conf pour limiter la profondeur d’analyse des archives contenues dans d’autres archives, un vecteur d’attaque connu sous le nom de « zip bomb ». Si votre infrastructure combine ClamAV avec une passerelle réseau plus large, envisagez également une association avec un pare-feu applicatif comme OPNsense pour filtrer le trafic en amont, réduisant d’autant le volume de fichiers que ClamAV doit traiter.

Pour les organisations exploitant plusieurs serveurs ClamAV, un miroir de signatures interne via `cvdupdate` évite de multiplier les requêtes sortantes vers les serveurs officiels ClamAV, un gain particulier sur les connexions internet limitées ou fortement filtrées. Enfin, si votre passerelle mail gère un volume important, envisagez de combiner ClamAV avec un système de protection contre les tentatives d’intrusion réseau comme CrowdSec, qui couvre un périmètre complémentaire (comportement réseau) à celui de ClamAV (contenu de fichiers).

## Questions fréquentes

**ClamAV est-il un antivirus complet pour poste de travail Windows ou macOS ?**

Non. ClamAV cible en priorité les serveurs et passerelles (mail, fichiers, uploads). Sur un poste de travail, il ne remplace pas une suite de sécurité complète avec protection en temps réel avancée et détection comportementale.

**Quelle est la différence entre clamscan et clamdscan ?**

clamscan recharge la base de signatures à chaque exécution, ce qui le rend lent sur de gros volumes. clamdscan délègue le scan au démon clamd déjà chargé en mémoire, donc nettement plus rapide pour les scans répétés ou planifiés.

**Faut-il choisir la branche 1.4 LTS ou la branche 1.5 stable ?**

La branche 1.4 LTS est supportée jusqu’au 15 août 2027 et convient aux environnements qui privilégient un cycle de mise à jour prévisible. La branche 1.5, taguée « latest » et « stable » par l’équipe ClamAV, apporte les dernières optimisations, notamment la réduction de taille des signatures amorcée fin 2025.

**ClamAV peut-il bloquer un mail contenant un virus avant qu’il n’arrive en boîte de réception ?**

Oui, via l’intégration Milter avec Postfix décrite dans ce tutoriel. En réglant `OnInfected` sur `Reject`, le message est rejeté au niveau SMTP avant toute remise en boîte.

**Combien de temps prend le premier téléchargement de la base de signatures ?**

Entre 2 et 10 minutes selon la bande passante disponible, la base virale complète pesant plusieurs centaines de mégaoctets. Les mises à jour suivantes sont incrémentales et beaucoup plus rapides.

**ClamAV consomme-t-il beaucoup de ressources en fonctionnement normal ?**

Le pic de consommation mémoire survient au démarrage de clamd, lors du chargement de la base en mémoire. En fonctionnement courant, la consommation CPU reste proportionnelle au volume de fichiers scannés ; ajustez `MaxThreads` selon les ressources disponibles.

**Que faire si ClamAV signale un faux positif sur un fichier légitime ?**

Ne restaurez jamais le fichier sans vérification. Identifiez le nom de la signature déclenchée dans les logs, croisez-le avec une seconde source (comme VirusTotal), signalez le cas à l’équipe ClamAV via leur processus officiel de soumission, puis restaurez uniquement après confirmation.

**ClamAV suffit-il à lui seul pour la conformité NIS2 ?**

Non. ClamAV répond à l’exigence de détection de logiciels malveillants sur un périmètre donné (mail, fichiers), mais la conformité NIS2 couvre un ensemble de mesures bien plus large : gestion des accès, chiffrement, plan de continuité, notification d’incidents. ClamAV est une brique parmi d’autres dans une architecture de défense en profondeur.
