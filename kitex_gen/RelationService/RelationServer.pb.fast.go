// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package RelationService

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FollowCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.FollowerCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.IsFollow, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *UserInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UserInfo[number], err)
}

func (x *UserInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FollowCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.FollowerCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.IsFollow, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Avatar, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.BackgroundImage, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Signature, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.TotalFavorited, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	x.WorkCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UserInfo) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	x.FavoriteCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseResp[number], err)
}

func (x *BaseResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationActionRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationActionRequest[number], err)
}

func (x *RelationActionRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationActionRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationActionRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RelationActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationActionResponse[number], err)
}

func (x *RelationActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *RelationFollowListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFollowListRequest[number], err)
}

func (x *RelationFollowListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationFollowListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationFollowListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFollowListResponse[number], err)
}

func (x *RelationFollowListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *RelationFollowListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v UserInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.FollowList = append(x.FollowList, &v)
	return offset, nil
}

func (x *RelationFollowerListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFollowerListRequest[number], err)
}

func (x *RelationFollowerListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationFollowerListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationFollowerListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFollowerListResponse[number], err)
}

func (x *RelationFollowerListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *RelationFollowerListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v UserInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.FollowerList = append(x.FollowerList, &v)
	return offset, nil
}

func (x *RelationQueryRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationQueryRequest[number], err)
}

func (x *RelationQueryRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationQueryRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationQueryResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationQueryResponse[number], err)
}

func (x *RelationQueryResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *RelationQueryResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.IsFollow, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *QueryUserInfosWithRelationRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryUserInfosWithRelationRequest[number], err)
}

func (x *QueryUserInfosWithRelationRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *QueryUserInfosWithRelationRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.IdList = append(x.IdList, v)
			return offset, err
		})
	return offset, err
}

func (x *QueryUserInfosWithRelationResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryUserInfosWithRelationResponse[number], err)
}

func (x *QueryUserInfosWithRelationResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v UserInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserInfoList = append(x.UserInfoList, &v)
	return offset, nil
}

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.FollowCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetFollowCount())
	return offset
}

func (x *User) fastWriteField4(buf []byte) (offset int) {
	if x.FollowerCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetFollowerCount())
	return offset
}

func (x *User) fastWriteField5(buf []byte) (offset int) {
	if !x.IsFollow {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 5, x.GetIsFollow())
	return offset
}

func (x *UserInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	return offset
}

func (x *UserInfo) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *UserInfo) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *UserInfo) fastWriteField3(buf []byte) (offset int) {
	if x.FollowCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetFollowCount())
	return offset
}

func (x *UserInfo) fastWriteField4(buf []byte) (offset int) {
	if x.FollowerCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetFollowerCount())
	return offset
}

func (x *UserInfo) fastWriteField5(buf []byte) (offset int) {
	if !x.IsFollow {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 5, x.GetIsFollow())
	return offset
}

func (x *UserInfo) fastWriteField6(buf []byte) (offset int) {
	if x.Avatar == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetAvatar())
	return offset
}

func (x *UserInfo) fastWriteField7(buf []byte) (offset int) {
	if x.BackgroundImage == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetBackgroundImage())
	return offset
}

func (x *UserInfo) fastWriteField8(buf []byte) (offset int) {
	if x.Signature == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetSignature())
	return offset
}

func (x *UserInfo) fastWriteField9(buf []byte) (offset int) {
	if x.TotalFavorited == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 9, x.GetTotalFavorited())
	return offset
}

func (x *UserInfo) fastWriteField10(buf []byte) (offset int) {
	if x.WorkCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 10, x.GetWorkCount())
	return offset
}

func (x *UserInfo) fastWriteField11(buf []byte) (offset int) {
	if x.FavoriteCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 11, x.GetFavoriteCount())
	return offset
}

func (x *BaseResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *BaseResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *BaseResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *RelationActionRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RelationActionRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Uid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUid())
	return offset
}

func (x *RelationActionRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetToUserId())
	return offset
}

func (x *RelationActionRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetActionType())
	return offset
}

