///
//  Generated code. Do not modify.
//  source: service_quiz.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'rpc_create_quiz.pb.dart' as $2;

class QuizServiceApi {
  $pb.RpcClient _client;
  QuizServiceApi(this._client);

  $async.Future<$2.CreateQuizResponse> createQuiz($pb.ClientContext? ctx, $2.CreateQuizRequest request) {
    var emptyResponse = $2.CreateQuizResponse();
    return _client.invoke<$2.CreateQuizResponse>(ctx, 'QuizService', 'CreateQuiz', request, emptyResponse);
  }
}

