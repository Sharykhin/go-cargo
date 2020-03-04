Go Cargo
========

This is a simplified project where each domain is design as a seprarate microservice but which may share the same database.
It was done intentionally not to focus on details but rather on general directory structure in terms of each domain
and general workflow.

#### Requirements:
 - [Docker](https://www.docker.com/)

#### Usage:
1. Make a copy of .env.example file to .env. This file keep enviroment variable for local development, mostly for port mapping.
```bash
cp .env.example .env
```

#### Work-flow:
General workflow:
Company want to get some goods of another company. It makes an order and then application finds the nearest warehouse for where goods will be kept for some period of time.

1. Domain Company makes an order of some goods and send an event
3. Domain Router catches that event and based on criteria like total weights of all order concurrenly try to find the nearest warehour that fits all requirements (days to keep, available space at date arrival) and register that order in a warehouse


##### Author: 
Sharykhin Sergey (chapal@inbox.ru)
