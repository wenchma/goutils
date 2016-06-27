package goutils

import (
    "fmt"
)

type Table struct {
	Fields     []string
	Rows       []map[string]string
	fieldSizes map[string]int
}

// NewTable - Creates a new table.
func NewTable(fields []string) *Table {
	return &Table{
		Fields:     fields,
		Rows:       make([]map[string]string, 0),
		fieldSizes: make(map[string]int),
	}
}

// AddRow - Adds table row.
func (t *Table) AddRow(row map[string]interface{}) {
	newRow := make(map[string]string)
	for _, k := range t.Fields {
		v := row[k]
		// If is not nil format
		// else value is empty string
		var val string
		if v == nil {
			val = ""
		} else {
            val = fmt.Sprintf("%v", v)
		}

		newRow[k] = val
	}
	t.calculateSizes(newRow)
	if len(newRow) > 0 {
		t.Rows = append(t.Rows, newRow)
	}
}

// Print - Pretty print table
func (t *Table) Print() {
	t.printDash()
	t.printField()
	t.printDash()
	for _, row := range t.Rows {
		t.printRow(row)
	}
	t.printDash()

}

func (t *Table) printDash() {
	s := "+"
	for _, k := range t.Fields {
		for i := 0; i < t.fieldSizes[k]; i++ {
			s += "-"
		}
		s += "+"
	}
	fmt.Println(s)
}

func (t *Table) printField() {
	s := "|"
	for _, k := range t.Fields {
		s += fmt.Sprintf(" %s ", k)
		for i := 0; i < (t.fieldSizes[k] - len(k) - 2); i++ {
			s += " "
		}
		s += "|"
	}
	fmt.Println(s)

}

func (t *Table) printRow(row map[string]string) {
	s := "|"
	for _, k := range t.Fields {
		s += fmt.Sprintf(" %s ", row[k])
		for i := 0; i < (t.fieldSizes[k] - len(row[k]) - 2); i++ {
			s += fmt.Sprintf(" ")
		}
		s += "|"
	}
	fmt.Println(s)
}

func (t *Table) calculateSizes(row map[string]string) {
	for _, k := range t.Fields {
		v, ok := row[k]
		if !ok {
			continue
		}

		vlen := len(v)
		// align to field name length
		if klen := len(k); vlen < klen {
			vlen = klen
		}
		vlen += 2 // + 2 spaces
		if t.fieldSizes[k] < vlen {
			t.fieldSizes[k] = vlen
		}
	}
}

