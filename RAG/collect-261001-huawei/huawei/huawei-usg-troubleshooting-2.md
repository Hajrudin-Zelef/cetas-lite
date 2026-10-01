---
id: collect-261001-huawei/huawei/huawei-usg-troubleshooting-2
title: "Huawei USG — Guide de dépannage (troubleshooting récurrent)"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr", "license", "memory"]
source: docs/RAG/collect-261001-huawei/huawei_usg_troubleshooting.md
source_anchor: ""
source_lines: [207, 437]
sha256: 9cda8f03f3e57ccb670601c8b0afd9da65cf71e5791f026930ab36673828f6cc
---

# Huawei USG — Guide de dépannage (troubleshooting récurrent)

```shell
display ipsec sa
display acl 3000        # vérifier le contenu de l'ACL
```

---

## 7. VPN monté mais pas de trafic

**Symptômes** : `display ipsec sa` montre des SA actifs, mais les LAN ne
se pinguent pas.

```shell
display ipsec sa                         # encaps/décapsulent ?
display firewall session table           # sessions visibles ?
```

**Causes** :
1. Security-policy : le trafic entre zones (trust→untrust ou vers la
   zone du tunnel) n'est pas autorisé. **Le VPN ne bypass pas la
   politique !**
2. Routage : il manque la route vers le LAN distant via le tunnel
   (en mode policy-based, le routage suit l'ACL ; en route-based,
   route statique vers l'interface Tunnel).
3. NAT qui « mange » le trafic VPN : exclure le trafic IPsec du NAT
   source (règle `nat-policy` avec `action no-nat` placée **avant**
   la règle de SNAT, ou ACL de NAT qui exclut les LAN distants).
4. MTU : les paquets ESP fragmentés sont parfois jetés → tester avec
   `ping -s 1400` puis ajuster (`mtu`, `tcp adjust-mss`).
5. Chevauchement d'adressage : les deux LAN utilisent le même
   sous-réseau → il faut du NAT avant chiffrement (rare, à éviter).

---

## 8. Lenteurs / CPU ou mémoire à 100 %

```shell
display cpu-usage
display memory-usage
display firewall session table | count   # nombre de sessions (approx)
display device
display logbuffer
```

**Causes fréquentes** :
- **Table de sessions saturée** (attaque ou P2P) : identifiez les
  tops parleurs (`display firewall session table` trié), bloquez en
  `firewall blacklist`, réduisez les timeouts.
- **UTM sur tout le trafic** : l'IPS/AV sur des flux lourds (sauvegardes,
  vidéo) coûte cher → excluez les flux de confiance des profils, ou
  passez l'IPS en détection seule d'abord.
- **Logs excessifs** : session-logging sur une règle très sollicitée =
  I/O disque/CPU → ne loguez que les règles sensibles.
- **Debugging oublié** : `undo debugging all` + `undo terminal debugging`
  (un debug resté actif plombe les performances).
- Boucle réseau en amont : broadcast storm → `display interface`
  (erreurs/overruns en forte hausse).
- Version logicielle buggée : comparez avec les release notes Huawei.

**Fix d'urgence** : `reset firewall session table` (casse les sessions
établies !) pour respirer, puis traitez la cause.

---

## 9. Problèmes DNS

**Symptômes** : `ping 8.8.8.8` OK mais aucun nom ne se résout.

```shell
nslookup www.google.com
display dns server          # selon version
display ip routing-table 8.8.8.8
```

- `dns resolve` oublié dans la config.
- Serveur DNS injoignable ou filtré par la security-policy
  (autoriser UDP 53 trust→untrust).
- Sur le LAN : le DHCP distribue-t-il les bons DNS ?
  (`display ip pool`, `dns-list`).
- DNS hijack de l'opérateur : testez avec `8.8.8.8` en direct.

---

## 10. Le DHCP ne distribue pas d'adresses

```shell
display ip pool name LAN_POOL
# (used/total : pool plein ?)
display interface GigabitEthernet 0/0/1   # dhcp select global présent ?
```

- `dhcp enable` oublié (erreur classique).
- Interface pas en `dhcp select global` (ou `interface`).
- Pool épuisé : élargissez le `network` ou réduisez le `lease`.
- Conflit : une IP statique dans la plage du pool.
- Le client est dans une autre VLAN/zone sans relay DHCP.

---

## 11. Impossible d'administrer le USG (web/SSH/ping)

```shell
# En console (toujours accessible) :
display interface brief
display zone
display security-policy rule all
```

