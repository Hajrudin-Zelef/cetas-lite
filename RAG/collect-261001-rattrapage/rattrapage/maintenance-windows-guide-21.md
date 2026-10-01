---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-21
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "arr", "incident"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [3916, 4046]
sha256: 0346c7c3464e1c544b67a5a8fb4911ec56f7a3ec908a5d7841a204327a3060f1
---

# Maintenance et exploitation Windows en entreprise

1. Votre serveur affiche « uptime 412 jours ». Pourquoi est-ce un problème
   d'exploitation, et que faites-vous ?
2. Un correctif du Patch Tuesday fait planter l'application métier sur l'anneau
   pilote. Décrivez la conduite à tenir pour l'anneau production.
3. `sfc /scannow` rapporte des fichiers irréparables. Quelle est la séquence
   correcte avant de le relancer ?
4. Le journal Sécurité d'un DC s'écrase tous les 2 jours. Quel est le risque et
   les deux corrections possibles ?
5. Un poste affiche `169.254.12.34`. Que signifie cette adresse et quelles sont
   les 3 vérifications dans l'ordre ?
6. `Test-NetConnection srv -Port 443` échoue mais le ping passe. Listez les 3
   causes les plus probables dans l'ordre de vérification.
7. Un utilisateur signale « la relation d'approbation a échoué » à l'ouverture de
   session. Donnez les deux méthodes de réparation, de la moins intrusive à la
   plus intrusive.
8. Pourquoi `Get-CimInstance Win32_Product` est-il à proscrire en inventaire ?
   Que faut-il utiliser à la place ?
9. Définissez RTO et RPO, et expliquez qui les valide dans l'entreprise.
10. Citez 3 raisons pour lesquelles une tâche planifiée peut échouer avec une
    erreur d'ouverture de session, et la solution recommandée.

---

## 136. Réponses du quiz

1. **Uptime 412 jours** = correctifs jamais finalisés (redémarrages en attente),
   risque de fuite mémoire et de panne au prochain reboot « surprise ». Action :
   planifier un redémarrage en fenêtre de maintenance (§8), vérifier la conformité
   patch après reboot.
2. **Bloquer le KB** pour la production (refus WSUS / pause WUfB), documenter
   (KB, symptômes, contournement), informer l'éditeur, re-tester au patch suivant
   (§11, §20). Ne jamais « forcer » un KB qui casse le métier.
3. **DISM d'abord** : `DISM /Online /Cleanup-Image /RestoreHealth` répare le
   magasin de composants dans lequel SFC puise, **puis** relancer `sfc /scannow`
   (§21-24).
4. **Risque** : aucune capacité d'enquête a posteriori (incident, audit).
   Corrections : augmenter la taille max (`wevtutil sl Security /ms:…`, §41)
   et/ou transférer vers un collecteur/SIEM (WEF §36, §110).
5. **APIPA** = aucun serveur DHCP n'a répondu. Vérifications : (1) connectivité
   locale (câble, `ping` passerelle), (2) `ipconfig /renew` + journal DHCP
   client, (3) côté infra : service DHCP, étendue saturée, relais IP helper du
   VLAN (§86, §83, §40).
6. **Port 443 fermé, ping OK** : (1) le service n'écoute pas / est arrêté,
   (2) pare-feu Windows, (3) pare-feu réseau inter-VLAN (§93). Vérifier l'écoute
   avec `Get-NetTCPConnection -State Listen` sur le serveur.
7. (1) `Test-ComputerSecureChannel -Repair` (répare le canal sans quitter le
   domaine), (2) quitter puis réintégrer le domaine (`Remove-Computer` /
   `Add-Computer`) si la réparation échoue (§103).
8. `Win32_Product` **déclenche une reconfiguration de chaque MSI** (lent, peut
   casser des applications). Utiliser le registre `Uninstall` ou `Get-Package`
   (§116).
9. **RTO** = durée max d'interruption acceptable ; **RPO** = perte de données max
   acceptable. Ils sont **négociés et signés avec la direction et les métiers**,
   pas décidés par l'IT seule (§122).
10. (1) mot de passe du compte expiré/modifié, (2) compte sans le droit « Ouvrir
    une session en tant que tâche », (3) compte désactivé. **Solution** : migrer
    vers un **gMSA** (mot de passe géré par AD) ou SYSTEM selon le besoin (§45-46).

