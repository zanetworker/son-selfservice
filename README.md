
[![Build Status](https://travis-ci.org/zanetworker/son-selfservice.svg?branch=master)](https://travis-ci.org/zanetworker/son-selfservice)
[![Go Report Card](https://goreportcard.com/badge/github.com/zanetworker/son-selfservice)](https://goreportcard.com/report/github.com/zanetworker/son-selfservice)

# Son-Selfservice
Sonata Self-service Application


### Architecture
![Architecture](self-service-portal.png?raw=true "Self-service Portal Architecture")


### How do I run this?
You need Docker installed on your machine and... that's it. Build the project:

```bash
cd selfservice-backend
chmod +x build.sh
./build.sh
```

After you run the build file, go back to the main project directory then run

```bash
docker-compose up --build
```
# License
son-selfservice is published under Apache 2.0 license. Please see the LICENSE file for more details.



# Lead Developers
The following lead developers are responsible for this repository and have admin rights. They can, for example, merge pull requests.
- [Adel Zaalouk](https://github.com/zanetworker)
