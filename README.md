# URF

## Description
URF is a universal hot-reload tool, similar to various NPM offerings. It watches a directory and executes a makefile target in that directory upon change. In the spirit of Ken, URF was made to be as simple as possible. Just `urf` in a directory with a makefile and you're hot reloading.

## Installation
`go get github.com/khodges42/urf`

## Usage
```
urf
  -dir string
    	Directory to watch (Default is .) (default ".")
  -maketarget string
    	Make Target (Default is urf) (default "urf")
  -rate-limiter int
    	Reload time, in ms (default is 2000) (default 2000)
```


### Makefile setup
URF will be looking for a file in the specified directory named `Makefile`. In this makefile, there should be a target named `urf:` that specifies what gets executed.

For those unfamiliar with makefiles, or those looking for examples, here are a few.

    ```
    # DOCKERFILE
    urf:
        docker stop foo
        docker rm foo
        docker build --no-cache -t foo .
        docker run -i -t --rm --env-file=./config.env -p=8080:8080 --name="foo" foo
    ```


    ```
    # GO
    urf:
        go build main.go
        ./foo
    ```


    ```
    # PYTHON
    urf:
        python -m unittest -v -b .
    ```

