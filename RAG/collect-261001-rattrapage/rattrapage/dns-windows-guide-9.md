---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-9
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [1082, 1218]
sha256: 890420a681210f2118aaec87f6986156adc8dd6919f818b471fa819f93cffaf6
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

- **KSK** (Key Signing Key, longue, ex : 2048 bits) : signe uniquement les clés. Sa *empreinte* (DS) est publiée chez le **parent**.
- **ZSK** (Zone Signing Key, plus courte, ex : 1280 bits) : signe les enregistrements de la zone. Rollover fréquent.
- **DS** (Delegation Signer) : l'enregistrement chez le parent qui ancre la confiance vers votre KSK.
- **DNSKEY / RRSIG / NSEC(3)** : les enregistrements techniques de la signature, générés automatiquement.

Le validateur part d'une **ancre de confiance** (trust anchor, ex : la clé de la racine, préinstallée) et descend la chaîne.

## 71. Signer une zone sur Windows Server

Méthode GUI (recommandée pour la première fois) : clic droit sur la zone → **DNSSEC → Signer la zone** → assistant (choix KSK/ZSK, algorithmes, NSEC3). Le serveur génère les clés, signe la zone, et gère le rollover.

En PowerShell :

```powershell
# 1. Ajouter les clés de signature (KSK 2048 bits + ZSK)
Add-DnsServerSigningKey -ZoneName "contoso.com" -CryptoAlgorithm RsaSha256 -KeyType KeySigningKey -KeyLength 2048
Add-DnsServerSigningKey -ZoneName "contoso.com" -CryptoAlgorithm RsaSha256 -KeyType ZoneSigningKey -KeyLength 1280

# 2. Paramétrer la signature (NSEC3 contre l'énumération de zone)
Set-DnsServerDnsSecZoneSetting -ZoneName "contoso.com" -SignWithNSEC3 $true -NSEC3HashAlgorithm SHA1 -NSEC3Iterations 10

# 3. Vérifier l'état de signature
Get-DnsServerDnsSecZoneSetting -ZoneName "contoso.com" | Format-List ZoneName, IsSigned, *
Get-DnsServerSigningKey -ZoneName "contoso.com" | Format-Table KeyId, KeyType, CryptoAlgorithm, IsActive
```

