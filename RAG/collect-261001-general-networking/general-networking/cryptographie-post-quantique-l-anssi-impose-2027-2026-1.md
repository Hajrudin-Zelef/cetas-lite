---
id: collect-261001-general-networking/general-networking/cryptographie-post-quantique-l-anssi-impose-2027-2026-1
title: "OpenSSL 3.5+ intègre nativement ML-KEM, ML-DSA et SLH-DSA"
domain: general-networking
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["acquisition", "arr", "mai", "research"]
source: docs/RAG/collect-261001-general-networking/cryptographie-post-quantique-l-anssi-impose-2027-2026.md
source_anchor: ""
source_lines: [1, 41]
sha256: 1088bada7b6c339a44b256e4a8bf7e90f3e9d910b0273a71a4d77dc6eb36508e
---

# OpenSSL 3.5+ intègre nativement ML-KEM, ML-DSA et SLH-DSA

L’Agence nationale de la sécurité des systèmes d’information (ANSSI) a franchi un cap que peu de régulateurs osaient encore : elle ne certifiera plus les produits de sécurité dépourvus de **cryptographie post-quantique**. L’annonce, faite à la mi-juin 2026 lors de la conférence France Quantum 2026 à Station F (Paris), transforme une recommandation prudente en obligation de marché. À partir de 2027, un produit qui ne résiste pas à un futur **ordinateur quantique** perdra son sésame le plus précieux : la qualification ANSSI, indispensable pour vendre à l’État et aux opérateurs d’infrastructures critiques.

Derrière le calendrier – arrêt des certifications en 2027, achats exclusivement « quantum-safe » d’ici 2030 – se joue une bascule industrielle majeure que plusieurs cabinets d’analystes commencent à chiffrer, avec des fourchettes qui varient selon la méthodologie mais convergent toutes sur une croissance à deux chiffres. Selon Juniper Research, dont le rapport de janvier 2026 a été révisé à la hausse, les investissements mondiaux dans la cryptographie post-quantique devraient désormais atteindre 13,3 milliards de dollars sur la seule année 2026 ; MarketsandMarkets l’évaluait pour sa part à 0,42 milliard de dollars en octobre 2025, avec un taux de croissance annuel de 46,2 % projeté jusqu’en 2030, tandis que Kaiso Research le situait à 470,2 millions de dollars en juin 2025, en route vers 26,15 milliards de dollars d’ici 2035. Les éditeurs français et européens de VPN, de HSM, de cartes à puce, de messageries chiffrées ou de solutions cloud souveraines doivent désormais réécrire leur cryptographie avant que le tampon de l’agence ne devienne inaccessible. Cet article décortique la décision, ses échéances, les algorithmes concernés et l’onde de choc qu’elle provoque déjà sur le marché européen de la **certification ANSSI**.

## Ce que l’ANSSI a réellement annoncé à France Quantum 2026

Le 16 juin 2026, sur la scène de la conférence France Quantum 2026, Samih Souissi, chef de cabinet de l’ANSSI, a confirmé une orientation que l’agence préparait depuis des années : à compter de 2027, l’ANSSI cessera progressivement de délivrer ses certifications et qualifications aux produits de sécurité qui n’intègrent pas de **chiffrement post-quantique**. Le message aux entreprises est sans ambiguïté : d’ici 2030, elles ne devraient plus acheter que des solutions résistantes au calcul quantique.

La portée est stratégique. La qualification ANSSI – la déclinaison française la plus exigeante de la certification – conditionne l’accès au marché des administrations et des opérateurs critiques. Perdre son éligibilité, c’est perdre l’accès à l’un des plus grands marchés technologiques publics d’Europe. Selon les précisions relayées par la presse spécialisée, l’agence recommande fortement des mécanismes *hybrides* (algorithme classique + algorithme post-quantique) pour toute solution censée protéger des données au-delà de 2030, cette exigence s’étendant aussi bien à l’échange de clés qu’aux signatures numériques. Le média français IT Social et le site spécialisé The Quantum Insider ont tous deux confirmé l’échéance de 2027.

Fait notable, l’ANSSI ne part pas d’une page blanche. L’agence avait publié dès 2022 un premier avis sur la transition vers la **cryptographie post-quantique**, suivi fin 2023 d’une position plus ferme annonçant qu’elle cesserait de délivrer certains labels de sécurité longue durée sans PQC. L’annonce de juin 2026 ne fait que fixer une date à un engagement déjà pris.

