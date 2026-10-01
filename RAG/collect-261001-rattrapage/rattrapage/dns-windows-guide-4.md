---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-4
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [368, 527]
sha256: 10b6fb2558f8a5a22c71471ee83e001b716a1c0e4d9f9f218254dde371014874
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

Si l'étape 2 échoue mais pas la 1 : problème de redirecteurs/racine ou de pare-feu sortant (voir cas pratique n°7/§107).

---

# Chapitre 3 — Les zones DNS

## 23. Les types de zones : panorama

| Type | Autoritaire ? | Modifiable localement ? | Usage |
|---|---|---|---|
| Primaire (standard) | Oui | Oui (fichier `.dns`) | Petit site sans AD, ou zone exposée |
| Primaire **AD-integrated** | Oui | Oui (répliquée via AD) | **Standard en entreprise** : multi-maîtres, sécurisée |
| Secondaire | Oui (lecture seule) | Non (copie du primaire) | Redondance, serveurs hors AD, DMZ |
| Stub | Non | Non (NS + SOA + glue uniquement) | Connaître les NS autoritaires d'une zone distante |
| Redirecteur conditionnel | Non | Non (renvoi vers des IP) | Router une zone vers des DNS spécifiques |

Règle : en domaine AD, **toutes les zones internes sont AD-integrated** sauf besoin particulier (zone exposée en DMZ → primaire standard + secondaire).

## 24. Zones primaires standard

Stockées dans `%SystemRoot%\System32\dns\<zone>.dns` (format texte). Un seul serveur inscriptible ; les autres sont secondaires.

```powershell
Add-DnsServerPrimaryZone -Name "dmz.contoso.local" -ZoneFile "dmz.contoso.local.dns" -DynamicUpdate None
```

Limites : fichier local = point de défaillance unique pour l'écriture, transferts de zone à configurer à la main, pas de mises à jour dynamiques sécurisées.

## 25. Zones secondaires

Copie en lecture seule d'une zone primaire, synchronisée par **transfert de zone** (AXFR/IXFR, chapitre 5). Utile pour :
- un DNS en DMZ qui ne doit pas être DC ;
- un site distant sans AD (liaison lente) ;
- un serveur BIND/Linux qui fait office de secondaire.

```powershell
Add-DnsServerSecondaryZone -Name "contoso.local" -ZoneFile "contoso.local.dns" -MasterServers "10.0.0.10"
# Forcer le rechargement depuis le maître
Sync-DnsServerZone -Name "contoso.local" -PassThru
```

Le secondaire **ne répond pas** aux mises à jour dynamiques : il les *refère* vers le primaire (les clients Windows suivent la référence automatiquement).

## 26. Zones stub

Une zone stub ne contient que les enregistrements **NS** de la zone distante (+ SOA + glue A). Le serveur stub interroge ensuite directement les NS distants pour chaque requête : il a toujours l'info à jour sans transfert complet.

```powershell
Add-DnsServerStubZone -Name "partenaire.local" -MasterServers "192.168.50.10" -ZoneFile "partenaire.local.dns"
```

**Stub vs redirecteur conditionnel** (erreur n°14, §136) : le stub suit automatiquement les changements de NS distants (il recharge les NS) ; le redirecteur conditionnel pointe vers des IP fixes. Préférez le stub quand l'autre partie administre ses NS et peut en ajouter.

## 27. Zones de recherche directe vs inversée

- **Directe** (`contoso.local`) : nom → IP (A, AAAA, CNAME, MX, SRV…).
- **Inversée** (`10.in-addr.arpa`, `1.0.10.in-addr.arpa`) : IP → nom (PTR). Indispensable pour : les logs lisibles, certains contrôles anti-spam, `nslookup` inverse, et des applications qui vérifient le reverse (SSH, monitoring).

```powershell
# Zone inversée pour 10.1.0.0/24 (fichier 1.10.in-addr.arpa.dns)
Add-DnsServerPrimaryZone -NetworkID "10.1.0.0/24" -ZoneFile "1.10.in-addr.arpa.dns" -DynamicUpdate Secure
```

**Créez une zone inversée pour chaque sous-réseau interne.** C'est l'oubli le plus fréquent (erreur n°12, §134).

## 28. Zones intégrées à Active Directory

