---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-16
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [2338, 2423]
sha256: 147cc565234be03ef1a488e5acafb6b3980499b3ef151e82b4357506808c203a
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

- **Symptômes** : le SYN part, le SYN-ACK ne revient pas (ou l'inverse) ; applis qui « se connectent puis figent ».
- **Diagnostic** : `display firewall session table` → session **incomplète** (un seul sens) ; vérifier les **routes retour** (`display ip routing-table` des deux côtés).
- **Causes** : route de retour qui **ne repasse pas par l'USG** (le firewall ne voit qu'un sens → droppe, normal pour un stateful) ; **deux chemins** (lien principal + secours) avec routage asymétrique.
- **Solution** : corriger le routage pour la **symétrie** ; en dernier recours, ajuster les timeouts/options de session (à éviter — masque le problème).

## 147. Cas 4 : le VPN IPSec ne monte pas (phase 1)

- **Symptômes** : `display ike sa` vide.
- **Diagnostic par étapes** :
  1. `ping` entre les IP publiques → si KO : réseau ou politique local (ICMP).
  2. Politique **untrust→local** : UDP 500, UDP 4500, protocole ESP autorisés ?
  3. `debugging ike packet` → regarder où ça bloque : pas de réponse du tout (réseau/ACL) ou **payload mismatch** (propositions IKE différentes).
  4. Vérifier : **PSK identique** (la faute de frappe n°1 — la re-saisir des deux côtés), **proposals** compatibles, **remote-address** correcte, **heure** synchronisée.
- **Solution** : corriger le paramètre divergent, `reset ike sa` / `reset ipsec sa` pour forcer la renégociation proprement.

## 148. Cas 5 : phase 1 OK, phase 2 KO (« le tunnel monte mais rien ne passe »)

