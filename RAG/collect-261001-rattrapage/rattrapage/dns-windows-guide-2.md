---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-2
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [58, 199]
sha256: b38565d17d7ddfff9f2f2a9e7a9b14ccee485cd4d586064208484795bd7374da
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

C'est la distinction la plus mal comprise — et la source de la moitié des erreurs d'architecture.

**Résolution récursive (côté client → résolveur) :** le client pose une question et attend *la* réponse finale. Le résolveur fait tout le travail. C'est ce que fait un poste Windows quand il interroge son serveur DNS configuré.

**Résolution itérative (résolveur → serveurs autoritaires) :** le résolveur interroge la racine, qui répond « je ne sais pas, demande au serveur du TLD », qui répond « demande au serveur autoritaire », etc. Chaque serveur ne donne que le meilleur indice suivant (*referral*).

```powershell
# Voir la différence en pratique : trace complète de la résolution itérative
Resolve-DnsName -Name "www.contoso.local" -Server "10.0.0.10" -DnsOnly
# Avec -Trace (Windows 10/Server 2019+) : affiche chaque étape itérative
Resolve-DnsName -Name "www.example.com" -Trace
```

**Règle d'architecture :** vos serveurs DNS internes doivent offrir la récursion à vos clients internes (sinon rien ne se résout), mais **jamais** au monde entier (serveur exposé = récursion désactivée, §117).

## 4. Hiérarchie DNS : racine, TLD, domaines

```
.  (racine, 13 identités a.root-servers.net … m.root-servers.net)
└── local          ← votre forêt AD (non routé sur Internet, c'est voulu)
│   └── contoso
│       ├── _msdcs              ← enregistrements de localisation des DC
│       ├── Paris-Siege._sites ← SRV spécifiques au site
│       └── srv-fichiers        ← A : 10.0.1.20
└── fr / com / …               ← Internet, via redirecteurs ou racine
```

Le nom de votre forêt AD (`contoso.local`) est une **zone autoritaire** sur vos DC/DNS. Tout ce qui n'est pas dans vos zones est résolu par récursion (redirecteurs ou indications de racine, chapitre 7).

## 5. Le rôle du DNS dans Active Directory

Active Directory **ne fonctionne pas sans DNS**. Kerberos, LDAP, la réplication, l'ouverture de session : tout commence par une requête DNS. Concrètement :

1. Le client cherche un contrôleur de domaine via des requêtes **SRV** (`_ldap._tcp.dc._msdcs.contoso.local`).
2. Il affine par **site AD** (`_ldap._tcp.Paris-Siege._sites.dc._msdcs.contoso.local`) pour contacter le DC le plus proche.
3. Il résout le nom du DC en IP via un enregistrement **A**.
4. Kerberos utilise ensuite ces noms pour obtenir des tickets.

**Conséquence opérationnelle :** un DNS qui répond « vite mais faux » (vieux cache, zone périmée, split-brain mal géré) provoque des ouvertures de session lentes, des GPO non appliquées et des erreurs de réplication. Le DNS est un composant d'infrastructure critique, pas un « service annexe ».

## 6. Les enregistrements SRV d'AD en détail

Le service Netlogon de chaque DC enregistre automatiquement (toutes les heures par défaut, ou via `nltest /dsregdns`) une série de SRV. Syntaxe générique :

```
_service._proto.nom.  TTL  IN  SRV  priorité  poids  port  cible
```

Les plus importants pour `contoso.local` :

| Enregistrement SRV | Rôle |
|---|---|
| `_ldap._tcp.dc._msdcs` | Localiser *un* DC (LDAP, port 389) |
| `_ldap._tcp.<Site>._sites.dc._msdcs` | Localiser un DC *du site* |
| `_kerberos._tcp.dc._msdcs` | Localiser un DC pour Kerberos (port 88) |
| `_kerberos._tcp.<Site>._sites.dc._msdcs` | Kerberos du site |
| `_gc._tcp` | Catalogue global (port 3268) |
| `_gc._tcp.<Site>._sites` | Catalogue global du site |
| `_kpasswd._tcp` / `_kpasswd._udp` | Changement de mot de passe Kerberos (port 464) |
| `_ldap._tcp.pdc._msdcs` | Émulateur PDC (port 389) |

Priorité basse = préféré ; à priorité égale, le **poids** répartit la charge.

## 7. La zone `_msdcs.contoso.local`

`_msdcs` est une **zone séparée**, déléguée depuis `contoso.local`, qui contient les enregistrements critiques pour la forêt entière :

