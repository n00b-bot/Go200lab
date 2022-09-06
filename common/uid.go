package common

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/btcsuite/btcutil/base58"
)

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}

func (u *UID) GetLocalID() uint32 {
	return u.localID
}

func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return UID{}, nil
	}
	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}
	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid & 0x3FFFF),
	}
	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (u *UID) MarshalJSON() ([]byte, error) {
	return []byte("\"" + u.String() + "\""), nil
}

func (u *UID) UnmarshalJSON(data []byte) error {
	decodeUID, err := DecomposeUID(string(data))
	if err != nil {
		return err
	}
	*u = decodeUID
	return nil
}
