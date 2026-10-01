---
id: collect-261001-general-networking/general-networking/clamav-antivirus-open-source-linux-10-etapes-2026-4
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2025-09-14"]
keywords: ["distribution", "incident", "open source"]
source: docs/RAG/collect-261001-general-networking/clamav-antivirus-open-source-linux-10-etapes-2026.md
source_anchor: ""
source_lines: [317, 371]
sha256: 0d2cbaa6ffb030ab163cea6403de2a5811b4ab3b947aa041b7e140f823a9d437
---

# Debian / Ubuntu

| Branche | Type | Date de sortie | Fin de support sécurité | Dernier patch connu | 
|---|---|---|---|---|
| 0.103 | LTS | 14 septembre 2020 | 14 septembre 2024 (signatures jusqu’au 14/09/2025) | Fin de vie atteinte | 
| 1.0 | LTS | 28 novembre 2022 | 28 novembre 2025 | 1.0.9 (18 juin 2025) | 
| 1.2 | Non-LTS | 27 août 2023 | 15 décembre 2024 | Fin de vie atteinte | 
| 1.3 | Non-LTS | 7 février 2024 | 7 février 2026 | 1.3.2 (4 septembre 2024) | 
| 1.4 | LTS | 15 août 2024 | 15 août 2027 | 1.4.6 (7 août 2026) | 
| 1.5 | Stable / latest | 2025 | En cours | 1.5.4 (7 août 2026) | 

Pour un déploiement en production, deux stratégies raisonnables se dégagent de ce tableau. Si la stabilité et un cycle de mises à jour prévisible priment, choisissez la branche 1.4 LTS, supportée jusqu’en août 2027. Si vous préférez bénéficier des dernières fonctionnalités et de la réduction de taille des signatures amorcée fin 2025, la branche 1.5 (taguée `latest` et `stable` par l’équipe ClamAV) est le choix par défaut. Dans les deux cas, planifiez dès maintenant la migration hors des branches 0.103, 1.0 et 1.2 si elles tournent encore chez vous : elles ne reçoivent plus de correctifs de sécurité.

## Erreurs courantes à éviter avec ClamAV

Cinq erreurs reviennent systématiquement dans les déploiements ClamAV mal préparés, souvent découvertes trop tard, au moment d’un incident.

- **Oublier NotifyClamd dans freshclam.conf** — sans cette directive, clamd continue de tourner avec une base de signatures obsolète après chaque mise à jour de freshclam, jusqu’au prochain redémarrage manuel du démon.
- **Scanner uniquement à l’arrivée, jamais en continu** — un scan ponctuel au moment de l’upload ne protège pas contre un fichier déposé avant l’activation de ClamAV, ou contre une signature ajoutée après coup pour un malware déjà présent sur le disque. Un scan planifié régulier reste indispensable en complément.
- **Confondre clamscan et clamdscan sur de gros volumes** — utiliser clamscan pour un scan récurrent sur des dizaines de milliers de fichiers recharge la base à chaque exécution et peut multiplier le temps de traitement par dix par rapport à clamdscan, qui délègue au démon déjà en mémoire.
- **Négliger la mise à jour du système hôte** — un ClamAV à jour sur un OS non patché reste vulnérable aux failles du noyau ou des bibliothèques système que le malware analysé pourrait exploiter une fois exécuté, indépendamment de la détection antivirus.
- **Ignorer les dates de fin de vie des branches** — faire tourner une version 0.103 ou 1.2 sans plan de migration expose à des vulnérabilités connues et déjà publiées, comme celles corrigées par les patchs 1.4.6 et 1.5.4 d’août 2026.

## Dépannage : les problèmes ClamAV les plus fréquents

Voici les incidents les plus fréquemment rencontrés en production, avec leur cause probable et la marche à suivre.

