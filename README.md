<img src="assets/gallows.png" align="right" width="100" height="auto"/>

# Hangman Istio Style

This is an implementation of the traditional game of Hangman. This application
is broken into 3 separate web services:

* Dictionary
* Game
* Hangman

This application is driven via a CLI which connects to the Hangman service and
tracks game progress and status.

The demo is deployed via Kubernetes and leverages an [Istio](http://istio.io)
service mesh to orchestrate the cluster. Inter-service communication can be dialed
in by injecting envoys policies in the cluster.

The CLI connects to the Hangman service to start new games and issue guesses. The game
service tracks the game state as a user takes turn to guess a word. Words are provided by
the dictionary service which loads a collection of words from disk and selects random
words to guess.

---
<img src="assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)