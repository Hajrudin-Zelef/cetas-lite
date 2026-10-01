---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-5
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [630, 821]
sha256: efa08ea4bfa6f580b70ed5450f422af186d447ea445ac3a6e1933d316e9be542
---

# Alerte si une étendue dépasse 85 % (à mettre en supervision, voir section 108)
$seuil = 85
Get-DhcpServerv4ScopeStatistics | Where-Object { $_.PercentageInUse -gt $seuil } |
    ForEach-Object {
        Write-Warning "Étendue $($_.ScopeId) à $([math]::Round($_.PercentageInUse,1)) % — prévoir une extension !"
    }
```

> 🎯 **Seuil d'alerte** : 80 % = vigilance, 90 % = action (étendre la plage ou réduire la durée des baux). À 100 %, les nouveaux clients **n'obtiennent plus d'IP** (cas pratique n°2, section 78).

## 25. Étendre ou réduire une étendue

```powershell
# Étendre : passer de .50-.200 à .50-.250
Set-DhcpServerv4Scope -ScopeId 192.0.2.0 -StartRange 192.0.2.50 -EndRange 192.0.2.250

# Réduire : attention, les baux existants hors nouvelle plage restent valides jusqu'à expiration
Set-DhcpServerv4Scope -ScopeId 192.0.2.0 -StartRange 192.0.2.50 -EndRange 192.0.2.150
```

> ⚠️ On ne peut étendre que **dans les limites du masque** de l'étendue. Pour passer de /24 à /23, il faut **recréer l'étendue** (supprimer + recréer) — planifier en heure creuse, avec basculement (section 59) pour éviter la coupure.

## 26. Supprimer une étendue (procédure propre)

```powershell
# 1. Désactiver d'abord (les clients gardent leur IP jusqu'à expiration)
Set-DhcpServerv4Scope -ScopeId 192.0.5.0 -State Inactive

# 2. Supprimer les baux si besoin de libérer vite
# Get-DhcpServerv4Lease -ScopeId 192.0.5.0 | Remove-DhcpServerv4Lease

# 3. Supprimer l'étendue
Remove-DhcpServerv4Scope -ScopeId 192.0.5.0
```

## 27. Pool d'adresses : lecture et interprétation

La console affiche le **pool d'adresses** = plage totale + exclusions + état. En PowerShell :

```powershell
# Vue synthétique : plage, exclusions, baux
$scope = Get-DhcpServerv4Scope -ScopeId 192.0.2.0
$stats = Get-DhcpServerv4ScopeStatistics -ScopeId 192.0.2.0
[PSCustomObject]@{
    Étendue     = $scope.Name
    Plage       = "$($scope.StartRange) - $($scope.EndRange)"
    Masque      = $scope.SubnetMask
    Bail        = $scope.LeaseDuration
    État        = $scope.State
    Libres      = $stats.Free
    Utilisées   = $stats.InUse
    Réservées   = $stats.Reserved
    Utilisation = "$([math]::Round($stats.PercentageInUse,1)) %"
} | Format-List
```

## 28. Baux : nettoyage automatique des baux expirés

Windows DHCP **ne supprime pas automatiquement** les baux expirés de la base (ils restent en état `Expired`). Pour une base propre :

```powershell
# Script de nettoyage hebdomadaire (à planifier, voir section 108)
foreach ($scope in Get-DhcpServerv4Scope) {
    $expired = Get-DhcpServerv4Lease -ScopeId $scope.ScopeId |
        Where-Object { $_.AddressState -eq "Expired" -or $_.AddressState -eq "Released" }
    if ($expired) {
        Write-Output "Étendue $($scope.ScopeId) : $($expired.Count) baux expirés supprimés"
        $expired | Remove-DhcpServerv4Lease
    }
}
```

## 29. Identifier un client par sa MAC (ClientId)

Le `ClientId` affiché par `Get-DhcpServerv4Lease` est la MAC en hexadécimal avec tirets. Pour retrouver un équipement :

```powershell
# Rechercher tous les baux d'une MAC (ex : un copieur qui change de VLAN)
$mac = "00-1B-A9-3F-2C-7D"
Get-DhcpServerv4Scope | ForEach-Object {
    Get-DhcpServerv4Lease -ScopeId $_.ScopeId |
        Where-Object { $_.ClientId -eq $mac }
} | Select-Object ScopeId, IPAddress, HostName, LeaseExpiryTime
```

> 💡 Sur le client Windows : `ipconfig /all` → "Adresse physique". Sur un copieur : page de configuration réseau (imprimée depuis le panneau) ou étiquette sous l'appareil.

## 30. Bonnes pratiques — Étendues et baux (récap)

1. ✅ Une étendue par VLAN, nommée explicitement (`VLAN10-LAN-Bureaux`).
2. ✅ Exclusions systématiques pour l'infrastructure en IP statique.
3. ✅ Durée de bail adaptée au profil (8 j fixe, 1 j Wi-Fi, 4 h invités).
4. ✅ Réservations pour les MFP/imprimantes, jamais d'IP statique "sauvage" sur les copieurs.
5. ✅ Supervision du taux d'utilisation (alerte à 80 %).
6. ✅ Documentation du plan d'adressage (section 107).
7. ❌ Jamais deux serveurs DHCP **indépendants** sur la même étendue sans basculement configuré (conflits garantis).
8. ❌ Jamais d'étendue sans agent de relais sur les VLANs distants.

---

# Bloc C — Options DHCP

## 31. Les options DHCP : principe général

Les options sont des paramètres réseau envoyés **avec** le bail. Le client les applique à sa pile TCP/IP. Chaque option a un **numéro** (RFC 2132) :

```text
DHCPACK : IP=192.0.2.50, masque=255.255.255.0, bail=8 jours
          + option 003 (routeur) = 192.0.2.1
          + option 006 (DNS)     = 192.0.2.53, 192.0.2.54
          + option 015 (domaine) = contoso.local
