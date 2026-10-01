---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-8
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [959, 1081]
sha256: 07468dc2c92be7ba9d644c41a738ac4b0504e9eaf376e634831cbf5638b44793
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

**Boucle de redirecteurs** : A redirige vers B, B redirige vers A → timeouts, logs saturés. Ça arrive lors de fusions d'entreprises (chaque partie pointe vers l'autre « pour résoudre ses noms »).

**Règles anti-boucle :**
1. Documentez la **chaîne complète** de résolution (qui redirige vers qui) dans le dossier d'exploitation.
2. Un redirecteur doit toujours être « plus proche de la réponse » que vous (résolveur Internet, ou DNS autoritaire de la zone cible).
3. Testez avec `Resolve-DnsName -Trace` : une boucle se voit (requêtes qui tournent).
4. Méfiez-vous des redirecteurs conditionnels croisés entre deux forêts : préférez une **approbation de forêt + DNS intégré**, ou des zones stub.

```powershell
# Détecter une boucle : la trace montre des allers-retours entre les mêmes IP
Resolve-DnsName -Name "srv.partenaire.local" -Trace | Select-Object -Last 20
```

---

# Chapitre 8 — Split-brain DNS et stratégies DNS

## 63. Le problème du split-brain

Le **split-brain** (ou DNS à « double visage ») : le même nom de zone (`contoso.com`) doit résoudre **différemment** selon qu'on est dedans ou dehors.
Exemple : `www.contoso.com` = `203.0.113.10` (IP publique, site web) pour Internet, mais `10.0.1.50` (IP interne) pour les postes du LAN — parce que le trafic interne ne doit pas faire l'aller-retour par le pare-feu (hairpinning), ou parce que l'appli n'est pas exposée.

Sans split-brain géré, les postes internes résolvent l'IP publique → latence, pare-feu qui bloque le hairpin, ou pire : injoignable.

## 64. Solutions classiques au split-brain

| Solution | Principe | Avantages / limites |
|---|---|---|
| **Zone interne dédiée** | Créer `contoso.com` en zone AD-integrated avec *uniquement* les enregistrements internes nécessaires | Simple, robuste. ⚠️ Il faut dupliquer chaque enregistrement public utilisé en interne (oubli = panne). |
| **Redirecteur conditionnel** | `contoso.com` → DNS publics de l'entreprise pour l'externe | Ne permet pas de différencier interne/externe : tout le monde a la même vue. |
| **Sous-zone** | `interne.contoso.com` en interne, le reste via redirecteurs | Propre si on peut imposer un nommage distinct. |
| **Stratégies DNS (2016+)** | Même zone, **vues différentes** selon le sous-réseau client | Le plus élégant : une seule zone à maintenir (§65). |

**Règle :** si votre nom de domaine AD est déjà un nom public (ex : forêt `contoso.com` au lieu de `contoso.local`), le split-brain n'est pas une option, c'est une obligation.

## 65. Split-brain avec les stratégies DNS (zone scopes)

Depuis Windows Server 2016, une zone peut avoir plusieurs **scopes** (vues) : chaque scope contient ses propres enregistrements, et une **stratégie de résolution** choisit le scope selon le client (sous-réseau, heure, etc.).

```powershell
# 1. Déclarer les sous-réseaux clients
Add-DnsServerClientSubnet -Name "Subnet-LAN" -IPv4Subnet @("10.0.0.0/8")
Add-DnsServerClientSubnet -Name "Subnet-VPN" -IPv4Subnet @("172.16.0.0/16")

# 2. Créer les scopes dans la zone (le scope par défaut "." existe déjà)
Add-DnsServerZoneScope -ZoneName "contoso.com" -Name "Scope-Interne"
Add-DnsServerZoneScope -ZoneName "contoso.com" -Name "Scope-VPN"

# 3. Enregistrements différenciés par scope
Add-DnsServerResourceRecordA -Name "www" -ZoneName "contoso.com" -IPv4Address "10.0.1.50" -ZoneScope "Scope-Interne"
Add-DnsServerResourceRecordA -Name "www" -ZoneName "contoso.com" -IPv4Address "172.16.0.50" -ZoneScope "Scope-VPN"
Add-DnsServerResourceRecordA -Name "www" -ZoneName "contoso.com" -IPv4Address "203.0.113.10"   # scope par défaut = externe

# 4. Stratégies : quel client voit quel scope (le chiffre = ordre d'évaluation)
Add-DnsServerQueryResolutionPolicy -Name "Policy-Interne" -Action ALLOW -ClientSubnet "EQ,Subnet-LAN" -ZoneName "contoso.com" -ZoneScope "Scope-Interne,1"
Add-DnsServerQueryResolutionPolicy -Name "Policy-VPN"     -Action ALLOW -ClientSubnet "EQ,Subnet-VPN" -ZoneName "contoso.com" -ZoneScope "Scope-VPN,1"

# 5. Vérifier
Get-DnsServerQueryResolutionPolicy -ZoneName "contoso.com" | Format-Table Name, Action, ProcessingOrder
Get-DnsServerResourceRecord -ZoneName "contoso.com" -Name "www" -RRType A
```

