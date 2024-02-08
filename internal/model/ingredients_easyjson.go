// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC156496aDecodeGithubComHellKitchenApiGatewayInternalModel(in *jlexer.Lexer, out *IngredientInRecipeDTO) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "measurement_unit":
			out.MeasurementUnit = string(in.String())
		case "amount":
			out.Amount = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC156496aEncodeGithubComHellKitchenApiGatewayInternalModel(out *jwriter.Writer, in IngredientInRecipeDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"measurement_unit\":"
		out.RawString(prefix)
		out.String(string(in.MeasurementUnit))
	}
	{
		const prefix string = ",\"amount\":"
		out.RawString(prefix)
		out.Int(int(in.Amount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IngredientInRecipeDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC156496aEncodeGithubComHellKitchenApiGatewayInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IngredientInRecipeDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC156496aEncodeGithubComHellKitchenApiGatewayInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IngredientInRecipeDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC156496aDecodeGithubComHellKitchenApiGatewayInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IngredientInRecipeDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC156496aDecodeGithubComHellKitchenApiGatewayInternalModel(l, v)
}
func easyjsonC156496aDecodeGithubComHellKitchenApiGatewayInternalModel1(in *jlexer.Lexer, out *IngredientInRecipeCreationDTO) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "amount":
			out.Amount = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC156496aEncodeGithubComHellKitchenApiGatewayInternalModel1(out *jwriter.Writer, in IngredientInRecipeCreationDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"amount\":"
		out.RawString(prefix)
		out.Int(int(in.Amount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IngredientInRecipeCreationDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC156496aEncodeGithubComHellKitchenApiGatewayInternalModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IngredientInRecipeCreationDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC156496aEncodeGithubComHellKitchenApiGatewayInternalModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IngredientInRecipeCreationDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC156496aDecodeGithubComHellKitchenApiGatewayInternalModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IngredientInRecipeCreationDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC156496aDecodeGithubComHellKitchenApiGatewayInternalModel1(l, v)
}
