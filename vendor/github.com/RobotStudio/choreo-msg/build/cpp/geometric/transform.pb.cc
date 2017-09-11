// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: geometric/transform.proto

#define INTERNAL_SUPPRESS_PROTOBUF_FIELD_DEPRECATION
#include "geometric/transform.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/port.h>
#include <google/protobuf/stubs/once.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)

namespace choreo {
class TransformStampedDefaultTypeInternal : public ::google::protobuf::internal::ExplicitlyConstructed<TransformStamped> {
} _TransformStamped_default_instance_;
class TransformDefaultTypeInternal : public ::google::protobuf::internal::ExplicitlyConstructed<Transform> {
} _Transform_default_instance_;

namespace protobuf_geometric_2ftransform_2eproto {


namespace {

::google::protobuf::Metadata file_level_metadata[2];

}  // namespace

PROTOBUF_CONSTEXPR_VAR ::google::protobuf::internal::ParseTableField
    const TableStruct::entries[] = {
  {0, 0, 0, ::google::protobuf::internal::kInvalidMask, 0, 0},
};

PROTOBUF_CONSTEXPR_VAR ::google::protobuf::internal::AuxillaryParseTableField
    const TableStruct::aux[] = {
  ::google::protobuf::internal::AuxillaryParseTableField(),
};
PROTOBUF_CONSTEXPR_VAR ::google::protobuf::internal::ParseTable const
    TableStruct::schema[] = {
  { NULL, NULL, 0, -1, -1, false },
  { NULL, NULL, 0, -1, -1, false },
};

const ::google::protobuf::uint32 TableStruct::offsets[] = {
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TransformStamped, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TransformStamped, header_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TransformStamped, transform_),
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Transform, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Transform, translation_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Transform, rotation_),
};

static const ::google::protobuf::internal::MigrationSchema schemas[] = {
  { 0, -1, sizeof(TransformStamped)},
  { 7, -1, sizeof(Transform)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&_TransformStamped_default_instance_),
  reinterpret_cast<const ::google::protobuf::Message*>(&_Transform_default_instance_),
};

namespace {

void protobuf_AssignDescriptors() {
  AddDescriptors();
  ::google::protobuf::MessageFactory* factory = NULL;
  AssignDescriptors(
      "geometric/transform.proto", schemas, file_default_instances, TableStruct::offsets, factory,
      file_level_metadata, NULL, NULL);
}

void protobuf_AssignDescriptorsOnce() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &protobuf_AssignDescriptors);
}

void protobuf_RegisterTypes(const ::std::string&) GOOGLE_ATTRIBUTE_COLD;
void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::internal::RegisterAllTypes(file_level_metadata, 2);
}

}  // namespace

void TableStruct::Shutdown() {
  _TransformStamped_default_instance_.Shutdown();
  delete file_level_metadata[0].reflection;
  _Transform_default_instance_.Shutdown();
  delete file_level_metadata[1].reflection;
}

void TableStruct::InitDefaultsImpl() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  ::google::protobuf::internal::InitProtobufDefaults();
  ::choreo::protobuf_primitive_2fheader_2eproto::InitDefaults();
  ::choreo::protobuf_geometric_2fvector_2eproto::InitDefaults();
  ::choreo::protobuf_geometric_2fquaternion_2eproto::InitDefaults();
  _TransformStamped_default_instance_.DefaultConstruct();
  _Transform_default_instance_.DefaultConstruct();
  _TransformStamped_default_instance_.get_mutable()->header_ = const_cast< ::choreo::Header*>(
      ::choreo::Header::internal_default_instance());
  _TransformStamped_default_instance_.get_mutable()->transform_ = const_cast< ::choreo::Transform*>(
      ::choreo::Transform::internal_default_instance());
  _Transform_default_instance_.get_mutable()->translation_ = const_cast< ::choreo::Vector3*>(
      ::choreo::Vector3::internal_default_instance());
  _Transform_default_instance_.get_mutable()->rotation_ = const_cast< ::choreo::Quaternion*>(
      ::choreo::Quaternion::internal_default_instance());
}

