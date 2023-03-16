///
//  Generated code. Do not modify.
//  source: rpc_create_quiz.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use createQuizRequestDescriptor instead')
const CreateQuizRequest$json = const {
  '1': 'CreateQuizRequest',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'created_by', '3': 3, '4': 1, '5': 9, '10': 'createdBy'},
  ],
};

/// Descriptor for `CreateQuizRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createQuizRequestDescriptor = $convert.base64Decode('ChFDcmVhdGVRdWl6UmVxdWVzdBISCgRuYW1lGAEgASgJUgRuYW1lEiAKC2Rlc2NyaXB0aW9uGAIgASgJUgtkZXNjcmlwdGlvbhIdCgpjcmVhdGVkX2J5GAMgASgJUgljcmVhdGVkQnk=');
@$core.Deprecated('Use createQuizResponseDescriptor instead')
const CreateQuizResponse$json = const {
  '1': 'CreateQuizResponse',
  '2': const [
    const {'1': 'quiz', '3': 1, '4': 1, '5': 11, '6': '.pb.Quiz', '10': 'quiz'},
  ],
};

/// Descriptor for `CreateQuizResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createQuizResponseDescriptor = $convert.base64Decode('ChJDcmVhdGVRdWl6UmVzcG9uc2USHAoEcXVpehgBIAEoCzIILnBiLlF1aXpSBHF1aXo=');
