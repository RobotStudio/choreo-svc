// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: ping.proto

package choreo;

public final class PingOuterClass {
  private PingOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\nping.proto\022\006choreo\032\024primitive/bool.pro" +
      "to2,\n\004Ping\022$\n\004Ping\022\014.choreo.Bool\032\014.chore" +
      "o.Bool\"\000B+Z)github.com/RobotStudio/chore" +
      "o-svc/svc;svcb\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          choreo.BoolOuterClass.getDescriptor(),
        }, assigner);
    choreo.BoolOuterClass.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}