## Pourquoi 2027 ? La menace « récolter maintenant, déchiffrer plus tard »

Aucun **ordinateur quantique** capable de casser le chiffrement RSA ou les courbes elliptiques n’existe aujourd’hui. Pourquoi, alors, agir dès 2027 ? La réponse tient en un acronyme : HNDL, pour *Harvest Now, Decrypt Later* – « récolter maintenant, déchiffrer plus tard ». Des acteurs étatiques collectent dès aujourd’hui des flux chiffrés (communications diplomatiques, dossiers médicaux, secrets industriels) dans l’espoir de les déchiffrer le jour où une machine quantique suffisamment puissante – le fameux « Q-Day » – verra le jour.

Le raisonnement s’appuie sur l’inégalité formulée par le cryptographe Michele Mosca : si la durée pendant laquelle une donnée doit rester secrète (X), additionnée au temps nécessaire pour migrer son infrastructure (Y), dépasse le délai avant l’arrivée d’un ordinateur quantique cryptographiquement pertinent (Z), alors il est déjà trop tard. Pour des données de santé ou de défense qui doivent rester confidentielles trente à cinquante ans, la fenêtre de risque est *déjà* ouverte. L’algorithme de Shor, connu depuis 1994, garantit qu’une machine quantique mature réduirait à néant la sécurité de RSA et de l’ECC. La cryptographie résistante au quantique n’est donc pas une assurance contre un risque hypothétique de 2040 : c’est une réponse à une captation qui a commencé. Les calendriers de migration publiés en mai 2026 traduisent déjà cette urgence en obligation concrète : RSA-2048 doit être abandonné pour tout nouveau système d’ici 2030, avec une transition complète visée pour 2035.

L’ANSSI n’est pas isolée dans ce raisonnement. Le calendrier de 2027 converge, presque à la date près, avec le « procurement gate » de la NSA américaine (CNSA 2.0), qui impose que toute nouvelle acquisition destinée aux systèmes de sécurité nationale prenne en charge les algorithmes post-quantiques à compter du 1er janvier 2027. Deux des autorités de certification les plus exigeantes au monde ont convergé indépendamment sur la même année.

## Le calendrier de la transition post-quantique

La décision française s’inscrit dans une chronologie mondiale qui s’accélère depuis la standardisation des premiers algorithmes par le NIST – la consultation publique sur le rapport NIST IR 8547, consacré à la transition post-quantique, s’est d’ailleurs close en janvier 2025 après environ deux mois de retours de la communauté technique. Le tableau ci-dessous retrace les jalons qui mènent de l’appel à candidatures de 2016 à l’objectif européen de transition complète en 2035.

| Date | Jalon | Acteur | 
|---|---|---|
| 2016 | Lancement de l’appel à candidatures pour standardiser la PQC | NIST (États-Unis) | 
| 2022 | Premier avis sur la transition post-quantique | ANSSI | 
| Avril 2024 | Recommandation pour une feuille de route coordonnée | Commission européenne | 
| 13 août 2024 | Finalisation des standards FIPS 203, 204 et 205 | NIST | 
| Mai 2025 | « Agreed Cryptographic Mechanisms v2 » (usage hybride imposé) | ECCG (UE) | 
| Juin 2025 | Premier livrable de la feuille de route PQC | Groupe de coopération NIS | 
| 16 juin 2026 | Annonce de l’arrêt des certifications sans PQC | ANSSI | 
| 1er janv. 2027 | Obligation CNSA 2.0 pour les nouveaux achats + début de l’arrêt ANSSI | NSA / ANSSI | 
| 31 déc. 2030 | Systèmes à haut risque migrés ; achats « quantum-safe » uniquement | UE / ANSSI | 
| 31 déc. 2035 | Transition complète, autant que possible | UE | 

Ce séquençage révèle une logique : la standardisation d’abord (2024), les feuilles de route ensuite (2025-2026), puis les obligations contraignantes (2027-2030). L’ANSSI se positionne à la charnière, au moment où la théorie devient contrainte réglementaire.

## Qu’est-ce que la cryptographie post-quantique ?

