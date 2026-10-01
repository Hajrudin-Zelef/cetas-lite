---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-8
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [1196, 1386]
sha256: dba7b7ea4976681b988c5360078de82440da3e3bbb984ca254839f522e34c220
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

## 54. Désactiver / supprimer stratégies et filtres

```powershell
# Désactiver une stratégie (sans la supprimer)
Set-DhcpServerv4Policy -Name "PXE-BIOS" -ScopeId 192.0.2.0 -Enabled $false

# Supprimer
Remove-DhcpServerv4Policy -Name "PXE-BIOS" -ScopeId 192.0.2.0

# Retirer un filtre
Remove-DhcpServerv4Filter -List Deny -MacAddress "AA-BB-CC-DD-EE-FF"

# Désactiver une liste de filtres
Set-DhcpServerv4FilterList -Allow $false
```

## 55. Cas d'usage : VLAN imprimantes verrouillé (lien métier copieurs)

Objectif : sur le VLAN 50 (MFP), **seuls les copieurs enregistrés** obtiennent une IP.

```powershell
# 1. Étendue dédiée MFP (plage = que des réservations potentielles)
Add-DhcpServerv4Scope -Name "VLAN50-MFP" `
    -StartRange 192.0.6.10 -EndRange 192.0.6.100 `
    -SubnetMask 255.255.255.0 `
    -LeaseDuration (New-TimeSpan -Days 30) -State Active

# 2. Exclure toute la plage des attributions dynamiques...
Add-DhcpServerv4ExclusionRange -ScopeId 192.0.6.0 -StartRange 192.0.6.10 -EndRange 192.0.6.100

# 3. ... et ne servir que via réservations (une par MFP, voir section 22)
# (les adresses exclues restent attribuables par réservation !)

# 4. Options MFP : pas de passerelle Internet, DNS interne, SMTP dédié
Set-DhcpServerv4OptionValue -ScopeId 192.0.6.0 `
    -DnsServer 192.0.2.53 -DnsDomain "contoso.local" -OptionId 69 -Value 192.0.6.25
```

> 💡 Une adresse **exclue** peut quand même faire l'objet d'une **réservation** : c'est le pattern idéal pour un parc de copieurs — zéro attribution sauvage, 100 % traçable.

## 56. Tester une stratégie avant de la généraliser

```powershell
# Simuler : quel bail recevrait cette MAC ?
# (pas de cmdlet natif de simulation — méthode manuelle :)

# 1. Créer la stratégie désactivée
Add-DhcpServerv4Policy -Name "TEST-Telephone" -ScopeId 192.0.4.0 `
    -Condition OR -MacAddress EQ, "001B2A*" -Enabled $false

# 2. Vérifier la correspondance du pattern avec un bail existant
Get-DhcpServerv4Lease -ScopeId 192.0.4.0 |
    Where-Object { $_.ClientId -like "00-1B-2A*" } |
    Select-Object IPAddress, HostName, ClientId

# 3. Activer après validation
Set-DhcpServerv4Policy -Name "TEST-Telephone" -ScopeId 192.0.4.0 -Enabled $true
```

## 57. Ordre de traitement complet (précédence finale)

Quand un DISCOVER arrive, le serveur évalue dans cet ordre :

```text
1. Filtres MAC (Allow/Deny) → refusé ? STOP
2. Stratégies de l'étendue (dans l'ordre ProcessingOrder)
3. Réservation (MAC connue ?)
4. Classe utilisateur / fournisseur du client
5. Options d'étendue → options de serveur (héritage)
```

## 58. Bonnes pratiques — Scopes avancés (récap)

1. ✅ Superscope uniquement pour migration ou multinet justifié — pas comme solution au manque d'adresses (préférer un masque plus large).
2. ✅ Stratégies PXE testées sur une machine avant généralisation.
3. ✅ Filtre Allow réservé aux VLANs contrôlés (MFP, labos).
4. ✅ Nommer les stratégies explicitement (`PXE-UEFI`, `Telephones-ConstructeurX`).
5. ❌ Pas de chevauchement de plages entre stratégies d'une même étendue.
6. ❌ Pas de filtre Deny comme seule sécurité (MAC spoofable).

---

# Bloc E — Haute disponibilité et multi-VLAN

## 59. Basculement DHCP (failover) : principes

Depuis Windows Server 2012, le basculement **natif** remplace l'ancien cluster DHCP ou le "80/20". Deux modes :

| | Hot Standby (actif/passif) | Load Balance (actif/actif) |
|---|---|---|
| Principe | Primaire actif, secondaire en attente | Charge répartie (ratio configurable) |
| Ratio | — | 50/50 par défaut |
| Bascule | Automatique si primaire injoignable | Les deux servent en permanence |
| Cas d'usage | Secondaire moins puissant, simplicité | Deux serveurs équivalents, répartition |
| Réplication | Continue (baux + config) | Continue (baux + config) |

> 💡 **Recommandation** : **Load Balance 50/50** en général (les deux serveurs travaillent, pas de capacité dormante). **Hot Standby** si le secondaire est une VM légère ou sur un site distant.

Prérequis :

- Les 2 serveurs **autorisées dans AD** (section 9).
- L'étendue existe **déjà sur les deux serveurs** (même ScopeId) — ou la créer via le failover (elle sera répliquée).
- Connectivité TCP **647** entre les deux serveurs (port de réplication failover).

## 60. Configurer le basculement — Load Balance pas à pas

```powershell
# Sur le PRIMAIRE (srv-dhcp-01) : créer la relation de basculement
# L'étendue 192.0.2.0 doit exister sur les deux serveurs

