package packetPaddingContainer

import "encoding/binary"

func New() PacketPaddingContainer {
	return packetPaddingContainer{}
}

type packetPaddingContainer struct {
}

func (c packetPaddingContainer) Pack(data_OWNERSHIP_RELINQUISHED []byte, padding int) []byte {
	data := append(data_OWNERSHIP_RELINQUISHED, make([]byte, padding)...)
	data_length := len(data_OWNERSHIP_RELINQUISHED)
	data = append(data, byte(data_length>>8), byte(data_length))
	return data
}

func (c packetPaddingContainer) Pad(padding int) []byte {
	if assertPaddingLengthIsNotNegative := padding < 0; assertPaddingLengthIsNotNegative {
		return nil
	}
	switch padding {
	case 0:
		return []byte{}
	case 1:
		return []byte{0}
	case 2:
		return []byte{0, 0}
	default:
		return append(make([]byte, padding-2), byte(padding>>8), byte(padding))
	}

}

func (c packetPaddingContainer) Unpack(wrappedData_OWNERSHIP_RELINQUISHED []byte) ([]byte, int) {
	if len(wrappedData_OWNERSHIP_RELINQUISHED) < 2 {
		return nil, len(wrappedData_OWNERSHIP_RELINQUISHED)
	}
	wrappedData_tail := wrappedData_OWNERSHIP_RELINQUISHED[len(wrappedData_OWNERSHIP_RELINQUISHED)-2:]
	dataLength := int(binary.BigEndian.Uint16(wrappedData_tail))
	paddingLength := len(wrappedData_OWNERSHIP_RELINQUISHED) - dataLength - 2
	if paddingLength < 0 {
		return nil, paddingLength
	}
	return wrappedData_OWNERSHIP_RELINQUISHED[:dataLength], paddingLength
}
