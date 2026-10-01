---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-17
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [3133, 3344]
sha256: 26163a02dcf3fd6662397a5fb68e0dc2d4b6ad1acd23de9be6ec0040980e7eff
---

# Maintenance et exploitation Windows en entreprise

**Services** : utilisez la découverte `service.discovery` + prototypes pour
surveiller automatiquement les services critiques par rôle (liste par host group :
`Windows-DC`, `Windows-FileServer`…).

---

## 109. Prometheus : windows_exporter

**windows_exporter** expose les métriques Windows au format Prometheus
(port **9182** par défaut). Collecteurs utiles à activer :

```cmd
:: Installation en service (exemple)
windows_exporter.exe --collectors.enabled="cpu,cs,logical_disk,mem,net,os,service,system,textfile"
```

| Collecteur | Métriques |
|---|---|
| `cpu`, `cs`, `os`, `system` | Charge, mémoire, uptime |
| `logical_disk` | Espace libre/utilisé par volume |
| `net` | Débit/erreurs interfaces |
| `service` | État des services Windows |
| `textfile` | **Vos métriques maison** : déposez des `.prom` générés par script |

Exemple `textfile` : le script §106 (reboot en attente) écrit
`C:\Program Files\windows_exporter\textfile_inputs\maintenance.prom` :

```text
# HELP windows_pending_reboot Redémarrage en attente (1=oui)
# TYPE windows_pending_reboot gauge
windows_pending_reboot 0
```

```powershell
# Génération du fichier .prom par tâche planifiée
$pending = if (Test-Path 'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Component Based Servicing\RebootPending') { 1 } else { 0 }
@'
# HELP windows_pending_reboot Redémarrage en attente (1=oui)
# TYPE windows_pending_reboot gauge
windows_pending_reboot {0}
'@ -f $pending | Out-File 'C:\Program Files\windows_exporter\textfile_inputs\maintenance.prom' -Encoding ascii
```

Alertes Prometheus type : `windows_logical_disk_free_bytes / windows_logical_disk_size_bytes < 0.15`,
`windows_service_state{state="stopped"} == 1` sur services critiques.

---

## 110. Centralisation des journaux (syslog / Graylog / Loki)

Options pour centraliser sans WEF (§36) ou en complément :

| Solution | Protocole | Remarque |
|---|---|---|
| **WEF natif** | WinRM/HTTPS | Sans agent, parfait en domaine |
| **NXLog / Winlogbeat** | Syslog / Beats | Vers Graylog, ELK, Loki |
| **Agent Zabbix** | Zabbix | `eventlog[]` pour les journaux |
| **Promtail + Loki** | HTTP | Couplé à windows_exporter/Grafana |

Exemple Winlogbeat minimal (vers Logstash/Elasticsearch) :

```yaml
winlogbeat.event_logs:
  - name: System
    level: critical, error, warning
  - name: Security
    event_id: 4625, 4740, 4720, 4728, 4732, 4756, 1102
  - name: Microsoft-Windows-Sysmon/Operational
output.logstash:
  hosts: ["logstash.contoso.local:5044"]
```

**Sysmon** (Sysinternals) mérite une mention : il enrichit énormément le journal
`Microsoft-Windows-Sysmon/Operational` (créations de processus avec ligne de
commande, connexions réseau par processus) — précieux en sécurité, à déployer
avec une configuration de référence (ex. SwiftOnSecurity) adaptée.

---

## 111. PowerShell : fondamentaux d'administration

```powershell
# Aide et découverte (vos deux meilleurs amis)
Get-Help Get-Service -Full
Get-Command -Module NetAdapter
Get-Command *disk*

# Pipeline : filtrer, trier, sélectionner
Get-Service | Where-Object Status -eq 'Running' | Sort-Object DisplayName |
  Select-Object Name, DisplayName, StartType | Format-Table -AutoSize

# CIM plutôt que WMI (moderne, compatible PSRemoting)
Get-CimInstance Win32_OperatingSystem | Select-Object Caption, Version, LastBootUpTime

# Exporter proprement
Get-Service | Export-Csv C:\Temp\services.csv -NoTypeInformation -Encoding UTF8
Get-Service | ConvertTo-Json | Out-File C:\Temp\services.json
```

**Politique d'exécution** (pour vos scripts signés ou internes) :

```powershell
Set-ExecutionPolicy RemoteSigned -Scope LocalMachine -Force
# RemoteSigned : scripts locaux OK, scripts téléchargés = signature requise
```

