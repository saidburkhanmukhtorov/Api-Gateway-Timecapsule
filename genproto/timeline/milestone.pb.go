// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: submodule-for-timecapsule/timeline_service/milestone.proto

package timeline

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Milestone represents a significant event in a user's life.
type Milestone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Date      string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"` // Use string for date (YYYY-MM-DD format)
	Category  string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Milestone) Reset() {
	*x = Milestone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Milestone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Milestone) ProtoMessage() {}

func (x *Milestone) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Milestone.ProtoReflect.Descriptor instead.
func (*Milestone) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescGZIP(), []int{0}
}

func (x *Milestone) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Milestone) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Milestone) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Milestone) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Milestone) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Milestone) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Milestone) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// GetMilestoneByIdRequest represents a request to retrieve a milestone by its ID.
type GetMilestoneByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMilestoneByIdRequest) Reset() {
	*x = GetMilestoneByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMilestoneByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMilestoneByIdRequest) ProtoMessage() {}

func (x *GetMilestoneByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMilestoneByIdRequest.ProtoReflect.Descriptor instead.
func (*GetMilestoneByIdRequest) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescGZIP(), []int{1}
}

func (x *GetMilestoneByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteMilestoneRequest represents a request to delete a milestone by its ID.
type DeleteMilestoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMilestoneRequest) Reset() {
	*x = DeleteMilestoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMilestoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMilestoneRequest) ProtoMessage() {}

func (x *DeleteMilestoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMilestoneRequest.ProtoReflect.Descriptor instead.
func (*DeleteMilestoneRequest) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteMilestoneRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteMilestoneResponse represents a response to a milestone deletion request.
type DeleteMilestoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteMilestoneResponse) Reset() {
	*x = DeleteMilestoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMilestoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMilestoneResponse) ProtoMessage() {}

func (x *DeleteMilestoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMilestoneResponse.ProtoReflect.Descriptor instead.
func (*DeleteMilestoneResponse) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteMilestoneResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// GetAllMilestonesRequest represents a request to retrieve all milestones.
type GetAllMilestonesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit     int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // Filter by user ID
	Title     string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`                          // Filter by title
	Category  string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`                    // Filter by category
	StartDate string `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"` // Filter by start date (inclusive, YYYY-MM-DD)
	EndDate   string `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`       // Filter by end date (inclusive, YYYY-MM-DD)
}

func (x *GetAllMilestonesRequest) Reset() {
	*x = GetAllMilestonesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMilestonesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMilestonesRequest) ProtoMessage() {}

func (x *GetAllMilestonesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMilestonesRequest.ProtoReflect.Descriptor instead.
func (*GetAllMilestonesRequest) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllMilestonesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllMilestonesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllMilestonesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetAllMilestonesRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetAllMilestonesRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *GetAllMilestonesRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetAllMilestonesRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

// GetAllMilestonesResponse represents a response containing a list of milestones.
type GetAllMilestonesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Milestones []*Milestone `protobuf:"bytes,1,rep,name=milestones,proto3" json:"milestones,omitempty"`
	Count      int32        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllMilestonesResponse) Reset() {
	*x = GetAllMilestonesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMilestonesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMilestonesResponse) ProtoMessage() {}

func (x *GetAllMilestonesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMilestonesResponse.ProtoReflect.Descriptor instead.
func (*GetAllMilestonesResponse) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllMilestonesResponse) GetMilestones() []*Milestone {
	if x != nil {
		return x.Milestones
	}
	return nil
}

func (x *GetAllMilestonesResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_submodule_for_timecapsule_timeline_service_milestone_proto protoreflect.FileDescriptor

var file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2d, 0x66, 0x6f, 0x72, 0x2d,
	0x74, 0x69, 0x6d, 0x65, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x69, 0x6c,
	0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x09, 0x4d, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x65, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x52, 0x0a, 0x6d, 0x69, 0x6c,
	0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x91, 0x02,
	0x0a, 0x10, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x21, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x12, 0x56,
	0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x12, 0x20, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x69, 0x6c, 0x65,
	0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d,
	0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescOnce sync.Once
	file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescData = file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDesc
)

func file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescGZIP() []byte {
	file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescOnce.Do(func() {
		file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescData = protoimpl.X.CompressGZIP(file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescData)
	})
	return file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDescData
}

var file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_submodule_for_timecapsule_timeline_service_milestone_proto_goTypes = []any{
	(*Milestone)(nil),                // 0: timeline.Milestone
	(*GetMilestoneByIdRequest)(nil),  // 1: timeline.GetMilestoneByIdRequest
	(*DeleteMilestoneRequest)(nil),   // 2: timeline.DeleteMilestoneRequest
	(*DeleteMilestoneResponse)(nil),  // 3: timeline.DeleteMilestoneResponse
	(*GetAllMilestonesRequest)(nil),  // 4: timeline.GetAllMilestonesRequest
	(*GetAllMilestonesResponse)(nil), // 5: timeline.GetAllMilestonesResponse
}
var file_submodule_for_timecapsule_timeline_service_milestone_proto_depIdxs = []int32{
	0, // 0: timeline.GetAllMilestonesResponse.milestones:type_name -> timeline.Milestone
	1, // 1: timeline.MilestoneService.GetMilestoneById:input_type -> timeline.GetMilestoneByIdRequest
	2, // 2: timeline.MilestoneService.DeleteMilestone:input_type -> timeline.DeleteMilestoneRequest
	4, // 3: timeline.MilestoneService.GetAllMilestones:input_type -> timeline.GetAllMilestonesRequest
	0, // 4: timeline.MilestoneService.GetMilestoneById:output_type -> timeline.Milestone
	3, // 5: timeline.MilestoneService.DeleteMilestone:output_type -> timeline.DeleteMilestoneResponse
	5, // 6: timeline.MilestoneService.GetAllMilestones:output_type -> timeline.GetAllMilestonesResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_submodule_for_timecapsule_timeline_service_milestone_proto_init() }
func file_submodule_for_timecapsule_timeline_service_milestone_proto_init() {
	if File_submodule_for_timecapsule_timeline_service_milestone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Milestone); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetMilestoneByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteMilestoneRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteMilestoneResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllMilestonesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllMilestonesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submodule_for_timecapsule_timeline_service_milestone_proto_goTypes,
		DependencyIndexes: file_submodule_for_timecapsule_timeline_service_milestone_proto_depIdxs,
		MessageInfos:      file_submodule_for_timecapsule_timeline_service_milestone_proto_msgTypes,
	}.Build()
	File_submodule_for_timecapsule_timeline_service_milestone_proto = out.File
	file_submodule_for_timecapsule_timeline_service_milestone_proto_rawDesc = nil
	file_submodule_for_timecapsule_timeline_service_milestone_proto_goTypes = nil
	file_submodule_for_timecapsule_timeline_service_milestone_proto_depIdxs = nil
}