```

```powershell
# Lister toutes les options définies sur le serveur (catalogue)
Get-DhcpServerv4OptionDefinition | Select-Object OptionId, Name, Type | Format-Table -AutoSize
```

## 32. Les 6 options indispensables en entreprise

| Option | Nom | Valeur type | Obligatoire ? |
|--------|-----|-------------|---------------|
| 003 | Routeur (passerelle) | 192.0.2.1 | ✅ Oui |
| 006 | Serveurs DNS | 192.0.2.53, 192.0.2.54 | ✅ Oui |
| 015 | Nom de domaine DNS | contoso.local | ✅ Oui (AD) |
| 044 | Serveurs WINS/NBNS | 192.0.2.60 | ⚠️ Si WINS encore utilisé |
| 046 | Type de nœud WINS/NBT | 0x8 (hybride) | ⚠️ Avec 044 |
| 051 | Durée du bail | (gérée par l'étendue) | Automatique |

```powershell
# Configurer les options d'une étendue en une commande
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 `
    -Router 192.0.2.1 `
    -DnsServer 192.0.2.53, 192.0.2.54 `
    -DnsDomain "contoso.local" `
    -WinsServer 192.0.2.60 `
    -WinsNodeType 8

# Vérifier
Get-DhcpServerv4OptionValue -ScopeId 192.0.2.0 | Format-Table OptionId, Name, Value -AutoSize
```

## 33. Option 003 — Routeur / passerelle par défaut

```powershell
# Une seule passerelle (cas standard)
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 3 -Value 192.0.2.1

# Plusieurs passerelles (redondance, ordre de préférence)
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 3 -Value 192.0.2.1, 192.0.2.2
```

> ⚠️ La redondance de passerelle via l'option 003 est **médiocre** (pas de détection de panne rapide côté client). En entreprise, préférer **HSRP/VRRP** sur les routeurs : une seule IP virtuelle dans l'option 003.

## 34. Option 006 — Serveurs DNS

```powershell
# DNS primaire + secondaire (toujours 2 minimum en entreprise)
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 6 -Value 192.0.2.53, 192.0.2.54
```

Ordre = ordre d'interrogation par le client. Mettre les **contrôleurs de domaine** (qui hébergent le DNS AD) en premier.

> ⚠️ Ne jamais mettre un DNS public (8.8.8.8) dans l'option 006 d'un poste du domaine : le client ne résoudrait plus les noms AD (`contoso.local`) → ouverture de session impossible.

## 35. Option 015 — Nom de domaine DNS

```powershell
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 15 -Value "contoso.local"
```

Effets :

- Le client complète les noms courts (`ping srv-print` → `srv-print.contoso.local`).
- Utilisé pour l'enregistrement DNS dynamique (section 69).

## 36. Options 044/046 — WINS (héritage, encore présent)

WINS (Windows Internet Name Service) est obsolète mais survit dans les environnements avec de vieilles applications NetBIOS.

```powershell
# Serveurs WINS + type de nœud hybride (0x8 = H-node : WINS d'abord, broadcast ensuite)
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 44 -Value 192.0.2.60
Set-DhcpServerv4OptionValue -ScopeId 192.0.2.0 -OptionId 46 -Value 8
```

Types de nœud : 1 = B-node (broadcast), 2 = P-node (WINS seul), 4 = M-node (mixte), **8 = H-node (hybride, recommandé)**.

> 💡 Si aucun serveur WINS dans l'entreprise : **ne pas configurer** 044/046. Inutile et source de timeouts NetBIOS.

## 37. Options 066/067 — Serveur de démarrage et PXE (lien WDS/MDT)