La zone est stockée **dans la base AD** (partition d'application), pas dans un fichier. Chaque DC/DNS qui héberge la zone est inscriptible (multi-maîtres) : les mises à jour dynamiques sont acceptées partout et répliquées par AD.

```powershell
Add-DnsServerPrimaryZone -Name "contoso.local" -ReplicationScope "Forest" -DynamicUpdate "Secure"
```

`-ReplicationScope` : `Forest` (tous les DNS de la forêt), `Domain` (tous les DNS du domaine), `Legacy` (tous les DC, même sans DNS — compatibilité 2000). **Utilisez `Forest` pour `_msdcs` et les zones partagées, `Domain` par défaut sinon.**

## 29. Portées de réplication AD

| Portée | Partition AD | Répliqué vers |
|---|---|---|
| `Forest` | `ForestDnsZones` | Tous les DC/DNS de la **forêt** |
| `Domain` | `DomainDnsZones` | Tous les DC/DNS du **domaine** |
| `Legacy` | Partition de domaine | Tous les **DC** (compat Windows 2000) |
| Personnalisée | Partition d'application dédiée | DC inscrits manuellement (`dnscmd /EnlistDirectoryPartition`) |

```powershell
# Voir la partition d'une zone
Get-DnsServerZone -Name "contoso.local" | Select-Object ZoneName, IsDsIntegrated, DirectoryPartitionName
```

## 30. Avantages et limites des zones AD-integrated

**Avantages :** multi-maîtres (pas de primaire unique), réplication AD chiffrée et compressée (pas d'AXFR en clair), mises à jour dynamiques **sécurisées** (ACL par enregistrement), pas de fichier à sauvegarder séparément (sauvegarde = sauvegarde AD), convergence rapide intra-site.

**Limites :** uniquement entre DC Windows ; un secondaire BIND ne peut pas répliquer depuis AD (il faut un primaire standard ou exposer un transfert depuis un DC) ; la zone n'existe que si AD est sain (un problème de réplication AD = un problème DNS).

## 31. Créer une zone primaire (GUI + PowerShell)

GUI : clic droit **Zones de recherche directe → Nouvelle zone → Zone principale →** cocher **« Stocker la zone dans Active Directory »** → portée de réplication → mises à jour dynamiques.

```powershell
# AD-integrated, mises à jour sécurisées (le standard)
Add-DnsServerPrimaryZone -Name "contoso.local" -ReplicationScope "Domain" -DynamicUpdate "Secure" -PassThru

# Vérification
Get-DnsServerZone -Name "contoso.local" | Format-List ZoneName, ZoneType, IsDsIntegrated, DynamicUpdate, IsAutoCreated
```

`-DynamicUpdate` accepte : `None`, `NonsecureAndSecure`, `Secure`. En AD : toujours `Secure` (voir §78).

## 32. Créer une zone inversée

```powershell
# IPv4 : 10.2.0.0/24
Add-DnsServerPrimaryZone -NetworkID "10.2.0.0/24" -ReplicationScope "Domain" -DynamicUpdate "Secure"

# IPv6 : 2001:db8:10::/64 → la zone ip6.arpa est générée en nibbles inversés
Add-DnsServerPrimaryZone -NetworkID "2001:db8:10::/64" -ReplicationScope "Domain" -DynamicUpdate "Secure"

# Vérifier le nom généré
Get-DnsServerZone | Where-Object { $_.IsReverseLookupZone } | Select-Object ZoneName
```

Convention de nommage : pour `10.2.0.0/24`, la zone est `2.10.in-addr.arpa` (octets inversés). Pour un `/16`, c'est `10.in-addr.arpa`.

## 33. Convertir une zone standard en zone AD-integrated

GUI (méthode recommandée) : Propriétés de la zone → onglet **Général** → bouton **Modifier** à côté de « Type : Principal standard » → cocher **« Stocker la zone dans Active Directory »** → choisir la portée.

En ligne de commande (hérité) :

```powershell
dnscmd SRV-DNS-01 /ZoneChangeDirectoryPartition "contoso.local" /Forest
# Valeurs possibles : /Forest | /Domain | /Legacy
Get-DnsServerZone -Name "contoso.local" | Select-Object ZoneName, IsDsIntegrated
```

⚠️ La conversion inverse (AD → fichier) existe aussi mais fait perdre les mises à jour sécurisées : à éviter sauf besoin explicite (serveur DNS non-DC).

## 34. Suspendre, recharger, supprimer une zone

```powershell
# Suspendre (la zone ne répond plus mais reste configurée)
Suspend-DnsServerZone -Name "old.contoso.local" -PassThru
Resume-DnsServerZone  -Name "old.contoso.local" -PassThru

# Recharger depuis AD ou depuis le fichier
Sync-DnsServerZone -Name "contoso.local"

# Supprimer (irréversible : -Force ne demande pas confirmation)
Remove-DnsServerZone -Name "old.contoso.local" -Force
```

⚠️ `Remove-DnsServerZone` sur une zone AD-integrated la supprime de **tout l'annuaire** (réplication). Vérifiez deux fois le nom.

---

# Chapitre 4 — Les enregistrements de ressources

## 35. Anatomie d'un enregistrement

