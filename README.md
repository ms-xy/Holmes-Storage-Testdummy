# Holmes-Storage: A Storage Planner Test-Dummy for Holmes Processing

## Overview


## Dependencies


## Compilation


## Installation
* Copy the default configuration file located in config/storage.conf.example and
  change it according to your needs.
* Execute storage by calling
  `./Holmes-Storage --config <path_to_config>`

### Supported Databases
Holmes-Storage supports multiple databases and splits them into two categories:
Object Stores and Document Stores. This was done to provide users to more easily
select their preferred solutions while also allowing the mixing of databases for
optimization purposes.
The Test-Dummy implementation only supports in-memory storage, no configuration
for this one is required or possible.