- Les SRV `_tcp.dc`, `_tcp.gc`, `_tcp.pdc`, `_udp.kerberos`…
- Un **CNAME par DC** nommé d'après son GUID d'objet NTDS (`<GUID>._msdcs.contoso.local` → nom du DC). La réplication AD utilise ces CNAME : si l'un manque, la réplication vers ce DC casse.

```powershell
# Lister la zone _msdcs
Get-DnsServerResourceRecord -ZoneName "_msdcs.contoso.local" -RRType SRV |
    Select-Object HostName, @{n='Cible';e={$_.RecordData.DomainName}}
# Vérifier le CNAME GUID d'un DC
Resolve-DnsName -Name "<GUID>._msdcs.contoso.local" -Type CNAME
```

⚠️ **Ne supprimez jamais la zone `_msdcs`**, même si elle « semble vide » dans la console (erreur n°8, §130).

## 8. Sites AD et SRV par site

Les sites AD (`Paris-Siege`, `Lyon-Usine`) servent à orienter les clients vers le DC le plus proche via les SRV `_sites`. Le client détermine son site grâce à son sous-réseau (Sites et services AD → Subnets).

```powershell
# Voir les enregistrements SRV d'un site
Get-DnsServerResourceRecord -ZoneName "_msdcs.contoso.local" |
    Where-Object { $_.HostName -like "*Paris-Siege*" } |
    Select-Object HostName, RecordType
```

**Bonnes pratiques :**
- Un sous-réseau = un site, toujours. Un sous-réseau oublié = des clients qui traversent le WAN pour s'authentifier.
- Au moins un DC (idéalement deux) par site de production + un catalogue global par site.
- Les enregistrements `_sites` sont enregistrés par Netlogon en fonction du site du DC lui-même.

## 9. Enregistrements A des contrôleurs de domaine

Chaque DC enregistre aussi :

- Un **A** à son nom (`SRV-DC-01.contoso.local` → `10.0.0.10`) — utilisé par Kerberos et la réplication.
- Un **A** au nom du **domaine nu** (`contoso.local` → IP de *chaque* DC) : c'est du round-robin DNS qui répartit les requêtes LDAP/Kerberos sur tous les DC. C'est normal d'y voir plusieurs IP.

```powershell
# Les A du domaine nu (round-robin normal)
Resolve-DnsName -Name "contoso.local" -Type A
```

Si vous filtrez ou supprimez ces A « parce qu'il y a trop d'IP », vous cassez la répartition de charge et la localisation des DC.

## 10. Comment un client localise un DC (flux complet)

1. Le client lit son suffixe DNS (`contoso.local`) et son site (via son IP → subnet → site).
2. Requête SRV : `_ldap._tcp.Paris-Siege._sites.dc._msdcs.contoso.local`.
3. Si aucun DC dans le site : repli sur `_ldap._tcp.dc._msdcs.contoso.local` (n'importe quel DC).
4. Tri par priorité/poids, résolution A du DC choisi, connexion LDAP port 389.
5. Kerberos : même logique avec `_kerberos._tcp…` (port 88).

```powershell
# Simuler ce que fait un client (depuis le poste)
nltest /dsgetdc:contoso.local /site:Paris-Siege
nltest /dnsgetdc:contoso.local
```

## 11. DNS et stratégie de nommage

- **Nom de forêt :** un suffixe non routable (`.local`, `.interne`, `.lan`) évite les collisions avec l'Internet public et le besoin de split-brain sur le nom de domaine lui-même. ⚠️ `.local` est aussi utilisé par mDNS/Bonjour : en pratique ça cohabite, mais documentez-le.
- **Noms de serveurs :** courts, sans accents, sans underscore (sauf les noms réservés type `_msdcs`, `_tcp`… qui sont légaux en DNS mais pas en NetBIOS).
- **Évitez les noms à un seul label** (`serveur` sans domaine) : ça casse Kerberos et les certificats. Toujours un FQDN.
- **Cohérence DHCP/DNS :** le suffixe DNS distribué par DHCP (option 015) doit correspondre à la zone AD.

## 12. TTL : principes

Le TTL (Time To Live, en secondes) dit aux résolveurs combien de temps garder l'enregistrement en cache.

| Usage | TTL conseillé |
|---|---|
| Enregistrements d'infrastructure stables (NS, SOA, A de serveurs) | 1 heure à 1 jour (3600–86400) |
| Enregistrements SRV d'AD | 600 s (10 min, valeur Netlogon par défaut) |
| Enregistrements appelés à changer (bascule, migration) | 300 s (5 min) **avant** l'opération |
| Enregistrements DHCP dynamiques | 1200 s (20 min) par défaut |

