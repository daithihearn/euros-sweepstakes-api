# Euros Sweepstakes API


## Stack

- Go
- Redis

## API
You will also require `make` to be installed.

Then to run locally simply run:

```bash
make run
```

To build the executable binaries locally run:

```bash
make build
```
The binaries will be installed in the build folder and can be run directly.

If you want to build the docker image run:
    
```bash
make image
```

## Sync Job
To run the sync locally run:

```bash
make sync
```

The `make build` command described in the API section will build both binaries.

The `make image` command described in the API section will build a single docker image for both the API and sync job.

To run the docker image for the sync job run like so:

```bash
docker run -d --rm sweepstakes-sync ./sync
```

Running without the `./sync` will run the API.

## CORs
You must configure CORs by setting an environment variable `CORS_ALLOWED_ORIGINS` to a comma separated list of origins. For example:

```bash
CORS_ALLOWED_ORIGINS=http://localhost:3000,https://sweepstakes.es
```

Please ensure there are no spaces in the list.