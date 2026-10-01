---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-25
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [2155, 2250]
sha256: 5ce519df875555868f3308b14569ad2fd49e21847abda88aa10a96f409d584f6
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

1. **Volume annuel** : consolide tes achats chez un seul Gold Partner — la remise suit.
2. **Stock tampon** : négocie un dépôt-vente ou un stock dédié pour tes clients critiques.
3. **RMA** : délais écrits, procédure express pour les sites Gold (tes clients).
4. **Formation** : places gratuites aux sessions eKit (argument pour ton équipe, section 158).
5. **Démo** : kit de démo (AR + switch + 2 AP) pour tes avant-ventes — un client qui voit l'app en action signe plus vite.
6. **Co-marketing** : cas clients communs (le distributeur adore, et ça te fait de la pub).
7. **Roadmap** : être prévenu des fins de vie et nouveautés (184).
8. **Paiement** : délais alignés sur tes propres encaissements clients — le cash-flow tue plus de prestataires que la technique.

## 202. TCO 3 ans : le vrai coût d'un site (méthode)

```
TCO 3 ans = Matériel + Câblage/pose + Onduleur/élec + Main-d'œuvre config
          + 3 × Maintenance annuelle + 3 × (énergie + 4G secours éventuel)
          + 1 × Renouvellement partiel provisionné (ex. : batteries onduleur an 3)
Exemple de structure (ordres de grandeur, prix à vérifier) :
  Boutique type C : matériel 1×, pose 0,5×, maintenance 0,2×/an
  → TCO 3 ans ≈ 2,1× le prix du matériel seul.
  Hôtel 40 ch. : câblage souvent = 40-60 % du matériel en rénovation.
```
**Leçon** : le matériel n'est que la moitié du coût. Un devis qui oublie la maintenance et l'énergie est un devis qui fait perdre de l'argent — au client d'abord, à toi ensuite (retours SAV non facturés).

## 203. Les 5 erreurs de chiffrage qui coûtent cher

1. **Oublier le câblage** : en rénovation, c'est 30-60 % du budget — toujours une visite technique (159).
2. **Sous-dimensionner le PoE** : le 2e switch acheté en urgence coûte plus cher que le bon switch du départ.
3. **Oublier l'onduleur** : « on verra plus tard » = première coupure = premier incident non couvert.
4. **Ne pas chiffrer la maintenance** : 3 ans sans contrat = 3 ans d'interventions gratuites qui tuent ta marge.
5. **Chiffrer au prix catalogue sans remise** : ou l'inverse — figer un prix client avant d'avoir le prix distributeur. Toujours : prix fournisseur écrit → marge → prix client.

## 204. Appels d'offres : les 5 pièges

1. **Le cahier des charges copié d'un concurrent** : s'il exige du Cisco/Meraki nommément, tu ne gagneras pas avec eKit — sauf à faire valoir l'équivalence fonctionnelle (rarement gagné, ne perds pas 3 jours dessus).
2. **Les critères flous** : « Wi-Fi performant » ne veut rien dire — demande des critères mesurables (débit, couverture, utilisateurs) ou propose les tiens.
3. **Le moins-disant** : si le seul critère est le prix, quelqu'un proposera toujours moins cher — avec moins de service. Ne brade pas ta maintenance pour gagner : tu paieras pendant 3 ans.
4. **Les délais irréalistes** : « installation en 1 semaine » pour un hôtel 40 chambres sans câblage existant = impossible proprement. Écris ton planning réel ; un délai tenu vaut mieux qu'une promesse intenable.
5. **L'absence de visite technique** : un AO sans visite = un piège (murs en béton armé découverts le jour J). Si la visite est impossible, mets des **hypothèses écrites** et des prix conditionnels.

## 205. Planning type d'un déploiement (modèle)

```
Semaine -2 : visite technique, devis signé, commande matériel
Semaine -1 : réception matériel, pré-staging (onboarding au bureau), plan final
Jour J-1   : câblage réceptionné (187), baie prête, onduleur en place
Jour J     : AM : tête (modem->AR->switch), PM : AP + config + tests (70, 74)
Jour J+1   : tests utilisateurs, formation client, PV de recette (160)
J+2 / J+7  : inspection cloud, ajustements radio, clôture
```
Adapte les durées (×3 pour un hôtel), mais garde la **structure** : jamais de config avant le câblage réceptionné, jamais de PV sans tests utilisateurs.