func (x *RelationActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *RelationActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetBaseResp())
	return offset
}

func (x *RelationFollowListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationFollowListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Uid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUid())
	return offset
}

func (x *RelationFollowListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetToUid())
	return offset
}

func (x *RelationFollowListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationFollowListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetBaseResp())
	return offset
}

func (x *RelationFollowListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.FollowList == nil {
		return offset
	}
	for i := range x.GetFollowList() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetFollowList()[i])
	}
	return offset
}

func (x *RelationFollowerListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationFollowerListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Uid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUid())
	return offset
}

func (x *RelationFollowerListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetToUid())
	return offset
}

func (x *RelationFollowerListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationFollowerListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetBaseResp())
	return offset
}

func (x *RelationFollowerListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.FollowerList == nil {
		return offset
	}
	for i := range x.GetFollowerList() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetFollowerList()[i])
	}
	return offset
}

func (x *RelationQueryRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationQueryRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Uid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUid())
	return offset
}

func (x *RelationQueryRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetToUserId())
	return offset
}

func (x *RelationQueryResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationQueryResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetBaseResp())
	return offset
}

func (x *RelationQueryResponse) fastWriteField2(buf []byte) (offset int) {
	if !x.IsFollow {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 2, x.GetIsFollow())
	return offset
}

func (x *QueryUserInfosWithRelationRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *QueryUserInfosWithRelationRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Uid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUid())
	return offset
}

func (x *QueryUserInfosWithRelationRequest) fastWriteField2(buf []byte) (offset int) {
	if len(x.IdList) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 2, len(x.GetIdList()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.GetIdList()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *QueryUserInfosWithRelationResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *QueryUserInfosWithRelationResponse) fastWriteField1(buf []byte) (offset int) {
	if x.UserInfoList == nil {
		return offset
	}
	for i := range x.GetUserInfoList() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUserInfoList()[i])
	}
	return offset
}

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *User) sizeField3() (n int) {
	if x.FollowCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetFollowCount())
	return n
}

func (x *User) sizeField4() (n int) {
	if x.FollowerCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetFollowerCount())
	return n
}

func (x *User) sizeField5() (n int) {
	if !x.IsFollow {
		return n
	}
	n += fastpb.SizeBool(5, x.GetIsFollow())
	return n
}

func (x *UserInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	return n
}

func (x *UserInfo) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *UserInfo) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *UserInfo) sizeField3() (n int) {
	if x.FollowCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetFollowCount())
	return n
}

func (x *UserInfo) sizeField4() (n int) {
	if x.FollowerCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetFollowerCount())
	return n
}

func (x *UserInfo) sizeField5() (n int) {
	if !x.IsFollow {
		return n
	}
	n += fastpb.SizeBool(5, x.GetIsFollow())
	return n
}

func (x *UserInfo) sizeField6() (n int) {
	if x.Avatar == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetAvatar())
	return n
}

func (x *UserInfo) sizeField7() (n int) {
	if x.BackgroundImage == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetBackgroundImage())
	return n
}

func (x *UserInfo) sizeField8() (n int) {
	if x.Signature == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetSignature())
	return n
}

func (x *UserInfo) sizeField9() (n int) {
	if x.TotalFavorited == 0 {
		return n
	}
	n += fastpb.SizeInt64(9, x.GetTotalFavorited())
	return n
}

func (x *UserInfo) sizeField10() (n int) {
	if x.WorkCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(10, x.GetWorkCount())
	return n
}

func (x *UserInfo) sizeField11() (n int) {
	if x.FavoriteCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(11, x.GetFavoriteCount())
	return n
}

func (x *BaseResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *BaseResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetStatusCode())
	return n
}

func (x *BaseResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *RelationActionRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *RelationActionRequest) sizeField1() (n int) {
	if x.Uid == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUid())
	return n
}

func (x *RelationActionRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetToUserId())
	return n
}

func (x *RelationActionRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetActionType())
	return n
}

func (x *RelationActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *RelationActionResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetBaseResp())
	return n
}

func (x *RelationFollowListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationFollowListRequest) sizeField1() (n int) {
	if x.Uid == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUid())
	return n
}