⚠️ **Prérequis** : zone **AD-integrated** (requis par Windows pour la signature), horloges synchronisées (la validation échoue si l'horloge dérive — les signatures ont une période de validité).

## 72. Gestion des clés et rollover

Le rollover (renouvellement) est le moment dangereux : une clé qui change sans coordination = **SERVFAIL** pour tous les validateurs.

| Clé | Rollover conseillé | Mécanisme Windows |
|---|---|---|
| ZSK | Tous les 90 jours | Automatique (pré-publication) |
| KSK | Tous les 1–2 ans | **Manuel** : générez la nouvelle KSK, publiez le nouveau **DS chez le parent/registrar**, attendez 2× TTL du DS, puis activez |

```powershell
# Rollover manuel de la KSK (après publication du DS chez le registrar)
Add-DnsServerSigningKey -ZoneName "contoso.com" -KeyType KeySigningKey -CryptoAlgorithm RsaSha256 -KeyLength 2048
# ... attendre la propagation du DS ...
# Puis retirer l'ancienne :
Remove-DnsServerSigningKey -ZoneName "contoso.com" -KeyId "<ancien KeyId>" -Force
```

**Checklist rollover KSK :** nouveau DS publié chez le registrar → `dig DS contoso.com` OK depuis l'extérieur → attendre 48 h → activer → surveiller les SERVFAIL → supprimer l'ancienne clé.

## 73. Ancres de confiance (trust anchors)

Côté **validateur** (votre résolveur interne qui vérifie les signatures), il faut les clés publiques de départ :

```powershell
# Importer l'ancre de la racine (pré-distribuée par Windows, à vérifier)
Get-DnsServerTrustAnchor -Name "." | Format-Table Name, TrustAnchorType

# Ajouter manuellement une ancre (ex : zone partenaire signée, en DS)
Add-DnsServerTrustAnchor -Name "partenaire.local" -Type DS -KeyTag 12345 `
    -CryptoAlgorithm RsaSha256 -DigestType Sha256 -Digest "ABCDEF0123456789..."

# Lister toutes les ancres
Get-DnsServerTrustAnchor | Format-Table Name, TrustAnchorType
```

Sans ancre pour une zone signée, le validateur la considère comme **non sécurisée** (pas d'erreur, mais pas de protection non plus) — sauf si vous forcez la validation.

## 74. Validation DNSSEC côté résolveur

Activer la validation sur vos résolveurs internes : toute réponse d'une zone signée est vérifiée ; signature invalide → **SERVFAIL** (plutôt qu'une fausse réponse).

```powershell
# GUI : Propriétés du serveur → onglet Avancé → "Activer la validation DNSSEC pour les réponses distantes"
# Vérifier l'état :
Get-DnsServerRecursion | Select-Object Enable, EnableDnsSecValidation
```

Testez :
```powershell
# Doit réussir (zone signée valide)
Resolve-DnsName -Name "dnssec-tools.org" -Type A -DnssecOk -Server "10.0.0.10"
# Test de validation négative (zone volontairement mal signée)
Resolve-DnsName -Name "www.dnssec-failed.org" -Type A -DnssecOk -Server "10.0.0.10"  # → SERVFAIL attendu
```

⚠️ Avant d'activer la validation en production : vérifiez que vos **redirecteurs** valident ou au moins transmettent les bits DNSSEC (sinon, vos validations échouent à cause du redirecteur, pas de la zone).

## 75. NSEC vs NSEC3

Pour prouver la **non-existence** d'un nom (réponse NXDOMAIN signée), deux mécanismes :

| | NSEC | NSEC3 |
|---|---|---|
| Principe | Chaîne les noms existants en clair | Chaîne des **hachés** des noms |
| Risque | Permet l'**énumération de zone** (zone walking : lister tous les noms) | Protège contre l'énumération (haché + sel + itérations) |
| Coût | Léger | Un peu plus lourd (calculs) |

**Recommandation : NSEC3** systématiquement pour les zones d'entreprise (vos noms de serveurs internes sont de l'information). Windows le propose dans l'assistant de signature et via `-SignWithNSEC3 $true` (§71).

## 76. Limites et points d'attention DNSSEC

1. **Complexité opérationnelle** : rollover KSK, DS chez le registrar, horloges → prévoyez des runbooks.
2. **Taille des réponses** : les réponses signées sont grosses → bascule UDP→TCP fréquente → ouvrez **TCP 53** partout (pare-feu !).
3. **Dernier kilomètre non protégé** : entre le poste client et votre résolveur, pas de DNSSEC (le bit DO est positionné par le résolveur). Pour chiffrer ce segment : DNS sur TLS/HTTPS (voir §149).
4. **Dépannage** : un SERVFAIL DNSSEC ressemble à une panne DNS classique → ayez le réflexe `Resolve-DnsName -DnssecOk` et vérifiez la chaîne (cas n°11, §110).
5. **Zones AD internes** : la signature des zones internes est possible mais rarement prioritaire — le risque principal en interne est ailleurs (cf. chapitre 13).

---

# Chapitre 10 — DNS, Active Directory et DHCP

## 77. Mises à jour dynamiques : principe

Au lieu de créer chaque enregistrement à la main, les machines **s'enregistrent elles-mêmes** (RFC 2136) : au démarrage, à chaque renouvellement DHCP, toutes les 24 h. Le client envoie un UPDATE signé (Kerberos en mode sécurisé) au serveur DNS autoritaire.

```powershell
# Voir le mode de mise à jour dynamique d'une zone
Get-DnsServerZone -Name "contoso.local" | Select-Object ZoneName, DynamicUpdate
# Valeurs : None | NonsecureAndSecure | Secure
```

## 78. Secure vs non-secure vs none

| Mode | Authentification | Usage |
|---|---|---|
| `None` | — | Zones statiques (DMZ exposée, zone publique) |
| `NonsecureAndSecure` | Optionnelle | ⚠️ À éviter : n'importe qui peut écraser n'importe quel enregistrement |
| `Secure` | Kerberos obligatoire, ACL par enregistrement | **Standard AD** : seul le propriétaire (compte machine) peut modifier son enregistrement |

**Règle :** toute zone AD-integrated = `Secure`. Le mode non-sécurisé n'a de sens que pour des zones hébergeant des équipements incapables de Kerberos **et** isolées — et encore, préférez des enregistrements statiques.

## 79. DNS et DHCP : qui enregistre quoi

Deux modèles (propriétés du serveur DHCP → onglet **DNS**) :

**Modèle A — Le client s'enregistre lui-même** (défaut Windows) : le client DHCP met à jour son **A** (nom → IP) ; le serveur DHCP met à jour le **PTR** (IP → nom). Simple, pas de compte dédié.

**Modèle B — Le DHCP enregistre pour le client** (« Toujours mettre à jour dynamiquement… » coché) : le DHCP fait A + PTR pour tous les clients, y compris non-Windows (imprimantes, téléphones, Linux sans nsupdate). **Requiert un compte dédié** (§80).

