// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumChar field model
type FieldModelEnumChar struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int
}

// Create a new EnumChar field model
func NewFieldModelEnumChar(buffer *fbe.Buffer, offset int) *FieldModelEnumChar {
    return &FieldModelEnumChar{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelEnumChar) FBESize() int { return 1 }
// Get the field extra size
func (fm *FieldModelEnumChar) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelEnumChar) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumChar) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumChar) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumChar) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelEnumChar) Verify() bool { return true }

// Get the value
func (fm *FieldModelEnumChar) Get() (*EnumChar, error) {
    var value EnumChar
    return &value, fm.GetValueDefault(&value, EnumChar(0))
}

// Get the value with provided default value
func (fm *FieldModelEnumChar) GetDefault(defaults EnumChar) (*EnumChar, error) {
    var value EnumChar
    err := fm.GetValueDefault(&value, defaults)
    return &value, err
}

// Get the value by the given pointer
func (fm *FieldModelEnumChar) GetValue(value *EnumChar) error {
    return fm.GetValueDefault(value, EnumChar(0))
}

// Get the value by the given pointer with provided default value
func (fm *FieldModelEnumChar) GetValueDefault(value *EnumChar, defaults EnumChar) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return nil
    }

    *value = EnumChar(fbe.ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return nil
}

// Set the value by the given pointer
func (fm *FieldModelEnumChar) Set(value *EnumChar) error {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FieldModelEnumChar) SetValue(value EnumChar) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint8(value))
    return nil
}
