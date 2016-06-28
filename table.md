# Pretty Table in console for golang

## Usage


### Print table

```bash
import "github.com/wenchma/goutils"

table := gotuils.NewTable([]string{"Id", "Name"})
table.AddRow(map[string]interface{}{"Id": "100", "Name": "Mars"})
table.AddRow(map[string]interface{}{"Id": "200", "Name": "Helloworld"})
table.Print()
```

### Output

```bash
+-----+------------+
| Id  | Name       |
+-----+------------+
| 100 | Mars       |
| 200 | Helloworld |
+-----+------------+
```
