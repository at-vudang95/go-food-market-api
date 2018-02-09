package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	validator "gopkg.in/go-playground/validator.v9"
)

func TestRegisterCustomValidationRequiredIf(t *testing.T) {
	type exRequest struct {
		Field1 string `validate:"required_if=Field2"`
		Field2 string `validate:"required_if=Field1"`
	}
	tests := []struct {
		name     string
		wantFunc func(*testing.T, *validator.Validate)
	}{
		{
			"requiredIf valid in case including Field1 and Field2",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest{
					"field 1",
					"field 2",
				}
				err := got.Struct(case1)
				assert.Nil(t, err)
			},
		},
		{
			"requiredIf valid in case only Field1 is not empty, Field2 is empty ",
			func(t *testing.T, got *validator.Validate) {
				case2 := exRequest{
					Field1: "field 1",
				}
				err := got.Struct(case2)
				assert.Nil(t, err)
			},
		},
		{
			"requiredIf valid in case Field1 is empty, Field2 is not empty",
			func(t *testing.T, got *validator.Validate) {
				case3 := exRequest{
					Field2: "field 2",
				}
				err := got.Struct(case3)
				assert.Nil(t, err)
			},
		},
		{
			"requiredIf invalid incase Field1 and Field2 are both empty",
			func(t *testing.T, got *validator.Validate) {
				case4 := exRequest{}
				err := got.Struct(case4)
				assert.NotNil(t, err)
			},
		},
	}
	// run testcases one by one
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			tt.wantFunc(t, got)
		})
	}
}

func TestRegisterCustomValidationSliceEq(t *testing.T) {
	type exRequest struct {
		Field1 []string `validate:"s_eq=a@b"`
	}
	type exRequest2 struct {
		Field1 string `validate:"s_eq=a"`
	}
	tests := []struct {
		name     string
		wantFunc func(*testing.T, *validator.Validate)
	}{
		// TODO: Add test cases.
		{
			"s_eq valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest{
					[]string{"a"},
				}
				err := got.Struct(case1)
				assert.Nil(t, err)
			},
		},
		{
			"s_eq valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case2 := exRequest{
					[]string{"b"},
				}
				err := got.Struct(case2)
				assert.Nil(t, err)
			},
		},
		{
			"s_eq valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case2 := exRequest{
					[]string{"a", "b"},
				}
				err := got.Struct(case2)
				assert.Nil(t, err)
			},
		},
		{
			"requiredIf valid in case Field1 is empty, Field2 is not empty",
			func(t *testing.T, got *validator.Validate) {
				case3 := exRequest{
					[]string{"c"},
				}
				err := got.Struct(case3)
				assert.NotNil(t, err)
			},
		},
		{
			"requiredIf valid in case Field1 is empty, Field2 is not empty",
			func(t *testing.T, got *validator.Validate) {
				case3 := exRequest{
					[]string{"a", "b", "c"},
				}
				err := got.Struct(case3)
				assert.NotNil(t, err)
			},
		},
	}
	// run testcases one by one
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			tt.wantFunc(t, got)
		})
	}

	tests2 := []struct {
		name     string
		wantFunc func(*testing.T, *validator.Validate)
	}{
		// TODO: Add test cases.
		{
			"s_eq valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest2{
					"a",
				}
				err := got.Struct(case1)
				assert.NotNil(t, err)
			},
		},
	} // run testcases one by one
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			tt.wantFunc(t, got)
		})
	}

}

func TestRegisterCustomValidationStringGt(t *testing.T) {
	type exRequest struct {
		Field1 string `validate:"string_gt=0"`
	}
	type exRequest2 struct {
		Field1 string `validate:"string_gt=-1"`
	}
	type exRequest3 struct {
		Field1 string `validate:"string_gt=a"`
	}
	type exRequest4 struct {
		Field1 int `validate:"string_gt=0"`
	}
	tests := []struct {
		name     string
		wantFunc func(*testing.T, *validator.Validate)
	}{
		{
			"string_gt valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest{
					"1",
				}
				err := got.Struct(case1)
				assert.Nil(t, err)
			},
		},
		{
			"string_gt valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case2 := exRequest{
					"0",
				}
				err := got.Struct(case2)
				assert.NotNil(t, err)
			},
		},
		{
			"string_gt valid  in case Field1 is a",
			func(t *testing.T, got *validator.Validate) {
				case2 := exRequest{
					"a",
				}
				err := got.Struct(case2)
				assert.NotNil(t, err)
			},
		},
		{
			"string_gt valid  in case Field1 is -1",
			func(t *testing.T, got *validator.Validate) {
				case3 := exRequest{
					"-1",
				}
				err := got.Struct(case3)
				assert.NotNil(t, err)
			},
		},
		{
			"string_gt valid  in case Field1 is empty",
			func(t *testing.T, got *validator.Validate) {
				case3 := exRequest{
					"",
				}
				err := got.Struct(case3)
				assert.NotNil(t, err)
			},
		},
	}
	// run testcases one by one
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			tt.wantFunc(t, got)
		})
	}

	tests2 := []struct {
		name     string
		wantFunc func(*testing.T, *validator.Validate)
	}{
		{
			"string_gt valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest2{
					"0",
				}
				err := got.Struct(case1)
				assert.Nil(t, err)
			},
		},
		{
			"string_gt valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest2{
					"-1",
				}
				err := got.Struct(case1)
				assert.NotNil(t, err)
			},
		},
		{
			"string_gt valid in case including struct params error",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest3{
					"0",
				}
				err := got.Struct(case1)
				assert.NotNil(t, err)
			},
		},
		{
			"string_gt valid in case including struct params error",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest4{
					0,
				}
				err := got.Struct(case1)
				assert.NotNil(t, err)
			},
		},
	} // run testcases one by one
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			tt.wantFunc(t, got)
		})
	}

}