func (x *RelationFollowListRequest) sizeField2() (n int) {
	if x.ToUid == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetToUid())
	return n
}

func (x *RelationFollowListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationFollowListResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetBaseResp())
	return n
}

func (x *RelationFollowListResponse) sizeField2() (n int) {
	if x.FollowList == nil {
		return n
	}
	for i := range x.GetFollowList() {
		n += fastpb.SizeMessage(2, x.GetFollowList()[i])
	}
	return n
}

func (x *RelationFollowerListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationFollowerListRequest) sizeField1() (n int) {
	if x.Uid == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUid())
	return n
}

func (x *RelationFollowerListRequest) sizeField2() (n int) {
	if x.ToUid == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetToUid())
	return n
}

func (x *RelationFollowerListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationFollowerListResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetBaseResp())
	return n
}

func (x *RelationFollowerListResponse) sizeField2() (n int) {
	if x.FollowerList == nil {
		return n
	}
	for i := range x.GetFollowerList() {
		n += fastpb.SizeMessage(2, x.GetFollowerList()[i])
	}
	return n
}

func (x *RelationQueryRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationQueryRequest) sizeField1() (n int) {
	if x.Uid == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUid())
	return n
}

func (x *RelationQueryRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetToUserId())
	return n
}

func (x *RelationQueryResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationQueryResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetBaseResp())
	return n
}

func (x *RelationQueryResponse) sizeField2() (n int) {
	if !x.IsFollow {
		return n
	}
	n += fastpb.SizeBool(2, x.GetIsFollow())
	return n
}

func (x *QueryUserInfosWithRelationRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *QueryUserInfosWithRelationRequest) sizeField1() (n int) {
	if x.Uid == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUid())
	return n
}

func (x *QueryUserInfosWithRelationRequest) sizeField2() (n int) {
	if len(x.IdList) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(2, len(x.GetIdList()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.GetIdList()[numIdxOrVal])
			return n
		})
	return n
}

func (x *QueryUserInfosWithRelationResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *QueryUserInfosWithRelationResponse) sizeField1() (n int) {
	if x.UserInfoList == nil {
		return n
	}
	for i := range x.GetUserInfoList() {
		n += fastpb.SizeMessage(1, x.GetUserInfoList()[i])
	}
	return n
}

var fieldIDToName_User = map[int32]string{
	1: "Id",
	2: "Name",
	3: "FollowCount",
	4: "FollowerCount",
	5: "IsFollow",
}

var fieldIDToName_UserInfo = map[int32]string{
	1:  "Id",
	2:  "Name",
	3:  "FollowCount",
	4:  "FollowerCount",
	5:  "IsFollow",
	6:  "Avatar",
	7:  "BackgroundImage",
	8:  "Signature",
	9:  "TotalFavorited",
	10: "WorkCount",
	11: "FavoriteCount",
}

var fieldIDToName_BaseResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var fieldIDToName_RelationActionRequest = map[int32]string{
	1: "Uid",
	2: "ToUserId",
	3: "ActionType",
}

var fieldIDToName_RelationActionResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_RelationFollowListRequest = map[int32]string{
	1: "Uid",
	2: "ToUid",
}

var fieldIDToName_RelationFollowListResponse = map[int32]string{
	1: "BaseResp",
	2: "FollowList",
}

var fieldIDToName_RelationFollowerListRequest = map[int32]string{
	1: "Uid",
	2: "ToUid",
}

var fieldIDToName_RelationFollowerListResponse = map[int32]string{
	1: "BaseResp",
	2: "FollowerList",
}

var fieldIDToName_RelationQueryRequest = map[int32]string{
	1: "Uid",
	2: "ToUserId",
}

var fieldIDToName_RelationQueryResponse = map[int32]string{
	1: "BaseResp",
	2: "IsFollow",
}

var fieldIDToName_QueryUserInfosWithRelationRequest = map[int32]string{
	1: "Uid",
	2: "IdList",
}

var fieldIDToName_QueryUserInfosWithRelationResponse = map[int32]string{
	1: "UserInfoList",
}