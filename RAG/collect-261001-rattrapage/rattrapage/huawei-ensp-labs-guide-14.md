---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-14
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [2039, 2176]
sha256: a8475e419ee06b2df85d323ab035078201db02dc68636f42ce44910917ce049d
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

```
# Si après les réparations 1-4 certains flux restent bloqués de façon "globale"...
[R1]display current-configuration | include traffic-filter
# traffic-filter inbound acl 3000 sur GE0/0/2 !
[R1]display acl 3000
# rule deny → tout est bloqué en entrée de l'interface.
[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]undo traffic-filter inbound
# Re-test complet.
```

**Validation finale :** matrice de pings complète au tableau, tout au vert.

### Vérifications (checklist de fin de TP)

```
[SW1/SW2]display port vlan          # VLAN et trunks symétriques
[R1/R2]display ospf peer brief      # Full des deux côtés
[R1/R2]display ip routing-table protocol ospf
[R2]display nat session all          # sessions pendant un ping Internet
[R1]display current-configuration | include traffic-filter   # aucune ACL parasite
```

### Pièges classiques (côté stagiaire)

1. **Corriger plusieurs choses à la fois** sans re-tester entre chaque → on ne sait plus ce qui a réparé quoi. Discipline : une correction = un test.
2. **Sauter les couches basses** : chercher du côté d'OSPF alors qu'un lien est down. Toujours commencer par `display interface brief`.
3. **Le "ça marchait hier"** : en dépannage réel, la première question est "qu'est-ce qui a changé ?" (`display current-configuration`, historique des `save`).
4. **Confondre symptôme et cause** : "Internet est coupé" peut être du NAT (panne 4), du routage (panne 3) ou une ACL (panne 5) → la méthode par couches départage.
5. **Négliger la documentation** : chaque panne doit donner lieu à une fiche (symptôme → diagnostic → correction → validation). C'est noté.

### Barème indicatif (30 points — TP noté)

| Critère | Points |
|---|---|
| Méthode suivie et documentée (6 étapes appliquées) | 6 |
| Panne 1 (VLAN) trouvée et corrigée | 4 |
| Panne 2 (trunk asymétrique) trouvée et corrigée | 5 |
| Panne 3 (Router-ID dupliqué) trouvée et corrigée | 5 |
| Panne 4 (NAT/ACL) trouvée et corrigée | 5 |
| Panne 5 (ACL parasite) trouvée et corrigée | 5 |

### Durée estimée
**1 h 30 à 2 h** (dont 15 min de débrief collectif).

### Fiche animateur — points à insister
- Ce TP est **le plus important** du parcours : en production, on dépanne plus qu'on ne configure. La méthode compte plus que la vitesse.
- Imposer l'écrit : fiche par panne. Un dépanneur qui ne documente pas fait perdre à l'équipe.
- Débrief collectif en fin de séance : chaque groupe présente UNE panne (symptôme → raisonnement → correction). C'est là que l'apprentissage se consolide.
- Adapter la difficulté : pour un groupe débutant, donner le nombre de pannes (5) et leurs couches ; pour un groupe avancé, ne rien dire ("le réseau est en panne, débrouillez-vous").
- Lien terrain : cette méthode (couches 1→2→3→fonctions) est exactement celle à appliquer sur un ticket réel, que ce soit un S310, un AR720 ou un USG6000.
- Insister sur `display` avant `config` : **on ne configure jamais pour diagnostiquer, on observe d'abord**.

---

## 15. Quiz final (10 questions corrigées)

> À faire passer en fin de parcours (30 minutes, sans notes). Barème : 2 points par question, total 20.

**Q1. Sur un switch Huawei, un port en `port link-type trunk` avec `port trunk allow-pass vlan 10 20` : que se passe-t-il pour une trame du VLAN 30 qui arrive sur ce port ?**

<details>
<summary>Réponse</summary>

Elle est **rejetée (filtrée)** : le trunk ne laisse passer que les VLAN explicitement autorisés par `allow-pass`. Par défaut, seul le VLAN 1 passe. C'est le mécanisme du "VLAN qui ne passe pas" (TP2).
</details>

**Q2. Vous configurez une sous-interface `GigabitEthernet0/0/1.10` avec `dot1q termination vid 10` et une adresse IP, mais les PC du VLAN 10 ne pingent pas la passerelle. Quelle commande manque très probablement ?**

