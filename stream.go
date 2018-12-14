/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Aspose">
 *   Copyright (c) 2018 Aspose.Slides for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */

package asposeslidescloud

import (
	"encoding/json"
)


type IStream interface {

	getCanRead() bool
	setCanRead(newValue bool)

	getCanSeek() bool
	setCanSeek(newValue bool)

	getCanTimeout() bool
	setCanTimeout(newValue bool)

	getCanWrite() bool
	setCanWrite(newValue bool)

	getLength() int64
	setLength(newValue int64)

	getPosition() int64
	setPosition(newValue int64)

	getReadTimeout() int32
	setReadTimeout(newValue int32)

	getWriteTimeout() int32
	setWriteTimeout(newValue int32)

	getNull() IStream
	setNull(newValue IStream)
}

type Stream struct {

	CanRead bool `json:"CanRead,omitempty"`

	CanSeek bool `json:"CanSeek,omitempty"`

	CanTimeout bool `json:"CanTimeout,omitempty"`

	CanWrite bool `json:"CanWrite,omitempty"`

	Length int64 `json:"Length,omitempty"`

	Position int64 `json:"Position,omitempty"`

	ReadTimeout int32 `json:"ReadTimeout,omitempty"`

	WriteTimeout int32 `json:"WriteTimeout,omitempty"`

	Null IStream `json:"Null,omitempty"`
}

func (this Stream) getCanRead() bool {
	return this.CanRead
}

func (this Stream) setCanRead(newValue bool) {
	this.CanRead = newValue
}
func (this Stream) getCanSeek() bool {
	return this.CanSeek
}

func (this Stream) setCanSeek(newValue bool) {
	this.CanSeek = newValue
}
func (this Stream) getCanTimeout() bool {
	return this.CanTimeout
}

func (this Stream) setCanTimeout(newValue bool) {
	this.CanTimeout = newValue
}
func (this Stream) getCanWrite() bool {
	return this.CanWrite
}

func (this Stream) setCanWrite(newValue bool) {
	this.CanWrite = newValue
}
func (this Stream) getLength() int64 {
	return this.Length
}

func (this Stream) setLength(newValue int64) {
	this.Length = newValue
}
func (this Stream) getPosition() int64 {
	return this.Position
}

func (this Stream) setPosition(newValue int64) {
	this.Position = newValue
}
func (this Stream) getReadTimeout() int32 {
	return this.ReadTimeout
}

func (this Stream) setReadTimeout(newValue int32) {
	this.ReadTimeout = newValue
}
func (this Stream) getWriteTimeout() int32 {
	return this.WriteTimeout
}

func (this Stream) setWriteTimeout(newValue int32) {
	this.WriteTimeout = newValue
}
func (this Stream) getNull() IStream {
	return this.Null
}

func (this Stream) setNull(newValue IStream) {
	this.Null = newValue
}

func (this *Stream) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valCanRead, ok := objMap["CanRead"]; ok {
		if valCanRead != nil {
			var valueForCanRead bool
			err = json.Unmarshal(*valCanRead, &valueForCanRead)
			if err != nil {
				return err
			}
			this.CanRead = valueForCanRead
		}
	}

	if valCanSeek, ok := objMap["CanSeek"]; ok {
		if valCanSeek != nil {
			var valueForCanSeek bool
			err = json.Unmarshal(*valCanSeek, &valueForCanSeek)
			if err != nil {
				return err
			}
			this.CanSeek = valueForCanSeek
		}
	}

	if valCanTimeout, ok := objMap["CanTimeout"]; ok {
		if valCanTimeout != nil {
			var valueForCanTimeout bool
			err = json.Unmarshal(*valCanTimeout, &valueForCanTimeout)
			if err != nil {
				return err
			}
			this.CanTimeout = valueForCanTimeout
		}
	}

	if valCanWrite, ok := objMap["CanWrite"]; ok {
		if valCanWrite != nil {
			var valueForCanWrite bool
			err = json.Unmarshal(*valCanWrite, &valueForCanWrite)
			if err != nil {
				return err
			}
			this.CanWrite = valueForCanWrite
		}
	}

	if valLength, ok := objMap["Length"]; ok {
		if valLength != nil {
			var valueForLength int64
			err = json.Unmarshal(*valLength, &valueForLength)
			if err != nil {
				return err
			}
			this.Length = valueForLength
		}
	}

	if valPosition, ok := objMap["Position"]; ok {
		if valPosition != nil {
			var valueForPosition int64
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}

	if valReadTimeout, ok := objMap["ReadTimeout"]; ok {
		if valReadTimeout != nil {
			var valueForReadTimeout int32
			err = json.Unmarshal(*valReadTimeout, &valueForReadTimeout)
			if err != nil {
				return err
			}
			this.ReadTimeout = valueForReadTimeout
		}
	}

	if valWriteTimeout, ok := objMap["WriteTimeout"]; ok {
		if valWriteTimeout != nil {
			var valueForWriteTimeout int32
			err = json.Unmarshal(*valWriteTimeout, &valueForWriteTimeout)
			if err != nil {
				return err
			}
			this.WriteTimeout = valueForWriteTimeout
		}
	}

	if valNull, ok := objMap["Null"]; ok {
		if valNull != nil {
			var valueForNull Stream
			err = json.Unmarshal(*valNull, &valueForNull)
			if err != nil {
				return err
			}
			this.Null = valueForNull
		}
	}

    return nil
}
