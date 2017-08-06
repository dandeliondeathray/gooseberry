# Terminology for CI systems

## Nouns

* **Job**: Instance of a *job specification*.

* **Job specification**: Description of something to be done by a *worker*.

* **Source**: Somewhere the CI system can fetch the latest version of the source code.

## Verbs

* **Trigger job**: Start an instance of a job, 

## Components

* **Source watcher**: Watches for changes in the source code (nominally pushes to a git repo), and emits some type of event when an update is detected.

* **Trigger**: Listens for events in the system, and triggers appropriate actions. Most common (and maybe only) use case: triggering a *worker* when the *source watcher* notifies about a repo update. If pipelines are supported, might also listen for job completion events from workers to trigger the next step of the pipeline.

* **Worker**: Runs jobs. As simple and convention-based as possible; anything more than bare-bones basic has to come from inside the repo.

* **UI**: Displays status about various things, e.g. progress and results of builds and tests.
