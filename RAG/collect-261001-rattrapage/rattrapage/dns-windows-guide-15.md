---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-15
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [1913, 2040]
sha256: c3d34a38ea41bb2b43bfd9fa8c8b58c6fc47df31926abe431c08de636443a842
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

```powershell
# Règle pare-feu : autoriser le DNS entrant (si vous gérez le pare-feu Windows)
New-NetFirewallRule -DisplayName "DNS entrant UDP/TCP 53" -Direction Inbound `
    -Protocol UDP -LocalPort 53 -Action Allow -Profile Domain
New-NetFirewallRule -DisplayName "DNS entrant TCP 53" -Direction Inbound `
    -Protocol TCP -LocalPort 53 -Action Allow -Profile Domain
```

⚠️ Le piège classique : on ouvre l'UDP 53 mais pas le **TCP 53** → les transferts de zone et les réponses DNSSEC échouent mystérieusement.

---

# Chapitre 14 — 18 erreurs classiques

## 123. Erreur n°1 : le serveur DNS ne se pointe pas lui-même en premier

**Symptôme :** lenteurs d'ouverture de session, erreurs Netlogon, le DC « ne se trouve pas lui-même ».
**Cause :** la carte du DC pointe vers un autre DNS (ou un externe) en préféré.
**Correction :** lui-même (`127.0.0.1`) en premier, le partenaire en second (§20). À vérifier **après chaque changement d'IP**.

## 124. Erreur n°2 : l'îlot DNS (DNS island)

**Symptôme :** après la promotion d'un nouveau DC/DNS, il ne réplique pas et ne trouve pas les autres DC.
**Cause :** le nouveau serveur se pointe **lui-même en premier** alors que sa zone est encore vide → il s'isole (il ne connaît que lui).
**Correction :** pendant la promotion, pointez le futur DC vers un DNS **existant** en préféré ; ne le faites se pointer lui-même qu'**après** la réplication initiale. Puis remettez la config standard (§20).

## 125. Erreur n°3 : scavenging activé avec des intervalles incohérents

**Symptôme :** des enregistrements valides disparaissent (voir cas n°5, §104).
**Cause :** `no-refresh + refresh` < bail DHCP ou < absences typiques.
**Correction :** alignez les trois (bail DHCP ≤ no-refresh + refresh, scavenging ≥ 7 j), et un seul serveur scavenger (§56).

## 126. Erreur n°4 : des enregistrements « statiques » qui ne le sont pas

**Symptôme :** l'imprimante du service compta disparaît du DNS tous les mois.
**Cause :** enregistrement créé par DHCP (dynamique, timestampé) pour un équipement passé ensuite en IP fixe sans repasser l'enregistrement en statique.
**Correction :** recréez l'enregistrement à la main (timestamp à 0, §84) ou décochez « Supprimer cet enregistrement lorsqu'il devient périmé ».

## 127. Erreur n°5 : un CNAME au sommet de zone (apex)

**Symptôme :** la zone ne charge plus, ou les MX du domaine ne fonctionnent plus.
**Cause :** un CNAME à l'apex cohabite avec le SOA et les NS — interdit par les RFC, le serveur Windows le refuse ou produit une zone invalide.
**Correction :** utilisez un A à l'apex (quitte à le maintenir via script/API), jamais un CNAME (§37).

## 128. Erreur n°6 : un TTL d'une heure sur un enregistrement critique avant migration

**Symptôme :** le jour de la bascule, la moitié des clients va encore sur l'ancienne IP pendant des heures.
**Cause :** TTL trop long + pas de baisse préalable.
**Correction :** 48 h avant : TTL à 300 s ; jour J : changez l'IP ; J+2 : remontez le TTL (§12).

## 129. Erreur n°7 : des redirecteurs lents ou morts

**Symptôme :** toute la résolution Internet est lente par intermittence.
**Cause :** redirecteur injoignable (timeout 3 s à chaque requête froide) ou surchargé.
**Correction :** `Test-DnsServer -Context Forwarder`, remplacez par des résolveurs rapides et redondants, réglez `-Timeout` (§60).

## 130. Erreur n°8 : la zone `_msdcs` supprimée ou corrompue

**Symptôme :** plus aucun client ne trouve de DC ; `dcdiag /test:dns` catastrophique.
**Cause :** suppression « de ménage » de la zone `_msdcs` (elle paraît redondante avec `contoso.local`).
**Correction :** recréez la zone AD-integrated `_msdcs.contoso.local` (portée **Forest**), redémarrez Netlogon sur chaque DC puis `nltest /dsregdns` partout. Et ne recommencez jamais (§7).

