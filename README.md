## Trying out wasip2 + TinyGo + wasm-tools-go

#### generate bindings
> using https://github.com/ydnar/wasm-tools-go/

```
wit-bindgen-go generate --exports ./foo-wit
```


#### Building
> using tinygo from here: https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2

```
export PATH=/Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/bin:$PATH
export TINYGOROOT=~/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/

tinygo build -target=wasip2 -wit-package ./foo-wit -wit-world foo-namespace:pkg/foo-world -o main.wasm -x -work main.go


WORK=/var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887
wasm-ld --stack-first --no-demangle -L /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/ -o /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887/main /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887/main.o /Users/rajatjindal/Library/Caches/tinygo/compiler-rt-wasm32-unknown-wasi-generic/lib.a /Users/rajatjindal/Library/Caches/tinygo/obj-363c24a529264e57aec91eacb3cbc247bacb915b67966242e30f7fb9.bc /Users/rajatjindal/Library/Caches/tinygo/obj-870ef2ac5a778e9704c61e87bde8b24392ff596a3b0126fe5821e7b1.bc /Users/rajatjindal/Library/Caches/tinygo/wasmbuiltins-wasm32-unknown-wasi-generic/lib.a -mllvm -mcpu=generic --lto-O2 --thinlto-cache-dir=/Users/rajatjindal/Library/Caches/tinygo/thinlto -mllvm --rotation-max-header-size=0
/opt/homebrew/bin/wasm-opt --asyncify -Oz -g /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887/main --output /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887/main
wasm-tools component embed -w foo-namespace:pkg/foo-world ./foo-wit /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887/main -o /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887/main
wasm-tools component new /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887/main -o /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887/main
FuncType { params: [I32, I32], results: [] }
error: failed to encode a component from module

Caused by:
    0: failed to decode world from module
    1: module was not valid
    2: failed to validate exported interface `foo-namespace:pkg/foo@2.0.0`
    3: type mismatch for function `greet`: expected `[] -> [I32]` but found `[I32, I32] -> []`
error: wasm-tools failed: exit status 1
```

#### Debugging
```
wasm-tools dump ./target/wasm32-wasi/release/http_trigger.wasm | grep 'Import.*key-value.*open'                 
```



Hi @dgryski, I have been trying out tinygo/wasm-tools-go repo, and running into a problem, which seems like a tinygo problem (or limitation that it is not supported yet). Like I mentioned earlier, that I am still very new to all these tooling, but I just wanted to bring it to your attention incase if it is interesting to you.

foo.wit
interface foo {
    greet: func() -> string;
}
from running wasm-tools component wit ./foo-wit -j > wit.json:

    {
      "name": "foo",
      "types": {},
      "functions": {
        "greet": {
          "name": "greet",
          "kind": "freestanding",
          "params": [],
          "results": [
            {
              "type": "string"
            }
          ]
        }
      },
      "package": 7
    }

from running `tinygo build -target=wasip2 -wit-package ./foo-wit -wit-world foo-namespace:pkg/foo-world -o main.wasm -x -work main.go :


error: failed to encode a component from module

Caused by:
    0: failed to decode world from module
    1: module was not valid
    2: failed to validate exported interface `foo-namespace:pkg/foo@2.0.0`
    3: type mismatch for function `greet`: expected `[] -> [I32]` but found `[I32, I32] -> []`
error: wasm-tools failed: exit status 1

from running: wasm-tools mutate /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo226195887/main -t | grep -i greet | head -1:

(Notice the param i32 i32, I was expecting it to be something like (param i32) (result i32)
  (func $foo-namespace:pkg/foo@2.0.0#greet (;259;) (type 0) (param i32 i32)

kindly let me know if it makes sense. The repo where I was trying this out is here: https://github.com/rajatjindal/foo-wasip2
```