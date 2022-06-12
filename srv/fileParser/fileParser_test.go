package fileParser

import (
	"reflect"
	"testing"
)

func Test_readFormatLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    Format
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Empty string",
			args{""},
			Format{},
			false,
		},
		{
			"invalid input format",
			args{","},
			Format{},
			false,
		},
		{
			"invalid input format",
			args{",,"},
			Format{},
			false,
		},
		{
			"invalid input format",
			args{",,,"},
			Format{},
			false,
		},
		{
			"invalid input format",
			args{"a,,"},
			Format{},
			false,
		},
		{
			"invalid input format",
			args{",a,"},
			Format{},
			false,
		},
		{
			"invalid input format",
			args{",,a"},
			Format{},
			false,
		},
		{
			"invalid input format",
			args{"a,b,c"},
			Format{},
			false,
		},
		{
			"invalid input format",
			args{",10,c"},
			Format{},
			false,
		},
		{
			"invalid input format",
			args{"a,10,c,10"},
			Format{},
			false,
		},
		{
			"valid data",
			args{"name ,10 ,TEXT "},
			Format{
				ColumnName: "name",
				Width:      10,
				DateType:   "TEXT",
			},
			false,
		},
		{
			"valid data",
			args{"name,10,TEXT"},
			Format{
				ColumnName: "name",
				Width:      10,
				DateType:   "TEXT",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFormatLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFormatLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFormatLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readDataLine(t *testing.T) {
	type args struct {
		line    string
		formats []Format
	}
	tests := []struct {
		name    string
		args    args
		want    []Data
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"valid data",
			args{
				line: "Diabetes  1  1",
				formats: []Format{
					{
						ColumnName: "name",
						Width:      10,
						DateType:   "TEXT",
					},
					{
						ColumnName: "valid",
						Width:      1,
						DateType:   "BOOLEAN",
					},
					{
						ColumnName: "count",
						Width:      3,
						DateType:   "INTEGER",
					},
				},
			},
			[]Data{
				{
					ColumnName: "name",
					Value:      "Diabetes",
				},
				{
					ColumnName: "valid",
					Value:      true,
				},
				{
					ColumnName: "count",
					Value:      1,
				},
			},
			false,
		},
		{
			"empty line",
			args{
				line: "",
				formats: []Format{
					{
						ColumnName: "name",
						Width:      10,
						DateType:   "TEXT",
					},
					{
						ColumnName: "valid",
						Width:      1,
						DateType:   "BOOLEAN",
					},
					{
						ColumnName: "count",
						Width:      3,
						DateType:   "INTEGER",
					},
				},
			},
			[]Data{},
			false,
		},
		{
			"invalid data",
			args{
				line: "Diabetes 1  1",
				formats: []Format{
					{
						ColumnName: "name",
						Width:      10,
						DateType:   "TEXT",
					},
					{
						ColumnName: "valid",
						Width:      1,
						DateType:   "BOOLEAN",
					},
					{
						ColumnName: "count",
						Width:      3,
						DateType:   "INTEGER",
					},
				},
			},
			[]Data{},
			false,
		},
		{
			"invalid data",
			args{
				line: "Diabetes   1  1",
				formats: []Format{
					{
						ColumnName: "name",
						Width:      10,
						DateType:   "TEXT",
					},
					{
						ColumnName: "valid",
						Width:      1,
						DateType:   "BOOLEAN",
					},
					{
						ColumnName: "count",
						Width:      3,
						DateType:   "INTEGER",
					},
				},
			},
			[]Data{},
			false,
		},
		{
			"invalid data",
			args{
				line: "Diabetes  1",
				formats: []Format{
					{
						ColumnName: "name",
						Width:      10,
						DateType:   "TEXT",
					},
					{
						ColumnName: "valid",
						Width:      1,
						DateType:   "BOOLEAN",
					},
					{
						ColumnName: "count",
						Width:      3,
						DateType:   "INTEGER",
					},
				},
			},
			[]Data{},
			false,
		},
		{
			"invalid data",
			args{
				line: "Diabetes  -  1",
				formats: []Format{
					{
						ColumnName: "name",
						Width:      10,
						DateType:   "TEXT",
					},
					{
						ColumnName: "valid",
						Width:      1,
						DateType:   "BOOLEAN",
					},
					{
						ColumnName: "count",
						Width:      3,
						DateType:   "INTEGER",
					},
				},
			},
			[]Data{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readDataLine(tt.args.line, tt.args.formats)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDataLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDataLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