<details>
<summary>Réponse</summary>

`arp broadcast enable` : sans elle, le routeur VRP ne traite pas les requêtes ARP entrantes sur la sous-interface (TP4). Vérification : `display interface GigabitEthernet0/0/1.10`.
</details>

**Q3. En OSPF, deux routeurs ont la même area (0), des timers identiques, mais l'adjacence reste Down. Citez deux autres causes possibles.**

<details>
<summary>Réponse</summary>

- **Router-ID dupliqué** (ou 0.0.0.0) : deux routeurs avec le même Router-ID ne peuvent pas devenir voisins (TP12, panne 3).
- **Masques/MTU incohérents** sur le lien, ou **sous-réseaux différents** des deux côtés du lien, ou une ACL filtrant les Hello (multicast 224.0.0.5).
</details>

**Q4. Quelle est la différence entre `preference 60` et `preference 100` sur deux routes statiques vers la même destination ?**

<details>
<summary>Réponse</summary>

La route avec la **préférence la plus basse** est préférée : celle à 60 est active, celle à 100 est en backup (floating static). Si la route à 60 disparaît (interface down), celle à 100 prend automatiquement le relais (TP5).
</details>

**Q5. Dans un tunnel IPSec policy-based, à quoi sert l'ACL (ex. `acl 3000`) et pourquoi doit-elle être "miroir" sur les deux routeurs ?**

<details>
<summary>Réponse</summary>

L'ACL définit l'**interesting traffic** : seul le trafic qui la matche est chiffré dans le tunnel. Elle doit être miroir (source/destination inversées) car chaque routeur la lit dans son propre sens : R1 protège `10.0→20.0`, R2 doit protéger `20.0→10.0`. Sinon la phase 2 (proxy-ID) ne négocie pas (TP9).
</details>

**Q6. Sur un USG, vous créez une règle `trust → untrust` en `action permit`. Faut-il aussi créer la règle retour `untrust → trust` pour que la navigation web fonctionne ? Pourquoi ?**

<details>
<summary>Réponse</summary>

**Non** : le firewall est **stateful** — il suit les sessions et autorise automatiquement le trafic retour d'une session initiée depuis trust. Une règle untrust→trust ne serait nécessaire que pour un flux **initié** depuis l'extérieur (ex. serveur exposé, TP11).
</details>

**Q7. Un AP en mode Fit reste à l'état `idle` dans `display ap all`. Citez deux vérifications à faire, dans l'ordre.**

<details>
<summary>Réponse</summary>

1. **Connectivité IP** entre l'AP et l'AC (ping dans le VLAN de gestion) : sans IP, pas de tunnel CAPWAP (TP10).
2. **Déclaration de l'AP** sur l'AC (bonne MAC / bon SN, AP rattaché à un groupe avec un VAP diffusé). Si l'AP est en `configFailed`, vérifier les noms de profils référencés.
</details>

**Q8. Expliquez la séquence DHCP DORA et dites à quelle étape le relais DHCP intervient.**

<details>
<summary>Réponse</summary>

**D**iscover (broadcast client) → **O**ffer (proposition du serveur) → **R**equest (le client choisit) → **A**CK (confirmation). Le **relais** intervient dès le Discover : le routeur du site distant intercepte le broadcast, le convertit en unicast vers le serveur central en ajoutant son adresse (`giaddr`), ce qui permet au serveur de choisir le bon pool (TP4).
</details>

**Q9. Vous activez RSTP sur trois switches en triangle sans fixer les priorités. Qui devient root bridge, et pourquoi est-ce un problème en production ?**

<details>
<summary>Réponse</summary>

Le switch avec le **Bridge ID le plus petit** (priorité 32768 partout par défaut, donc la **plus petite adresse MAC**). Problème : le root peut être un switch d'accès en bout de chaîne, mal placé et peu puissant → chemins sous-optimaux et convergence fragile. En production, on fixe toujours `stp root primary/secondary` sur les switches de cœur (TP3).
</details>

**Q10. Un utilisateur se plaint : "J'ai une adresse IP mais pas d'Internet". Citez, dans l'ordre de la méthode de diagnostic, les 4 vérifications à faire sur son poste puis sur le réseau.**

<details>
<summary>Réponse</summary>

