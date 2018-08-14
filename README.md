# scratch

This is a blank Go "Hello World" template.

## Building

With `Make`:

`make`

With Docker:

`docker build -t ehotinger/scratch .`

With acb locally:

- `docker run -v $(pwd):/workspace --workdir /workspace -v /var/run/docker.sock:/var/run/docker.sock acb exec --homevol $(pwd) --task acb.yaml --debug`
- `docker run -v $(pwd):/workspace --workdir /workspace -v /var/run/docker.sock:/var/run/docker.sock acb exec --homevol $(pwd) --task inline.yaml --debug`
- `docker run -v $(pwd):/workspace --workdir /workspace -v /var/run/docker.sock:/var/run/docker.sock acb exec --homevol $(pwd) --task acb-img.yaml --debug`