## 206. Double WAN et VPN site-à-site : l'architecture qui tient

```
        [Site A - Siège]                    [Site B - Agence]
   [Fibre Op1]   [4G secours]          [Fibre] (mono-lien suffit souvent)
        \           /                        |
     [USG6000F-S125]  <-- IPsec VPN -->  [AR280]
           |                                  |
     [LAN A + VLAN]                     [LAN B + VLAN]
```
- Le VPN IPsec entre USG/AR eKit : paramètres IKE à aligner (phase 1/2, PSK, sous-réseaux) — **une ligne différente = tunnel qui ne monte pas** (cas n°14).
- Chaque site garde son **adressage propre** (jamais le même sous-réseau des deux côtés !).
- En secours 4G : le VPN doit **remonter** sur le lien de secours — à tester explicitement (le failover WAN sans failover VPN = agence isolée).
- Alternative : VPN **hub-and-spoke** vers le siège pour 5+ agences (plus simple à administrer que le maillé complet).

## 207. Supervision SNMP pratique : par où commencer

- Les switchs eKit exposent **SNMPv1/v2c/v3** (fiches S220/S310/S620) : active **v3** (auth + priv) avec un utilisateur dédié au monitoring.
- Dans ton NMS (Zabbix/PRTG/LibreNMS) : ajoute chaque switch/AR en SNMP, supervise au minimum : **disponibilité (ping), état des interfaces, erreurs/compteurs, charge CPU/mémoire** (selon les MIB exposées — **à vérifier sur la documentation officielle** par modèle, ne promets pas une métrique sans l'avoir vue).
- **Ne scrape pas** toutes les 30 secondes un parc de 50 sites : 5 min d'intervalle suffit en PME (le temps réel, c'est le cloud eKit qui le fait).
- Complémentarité : eKit = alertes temps réel + config ; NMS = historique + tendances + preuve SLA (section 78-79).

## 208. eKit et la GTB/domotique : cohabiter intelligemment

- Climatisation, contrôle d'accès, alarmes : de plus en plus **IP** — donc sur **ton** réseau.
- Règle : **un VLAN par famille** (GTB, contrôle d'accès, alarme), avec des règles inter-VLAN strictes (la clim n'a pas besoin de voir les PC).
- Les équipements GTB sont souvent **fragiles** réseau (vieilles piles IP, pas de DHCP fiable) : IP fixes, documentées, et **ne jamais** les mettre sur un VLAN avec portail captif.
- Coordination : le chauffagiste/alarmeur n'est pas un administrateur réseau — **tu** gères le VLAN, **il** gère son application. Écris-le.
- PoE : beaucoup d'équipements GTB sont PoE — intègre-les dans ton budget (68).

## 209. Wi-Fi invités multi-sites : centraliser sans se tromper

- **Modèle unique** de SSID invités (nom, portail, quotas) appliqué à tous les sites : l'expérience client est identique partout, et tu ne gères qu'un modèle.
- **Attention** : les plages IP peuvent être identiques entre sites (LAN indépendants — section 57), mais les **noms de sites** dans les alertes doivent être explicites (sinon « AP hors ligne » ne veut rien dire).
- Portail : si tu veux une page **par enseigne** avec le logo de chaque boutique, vérifie la personnalisation par site (**à vérifier sur la documentation officielle**).
- **Données** : les logs de connexion des invités sont des données personnelles — mentions légales + durée de conservation par site/pays (163).

## 210. Tester avant de promettre : la maquette d'atelier

- **Investissement** : 1 AR180, 1 S220-8P4S, 2 AP361 (ton kit de démo, section 201) — amorti à la première avant-vente gagnée.
- **Usages** : pré-staging (19), formation équipe (158), reproduction de bugs avant d'aller sur site, test des MAJ (42), démo client.
- **Rituel** : chaque nouvelle version firmware passe 1 semaine sur la maquette avant le site pilote.
- Une maquette qui tourne en permanence dans ton atelier, c'est aussi une **vitrine** : le client qui la voit comprend immédiatement ce qu'il achète.

## 211. Recruter et faire grandir un technicien réseau PME

