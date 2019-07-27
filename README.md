# gopy-pr180

This is a poc to get [PR-180](https://github.com/go-python/gopy/pull/180) working for [gopy](https://github.com/go-python/gopy).

## Getting Started

Clone the repo and build with docker...

```
git clone git@github.com:nickpoorman/gopy-pr180.git
cd gopy-pr180
make docker-build
make docker-run
```

In the docker image run...

```
make build-hi
```

You should see that it failed...

```
root@2178f8676bfa:/src/gopy-pr180# make build-hi
gopy build -vm="python3" -output=/out github.com/go-python/gopy/_examples/hi

--- Processing package: github.com/go-python/gopy/_examples/hi ---

--- building package ---
gopy build -vm=python3 -output=/out github.com/go-python/gopy/_examples/hi
goimports -w hi.go
go build -buildmode=c-shared -o hi_go.so .
python3 build.py
go env CC
python3-config --cflags
python3-config --ldflags
gcc hi.c  hi_go.so -o _hi.so -I/usr/include/python3.6m -I/usr/include/python3.6m  -Wno-unused-result -Wsign-compare -g -fdebug-prefix-map=/build/python3.6-PEMn0O/python3.6-3.6.8=. -specs=/usr/share/dpkg/no-pie-compile.specs -fstack-protector -Wformat -Werror=format-security  -DNDEBUG -g -fwrapv -O3 -Wall -L/usr/lib/python3.6/config-3.6m-x86_64-linux-gnu -L/usr/lib -lpython3.6m -lpthread -ldl  -lutil -lm  -Xlinker -export-dynamic -Wl,-O1 -Wl,-Bsymbolic-functions
cmd had error: exit status 1
output: gcc: error: : No such file or directory
gcc: error: : No such file or directory
gcc: error: : No such file or directory
gcc: error: : No such file or directory
gcc: error: : No such file or directory

2019/07/27 18:41:36 error dispatching command: exit status 1
Makefile:14: recipe for target 'build-hi' failed
make: *** [build-hi] Error 1
```

Modify the generated Makefile in `/out` to include some shared flags.

```
# Add "-fPIC --shared" to the line in the Makefile that builds the _hi.so library.
sed -i -e 's/$(CFLAGS) $(LDFLAGS)/$(CFLAGS) $(LDFLAGS) -fPIC --shared/g' /out/Makefile
```

Rerun the generated Makefile build.

```
cd /out
make build
```

Test that it worked...

```
python3
>>> import hi
>>> dir(hi)
>>> hi.Hello("you")
```
