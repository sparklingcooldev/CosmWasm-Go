// Code generated by tinyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

import (
	tinyjson "github.com/CosmWasm/tinyjson"
	jlexer "github.com/CosmWasm/tinyjson/jlexer"
	jwriter "github.com/CosmWasm/tinyjson/jwriter"
)

// suppress unused package warning
var (
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ tinyjson.Marshaler
)

func tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes(in *jlexer.Lexer, out *TransactionInfo) {
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
		case "index":
			out.Index = uint32(in.Uint32())
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
func tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes(out *jwriter.Writer, in TransactionInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"index\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.Index))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TransactionInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v TransactionInfo) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TransactionInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *TransactionInfo) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes(l, v)
}
func tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes1(in *jlexer.Lexer, out *MessageInfo) {
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
		case "sender":
			out.Sender = string(in.String())
		case "funds":
			if in.IsNull() {
				in.Skip()
				out.Funds = nil
			} else {
				in.Delim('[')
				if out.Funds == nil {
					if !in.IsDelim(']') {
						out.Funds = make([]Coin, 0, 2)
					} else {
						out.Funds = []Coin{}
					}
				} else {
					out.Funds = (out.Funds)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Coin
					(v1).UnmarshalTinyJSON(in)
					out.Funds = append(out.Funds, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes1(out *jwriter.Writer, in MessageInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"sender\":"
		out.RawString(prefix[1:])
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"funds\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.Funds {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalTinyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessageInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v MessageInfo) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessageInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes1(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *MessageInfo) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes1(l, v)
}
func tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes2(in *jlexer.Lexer, out *Env) {
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
		case "block":
			(out.Block).UnmarshalTinyJSON(in)
		case "contract":
			(out.Contract).UnmarshalTinyJSON(in)
		case "transaction_info":
			if in.IsNull() {
				in.Skip()
				out.Transaction = nil
			} else {
				if out.Transaction == nil {
					out.Transaction = new(TransactionInfo)
				}
				(*out.Transaction).UnmarshalTinyJSON(in)
			}
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
func tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes2(out *jwriter.Writer, in Env) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"block\":"
		out.RawString(prefix[1:])
		(in.Block).MarshalTinyJSON(out)
	}
	{
		const prefix string = ",\"contract\":"
		out.RawString(prefix)
		(in.Contract).MarshalTinyJSON(out)
	}
	if in.Transaction != nil {
		const prefix string = ",\"transaction_info\":"
		out.RawString(prefix)
		(*in.Transaction).MarshalTinyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Env) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v Env) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Env) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes2(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *Env) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes2(l, v)
}
func tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes3(in *jlexer.Lexer, out *ContractInfo) {
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
		case "address":
			out.Address = string(in.String())
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
func tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes3(out *jwriter.Writer, in ContractInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix[1:])
		out.String(string(in.Address))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContractInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v ContractInfo) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContractInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes3(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *ContractInfo) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes3(l, v)
}
func tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes4(in *jlexer.Lexer, out *BlockInfo) {
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
		case "height":
			out.Height = uint64(in.Uint64())
		case "time":
			out.Time = uint64(in.Uint64Str())
		case "chain_id":
			out.ChainID = string(in.String())
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
func tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes4(out *jwriter.Writer, in BlockInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"height\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Height))
	}
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix)
		out.Uint64Str(uint64(in.Time))
	}
	{
		const prefix string = ",\"chain_id\":"
		out.RawString(prefix)
		out.String(string(in.ChainID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BlockInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v BlockInfo) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson2cc1611bEncodeGithubComCosmwasmCosmwasmGoStdTypes4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BlockInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes4(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *BlockInfo) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson2cc1611bDecodeGithubComCosmwasmCosmwasmGoStdTypes4(l, v)
}
