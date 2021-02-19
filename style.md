## Naming
### file
檔案命名一律採用小寫，不用駝峰式
除測試檔案外，不用底線
```
stringutil.go
stringutil_test.go
```

### package
小寫，不用駝峰式
盡量與資料夾名稱相同

###variable
採用駝峰式，以public, private決定首字大小寫
```
apiClient
URLString
```

###constant, function
原則上同變數

## Struct
struct methods一律使用pointer
```
type Cart struct {
    Name  string
    Price int
}

func (c *Cart) UpdatePrice(price int) {
    c.Price = price
}
```

## Error Handling
### 不要濫用 panic
如果預期可能發生錯誤，使用error來處理
```
func run(args []string) error {
  if len(args) == 0 {
    return errors.New("an argument is required")
  }
  // ...
  return nil
}

func main() {
  if err := run(os.Args[1:]); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
```
### type assertion
使用comma ok來handle error
```
t, ok := i.(string)
if !ok {
  // gracefully handle error
}
```

### goroutine
任何一個goroutine都應該有recover來保護程序不會因為panic而crash。

## Declaration
### Slice
盡可能為make提供一個初始容量
```
for n := 0; n < b.N; n++ {
  data := make([]int, 0, size)
  for k := 0; k < size; k++{
    data = append(data, k)
  }
}
```

### Map
使用make來初始化map
```
m1 := map[string]string{} (X)
m1 := make(map[string]string) (O)
```

## Indent
```
// Bad
if err != nil {
    error handling
} else {
    normal code
}

// Good
if err != nil {
    error handling
    return
}
normal code
```