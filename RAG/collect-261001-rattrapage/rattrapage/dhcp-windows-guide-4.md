---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-4
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [452, 629]
sha256: 4f054d6ced226bd721740f392fe28031f2cbd1c1ef6bb470bd29e018d1efac0d
---

# Ajouter une exclusion unique (ex : un serveur en .25)
Add-DhcpServerv4ExclusionRange -ScopeId 192.0.2.0 -StartRange 192.0.2.25 -EndRange 192.0.2.25

# Lister les exclusions
Get-DhcpServerv4ExclusionRange -ScopeId 192.0.2.0

# Supprimer une exclusion (quand l'équipement passe en réservation DHCP)
Remove-DhcpServerv4ExclusionRange -ScopeId 192.0.2.0 -StartRange 192.0.2.25 -EndRange 192.0.2.25
```

> ⚠️ **Erreur classique n°1** : oublier d'exclure les IP statiques → le DHCP attribue une IP déjà utilisée → **conflit d'adresses** (voir cas pratique n°1, section 77).

## 18. Durée du bail : comment la choisir

La durée du bail détermine la fréquence des renouvellements et la vitesse de récupération des adresses.

| Contexte | Durée recommandée | Pourquoi |
|----------|-------------------|----------|
| Postes fixes (LAN) | 8 jours (défaut) | Stable, peu de trafic DHCP |
| Portables / Wi-Fi corporate | 1 jour | Les utilisateurs bougent, libère vite les adresses |
| Wi-Fi invités | 4 heures (voire 1 h) | Forte rotation, évite l'épuisement |
| Téléphonie IP | 8 jours | Équipements fixes |
| IoT / GTB / capteurs | 30 jours | Équipements rarement redémarrés |
| Événementiel / formation (salle) | 8 heures | Pic ponctuel d'utilisateurs |

```powershell
# Modifier la durée du bail d'une étendue
Set-DhcpServerv4Scope -ScopeId 192.0.3.0 -LeaseDuration (New-TimeSpan -Days 1)

# Bail invité : 4 heures
Set-DhcpServerv4Scope -ScopeId 192.0.5.0 -LeaseDuration (New-TimeSpan -Hours 4)
```

> 💡 La durée du bail **ne s'applique qu'aux nouveaux baux et renouvellements**. Pour forcer l'application immédiate : supprimer les baux existants (section 20) — à faire en heure creuse.

## 19. Visualiser les baux

```powershell
# Tous les baux d'une étendue
Get-DhcpServerv4Lease -ScopeId 192.0.2.0 |
    Select-Object IPAddress, HostName, ClientId, LeaseExpiryTime, AddressState |
    Format-Table -AutoSize

# Baux actifs uniquement
Get-DhcpServerv4Lease -ScopeId 192.0.2.0 |
    Where-Object { $_.AddressState -eq "Active" }

# Trouver qui a une IP donnée (utile en dépannage)
Get-DhcpServerv4Lease -ScopeId 192.0.2.0 |
    Where-Object { $_.IPAddress -eq "192.0.2.87" } |
    Format-List IPAddress, HostName, ClientId, LeaseExpiryTime

# Compter les baux par état
Get-DhcpServerv4Lease -ScopeId 192.0.2.0 |
    Group-Object AddressState | Select-Object Name, Count
```

États possibles d'un bail (`AddressState`) :

| État | Signification |
|------|---------------|
| `Active` | Bail attribué et valide |
| `Expired` | Bail expiré, adresse récupérable |
| `Released` | Client a libéré (`ipconfig /release`) |
| `ReservationActive` / `ReservationInactive` | Réservation (voir section 22) |

## 20. Supprimer un bail (libérer une adresse)

```powershell
# Supprimer un bail précis (l'adresse redevient disponible)
Remove-DhcpServerv4Lease -ScopeId 192.0.2.0 -IPAddress 192.0.2.87

# Supprimer TOUS les baux expirés/inactifs d'une étendue (nettoyage)
Get-DhcpServerv4Lease -ScopeId 192.0.2.0 |
    Where-Object { $_.AddressState -ne "Active" -and $_.AddressState -notlike "Reservation*" } |
    Remove-DhcpServerv4Lease

# Forcer un client à renouveler : côté client
# ipconfig /release puis ipconfig /renew
```

> ⚠️ Supprimer un bail **actif** ne déconnecte pas le client immédiatement : il garde son IP jusqu'au prochain renouvellement (T1). Pour un effet immédiat, agir côté client (`ipconfig /release`).

## 21. Réconciliation : réparer la base des baux

La base DHCP (`dhcp.mdb` sous `C:\Windows\System32\dhcp`) peut se désynchroniser du registre (après un crash, une restauration...). La **réconciliation** compare et corrige.

```powershell
# Réconcilier une étendue (corrige les incohérences)
Invoke-DhcpServerv4ScopeReconciliation -ScopeId 192.0.2.0