void InitDefaults() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &TableStruct::InitDefaultsImpl);
}
void AddDescriptorsImpl() {
  InitDefaults();
  static const char descriptor[] = {
      "\n\031geometric/transform.proto\022\006choreo\032\026pri"
      "mitive/header.proto\032\026geometric/vector.pr"
      "oto\032\032geometric/quaternion.proto\"X\n\020Trans"
      "formStamped\022\036\n\006header\030\001 \001(\0132\016.choreo.Hea"
      "der\022$\n\ttransform\030\002 \001(\0132\021.choreo.Transfor"
      "m\"W\n\tTransform\022$\n\013translation\030\001 \001(\0132\017.ch"
      "oreo.Vector3\022$\n\010rotation\030\002 \001(\0132\022.choreo."
      "QuaternionB+Z)github.com/RobotStudio/cho"
      "reo-msg/msg;msgb\006proto3"
  };
  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
      descriptor, 343);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "geometric/transform.proto", &protobuf_RegisterTypes);
  ::choreo::protobuf_primitive_2fheader_2eproto::AddDescriptors();
  ::choreo::protobuf_geometric_2fvector_2eproto::AddDescriptors();
  ::choreo::protobuf_geometric_2fquaternion_2eproto::AddDescriptors();
  ::google::protobuf::internal::OnShutdown(&TableStruct::Shutdown);
}

void AddDescriptors() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &AddDescriptorsImpl);
}
// Force AddDescriptors() to be called at static initialization time.
struct StaticDescriptorInitializer {
  StaticDescriptorInitializer() {
    AddDescriptors();
  }
} static_descriptor_initializer;

}  // namespace protobuf_geometric_2ftransform_2eproto


// ===================================================================

#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int TransformStamped::kHeaderFieldNumber;
const int TransformStamped::kTransformFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

TransformStamped::TransformStamped()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  if (GOOGLE_PREDICT_TRUE(this != internal_default_instance())) {
    protobuf_geometric_2ftransform_2eproto::InitDefaults();
  }
  SharedCtor();
  // @@protoc_insertion_point(constructor:choreo.TransformStamped)
}
TransformStamped::TransformStamped(const TransformStamped& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL),
      _cached_size_(0) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  if (from.has_header()) {
    header_ = new ::choreo::Header(*from.header_);
  } else {
    header_ = NULL;
  }
  if (from.has_transform()) {
    transform_ = new ::choreo::Transform(*from.transform_);
  } else {
    transform_ = NULL;
  }
  // @@protoc_insertion_point(copy_constructor:choreo.TransformStamped)
}

void TransformStamped::SharedCtor() {
  ::memset(&header_, 0, reinterpret_cast<char*>(&transform_) -
    reinterpret_cast<char*>(&header_) + sizeof(transform_));
  _cached_size_ = 0;
}

TransformStamped::~TransformStamped() {
  // @@protoc_insertion_point(destructor:choreo.TransformStamped)
  SharedDtor();
}

void TransformStamped::SharedDtor() {
  if (this != internal_default_instance()) {
    delete header_;
  }
  if (this != internal_default_instance()) {
    delete transform_;
  }
}

void TransformStamped::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* TransformStamped::descriptor() {
  protobuf_geometric_2ftransform_2eproto::protobuf_AssignDescriptorsOnce();
  return protobuf_geometric_2ftransform_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const TransformStamped& TransformStamped::default_instance() {
  protobuf_geometric_2ftransform_2eproto::InitDefaults();
  return *internal_default_instance();
}

TransformStamped* TransformStamped::New(::google::protobuf::Arena* arena) const {
  TransformStamped* n = new TransformStamped;
  if (arena != NULL) {
    arena->Own(n);
  }
  return n;
}

void TransformStamped::Clear() {
// @@protoc_insertion_point(message_clear_start:choreo.TransformStamped)
  if (GetArenaNoVirtual() == NULL && header_ != NULL) {
    delete header_;
  }
  header_ = NULL;
  if (GetArenaNoVirtual() == NULL && transform_ != NULL) {
    delete transform_;
  }
  transform_ = NULL;
}

bool TransformStamped::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:choreo.TransformStamped)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // .choreo.Header header = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(10u)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessageNoVirtual(
               input, mutable_header()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // .choreo.Transform transform = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(18u)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessageNoVirtual(
               input, mutable_transform()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0 ||
            ::google::protobuf::internal::WireFormatLite::GetTagWireType(tag) ==
            ::google::protobuf::internal::WireFormatLite::WIRETYPE_END_GROUP) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormatLite::SkipField(input, tag));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:choreo.TransformStamped)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:choreo.TransformStamped)
  return false;
#undef DO_
}

