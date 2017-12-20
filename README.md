<img src="assets/gallows.png" align="right" width="100" height="auto"/>

# Hangman Istio Style

This is an implementation of the traditional game of Hangman. This application
is broken into 3 separate web services:

* Dictionary Service
* Game Service
* Hangman Service

This application is driven via a cli which connects to the Hangman service and
tracks game progress and status.

This demo application is deployed via Kubernetes and leverages an [Istio](http://istio.io)
service mesh to orchestrate the cluster and manage inter-service communication by
dialing various knobs.


---
<img src="assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)