---
id: collect-261001-rattrapage/rattrapage/canon-copieurs-guide-7
title: "Canon — Guide ultra-complet copieurs (maintenance au cœur)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/canon_copieurs_guide.md
source_anchor: ""
source_lines: [1000, 1184]
sha256: 9e2fa96dba781c760fcab55127a5791cf9d917d72865d00ab7d065ab84bbebf0
---

# Canon — Guide ultra-complet copieurs (maintenance au cœur)

Formation express de 15 minutes à chaque installation :
1. Charger le papier (ventiler, guides, pas trop remplir).
2. Sortir un bourrage (suivre l'écran, sens du papier, tout refermer).
3. Changer le toner (sans secouer comme un shaker).
4. Ce qu'il ne faut JAMAIS faire : agrafes dans l'ADF, papier humide,
   éteindre en arrachant la prise, toner compatible.
5. Qui appeler et quoi noter (code affiché à l'écran).

Un client formé = 50 % d'appels en moins = votre marge.

---

---

## 28. Le chemin papier en détail — chaque rouleau compte

> Comprendre le trajet exact du papier = localiser tout bourrage
> sans hésiter.

```
Cassette → [Pickup roller] → [Feed roller + Separation roller]
         → [Registration rollers] → [Transfert (tambour/ITB)]
         → [Séparation] → [Fixing (fusion)] → [Sortie]
                                                    ↓ (duplex)
                                              [Inverseur] → [Duplex]
                                                    ↓ (finisher)
                                              [Finisher]
```

### 28.1 Rôle de chaque rouleau
- **Pickup** : prend la feuille en haut de la pile (usure = 0108).
- **Feed** : avance la feuille (usure = retard 0101).
- **Separation** : retient les feuilles suivantes (usure = multi-prise).
- **Registration** : synchronise la feuille avec l'image (sale =
  décalage d'image, bourrage 0201).
- **Sortie/duplex** : éjecte ou réintroduit (usure = 0301/0A01).

### 28.2 Quand remplacer quoi (symptômes)
- Multi-prises → **separation** en priorité.
- Non-prise → **pickup** en priorité.
- Retards sans cause papier → **feed**.
- Toujours remplacer en **kit complet** (les 3) : un neuf + deux usés
  = déséquilibre de friction = nouveaux bourrages.

---

## 29. L'unité laser (LSU) — diagnostic

- Rôle : « dessine » l'image sur le tambour (étape 2 du §2).
- Symptômes : page blanche totale, image pâle uniforme, bandes
  blanches verticales (poussière sur la vitre du laser).
- **Nettoyage** : la vitre/lentille du laser s'encrasse (toner en
  suspension) → chiffon sec non pelucheux, jamais de liquide à
  l'intérieur.
- **Sécurité** : ne jamais alimenter le laser capot ouvert sans les
  protections (rayonnement classe 3B).
- Si le polygone (miroir rotatif) est bruyant → unité à remplacer
  (pas réparable en atelier).

---

## 30. Haute tension — charge, transfert, séparation

- La machine génère des **hautes tensions** (jusqu'à −6 kV pour la
  charge primaire) via la carte **HVT** (High Voltage Transformer).
- Symptômes :
  - Fond gris/noir → charge primaire trop faible.
  - Image pâle/effacée → transfert trop faible.
  - Papier qui colle au tambour → séparation trop faible.
- Diagnostic :
  1. Nettoyer les coronas/fils de charge (outil de nettoyage fourni,
     ou coton-tige + alcool).
  2. Vérifier les contacts (lames de contact pliées = classique après
     remplacement tambour).
  3. Mesurer les sorties HVT (multimètre HT ou selon manuel — **danger**,
     personnel qualifié uniquement).
- **Grille de charge (grid)** : encrassée = charge irrégulière =
  bandes. Nettoyage systématique à chaque remplacement tambour.

---

## 31. FAX — installation et pannes

- Ligne analogique RTC requise (les box « voix sur IP » posent
  problème : **désactiver l'ECM** si échecs, passer en 9 600 bps).
- Réglages : n° de fax, nom d'émetteur, mode de réception (auto/manuel),
  renvoi vers email/dossier (pratique : fax → PDF par mail).
- Pannes classiques :
  - **E674** → carte fax (réinstaller, vérifier).
  - Échec d'envoi → ligne (tester avec un téléphone), ECM, vitesse.
  - Réception noire → cartouche/toner (côté réception c'est une
    impression normale).
- En Afrique : le fax décline, mais les administrations et banques
  l'exigent encore → savoir le configurer reste utile.

---

## 32. Impression suiveuse et uniFLOW

- **Problème** : documents abandonnés sur l'imprimante (confidentialité,
  gaspillage).
- **Solution Canon : uniFLOW** — l'utilisateur s'authentifie (badge,
  code) et libère ses impressions sur n'importe quel copieur du parc.
- Avantages : confidentialité, quotas, rapports par utilisateur,
  suppression auto des travaux non libérés.
- En maintenance : savoir réenregistrer un copieur dans uniFLOW
  (adresse serveur, certificat) après remplacement de carte/HDD.

---

## 33. Dépannage réseau avancé

```
1. Le copieur a-t-il une IP ? → Panneau > Réseau > TCP/IP
2. Ping depuis un PC → OK ? sinon : câble, switch, VLAN
3. Page web embarquée (http://IP) accessible ? → si non : HTTP désactivé ou carte NIC
4. Port 9100 (RAW) / 515 (LPD) ouverts ? → pare-feu
5. File d'impression Windows : erreur ? → pilote, spouleur (redémarrer le spouleur)
6. Wireshark si besoin : le PC envoie-t-il au copieur ?
```

- **VLAN** : en entreprise, les copieurs sont souvent sur un VLAN
  dédié → le PC d'un autre VLAN ne les « voit » pas sans routage.
- **Wi-Fi** : les copieurs en Wi-Fi perdent la connexion (mode éco) →
  privilégiez le **câble**.
- **IPv6** : source de bizarreries → désactivez-le si non utilisé.

---

## 34. Gestion de parc (flotte de copieurs)

- **Outil Canon : e-Maintenance / imageWARE** — remontée automatique
  des compteurs, alertes toner, historiques. Proposez-le à vos clients
  multi-sites : vous êtes alerté **avant** la panne.
- **Standardisation** : un seul modèle par client si possible
  (stock de pièces unique, formation unique).
- **Rotation** : les machines à faible volume vieillissent mal
  (toner qui s'agglutine) → planifiez des impressions de test.
- **Fin de vie** : à 5–7 ans ou 1M de pages, le coût de maintenance
  dépasse le remplacement → proposez le renouvellement **avant**
  la panne fatale (argument chiffré avec l'historique).

---

## 35. Pièces détachées — références et compatibilités

- **Toujours commander par référence exacte** (ex. : `GPR-57` pour les
  toners — la référence change par série !).
- Compatibilités : un tambour de C5535 ne va pas sur un C5560 même si
  ça ressemble → vérifiez la **parts list** du manuel de service.
- **Pièces d'occasion** : acceptable pour cartes/cartes mères (testées),
  **jamais** pour tambour/développeur/fusion (usure inconnue).
- Fournisseurs : distributeur Canon agréé (garantie) vs grossistes
  (prix). En Afrique : anticipez **3–6 semaines** de délai → stock
  tampon obligatoire.

---

## 36. Questions pièges (pour valider vos connaissances)

1. Pourquoi ne faut-il pas secouer une cartouche de toner neuve ?
   → Poussière en suspension, risque d'inhalation, bourrage de la vis sans fin.
2. Un client a des bourrages 0108 uniquement le matin. Pourquoi ?
   → Humidité nocturne du papier (local non climatisé).
3. Après un remplacement de développeur, les copies sont trop foncées.
   → Initialisation STIR/TONER-S non faite.
4. E602 après un orage. Premier réflexe ?
   → CHK-HDD en mode service, puis remplacement + SST si échec.
5. Le scan SMB marchait hier, plus aujourd'hui, rien n'a changé sur le copieur.
   → Mot de passe Windows expiré ou mise à jour Windows (SMB).
6. Des points se répètent tous les 94 mm.
   → Tambour (circonférence).
7. Pourquoi reset les compteurs PARTS après remplacement ?
   → Sinon la machine applique les corrections d'usure de l'ancienne pièce.
8. Un copieur réformé part chez un ferrailleur. Qu'oubliez-vous ?
   → Effacement du disque dur (données clients).

---

---

## 37. Démontage — méthode et accès aux cartes

> 80 % des dégâts en maintenance viennent du démontage, pas de la panne.

