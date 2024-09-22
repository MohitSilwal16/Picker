// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/file_watcher.proto

package pb

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

type CreateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath     string `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
	SessionToken string `protobuf:"bytes,2,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *CreateFileRequest) Reset() {
	*x = CreateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileRequest) ProtoMessage() {}

func (x *CreateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileRequest.ProtoReflect.Descriptor instead.
func (*CreateFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFileRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *CreateFileRequest) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type CreateFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileCreated bool `protobuf:"varint,1,opt,name=fileCreated,proto3" json:"fileCreated,omitempty"`
}

func (x *CreateFileResponse) Reset() {
	*x = CreateFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileResponse) ProtoMessage() {}

func (x *CreateFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileResponse.ProtoReflect.Descriptor instead.
func (*CreateFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFileResponse) GetFileCreated() bool {
	if x != nil {
		return x.FileCreated
	}
	return false
}

type CreateDirRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirPath      string `protobuf:"bytes,1,opt,name=dirPath,proto3" json:"dirPath,omitempty"`
	SessionToken string `protobuf:"bytes,2,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *CreateDirRequest) Reset() {
	*x = CreateDirRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirRequest) ProtoMessage() {}

func (x *CreateDirRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirRequest.ProtoReflect.Descriptor instead.
func (*CreateDirRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDirRequest) GetDirPath() string {
	if x != nil {
		return x.DirPath
	}
	return ""
}

func (x *CreateDirRequest) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type CreateDirResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirCreated bool `protobuf:"varint,1,opt,name=dirCreated,proto3" json:"dirCreated,omitempty"`
}

func (x *CreateDirResponse) Reset() {
	*x = CreateDirResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirResponse) ProtoMessage() {}

func (x *CreateDirResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirResponse.ProtoReflect.Descriptor instead.
func (*CreateDirResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDirResponse) GetDirCreated() bool {
	if x != nil {
		return x.DirCreated
	}
	return false
}

type WriteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath     string `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
	FileContent  []byte `protobuf:"bytes,2,opt,name=fileContent,proto3" json:"fileContent,omitempty"`
	SessionToken string `protobuf:"bytes,3,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *WriteFileRequest) Reset() {
	*x = WriteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileRequest) ProtoMessage() {}

func (x *WriteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileRequest.ProtoReflect.Descriptor instead.
func (*WriteFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{4}
}

func (x *WriteFileRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *WriteFileRequest) GetFileContent() []byte {
	if x != nil {
		return x.FileContent
	}
	return nil
}

func (x *WriteFileRequest) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type WriteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileWritten bool `protobuf:"varint,1,opt,name=fileWritten,proto3" json:"fileWritten,omitempty"`
}

func (x *WriteFileResponse) Reset() {
	*x = WriteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileResponse) ProtoMessage() {}

func (x *WriteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileResponse.ProtoReflect.Descriptor instead.
func (*WriteFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{5}
}

func (x *WriteFileResponse) GetFileWritten() bool {
	if x != nil {
		return x.FileWritten
	}
	return false
}

type RemoveFileDirRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileDirPath  string `protobuf:"bytes,1,opt,name=fileDirPath,proto3" json:"fileDirPath,omitempty"`
	SessionToken string `protobuf:"bytes,2,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *RemoveFileDirRequest) Reset() {
	*x = RemoveFileDirRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFileDirRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFileDirRequest) ProtoMessage() {}

func (x *RemoveFileDirRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFileDirRequest.ProtoReflect.Descriptor instead.
func (*RemoveFileDirRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveFileDirRequest) GetFileDirPath() string {
	if x != nil {
		return x.FileDirPath
	}
	return ""
}

func (x *RemoveFileDirRequest) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type RemoveFileDirResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileRemoved bool `protobuf:"varint,1,opt,name=fileRemoved,proto3" json:"fileRemoved,omitempty"`
}

func (x *RemoveFileDirResponse) Reset() {
	*x = RemoveFileDirResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFileDirResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFileDirResponse) ProtoMessage() {}