func TestRegisterCustomValidationStringLt(t *testing.T) {
	type exRequest struct {
		Field1 string `validate:"string_lt=0"`
	}
	type exRequest2 struct {
		Field1 string `validate:"string_lt=1"`
	}
	type exRequest3 struct {
		Field1 string `validate:"string_lt=a"`
	}
	type exRequest4 struct {
		Field1 int `validate:"string_lt=0"`
	}
	tests := []struct {
		name     string
		wantFunc func(*testing.T, *validator.Validate)
	}{
		{
			"string_lt valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest{
					"-1",
				}
				err := got.Struct(case1)
				assert.Nil(t, err)
			},
		},
		{
			"string_lt valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case2 := exRequest{
					"0",
				}
				err := got.Struct(case2)
				assert.NotNil(t, err)
			},
		},
		{
			"string_lt valid  in case Field1 is a",
			func(t *testing.T, got *validator.Validate) {
				case2 := exRequest{
					"a",
				}
				err := got.Struct(case2)
				assert.NotNil(t, err)
			},
		},
		{
			"string_lt valid in case including struct int",
			func(t *testing.T, got *validator.Validate) {
				case3 := exRequest4{
					0,
				}
				err := got.Struct(case3)
				assert.NotNil(t, err)
			},
		},
	}
	// run testcases one by one
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			tt.wantFunc(t, got)
		})
	}

	tests2 := []struct {
		name     string
		wantFunc func(*testing.T, *validator.Validate)
	}{
		{
			"string_lt valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest2{
					"0",
				}
				err := got.Struct(case1)
				assert.Nil(t, err)
			},
		},
		{
			"string_lt valid in case including Field1",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest2{
					"0",
				}
				err := got.Struct(case1)
				assert.Nil(t, err)
			},
		},
		{
			"string_lt valid in case including struct params error",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest3{
					"0",
				}
				err := got.Struct(case1)
				assert.NotNil(t, err)
			},
		},
		{
			"string_lt valid in case including struct int",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest4{
					0,
				}
				err := got.Struct(case1)
				assert.NotNil(t, err)
			},
		},
	} // run testcases one by one
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			tt.wantFunc(t, got)
		})
	}

}

func TestRegisterCustomValidationIsImageName(t *testing.T) {
	type exRequest struct {
		Field1 string `validate:"image_name"`
	}
	type exRequest2 struct {
		Field1 int `validate:"image_name"`
	}
	testImNameisString := []struct {
		name     string
		wantFunc func(*testing.T, *validator.Validate)
	}{
		{
			"image_name valid in case including Field1 has imagenamejpg",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest{
					"imagenamejpg",
				}
				err := got.Struct(case1)
				assert.Nil(t, err)
			},
		},
		{
			"image_name valid in case including Field1 has imagename1jpg",
			func(t *testing.T, got *validator.Validate) {
				case2 := exRequest{
					"imagename1jpg",
				}
				err := got.Struct(case2)
				assert.Nil(t, err)
			},
		},
		{
			"image_name valid in case including Field1 has image-name1.jpg",
			func(t *testing.T, got *validator.Validate) {
				case3 := exRequest{
					"image-name1.jpg",
				}
				err := got.Struct(case3)
				assert.Nil(t, err)
			},
		},
		{
			"image_name valid in case including Field1 has 403540-69",
			func(t *testing.T, got *validator.Validate) {
				case4 := exRequest{
					"403540-69",
				}
				err := got.Struct(case4)
				assert.Nil(t, err)
			},
		},
		{
			"image_name invalid in case including Field1 has image@",
			func(t *testing.T, got *validator.Validate) {
				case5 := exRequest{
					"image@",
				}
				err := got.Struct(case5)
				assert.NotNil(t, err)
			},
		},
		{
			"image_name invalid in case including Field1 has image.!#",
			func(t *testing.T, got *validator.Validate) {
				case6 := exRequest{
					"image.!#",
				}
				err := got.Struct(case6)
				assert.NotNil(t, err)
			},
		},
		{
			"image_name invalid in case including Field1 has @$123",
			func(t *testing.T, got *validator.Validate) {
				case7 := exRequest{
					"@$123",
				}
				err := got.Struct(case7)
				assert.NotNil(t, err)
			},
		},
	}
	// run testcases one by one
	for _, tt := range testImNameisString {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			tt.wantFunc(t, got)
		})
	}

	testImNameIsInt := []struct {
		name     string
		wantFunc func(*testing.T, *validator.Validate)
	}{
		{
			"string_lt invalid in case including Field1 has 10",
			func(t *testing.T, got *validator.Validate) {
				case1 := exRequest2{
					10,
				}
				err := got.Struct(case1)
				assert.NotNil(t, err)
			},
		},
	} // run testcases one by one
	for _, tt := range testImNameIsInt {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			tt.wantFunc(t, got)
		})
	}
}

func TestRegisterCustomTagNameByFormTag(t *testing.T) {
	tests := []struct {
		name          string
		exRequest     interface{}
		wantFieldName string
	}{
		{
			"test returned field is form tag value",
			struct {
				FieldName string `form:"field_name" validate:"required"`
			}{},
			"field_name",
		},
		{
			"test returned struct field name when not set form tag",
			struct {
				FieldName string `validate:"required"`
			}{},
			"FieldName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			err := got.Struct(tt.exRequest)
			assert.Equal(t, tt.wantFieldName, err.(validator.ValidationErrors)[0].Field())
		})
	}
}