void TransformStamped::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:choreo.TransformStamped)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // .choreo.Header header = 1;
  if (this->has_header()) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      1, *this->header_, output);
  }

  // .choreo.Transform transform = 2;
  if (this->has_transform()) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      2, *this->transform_, output);
  }

  // @@protoc_insertion_point(serialize_end:choreo.TransformStamped)
}

::google::protobuf::uint8* TransformStamped::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:choreo.TransformStamped)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // .choreo.Header header = 1;
  if (this->has_header()) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageNoVirtualToArray(
        1, *this->header_, deterministic, target);
  }

  // .choreo.Transform transform = 2;
  if (this->has_transform()) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageNoVirtualToArray(
        2, *this->transform_, deterministic, target);
  }

  // @@protoc_insertion_point(serialize_to_array_end:choreo.TransformStamped)
  return target;
}

size_t TransformStamped::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:choreo.TransformStamped)
  size_t total_size = 0;

  // .choreo.Header header = 1;
  if (this->has_header()) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::MessageSizeNoVirtual(
        *this->header_);
  }

  // .choreo.Transform transform = 2;
  if (this->has_transform()) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::MessageSizeNoVirtual(
        *this->transform_);
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = cached_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void TransformStamped::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:choreo.TransformStamped)
  GOOGLE_DCHECK_NE(&from, this);
  const TransformStamped* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const TransformStamped>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:choreo.TransformStamped)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:choreo.TransformStamped)
    MergeFrom(*source);
  }
}

void TransformStamped::MergeFrom(const TransformStamped& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:choreo.TransformStamped)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.has_header()) {
    mutable_header()->::choreo::Header::MergeFrom(from.header());
  }
  if (from.has_transform()) {
    mutable_transform()->::choreo::Transform::MergeFrom(from.transform());
  }
}

void TransformStamped::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:choreo.TransformStamped)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void TransformStamped::CopyFrom(const TransformStamped& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:choreo.TransformStamped)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool TransformStamped::IsInitialized() const {
  return true;
}

void TransformStamped::Swap(TransformStamped* other) {
  if (other == this) return;
  InternalSwap(other);
}
void TransformStamped::InternalSwap(TransformStamped* other) {
  std::swap(header_, other->header_);
  std::swap(transform_, other->transform_);
  std::swap(_cached_size_, other->_cached_size_);
}

::google::protobuf::Metadata TransformStamped::GetMetadata() const {
  protobuf_geometric_2ftransform_2eproto::protobuf_AssignDescriptorsOnce();
  return protobuf_geometric_2ftransform_2eproto::file_level_metadata[kIndexInFileMessages];
}

#if PROTOBUF_INLINE_NOT_IN_HEADERS
// TransformStamped

// .choreo.Header header = 1;
bool TransformStamped::has_header() const {
  return this != internal_default_instance() && header_ != NULL;
}
void TransformStamped::clear_header() {
  if (GetArenaNoVirtual() == NULL && header_ != NULL) delete header_;
  header_ = NULL;
}
const ::choreo::Header& TransformStamped::header() const {
  // @@protoc_insertion_point(field_get:choreo.TransformStamped.header)
  return header_ != NULL ? *header_
                         : *::choreo::Header::internal_default_instance();
}
::choreo::Header* TransformStamped::mutable_header() {
  
  if (header_ == NULL) {
    header_ = new ::choreo::Header;
  }
  // @@protoc_insertion_point(field_mutable:choreo.TransformStamped.header)
  return header_;
}
::choreo::Header* TransformStamped::release_header() {
  // @@protoc_insertion_point(field_release:choreo.TransformStamped.header)
  
  ::choreo::Header* temp = header_;
  header_ = NULL;
  return temp;
}
void TransformStamped::set_allocated_header(::choreo::Header* header) {
  delete header_;
  header_ = header;
  if (header) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_set_allocated:choreo.TransformStamped.header)
}

// .choreo.Transform transform = 2;
bool TransformStamped::has_transform() const {
  return this != internal_default_instance() && transform_ != NULL;
}
void TransformStamped::clear_transform() {
  if (GetArenaNoVirtual() == NULL && transform_ != NULL) delete transform_;
  transform_ = NULL;
}
const ::choreo::Transform& TransformStamped::transform() const {
  // @@protoc_insertion_point(field_get:choreo.TransformStamped.transform)
  return transform_ != NULL ? *transform_
                         : *::choreo::Transform::internal_default_instance();
}
::choreo::Transform* TransformStamped::mutable_transform() {
  
  if (transform_ == NULL) {
    transform_ = new ::choreo::Transform;
  }
  // @@protoc_insertion_point(field_mutable:choreo.TransformStamped.transform)
  return transform_;
}
::choreo::Transform* TransformStamped::release_transform() {
  // @@protoc_insertion_point(field_release:choreo.TransformStamped.transform)
  
  ::choreo::Transform* temp = transform_;
  transform_ = NULL;
  return temp;
}
void TransformStamped::set_allocated_transform(::choreo::Transform* transform) {
  delete transform_;
  transform_ = transform;
  if (transform) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_set_allocated:choreo.TransformStamped.transform)
}

