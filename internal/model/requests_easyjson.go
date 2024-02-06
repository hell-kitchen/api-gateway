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

func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel(in *jlexer.Lexer, out *UsersUnsubscribeRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel(out *jwriter.Writer, in UsersUnsubscribeRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersUnsubscribeRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersUnsubscribeRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersUnsubscribeRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersUnsubscribeRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel1(in *jlexer.Lexer, out *UsersSubscribeRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel1(out *jwriter.Writer, in UsersSubscribeRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersSubscribeRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersSubscribeRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersSubscribeRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersSubscribeRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel1(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel2(in *jlexer.Lexer, out *UsersSetPasswordRequest) {
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
		case "new_password":
			out.NewPassword = string(in.String())
		case "current_password":
			out.CurrentPassword = string(in.String())
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel2(out *jwriter.Writer, in UsersSetPasswordRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"new_password\":"
		out.RawString(prefix[1:])
		out.String(string(in.NewPassword))
	}
	{
		const prefix string = ",\"current_password\":"
		out.RawString(prefix)
		out.String(string(in.CurrentPassword))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersSetPasswordRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersSetPasswordRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersSetPasswordRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersSetPasswordRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel2(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel3(in *jlexer.Lexer, out *UsersGetSubscriptionsRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel3(out *jwriter.Writer, in UsersGetSubscriptionsRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersGetSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersGetSubscriptionsRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersGetSubscriptionsRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersGetSubscriptionsRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel3(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel4(in *jlexer.Lexer, out *UsersGetByIDRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel4(out *jwriter.Writer, in UsersGetByIDRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersGetByIDRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersGetByIDRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersGetByIDRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersGetByIDRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel4(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel5(in *jlexer.Lexer, out *UsersGetAllRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel5(out *jwriter.Writer, in UsersGetAllRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersGetAllRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersGetAllRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersGetAllRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersGetAllRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel5(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel6(in *jlexer.Lexer, out *UsersCreateRequest) {
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
		case "username":
			out.Username = string(in.String())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel6(out *jwriter.Writer, in UsersCreateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
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
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersCreateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersCreateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersCreateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersCreateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel6(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel7(in *jlexer.Lexer, out *TokensLoginRequest) {
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
		case "password":
			out.Password = string(in.String())
		case "email":
			out.Email = string(in.String())
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel7(out *jwriter.Writer, in TokensLoginRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix[1:])
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TokensLoginRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TokensLoginRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TokensLoginRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TokensLoginRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel7(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel8(in *jlexer.Lexer, out *TagsGetByIDRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel8(out *jwriter.Writer, in TagsGetByIDRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TagsGetByIDRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TagsGetByIDRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TagsGetByIDRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TagsGetByIDRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel8(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel9(in *jlexer.Lexer, out *RecipesUpdateByIDRequest) {
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
		case "ingredients":
			if in.IsNull() {
				in.Skip()
				out.Ingredients = nil
			} else {
				in.Delim('[')
				if out.Ingredients == nil {
					if !in.IsDelim(']') {
						out.Ingredients = make([]IngredientInRecipeDTO, 0, 2)
					} else {
						out.Ingredients = []IngredientInRecipeDTO{}
					}
				} else {
					out.Ingredients = (out.Ingredients)[:0]
				}
				for !in.IsDelim(']') {
					var v1 IngredientInRecipeDTO
					(v1).UnmarshalEasyJSON(in)
					out.Ingredients = append(out.Ingredients, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]int, 0, 8)
					} else {
						out.Tags = []int{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v2 int
					v2 = int(in.Int())
					out.Tags = append(out.Tags, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "image":
			out.Image = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "cooking_time":
			out.CookingTime = int(in.Int())
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel9(out *jwriter.Writer, in RecipesUpdateByIDRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ingredients\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Ingredients == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Ingredients {
				if v3 > 0 {
					out.RawByte(',')
				}
				(v4).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Tags {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"cooking_time\":"
		out.RawString(prefix)
		out.Int(int(in.CookingTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipesUpdateByIDRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipesUpdateByIDRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipesUpdateByIDRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipesUpdateByIDRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel9(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel10(in *jlexer.Lexer, out *RecipesRemoveRecipeFromShoppingCartRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel10(out *jwriter.Writer, in RecipesRemoveRecipeFromShoppingCartRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipesRemoveRecipeFromShoppingCartRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipesRemoveRecipeFromShoppingCartRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipesRemoveRecipeFromShoppingCartRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipesRemoveRecipeFromShoppingCartRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel10(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel11(in *jlexer.Lexer, out *RecipesRemoveRecipeFromFavoriteRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel11(out *jwriter.Writer, in RecipesRemoveRecipeFromFavoriteRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipesRemoveRecipeFromFavoriteRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipesRemoveRecipeFromFavoriteRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipesRemoveRecipeFromFavoriteRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipesRemoveRecipeFromFavoriteRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel11(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel12(in *jlexer.Lexer, out *RecipesGetByIDRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel12(out *jwriter.Writer, in RecipesGetByIDRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipesGetByIDRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipesGetByIDRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipesGetByIDRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipesGetByIDRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel12(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel13(in *jlexer.Lexer, out *RecipesGetAllRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel13(out *jwriter.Writer, in RecipesGetAllRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipesGetAllRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipesGetAllRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipesGetAllRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipesGetAllRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel13(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel14(in *jlexer.Lexer, out *RecipesDeleteByIDRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel14(out *jwriter.Writer, in RecipesDeleteByIDRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipesDeleteByIDRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipesDeleteByIDRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipesDeleteByIDRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipesDeleteByIDRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel14(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel15(in *jlexer.Lexer, out *RecipesCreateRequest) {
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
		case "ingredients":
			if in.IsNull() {
				in.Skip()
				out.Ingredients = nil
			} else {
				in.Delim('[')
				if out.Ingredients == nil {
					if !in.IsDelim(']') {
						out.Ingredients = make([]IngredientInRecipeDTO, 0, 2)
					} else {
						out.Ingredients = []IngredientInRecipeDTO{}
					}
				} else {
					out.Ingredients = (out.Ingredients)[:0]
				}
				for !in.IsDelim(']') {
					var v7 IngredientInRecipeDTO
					(v7).UnmarshalEasyJSON(in)
					out.Ingredients = append(out.Ingredients, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]int, 0, 8)
					} else {
						out.Tags = []int{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v8 int
					v8 = int(in.Int())
					out.Tags = append(out.Tags, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "image":
			out.Image = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "cooking_time":
			out.CookingTime = int(in.Int())
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel15(out *jwriter.Writer, in RecipesCreateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ingredients\":"
		out.RawString(prefix[1:])
		if in.Ingredients == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Ingredients {
				if v9 > 0 {
					out.RawByte(',')
				}
				(v10).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Tags {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v12))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"cooking_time\":"
		out.RawString(prefix)
		out.Int(int(in.CookingTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipesCreateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel15(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipesCreateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel15(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipesCreateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel15(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipesCreateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel15(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel16(in *jlexer.Lexer, out *RecipesAddRecipeToShoppingCartRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel16(out *jwriter.Writer, in RecipesAddRecipeToShoppingCartRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipesAddRecipeToShoppingCartRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel16(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipesAddRecipeToShoppingCartRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel16(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipesAddRecipeToShoppingCartRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel16(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipesAddRecipeToShoppingCartRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel16(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel17(in *jlexer.Lexer, out *RecipesAddRecipeToFavoriteRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel17(out *jwriter.Writer, in RecipesAddRecipeToFavoriteRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipesAddRecipeToFavoriteRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel17(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipesAddRecipeToFavoriteRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel17(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipesAddRecipeToFavoriteRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel17(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipesAddRecipeToFavoriteRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel17(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel18(in *jlexer.Lexer, out *IngredientsGetByIDRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel18(out *jwriter.Writer, in IngredientsGetByIDRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IngredientsGetByIDRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel18(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IngredientsGetByIDRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel18(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IngredientsGetByIDRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel18(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IngredientsGetByIDRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel18(l, v)
}
func easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel19(in *jlexer.Lexer, out *IngredientsGetAllRequest) {
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
func easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel19(out *jwriter.Writer, in IngredientsGetAllRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IngredientsGetAllRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel19(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IngredientsGetAllRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComHellKitchenApiGatewayInternalModel19(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IngredientsGetAllRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel19(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IngredientsGetAllRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComHellKitchenApiGatewayInternalModel19(l, v)
}
