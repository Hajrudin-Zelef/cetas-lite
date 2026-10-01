---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-17
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [3411, 3496]
sha256: 553d4947670542742ffe974245c0f68005b092e76de5adc1968d9108f2964191
---

# Guide ultra-complet — Huawei eKit S310

## 126. Quiz : 10 questions pour valider (réponses en section 127)

**Q1.** Quelle est la différence entre un S310-24T4S et un S310-24P4S ?
**Q2.** Quel est le budget PoE total d'un S310-24P4S ? Combien d'AP761 (17,7 W)
peut-il alimenter au maximum ?
**Q3.** Pourquoi ne faut-il jamais laisser le mot de passe d'usine, et où
stocke-t-on les mots de passe admin ?
**Q4.** Un port hybrid avec `pvid vlan 10`, `tagged vlan 20`, `untagged vlan 10` :
que devient une trame non taggée qui entre ? Une trame taggée VLAN 20 qui sort ?
**Q5.** Sur un trunk, quel oubli classique explique que « le VLAN 20 ne passe pas »
alors que le VLAN 10 passe ?
**Q6.** Cite les 3 protections à mettre sur chaque port utilisateur contre les
boucles accidentelles.
**Q7.** Après avoir activé DHCP snooping, plus personne n'a d'IP. Quelle est la
cause la plus probable ?
**Q8.** À quoi sert la table de binding du DHCP snooping, et quelles fonctions
s'appuient dessus ?
**Q9.** Pourquoi faut-il faire `save` après chaque session de configuration ?
**Q10.** Ton uplink 1G est saturé à 100 % en continu. Cite 2 solutions
structurelles (pas la QoS).

---

## 127. Réponses du quiz

**R1.** Le **T** = cuivre sans PoE ; le **P** = cuivre **avec PoE+** (802.3at,
budget 400 W). Le reste (24 ports GE + 4 SFP) est identique.
**R2.** **400 W**. 400 ÷ 17,7 = 22,6 → **22 AP761 maximum** (389,4 W). Le 23e ne
démarrera pas.
**R3.** Parce qu'un mot de passe connu/publié = accès admin à quiconque est sur
le réseau. Mots de passe **uniques, forts, au coffre d'entreprise**, jamais en
clair ni sur post-it.
**R4.** La trame non taggée entrante est marquée **VLAN 10** (PVID). La trame
VLAN 20 sortante sort **taggée** (et une trame VLAN 10 sortante sortirait non
taggée).
**R5.** Le VLAN 20 a été **oublié dans `port trunk allow-pass vlan ...`**
d'un côté (ou des deux). Vérifier avec `display port vlan` des deux côtés.
**R6.** **Edge port** (`stp edged-port enable`) + **BPDU guard**
(`stp bpdu-protection`) + **storm control**. (RSTP actif globalement, évidemment.)
**R7.** Le port vers le **serveur DHCP** (ou le relais) n'est pas en
`dhcp snooping trusted` — sur **chaque** switch traversé en cascade.
**R8.** Elle associe (IP, MAC, port, VLAN, bail) de chaque client légitime.
**DAI** (vérification des ARP) et **IPSG** (filtrage du trafic source) s'appuient
dessus. C'est aussi un inventaire vivant du réseau.
**R9.** Parce que la configuration active vit en **RAM** : sans `save`, elle est
**perdue au reboot**. `display current-configuration` ≠ `display
saved-configuration` tant qu'on n'a pas sauvegardé.
**R10.** Passer l'uplink en **10G** (modèle X / SFP+) ou monter un **Eth-Trunk
LACP** (2× 1G). Et traiter la cause (sauvegardes en heures creuses, caméras...).

**Score :** 9–10 : tu peux former les autres. 7–8 : solide, relis tes points
faibles. < 7 : reprends les sections correspondantes avant de toucher à la prod.

---

## 128. Pour aller plus loin : ressources officielles

- **Fiche produit officielle (eKit) :**
  `https://ekit.huawei.com/ekit/front/ssr/en/product/1239139538459211136`
- **Datasheet série eKitEngine S310 (PDF, spécifications complètes) :**
  `https://i.pcdeacitec.com/datasheets/syscom/S310-24U4X/f2112277465062598bb47bdc547534d7.pdf`
- **Datasheet miroir :**
  `https://resource.ewe.rs/media/documents/2026/08/2026-08-07573726.pdf`
- **Datasheet S310-24T4S :**
  `https://shop.aodatacloud.es/medias/datasheet_s310-24t4s.pdf`
- **Support entreprise Huawei** (firmwares, guides de configuration par version) :
  via ton revendeur / compte support — les images logicielles ne se téléchargent
  pas sans compte.
- **Communauté / forum Huawei** : utile pour les cas tordus déjà rencontrés
  par d'autres (toujours recouper avec la doc officielle).

> ⚠️ **Avertissement versions :** ce guide décrit des commandes **syntaxiquement
> plausibles pour VRP Huawei** et des comportements standard de la gamme eKit
> S310. La syntaxe exacte de certaines fonctions (MQC/QoS avancée, DAI/IPSG,
> MUX VLAN, 802.1X, ports de stack, `startup system-software`, VCT) peut varier
> selon la **version logicielle** : devant toute commande marquée « à vérifier
> sur la version du modèle exact », tape `?` en contexte et consulte le guide
> de configuration de ta version avant la production. En cas de doute, la
> maquette (un switch de test) tranche toujours.

---

*Fin du guide — Huawei eKit S310. Version 1.0 — septembre 2026.*
*Prochaines étapes suggérées : adapter les exemples à ton plan d'adressage réel,
imprimer la checklist (118) et le pense-bête (124), et planifier la première
visite de maintenance (117).*
