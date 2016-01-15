# env-tpl

[![Build Status](https://travis-ci.org/arschles/envtpl.svg?branch=master)](https://travis-ci.org/arschles/envtpl) [![Docker Repository on Quay](https://quay.io/repository/arschles/envtpl/status?token=f8121f5f-0c36-4241-8f33-aa3423ca519f "Docker Repository on Quay")](https://quay.io/repository/arschles/envtpl)

Render Go templates from Environment Variables. Let's say you have the following template file called `pwd.tpl`:

```
The current working directory is {{.PWD}}
```

Run `envtpl -in pwd.tpl`, and you'll see the following printed to STDOUT:

```
The current working directory is /path/you/executed/from
```