## 131. Erreur n°9 : DHCP sans compte de service dédié

**Symptômes :** après remplacement d'un serveur DHCP, les mises à jour DNS échouent (« accès refusé »), doublons d'enregistrements.
**Cause :** les enregistrements appartiennent à l'ancien compte machine du DHCP.
**Correction :** compte de service dédié + groupe DnsUpdateProxy **avant** la mise en production (§80).

## 132. Erreur n°10 : transferts de zone ouverts à « tout serveur »

**Symptôme :** aucun — c'est silencieux. Un `dig AXFR contoso.local @votre-DNS` liste toute votre infrastructure (noms de serveurs, parfois leur rôle).
**Cause :** liste blanche non configurée.
**Correction :** liste blanche d'IP uniquement (§50). Testez depuis l'extérieur après chaque changement.

## 133. Erreur n°11 : mises à jour dynamiques non sécurisées sur une zone AD

**Symptôme :** aucun — jusqu'au jour où un poste vérolé écrase l'enregistrement d'un serveur.
**Cause :** zone en `NonsecureAndSecure` « pour que les imprimantes puissent s'enregistrer ».
**Correction :** repassez en `Secure` ; les équipements qui ne font pas Kerberos passent par le DHCP (modèle B, §79) ou des enregistrements statiques.

## 134. Erreur n°12 : pas de zone inversée

**Symptôme :** `nslookup 10.1.0.50` répond « *** can't find », logs illisibles, applications qui vérifient le reverse en échec.
**Cause :** zone `in-addr.arpa` jamais créée.
**Correction :** une zone inversée par sous-réseau, mises à jour sécurisées, PTR via DHCP (§32, §40).

## 135. Erreur n°13 : un seul serveur DNS distribué par DHCP

**Symptôme :** quand `SRV-DNS-01` est en maintenance, tout le site est aveugle.
**Cause :** option DHCP 006 avec une seule IP.
**Correction :** toujours **deux** DNS dans le DHCP (et dans les configs statiques), sur deux serveurs physiques/VM distincts (§85).

## 136. Erreur n°14 : confondre zone stub et redirecteur conditionnel

**Symptôme :** après un changement de NS chez le partenaire, la résolution casse (ou au contraire on s'étonne qu'elle suive toute seule).
**Cause :** on attendait le comportement de l'autre mécanisme.
**Correction :** stub = suit les NS distants automatiquement ; conditional forwarder = IP fixes, à maintenir (§26, §61). Choisissez en conscience.

## 137. Erreur n°15 : ne jamais vider ni tester le cache

**Symptôme :** « on a corrigé l'enregistrement mais ça ne marche toujours pas » pendant des heures.
**Cause :** cache négatif (NXDOMAIN) ou cache serveur périmé ; personne ne pense à `Clear-DnsServerCache`.
**Correction :** réflexe `-DnsOnly` vs `-CacheOnly` pour diagnostiquer (§92), vidage ciblé après correction (§95).

## 138. Erreur n°16 : TTL trop long pendant une migration (variante)

**Symptôme :** rollback impossible proprement : l'ancienne IP reste en cache partout.
**Cause :** on a remonté le TTL trop tôt, ou on n'a jamais baissé celui des enregistrements SRV d'AD (600 s par défaut — correct, mais vérifiez).
**Correction :** plan de TTL écrit dans le dossier de changement, avec les heures de baisse/remontée (§12).

## 139. Erreur n°17 : un DC dans le groupe DnsUpdateProxy

**Symptôme :** aucun — jusqu'à l'audit de sécurité (ou l'incident).
**Cause :** « pour que ça marche » on y a mis le compte machine du DC/DHCP.
**Correction :** retirez-le immédiatement : ses enregistrements deviennent non sécurisés et modifiables par tout membre du groupe (§81).

## 140. Erreur n°18 : récursion ouverte sur un serveur exposé

**Symptôme :** votre DNS sert d'amplificateur dans un DDoS ; votre FAI vous appelle.
**Cause :** serveur autoritaire public avec récursion activée pour tout le monde.
**Correction :** `Set-DnsServerRecursion -Enable $false` sur l'exposé, ou recursion scopes limités au LAN (§117), + RRL (§118).

---

# Chapitre 15 — Exploitation

## 141. Checklist de mise en production