- **Règle vers la zone `local` manquante** : c'est le cas n°1 après un
  reset ou une nouvelle config (voir guide CLI §6, règle ADMIN_TO_FW).
- Web : `display web-manager` — service désactivé ou port changé ?
  URL = `https://<ip>:8443`.
- SSH : `display ssh server status`, `display ssh user-information`.
  Protocole VTY restreint à telnet ? (`protocol inbound ssh`).
- Compte verrouillé après échecs : attendez ou débloquez en console.
- Vous êtes dans la zone `untrust` : l'admin depuis Internet est
  (heureusement) refusé par défaut → passez par VPN.

---

## 12. HA : pas de bascule / split-brain

```shell
display hrp state verbose
display hrp interface
display vrrp
```

- Lien heartbeat DOWN : câble, VLAN, ou IP `remote` fausse dans
  `hrp interface ... remote ...`.
- `hrp enable` oublié sur un des deux nœuds.
- Versions logicielles **différentes** : la synchro échoue.
- Les deux nœuds en `active` (split-brain) : heartbeat coupé des deux
  côtés → rétablissez le lien, un seul doit rester actif.
- Config non synchronisée : `hrp auto-sync config` puis vérifiez
  `display hrp state verbose` (state = `normal`).
- VGMP : si un groupe VRRP ne bascule pas avec les autres, vérifiez
  son rattachement au groupe VGMP.

---

## 13. OSPF / BGP qui ne s'établit pas

```shell
display ospf peer
display ospf interface
display bgp peer
display ip routing-table
```

- Interfaces pas dans la même zone avec politique autorisée : **OSPF/BGP
  entre zones passe par la security-policy** (trust→local / local→trust
  pour les paquets à destination du USG lui-même).
- `network` mal déclaré (wildcard inversé en OSPF).
- MTU mismatch (voisin reste en ExStart) : `display ospf peer verbose`.
- Authentification MD5 : mdp différent des deux côtés.
- BGP : `as-number` du peer faux, ou `ebgp` sans `peer ebgp-max-hop`
  sur plusieurs sauts.
- Router-ID dupliqué dans la zone OSPF.

---

## 14. L'UTM bloque à tort (faux positifs)

**Symptômes** : une appli métier cassée depuis l'activation IPS/AV/URL.

1. Identifiez le moteur : désactivez les profils **un par un** dans la
   règle (d'abord URL-filter, puis AV, puis IPS) jusqu'à trouver le
   coupable.
2. IPS : passez le profil en **alerte seule** d'abord, analysez les logs
   (`display logbuffer`), puis créez des exceptions pour les signatures
   fautives.
3. URL-filter : mettez le domaine métier en liste blanche.
4. AV : excluez les types de fichiers internes volumineux.
5. Vérifiez que la **base de signatures est à jour** : une base
   périmée = faux positifs + trous de sécurité (`display utm`,
   Update Center).
6. HTTPS : sans inspection SSL, l'UTM est aveugle sur le chiffré ;
   avec, un certificat mal déployé casse la navigation → vérifiez le
   CA sur les postes.

---

## 15. Licence / signatures expirées

```shell
display license
```

- Licence UTM/IPS expirée : les moteurs passent en mode dégradé ou
  s'arrêtent → renouvelez via le portail Huawei (fichier `.dat`,
  import en web UI).
- Après import : `reboot` parfois nécessaire selon le type de licence.
- Signatures non mises à jour depuis des mois : planifiez la mise à
  jour automatique (Update Center > Schedule).

---

## 16. Trafic asymétrique (le tueur silencieux)

**Symptômes** : sessions qui se montent puis tombent, applis qui
« rament », VPN instables.

- Le USG est **stateful** : si le retour ne repasse pas par lui
  (double chemin, deux FAI sans PBR), il jette les paquets.
- Diagnostic : `display firewall session table` montre des sessions
  à moitié établies ; `tracert` aller vs retour divergent.
- Fix : symétrisez les chemins (PBR, routage), ou en dernier recours
  désactivez la vérification d'état sur le flux concerné (à éviter en
  production).

---

## 17. Table de sessions pleine / attaque en cours

```shell
display firewall session table | include <ip-suspecte>
display firewall statistics
display logbuffer | include attack
```

- En cas de flood : `firewall blacklist <ip>` immédiatement.
- Vérifiez `Attack Defense` (SYN flood, UDP flood) : seuils trop hauts ?
- Rate-limitez en amont si le lien est saturé avant le USG.
- Après l'attaque : `reset firewall session table` pour nettoyer
  (coupure brève des sessions légitimes).

---

## 18. Échec d'upgrade / boot en boucle

