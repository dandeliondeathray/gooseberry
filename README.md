# Gooseberry
An experimental continuous integration system. You should not use this.

## Motivation
Why make another CI system?

- Learn by doing.
- All CI systems suck, because all software sucks.
- I need a CI system for the niancat-micro project.

No one else should use this.

## Features
These are the features I want in a CI system:

- All configuration is done via version controlled files.
- Admins and developers have different user interfaces, as they use the system differently.
- Explicitly designed for build pipelines.
- Simple support for building Docker images.
- Each build step produces an artifact.

## Build pipelines
Assume you have a pipeline that builds an executable `hello`, which depends on a library `greeting`.
The library has unit tests. The pipeline might look something like

1. Build `greeting`.
2. Run unit tests for `greeting`.
3. Build `hello`.

In many systems you would trigger the build at step 1, and on success each step would trigger the
next, in a forward chaining manner. In Gooseberry the trigger is done for step 3, and that part of
the pipeline determines the sequence of builds on which it depends. For this example the end result
is the same, but for build piplines with branches, this approach ensures you know exactly what
branch you need to build.

## Specification
The specifications for Gooseberry is written in Gherkin, available in the `features` directory.

## License
Apache License, version 2.0. See LICENSE.