> Écrivez vos scripts en **fonctions avancées** (`[CmdletBinding()]`,
> paramètres typés, `-WhatIf` supporté) dès qu'ils dépassent 20 lignes ou qu'ils
> modifient le système. Un script « one-shot » devient toujours un script de
> production un jour.

---

## 112. PSRemoting et WinRM : mise en place sécurisée

PSRemoting = exécuter du PowerShell **à distance** via WinRM. Activation :

```powershell
# Sur chaque machine cible (ou par GPO : voir ci-dessous)
Enable-PSRemoting -Force
# Vérifier
Test-WSMan srv-fic-01
```

Par GPO (recommandé en domaine) :

```text
Configuration ordinateur > Stratégies > Modèles d'administration >
Composants Windows > WinRM > Service WinRM
  - Autoriser la gestion de serveur à distance via WinRM : Activé (IPv4: *)
+ Pare-feu Windows : autoriser « Gestion à distance Windows » (TCP 5985)
+ Service WinRM en démarrage automatique
```

**Sécurité :**

- En domaine avec Kerberos, le trafic WinRM est **chiffré** par défaut :
  n'activez HTTPS/5986 que si vous traversez des zones non maîtrisées.
- **N'utilisez pas TrustedHosts avec `*`** en production : c'est l'équivalent de
  désactiver l'authentification mutuelle.
- Restreignez qui peut se connecter : groupe « Utilisateurs de la gestion à
  distance » (Remote Management Users), jamais les comptes utilisateurs standard.
- Désactivez l'authentification de base (`AllowBasic = false`) côté serveur.

```powershell
# Durcir : vérifier la config WinRM
winrm get winrm/config/service
# Doit montrer : AllowBasic=false (sauf besoin explicite), AllowUnencryptedTraffic=false
```

---

## 113. Invoke-Command : exécution à distance

```powershell
# Une commande sur N machines (Kerberos implicite en domaine)
Invoke-Command -ComputerName srv-fic-01, srv-prn-01 {
    Get-Service Spooler | Select-Object MachineName, Status
}   # Note : $env:COMPUTERNAME dans le bloc distant = la cible

# Avec identifiants explicites (hors domaine ou compte dédié)
$cred = Get-Credential -UserName 'contoso\admin-expl' -Message 'Compte exploitation'
Invoke-Command -ComputerName srv-dmz-01 -Credential $cred -ScriptBlock {
    Get-HotFix | Sort-Object InstalledOn -Descending | Select-Object -First 3
}

# Exécuter un script local SUR les machines distantes
Invoke-Command -ComputerName (Get-Content C:\Scripts\serveurs.txt) `
  -FilePath C:\Scripts\Collecte-Sante.ps1

# Récupérer des fichiers distants (rapports générés à distance)
Invoke-Command -ComputerName srv-fic-01 {
    Get-ChildItem C:\Temp\rapport-*.csv | Select-Object -ExpandProperty FullName
} | ForEach-Object {
    Copy-Item -Path $_ -Destination C:\Temp\Central\ -FromSession $s
}
```

> Les objets retournés sont **désérialisés** : les méthodes ne fonctionnent plus
> (ex. `Stop()` sur un service distant). Faites tout le traitement **dans** le
> bloc distant, ne retournez que des données.

---

## 114. Sessions persistantes et fichiers distants

```powershell
# Session réutilisable (plus rapide que N connexions)
$s = New-PSSession -ComputerName srv-fic-01
Invoke-Command -Session $s { Get-PSDrive C }
Invoke-Command -Session $s { Get-Service | Where-Object Status -eq 'Stopped' }
Remove-PSSession $s

# Copier des fichiers vers/depuis la machine distante (PS 5.1+)
Copy-Item -Path C:\Scripts\Deploy.ps1 -Destination C:\Temp\ -ToSession $s
Copy-Item -Path C:\Temp\resultat.csv -Destination C:\Temp\Central\ -FromSession $s

# Entrer en interactif (dépannage)
Enter-PSSession -ComputerName srv-fic-01
# ... commandes exécutées à distance ...
Exit-PSSession
```

**Limites** : double-hop (rebond vers une 3ᵉ machine) bloqué par Kerberos par
défaut → utilisez CredSSP (à vos risques, à restreindre) ou mieux : exécutez
depuis la cible avec un gMSA (§46). Et rappelez-vous : **WinRM ≠ RDP** — pour une
console interactive complète, RDP reste nécessaire.

---

## 115. Inventaire matériel en PowerShell

