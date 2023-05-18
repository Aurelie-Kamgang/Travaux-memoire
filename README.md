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
    
    
**Voici l'architecture du scénario décrit plus haut:**    
    
**Architecture**

![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/98de52a2-25a9-402b-8f4a-55f84ee2e0c9)



##  Justification du choix de l'outils Apache Jmeter

Ci-dessous nous avons un tableau comparatif de certains outils de benchmark http:


![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/ca787fd6-fed1-43ae-81e9-2ce733a2bbe8)



 1. D'après le tableau précédent, les outils tel que le Temps réel = Non ,  les Fonctionnalités d'analyse = Avancées et  la Flexibilité = Elevée sont:
				
				- jmeter
				- k6
				- locust
				
 2. Apache JMeter :
 
       Avantages par rapport au travail à effectuer:
       
       -    Capacité à simuler des charges élevées grâce à la possibilité d'exécuter des tests distribués.
       
      - Apache JMeter conserve sa popularité en raison de son historique, de sa polyvalence et de son support étendu.
      
      -   Offre des fonctionnalités avancées d'analyse des résultats, notamment des graphiques, des rapports et des statistiques détaillées.
      
     Inconvénients :
     
      -    Peut nécessiter une consommation de mémoire et de ressources élevée pour exécuter des tests de charge importants.
      
     -   La création de scénarios de test complexes peut nécessiter une expertise approfondie.
 
 
 3. k6

      Avantages par rapport au travail à effectuer:

      -  Capacité à gérer des charges importantes en utilisant une architecture légère et efficace.
      
     -   Prise en charge des tests distribués pour simuler des charges réparties sur plusieurs machines

       Inconvénients:

       -    Certaines fonctionnalités avancées peuvent nécessiter une personnalisation via des scripts JavaScript.
       
        -   Moins de fonctionnalités d'analyse et de rapports prédéfinis par rapport à d'autres outils.
        
        -   La courbe d'apprentissage peut être légèrement plus élevée pour les utilisateurs qui ne sont pas familiers avec JavaScript.

 4. Locust

     Avantages par rapport au travail à effectuer;

     -    Capacité à simuler des charges élevées grâce à une architecture légère et évolutive.
     
      -   Prise en charge de la distribution des tests de charge sur plusieurs machines.

     Inconvénients:

    -    Moins de fonctionnalités d'analyse et de rapports prédéfinis par rapport à certains autres outils.
    
    -   L'interface utilisateur peut sembler moins conviviale pour les utilisateurs novices.

    -   Certaines fonctionnalités avancées peuvent nécessiter une personnalisation via des scripts Python.
 
### Choix

Par rapport au travail à faire **Apache Jmeter** est pertinent car:

- Il est flexible et peut s'adapter et  répondre aux besoins spécifiques des utilisateurs dans différents scénarios d'utilisation (Prise en charge de différents protocoles, Personnalisation des requêtes, Personnalisation des requêtes, etc.)

- Il collecte les données qui seront analysées plus tard donc le paramètre temps ici n'est pas important.

- Apache JMeter est généralement considéré comme l'outil le plus populaire et bénéficie d'une communauté active.

**NB:** Le choix d'utiliser **pandas**, **numpy** et **matplotlib** pour l'analyse est du au fait que:

Apache JMeter est plus adapté pour l'analyse des performances des applications, les bibliothèques Python telles que Pandas, NumPy et Matplotlib peuvent être utilisées pour compléter l'analyse de données liées au réseau, en utilisant les données collectées par JMeter comme source de données.

