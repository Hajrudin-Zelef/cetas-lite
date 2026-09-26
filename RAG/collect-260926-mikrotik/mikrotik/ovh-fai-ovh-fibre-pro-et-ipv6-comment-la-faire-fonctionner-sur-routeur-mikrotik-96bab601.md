---
id: collect-260926-mikrotik/mikrotik/ovh-fai-ovh-fibre-pro-et-ipv6-comment-la-faire-fonctionner-sur-routeur-mikrotik-96bab601
title: "OVH Fibre Pro et IPv6 : comment la faire fonctionner sur routeur Mikrotik ?"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/ovh-fai-ovh-fibre-pro-et-ipv6-comment-la-faire-fonctionner-sur-routeur-mikrotik-96bab601.md
source_anchor: ""
source_lines: [1, 69]
sha256: ea5f7e4fa8d131bf098b70a5c732e9c3aa4db3f22d4ecace46309b27a49c2a09
---

# OVH Fibre Pro et IPv6 : comment la faire fonctionner sur routeur Mikrotik ?

*Source : https://lafibre.info/ovh-fai/ovh-fibre-pro-et-ipv6-comment-la-faire-fonctionner-sur-routeur-mikrotik/*

**Question initiale :** J'utilise le routeur Mikrotik hEX (RB750Gr3) sous RouterOS 7.15.3. Connectivité IPv4 OK mais pas d'IPv6. L'IPv6 est activée dans OVH Manager avec un bloc /56. La section IPv6 du WebFig est configurée sans succès. Statut de la ligne OVH : « En cours de création », reverse DNS IPv4 et IPv6 en erreur 403.

**EDIT (résolu) :** la ligne est passée à l'état actif, ce qui a permis de faire fonctionner le reverse DNS et la connectivité IPv6. Guide ajouté ci-dessous.

---

**Réponse — piste IPCPv6 et DHCPv6-PD :**

D'après la doc OVH (KB0059164, si à jour), il faut faire de l'IPCPv6 : vérifier côté Mikrotik que c'est supporté (a priori oui, regarder dans le profil PPPoE si « use ipv6 » est activé). Selon le profil de la connexion (standard, Orange ou Bouygues), utiliser le bon VLAN. Et surtout mettre le bon MTU sur l'interface PPPoE.

Si on a une IPv6 link-local (fe80...) sur l'interface PPPoE, c'est que cette partie fonctionne.

Ensuite, utiliser le bloc /56 — tester si c'est du DHCPv6-PD : dans WebFig, IPv6 > DHCP client > Add new, choisir l'interface PPPoE vers OVH, cocher « prefix » dans la zone Request, nommer le pool au choix, laisser 64 ou mettre 56 dans « pool prefix length », cocher « Add Default Route », éventuellement décocher « rapid commit » (peut bugger les renouvellements selon l'équipement en face).

Si on obtient un pool, c'est bon. L'utiliser ensuite sur le LAN en ajoutant une adresse dans IPv6 > Addresses > Add new : choisir « from pool », éventuellement une partie basse manuelle (par défaut ::/64 génère une IP avec le préfixe + EUI64 ; avec « ::5/64 » ça combine le préfixe et ::5). Cocher « advertise » annonce le préfixe sur cette interface (typiquement le LAN) pour le SLAAC des équipements du réseau.

---

**Guide IPv4/IPv6 OVH Fibre Pro + Mikrotik — Profil Standard (aucun VLAN) :**

**PPP > Add New > PPPoE client**
- General, Name: pppoe-ovh, Max MTU: 1492, Interfaces: ether1, Dial Out
- User: your PPPoE ID @ovh.kosc
- Password: your PPPoE password
- Profile: default
- Use Peer DNS: no
- Add Default Route: yes

**PPP > Profiles > default**
- General : DHCPv6 PD Pool: ovh-pd ; DNS Server: none
- Protocols : Use IPv6: yes

**IP > Firewall > NAT > Add New**
- Out. Interface: pppoe-ovh — Action: masquerade

**IPv6 > DHCP client > Add New**
- Interface: pppoe-ovh
- Request: prefix coché (mais pas address)
- Pool Name: ovh-pd ; Pool Prefix Length: 64
- Use Peer DNS: no ; Rapid Commit: no ; Add Default Route: yes

**IPv6 Addresses > Add New**
- Address: ::/64, From Pool: ovh-pd, Interface: bridge, Advertise: yes

**IP > DNS**
- Servers: 2620:fe::fe, 2620:fe::9 (Quad9 — malware blocking, validation DNSSEC)
- Allow Remote Requests: yes

**Bridge > bridge** — MTU: 1492

**IPv6 > Firewall > Mangle**
- Chain: forward, Protocol: tcp, Advanced, TCP Flags: syn
- Action: Change MSS — New TCP MSS: clamp to pmtu — Passthrough: yes

---

**Discussion DNS :** Mettre deux fois un même serveur DNS en IPv4 et IPv6 n'est pas opportun. Si confiant dans l'IPv6, ne mettre que des serveurs IPv6 : 2606:4700:4700::1111 et 2606:4700:4700::1001 (le DNS joint en IPv6 répond bien pour les sites IPv4-only). Avec fallback IPv4 : primaire en IPv6 (2606:4700:4700::1111), secondaire en IPv4 (1.0.0.1).

**Remarque vie privée :** pour des guides/tutoriels, mieux vaut spécifier « Use Peer DNS: yes » plutôt que des DNS tiers en dur — les utilisateurs ne réalisent pas forcément qu'ils envoient tout leur trafic DNS à une entreprise tierce.

**Suivi :** recommandations DNS appliquées (résolveurs joints en IPv6 uniquement, pas de doublons, Quad9 — fondation suisse à but non lucratif — au lieu de Cloudflare, entreprise privée). La censure DNS demandée par l'État français aux FAI rend les résolveurs des FAI trop menteurs pour un usage exigeant. Règle mangle IPv6 ajoutée (nécessaire au bon fonctionnement).

**Débits :** speedtest à 864 Mbit/s en IPv4 contre 327 Mbit/s en IPv6 — signalement ARCEP effectué (blocage d'ICMPv6 par OVH avec impact sur PMTUd évoqué ; pas de support IPv6 FastTrack côté Mikrotik ?). Mikrotik contacté. Un routeur plus puissant (ex. RB5009UG+S+IN) pourrait aider.

**Retour d'expérience CPU :** vérifier la charge CPU du routeur pendant le speedtest IPv6. Exemple : un hAP² sur FTTH Milkywan tenait 1 Gb/s down / 500 Mb/s up en IPv4 mais plafonnait à 250-300 Mb/s en IPv6 avec un cœur CPU à 100 %. Après passage sur un Cisco 1100 d'occasion, plus de différence de débit IPv6 vs IPv4. Sur le hEX du fil : le CPU load (à 1 % au repos) monte à 49 % pendant le test IPv6 OVH — migration vers un routeur plus puissant à prévoir.
