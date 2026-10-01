---
id: collect-261001-general-networking/general-networking/cryptographie-post-quantique-l-anssi-impose-2027-2026-4
title: "OpenSSL 3.5+ intègre nativement ML-KEM, ML-DSA et SLH-DSA"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "valuation"]
source: docs/RAG/collect-261001-general-networking/cryptographie-post-quantique-l-anssi-impose-2027-2026.md
source_anchor: ""
source_lines: [134, 178]
sha256: f8e089595b8893bb06d4f761981d3d65c4816456099b37ce1d86ef82cbc59e2c
---

# OpenSSL 3.5+ intègre nativement ML-KEM, ML-DSA et SLH-DSA

Au-delà des algorithmes, Samih Souissi a inscrit la décision dans un cadre plus large mêlant gouvernance, planification industrielle, réglementation et souveraineté. Ce vocabulaire n’est pas anodin : il révèle que la transition post-quantique est autant une politique industrielle qu’une mesure de sécurité. En fixant le tempo par la certification, l’État français conserve la maîtrise du calendrier plutôt que de le subir de fournisseurs extra-européens.

Cette dimension souveraine explique aussi la préférence pour l’hybridation et la volonté de conserver la possibilité de certifier des algorithmes alternatifs (comme FrodoKEM), au-delà des seuls choix du NIST. L’Europe ne veut pas dépendre d’un unique jeu d’algorithmes validé outre-Atlantique. La **certification ANSSI** devient ainsi un instrument de politique cryptographique : un moyen d’imposer des standards tout en préservant une marge d’autonomie technologique.

## Cinq prédictions pour la transition post-quantique

- **Une ruée sur les cycles de qualification en 2026.** Les éditeurs vont saturer les laboratoires d’évaluation agréés pour décrocher leur label avant la fermeture de 2027, allongeant mécaniquement les délais.
- **La crypto-agilité deviendra un critère d’achat.** Au-delà de « supporte-t-il la PQC ? », les acheteurs publics exigeront la capacité à changer d’algorithme sans remplacer le matériel – un argument commercial différenciant dès 2027.
- **L’hybride s’imposera comme norme de fait.** Comme pour X25519MLKEM768 en TLS, les déploiements combineront systématiquement classique et post-quantique jusqu’à ce que la confiance dans les algorithmes récents soit consolidée.
- **Un effet d’entraînement sur le privé.** Bien que ciblant les produits qualifiés, la mesure poussera banques, assureurs et santé à exiger la PQC de leurs fournisseurs, bien avant toute obligation légale directe.
- **D’autres régulateurs européens suivront.** L’Allemagne (BSI) et d’autres agences nationales devraient durcir leurs propres schémas de certification d’ici 2027-2028, sous l’impulsion de la feuille de route coordonnée.

## FAQ : cryptographie post-quantique et certification ANSSI

### Qu’a exactement annoncé l’ANSSI en juin 2026 ?

L’agence a annoncé qu’à partir de 2027, elle cesserait progressivement de certifier et de qualifier les produits de sécurité qui n’intègrent pas de cryptographie post-quantique, et a appelé les entreprises à n’acheter que des solutions « quantum-safe » d’ici 2030. L’annonce a été faite le 16 juin 2026 lors de la conférence France Quantum 2026 à Paris.

### La cryptographie post-quantique est-elle différente de la cryptographie quantique ?

Oui. La cryptographie post-quantique repose sur des algorithmes classiques (logiciels) résistants aux attaques quantiques, déployables sur les ordinateurs actuels. La cryptographie quantique, comme la distribution quantique de clés (QKD), nécessite du matériel photonique spécialisé. L’obligation de l’ANSSI porte sur la première.

### Quels algorithmes faut-il adopter ?

Les standards NIST finalisés en 2024 servent de référence : ML-KEM (FIPS 203) pour l’établissement de clés, ML-DSA (FIPS 204) et SLH-DSA (FIPS 205) pour les signatures. L’ANSSI recommande de les utiliser en mode hybride, combinés à un algorithme classique éprouvé.

### Un ordinateur quantique peut-il déjà casser mon chiffrement ?

Pas aujourd’hui. Aucune machine quantique capable de casser RSA ou l’ECC n’existe encore. Le risque immédiat vient de la stratégie « récolter maintenant, déchiffrer plus tard » : des données interceptées aujourd’hui pourraient être déchiffrées demain. C’est pourquoi les données à longue durée de vie doivent migrer sans attendre.

### Mon entreprise est-elle concernée si elle ne vend pas à l’État ?

Directement, la mesure vise les produits soumis à la certification ANSSI. Indirectement, elle façonnera l’ensemble du marché : les fournisseurs alignant leurs gammes sur les exigences publiques, la PQC deviendra le standard proposé à tous les clients, y compris privés.

### Par où commencer une migration post-quantique ?

Par l’inventaire cryptographique : cartographier où et comment le chiffrement est utilisé (TLS, VPN, signatures, clés SSH). Ensuite, prioriser les systèmes exposés et les données sensibles à longue durée de vie, puis déployer des mécanismes hybrides pour assurer la compatibilité pendant la transition.

### Quel est le lien avec la feuille de route européenne ?

La décision française s’aligne sur la feuille de route coordonnée de l’UE, qui vise la migration des systèmes à haut risque d’ici fin 2030 et une transition complète d’ici 2035. L’ANSSI utilise la certification comme levier national pour atteindre ces objectifs communs.

### Related Coverage : pour aller plus loin

*Sources et références : NIST – First Post-Quantum Encryption Standards, NIST Post-Quantum Cryptography Project, Commission européenne – Coordinated Implementation Roadmap, ANSSI / Groupe NIS, IT Social, The Quantum Insider. Article publié le 17 juin 2026.*