func (x *RemoveFileDirResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFileDirResponse.ProtoReflect.Descriptor instead.
func (*RemoveFileDirResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveFileDirResponse) GetFileRemoved() bool {
	if x != nil {
		return x.FileRemoved
	}
	return false
}

type RenameFileDirRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldFileDirName string `protobuf:"bytes,1,opt,name=oldFileDirName,proto3" json:"oldFileDirName,omitempty"`
	NewFileDirName string `protobuf:"bytes,2,opt,name=newFileDirName,proto3" json:"newFileDirName,omitempty"`
	SessionToken   string `protobuf:"bytes,3,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *RenameFileDirRequest) Reset() {
	*x = RenameFileDirRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameFileDirRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameFileDirRequest) ProtoMessage() {}

func (x *RenameFileDirRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameFileDirRequest.ProtoReflect.Descriptor instead.
func (*RenameFileDirRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{8}
}

func (x *RenameFileDirRequest) GetOldFileDirName() string {
	if x != nil {
		return x.OldFileDirName
	}
	return ""
}

func (x *RenameFileDirRequest) GetNewFileDirName() string {
	if x != nil {
		return x.NewFileDirName
	}
	return ""
}

func (x *RenameFileDirRequest) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type RenameFileDirResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileRenamed bool `protobuf:"varint,1,opt,name=fileRenamed,proto3" json:"fileRenamed,omitempty"`
}

func (x *RenameFileDirResponse) Reset() {
	*x = RenameFileDirResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_watcher_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameFileDirResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameFileDirResponse) ProtoMessage() {}

func (x *RenameFileDirResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_watcher_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameFileDirResponse.ProtoReflect.Descriptor instead.
func (*RenameFileDirResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_watcher_proto_rawDescGZIP(), []int{9}
}

func (x *RenameFileDirResponse) GetFileRenamed() bool {
	if x != nil {
		return x.FileRenamed
	}
	return false
}

var File_proto_file_watcher_proto protoreflect.FileDescriptor

var file_proto_file_watcher_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x53, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x36, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x66, 0x69, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x50,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x33, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x69, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x74, 0x0a, 0x10, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x35, 0x0a, 0x11, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x74,
	0x65, 0x6e, 0x22, 0x5c, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69,
	0x6c, 0x65, 0x44, 0x69, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x39, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x14,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x69, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x6c,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x6e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x15, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x64, 0x32, 0xf2, 0x02, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x12, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_file_watcher_proto_rawDescOnce sync.Once
	file_proto_file_watcher_proto_rawDescData = file_proto_file_watcher_proto_rawDesc
)

func file_proto_file_watcher_proto_rawDescGZIP() []byte {
	file_proto_file_watcher_proto_rawDescOnce.Do(func() {
		file_proto_file_watcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_file_watcher_proto_rawDescData)
	})
	return file_proto_file_watcher_proto_rawDescData
}

var file_proto_file_watcher_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_file_watcher_proto_goTypes = []any{
	(*CreateFileRequest)(nil),     // 0: proto.CreateFileRequest
	(*CreateFileResponse)(nil),    // 1: proto.CreateFileResponse
	(*CreateDirRequest)(nil),      // 2: proto.CreateDirRequest
	(*CreateDirResponse)(nil),     // 3: proto.CreateDirResponse
	(*WriteFileRequest)(nil),      // 4: proto.WriteFileRequest
	(*WriteFileResponse)(nil),     // 5: proto.WriteFileResponse
	(*RemoveFileDirRequest)(nil),  // 6: proto.RemoveFileDirRequest
	(*RemoveFileDirResponse)(nil), // 7: proto.RemoveFileDirResponse
	(*RenameFileDirRequest)(nil),  // 8: proto.RenameFileDirRequest
	(*RenameFileDirResponse)(nil), // 9: proto.RenameFileDirResponse
}
var file_proto_file_watcher_proto_depIdxs = []int32{
	0, // 0: proto.FileWatcher.CreateFile:input_type -> proto.CreateFileRequest
	2, // 1: proto.FileWatcher.CreateDir:input_type -> proto.CreateDirRequest
	4, // 2: proto.FileWatcher.WriteFile:input_type -> proto.WriteFileRequest
	6, // 3: proto.FileWatcher.RemoveFileDir:input_type -> proto.RemoveFileDirRequest
	8, // 4: proto.FileWatcher.RenameFileDir:input_type -> proto.RenameFileDirRequest
	1, // 5: proto.FileWatcher.CreateFile:output_type -> proto.CreateFileResponse
	3, // 6: proto.FileWatcher.CreateDir:output_type -> proto.CreateDirResponse
	5, // 7: proto.FileWatcher.WriteFile:output_type -> proto.WriteFileResponse
	7, // 8: proto.FileWatcher.RemoveFileDir:output_type -> proto.RemoveFileDirResponse
	9, // 9: proto.FileWatcher.RenameFileDir:output_type -> proto.RenameFileDirResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_file_watcher_proto_init() }
func file_proto_file_watcher_proto_init() {
	if File_proto_file_watcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_file_watcher_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFileRequest); i {
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
		file_proto_file_watcher_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFileResponse); i {
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
		file_proto_file_watcher_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDirRequest); i {
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
		file_proto_file_watcher_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDirResponse); i {
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
		file_proto_file_watcher_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*WriteFileRequest); i {
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
		file_proto_file_watcher_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*WriteFileResponse); i {
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
		file_proto_file_watcher_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveFileDirRequest); i {
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
		file_proto_file_watcher_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveFileDirResponse); i {
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
		file_proto_file_watcher_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RenameFileDirRequest); i {
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
		file_proto_file_watcher_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*RenameFileDirResponse); i {
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
			RawDescriptor: file_proto_file_watcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_file_watcher_proto_goTypes,
		DependencyIndexes: file_proto_file_watcher_proto_depIdxs,
		MessageInfos:      file_proto_file_watcher_proto_msgTypes,
	}.Build()
	File_proto_file_watcher_proto = out.File
	file_proto_file_watcher_proto_rawDesc = nil
	file_proto_file_watcher_proto_goTypes = nil
	file_proto_file_watcher_proto_depIdxs = nil
}
