// Code generated by protoc-gen-gogo.
// source: raft_internal.proto
// DO NOT EDIT!

package etcdserverpb

import (
	"fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An InternalRaftRequest is the union of all requests which can be
// sent via raft.
type InternalRaftRequest struct {
	ID                     uint64                         `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	V2                     *Request                       `protobuf:"bytes,2,opt,name=v2" json:"v2,omitempty"`
	Range                  *RangeRequest                  `protobuf:"bytes,3,opt,name=range" json:"range,omitempty"`
	Put                    *PutRequest                    `protobuf:"bytes,4,opt,name=put" json:"put,omitempty"`
	DeleteRange            *DeleteRangeRequest            `protobuf:"bytes,5,opt,name=delete_range,json=deleteRange" json:"delete_range,omitempty"`
	Txn                    *TxnRequest                    `protobuf:"bytes,6,opt,name=txn" json:"txn,omitempty"`
	Compaction             *CompactionRequest             `protobuf:"bytes,7,opt,name=compaction" json:"compaction,omitempty"`
	LeaseGrant             *LeaseGrantRequest             `protobuf:"bytes,8,opt,name=lease_grant,json=leaseGrant" json:"lease_grant,omitempty"`
	LeaseRevoke            *LeaseRevokeRequest            `protobuf:"bytes,9,opt,name=lease_revoke,json=leaseRevoke" json:"lease_revoke,omitempty"`
	AuthEnable             *AuthEnableRequest             `protobuf:"bytes,10,opt,name=auth_enable,json=authEnable" json:"auth_enable,omitempty"`
	AuthDisable            *AuthDisableRequest            `protobuf:"bytes,11,opt,name=auth_disable,json=authDisable" json:"auth_disable,omitempty"`
	AuthUserAdd            *AuthUserAddRequest            `protobuf:"bytes,12,opt,name=auth_user_add,json=authUserAdd" json:"auth_user_add,omitempty"`
	AuthUserDelete         *AuthUserDeleteRequest         `protobuf:"bytes,13,opt,name=auth_user_delete,json=authUserDelete" json:"auth_user_delete,omitempty"`
	AuthUserChangePassword *AuthUserChangePasswordRequest `protobuf:"bytes,14,opt,name=auth_user_change_password,json=authUserChangePassword" json:"auth_user_change_password,omitempty"`
	AuthUserGrant          *AuthUserGrantRequest          `protobuf:"bytes,15,opt,name=auth_user_grant,json=authUserGrant" json:"auth_user_grant,omitempty"`
	AuthRoleAdd            *AuthRoleAddRequest            `protobuf:"bytes,16,opt,name=auth_role_add,json=authRoleAdd" json:"auth_role_add,omitempty"`
	AuthRoleGrant          *AuthRoleGrantRequest          `protobuf:"bytes,17,opt,name=auth_role_grant,json=authRoleGrant" json:"auth_role_grant,omitempty"`
	Authenticate           *AuthenticateRequest           `protobuf:"bytes,18,opt,name=authenticate" json:"authenticate,omitempty"`
	Alarm                  *AlarmRequest                  `protobuf:"bytes,19,opt,name=alarm" json:"alarm,omitempty"`
}

func (m *InternalRaftRequest) Reset()                    { *m = InternalRaftRequest{} }
func (m *InternalRaftRequest) String() string            { return proto.CompactTextString(m) }
func (*InternalRaftRequest) ProtoMessage()               {}
func (*InternalRaftRequest) Descriptor() ([]byte, []int) { return fileDescriptorRaftInternal, []int{0} }

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptorRaftInternal, []int{1} }

func init() {
	proto.RegisterType((*InternalRaftRequest)(nil), "etcdserverpb.InternalRaftRequest")
	proto.RegisterType((*EmptyResponse)(nil), "etcdserverpb.EmptyResponse")
}
func (m *InternalRaftRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *InternalRaftRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.ID))
	}
	if m.V2 != nil {
		data[i] = 0x12
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.V2.Size()))
		n1, err := m.V2.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Range != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Range.Size()))
		n2, err := m.Range.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Put != nil {
		data[i] = 0x22
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Put.Size()))
		n3, err := m.Put.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.DeleteRange != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.DeleteRange.Size()))
		n4, err := m.DeleteRange.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Txn != nil {
		data[i] = 0x32
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Txn.Size()))
		n5, err := m.Txn.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Compaction != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Compaction.Size()))
		n6, err := m.Compaction.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.LeaseGrant != nil {
		data[i] = 0x42
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.LeaseGrant.Size()))
		n7, err := m.LeaseGrant.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.LeaseRevoke != nil {
		data[i] = 0x4a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.LeaseRevoke.Size()))
		n8, err := m.LeaseRevoke.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.AuthEnable != nil {
		data[i] = 0x52
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.AuthEnable.Size()))
		n9, err := m.AuthEnable.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.AuthDisable != nil {
		data[i] = 0x5a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.AuthDisable.Size()))
		n10, err := m.AuthDisable.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.AuthUserAdd != nil {
		data[i] = 0x62
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.AuthUserAdd.Size()))
		n11, err := m.AuthUserAdd.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	if m.AuthUserDelete != nil {
		data[i] = 0x6a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.AuthUserDelete.Size()))
		n12, err := m.AuthUserDelete.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	if m.AuthUserChangePassword != nil {
		data[i] = 0x72
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.AuthUserChangePassword.Size()))
		n13, err := m.AuthUserChangePassword.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	if m.AuthUserGrant != nil {
		data[i] = 0x7a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.AuthUserGrant.Size()))
		n14, err := m.AuthUserGrant.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n14
	}
	if m.AuthRoleAdd != nil {
		data[i] = 0x82
		i++
		data[i] = 0x1
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.AuthRoleAdd.Size()))
		n15, err := m.AuthRoleAdd.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n15
	}
	if m.AuthRoleGrant != nil {
		data[i] = 0x8a
		i++
		data[i] = 0x1
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.AuthRoleGrant.Size()))
		n16, err := m.AuthRoleGrant.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n16
	}
	if m.Authenticate != nil {
		data[i] = 0x92
		i++
		data[i] = 0x1
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Authenticate.Size()))
		n17, err := m.Authenticate.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n17
	}
	if m.Alarm != nil {
		data[i] = 0x9a
		i++
		data[i] = 0x1
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Alarm.Size()))
		n18, err := m.Alarm.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n18
	}
	return i, nil
}

func (m *EmptyResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EmptyResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64RaftInternal(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32RaftInternal(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRaftInternal(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *InternalRaftRequest) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRaftInternal(uint64(m.ID))
	}
	if m.V2 != nil {
		l = m.V2.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Put != nil {
		l = m.Put.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.DeleteRange != nil {
		l = m.DeleteRange.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Txn != nil {
		l = m.Txn.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Compaction != nil {
		l = m.Compaction.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseGrant != nil {
		l = m.LeaseGrant.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseRevoke != nil {
		l = m.LeaseRevoke.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthEnable != nil {
		l = m.AuthEnable.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthDisable != nil {
		l = m.AuthDisable.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserAdd != nil {
		l = m.AuthUserAdd.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserDelete != nil {
		l = m.AuthUserDelete.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserChangePassword != nil {
		l = m.AuthUserChangePassword.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserGrant != nil {
		l = m.AuthUserGrant.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleAdd != nil {
		l = m.AuthRoleAdd.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleGrant != nil {
		l = m.AuthRoleGrant.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.Authenticate != nil {
		l = m.Authenticate.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.Alarm != nil {
		l = m.Alarm.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	return n
}

func (m *EmptyResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovRaftInternal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRaftInternal(x uint64) (n int) {
	return sovRaftInternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InternalRaftRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InternalRaftRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalRaftRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.V2 == nil {
				m.V2 = &Request{}
			}
			if err := m.V2.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Range == nil {
				m.Range = &RangeRequest{}
			}
			if err := m.Range.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Put", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Put == nil {
				m.Put = &PutRequest{}
			}
			if err := m.Put.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeleteRange == nil {
				m.DeleteRange = &DeleteRangeRequest{}
			}
			if err := m.DeleteRange.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Txn == nil {
				m.Txn = &TxnRequest{}
			}
			if err := m.Txn.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compaction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Compaction == nil {
				m.Compaction = &CompactionRequest{}
			}
			if err := m.Compaction.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseGrant", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseGrant == nil {
				m.LeaseGrant = &LeaseGrantRequest{}
			}
			if err := m.LeaseGrant.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseRevoke", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseRevoke == nil {
				m.LeaseRevoke = &LeaseRevokeRequest{}
			}
			if err := m.LeaseRevoke.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthEnable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthEnable == nil {
				m.AuthEnable = &AuthEnableRequest{}
			}
			if err := m.AuthEnable.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthDisable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthDisable == nil {
				m.AuthDisable = &AuthDisableRequest{}
			}
			if err := m.AuthDisable.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserAdd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserAdd == nil {
				m.AuthUserAdd = &AuthUserAddRequest{}
			}
			if err := m.AuthUserAdd.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserDelete", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserDelete == nil {
				m.AuthUserDelete = &AuthUserDeleteRequest{}
			}
			if err := m.AuthUserDelete.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserChangePassword", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserChangePassword == nil {
				m.AuthUserChangePassword = &AuthUserChangePasswordRequest{}
			}
			if err := m.AuthUserChangePassword.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserGrant", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserGrant == nil {
				m.AuthUserGrant = &AuthUserGrantRequest{}
			}
			if err := m.AuthUserGrant.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleAdd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthRoleAdd == nil {
				m.AuthRoleAdd = &AuthRoleAddRequest{}
			}
			if err := m.AuthRoleAdd.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleGrant", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthRoleGrant == nil {
				m.AuthRoleGrant = &AuthRoleGrantRequest{}
			}
			if err := m.AuthRoleGrant.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authenticate == nil {
				m.Authenticate = &AuthenticateRequest{}
			}
			if err := m.Authenticate.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alarm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Alarm == nil {
				m.Alarm = &AlarmRequest{}
			}
			if err := m.Alarm.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmptyResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmptyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRaftInternal(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRaftInternal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaftInternal
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRaftInternal(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRaftInternal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaftInternal   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorRaftInternal = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x69, 0xf7, 0x07, 0xea, 0xfe, 0xc5, 0x05, 0x64, 0x7a, 0x31, 0x46, 0x11, 0x12, 0x02,
	0xa9, 0xa0, 0xf1, 0x00, 0x50, 0xda, 0x0a, 0x0d, 0x81, 0x34, 0x45, 0x70, 0x1d, 0xb9, 0xc9, 0x59,
	0x57, 0x91, 0x26, 0xc1, 0x71, 0xcb, 0x78, 0x39, 0xae, 0x77, 0xc9, 0x23, 0x00, 0x4f, 0x32, 0xfb,
	0x38, 0x71, 0x9a, 0xcd, 0xbd, 0x88, 0x94, 0x7c, 0xe7, 0x3b, 0xbf, 0xf3, 0xd9, 0xb1, 0x4c, 0xfa,
	0x82, 0x9f, 0x4b, 0x7f, 0x19, 0x4b, 0x10, 0x31, 0x8f, 0x46, 0xa9, 0x48, 0x64, 0x42, 0x5b, 0x20,
	0x83, 0x30, 0x03, 0xb1, 0x01, 0x91, 0xce, 0x07, 0x0f, 0x16, 0xc9, 0x22, 0xc1, 0xc2, 0x6b, 0xfd,
	0x66, 0x3c, 0x83, 0x5e, 0xe9, 0xc9, 0x95, 0x86, 0x48, 0x03, 0xf3, 0x3a, 0xfc, 0xdd, 0x20, 0xfd,
	0xd3, 0x9c, 0xe9, 0xa9, 0x01, 0x1e, 0xfc, 0x58, 0x43, 0x26, 0x69, 0x87, 0xd4, 0x4f, 0xa7, 0xac,
	0x76, 0x5c, 0x7b, 0xb1, 0xef, 0xd5, 0x97, 0x53, 0xfa, 0x9c, 0xd4, 0x37, 0x27, 0xac, 0xae, 0xbe,
	0x9b, 0x27, 0x0f, 0x47, 0xdb, 0x53, 0x47, 0x79, 0x8b, 0xa7, 0x0c, 0xf4, 0x0d, 0x39, 0x10, 0x3c,
	0x5e, 0x00, 0xdb, 0x43, 0xe7, 0xe0, 0x86, 0x53, 0x97, 0x0a, 0xbb, 0x31, 0xd2, 0x97, 0x64, 0x2f,
	0x5d, 0x4b, 0xb6, 0x8f, 0x7e, 0x56, 0xf5, 0x9f, 0xad, 0x8b, 0x3c, 0x9e, 0x36, 0xd1, 0x09, 0x69,
	0x85, 0x10, 0x81, 0x04, 0xdf, 0x0c, 0x39, 0xc0, 0xa6, 0xe3, 0x6a, 0xd3, 0x14, 0x1d, 0x95, 0x51,
	0xcd, 0xb0, 0xd4, 0xf4, 0x40, 0x79, 0x19, 0xb3, 0x43, 0xd7, 0xc0, 0xaf, 0x97, 0xb1, 0x1d, 0xa8,
	0x4c, 0xf4, 0x1d, 0x21, 0x41, 0xb2, 0x4a, 0x79, 0x20, 0x97, 0x49, 0xcc, 0xee, 0x62, 0xcb, 0x93,
	0x6a, 0xcb, 0xc4, 0xd6, 0x8b, 0xce, 0xad, 0x16, 0xfa, 0x9e, 0x34, 0x23, 0xe0, 0x19, 0xf8, 0x0b,
	0x95, 0x58, 0xb2, 0x7b, 0x2e, 0xc2, 0x67, 0x6d, 0xf8, 0xa8, 0xeb, 0x96, 0x10, 0x59, 0x49, 0xaf,
	0xd9, 0x10, 0x04, 0x6c, 0x92, 0xef, 0xc0, 0x1a, 0xae, 0x35, 0x23, 0xc2, 0x43, 0x83, 0x5d, 0x73,
	0x54, 0x6a, 0x3a, 0x06, 0x5f, 0xcb, 0x0b, 0x1f, 0x62, 0x3e, 0x8f, 0x80, 0x11, 0x57, 0x8c, 0xb1,
	0x32, 0xcc, 0xb0, 0x6e, 0x63, 0x70, 0x2b, 0xe9, 0x18, 0x48, 0x08, 0x97, 0x19, 0x22, 0x9a, 0xae,
	0x18, 0x1a, 0x31, 0x35, 0x06, 0x1b, 0x83, 0x97, 0x1a, 0x9d, 0x92, 0x36, 0x42, 0xd6, 0xaa, 0xc3,
	0xe7, 0x61, 0xc8, 0x5a, 0xbb, 0x28, 0xdf, 0xd4, 0xd7, 0x38, 0x0c, 0x2b, 0x94, 0x5c, 0xa3, 0x5f,
	0x48, 0xaf, 0xa4, 0x98, 0x3f, 0xcb, 0xda, 0x08, 0x7a, 0xe6, 0x06, 0xe5, 0x27, 0x22, 0x67, 0x75,
	0x78, 0x45, 0xa6, 0xe7, 0xe4, 0x71, 0x89, 0x0b, 0x2e, 0xf4, 0x19, 0xf1, 0x53, 0x9e, 0x65, 0x3f,
	0x13, 0x11, 0xb2, 0x0e, 0x72, 0x5f, 0xb9, 0xb9, 0x13, 0x34, 0x9f, 0xe5, 0xde, 0x82, 0xff, 0x88,
	0x3b, 0xcb, 0xf4, 0x13, 0xe9, 0x96, 0x73, 0xcc, 0x71, 0xe8, 0x22, 0x7d, 0xe8, 0xa6, 0x57, 0x4e,
	0x44, 0x9b, 0x6f, 0xab, 0x76, 0x23, 0x45, 0x12, 0x01, 0x6e, 0x64, 0x6f, 0xd7, 0x46, 0x7a, 0xca,
	0x71, 0x73, 0x23, 0x73, 0xcd, 0x26, 0x42, 0x8a, 0x49, 0x74, 0x7f, 0x57, 0x22, 0xdd, 0x73, 0x3b,
	0x91, 0x55, 0xe9, 0xcc, 0x9c, 0x0f, 0x88, 0xe5, 0x32, 0xe0, 0xea, 0x87, 0x50, 0x04, 0x3d, 0xbd,
	0x0d, 0x2a, 0x1c, 0x05, 0xa7, 0xd2, 0xa6, 0xef, 0x0f, 0x1e, 0x71, 0xb1, 0x62, 0x7d, 0xd7, 0xfd,
	0x31, 0xd6, 0x25, 0x7b, 0x7f, 0xa0, 0x71, 0xd8, 0x25, 0xed, 0xd9, 0x2a, 0x95, 0xbf, 0x3c, 0xc8,
	0xd2, 0x24, 0xce, 0xe0, 0x43, 0xef, 0xea, 0xdf, 0xd1, 0x9d, 0xab, 0xff, 0x47, 0xb5, 0x3f, 0xea,
	0xf9, 0xab, 0x9e, 0xf9, 0x21, 0x5e, 0x75, 0x6f, 0xaf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x6d,
	0xd8, 0x5b, 0x42, 0x05, 0x00, 0x00,
}
