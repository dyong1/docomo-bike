package http

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-chi/chi"
)

func bindURLParams(params interface{}, r *http.Request) error {
	t := reflect.TypeOf(params).Elem()
	s := reflect.ValueOf(params).Elem()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fv := s.Field(i)
		if v, ok := ft.Tag.Lookup("urlParam"); ok {
			if !fv.IsValid() {
				continue
			}
			if !fv.CanSet() {
				continue
			}
			val := chi.URLParam(r, v)
			if ft.Type.Kind() == reflect.String {
				fv.SetString(val)
				continue
			}
			if ft.Type.Kind() == reflect.Int {
				ival, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				fv.SetInt(int64(ival))
				continue
			}
			if ft.Type.Kind() == reflect.Int64 {
				ival, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					return err
				}
				fv.SetInt(ival)
				continue
			}
			fv.Set(reflect.ValueOf(val))
		}
	}
	return nil
}
