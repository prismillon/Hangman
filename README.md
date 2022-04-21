# HANGMAN PROJECT

**Usage:**

Il est recommandé d'utiliser le programme dans un terminal externe à celui de vs code, sinon il est recommandé d'argandir celui-ci

```
git clone https://git.ytrack.learn.ynov.com/PPOIRIER/hangman
cd hangman/
go run . word.txt
```

``word.txt`` est le chemin relatif vers un fichier de mots
un mot sera aléatoirement choisit dans ce fichier et le jeu
commencera

**Sauvegarde:**
```
Il reste 10 essais.







v _ l l _ 
Entrez une nouvelle lettre: save name
```
ou
```
Entrez une nouvelle lettre: stop name
```

**Jouer la sauvegarde**
```
go run . -sw name
```
ou
```
go run . --startWith name
```

**Jouer en ascii**
```
go run . word.txt -lf standard.txt
```
ou
```
go run . word.txt --letterFile standard.txt
```