#endif  // PROTOBUF_INLINE_NOT_IN_HEADERS

// ===================================================================

#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int Transform::kTranslationFieldNumber;
const int Transform::kRotationFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

Transform::Transform()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  if (GOOGLE_PREDICT_TRUE(this != internal_default_instance())) {
    protobuf_geometric_2ftransform_2eproto::InitDefaults();
  }
  SharedCtor();
  // @@protoc_insertion_point(constructor:choreo.Transform)
}
Transform::Transform(const Transform& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL),
      _cached_size_(0) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  if (from.has_translation()) {
    translation_ = new ::choreo::Vector3(*from.translation_);
  } else {
    translation_ = NULL;
  }
  if (from.has_rotation()) {
    rotation_ = new ::choreo::Quaternion(*from.rotation_);
  } else {
    rotation_ = NULL;
  }
  // @@protoc_insertion_point(copy_constructor:choreo.Transform)
}

void Transform::SharedCtor() {
  ::memset(&translation_, 0, reinterpret_cast<char*>(&rotation_) -
    reinterpret_cast<char*>(&translation_) + sizeof(rotation_));
  _cached_size_ = 0;
}

Transform::~Transform() {
  // @@protoc_insertion_point(destructor:choreo.Transform)
  SharedDtor();
}

void Transform::SharedDtor() {
  if (this != internal_default_instance()) {
    delete translation_;
  }
  if (this != internal_default_instance()) {
    delete rotation_;
  }
}

void Transform::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* Transform::descriptor() {
  protobuf_geometric_2ftransform_2eproto::protobuf_AssignDescriptorsOnce();
  return protobuf_geometric_2ftransform_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const Transform& Transform::default_instance() {
  protobuf_geometric_2ftransform_2eproto::InitDefaults();
  return *internal_default_instance();
}

Transform* Transform::New(::google::protobuf::Arena* arena) const {
  Transform* n = new Transform;
  if (arena != NULL) {
    arena->Own(n);
  }
  return n;
}

void Transform::Clear() {
// @@protoc_insertion_point(message_clear_start:choreo.Transform)
  if (GetArenaNoVirtual() == NULL && translation_ != NULL) {
    delete translation_;
  }
  translation_ = NULL;
  if (GetArenaNoVirtual() == NULL && rotation_ != NULL) {
    delete rotation_;
  }
  rotation_ = NULL;
}

bool Transform::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:choreo.Transform)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // .choreo.Vector3 translation = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(10u)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessageNoVirtual(
               input, mutable_translation()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // .choreo.Quaternion rotation = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(18u)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessageNoVirtual(
               input, mutable_rotation()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0 ||
            ::google::protobuf::internal::WireFormatLite::GetTagWireType(tag) ==
            ::google::protobuf::internal::WireFormatLite::WIRETYPE_END_GROUP) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormatLite::SkipField(input, tag));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:choreo.Transform)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:choreo.Transform)
  return false;
#undef DO_
}

void Transform::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:choreo.Transform)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // .choreo.Vector3 translation = 1;
  if (this->has_translation()) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      1, *this->translation_, output);
  }

  // .choreo.Quaternion rotation = 2;
  if (this->has_rotation()) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      2, *this->rotation_, output);
  }

  // @@protoc_insertion_point(serialize_end:choreo.Transform)
}

::google::protobuf::uint8* Transform::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:choreo.Transform)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // .choreo.Vector3 translation = 1;
  if (this->has_translation()) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageNoVirtualToArray(
        1, *this->translation_, deterministic, target);
  }

  // .choreo.Quaternion rotation = 2;
  if (this->has_rotation()) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageNoVirtualToArray(
        2, *this->rotation_, deterministic, target);
  }

  // @@protoc_insertion_point(serialize_to_array_end:choreo.Transform)
  return target;
}

