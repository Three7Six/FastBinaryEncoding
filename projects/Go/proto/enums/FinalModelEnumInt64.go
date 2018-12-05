// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumInt64 final model
type FinalModelEnumInt64 struct {
    // Final model buffer
    buffer *fbe.Buffer
    // Final model buffer offset
    offset int
}

// Create a new EnumInt64 final model
func NewFinalModelEnumInt64(buffer *fbe.Buffer, offset int) *FinalModelEnumInt64 {
    return &FinalModelEnumInt64{buffer: buffer, offset: offset}
}

// Get the allocation size
func (fm *FinalModelEnumInt64) FBEAllocationSize(value *EnumInt64) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelEnumInt64) FBESize() int { return 8 }

// Get the final offset
func (fm *FinalModelEnumInt64) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumInt64) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumInt64) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumInt64) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FinalModelEnumInt64) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return fbe.MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm *FinalModelEnumInt64) Get() (*EnumInt64, int, error) {
    var value EnumInt64
    size, err := fm.GetValueDefault(&value, EnumInt64(0))
    return &value, size, err
}

// Get the value with provided default value
func (fm *FinalModelEnumInt64) GetDefault(defaults EnumInt64) (*EnumInt64, int, error) {
    var value EnumInt64
    size, err := fm.GetValueDefault(&value, defaults)
    return &value, size, err
}

// Get the value by the given pointer
func (fm *FinalModelEnumInt64) GetValue(value *EnumInt64) (int, error) {
    return fm.GetValueDefault(value, EnumInt64(0))
}

// Get the value by the given pointer with provided default value
func (fm *FinalModelEnumInt64) GetValueDefault(value *EnumInt64, defaults EnumInt64) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return 0, errors.New("model is broken")
    }

    *value = EnumInt64(fbe.ReadInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return fm.FBESize(), nil
}

// Set the value by the given pointer
func (fm *FinalModelEnumInt64) Set(value *EnumInt64) (int, error) {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FinalModelEnumInt64) SetValue(value EnumInt64) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int64(value))
    return fm.FBESize(), nil
}
