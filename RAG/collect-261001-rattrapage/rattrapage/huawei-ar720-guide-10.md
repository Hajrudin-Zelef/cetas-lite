---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-10
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [1690, 1895]
sha256: 54ce1b9993a1db4c98ad75764ff8162261c07af4d922a14f87cc3b3b52626682
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

Sur la gamme AR, certaines fonctions avancées (contrôleur WLAN avec beaucoup d'AP,
fonctions de sécurité avancées type IPS/AV avec signatures, parfois le nombre de
tunnels) peuvent nécessiter des licences ou des abonnements de signatures — **à vérifier
sur la fiche du modèle exact** et le devis du fournisseur.

Procédure type (quand applicable) :

```
display license
```

- Noter les dates d'expiration des licences/signatures dans le planning de maintenance.
- Une signature IPS/AV expirée = protection qui ne se met plus à jour (silencieusement).
- Renouveler **avant** expiration : prévoir 30 jours de marge dans le suivi.

## 88. Mot de passe perdu : procédure de récupération

Si plus personne ne connaît le mot de passe admin :

1. Accès physique + console obligatoires.
2. Au boot, interrompre la séquence (touche selon version, souvent `Ctrl+B`) pour entrer
   dans le BootROM.
3. Choisir l'option « skip current configuration » / ignorer la config au démarrage
   (le libellé exact dépend de la version — **à vérifier sur la fiche du modèle exact**).
4. Le routeur démarre sans config : reconfigurer les accès, puis recharger la config
   sauvegardée (d'où l'importance de la sauvegarde externe, section 83).
5. `save` avec les nouveaux mots de passe.

C'est pour ça que l'accès physique au local technique doit être contrôlé : quiconque a
la console peut réinitialiser les mots de passe.

---
---

# Partie K — Dépannage terrain (15 cas détaillés)

> Méthode générale : **1)** qualifier le symptôme (quoi, quand, qui est impacté),
> **2)** vérifier la couche basse (LED, lien, IP), **3)** remonter les couches
> (routage → NAT → firewall → applicatif), **4)** une hypothèse à la fois,
> **5)** noter la cause et la solution dans le dossier du site.

## 89. Cas n°1 — Le tunnel IPSec ne monte pas : phase 1 KO

**Symptômes :** `display ike sa` vide, aucun trafic vers le site distant.

**Diagnostic :**

```
display ike sa
display ike peer
ping <ip-publique-distante>
display security-policy all | include IKE
```

**Causes fréquentes → solutions :**

| Cause | Indice | Solution |
|---|---|---|
| Pas de connectivité IP | ping KO | Régler le WAN d'abord (Partie B) |
| UDP 500/4500 bloqués | ping OK, IKE muet | Ouvrir dans la politique untrust→local (section 68) + box FAI en face |
| PSK différente | logs IKE « authentication failed » | Ressaisir la même PSK des 2 côtés |
| Proposal incompatible | échec de négociation | Aligner chiffrement/hash/DH/ike-version |
| `ike-version` différente (v1 vs v2) | pas de SA du tout | Mettre la même version des 2 côtés |
| remote-address erronée | — | Vérifier l'IP publique distante (elle a peut-être changé en DHCP) |

## 90. Cas n°2 — Phase 1 OK, phase 2 KO (pas de SA IPSec)

**Symptômes :** `display ike sa` montre une SA établie, mais `display ipsec sa` est vide.

**Diagnostic :** `display ipsec sa`, `display acl 3100` des deux côtés.

**Causes → solutions :**

- **Crypto ACL non miroir** : l'ACL locale dit 192.168.10.0→192.168.0.0 mais la distante
  dit autre chose. Réécrire les deux ACL en miroir exact (sections 55–56).
- **Transform-set incompatible** (AES-256 d'un côté, 3DES de l'autre) : aligner.
- **PFS différent** (dh-group14 vs none) : aligner.
- **Politique appliquée sur la mauvaise interface** : `display ipsec policy` et vérifier
  qu'elle est bien sur l'interface WAN active.

## 91. Cas n°3 — Tunnel monté mais aucun trafic ne passe

**Symptômes :** SA IPSec établies, compteurs à 0 ou qui n'augmentent pas, ping KO.

**Diagnostic :**

```
display ipsec sa
display ip routing-table | include 192.168.0.0
display nat session table
display firewall session table
```

**Causes → solutions :**

