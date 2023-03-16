///
//  Generated code. Do not modify.
//  source: service_quiz.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import 'rpc_create_quiz.pbjson.dart' as $2;
import 'quiz.pbjson.dart' as $1;
import 'google/protobuf/timestamp.pbjson.dart' as $0;

const $core.Map<$core.String, $core.dynamic> QuizServiceBase$json = const {
  '1': 'QuizService',
  '2': const [
    const {'1': 'CreateQuiz', '2': '.pb.CreateQuizRequest', '3': '.pb.CreateQuizResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use quizServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> QuizServiceBase$messageJson = const {
  '.pb.CreateQuizRequest': $2.CreateQuizRequest$json,
  '.pb.CreateQuizResponse': $2.CreateQuizResponse$json,
  '.pb.Quiz': $1.Quiz$json,
  '.google.protobuf.Timestamp': $0.Timestamp$json,
};

/// Descriptor for `QuizService`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List quizServiceDescriptor = $convert.base64Decode('CgtRdWl6U2VydmljZRI9CgpDcmVhdGVRdWl6EhUucGIuQ3JlYXRlUXVpelJlcXVlc3QaFi5wYi5DcmVhdGVRdWl6UmVzcG9uc2UiAA==');
