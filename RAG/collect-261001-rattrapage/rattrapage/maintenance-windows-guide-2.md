---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-2
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2025-11-11", "2026-10-13", "2029-01-09", "2031-10-14", "2034-10-10"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [182, 348]
sha256: 63738a4283ce7a5633ca4e96bbe0f4138d0daf439eaadb237ee648b512e96a6c
---

# Maintenance et exploitation Windows en entreprise

| Plateforme | Versions couvertes | Fin de support (rappel) |
|---|---|---|
| Windows Server | 2019 (1809), 2022 (21H2), 2025 (24H2) | 2019 : 09/01/2029 (étendu) ; 2022 : 14/10/2031 ; 2025 : 10/10/2034 |
| Windows 11 | 23H2 (22631), 24H2 (26100) | 23H2 : 11/11/2025 (Pro) ; 24H2 : 13/10/2026 (Pro) — vérifiez le cycle en cours |

**Conventions du guide :**

- `PS>` = invite PowerShell (exécuté en administrateur sauf mention contraire).
- `CMD>` = invite de commandes classique.
- Les cmdlets sont données pour **PowerShell 5.1** (inclus dans Windows) sauf mention
  de PowerShell 7.
- Les chemins de GPO sont donnés sous la forme
  `Configuration ordinateur > Stratégies > Modèles d'administration > ...`.
- ⚠️ signale une opération à risque ou dépendante de la version.

> ⚠️ **Windows Server 2025** : WSUS y est toujours présent mais Microsoft a annoncé en
> 2024 sa **dépréciation** (plus aucune fonctionnalité nouvelle). Pour un nouveau
> déploiement, privilégiez **Windows Update for Business** (§16) ou une solution
> tierce, et ne basez plus votre stratégie à 5 ans sur WSUS seul.

---

## 3. Inventaire : la base de toute exploitation

Pas d'inventaire = pas de maintenance sérieuse. L'inventaire minimal par machine :

| Champ | Exemple | Source |
|---|---|---|
| Nom / rôle | `srv-fic-01` / serveur de fichiers | CMDB |
| OS + build | Server 2022 21H2 (20348.2402) | `Get-ComputerInfo` |
| IP / VLAN | 10.10.20.11 / VLAN 20 | IPAM |
| Garantie / contrat | Dell ProSupport jusqu'au 2028-06 | Fournisseur |
| Criticité | P1 (production) | Analyse d'impact |
| Fenêtre de maintenance | Dim 02:00–05:00 | Planning |
| Responsable | Équipe infra | Annuaire |

Collecte automatisée (détail §115-117). Stockez le résultat dans une CMDB, un tableur
versionné, ou à défaut un partage avec historique. **Règle** : tout serveur ajouté au
domaine est inventorié le jour même ; revue complète trimestrielle.

Script express d'inventaire (à lancer via PSRemoting, §113) :

```powershell
$servers = 'srv-fic-01','srv-ad-01','srv-prn-01'
Invoke-Command -ComputerName $servers {
    [pscustomobject]@{
        Machine   = $env:COMPUTERNAME
        OS        = (Get-ComputerInfo).WindowsProductName
        Build     = (Get-ComputerInfo).OsBuildNumber
        IP        = (Get-NetIPAddress -AddressFamily IPv4 |
                     Where-Object { $_.InterfaceAlias -notlike '*Loopback*' }).IPAddress -join ', '
        DisqueC_Go = [math]::Round((Get-PSDrive C).Free / 1GB, 1)
        MAJ       = (Get-HotFix | Sort-Object InstalledOn -Descending |
                     Select-Object -First 1).HotFixID
    }
} | Format-Table -AutoSize
```

---

## 4. Plan de maintenance : le calendrier annuel

Le plan de maintenance est **le** document du chef de service. Il fixe qui fait quoi,
quand, avec quelle fenêtre d'intervention. Modèle type :

| Périodicité | Activités | Responsable | Fenêtre |
|---|---|---|---|
| Quotidienne | Checklists §5-6, supervision | Astreinte | 08:30 |
| Hebdomadaire | Checklist §7, revue des alertes | Technicien | Ven 16:00 |
| Mensuelle | Patch Tuesday (§10-11), checklist §8 | Équipe infra | Sam/Dim |
| Trimestrielle | Checklist §9, revue inventaire, test PRA partiel | Chef de service | Planifié |
| Annuelle | Test PRA complet, revue contrats, audit AD | Chef de service | Planifié |

**Principes :**

1. **Une fenêtre de maintenance contractuelle** par classe de serveurs (ex : serveurs
   applicatifs le 2ᵉ dimanche du mois 02:00–06:00). Communiquez-la aux métiers.
2. **Gel des changements** pendant les périodes critiques métier (clôtures, paies,
   inventaires).