size_t Transform::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:choreo.Transform)
  size_t total_size = 0;

  // .choreo.Vector3 translation = 1;
  if (this->has_translation()) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::MessageSizeNoVirtual(
        *this->translation_);
  }

  // .choreo.Quaternion rotation = 2;
  if (this->has_rotation()) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::MessageSizeNoVirtual(
        *this->rotation_);
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = cached_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void Transform::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:choreo.Transform)
  GOOGLE_DCHECK_NE(&from, this);
  const Transform* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const Transform>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:choreo.Transform)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:choreo.Transform)
    MergeFrom(*source);
  }
}

void Transform::MergeFrom(const Transform& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:choreo.Transform)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.has_translation()) {
    mutable_translation()->::choreo::Vector3::MergeFrom(from.translation());
  }
  if (from.has_rotation()) {
    mutable_rotation()->::choreo::Quaternion::MergeFrom(from.rotation());
  }
}

void Transform::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:choreo.Transform)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void Transform::CopyFrom(const Transform& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:choreo.Transform)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Transform::IsInitialized() const {
  return true;
}

void Transform::Swap(Transform* other) {
  if (other == this) return;
  InternalSwap(other);
}
void Transform::InternalSwap(Transform* other) {
  std::swap(translation_, other->translation_);
  std::swap(rotation_, other->rotation_);
  std::swap(_cached_size_, other->_cached_size_);
}

::google::protobuf::Metadata Transform::GetMetadata() const {
  protobuf_geometric_2ftransform_2eproto::protobuf_AssignDescriptorsOnce();
  return protobuf_geometric_2ftransform_2eproto::file_level_metadata[kIndexInFileMessages];
}

#if PROTOBUF_INLINE_NOT_IN_HEADERS
// Transform

// .choreo.Vector3 translation = 1;
bool Transform::has_translation() const {
  return this != internal_default_instance() && translation_ != NULL;
}
void Transform::clear_translation() {
  if (GetArenaNoVirtual() == NULL && translation_ != NULL) delete translation_;
  translation_ = NULL;
}
const ::choreo::Vector3& Transform::translation() const {
  // @@protoc_insertion_point(field_get:choreo.Transform.translation)
  return translation_ != NULL ? *translation_
                         : *::choreo::Vector3::internal_default_instance();
}
::choreo::Vector3* Transform::mutable_translation() {
  
  if (translation_ == NULL) {
    translation_ = new ::choreo::Vector3;
  }
  // @@protoc_insertion_point(field_mutable:choreo.Transform.translation)
  return translation_;
}
::choreo::Vector3* Transform::release_translation() {
  // @@protoc_insertion_point(field_release:choreo.Transform.translation)
  
  ::choreo::Vector3* temp = translation_;
  translation_ = NULL;
  return temp;
}
void Transform::set_allocated_translation(::choreo::Vector3* translation) {
  delete translation_;
  translation_ = translation;
  if (translation) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_set_allocated:choreo.Transform.translation)
}

// .choreo.Quaternion rotation = 2;
bool Transform::has_rotation() const {
  return this != internal_default_instance() && rotation_ != NULL;
}
void Transform::clear_rotation() {
  if (GetArenaNoVirtual() == NULL && rotation_ != NULL) delete rotation_;
  rotation_ = NULL;
}
const ::choreo::Quaternion& Transform::rotation() const {
  // @@protoc_insertion_point(field_get:choreo.Transform.rotation)
  return rotation_ != NULL ? *rotation_
                         : *::choreo::Quaternion::internal_default_instance();
}
::choreo::Quaternion* Transform::mutable_rotation() {
  
  if (rotation_ == NULL) {
    rotation_ = new ::choreo::Quaternion;
  }
  // @@protoc_insertion_point(field_mutable:choreo.Transform.rotation)
  return rotation_;
}
::choreo::Quaternion* Transform::release_rotation() {
  // @@protoc_insertion_point(field_release:choreo.Transform.rotation)
  
  ::choreo::Quaternion* temp = rotation_;
  rotation_ = NULL;
  return temp;
}
void Transform::set_allocated_rotation(::choreo::Quaternion* rotation) {
  delete rotation_;
  rotation_ = rotation;
  if (rotation) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_set_allocated:choreo.Transform.rotation)
}

#endif  // PROTOBUF_INLINE_NOT_IN_HEADERS

// @@protoc_insertion_point(namespace_scope)

}  // namespace choreo

// @@protoc_insertion_point(global_scope)
