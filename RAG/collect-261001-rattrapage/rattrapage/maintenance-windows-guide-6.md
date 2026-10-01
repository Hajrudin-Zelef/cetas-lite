---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-6
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "arr", "incident", "memory"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [940, 1151]
sha256: 9c9c2cbec468962690446bd4eac76df541fce06434fec2623750142578846b67
---

# Maintenance et exploitation Windows en entreprise

> ⚠️ `/ResetBase` est **irréversible** : après son passage, impossible de
> désinstaller un correctif (§20). Ne l'utilisez qu'après validation complète du
> patch mensuel, jamais sur un serveur critique en période sensible.

Planificateur intégré : la tâche `Microsoft\Windows\Servicing\StartComponentCleanup`
tourne déjà automatiquement. Vérifiez qu'elle n'est pas désactivée (§47).

---

## 30. Journaux IIS et fichiers de trace : rotation

Les logs IIS (`C:\inetpub\logs\LogFiles\W3SVC*\`) grossissent vite sur un serveur
web exposé. Stratégie :

1. **Dans IIS** : limitez la rétention (journalisation > planifier : quotidien,
   taille max) ou centralisez vers un collecteur.
2. **Purge scriptée** des fichiers de plus de N jours :

```powershell
# Rotation des logs IIS : supprime les .log de +30 jours
$logRoot = 'C:\inetpub\logs\LogFiles'
Get-ChildItem $logRoot -Recurse -Filter '*.log' |
  Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } |
  Remove-Item -Force