Add-DhcpServerv4Failover `
    -Name "Failover-LAN" `
    -ScopeId 192.0.2.0 `
    -PartnerServer "srv-dhcp-02.contoso.local" `
    -Mode LoadBalance `
    -LoadBalancePercent 50 `
    -SharedSecret "MotDePasse-Partage-Securise-123!" `
    -ServerRole Active

# Vérifier
Get-DhcpServerv4Failover -ComputerName "srv-dhcp-01.contoso.local"
```

Paramètres clés :

- `-SharedSecret` : secret partagé pour authentifier la réplication (le changer régulièrement, le stocker en coffre).
- `-LoadBalancePercent` : % des requêtes traitées par le local (50 = équilibré).
- `-MaxClientLeadTime` : durée max d'extension de bail par le partenaire (défaut 1 h).

En GUI : console DHCP → clic droit sur l'étendue → **Configurer le basculement** → assistant (choix du partenaire, mode, secret partagé).

## 61. Configurer le basculement — Hot Standby pas à pas

```powershell
Add-DhcpServerv4Failover `
    -Name "Failover-WiFi" `
    -ScopeId 192.0.3.0, 192.0.4.0 `
    -PartnerServer "srv-dhcp-02.contoso.local" `
    -Mode HotStandby `
    -ReservePercent 5 `
    -SharedSecret "MotDePasse-Partage-Securise-123!" `
    -ServerRole Active `
    -SwitchoverInterval (New-TimeSpan -Minutes 60)

# -ReservePercent 5 : le standby garde 5 % des adresses pour les nouveaux clients
#                    si le primaire tombe (évite de tout donner d'un coup)
# -SwitchoverInterval : bascule automatique après 60 min d'indisponibilité du primaire
```

> ⚠️ En Hot Standby, le serveur en `Standby` ne répond pas aux DISCOVER tant que le primaire est joignable. Vérifier le rôle : `Get-DhcpServerv4Failover | Select-Object Name, Mode, ServerRole, State`.

## 62. Réplication du basculement : ce qui est répliqué (et ce qui ne l'est pas)

**Répliqué automatiquement** :

- Baux (leases) — en continu
- Réservations
- Exclusions
- Options d'étendue
- Stratégies
- Durée du bail

**NON répliqué** (à configurer manuellement sur chaque serveur) :

- Options de **serveur** (niveau serveur, section 38)
- Filtres MAC (Allow/Deny)
- Classes personnalisées
- Configuration du DNS dynamique (compte dédié, section 71)
- Paramètres d'audit/journalisation

```powershell
# Forcer une réplication immédiate
Invoke-DhcpServerv4FailoverReplication `
    -ComputerName "srv-dhcp-01.contoso.local" `
    -Name "Failover-LAN"

# Vérifier l'état de la réplication
Get-DhcpServerv4Failover -ComputerName "srv-dhcp-01.contoso.local" |
    Select-Object Name, Mode, State, ServerRole
# State doit être "Normal"
```

> 💡 **Piège classique** : on modifie une option de niveau **serveur** sur le primaire et on s'étonne qu'elle ne se réplique pas. Script de synchro manuelle en section 63.

## 63. Script : synchroniser les options de niveau serveur entre les 2 nœuds

```powershell
# À exécuter après toute modif d'option serveur / filtre / classe sur le primaire
$primary   = "srv-dhcp-01.contoso.local"
$secondary = "srv-dhcp-02.contoso.local"

