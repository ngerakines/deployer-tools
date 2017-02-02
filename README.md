# deployer-tools

A utility that evaluates a mapping of project branches to services in clusters, outputting docker swarm actions.

## Usage

```
Usage: deployer-tools [options]
  Helps deploy things.
Options:
  -event string
    	The location of the event file.
  -host-template string
    	The template used to build the host DNS entry.
  -mapping string
    	The location of the mapping file.
  -silent
    	Operate silently with not output. Defaults to true. (default true)
```

### Example

Given an event:

```json
{
  "Project":"hello-world",
  "Branch":"master",
  "Container":"ngerakines/hello-world-nodejs:CI1"
}
```

And a mapping:

```json
{
   "hello-world": {
      "master": {
         "greeting": ["hello"]
      }
   },
   "project": {
      "branch": {
         "cluster": ["service"]
      }
   }
}
```

    $ ./deployer-tools -mapping mapping.json -event event.json
    $ docker -H manager01.greeting.internal:2375 service update --image ngerakines/hello-world-nodejs:CI1 hello