3. **Toute intervention hors fenêtre** = changement exceptionnel (§130) avec validation.
4. **Journalisez tout** : un ticket ou une entrée de journal par intervention, même
   de 5 minutes. La mémoire est un mauvais outil d'exploitation.

---

## 5. Checklist quotidienne (serveurs)

À exécuter chaque matin ouvré, idéalement en partie automatisée (rapport e-mail).

- [ ] **Supervision** : 0 alerte critique non acquittée (Zabbix/Prometheus, §106).
- [ ] **Espace disque** : aucun volume < 15 % libre sur les serveurs P1/P2.
- [ ] **Sauvegardes** : tous les jobs de la nuit en succès (vérifier le rapport, pas
      seulement l'absence d'alerte).
- [ ] **Services critiques** : AD DS, DNS, DHCP démarrés sur les contrôleurs ;
      spouleur, partages sur les serveurs concernés.
- [ ] **Journal Système** : erreurs/disques (`disk`, `ntfs`), erreurs de démarrage.
- [ ] **Journal Sécurité** : pics d'échecs d'ouverture de session (ID 4625).
- [ ] **Réplication AD** : `repadmin /replsummary` sans erreur > 0.
- [ ] **Antivirus/EDR** : définitions à jour, 0 menace non traitée.
- [ ] **Onduleurs** : état normal si supervisés (lien avec l'exploitation énergie).

Script de contrôle matinal (extrait ; version complète §118) :

```powershell
# Contrôle express : disque + services critiques sur une liste de serveurs
$servers = Get-Content C:\Scripts\serveurs.txt
Invoke-Command -ComputerName $servers {
    $disk = Get-CimInstance Win32_LogicalDisk -Filter "DeviceID='C:'"
    [pscustomobject]@{
        Machine   = $env:COMPUTERNAME
        C_LibrePct = [math]::Round($disk.FreeSpace / $disk.Size * 100, 1)
        DNS       = (Get-Service Dnscache).Status
        Spooler   = (Get-Service Spooler).Status
    }
} | Where-Object { $_.C_LibrePct -lt 20 } | Format-Table -AutoSize
```

---

## 6. Checklist quotidienne (postes)

- [ ] **File de tickets** : nouveaux incidents utilisateurs triés avant 09:30.
- [ ] **Déploiements en échec** : rapports Intune/GPO/SCCM (logiciels, mises à jour).
- [ ] **Postes critiques** (accueil, ateliers, caisses) : démarrés et à jour.
- [ ] **Antivirus** : postes sans remontée > 48 h = à investiguer.
- [ ] **Espace disque** : alerte sous 10 % sur C: (profils utilisateurs).

> Astuce : un rapport PowerShell quotidien envoyé par e-mail (§121) remplace
> avantageusement la vérification manuelle poste par poste.

---

## 7. Checklist hebdomadaire

- [ ] Revue des **alertes** de la semaine : patterns récurrents ? actions correctives ?
- [ ] **Journaux d'événements** des serveurs critiques : erreurs nouvelles vs baseline.
- [ ] **Espace disque** : tendance (un volume qui perd 2 %/semaine sera plein dans X
      semaines — anticipez, §31).
- [ ] **Correctifs en attente** : serveurs/postes non conformes au patch du mois.
- [ ] **Tâches planifiées** : 0 échec sur les tâches critiques (§47).
- [ ] **Sauvegardes** : un contrôle de restauration de fichier test par semaine
      (tourniquet sur les serveurs).
- [ ] **Sécurité** : comptes verrouillés, nouveaux comptes créés, groupes sensibles
      modifiés (journal Sécurité des DC).
- [ ] **Nettoyage** : corbeilles, dossiers temp des serveurs TSE/RDS.

---

## 8. Checklist mensuelle

- [ ] **Patch Tuesday** : déploiement anneaux test → pilote → production (§11).
- [ ] **Rapport de conformité** : % de machines à jour par anneau (objectif ≥ 98 %).
- [ ] **Redémarrages** : tous les serveurs redémarrés au moins une fois dans le mois
      (les serveurs « uptime 400 jours » sont des bombes à retardement).
- [ ] **Comptes AD** : désactiver les comptes inactifs > 90 jours (§119).
- [ ] **Certificats** : liste des certificats expirant dans < 60 jours (§120).
- [ ] **Quotas et partages** : revue des dossiers > 80 % de quota.
- [ ] **GPO** : sauvegarde (Backup-GPO) avant toute modification.
- [ ] **Documentation** : procédures mises à jour après chaque changement notable.
- [ ] **Revue des changements** du mois avec l'équipe (15 min suffisent).

---

## 9. Checklist trimestrielle et annuelle

**Trimestrielle :**

