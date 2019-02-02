# URF

## Description
URF is a universal hot-reload tool, similar to various NPM offerings. It watches a specified directory and executes a makefile target in that directory upon change.

## Installation
`go get github.com/khodges42/urf`

## Usage
`urf ~/code/myProject/`

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