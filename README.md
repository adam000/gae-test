# gae-test

My attempts at using Google App Engine's Flexible Environment. I'm trying to make this as easy to reproduce as possible
for the sake of getting help with figuring things out.

## Problem #1: local import paths not working

I attempt to run `aedeploy gcloud app deploy` for my app and it fails to build, despite being able to run `go run main.go`
from the root directory successfully (note that `gae-test` is in the root folder of my `$GOPATH`, which is where all of my source code lives in local development).

Output of running deploy command is at OUTPUT.txt.
