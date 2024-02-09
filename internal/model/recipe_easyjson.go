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

func easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel(in *jlexer.Lexer, out *RecipeInUsersSubscriptions) {
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
		case "image":
			out.Image = string(in.String())
		case "cooking_time":
			out.CookingTime = uint32(in.Uint32())
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
func easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel(out *jwriter.Writer, in RecipeInUsersSubscriptions) {
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
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"cooking_time\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.CookingTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipeInUsersSubscriptions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipeInUsersSubscriptions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipeInUsersSubscriptions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipeInUsersSubscriptions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel(l, v)
}
func easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel1(in *jlexer.Lexer, out *RecipeDTO) {
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
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]TagDTO, 0, 1)
					} else {
						out.Tags = []TagDTO{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v1 TagDTO
					easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel2(in, &v1)
					out.Tags = append(out.Tags, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "author":
			easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel3(in, &out.Author)
		case "ingredients":
			if in.IsNull() {
				in.Skip()
				out.Ingredients = nil
			} else {
				in.Delim('[')
				if out.Ingredients == nil {
					if !in.IsDelim(']') {
						out.Ingredients = make([]IngredientInRecipeDTO, 0, 1)
					} else {
						out.Ingredients = []IngredientInRecipeDTO{}
					}
				} else {
					out.Ingredients = (out.Ingredients)[:0]
				}
				for !in.IsDelim(']') {
					var v2 IngredientInRecipeDTO
					(v2).UnmarshalEasyJSON(in)
					out.Ingredients = append(out.Ingredients, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "is_favorited":
			out.IsFavorited = bool(in.Bool())
		case "is_in_shopping_cart":
			out.IsInShoppingCart = bool(in.Bool())
		case "name":
			out.Name = string(in.String())
		case "image":
			out.Image = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "cooking_time":
			out.CookingTime = uint32(in.Uint32())
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
func easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel1(out *jwriter.Writer, in RecipeDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Tags {
				if v3 > 0 {
					out.RawByte(',')
				}
				easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel2(out, v4)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel3(out, in.Author)
	}
	{
		const prefix string = ",\"ingredients\":"
		out.RawString(prefix)
		if in.Ingredients == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Ingredients {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"is_favorited\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsFavorited))
	}
	{
		const prefix string = ",\"is_in_shopping_cart\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsInShoppingCart))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"cooking_time\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.CookingTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipeDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipeDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipeDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipeDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel1(l, v)
}
func easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel3(in *jlexer.Lexer, out *UserInRecipe) {
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
func easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel3(out *jwriter.Writer, in UserInRecipe) {
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
func easyjsonAf04a44aDecodeGithubComHellKitchenApiGatewayInternalModel2(in *jlexer.Lexer, out *TagDTO) {
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
		case "color":
			out.Color = string(in.String())
		case "slug":
			out.Slug = string(in.String())
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
func easyjsonAf04a44aEncodeGithubComHellKitchenApiGatewayInternalModel2(out *jwriter.Writer, in TagDTO) {
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
		const prefix string = ",\"color\":"
		out.RawString(prefix)
		out.String(string(in.Color))
	}
	{
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	out.RawByte('}')
}
