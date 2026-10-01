---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-1
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [1, 57]
sha256: 4bfc4c1377b77270e3ae01ffb3d8bf5e8017b1964566dedfed22b6ca533b18ce
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

> **Public :** administrateurs systèmes & réseaux, chefs de service.
> **Versions couvertes :** Windows Server 2019, 2022 et 2025.
> **Convention :** tous les exemples utilisent le domaine fictif `contoso.local` et des adresses RFC 1918 (`10.0.0.0/8`). Aucune donnée réelle.
> **Avertissement :** certaines applets PowerShell ou options GUI varient légèrement selon la version — les différences sont signalées par `⚠️ Version`.

---

## Sommaire

- Chapitre 1 — Fondamentaux DNS et DNS dans Active Directory (§1–§12)
- Chapitre 2 — Installation du rôle DNS et prise en main (§13–§22)
- Chapitre 3 — Les zones DNS (§23–§34)
- Chapitre 4 — Les enregistrements de ressources (§35–§46)
- Chapitre 5 — Transferts de zone AXFR/IXFR (§47–§51)
- Chapitre 6 — Vieillissement (aging) et scavenging (§52–§56)
- Chapitre 7 — Redirecteurs et indications de racine (§57–§62)
- Chapitre 8 — Split-brain DNS et stratégies DNS (§63–§68)
- Chapitre 9 — DNSSEC (§69–§76)
- Chapitre 10 — DNS, Active Directory et DHCP (§77–§84)
- Chapitre 11 — Haute disponibilité (§85–§89)
- Chapitre 12 — Dépannage (§90–§99) et 15 cas pratiques (§100–§114)
- Chapitre 13 — Sécurité (§115–§122)
- Chapitre 14 — 18 erreurs classiques (§123–§140)
- Chapitre 15 — Exploitation : checklists, sauvegarde, supervision (§141–§144)
- Pense-bête des commandes (§145) · Quiz (§146–§147) · Glossaire (§148) · Pour aller plus loin (§149)

**Détail des sections :**