1. **Exclusion NAT oubliée** (le grand classique) : le trafic vers le site distant est
   NATé avant d'entrer dans le tunnel. Ajouter la règle `deny` dans l'ACL du `nat
   outbound` (section 40).
2. **Route manquante** vers le réseau distant : ajouter la route statique ou vérifier
   OSPF (`display ospf peer`).
3. **Politique de sécurité** : autoriser trust→untrust et le retour (les flux IPSec
   eux-mêmes sont gérés, mais le trafic clair GRE/interne doit être autorisé selon la
   version).
4. **MTU** : paquets qui passent dans un sens mais pas l'autre → régler `tcp adjust-mss`
   sur le tunnel (section 58).

## 92. Cas n°4 — Le PPPoE décroche en boucle

**Symptômes :** Internet coupe et revient toutes les quelques minutes ; `display
pppoe-client session summary` montre des sessions qui se renégocient.

**Diagnostic :**

```
display pppoe-client session summary
display logbuffer | include PPP
display interface Dialer 1
```

**Causes → solutions :**

- Ligne FAI instable (marge SNR faible) : faire tester la ligne par le FAI, vérifier le
  câblage cuivre, les filtres, l'ONT.
- Négociation PAP/CHAP : forcer le bon mode si le FAI est strict.
- `dialer-group` / `dialer-rule` mal configurée : le Dialer ne redémarre pas tout seul.
  Vérifier `dialer bundle 1` et la règle `permit`.
- MTU : tester avec `ping -s 1472 -c 5` (ne pas confondre perte MTU et décrochage).

## 93. Cas n°5 — Le NAT ne translate pas (pas d'Internet malgré WAN OK)

**Symptômes :** le routeur ping Internet, les PC du LAN non.

**Diagnostic :**

```
display nat outbound
display acl 2000
display ip interface brief
```

**Causes → solutions :**

1. `nat outbound` oublié sur l'interface WAN (ou sur la **bonne** interface : Dialer et
   pas GE0/0/0 quand on est en PPPoE !).
2. Le réseau LAN oublié dans l'ACL 2000 (ajouter la règle `permit`).
3. Le `nat outbound` est sur l'interface physique alors que l'IP est sur le Dialer
   (PPPoE) : le NAT doit être sur l'interface qui porte l'IP publique.
4. Politique de sécurité trust→untrust manquante (section 66).

## 94. Cas n°6 — OSPF neighbor bloqué (ExStart, Exchange, Loading)

**Symptômes :** `display ospf peer` montre un voisin qui ne dépasse pas un état
intermédiaire.

**Diagnostic :** `display ospf peer verbose`, `display ospf error`.

**Causes → solutions :**

| État bloqué | Cause typique | Solution |
|---|---|---|
| Init | Hello non reçu en retour | ACL/firewall qui bloque le multicast, `silent-interface` d'un côté |
| ExStart | MTU différent | Aligner la MTU des deux interfaces |
| Exchange/Loading | Router-ID dupliqué | Changer le router-id (unique !) |
| — | Authentification MD5 différente | Même clé, même mode des 2 côtés |
| — | Timers Hello/Dead différents | Aligner |
| — | Area différente | Même area des 2 côtés |

## 95. Cas n°7 — Débit asymétrique ou très inférieur au contrat FAI

**Symptômes :** speedtest à 40 Mbit/s sur un lien 200 Mbit/s, ou montant correct mais
descendant effondré.

**Diagnostic :**

```
display interface GigabitEthernet 0/0/0
display cpu-usage
display traffic-policy statistics interface Dialer 1 outbound
```

**Causes → solutions :**

- **Duplex/négociation** : `display interface` montre des erreurs CRC, collisions →
  forcer `duplex full` + `speed 1000` des deux côtés (ou laisser l'auto-négociation si
  elle converge).
- **CPU à 100 %** : trop de services (IPS + VPN + QoS) → alléger ou accepter la limite
  plateforme.
- **CAR trop agressif** : vérifier les traffic-policy appliquées.
- **Câble** : un câble abîmé négocie parfois en 100 Mbit/s sans erreur franche →
  changer le câble, tester.
- Tester **au routeur** d'abord (`ping` + transfert depuis le routeur) pour isoler :
  problème WAN ou problème LAN ?

## 96. Cas n°8 — DHCP qui ne distribue plus

**Symptômes :** les nouveaux PC restent en 169.254.x.x.

**Diagnostic :**

```
display ip pool name lan-bureautique used
display dhcp server statistics
```

**Causes → solutions :**

