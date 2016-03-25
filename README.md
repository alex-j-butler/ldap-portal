# LDAP Portal

A Go application to manage Qixalite LDAP credentials.

[![Build Status](https://ci.qixalite.com/buildStatus/icon?job=LDAP-Portal)](https://ci.qixalite.com/job/LDAP-Portal)

## Starting the Go application

`make run` will compile to Go application and run it.

The server will be started at [localhost:4000](http://localhost:4000).

## Installing for production

This repository includes an example Systemd unit file for running the LDAP portal as a daemon for use with production servers, located at `support/ldap-portal.service`

