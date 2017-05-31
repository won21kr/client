// Auto-generated by avdl-compiler v1.3.16 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/common.avdl

package chat1

import (
	"errors"
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type ThreadID []byte

func (o ThreadID) DeepCopy() ThreadID {
	return append([]byte(nil), o...)
}

type MessageID uint

func (o MessageID) DeepCopy() MessageID {
	return o.DeepCopy()
}

type TopicID []byte

func (o TopicID) DeepCopy() TopicID {
	return append([]byte(nil), o...)
}

type ConversationID []byte

func (o ConversationID) DeepCopy() ConversationID {
	return append([]byte(nil), o...)
}

type TLFID []byte

func (o TLFID) DeepCopy() TLFID {
	return append([]byte(nil), o...)
}

type Hash []byte

func (o Hash) DeepCopy() Hash {
	return append([]byte(nil), o...)
}

type InboxVers uint64

func (o InboxVers) DeepCopy() InboxVers {
	return o.DeepCopy()
}

type OutboxID []byte

func (o OutboxID) DeepCopy() OutboxID {
	return append([]byte(nil), o...)
}

type ConversationMembersType int

const (
	ConversationMembersType_KBFS    ConversationMembersType = 0
	ConversationMembersType_TEAM    ConversationMembersType = 1
	ConversationMembersType_IMPTEAM ConversationMembersType = 2
)

func (o ConversationMembersType) DeepCopy() ConversationMembersType { return o }

var ConversationMembersTypeMap = map[string]ConversationMembersType{
	"KBFS":    0,
	"TEAM":    1,
	"IMPTEAM": 2,
}

var ConversationMembersTypeRevMap = map[ConversationMembersType]string{
	0: "KBFS",
	1: "TEAM",
	2: "IMPTEAM",
}

func (e ConversationMembersType) String() string {
	if v, ok := ConversationMembersTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type MessageType int

const (
	MessageType_NONE               MessageType = 0
	MessageType_TEXT               MessageType = 1
	MessageType_ATTACHMENT         MessageType = 2
	MessageType_EDIT               MessageType = 3
	MessageType_DELETE             MessageType = 4
	MessageType_METADATA           MessageType = 5
	MessageType_TLFNAME            MessageType = 6
	MessageType_HEADLINE           MessageType = 7
	MessageType_ATTACHMENTUPLOADED MessageType = 8
)

func (o MessageType) DeepCopy() MessageType { return o }

var MessageTypeMap = map[string]MessageType{
	"NONE":               0,
	"TEXT":               1,
	"ATTACHMENT":         2,
	"EDIT":               3,
	"DELETE":             4,
	"METADATA":           5,
	"TLFNAME":            6,
	"HEADLINE":           7,
	"ATTACHMENTUPLOADED": 8,
}

var MessageTypeRevMap = map[MessageType]string{
	0: "NONE",
	1: "TEXT",
	2: "ATTACHMENT",
	3: "EDIT",
	4: "DELETE",
	5: "METADATA",
	6: "TLFNAME",
	7: "HEADLINE",
	8: "ATTACHMENTUPLOADED",
}

type TopicType int

const (
	TopicType_NONE TopicType = 0
	TopicType_CHAT TopicType = 1
	TopicType_DEV  TopicType = 2
)

func (o TopicType) DeepCopy() TopicType { return o }

var TopicTypeMap = map[string]TopicType{
	"NONE": 0,
	"CHAT": 1,
	"DEV":  2,
}

var TopicTypeRevMap = map[TopicType]string{
	0: "NONE",
	1: "CHAT",
	2: "DEV",
}

type ConversationStatus int

const (
	ConversationStatus_UNFILED  ConversationStatus = 0
	ConversationStatus_FAVORITE ConversationStatus = 1
	ConversationStatus_IGNORED  ConversationStatus = 2
	ConversationStatus_BLOCKED  ConversationStatus = 3
	ConversationStatus_MUTED    ConversationStatus = 4
	ConversationStatus_REPORTED ConversationStatus = 5
)

func (o ConversationStatus) DeepCopy() ConversationStatus { return o }

var ConversationStatusMap = map[string]ConversationStatus{
	"UNFILED":  0,
	"FAVORITE": 1,
	"IGNORED":  2,
	"BLOCKED":  3,
	"MUTED":    4,
	"REPORTED": 5,
}

var ConversationStatusRevMap = map[ConversationStatus]string{
	0: "UNFILED",
	1: "FAVORITE",
	2: "IGNORED",
	3: "BLOCKED",
	4: "MUTED",
	5: "REPORTED",
}

func (e ConversationStatus) String() string {
	if v, ok := ConversationStatusRevMap[e]; ok {
		return v
	}
	return ""
}

type Pagination struct {
	Next     []byte `codec:"next" json:"next"`
	Previous []byte `codec:"previous" json:"previous"`
	Num      int    `codec:"num" json:"num"`
	Last     bool   `codec:"last" json:"last"`
}

func (o Pagination) DeepCopy() Pagination {
	return Pagination{
		Next:     append([]byte(nil), o.Next...),
		Previous: append([]byte(nil), o.Previous...),
		Num:      o.Num,
		Last:     o.Last,
	}
}

type RateLimit struct {
	Name           string `codec:"name" json:"name"`
	CallsRemaining int    `codec:"callsRemaining" json:"callsRemaining"`
	WindowReset    int    `codec:"windowReset" json:"windowReset"`
	MaxCalls       int    `codec:"maxCalls" json:"maxCalls"`
}

func (o RateLimit) DeepCopy() RateLimit {
	return RateLimit{
		Name:           o.Name,
		CallsRemaining: o.CallsRemaining,
		WindowReset:    o.WindowReset,
		MaxCalls:       o.MaxCalls,
	}
}

type TLFVisibility int

const (
	TLFVisibility_ANY     TLFVisibility = 0
	TLFVisibility_PUBLIC  TLFVisibility = 1
	TLFVisibility_PRIVATE TLFVisibility = 2
)

func (o TLFVisibility) DeepCopy() TLFVisibility { return o }

var TLFVisibilityMap = map[string]TLFVisibility{
	"ANY":     0,
	"PUBLIC":  1,
	"PRIVATE": 2,
}

var TLFVisibilityRevMap = map[TLFVisibility]string{
	0: "ANY",
	1: "PUBLIC",
	2: "PRIVATE",
}

func (e TLFVisibility) String() string {
	if v, ok := TLFVisibilityRevMap[e]; ok {
		return v
	}
	return ""
}

type GetInboxQuery struct {
	ConvID            *ConversationID      `codec:"convID,omitempty" json:"convID,omitempty"`
	TopicType         *TopicType           `codec:"topicType,omitempty" json:"topicType,omitempty"`
	TlfID             *TLFID               `codec:"tlfID,omitempty" json:"tlfID,omitempty"`
	TlfVisibility     *TLFVisibility       `codec:"tlfVisibility,omitempty" json:"tlfVisibility,omitempty"`
	Before            *gregor1.Time        `codec:"before,omitempty" json:"before,omitempty"`
	After             *gregor1.Time        `codec:"after,omitempty" json:"after,omitempty"`
	OneChatTypePerTLF *bool                `codec:"oneChatTypePerTLF,omitempty" json:"oneChatTypePerTLF,omitempty"`
	Status            []ConversationStatus `codec:"status" json:"status"`
	ConvIDs           []ConversationID     `codec:"convIDs" json:"convIDs"`
	UnreadOnly        bool                 `codec:"unreadOnly" json:"unreadOnly"`
	ReadOnly          bool                 `codec:"readOnly" json:"readOnly"`
	ComputeActiveList bool                 `codec:"computeActiveList" json:"computeActiveList"`
	SummarizeMaxMsgs  bool                 `codec:"summarizeMaxMsgs" json:"summarizeMaxMsgs"`
}

func (o GetInboxQuery) DeepCopy() GetInboxQuery {
	return GetInboxQuery{
		ConvID: (func(x *ConversationID) *ConversationID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ConvID),
		TopicType: (func(x *TopicType) *TopicType {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TopicType),
		TlfID: (func(x *TLFID) *TLFID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TlfID),
		TlfVisibility: (func(x *TLFVisibility) *TLFVisibility {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TlfVisibility),
		Before: (func(x *gregor1.Time) *gregor1.Time {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Before),
		After: (func(x *gregor1.Time) *gregor1.Time {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.After),
		OneChatTypePerTLF: (func(x *bool) *bool {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.OneChatTypePerTLF),
		Status: (func(x []ConversationStatus) []ConversationStatus {
			var ret []ConversationStatus
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Status),
		ConvIDs: (func(x []ConversationID) []ConversationID {
			var ret []ConversationID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ConvIDs),
		UnreadOnly:        o.UnreadOnly,
		ReadOnly:          o.ReadOnly,
		ComputeActiveList: o.ComputeActiveList,
		SummarizeMaxMsgs:  o.SummarizeMaxMsgs,
	}
}

type ConversationIDTriple struct {
	Tlfid     TLFID     `codec:"tlfid" json:"tlfid"`
	TopicType TopicType `codec:"topicType" json:"topicType"`
	TopicID   TopicID   `codec:"topicID" json:"topicID"`
}

func (o ConversationIDTriple) DeepCopy() ConversationIDTriple {
	return ConversationIDTriple{
		Tlfid:     o.Tlfid.DeepCopy(),
		TopicType: o.TopicType.DeepCopy(),
		TopicID:   o.TopicID.DeepCopy(),
	}
}

type ConversationFinalizeInfo struct {
	ResetUser      string       `codec:"resetUser" json:"resetUser"`
	ResetDate      string       `codec:"resetDate" json:"resetDate"`
	ResetFull      string       `codec:"resetFull" json:"resetFull"`
	ResetTimestamp gregor1.Time `codec:"resetTimestamp" json:"resetTimestamp"`
}

func (o ConversationFinalizeInfo) DeepCopy() ConversationFinalizeInfo {
	return ConversationFinalizeInfo{
		ResetUser:      o.ResetUser,
		ResetDate:      o.ResetDate,
		ResetFull:      o.ResetFull,
		ResetTimestamp: o.ResetTimestamp.DeepCopy(),
	}
}

type ConversationResolveInfo struct {
	NewTLFName string `codec:"newTLFName" json:"newTLFName"`
}

func (o ConversationResolveInfo) DeepCopy() ConversationResolveInfo {
	return ConversationResolveInfo{
		NewTLFName: o.NewTLFName,
	}
}

type ConversationMetadata struct {
	IdTriple       ConversationIDTriple      `codec:"idTriple" json:"idTriple"`
	ConversationID ConversationID            `codec:"conversationID" json:"conversationID"`
	Visibility     TLFVisibility             `codec:"visibility" json:"visibility"`
	Status         ConversationStatus        `codec:"status" json:"status"`
	MembersType    ConversationMembersType   `codec:"membersType" json:"membersType"`
	FinalizeInfo   *ConversationFinalizeInfo `codec:"finalizeInfo,omitempty" json:"finalizeInfo,omitempty"`
	Supersedes     []ConversationMetadata    `codec:"supersedes" json:"supersedes"`
	SupersededBy   []ConversationMetadata    `codec:"supersededBy" json:"supersededBy"`
	ActiveList     []gregor1.UID             `codec:"activeList" json:"activeList"`
}

func (o ConversationMetadata) DeepCopy() ConversationMetadata {
	return ConversationMetadata{
		IdTriple:       o.IdTriple.DeepCopy(),
		ConversationID: o.ConversationID.DeepCopy(),
		Visibility:     o.Visibility.DeepCopy(),
		Status:         o.Status.DeepCopy(),
		MembersType:    o.MembersType.DeepCopy(),
		FinalizeInfo: (func(x *ConversationFinalizeInfo) *ConversationFinalizeInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.FinalizeInfo),
		Supersedes: (func(x []ConversationMetadata) []ConversationMetadata {
			var ret []ConversationMetadata
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Supersedes),
		SupersededBy: (func(x []ConversationMetadata) []ConversationMetadata {
			var ret []ConversationMetadata
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.SupersededBy),
		ActiveList: (func(x []gregor1.UID) []gregor1.UID {
			var ret []gregor1.UID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ActiveList),
	}
}

type ConversationReaderInfo struct {
	Mtime     gregor1.Time `codec:"mtime" json:"mtime"`
	ReadMsgid MessageID    `codec:"readMsgid" json:"readMsgid"`
	MaxMsgid  MessageID    `codec:"maxMsgid" json:"maxMsgid"`
}

func (o ConversationReaderInfo) DeepCopy() ConversationReaderInfo {
	return ConversationReaderInfo{
		Mtime:     o.Mtime.DeepCopy(),
		ReadMsgid: o.ReadMsgid.DeepCopy(),
		MaxMsgid:  o.MaxMsgid.DeepCopy(),
	}
}

type Conversation struct {
	Metadata        ConversationMetadata    `codec:"metadata" json:"metadata"`
	ReaderInfo      *ConversationReaderInfo `codec:"readerInfo,omitempty" json:"readerInfo,omitempty"`
	MaxMsgs         []MessageBoxed          `codec:"maxMsgs" json:"maxMsgs"`
	MaxMsgSummaries []MessageSummary        `codec:"maxMsgSummaries" json:"maxMsgSummaries"`
}

func (o Conversation) DeepCopy() Conversation {
	return Conversation{
		Metadata: o.Metadata.DeepCopy(),
		ReaderInfo: (func(x *ConversationReaderInfo) *ConversationReaderInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ReaderInfo),
		MaxMsgs: (func(x []MessageBoxed) []MessageBoxed {
			var ret []MessageBoxed
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.MaxMsgs),
		MaxMsgSummaries: (func(x []MessageSummary) []MessageSummary {
			var ret []MessageSummary
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.MaxMsgSummaries),
	}
}

type MessageSummary struct {
	MsgID       MessageID    `codec:"msgID" json:"msgID"`
	MessageType MessageType  `codec:"messageType" json:"messageType"`
	TlfName     string       `codec:"tlfName" json:"tlfName"`
	TlfPublic   bool         `codec:"tlfPublic" json:"tlfPublic"`
	Ctime       gregor1.Time `codec:"ctime" json:"ctime"`
}

func (o MessageSummary) DeepCopy() MessageSummary {
	return MessageSummary{
		MsgID:       o.MsgID.DeepCopy(),
		MessageType: o.MessageType.DeepCopy(),
		TlfName:     o.TlfName,
		TlfPublic:   o.TlfPublic,
		Ctime:       o.Ctime.DeepCopy(),
	}
}

type MessageServerHeader struct {
	MessageID    MessageID    `codec:"messageID" json:"messageID"`
	SupersededBy MessageID    `codec:"supersededBy" json:"supersededBy"`
	Ctime        gregor1.Time `codec:"ctime" json:"ctime"`
}

func (o MessageServerHeader) DeepCopy() MessageServerHeader {
	return MessageServerHeader{
		MessageID:    o.MessageID.DeepCopy(),
		SupersededBy: o.SupersededBy.DeepCopy(),
		Ctime:        o.Ctime.DeepCopy(),
	}
}

type MessagePreviousPointer struct {
	Id   MessageID `codec:"id" json:"id"`
	Hash Hash      `codec:"hash" json:"hash"`
}

func (o MessagePreviousPointer) DeepCopy() MessagePreviousPointer {
	return MessagePreviousPointer{
		Id:   o.Id.DeepCopy(),
		Hash: o.Hash.DeepCopy(),
	}
}

type OutboxInfo struct {
	Prev        MessageID    `codec:"prev" json:"prev"`
	ComposeTime gregor1.Time `codec:"composeTime" json:"composeTime"`
}

func (o OutboxInfo) DeepCopy() OutboxInfo {
	return OutboxInfo{
		Prev:        o.Prev.DeepCopy(),
		ComposeTime: o.ComposeTime.DeepCopy(),
	}
}

type MessageClientHeader struct {
	Conv         ConversationIDTriple     `codec:"conv" json:"conv"`
	TlfName      string                   `codec:"tlfName" json:"tlfName"`
	TlfPublic    bool                     `codec:"tlfPublic" json:"tlfPublic"`
	MessageType  MessageType              `codec:"messageType" json:"messageType"`
	Supersedes   MessageID                `codec:"supersedes" json:"supersedes"`
	Deletes      []MessageID              `codec:"deletes" json:"deletes"`
	Prev         []MessagePreviousPointer `codec:"prev" json:"prev"`
	Sender       gregor1.UID              `codec:"sender" json:"sender"`
	SenderDevice gregor1.DeviceID         `codec:"senderDevice" json:"senderDevice"`
	MerkleRoot   *MerkleRoot              `codec:"merkleRoot,omitempty" json:"merkleRoot,omitempty"`
	OutboxID     *OutboxID                `codec:"outboxID,omitempty" json:"outboxID,omitempty"`
	OutboxInfo   *OutboxInfo              `codec:"outboxInfo,omitempty" json:"outboxInfo,omitempty"`
}

func (o MessageClientHeader) DeepCopy() MessageClientHeader {
	return MessageClientHeader{
		Conv:        o.Conv.DeepCopy(),
		TlfName:     o.TlfName,
		TlfPublic:   o.TlfPublic,
		MessageType: o.MessageType.DeepCopy(),
		Supersedes:  o.Supersedes.DeepCopy(),
		Deletes: (func(x []MessageID) []MessageID {
			var ret []MessageID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Deletes),
		Prev: (func(x []MessagePreviousPointer) []MessagePreviousPointer {
			var ret []MessagePreviousPointer
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Prev),
		Sender:       o.Sender.DeepCopy(),
		SenderDevice: o.SenderDevice.DeepCopy(),
		MerkleRoot: (func(x *MerkleRoot) *MerkleRoot {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.MerkleRoot),
		OutboxID: (func(x *OutboxID) *OutboxID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.OutboxID),
		OutboxInfo: (func(x *OutboxInfo) *OutboxInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.OutboxInfo),
	}
}

type MessageClientHeaderVerified struct {
	Conv         ConversationIDTriple     `codec:"conv" json:"conv"`
	TlfName      string                   `codec:"tlfName" json:"tlfName"`
	TlfPublic    bool                     `codec:"tlfPublic" json:"tlfPublic"`
	MessageType  MessageType              `codec:"messageType" json:"messageType"`
	Prev         []MessagePreviousPointer `codec:"prev" json:"prev"`
	Sender       gregor1.UID              `codec:"sender" json:"sender"`
	SenderDevice gregor1.DeviceID         `codec:"senderDevice" json:"senderDevice"`
	OutboxID     *OutboxID                `codec:"outboxID,omitempty" json:"outboxID,omitempty"`
	OutboxInfo   *OutboxInfo              `codec:"outboxInfo,omitempty" json:"outboxInfo,omitempty"`
}

func (o MessageClientHeaderVerified) DeepCopy() MessageClientHeaderVerified {
	return MessageClientHeaderVerified{
		Conv:        o.Conv.DeepCopy(),
		TlfName:     o.TlfName,
		TlfPublic:   o.TlfPublic,
		MessageType: o.MessageType.DeepCopy(),
		Prev: (func(x []MessagePreviousPointer) []MessagePreviousPointer {
			var ret []MessagePreviousPointer
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Prev),
		Sender:       o.Sender.DeepCopy(),
		SenderDevice: o.SenderDevice.DeepCopy(),
		OutboxID: (func(x *OutboxID) *OutboxID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.OutboxID),
		OutboxInfo: (func(x *OutboxInfo) *OutboxInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.OutboxInfo),
	}
}

type EncryptedData struct {
	V int    `codec:"v" json:"v"`
	E []byte `codec:"e" json:"e"`
	N []byte `codec:"n" json:"n"`
}

func (o EncryptedData) DeepCopy() EncryptedData {
	return EncryptedData{
		V: o.V,
		E: append([]byte(nil), o.E...),
		N: append([]byte(nil), o.N...),
	}
}

type SignEncryptedData struct {
	V int    `codec:"v" json:"v"`
	E []byte `codec:"e" json:"e"`
	N []byte `codec:"n" json:"n"`
}

func (o SignEncryptedData) DeepCopy() SignEncryptedData {
	return SignEncryptedData{
		V: o.V,
		E: append([]byte(nil), o.E...),
		N: append([]byte(nil), o.N...),
	}
}

type SealedData struct {
	V int    `codec:"v" json:"v"`
	E []byte `codec:"e" json:"e"`
	N []byte `codec:"n" json:"n"`
}

func (o SealedData) DeepCopy() SealedData {
	return SealedData{
		V: o.V,
		E: append([]byte(nil), o.E...),
		N: append([]byte(nil), o.N...),
	}
}

type SignatureInfo struct {
	V int    `codec:"v" json:"v"`
	S []byte `codec:"s" json:"s"`
	K []byte `codec:"k" json:"k"`
}

func (o SignatureInfo) DeepCopy() SignatureInfo {
	return SignatureInfo{
		V: o.V,
		S: append([]byte(nil), o.S...),
		K: append([]byte(nil), o.K...),
	}
}

type MerkleRoot struct {
	Seqno int64  `codec:"seqno" json:"seqno"`
	Hash  []byte `codec:"hash" json:"hash"`
}

func (o MerkleRoot) DeepCopy() MerkleRoot {
	return MerkleRoot{
		Seqno: o.Seqno,
		Hash:  append([]byte(nil), o.Hash...),
	}
}

type InboxResType int

const (
	InboxResType_VERSIONHIT InboxResType = 0
	InboxResType_FULL       InboxResType = 1
)

func (o InboxResType) DeepCopy() InboxResType { return o }

var InboxResTypeMap = map[string]InboxResType{
	"VERSIONHIT": 0,
	"FULL":       1,
}

var InboxResTypeRevMap = map[InboxResType]string{
	0: "VERSIONHIT",
	1: "FULL",
}

func (e InboxResType) String() string {
	if v, ok := InboxResTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type InboxViewFull struct {
	Vers          InboxVers      `codec:"vers" json:"vers"`
	Conversations []Conversation `codec:"conversations" json:"conversations"`
	Pagination    *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

func (o InboxViewFull) DeepCopy() InboxViewFull {
	return InboxViewFull{
		Vers: o.Vers.DeepCopy(),
		Conversations: (func(x []Conversation) []Conversation {
			var ret []Conversation
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Conversations),
		Pagination: (func(x *Pagination) *Pagination {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Pagination),
	}
}

type InboxView struct {
	Rtype__ InboxResType   `codec:"rtype" json:"rtype"`
	Full__  *InboxViewFull `codec:"full,omitempty" json:"full,omitempty"`
}

func (o *InboxView) Rtype() (ret InboxResType, err error) {
	switch o.Rtype__ {
	case InboxResType_FULL:
		if o.Full__ == nil {
			err = errors.New("unexpected nil value for Full__")
			return ret, err
		}
	}
	return o.Rtype__, nil
}

func (o InboxView) Full() (res InboxViewFull) {
	if o.Rtype__ != InboxResType_FULL {
		panic("wrong case accessed")
	}
	if o.Full__ == nil {
		return
	}
	return *o.Full__
}

func NewInboxViewWithVersionhit() InboxView {
	return InboxView{
		Rtype__: InboxResType_VERSIONHIT,
	}
}

func NewInboxViewWithFull(v InboxViewFull) InboxView {
	return InboxView{
		Rtype__: InboxResType_FULL,
		Full__:  &v,
	}
}

func (o InboxView) DeepCopy() InboxView {
	return InboxView{
		Rtype__: o.Rtype__.DeepCopy(),
		Full__: (func(x *InboxViewFull) *InboxViewFull {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Full__),
	}
}

type CommonInterface interface {
}

func CommonProtocol(i CommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "chat.1.common",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type CommonClient struct {
	Cli rpc.GenericClient
}
