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

func easyjson84c0690eDecodeGithubComHellKitchenApiGatewayInternalModel(in *jlexer.Lexer, out *UserInSubscriptionsResult) {
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
		case "email":
			out.Email = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "is_subscribed":
			out.IsSubscribed = bool(in.Bool())
		case "recipes":
			if in.IsNull() {
				in.Skip()
				out.Recipes = nil
			} else {
				in.Delim('[')
				if out.Recipes == nil {
					if !in.IsDelim(']') {
						out.Recipes = make([]RecipeInUsersSubscriptions, 0, 1)
					} else {
						out.Recipes = []RecipeInUsersSubscriptions{}
					}
				} else {
					out.Recipes = (out.Recipes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 RecipeInUsersSubscriptions
					(v1).UnmarshalEasyJSON(in)
					out.Recipes = append(out.Recipes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "recipes_count":
			out.RecipesCount = int(in.Int())
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
func easyjson84c0690eEncodeGithubComHellKitchenApiGatewayInternalModel(out *jwriter.Writer, in UserInSubscriptionsResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"is_subscribed\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSubscribed))
	}
	{
		const prefix string = ",\"recipes\":"
		out.RawString(prefix)
		if in.Recipes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Recipes {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"recipes_count\":"
		out.RawString(prefix)
		out.Int(int(in.RecipesCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInSubscriptionsResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGithubComHellKitchenApiGatewayInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInSubscriptionsResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGithubComHellKitchenApiGatewayInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInSubscriptionsResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGithubComHellKitchenApiGatewayInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInSubscriptionsResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGithubComHellKitchenApiGatewayInternalModel(l, v)
}
func easyjson84c0690eDecodeGithubComHellKitchenApiGatewayInternalModel1(in *jlexer.Lexer, out *UserInSubscriptions) {
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
		case "email":
			out.Email = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "is_subscribed":
			out.IsSubscribed = bool(in.Bool())
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
func easyjson84c0690eEncodeGithubComHellKitchenApiGatewayInternalModel1(out *jwriter.Writer, in UserInSubscriptions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"is_subscribed\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSubscribed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInSubscriptions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGithubComHellKitchenApiGatewayInternalModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInSubscriptions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGithubComHellKitchenApiGatewayInternalModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInSubscriptions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGithubComHellKitchenApiGatewayInternalModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInSubscriptions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGithubComHellKitchenApiGatewayInternalModel1(l, v)
}
func easyjson84c0690eDecodeGithubComHellKitchenApiGatewayInternalModel2(in *jlexer.Lexer, out *UserInRecipe) {
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
		case "email":
			out.Email = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "is_subscribed":
			out.IsSubscribed = bool(in.Bool())
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
func easyjson84c0690eEncodeGithubComHellKitchenApiGatewayInternalModel2(out *jwriter.Writer, in UserInRecipe) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"is_subscribed\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSubscribed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInRecipe) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGithubComHellKitchenApiGatewayInternalModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInRecipe) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGithubComHellKitchenApiGatewayInternalModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInRecipe) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGithubComHellKitchenApiGatewayInternalModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInRecipe) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGithubComHellKitchenApiGatewayInternalModel2(l, v)
}
