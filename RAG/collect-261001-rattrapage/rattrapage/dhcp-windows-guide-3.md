---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-3
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "attribution"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [280, 451]
sha256: 75fdb996ff025048e0883797037e9e1319543df7dd7b3fcaecd5ae9bac56ba9b
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

| Nœud | Actions clés |
|------|--------------|
| Serveur | Autoriser/retirer, sauvegarder/restaurer, configurer le basculement |
| Étendue | Activer/désactiver, créer, réconcilier, statistiques |
| Baux d'adresses | Afficher, supprimer un bail |
| Réservations | Ajouter (MAC + IP) |
| Options d'étendue | Configurer 003/006/015... |

## 11. Le module PowerShell DhcpServer — la vraie console de l'admin

En production, **tout se fait en PowerShell** : c'est scriptable, auditable, et identique sur Server Core.

```powershell
# Importer (automatique si le rôle est installé)
Import-Module DhcpServer

# Lister toutes les cmdlets par verbe
Get-Command -Module DhcpServer | Group-Object { $_.Name.Split('-')[0] } |
    Sort-Object Count -Descending | Format-Table Name, Count -AutoSize

# Familles principales :
#  Get-/Add-/Set-/Remove-DhcpServerv4Scope        -> étendues
#  Get-/Add-/Set-/Remove-DhcpServerv4Lease        -> baux
#  Get-/Add-/Set-/Remove-DhcpServerv4Reservation -> réservations
#  Get-/Set-DhcpServerv4OptionValue               -> options
#  Get-/Add-/Remove-DhcpServerv4ExclusionRange    -> exclusions
#  Get-/Add-DhcpServerv4Failover                  -> basculement
#  Get-/Add-DhcpServerv4Policy                    -> stratégies
#  Get-/Add-DhcpServerv4Filter                    -> filtres MAC
#  Export-/Import-DhcpServer                      -> sauvegarde
```

> 💡 Toutes les commandes de ce guide ciblent le serveur local par défaut. Pour cibler un serveur distant, ajoutez `-ComputerName "srv-dhcp-02.contoso.local"` (WinRM requis).

## 12. Journal des événements DHCP — où regarder en premier

| Journal | Chemin | Événements clés |
|---------|--------|-----------------|
| DHCP-Server | `Applications and Services Logs\Microsoft\Windows\DHCP-Server` | 1020 (étendue épuisée), 1046 (non autorisé), 1034 (basculement) |
| Système | `Windows Logs\System` source `DHCPServer` | Démarrage/arrêt du service |

```powershell
# Les 20 derniers événements DHCP-Server
Get-WinEvent -LogName "Microsoft-Windows-DHCP-Server/Admin" -MaxEvents 20 |
    Select-Object TimeCreated, Id, LevelDisplayName,
        @{N="Message";E={$_.Message.Substring(0,[Math]::Min(120,$_.Message.Length))}} |
    Format-Table -AutoSize

# Surveiller en continu les alertes d'épuisement (ID 1020)
Get-WinEvent -LogName "Microsoft-Windows-DHCP-Server/Admin" |
    Where-Object { $_.Id -eq 1020 } |
    Select-Object TimeCreated, Message | Format-List
```

## 12bis. DHCPv6 — ce qu'il faut savoir (sans s'y noyer)

