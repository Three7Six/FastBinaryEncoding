// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumUInt32 field model
type FieldModelEnumUInt32 struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int
}

// Create a new EnumUInt32 field model
func NewFieldModelEnumUInt32(buffer *fbe.Buffer, offset int) *FieldModelEnumUInt32 {
    return &FieldModelEnumUInt32{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelEnumUInt32) FBESize() int { return 4 }
// Get the field extra size
func (fm *FieldModelEnumUInt32) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelEnumUInt32) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumUInt32) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumUInt32) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumUInt32) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelEnumUInt32) Verify() bool { return true }

// Get the value
func (fm *FieldModelEnumUInt32) Get() (*EnumUInt32, error) {
    var value EnumUInt32
    return &value, fm.GetValueDefault(&value, EnumUInt32(0))
}

// Get the value with provided default value
func (fm *FieldModelEnumUInt32) GetDefault(defaults EnumUInt32) (*EnumUInt32, error) {
    var value EnumUInt32
    err := fm.GetValueDefault(&value, defaults)
    return &value, err
}

// Get the value by the given pointer
func (fm *FieldModelEnumUInt32) GetValue(value *EnumUInt32) error {
    return fm.GetValueDefault(value, EnumUInt32(0))
}

// Get the value by the given pointer with provided default value
func (fm *FieldModelEnumUInt32) GetValueDefault(value *EnumUInt32, defaults EnumUInt32) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return nil
    }

    *value = EnumUInt32(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return nil
}

// Set the value by the given pointer
func (fm *FieldModelEnumUInt32) Set(value *EnumUInt32) error {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FieldModelEnumUInt32) SetValue(value EnumUInt32) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(value))
    return nil
}
