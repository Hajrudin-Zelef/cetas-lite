---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-17
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [2182, 2275]
sha256: b47f6eace30b5c49cc134ddb90f805e0a223f1bb1621acf5ccc0fd7b29682483
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

1. Quelle est la différence entre une résolution **récursive** et une résolution **itérative**, et qui fait quoi ?
2. Citez 3 enregistrements SRV qu'un contrôleur de domaine enregistre automatiquement, avec leur rôle.
3. Pourquoi ne faut-il **jamais** mettre un DNS externe (ex : 8.8.8.8) dans la carte réseau d'un DC ?
4. Expliquez le mécanisme no-refresh / refresh avec des intervalles de 7 jours : quand un enregistrement créé à J0 devient-il éligible au scavenging ?
5. Quelle est la différence entre une **zone stub** et un **redirecteur conditionnel** ?
6. Dans quel cas préférer les **indications de racine** aux **redirecteurs** ?
7. Pourquoi un **CNAME à l'apex** d'une zone est-il interdit ?
8. Que vérifie-t-on avec `Test-DnsServer -Context "Forwarder"`, et que signifie un échec ?
9. Quel est le risque de mettre un compte dans le groupe **DnsAdmins** sur un DC qui est aussi serveur DNS ?
10. Un enregistrement doit changer d'IP dans 48 h pour une migration : détaillez la procédure TTL correcte.

---

## 147. Réponses du quiz

1. **Récursive** : le client demande au résolveur et attend la réponse finale (le résolveur fait tout le travail). **Itérative** : le résolveur interroge successivement racine → TLD → serveur autoritaire, chacun ne donnant qu'une indication (referral) vers l'étape suivante.
2. `_ldap._tcp.dc._msdcs` (localiser un DC, port 389), `_kerberos._tcp.dc._msdcs` (Kerberos, port 88), `_gc._tcp` (catalogue global, port 3268) — plus les variantes `_sites` par site AD, `_kpasswd`, `_ldap._tcp.pdc._msdcs` (émulateur PDC).
3. Parce que le DC doit résoudre les enregistrements SRV de l'AD en priorité. Un DNS externe ne connaît pas la zone AD → échecs de localisation de DC, ouvertures de session lentes, réplication en erreur. L'externe passe par les **redirecteurs**, jamais par la carte réseau.
4. J0→J7 : no-refresh (les rafraîchissements sont ignorés) ; J7→J14 : refresh (le client peut mettre à jour son timestamp) ; après J14 sans rafraîchissement : l'enregistrement est **périmé** et sera supprimé au prochain cycle de scavenging.
5. La **zone stub** ne contient que les NS de la zone distante et suit automatiquement leurs changements ; le **redirecteur conditionnel** envoie les requêtes vers des IP fixes qu'il faut maintenir à la main.
6. Quand on veut l'indépendance vis-à-vis d'un tiers (souveraineté, confidentialité des requêtes), quand le pare-feu autorise tout le 53 sortant, ou pour un serveur exposé où l'on ne veut aucune dépendance externe.
7. Parce qu'un nom avec CNAME ne peut porter **aucun autre** enregistrement (RFC), or l'apex porte déjà le SOA et les NS : la zone deviendrait invalide.
8. On teste que chaque redirecteur configuré répond correctement. Un échec = le redirecteur est injoignable/lent → à remplacer ou à investiguer (pare-feu, panne FAI), car il ralentit toute la résolution Internet (timeouts en cascade).
9. Un membre de DnsAdmins peut charger une DLL arbitraire dans le service DNS (`ServerLevelPluginDll`) → exécution de code **SYSTEM** sur le DC → compromission complète du domaine. DnsAdmins ≈ Domain Admins sur un DC/DNS.
10. **J-2** : baisser le TTL à 300 s (5 min). **Jour J** : changer l'IP. **J+2** (après 2× l'ancien TTL) : remonter le TTL à sa valeur normale. Baisser le TTL *pendant* la migration ne sert à rien car les caches ont déjà l'ancienne valeur.

---

## 148. Glossaire