Windows Server gère aussi le DHCPv6 (attribution d'adresses IPv6 + options comme DNS IPv6). En entreprise française typique, l'IPv6 interne reste rare ; le DHCPv6 sert surtout si tu déploies de l'IPv6 natif.

```powershell
# Créer une étendue DHCPv6 (exemple)
Add-DhcpServerv6Scope -Prefix 2001:db8:1:: -Name "LAN-IPv6" -Preference 255 -ValidLifeTime 8.00:00:00 -PreferredLifeTime 6.00:00:00

# Options DHCPv6 (ex : DNS IPv6 = option 23)
Set-DhcpServerv6OptionValue -ScopeId 2001:db8:1:: -OptionId 23 -Value 2001:db8::53
```

> Ce guide se concentre sur **DHCPv4**, qui reste le standard en entreprise. Les concepts (baux, réservations, basculement) sont transposables.

---

# Bloc B — Étendues et baux

## 13. Concevoir son plan d'adressage avant de créer la moindre étendue

Règle d'or : **on ne crée pas d'étendue sans plan d'adressage écrit et validé**. Modèle de découpage type pour une PME/ETI :

| VLAN | Usage | Sous-réseau | Passerelle | Plage DHCP | Exclusions | Bail |
|------|-------|-------------|------------|------------|------------|------|
| 10 | LAN bureaux | 192.0.2.0/24 | 192.0.2.1 | .50 – .200 | .1 – .49 (infra) | 8 jours |
| 20 | Wi-Fi corporate | 192.0.3.0/24 | 192.0.3.1 | .50 – .200 | .1 – .49 | 1 jour |
| 30 | Voix (ToIP) | 192.0.4.0/24 | 192.0.4.1 | .50 – .200 | .1 – .49 | 8 jours |
| 40 | Invités | 192.0.5.0/24 | 192.0.5.1 | .50 – .200 | .1 – .49 | 4 heures |
| 50 | Imprimantes/MFP | 192.0.6.0/24 | 192.0.6.1 | — (réservations) | tout sauf réservations | — |
| 60 | Serveurs | 192.0.7.0/24 | 192.0.7.1 | — (statique) | tout | — |
| 70 | IoT/GTB | 192.0.8.0/24 | 192.0.8.1 | .50 – .200 | .1 – .49 | 30 jours |

> ⚠️ 192.0.2.0/24, 198.51.100.0/24, 203.0.113.0/24 sont des plages **TEST-NET réservées à la documentation** (RFC 5737) — ne les utilise jamais en production. En réel : 10.0.0.0/8, 172.16.0.0/12 ou 192.168.0.0/16.

## 14. Créer une étendue — console GUI

1. Console DHCP → **Pool d'adresses IPv4** → clic droit → **Nouvelle étendue**.
2. Nom : `LAN-Bureaux` ; description : `VLAN 10 — postes fixes et portables`.
3. Plage : `192.0.2.50` → `192.0.2.200`, masque `255.255.255.0` (longueur 24).
4. **Exclusions** : `192.0.2.1` – `192.0.2.49` (passerelle, switchs, points d'accès en statique).
5. Durée du bail : 8 jours (défaut).
6. Options : routeur `192.0.2.1`, DNS `192.0.2.53`, nom de domaine `contoso.local`.
7. **Activer l'étendue** : Oui.

## 15. Créer une étendue — PowerShell (reproductible)

```powershell
# Créer l'étendue LAN bureaux
Add-DhcpServerv4Scope `
    -Name "LAN-Bureaux" `
    -Description "VLAN 10 - Postes fixes et portables" `
    -StartRange 192.0.2.50 `
    -EndRange 192.0.2.200 `
    -SubnetMask 255.255.255.0 `
    -LeaseDuration (New-TimeSpan -Days 8) `
    -State Active

# Ajouter l'exclusion pour l'infrastructure en statique
Add-DhcpServerv4ExclusionRange `
    -ScopeId 192.0.2.0 `
    -StartRange 192.0.2.1 `
    -EndRange 192.0.2.49

# Configurer les options d'étendue (routeur, DNS, domaine)
Set-DhcpServerv4OptionValue `
    -ScopeId 192.0.2.0 `
    -Router 192.0.2.1 `
    -DnsServer 192.0.2.53, 192.0.2.54 `
    -DnsDomain "contoso.local"

# Vérifier
Get-DhcpServerv4Scope -ScopeId 192.0.2.0 | Format-List *
Get-DhcpServerv4ExclusionRange -ScopeId 192.0.2.0
Get-DhcpServerv4OptionValue -ScopeId 192.0.2.0
```

Créer **plusieurs étendues** d'un coup (script de provisioning) :

```powershell
$scopes = @(
    @{ Name="LAN-Bureaux";  Id="192.0.2.0"; Start="192.0.2.50";  End="192.0.2.200"; Router="192.0.2.1"; LeaseDays=8 }
    @{ Name="WiFi-Corp";    Id="192.0.3.0"; Start="192.0.3.50";  End="192.0.3.200"; Router="192.0.3.1"; LeaseDays=1 }
    @{ Name="Voix-ToIP";    Id="192.0.4.0"; Start="192.0.4.50";  End="192.0.4.200"; Router="192.0.4.1"; LeaseDays=8 }
    @{ Name="Guest";         Id="192.0.5.0"; Start="192.0.5.50";  End="192.0.5.200"; Router="192.0.5.1"; LeaseDays=0 }
)

foreach ($s in $scopes) {
    $lease = if ($s.LeaseDays -eq 0) { New-TimeSpan -Hours 4 } else { New-TimeSpan -Days $s.LeaseDays }
    Add-DhcpServerv4Scope -Name $s.Name -StartRange $s.Start -EndRange $s.End `
        -SubnetMask 255.255.255.0 -LeaseDuration $lease -State Active
    Set-DhcpServerv4OptionValue -ScopeId $s.Id -Router $s.Router `
        -DnsServer 192.0.2.53 -DnsDomain "contoso.local"
}
```

## 16. Activer / désactiver une étendue

```powershell
# Désactiver (maintenance : le serveur ne distribue plus sur cette étendue)
Set-DhcpServerv4Scope -ScopeId 192.0.2.0 -State Inactive

# Réactiver
Set-DhcpServerv4Scope -ScopeId 192.0.2.0 -State Active

# Vérifier l'état de toutes les étendues
Get-DhcpServerv4Scope | Select-Object Name, ScopeId, State, StartRange, EndRange
```

> 💡 Désactiver une étendue **ne supprime pas les baux existants** : les clients gardent leur IP jusqu'à expiration. Pour forcer le renouvellement : `ipconfig /release` + `/renew` côté client, ou réduire temporairement la durée du bail.

## 17. Exclusions : protéger les IP statiques

Les exclusions définissent des adresses **jamais attribuées** par le DHCP. Usage : passerelle, switchs, imprimantes en statique, serveurs, points d'accès.

```powershell
# Ajouter une exclusion
Add-DhcpServerv4ExclusionRange -ScopeId 192.0.2.0 -StartRange 192.0.2.1 -EndRange 192.0.2.20

