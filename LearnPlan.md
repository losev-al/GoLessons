день  задание  примечание

# 1. Массивы и срезы  Теория

## Теория

знакомство с конвенцией по code style': Конвенция go команды

https://www.sohamkamani.com/blog/golang/arrays-vs-slices/

https://programming.guide/go/nil-slice-vs-empty-slice.html

https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94

https://www.callicoder.com/golang-slices/

https://github.com/golang/go/wiki/SliceTricks

https://go101.org/article/memory-leaking.html

В обязательном порядке изучаем исходники:

https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/slice.go

https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/slice_test.go


## Практика

Реализовать паттерн фасад https://en.wikipedia.org/wiki/Facade_pattern  в соответствии с конвенцией

Для практического задания необходимо создать публичный репозиторий на github'e в контрибьютеры добавить (frankegoesdown, aber4nod, ssmmxx)

Реализация паттерна должна быть следующая:

описать пакет facade в директории pkg/facade
в директории cmd/facade создать файл main.go, который фактически будет представлять собой интеграционное тестирование пакета 

# 2. Карты  

## Теория

https://medium.com/rungo/the-anatomy-of-maps-in-go-79b82836838b

https://www.callicoder.com/golang-maps/

к прочтению обязательно:

https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html

исходники:

основной файл:

https://github.com/golang/go/blob/master/src/runtime/map.go

https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/map_benchmark_test.go

https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/map_test.go

## Практика

Реализовать паттерн строитель https://en.wikipedia.org/wiki/Builder_pattern в соответствии с конвенцией

# 3. Контекст и линтеры  

## Теория

https://github.com/quii/learn-go-with-tests/blob/master/context.md

https://github.com/quii/learn-go-with-tests/tree/master/context

https://habr.com/ru/company/roistat/blog/413175/

https://habr.com/ru/post/457970/

https://golang.org/cmd/vet/

https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html

https://golang.org/pkg/context/

https://habr.com/ru/company/nixys/blog/461723/

## Практика

Реализовать паттерн посетитель https://en.wikipedia.org/wiki/Visitor_pattern в соответствии с конвенцией


# 4. Указатели и типы данных  

## Теория

поинтеры через тесты:

https://github.com/quii/learn-go-with-tests/blob/master/pointers-and-errors.md

https://github.com/quii/learn-go-with-tests/tree/master/pointers


ссылки для изучения

https://rtfm.co.ua/golang-ukazateli-podrobnyj-razbor/

https://medium.com/@vCabbage/go-are-pointers-a-performance-optimization-a95840d3ef85

https://habr.com/ru/post/339192/

https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963

https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html

включая подссылки в статье

https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html

типы в сорсах гошки:

https://github.com/golang/go/blob/0ae9389609f23dc905c58fc2ad7bcc16b770f337/src/runtime/type.go

## Практика

Реализовать паттерн комманда https://en.wikipedia.org/wiki/Command_pattern в соответствии с конвенцией


# 5. Каналы и планировщик  

## Теория

https://github.com/quii/learn-go-with-tests/blob/master/concurrency.md

https://github.com/quii/learn-go-with-tests/tree/master/concurrency

https://docs.google.com/document/d/1-n2O-c3C6kz-4vNKJdtR9bjh-kXqJe7RydSOTBxvB3g/edit?usp=sharing (что то здесь с форматированием все плохо)

https://habr.com/ru/post/308070/

https://github.com/quii/learn-go-with-tests/blob/master/select.md

https://github.com/quii/learn-go-with-tests/tree/master/select

https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html

https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/chan_test.go

https://blog.gopheracademy.com/advent-2019/directional-channels/

исходники: https://github.com/golang/go/blob/master/src/runtime/proc.go

если будет интересно, что за зверь такой sudog в гошном рантайме:

https://groups.google.com/forum/#!topic/golang-codereviews/rC9BLPFvkW8


Планировщик и sync pool

https://docs.google.com/presentation/d/1UCOEHw-0oSUiOA94aTX79Ki7VbsmPtlYlv1EfyI1ylE/edit#slide=id.p

https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html

https://habr.com/ru/post/333654/

https://habr.com/ru/post/277137/

https://dev-gang.ru/article/go-ponjat-dizain-syncpool-cpvecztx8e/

https://golang.org/src/sync/pool.go

https://medium.com/a-journey-with-go/go-understand-the-design-of-sync-pool-2dde3024e277

## Практика

Реализовать паттерн цепочка вызовов https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern в соответствии с конвенцией


# 6.  Сборщик мусора  

## Теория

https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html

https://blog.golang.org/ismmkeynote

https://www.facebook.com/watch/?v=871606259683390

https://habr.com/ru/post/265833/

https://github.com/golang/go/blob/50bd1c4d4eb4fac8ddeb5f063c099daccfb71b26/src/runtime/debug/garbage.go

https://github.com/golang/go/blob/7955ecebfc85851d43913f9358fa5f6a7bbb7c59/src/runtime/extern.go

## Практика

Реализовать паттерн фабричный метод https://en.wikipedia.org/wiki/Factory_method_pattern в соответствии с конвенцией


# 7. Escape analysis + приватные интерфейсы и solid  

## Теория:

escape analysis:
https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html

https://habr.com/ru/company/intel/blog/422447/

interfaces:

private:

blog.j7mbo.com/bypassing-golangs-lack-of-constructors/

Interface segregation principle:

https://en.wikipedia.org/wiki/Interface_segregation_principle

solid go design

https://dave.cheney.net/2016/08/20/solid-go-design

## Практика

Реализовать паттерн стратегия https://en.wikipedia.org/wiki/Strategy_pattern в соответствии с конвенцией


# 8. Профилирование системы  

## Теория:

https://www.ardanlabs.com/blog/2017/06/language-mechanics-on-memory-profiling.html

https://golang.org/pkg/runtime/pprof/

https://golang.org/doc/diagnostics.html

https://habr.com/ru/company/badoo/blog/301990/

очень полезная статья, к изучению обязательна

https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/