---

## 137. Glossaire

| Terme | Définition |
|---|---|
| **AD DS** | Active Directory Domain Services — annuaire d'entreprise |
| **APIPA** | Adressage automatique 169.254.x.x quand aucun DHCP ne répond |
| **BCD** | Boot Configuration Data — magasin de démarrage (remplace boot.ini) |
| **BSOD** | Blue Screen Of Death — écran bleu d'erreur fatale du noyau |
| **DISM** | Outil de maintenance des images et du magasin de composants |
| **DFSR** | DFS Replication — réplication SYSVOL entre DC |
| **DSRM** | Directory Services Restore Mode — mode de restauration AD |
| **EDR** | Endpoint Detection and Response — antivirus nouvelle génération |
| **FSRM** | File Server Resource Manager — quotas/rapports avancés |
| **FSMO** | Rôles maîtres d'opérations AD (PDC, RID, Infrastructure, Schéma, Nom de domaine) |
| **gMSA** | Group Managed Service Account — compte de service à mot de passe géré par AD |
| **GPO** | Group Policy Object — stratégie de groupe |
| **NLA** | Network Location Awareness — détection du profil réseau |
| **OOB** | Out-Of-Band — correctif publié hors cycle mensuel |
| **PAL** | Performance Analysis of Logs — analyse des journaux de compteurs |
| **PCA** | Plan de Continuité d'Activité |
| **PRA** | Plan de Reprise d'Activité |
| **PSRemoting** | Exécution PowerShell à distance via WinRM |
| **ReFS** | Resilient File System — système de fichiers résilient |
| **RPO** | Recovery Point Objective — perte de données max acceptable |
| **RSOP** | Resultant Set Of Policy — résultante des stratégies appliquées |
| **RTO** | Recovery Time Objective — durée max d'interruption acceptable |
| **SFC** | System File Checker — vérificateur des fichiers système |
| **SMART** | Self-Monitoring, Analysis and Reporting Technology — santé des disques |
| **SSU** | Servicing Stack Update — MAJ de la pile de maintenance Windows Update |
| **SYSVOL** | Partage répliqué contenant GPO et scripts d'ouverture de session |
| **VSS** | Volume Shadow Copy Service — clichés instantanés |
| **WEF** | Windows Event Forwarding — transfert d'événements vers un collecteur |
| **WinRE** | Windows Recovery Environment — environnement de récupération |
| **WSUS** | Windows Server Update Services — serveur local de mises à jour (déprécié) |
| **WUfB** | Windows Update for Business — pilotage des MAJ sans serveur local |

---

## 138. Pour aller plus loin

**Documentation officielle (à consulter en priorité) :**

- Microsoft Learn : *Windows Server* (déploiement, rôles, dépannage) —
  `learn.microsoft.com/windows-server`
- Référence des ID d'événements Windows et des compteurs de performances —
  Microsoft Learn
- Documentation **windows_exporter** (collecteurs, textfile) —
  dépôt GitHub `prometheus-community/windows_exporter`
- Documentation **Zabbix** : template *Windows by Zabbix agent*

**Outils :**

- **Sysinternals Suite** (Process Explorer, Autoruns, Sysmon, Procmon) —
  indispensables en dépannage avancé
- **WinDbg** (Microsoft Store) + symboles Microsoft — analyse de dumps (§71)
- **PAL** (Performance Analysis of Logs) — analyse des journaux perfmon (§60)
- **Wireshark** — capture réseau quand `tracert`/`Test-NetConnection` ne suffisent plus

**Sujets connexes à ce guide :**

- Supervision : *guide Zabbix* et *guide Prometheus/Grafana* (métriques, alerting)
- Sauvegarde : solutions d'entreprise (Veeam…), PRA/PCA — §122-127
- Sécurité : durcissement AD, LAPS, EDR — guides dédiés
- Automatisation : PowerShell avancé, DSC/Intune — §111-121
- Énergie : onduleurs, dimensionnement, arrêt automatique (NUT) — en lien avec
  l'exploitation des salles serveurs

**Pratique recommandée** : montez un **labo** (Hyper-V ou Proxmox) avec 1 DC +
2 serveurs membres + 1 client Windows 11, et rejouez chaque section « cas
pratique » de ce guide. C'est en cassant puis réparant qu'on devient
opérationnel — jamais en lisant seulement.