⚠️ Version : stratégies DNS = **Windows Server 2016 et +** (2019/2022/2025 OK). Les scopes ne se répliquent que si la zone est AD-integrated.

## 66. Stratégies DNS : client subnets

Le **sous-réseau client** est déterminé par l'IP source de la requête (ou par l'info EDNS Client Subnet si elle traverse un redirecteur qui la transmet — rare en interne, à ne pas supposer).

```powershell
# Lister / affiner les subnets
Get-DnsServerClientSubnet | Format-Table Name, IPv4Subnet
# Un subnet peut combiner plusieurs plages
Set-DnsServerClientSubnet -Name "Subnet-LAN" -IPv4Subnet @("10.0.0.0/8","192.168.0.0/16")
```

**Piège :** un client dont l'IP n'est dans aucun subnet tombe sur le **scope par défaut**. En split-brain, ça veut dire « vue externe » : vérifiez que c'est bien le comportement voulu (ou ajoutez une stratégie de rattrapage explicite).

## 67. Stratégies DNS : exemples concrets

```powershell
# Exemple 1 : Géo-DNS interne — Lyon voit le serveur de fichiers de Lyon
Add-DnsServerClientSubnet -Name "Subnet-Lyon" -IPv4Subnet "10.2.0.0/16"
Add-DnsServerZoneScope -ZoneName "contoso.local" -Name "Scope-Lyon"
Add-DnsServerResourceRecordA -Name "fichiers" -ZoneName "contoso.local" -IPv4Address "10.2.1.20" -ZoneScope "Scope-Lyon"
Add-DnsServerResourceRecordA -Name "fichiers" -ZoneName "contoso.local" -IPv4Address "10.1.1.20"  # défaut = Paris
Add-DnsServerQueryResolutionPolicy -Name "Policy-Lyon" -Action ALLOW -ClientSubnet "EQ,Subnet-Lyon" -ZoneName "contoso.local" -ZoneScope "Scope-Lyon,1"

# Exemple 2 : politique horaire — la nuit, le portail de maintenance répond "en maintenance"
Add-DnsServerQueryResolutionPolicy -Name "Policy-Nuit" -Action ALLOW -TimeOfDay "EQ,22:00-06:00" `
    -ZoneName "contoso.com" -ZoneScope "Scope-Maintenance,1" -PassThru

# Exemple 3 : bloquer (DENY) la résolution d'une zone pour un subnet invité
Add-DnsServerQueryResolutionPolicy -Name "Policy-Invite-Deny" -Action DENY `
    -ClientSubnet "EQ,Subnet-Invite" -FQDN "EQ,*.interne.contoso.local" -PassThru
```

## 68. Bonnes pratiques split-brain

1. **Documentez la matrice** : pour chaque nom critique, qui voit quoi (tableau LAN / VPN / Externe).
2. **Ne dupliquez que le nécessaire** : chaque enregistrement public utilisé en interne doit exister dans la vue interne (le classique oublié : `autodiscover`, `vpn`, les enregistrements de vérification).
3. **Testez les 3 vues** après chaque changement (depuis une vraie machine de chaque subnet).
4. Préférez les **stratégies DNS** à la duplication manuelle de zones : une seule zone = un seul endroit à maintenir.
5. Surveillez la **divergence** : un script hebdo qui compare les scopes (export + diff).

---

# Chapitre 9 — DNSSEC

## 69. DNSSEC : pourquoi

Le DNS classique ne garantit **ni l'authenticité ni l'intégrité** : une réponse peut être forgée en transit (cache poisoning, homme-du-milieu). **DNSSEC** ajoute des signatures cryptographiques aux enregistrements : le résolveur peut *vérifier* que la réponse vient bien du propriétaire de la zone et n'a pas été altérée.

Ce que DNSSEC **fait** : authentifie l'origine des données, prouve l'inexistence (NSEC/NSEC3).
Ce que DNSSEC **ne fait pas** : ne chiffre pas (les requêtes restent lisibles — pour la confidentialité, voir DoT/DoH), ne garantit pas la disponibilité.

En entreprise : signez au minimum vos **zones exposées** (publiques) ; en interne, DNSSEC a du sens si votre menace inclut un attaquant sur le réseau (VLAN invité, Wi-Fi).

## 70. Chaîne de confiance : KSK, ZSK, DS

```
Racine (.) signée → DS du TLD → TLD signé → DS de contoso.com → contoso.com signé
```