| Terme | Définition |
|---|---|
| AD-integrated (zone) | Zone stockée dans Active Directory, multi-maîtres, répliquée par AD |
| Aging (vieillissement) | Mécanisme d'horodatage des enregistrements dynamiques |
| Anycast | Plusieurs serveurs partageant une même IP annoncée en BGP |
| Apex (sommet de zone) | Le nom de la zone elle-même (ex : `contoso.local`) |
| Autoritaire | Serveur qui possède les données d'une zone |
| AXFR / IXFR | Transfert de zone complet / incrémental |
| CAA | Enregistrement déclarant les CA autorisées à émettre pour le domaine |
| Cache locking | Protection contre l'écrasement prématuré des entrées en cache |
| Cache négatif | Mise en cache des réponses NXDOMAIN |
| Conditional forwarder | Renvoi des requêtes d'une zone vers des serveurs désignés |
| DnsUpdateProxy | Groupe AD permettant aux DHCP d'enregistrer sans devenir propriétaires |
| DNSSEC | Signatures cryptographiques garantissant authenticité et intégrité |
| DS (Delegation Signer) | Enregistrement chez le parent ancrant la chaîne DNSSEC |
| EDNS | Extensions du DNS (réponses > 512 octets, bits DNSSEC…) |
| Forwarder (redirecteur) | Serveur vers qui on transfère les requêtes récursives |
| Glue record | A/AAAA d'un NS situé dans la zone qu'il délègue |
| KSK / ZSK | Clés DNSSEC : signature des clés / signature de la zone |
| _msdcs | Sous-zone contenant les enregistrements de localisation des DC |
| Notify | Notification poussée du primaire vers les secondaires après changement |
| NSEC / NSEC3 | Preuve d'inexistence signée (NSEC3 : hachée, anti-énumération) |
| NXDOMAIN | Réponse « le nom n'existe pas » |
| QNAME / QTYPE | Nom et type demandés dans une requête |
| RCODE | Code de réponse (NOERROR, NXDOMAIN, SERVFAIL, REFUSED…) |
| Récursion | Le serveur fait tout le travail de résolution pour le client |
| Referral | Indication « demande au serveur suivant » en résolution itérative |
| RRL | Response Rate Limiting : limite anti-amplification DDoS |
| Round-robin DNS | Plusieurs A/AAAA sur un nom, retournés en rotation |
| Scavenging | Suppression automatique des enregistrements dynamiques périmés |
| Scope (zone scope) | Vue d'une zone dans les stratégies DNS (split-brain) |
| Serial (SOA) | Numéro de version de la zone, déclencheur des transferts |
| Split-brain | Même zone, réponses différentes selon le client (interne/externe) |
| SRV | Enregistrement de localisation de service (priorité, poids, port) |
| Stub (zone) | Zone ne contenant que les NS d'une zone distante |
| TTL | Durée de vie en cache d'un enregistrement |
| Trust anchor | Clé de confiance de départ d'un validateur DNSSEC |
| Zone | Portion de l'espace de noms administrée par un serveur |

---

## 149. Pour aller plus loin

**Pratiquer en labo :**
- Montez une maquette : 2 DC/DNS (Server 2022/2025 Core), une zone AD-integrated, un DHCP avec compte dédié, activez le scavenging et observez les Events 2501.
- Simulez chaque cas pratique du chapitre 12 (cassez, diagnostiquez, réparez) : c'est le meilleur entraînement.
- Testez le split-brain par stratégies DNS avec 2 scopes, puis validez depuis 2 subnets.

**Approfondir :**
- **dcdiag /test:dns /v** et **repadmin** : le DNS et l'AD sont indissociables, maîtrisez les deux diagnostics.
- **DNS sur TLS (DoT, port 853) / DNS sur HTTPS (DoH)** : chiffrement du dernier kilomètre — natif côté client Windows 11/Server 2022+, à évaluer selon votre politique de confidentialité.
- **gMSA** (Group Managed Service Account) pour le compte DHCP-DNS : mot de passe géré automatiquement par AD, sans expiration manuelle.
- **IPAM** (IP Address Management, rôle Windows Server) : inventaire centralisé IP/DNS/DHCP pour les parcs importants.
- **BIND 9** en secondaire de vos zones : interopérabilité (transferts TSIG) pour les environnements mixtes.
- RFC de référence : 1034/1035 (bases), 2136 (updates dynamiques), 1996 (notify), 4033–4035 (DNSSEC), 6840, 7766 (DNS sur TCP).

**Runbooks à écrire pour votre équipe :**
1. « Ajout d'un serveur DNS au site X » (checklist §141 adaptée).
2. « Migration d'IP d'un serveur » (procédure TTL §12 + checklist §114).
3. « Rollover KSK DNSSEC » (§72).
4. « Le DNS ne répond plus » (arbre de décision basé sur §90–99).
5. « Audit trimestriel DNS » (§142, à planifier dans l'agenda d'équipe).

*Fin du guide — bon courage, et que vos serials soient cohérents.*
