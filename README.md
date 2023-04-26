# Thread-safe map

## Install

```sh
go get -u github.com/xprgv/tsmap-go
```

## Usage

```go
storage := tsmap.NewTsMap[int, string]()
storage.Set(1, "data1")
data, exist = storage.Get(1)
storage.Size()
storage.ForEach(func(value string) {
// some processing
})
_, _ = storage.Pop(1)
storage.Delete(1)
storage.Flush()
```
