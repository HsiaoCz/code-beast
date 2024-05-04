#!/bin/bash
sudo prometheus --web.listen-address=:9002 & sudo prometheus --config.file=prometheus.yml