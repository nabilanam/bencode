# bencode

bencode implementation in go

## example

Encode a map[string]interface{}
```
e := New(map[string]interface{}{
		"spam": []interface{}{
			"a",
			"b",
		},
	})
e.Encode()
```
Decode a dictionary
```
d := New([]byte("d4:spaml1:a1:bee"))
d.Decode()
```
