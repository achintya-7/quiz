///
//  Generated code. Do not modify.
//  source: service_quiz.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'rpc_create_quiz.pb.dart' as $2;
import 'service_quiz.pbjson.dart';

export 'service_quiz.pb.dart';

abstract class QuizServiceBase extends $pb.GeneratedService {
  $async.Future<$2.CreateQuizResponse> createQuiz($pb.ServerContext ctx, $2.CreateQuizRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'CreateQuiz': return $2.CreateQuizRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'CreateQuiz': return this.createQuiz(ctx, request as $2.CreateQuizRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => QuizServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => QuizServiceBase$messageJson;
}

