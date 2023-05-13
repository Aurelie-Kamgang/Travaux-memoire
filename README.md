# Travaux-mémoire
**Evaluation de l'impact de la connectivité internet sur l'accès aux services cloud depuis le Cameroun.**

## Implémentation 
Nous allons utilisé une procédure basée sur des mesures de performance effectuées sur une infrastructure Cloud comme banc d’essai.
Pour le moment, les aspects des services Cloud à évaluer sont sur le SaaS.

Les mesures de performance qui seront analysées sont :
- Temps de réponse du service
- Latence
- Bande passante
- Débit

## Etapes

**1.  Générer des charges de travail avec l'outil Benchmark http **Apache JMeter** :**
    
    -   Concevoir un plan de test dans JMeter qui inclut les requêtes HTTP à envoyer à l'application ou au service cloud.
    
    -   Configurer le nombre d'utilisateurs simulés et la fréquence des requêtes pour simuler une charge de travail réaliste.
    
    -   Exécuter le test et collecter les données de performance,  comprennant les temps de réponse, les débits, les latences et les bandes passantes.
    
**2.  Exporter les données collectées dans un fichier CSV :**
    
    -   Dans JMeter, exporter les résultats du test dans un fichier CSV.
    
    -   S'assurer que le fichier CSV inclut toutes les métriques de performance nécessaires pour l'analyse.
    
**3.  Prétraiter les données à l'aide de Pandas :**

    -   Charger les données du fichier CSV dans un DataFrame Pandas.
    
    -   Vérifier les données pour les valeurs manquantes ou les erreurs.
    
    -   Ajouter des colonnes supplémentaires pour les métriques de performance dérivées, telles que le taux de réussite, la moyenne et l'écart type des temps de réponse, etc.
    
**4.  Analyser les données à l'aide de NumPy et Matplotlib :**
    
    -   Utiliser NumPy pour effectuer des calculs statistiques avancés sur les données, tels que la régression linéaire, l'analyse de variance, etc.
    
    -   Utiliser Matplotlib pour créer des graphiques pour représenter les données de performance, y compris les histogrammes, les graphiques en boîte, les graphiques de densité, etc.
    
    -   Ajouter des légendes, des titres et des étiquettes d'axe pour rendre les graphiques plus compréhensibles.
    
    
**5.  Interpréter les résultats de l'analyse :**
    
    -   Analyser les graphiques et les résultats statistiques pour identifier les tendances et les anomalies dans les données de performance.
    -   Évaluer l'impact des différentes variables sur les performances, telles que le nombre d'utilisateurs simulés, la fréquence des requêtes, etc.
    -   Tirer des conclusions sur les performances de l'application ou du service cloud évalué, et identifier les domaines d'amélioration potentiels (perspectives).
    
    
**Architecture**

![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/98de52a2-25a9-402b-8f4a-55f84ee2e0c9)


