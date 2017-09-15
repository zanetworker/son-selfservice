
[![Build Status](https://travis-ci.org/zanetworker/son-selfservice.svg?branch=master)](https://travis-ci.org/zanetworker/son-selfservice)
[![Go Report Card](https://goreportcard.com/badge/github.com/zanetworker/son-selfservice)](https://goreportcard.com/report/github.com/zanetworker/son-selfservice)

# Son-Selfservice
Son-Selfservice is an independent sonata app that allows you to manage the state of FSMs and SSMs through a user interface.

![Selfservice](selfservice_portal.png?raw=true "Self-service Portal Architecture")

### Architecture
![Architecture](self-service-portal.png?raw=true "Self-service Portal Architecture")


### How do I run this?
You need Docker >= v17.03.1-ce installed on your machine as well as Docker Compose >= v1.11.2  and... that's it. Build the project:

```bash
chmod +x one_for_all.sh
./one_for_all.sh
```

### Test functionality
In addition to the selfservice application code, there is also a stub SSM, implementation that currently does nothing except populate a list of FSMs which should be similar to what the actual SSM would handle it.

To test it, you need to have python (>= Python 2.7.11) installed. Then simply do the following:

```bash
cd ssm_stub
pip install -r requirements.txt
python ssm_stub.py
```

The application needs to be running. Once you do that, you will find a list of FSMs populated on the UI.


# License
son-selfservice is published under Apache 2.0 license. Please see the LICENSE file for more details.


# Lead Developers
The following lead developers are responsible for this repository and have admin rights. They can, for example, merge pull requests.
- [Adel Zaalouk](https://github.com/zanetworker)