```

3. **Autres traces gourmandes** : `C:\Windows\System32\LogFiles` (notamment
   `HTTPERR`), logs d'applications métier, dumps `MEMORY.DMP` oubliés après un
   BSOD (§70) — un seul fichier peut peser la taille de la RAM.

> Mettez les logs applicatifs **hors du volume système** dès l'installation
> (ex. `D:\Logs`). Un C: plein = serveur instable (§31).

---

## 31. Politique d'espace disque : seuils et alertes

Seuils d'exploitation recommandés :

| Volume | Alerte (warning) | Critique | Action |
|---|---|---|---|
| C: système | < 20 % libre | < 10 % libre ou < 10 Go | Nettoyage + analyse immédiate |
| Données | < 15 % libre | < 5 Go | Extension / archivage |
| Logs dédié | < 20 % libre | < 10 % libre | Rotation forcée (§30) |
| Sauvegarde | < 20 % libre | Job en échec | Purge des anciennes sauvegardes |

Surveillance native sans superviseur (quota d'alerte par volume via tâche
planifiée, §43) :

```powershell
# Alerte si un volume fixe passe sous 15 % libre (à planifier chaque heure)
$seuil = 15
Get-CimInstance Win32_LogicalDisk -Filter "DriveType=3" | ForEach-Object {
    $pct = [math]::Round($_.FreeSpace / $_.Size * 100, 1)
    if ($pct -lt $seuil) {
        $msg = "$($_.DeviceID) : $pct % libre sur $($env:COMPUTERNAME)"
        Write-EventLog -LogName Application -Source 'SuiviDisque' `
          -EventId 1001 -EntryType Warning -Message $msg
    }
}
# (Créer la source une fois : New-EventLog -LogName Application -Source 'SuiviDisque')
```

Top 10 des dossiers les plus lourds (diagnostic express) :

```powershell
Get-ChildItem C:\ -Directory -ErrorAction SilentlyContinue | ForEach-Object {
    $size = (Get-ChildItem $_.FullName -Recurse -File -ErrorAction SilentlyContinue |
             Measure-Object Length -Sum).Sum
    [pscustomobject]@{ Dossier = $_.FullName; Go = [math]::Round($size / 1GB, 2) }
} | Sort-Object Go -Descending | Select-Object -First 10
```

---

## 32. Event Viewer : prise en main

L'**Observateur d'événements** (`eventvwr.msc`) est le point d'entrée visuel.
Journaux Windows essentiels :

| Journal | Contenu |
|---|---|
| **Application** | Erreurs des applications et services |
| **Système** | Pilotes, services, matériel, démarrage/arrêt |
| **Sécurité** | Audits d'ouverture de session, accès objets (si audit activé) |
| **Installation** | Historique des installations (MSI, correctifs) |
| Journaux des applications et services | Par rôle : `Microsoft > Windows > ...` (ex. TerminalServices, SMBClient) |

Réflexes :

- Triez par **niveau** (Critique/Erreur) puis par **date** : les erreurs en grappe
  au même moment = un incident unique, pas dix.
- L'onglet **Détails > convivial** donne les champs exploitables (noms, codes).
- Un ID d'événement + sa **source** suffisent à une recherche web ciblée
  (ex. « Event ID 55 Ntfs »).

---

## 33. Get-WinEvent : interroger les journaux en PowerShell

`Get-WinEvent` remplace l'ancien `Get-EventLog` (limité aux journaux classiques).
C'est l'outil n°1 d'exploitation scriptée.

```powershell
# 50 dernières erreurs du journal Système
Get-WinEvent -LogName System -MaxEvents 50 | Where-Object LevelDisplayName -eq 'Error'

# Filtre efficace côté fournisseur (beaucoup plus rapide) : dernières 24 h, niveaux 1-2
$filtre = @{ LogName = 'System'; Level = 1, 2; StartTime = (Get-Date).AddDays(-1) }
Get-WinEvent -FilterHashtable $filtre | Select-Object TimeCreated, Id, ProviderName, Message |
  Format-Table -AutoSize -Wrap

# Compter les erreurs par ID sur 7 jours (détection de patterns)
Get-WinEvent -FilterHashtable @{ LogName = 'System'; Level = 2; StartTime = (Get-Date).AddDays(-7) } |
  Group-Object Id | Sort-Object Count -Descending | Select-Object -First 15 Count, Name

# Journal Sécurité : échecs d'ouverture de session (4625) sur 24 h
Get-WinEvent -FilterHashtable @{ LogName = 'Security'; Id = 4625; StartTime = (Get-Date).AddDays(-1) } |
  Select-Object TimeCreated,
    @{n='Compte'; e={ $_.Properties[5].Value }},
    @{n='IP';     e={ $_.Properties[19].Value }} |
  Format-Table -AutoSize
```

> ⚠️ L'index des `Properties[]` varie selon la version : vérifiez une fois sur un
> événement réel (onglet Détails) avant d'industrialiser.

Interroger une machine distante (PSRemoting ou `-ComputerName` sur les journaux
classiques uniquement ; préférez `Invoke-Command`, §113).

---

## 34. wevtutil : l'outil en ligne de commande

`wevtutil` est le pendant CMD de Get-WinEvent, utile en WinRE ou en script batch.

```cmd
:: Lister les journaux disponibles
wevtutil el

:: 10 derniers événements du journal Système, texte lisible, plus récents d'abord
wevtutil qe System /c:10 /rd:true /f:text

:: Exporter un journal en .evtx (pour analyse hors ligne)
wevtutil epl System C:\Temp\system.evtx

:: Vider un journal (après export !)
wevtutil cl System

:: Taille et rétention d'un journal
wevtutil gl System
```

Effacer **tous** les journaux d'un poste en dépannage (radical, à tracer) :

```cmd
for /f "tokens=*" %l in ('wevtutil el') do wevtutil cl "%l"
```

---

## 35. Filtres XPath et vues personnalisées

Dans l'Observateur : clic droit sur un journal > **Filtrer le journal actuel** >
onglet **XML** > cocher « Modifier la requête manuellement ». Exemple : toutes les
erreurs disque et NTFS sur 7 jours :

```xml
<QueryList>
  <Query Id="0" Path="System">
    <Select Path="System">
      *[System[(Level=2 or Level=3)
        and (EventID=7 or EventID=11 or EventID=55 or EventID=98 or EventID=140 or EventID=153)
        and TimeCreated[timediff(@SystemTime) &lt;= 604800000]]]
    </Select>
  </Query>
</QueryList>
```

**Vues personnalisées** : enregistrez vos filtres usuels (clic droit > Créer une vue
personnalisée) : « Erreurs disque 7 j », « Échecs logon 24 h », « Redémarrages
inattendus ». Exportables en XML pour déploiement homogène sur le parc
(import via la console ou GPO de préférences).

En PowerShell, le même filtre s'exprime via `-FilterXml` :

```powershell
$xml = @'
<QueryList><Query Id="0" Path="System"><Select Path="System">
*[System[(Level=2) and TimeCreated[timediff(@SystemTime) &lt;= 86400000]]]
</Select></Query></QueryList>
'@
Get-WinEvent -FilterXml $xml | Select-Object TimeCreated, Id, ProviderName
```

---

## 36. Abonnements d'événements (WEF) et collecteur

Le **transfert d'événements Windows (WEF)** centralise les journaux vers un
**collecteur** sans agent tiers : les sources « poussent » (push) leurs événements.

**Côté collecteur** (un serveur dédié ou le serveur de supervision) :

```cmd
:: Initialiser le service Windows Event Collector
wecutil qc /q
```

Puis créer l'abonnement (console `eventvwr.msc` > Abonnements, ou `wecutil cs`
avec un XML). Exemple minimal d'abonnement « Erreurs Système de tous les serveurs » :