- **Symptômes** : `display ike sa` = RD, mais `display ipsec sa` vide ou sans trafic.
- **Diagnostic** : comparer les **ACL** des deux côtés (miroir exact ?), les **transform-set** (algorithmes compatibles ?), le **PFS** (activé d'un seul côté = échec).
- **Autres causes** : **NAT** qui attrape le trafic VPN (exemption `no-nat` manquante ou mal ordonnée — section 52) ; **politique de sécurité** qui bloque le trafic clair dans le tunnel.
- **Solution** : aligner ACL/transform-set/PFS, placer l'exemption NAT **avant** la règle NAT, autoriser le trafic dans la politique.

## 149. Cas 6 : tunnel IPSec instable (flapping)

- **Symptômes** : le tunnel tombe et remonte toutes les X minutes/heures.
- **Diagnostic** : corréler les coupures avec les **lifetimes** (renégociation qui échoue), vérifier le **DPD** (actif ?), regarder les logs des deux côtés à l'heure exacte.
- **Causes fréquentes** : DPD non configuré (pair « fantôme ») ; **double NAT** devant un site (NAT-T mal géré) ; lifetimes incompatibles ; **instabilité du lien** (perte de paquets — IKE n'aime pas).
- **Solution** : activer le DPD (section 62), aligner les lifetimes, stabiliser le lien (QoS, changement de FAI si récurrent).

## 150. Cas 7 : le portail SSL VPN ne s'ouvre pas

- **Symptômes** : `https://<ip-publique>` ne répond pas / timeout.
- **Diagnostic** :
  1. La **passerelle SSL VPN** est-elle démarrée ? (`display sslvpn gateway`)
  2. Politique **untrust→local** : port 443 (ou port configuré) autorisé ?
  3. Le port est-il en **conflit** (le web d'admin utilise aussi le 443 ? → changer l'un des deux) ?
  4. Test depuis l'intérieur (bypass la politique untrust) pour isoler.
- **Solution** : démarrer la passerelle, corriger la politique, changer le port (ex. 4433) et **documenter**.

## 151. Cas 8 : SSL VPN connecté mais aucun accès réseau

- **Symptômes** : login OK, portail OK, mais les ressources internes injoignables.
- **Diagnostic** : depuis quelle **zone** arrive le trafic VPN (souvent une zone dédiée ou trust) ? La **politique** autorise-t-elle zone-VPN→trust/dmz ? Le client a-t-il reçu une **IP du pool** (conflit avec le LAN ?) ? Les **routes** côté client (split tunneling : les routes sont-elles poussées ?).
- **Solution** : créer les politiques manquantes, vérifier le pool d'adresses (pas de chevauchement), vérifier le routage client.

## 152. Cas 9 : débit effondré depuis l'activation de l'UTM

- **Symptômes** : tout marchait, depuis l'activation AV/IPS c'est lent (ou ça timeout).
- **Diagnostic** : `display cpu` (un cœur à 100 % ?), `display memory`, comparer **avant/après** désactivation temporaire d'un profil (en heures creuses, avec accord).
- **Causes** : modèle **sous-dimensionné** pour l'UTM (section 73 — le cas classique) ; profil appliqué sur des règles à **gros volume** (backup, vidéosurveillance) ; **inspection SSL** sur tout le trafic.
- **Solution** : restreindre les profils aux règles utiles (matrice section 84), exclure les flux volumineux non à risque, **upgrader de modèle** si structurel (c'est un dimensionnement, pas un bug).

## 153. Cas 10 : l'IPS/AV bloque une application métier (faux positif)

- **Symptômes** : une appli interne/outil métier bloqué depuis une MAJ de signatures.
- **Diagnostic** : logs UTM → **signature ID** en cause ; confirmer que c'est bien un faux positif (l'appli est légitime, le comportement est normal).
- **Solution** : **exception ciblée** (désactiver la signature pour cette source/destination, ou whitelist), **jamais** désactiver tout le profil. Documenter (section 86), **revue trimestrielle**. Remonter le faux positif à Huawei si récurrent.

## 154. Cas 11 : session asymétrique en HA après bascule

- **Symptômes** : après bascule HA, certaines connexions tombent (surtout les longues : SSH, RDP, visio).
- **Diagnostic** : comparer `display firewall session table` sur les deux nœuds **avant** bascule (tables quasi identiques ?) ; `display hrp sync-status`.
- **Causes** : **synchro HRP** incomplète ou désactivée ; câblage **asymétrique** ; bascule pendant un pic (synchro en retard).
- **Solution** : activer/vérifier la synchro auto (section 102), corriger le câblage, re-tester la bascule (section 101).

## 155. Cas 12 : « Internet coupé » — le NAT ne traduit plus

- **Symptômes** : plus aucun accès Internet depuis le LAN, le reste (local, DMZ) OK.
- **Diagnostic** : `display nat-policy statistics` (compteurs à 0 ?), l'**IP de l'interface WAN** a-t-elle changé (DHCP/PPPoE du FAI ?) — avec Easy IP c'est transparent, mais vérifier ; le **pool NAT** est-il épuisé (`display nat address-group`) ?
- **Causes fréquentes** : FAI qui a changé d'IP et **route par défaut** tombée ; pool NAT **saturé** (trop petit pour le nombre de sessions) ; règle NAT **désactivée** par erreur.
- **Solution** : vérifier IP/route WAN, agrandir le pool ou passer en PAT, réactiver la règle.

## 156. Cas 13 : DNS qui ne résout plus à travers le firewall

- **Symptômes** : « Internet ne marche pas » mais les IP passent (ping 8.8.8.8 OK, google.com KO).
- **Diagnostic** : `nslookup` depuis un poste → timeout ? Politique : le **DNS (UDP/TCP 53)** est-il autorisé trust→untrust (section 35, règle 1 — l'oubli classique) ?
- **Solution** : ajouter la règle DNS. 💡 **Toujours** tester un accès Internet **par IP et par nom** pour distinguer « réseau » de « DNS ».

## 157. Cas 14 : FTP/SIP qui « marche à moitié »

- **Symptômes** : FTP : connexion OK, `ls` OK, transfert KO. SIP : enregistrement OK, pas d'audio.
- **Diagnostic** : ALG activés ? (`display firewall alg`) ; politique : les **ports de données** (FTP passif : plage du serveur ; SIP/RTP : UDP 10000-20000) sont-ils autorisés ?
- **Solution** : activer l'ALG (section 53), ouvrir la plage de ports de données (restreinte aux IP des serveurs), préférer le **mode passif** FTP.

## 158. Cas 15 : boucle / tempête après un mauvais câblage

