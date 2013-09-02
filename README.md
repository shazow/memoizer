# memoizer

Golang memoizer interface for caching your function calls.

This is a generic implementation which supports providing your own caching
mechanism.

**Status:** Use with skepticism. I'm not convinced the interface and reflection bits are ideal here—suggestions appreciated.


## Usage

Memoizer comes with a simple example caching implementation, called `MemoryCache`.

```go
memoize := Memoize{NewMemoryCache()}

// First call, uncached. `somefunc` gets evaluated with `args...`.
// `r1` gets cached if `err == nil`.
r1, err := memoize.Call(somefunc, args...)

// Second call for the same function and args, so it's already cached.
// `r1 == r2` and `err == nil`.
r2, err := memoize.Call(somefunc, args...)
```

The module is laid out so that you can provide your own caching mechanisms,
such as memoizing and caching directly into memcached or redis with gob
encoding/decoding.

Pull requests with other caching mechanisms are welcome!


## Todo

* Add a `Memoizer.Replace(...)` function (to complement the existing `Memoizer.Call(...)` function) which returns an equivalent memoized
  function.
* Figure out if there's a way to support functions with arbitrary number of
  return values.
* Implement more example caching mechanisms. Memcache, redis, etc.
* Write a benchmark to better understand the performance implications of using reflection.


## Contributing

1. [Check for open issues](https://github.com/shazow/memoizer/issues>) or open
   a fresh issue to start a discussion around a feature idea or a bug.
1. Fork the [memoizer repository on Github](https://github.com/shazow/urllib3>)
   to start making your changes.
1. Write a test which shows that the bug was fixed or that the feature works
   as expected.
1. Send a pull request and bug the maintainer until it gets merged and published.
   :) Make sure to add yourself to ``CONTRIBUTORS.md``.
