package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	v "gopkg.in/go-playground/validator.v9"
)

type fieldLevelMock struct {
	top    reflect.Value
	parent reflect.Value
	field  reflect.Value
	param  reflect.Value
	str    string
}

func (f *fieldLevelMock) Top() reflect.Value {
	return f.top
}

// returns the current fields parent struct, if any or
// the comparison value if called 'VarWithValue'
func (f *fieldLevelMock) Parent() reflect.Value {
	return f.parent
}

// returns current field for validation
func (f *fieldLevelMock) Field() reflect.Value {
	return f.field
}

// returns the field's name with the tag
// name takeing precedence over the fields actual name.
func (f *fieldLevelMock) FieldName() string {
	return f.str
}

// returns the struct field's name
func (f *fieldLevelMock) StructFieldName() string {
	return f.str
}

// returns param for validation ag)ainst current field
func (f *fieldLevelMock) Param() string {
	return f.param.String()
}

func (f *fieldLevelMock) ExtractType(field reflect.Value) (value reflect.Value, kind reflect.Kind, nullable bool) {
	return field, field.Kind(), false
}
func (f *fieldLevelMock) GetStructFieldOK() (reflect.Value, reflect.Kind, bool) {
	return f.param, f.param.Kind(), false
}

func TestRequiredIf(t *testing.T) {

	type args struct {
		fl validator.FieldLevel
	}
	var nilSlice []byte
	tests := []struct {
		caseName string
		args     args
		want     bool
	}{
		{
			"test return true if two params are not empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("value1"),
					param: reflect.ValueOf("value2"),
				},
			},
			true,
		},
		{
			"test return true if one of params is not empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(""),
					param: reflect.ValueOf("value2"),
				},
			},
			true,
		},
		{
			"test return false if both the params are empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(""),
					param: reflect.ValueOf(""),
				},
			},
			false,
		},
		{
			"test return true if one of params are empty and slice",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(nilSlice),
					param: reflect.ValueOf(""),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			if got := RequiredIf(tt.args.fl); got != tt.want {
				t.Errorf("RequiredIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceEq(t *testing.T) {

	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		caseName string
		args     args
		want     bool
	}{
		{
			"test return true if field contains param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf([]string{"a", "b"}),
					param: reflect.ValueOf("a@b"),
				},
			},
			true,
		},
		{
			"test return true if field contains param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf([]string{"a", "b"}),
					param: reflect.ValueOf("a@b@c"),
				},
			},
			true,
		},
		{
			"test return true if field not contains param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf([]string{"a", "d"}),
					param: reflect.ValueOf("a@b@c"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			if got := SliceEq(tt.args.fl); got != tt.want {
				t.Errorf("SliceEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringGt(t *testing.T) {

	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		caseName string
		args     args
		want     bool
	}{
		{
			"test return true if field > param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("1"),
					param: reflect.ValueOf("0"),
				},
			},
			true,
		},
		{
			"test return true if field = param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("1"),
					param: reflect.ValueOf("1"),
				},
			},
			false,
		},
		{
			"test return false if both the params are empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(""),
					param: reflect.ValueOf("0"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			if got := StringGt(tt.args.fl); got != tt.want {
				t.Errorf("StringGt() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestStringLt(t *testing.T) {

	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		caseName string
		args     args
		want     bool
	}{
		{
			"test return true if field < param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("-1"),
					param: reflect.ValueOf("0"),
				},
			},
			true,
		},
		{
			"test return true if field = param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("1"),
					param: reflect.ValueOf("1"),
				},
			},
			false,
		},
		{
			"test return false if both the params are empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(""),
					param: reflect.ValueOf("0"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			if got := StringLt(tt.args.fl); got != tt.want {
				t.Errorf("StringGt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsImageName(t *testing.T) {
	type args struct {
		fl v.FieldLevel
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test return true if ImageName has imagename value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("imagename"),
				},
			},
			true,
		},
		{
			"test return true if ImageName has imagename123 value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("imagename123"),
				},
			},
			true,
		},
		{
			"test return true if ImageName has image-name.jpg value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image-name.jpg"),
				},
			},
			true,
		},
		{
			"test return true if ImageName has image-_2name.jpg value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image-_2name.jpg"),
				},
			},
			true,
		},
		{
			"test return true if ImageName has 403540-69 value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("403540-69"),
				},
			},
			true,
		},
		{
			"test return false if ImageName has image# value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image#"),
				},
			},
			false,
		},
		{
			"test return false if ImageName has image@ value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image@"),
				},
			},
			false,
		},
		{
			"test return false if ImageName has image! value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image!"),
				},
			},
			false,
		},
		{
			"test return false if ImageName has image% value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image%"),
				},
			},
			false,
		},
		{
			"test return false if ImageName has image&* value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image&*"),
				},
			},
			false,
		},
		{
			"test return true if false has @$123 value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("@$123"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := isImageName(tt.args.fl)
			{
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