# Réconcilier toutes les étendues
Get-DhcpServerv4Scope | ForEach-Object {
    Invoke-DhcpServerv4ScopeReconciliation -ScopeId $_.ScopeId
}
```

En console : clic droit sur l'étendue → **Réconcilier**. Le rapport indique les baux orphelins ou manquants.

## 22. Réservations : des IP fixes gérées par le DHCP

Une réservation lie **une adresse IP à une adresse MAC** : le client reçoit toujours la même IP, mais la configuration reste centralisée (passerelle, DNS modifiables sans toucher l'équipement).

```powershell
# Créer une réservation : copieur MFP du service comptabilité
Add-DhcpServerv4Reservation `
    -ScopeId 192.0.6.0 `
    -IPAddress 192.0.6.10 `
    -ClientId "00-1B-A9-3F-2C-7D" `
    -Name "MFP-Compta-RDC" `
    -Description "Copieur multifonction - Comptabilité RDC"

# Format du ClientId : MAC en hexadécimal séparé par des tirets
# Pour trouver la MAC d'un équipement : regarder l'étiquette, ou depuis un bail existant :
Get-DhcpServerv4Lease -ScopeId 192.0.6.0 |
    Where-Object { $_.HostName -like "*MFP*" } |
    Select-Object IPAddress, HostName, ClientId

# Lister les réservations
Get-DhcpServerv4Reservation -ScopeId 192.0.6.0 | Format-Table IPAddress, Name, ClientId

# Supprimer une réservation (le copieur est remplacé)
Remove-DhcpServerv4Reservation -ScopeId 192.0.6.0 -ClientId "00-1B-A9-3F-2C-7D"
```

> **Lien métier copieurs** : chaque MFP du parc = **une réservation DHCP** avec un nom explicite (`MFP-<Service>-<Étage>`). Avantages :
>
> 1. **IP fixe garantie** : les pilotes d'impression déployés par GPO, les dossiers de numérisation SMB (`\\srv-print\scans`) et les connecteurs LDAP pointent vers une IP stable.
> 2. **Zéro config sur le copieur** : à l'installation, le technicien met le MFP en DHCP, la réservation fait le reste.
> 3. **Changement de DNS/passerelle** : une modification de l'option d'étendue se propage à tous les MFP au prochain renouvellement — pas besoin de passer sur chaque machine.
> 4. **Traçabilité** : `Get-DhcpServerv4Reservation` = inventaire IP des copieurs, exportable en CSV pour la CMDB.
>
> ⚠️ L'adresse réservée doit être **dans la plage de l'étendue** mais **hors des baux dynamiques** : soit exclue de la plage dynamique, soit dans une étendue dédiée aux MFP (VLAN 50 du plan section 13 — recommandé).

Script : générer les réservations MFP depuis un CSV d'inventaire :

```powershell
# mfp.csv : IPAddress,MacAddress,Name,Location
# 192.0.6.10,00-1B-A9-3F-2C-7D,MFP-Compta-RDC,Comptabilite RDC
Import-Csv "C:\Admin\mfp.csv" | ForEach-Object {
    Add-DhcpServerv4Reservation `
        -ScopeId 192.0.6.0 `
        -IPAddress $_.IPAddress `
        -ClientId $_.MacAddress `
        -Name $_.Name `
        -Description "MFP - $($_.Location)"
}
```

## 23. Réservations : options spécifiques par réservation

Une réservation peut avoir **ses propres options**, qui écrasent celles de l'étendue (voir précédence, section 38) :

```powershell
# Le MFP compta utilise un serveur d'impression dédié comme passerelle de numérisation (exemple)
Set-DhcpServerv4OptionValue `
    -ReservedIP 192.0.6.10 `
    -OptionId 003 `
    -Value 192.0.6.2
```

Cas d'usage : MFP avec **serveur SMTP dédié** pour le scan-to-mail (option 069 — serveurs SMTP), ou **serveur de temps** spécifique.

## 24. Statistiques d'étendue : surveiller le taux d'utilisation

```powershell
# Statistiques d'une étendue
Get-DhcpServerv4ScopeStatistics -ScopeId 192.0.2.0 |
    Select-Object ScopeId, Free, InUse, PercentageInUse, Reserved, Pending

# Toutes les étendues triées par taux d'utilisation (les plus pleines d'abord)
Get-DhcpServerv4ScopeStatistics |
    Select-Object ScopeId,
        @{N="Utilisation_%";E={[math]::Round($_.PercentageInUse,1)}},
        Free, InUse |
    Sort-Object "Utilisation_%" -Descending |
    Format-Table -AutoSize