1. Objectif de ce guide et périmètre · 2. Rappels : qu'est-ce que le DNS ? · 3. Résolution récursive vs itérative · 4. Hiérarchie DNS : racine, TLD, domaines · 5. Le rôle du DNS dans Active Directory · 6. Les enregistrements SRV d'AD en détail · 7. La zone `_msdcs` · 8. Sites AD et SRV par site · 9. Enregistrements A des contrôleurs de domaine · 10. Comment un client localise un DC (flux complet) · 11. DNS et stratégie de nommage · 12. TTL : principes · 13. Installation du rôle DNS (GUI) · 14. Installation via PowerShell et Server Core · 15. La console DNS (dnsmgmt.msc) · 16. Le module PowerShell DnsServer : tour d'horizon · 17. dnscmd.exe : l'outil historique · 18. Vérifier l'état du service DNS · 19. Configurer les adresses d'écoute · 20. Configurer la carte réseau du serveur DNS lui-même · 21. Checklist d'installation · 22. Premier test de résolution · 23. Les types de zones : panorama · 24. Zones primaires standard · 25. Zones secondaires · 26. Zones stub · 27. Zones de recherche directe vs inversée · 28. Zones intégrées à Active Directory · 29. Portées de réplication AD · 30. Avantages et limites des zones AD-integrated · 31. Créer une zone primaire (GUI + PowerShell) · 32. Créer une zone inversée · 33. Convertir une zone standard en zone AD-integrated · 34. Suspendre, recharger, supprimer une zone · 35. Anatomie d'un enregistrement · 36. A et AAAA · 37. CNAME : alias et pièges · 38. MX : messagerie · 39. SRV : services · 40. PTR : résolution inverse · 41. TXT : usages · 42. NS et glue records · 43. SOA : le début d'autorité · 44. CAA : autorité de certification · 45. Les enregistrements SRV d'AD : liste complète · 46. Créer des enregistrements en PowerShell : synthèse · 47. AXFR vs IXFR · 48. Configurer le transfert de zone (GUI) · 49. Configurer en PowerShell / dnscmd · 50. Sécuriser les transferts · 51. Notify, SOA serial et cohérence · 52. Le problème des enregistrements périmés · 53. No-refresh et refresh intervals : le mécanisme · 54. Activer le vieillissement sur une zone · 55. Activer le scavenging au niveau serveur · 56. Bonnes pratiques et pièges du scavenging · 57. Redirecteurs : principe · 58. Indications de racine : principe · 59. Redirecteurs vs racine : quand utiliser quoi · 60. Configurer les redirecteurs · 61. Redirecteurs conditionnels · 62. Chaînage et boucles : à éviter · 63. Le problème du split-brain · 64. Solutions classiques au split-brain · 65. Split-brain avec les stratégies DNS (zone scopes) · 66. Stratégies DNS : client subnets · 67. Stratégies DNS : exemples concrets · 68. Bonnes pratiques split-brain · 69. DNSSEC : pourquoi · 70. Chaîne de confiance : KSK, ZSK, DS · 71. Signer une zone sur Windows Server · 72. Gestion des clés et rollover · 73. Ancres de confiance (trust anchors) · 74. Validation DNSSEC côté résolveur · 75. NSEC vs NSEC3 · 76. Limites et points d'attention DNSSEC · 77. Mises à jour dynamiques : principe · 78. Secure vs non-secure vs none · 79. DNS et DHCP : qui enregistre quoi · 80. Configurer les credentials DHCP · 81. Le groupe DnsUpdateProxy : utilité et danger · 82. Protéger les enregistrements existants · 83. Scavenging et DHCP : le duo gagnant · 84. Cas des clients non-Windows et statiques · 85. Plusieurs serveurs DNS : architecture type · 86. AD-integrated : la HA naturelle · 87. Secondaires et transferts pour la HA · 88. Anycast DNS : principe · 89. Supervision de la disponibilité DNS · 90. Méthodologie de dépannage DNS · 91. nslookup : bases et pièges · 92. Resolve-DnsName : le successeur · 93. dnscmd pour le diagnostic · 94. Test-DnsServer · 95. Vider le cache : client et serveur · 96. Les journaux DNS : lequel regarder · 97. Journal analytique DNS · 98. Debug logging (dnscmd /LogLevel) · 99. Event IDs à connaître · 100–114. 15 cas pratiques commentés · 115. Surface d'attaque du DNS · 116. Cache poisoning et protections · 117. Restreindre la récursion · 118. Response Rate Limiting (RRL) · 119. Cache locking · 120. Le groupe DnsAdmins : risque d'escalade · 121. Durcissement : liste de mesures · 122. DNS et pare-feu : flux réseau · 123–140. 18 erreurs classiques · 141. Checklist de mise en production · 142. Checklist d'audit trimestriel · 143. Sauvegarde et restauration des zones · 144. Supervision : compteurs et alertes · 145. Pense-bête des commandes · 146. Quiz : 10 questions · 147. Réponses du quiz · 148. Glossaire · 149. Pour aller plus loin

---

# Chapitre 1 — Fondamentaux DNS et DNS dans Active Directory

## 1. Objectif de ce guide et périmètre

Ce guide couvre le DNS **tel qu'il tourne en entreprise sous Windows Server 2019/2022/2025**, avec un focus sur l'intégration Active Directory — parce qu'en pratique, 90 % des pannes DNS en entreprise sont des pannes d'AD déguisées (localisation de DC, réplication, scavenging agressif).
Chaque notion est accompagnée de commandes PowerShell testées syntaxiquement, de tableaux de référence et de cas réels.
Hypothèses : forêt `contoso.local`, deux sites AD (`Paris-Siege`, `Lyon-Usine`), serveurs `SRV-DNS-01`/`SRV-DNS-02` en `10.0.0.10`/`10.0.0.11`.

## 2. Rappels : qu'est-ce que le DNS ?

Le DNS (Domain Name System) est une base de données distribuée et hiérarchique qui associe des noms à des informations (adresses IP, serveurs de messagerie, services…).
Points clés à ne jamais perdre de vue en exploitation :

| Concept | Ce qu'il faut retenir |
|---|---|
| Hiérarchie | Racine `.` → TLD (`.fr`, `.local`) → domaines → sous-domaines. La délégation se fait par enregistrements NS. |
| Protocole | UDP 53 par défaut, TCP 53 pour les réponses > 512 octets (ou avec EDNS > 4096), les transferts de zone et DNSSEC. |
| Modèle | Autoritaire (le serveur *possède* la zone) vs récursif/résolveur (le serveur *cherche* pour le client). Un serveur Windows fait souvent les deux. |
| Cohérence | Pas de transaction globale : la donnée se propage par transferts de zone et réplication AD, avec un délai. |
| TTL | Chaque enregistrement a une durée de vie en cache. Baisser le TTL *avant* un changement, pas pendant. |

## 3. Résolution récursive vs itérative

