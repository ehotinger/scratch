# scratch

This is a blank Go "Hello World" template.

## Building

With `Make`:

`make`

With Docker:

`docker build -t ehotinger/scratch .`

With acb locally:

- `docker run -v $(pwd):/workspace --workdir /workspace -v /var/run/docker.sock:/var/run/docker.sock acb exec --homevol $(pwd) -f acb.yaml --debug`

- `docker run -v $(pwd):/workspace --workdir /workspace -v /var/run/docker.sock:/var/run/docker.sock acb exec --homevol $(pwd) -f acb-inline.yaml --debug`

- `docker run -v $(pwd):/workspace --workdir /workspace -v /var/run/docker.sock:/var/run/docker.sock acb exec --homevol $(pwd) -f acb-img.yaml --debug`

- `docker run -v $(pwd):/workspace --workdir /workspace -v /var/run/docker.sock:/var/run/docker.sock acb exec --homevol $(pwd) -f acb-multiline-build.yaml --debug --registry ehotinger.azurecr.io --id sample`

- `docker run -v $(pwd):/workspace --workdir /workspace -v /var/run/docker.sock:/var/run/docker.sock acb exec --homevol $(pwd) -f acb-push.yaml --debug --registry foo.azurecr.io --username user --password pw --id sample`