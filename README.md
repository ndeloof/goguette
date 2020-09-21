# goguette

![logo](logo.png)

GoGeuette is a golang code generator to produce collections-like types
for a Golang struct. This is an attempt to balance lack of generics in 
Golang 1.x

Generated code is hignly inspired by [Koltin's collections](https://kotlinlang.org/api/latest/jvm/stdlib/kotlin.collections)


## generated code

see [example generated code](example/foo_goguette.go).

With generics support, one would define a List of Foo elements as
`List[Foo]` ... but we're stuck in Golang 1.x, so we generate code
as `ListƒFoo` instead. 

If you wonder, `ƒ` is Latin f with hook 
[U+0192](http://www.fileformat.info/info/unicode/char/0192/index.htm). This has been chosen as it can be read "_List oƒ Foo_".

Can type on Mac's keyboard using `Alt+F`, not sure how portable this is :P 
IDE autocompletion does it for you anyway.

Some/None is implemented as a `FooɁ`, here again abusing Unicode to adopt [Latin Global Stop](http://www.fileformat.info/info/unicode/char/0241/index.htm) that looks like question mark and makes a valid Golang identifier. 

Please note we don't generate interface and a proper private implementation. This is by intent, so you can
add you own methods, typically cross types mapping functions that we can't generate. You'll have to write
those by yourself if you want to implement filter.map.reduce, we only can generate filter.reduce

```golang
func (o FooɁ) MapToBar(transformation func(Foo) Bar) BarɁ { ...
```

## About the name

pronounce just like `go get`

"goguette" is a French noun used to designate both a party where the rules are not in order and where everything is allowed, and humorous words that can be perceived as being inappropriate.