- **freshclam échoue avec « Can’t connect to port 443 »** — un pare-feu sortant bloque l’accès HTTPS vers les miroirs de signatures. Vérifiez avec`curl -v https://database.clamav.net` et ouvrez le port 443 en sortant si nécessaire, ou déclarez un proxy via`HTTPProxyServer` dans freshclam.conf.
- **clamd refuse de démarrer avec « ERROR: Can’t open/parse the config file »** — une erreur de syntaxe s’est glissée dans clamd.conf, souvent une ligne dupliquée ou un paramètre mal orthographié. Validez avec`clamd --config-file=/etc/clamav/clamd.conf --debug` qui affiche la ligne fautive.
- **Consommation mémoire excessive au démarrage** — normal sur une machine à 2 Go de RAM au chargement initial de la base ; si le service est tué par l’OOM killer, augmentez la RAM disponible ou limitez`MaxThreads` pour réduire l’empreinte mémoire globale.
- **Le milter ne bloque aucun mail malgré une détection dans les logs** — vérifiez que`OnInfected` est bien réglé sur`Reject` ou`Quarantine` et non`Pass` , et que Postfix pointe vers le bon port dans`smtpd_milters` .
- **clamonacc ne détecte aucun fichier créé** — le noyau doit exposer fanotify ; vérifiez avec`zgrep CONFIG_FANOTIFY /proc/config.gz` ou consultez la documentation de votre distribution pour l’activer.
- **Faux positif récurrent sur un type de fichier interne** — créez une exception locale via un fichier de signature d’exclusion personnalisé (`local.ign2` ) plutôt que de désactiver globalement une catégorie de détection.
- **Scan anormalement lent sur des archives volumineuses** — vérifiez les paramètres`MaxScanSize` ,`MaxFileSize` et`ArchiveBlockEncrypted` , qui peuvent forcer une analyse exhaustive d’archives imbriquées consommant beaucoup de CPU.
- **Erreur « ERROR: Malformed database » après une mise à jour** — le téléchargement d’une base de signatures a été interrompu ou corrompu. Supprimez les fichiers`.cvd` dans`/var/lib/clamav` et relancez`freshclam` pour forcer un téléchargement complet.
- **Socket clamd introuvable (« Could not connect to clamd »)** — vérifiez que le service clamav-daemon est bien actif avec`systemctl status clamav-daemon` , et que le chemin du socket dans clamdscan.conf correspond bien à celui déclaré dans clamd.conf.

## ClamAV face aux solutions commerciales : Sophos, Bitdefender GravityZone, ESET

ClamAV ne participe pas régulièrement aux tests indépendants publiés par AV-Test ou AV-Comparatives, contrairement aux suites commerciales grand public. Il n’existe donc pas de score de détection comparable et vérifiable à opposer directement à Sophos, Bitdefender GravityZone ou ESET. La comparaison pertinente n’est pas tant « qui détecte le plus de menaces » que « quel outil pour quel usage ».

| Critère | ClamAV | Solutions commerciales (Sophos, Bitdefender GravityZone, ESET) | 
|---|---|---|
| Coût de licence | Gratuit, open source | Abonnement par poste/serveur | 
| Cas d’usage principal | Serveurs, passerelles mail, scan de fichiers | Postes de travail, serveurs, EDR intégré | 
| Console centralisée | Aucune nativement, à construire soi-même | Console cloud ou on-premise incluse | 
| Sandboxing / ML avancé | Non | Souvent inclus dans les offres premium | 
| Tests indépendants AV-Test/AV-Comparatives | Non régulièrement évalué | Évalué mensuellement | 
| Support éditeur | Communautaire + Cisco Talos (maintenance) | Support contractuel SLA | 
| Intégration Postfix/Milter native | Oui, clamav-milter dédié | Variable selon l’éditeur | 

En pratique, beaucoup d’organisations combinent les deux approches : ClamAV en première ligne sur la passerelle mail pour filtrer le volume brut à coût nul, et une solution commerciale avec EDR sur les postes de travail pour la détection comportementale avancée et la réponse aux incidents. Ce n’est pas l’un ou l’autre, mais une architecture en couches où chaque outil couvre un périmètre différent.

## ClamAV et conformité NIS2 pour les PME françaises

La directive NIS2 impose aux entités essentielles et importantes des mesures techniques de détection et de prévention des logiciels malveillants, y compris sur les serveurs de messagerie et de partage de fichiers. ClamAV, déployé sur une passerelle mail Postfix comme décrit dans ce tutoriel, contribue directement à cette exigence, à condition d’être intégré dans une démarche documentée plutôt que laissé en configuration par défaut.

