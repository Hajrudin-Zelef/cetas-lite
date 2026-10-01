---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-18
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Google", "Intel", "Microsoft", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "amd", "gpu", "intel", "luna", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [2211, 2361]
sha256: 9f0d84626e5769dc6cdea7b2f7fb7e13fedd6eb243b9b96c23f54f57d5cc3758
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

Le **14 septembre 2026**, des chercheurs (KU Leuven, ETH Zurich, Durham, Google) ont
divulgué **DDRop** : avec un interposeur DDR5 à **moins de 200 $**, ils ont cassé les
garanties d'intégrité mémoire d'**Intel TDX, Intel Scalable SGX et AMD SEV-SNP** — en
faisant « disparaître » des écritures mémoire (le processeur relit des données chiffrées
périmées sans s'en apercevoir). Ils ont pu lire la mémoire d'une VM victime et **forger
la mesure d'attestation**. La leçon : le chiffrement mémoire **sans fraîcheur garantie**
a des limites face à un attaquant **physique** équipé. Pour un datacenter on-prem avec
accès physique contrôlé (partie 4), le risque est contenu ; pour du cloud public, c'est
un argument pour la **défense en profondeur** (contrôle d'accès + attestation + HSM),
jamais une technologie seule.

## 156. NVIDIA H100 CC, Arm CCA, Red Hat CoCo (vérifié)

- **NVIDIA H100 Confidential Computing** : mode CC du GPU — VRAM chiffrée
  (AES-256-GCM), transferts CPU-GPU chiffrés (Confidential CUDA), attestation du GPU.
  Surcoût : **~2-5 %** de débit (ordre de grandeur constaté). S'intègre avec TDX/SEV-SNP
  pour une chaîne CPU+GPU attestée (vérifié, 2026).
- **Red Hat OpenShift** : support **Confidential GPU** en Technology Preview dans
  « OpenShift sandboxed containers 1.12 » (avril 2026, vérifié) — inférence de modèles
  propriétaires sur infra non fiable, clés délivrées après double attestation CPU+GPU.
- **Arm CCA** : overhead **1-5 %** annoncé, intéressant pour l'edge industriel
  (voir section 154).

## 157. Ce qui n'est pas encore là : la liste honnête

Au 27/09/2026, **non trouvé / non mature** :

- **FIPS 206 (HQC)** finalisé — en cours (section 148).
- Migration **complète** des PKI vers ML-DSA (les certificats PQC restent rares en
  production ; le hybride domine).
- Support PQC **natif et complet** dans tous les HSM d'entrée de gamme (YubiHSM 2 :
  pilote uniquement).
- Attestation **standardisée et interopérable** entre tous les TEE (les formats
  divergent encore ; les frameworks comme Veraison/Trustee progressent).
- Ordinateur quantique cryptographiquement pertinent — **aucun au 27/09/2026**
  (les estimations crédibles parlent de 2029-2032 au plus tôt, avec une large
  incertitude).

Ne pas acheter sur des promesses : exiger des **références déployées** et des
**validations FIPS/Common Criteria** datées.

# PARTIE 10 — APPROFONDISSEMENTS HSM : L'INTÉGRATION AU QUOTIDIEN

## 158. Intégrer un HSM à Microsoft AD CS (pas-à-pas)

1. Installer le **KSP/CSP** du HSM sur le serveur AD CS (YubiHSM KSP, Luna CSP...).
2. Lors de l'installation du rôle AD CS, choisir « utiliser un module de stockage de
   clés existant » et sélectionner le provider du HSM.
3. Générer la clé **dans** le HSM (ne jamais générer en logiciel puis importer pour
   une AC — sauf migration documentée).
4. Vérifier : `certutil -store` doit montrer la clé avec le provider HSM ; un export
   doit **échouer** (clé non exportable).
5. Sauvegarde : dupliquer la clé vers le 2e HSM (cérémonie, section 133).
6. Tester l'émission et la révocation avant de mettre en production.
7. Documenter le provider, le slot et la procédure de bascule.

Piège : AD CS en **cluster** ou après migration de serveur — le KSP doit être installé
et la clé présente sur chaque nœud ; tester le failover.

## 159. OpenSSL provider + HSM : exemple nginx

Avec OpenSSL 3.x, le modèle **provider** (remplaçant les engines) permet à nginx de
faire les handshakes TLS avec la clé privée dans le HSM :

```
# openssl.cnf (extrait)
[provider_sect]
pkcs11 = pkcs11-provider

# nginx : la clé est référencée par URI PKCS#11, jamais copiée
ssl_certificate     /etc/nginx/certs/server.crt;
ssl_certificate_key "pkcs11:token=web;object=tls-key;type=private";
```

Points de vigilance : le PIN du token doit être fourni au démarrage (fichier protégé
ou variable d'environnement via un helper), la latence handshake augmente (~ms), et il
faut **cacher les sessions** (session tickets/resumption) pour ne pas saturer le HSM.
Dimensionner : chaque handshake RSA-2048 = 1 op asymétrique — à 1 000 handshakes/s,
prévoir ~500 tps utiles (voir section 36).

## 160. EJBCA + HSM : les points de vigilance

- Le **CryptoToken PKCS#11** d'EJBCA pointe vers la lib du HSM ; tester avec
  `p11tool` avant de configurer.
- Séparer les **tokens** : un token/slot par CA (racine, intermédiaire) — jamais le
  même slot pour la racine offline et l'émettrice en ligne.
- Les **workers** (signature OCSP, CMP) utilisent le même HSM : dimensionner le débit
  cumulé.
- Sauvegarde EJBCA (base de données) **et** HSM : restaurer l'un sans l'autre ne sert à
  rien — tester la restauration complète (section 138).
- Durcissement : EJBCA sur OS durci, accès admin via client cert, pas d'Internet.

## 161. Monter une TSA avec OpenTSA + HSM

**OpenTSA** (openssl ts) + clé dans le HSM via provider PKCS#11 :

1. Générer la clé TSA dans le HSM, émettre le certificat TSA (EKU = timeStamping,
   **critique**, un seul usage).
2. Configurer `openssl ts -reply` avec `-signer` pointant vers le provider.
3. Source de temps : **NTS** (Network Time Security) vers des serveurs de confiance,
   avec monitoring du décalage (une TSA dont l'horloge dérive émet des jetons
   contestables).
4. Dimensionner : 1 jeton = 1 signature — à 100 jetons/s, un petit HSM suffit ; à
   10 000/s, prévoir le modèle adapté.
5. Publier la politique de la TSA (précision, exactitude) — exigé pour le qualifié
   eIDAS.

## 162. DNSSEC avec HSM : BIND 9 / Knot

- La **KSK** dans le HSM (PKCS#11), la **ZSK** en logiciel avec rotation fréquente
  (90 jours) ou aussi en HSM selon la politique.
- BIND 9 : `dnssec-keyfromlabel` pour générer la clé **dans** le HSM, configuration
  `key-directory` + moteur PKCS#11.
- **Rollover KSK** : procédure RFC 6781 (double signature) — planifier des semaines à
  l'avance (propagations DNS, TTL).
- Sauvegarde : la KSK perdue = impossible de signer la zone — backup HSM obligatoire.
- Volume : quelques signatures par jour (signatures de RRset à chaque changement de
  zone) — le plus petit HSM suffit.

## 163. Cluster Luna HA : principes de configuration

Le mode **HA** de Thales (principe général, détails dans la doc constructeur) :

1. Créer un **groupe HA** : un membre primaire, un ou plusieurs secondaires.
2. Les objets (clés) sont **répliqués** chiffrés entre membres ; les clients voient un
   slot virtuel unique.
3. Le client bascule automatiquement en cas de panne d'un membre (timeout
   configurable).
4. **Ne pas confondre** : la HA protège de la panne, pas de la compromission logique
   ni de la perte de tous les membres (d'où le backup offline, section 41).
5. Tester la bascule **réellement** (débrancher un membre) avant la mise en production
   — une HA jamais testée est une supposition.

## 164. Négocier l'achat d'un HSM : retours terrain

- Les prix publics n'existent pas : tout se fait **sur devis** via partenaires. Demander
  **3 devis** (Thales, Utimaco, Entrust) avec la **même** fiche de besoin (débit,
  partitions, certifications, support).
- Négocier le **support** (15-20 %/an) : durée d'engagement, SLA de remplacement
  (J+1 vs 4 h), inclus ou non des mises à jour firmware majeures.
- Demander une **maquette** (POC) : les constructeurs prêtent du matériel 30-60 jours
  pour valider l'intégration (PKCS#11 avec vos applis).
- Vérifier la **fin de commercialisation** du modèle (pas d'achat sur une gamme en
  fin de vie sans remise massive et stock de pièces).
- Faire chiffrer la **prestation d'installation + cérémonie** : souvent 5-10 k€, à
  intégrer au budget dès le départ (section 48).

## 165. TCO comparé : 3 scénarios sur 5 ans (détaillé)

Hypothèses : besoin ~500 ops/s, 3 partitions, HA (2 nœuds).

